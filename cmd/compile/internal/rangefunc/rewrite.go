// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rangefunc rewrites range-over-func to code that doesn't use range-over-funcs.
Rewriting the construct in the front end, before noder, means the functions generated during
the rewrite are available in a noder-generated representation for inlining by the back end.

# Theory of Operation

The basic idea is to rewrite

	for x := range f {
		...
	}

into

	f(func(x T) bool {
		...
	})

But it's not usually that easy.

# Range variables

For a range not using :=, the assigned variables cannot be function parameters
in the generated body function. Instead, we allocate fake parameters and
start the body with an assignment. For example:

	for expr1, expr2 = range f {
		...
	}

becomes

	f(func(#p1 T1, #p2 T2) bool {
		expr1, expr2 = #p1, #p2
		...
	})

(All the generated variables have a # at the start to signal that they
are internal variables when looking at the generated code in a
debugger. Because variables have all been resolved to the specific
objects they represent, there is no danger of using plain "p1" and
colliding with a Go variable named "p1"; the # is just nice to have,
not for correctness.)

It can also happen that there are fewer range variables than function
arguments, in which case we end up with something like

	f(func(x T1, _ T2) bool {
		...
	})

or

	f(func(#p1 T1, #p2 T2, _ T3) bool {
		expr1, expr2 = #p1, #p2
		...
	})

# Return

If the body contains a "break", that break turns into "return false",
to tell f to stop. And if the body contains a "continue", that turns
into "return true", to tell f to proceed with the next value.
Those are the easy cases.

If the body contains a return or a break/continue/goto L, then we need
to rewrite that into code that breaks out of the loop and then
triggers that control flow. In general we rewrite

	for x := range f {
		...
	}

into

	{
		var #next int
		f(func(x T1) bool {
			...
			return true
		})
		... check #next ...
	}

The variable #next is an integer code that says what to do when f
returns. Each difficult statement sets #next and then returns false to
stop f.

A plain "return" rewrites to {#next = -1; return false}.
The return false breaks the loop. Then when f returns, the "check
#next" section includes

	if #next == -1 { return }

which causes the return we want.

Return with arguments is more involved, and has to deal with
corner cases involving panic, defer, and recover.  The results
of the enclosing function or closure are rewritten to give them
names if they don't have them already, and the names are assigned
at the return site.

	  func foo() (#rv1 A, #rv2 B) {

		{
			var (
				#next int
			)
			f(func(x T1) bool {
				...
				{
					// return a, b
					#rv1, #rv2 = a, b
					#next = -1
					return false
				}
				...
				return true
			})
			if #next == -1 { return }
		}

# Checking

To permit checking that an iterator is well-behaved -- that is, that
it does not call the loop body again after it has returned false or
after the entire loop has exited (it might retain a copy of the body
function, or pass it to another goroutine) -- each generated loop has
its own #stateK variable that is used to check for permitted call
patterns to the yield function for a loop body.

The state values are:

abi.RF_DONE = 0      // body of loop has exited in a non-panic way
abi.RF_READY = 1     // body of loop has not exited yet, is not running
abi.RF_PANIC = 2     // body of loop is either currently running, or has panicked
abi.RF_EXHAUSTED = 3 // iterator function call, e.g. f(func(x t){...}), returned so the sequence is "exhausted".

abi.RF_MISSING_PANIC = 4 // used to report errors.

The value of #stateK transitions
(1) before calling the iterator function,

	var #stateN = abi.RF_READY

(2) after the iterator function call returns,

	if #stateN == abi.RF_PANIC {
		panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
	}
	#stateN = abi.RF_EXHAUSTED

(3) at the beginning of the iteration of the loop body,

	if #stateN != abi.RF_READY { #stateN = abi.RF_PANIC ; runtime.panicrangestate(#stateN) }
	#stateN = abi.RF_PANIC
	// This is slightly rearranged below for better code generation.

(4) when loop iteration continues,

	#stateN = abi.RF_READY
	[return true]

(5) when control flow exits the loop body.

	#stateN = abi.RF_DONE
	[return false]

For example:

	for x := range f {
		...
		if ... { break }
		...
	}

becomes

		{
			var #state1 = abi.RF_READY
			f(func(x T1) bool {
				if #state1 != abi.RF_READY { #state1 = abi.RF_PANIC; runtime.panicrangestate(#state1) }
				#state1 = abi.RF_PANIC
				...
				if ... { #state1 = abi.RF_DONE ; return false }
				...
				#state1 = abi.RF_READY
				return true
			})
	        if #state1 == abi.RF_PANIC {
	        	// the code for the loop body did not return normally
	        	panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
	        }
			#state1 = abi.RF_EXHAUSTED
		}

# Nested Loops

So far we've only considered a single loop. If a function contains a
sequence of loops, each can be translated individually. But loops can
be nested. It would work to translate the innermost loop and then
translate the loop around it, and so on, except that there'd be a lot
of rewriting of rewritten code and the overall traversals could end up
taking time quadratic in the depth of the nesting. To avoid all that,
we use a single rewriting pass that handles a top-most range-over-func
loop and all the range-over-func loops it contains at the same time.

If we need to return from inside a doubly-nested loop, the rewrites
above stay the same, but the check after the inner loop only says

	if #next < 0 { return false }

to stop the outer loop so it can do the actual return. That is,

	for range f {
		for range g {
			...
			return a, b
			...
		}
	}

becomes

	{
		var (
			#next int
		)
		var #state1 = abi.RF_READY
		f(func() bool {
			if #state1 != abi.RF_READY { #state1 = abi.RF_PANIC; runtime.panicrangestate(#state1) }
			#state1 = abi.RF_PANIC
			var #state2 = abi.RF_READY
			g(func() bool {
				if #state2 != abi.RF_READY { #state2 = abi.RF_PANIC; runtime.panicrangestate(#state2) }
				...
				{
					// return a, b
					#rv1, #rv2 = a, b
					#next = -1
					#state2 = abi.RF_DONE
					return false
				}
				...
				#state2 = abi.RF_READY
				return true
			})
	        if #state2 == abi.RF_PANIC {
	        	panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
	        }
			#state2 = abi.RF_EXHAUSTED
			if #next < 0 {
				#state1 = abi.RF_DONE
				return false
			}
			#state1 = abi.RF_READY
			return true
		})
	    if #state1 == abi.RF_PANIC {
	       	panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
	    }
		#state1 = abi.RF_EXHAUSTED
		if #next == -1 {
			return
		}
	}

# Labeled break/continue of range-over-func loops

For a labeled break or continue of an outer range-over-func, we
use positive #next values.

Any such labeled break or continue
really means "do N breaks" or "do N breaks and 1 continue".

The positive #next value tells which level of loop N to target
with a break or continue, where perLoopStep*N means break out of
level N and perLoopStep*N-1 means continue into level N.  The
outermost loop has level 1, therefore #next == perLoopStep means
to break from the outermost loop, and #next == perLoopStep-1 means
to continue the outermost loop.

Loops that might need to propagate a labeled break or continue
add one or both of these to the #next checks:

	    // N == depth of this loop, one less than the one just exited.
		if #next != 0 {
		  if #next >= perLoopStep*N-1 { // break or continue this loop
		  	if #next >= perLoopStep*N+1 { // error checking
		  	   // TODO reason about what exactly can appear
		  	   // here given full  or partial checking.
	           runtime.panicrangestate(abi.RF_DONE)
		  	}
		  	rv := #next & 1 == 1 // code generates into #next&1
			#next = 0
			return rv
		  }
		  return false // or handle returns and gotos
		}

For example (with perLoopStep == 2)

	F: for range f { // 1, 2
		for range g { // 3, 4
			for range h {
				...
				break F
				...
				...
				continue F
				...
			}
		}
		...
	}

becomes

	{
		var #next int
		var #state1 = abi.RF_READY
		f(func() { // 1,2
			if #state1 != abi.RF_READY { #state1 = abi.RF_PANIC; runtime.panicrangestate(#state1) }
			#state1 = abi.RF_PANIC
			var #state2 = abi.RF_READY
			g(func() { // 3,4
				if #state2 != abi.RF_READY { #state2 = abi.RF_PANIC; runtime.panicrangestate(#state2) }
				#state2 = abi.RF_PANIC
				var #state3 = abi.RF_READY
				h(func() { // 5,6
					if #state3 != abi.RF_READY { #state3 = abi.RF_PANIC; runtime.panicrangestate(#state3) }
					#state3 = abi.RF_PANIC
					...
					{
						// break F
						#next = 2
						#state3 = abi.RF_DONE
						return false
					}
					...
					{
						// continue F
						#next = 1
						#state3 = abi.RF_DONE
						return false
					}
					...
					#state3 = abi.RF_READY
					return true
				})
				if #state3 == abi.RF_PANIC {
					panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
				}
				#state3 = abi.RF_EXHAUSTED
				if #next != 0 {
					// no breaks or continues targeting this loop
					#state2 = abi.RF_DONE
					return false
				}
				return true
			})
	    	if #state2 == abi.RF_PANIC {
	       		panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
	   		}
			#state2 = abi.RF_EXHAUSTED
			if #next != 0 { // just exited g, test for break/continue applied to f/F
				if #next >= 1 {
					if #next >= 3 { runtime.panicrangestate(abi.RF_DONE) } // error
					rv := #next&1 == 1
					#next = 0
					return rv
				}
				#state1 = abi.RF_DONE
				return false
			}
			...
			return true
		})
	    if #state1 == abi.RF_PANIC {
	       	panic(runtime.panicrangestate(abi.RF_MISSING_PANIC))
	    }
		#state1 = abi.RF_EXHAUSTED
	}

Note that the post-h checks only consider a break,
since no generated code tries to continue g.

# Gotos and other labeled break/continue

The final control flow translations are goto and break/continue of a
non-range-over-func statement. In both cases, we may need to break
out of one or more range-over-func loops before we can do the actual
control flow statement. Each such break/continue/goto L statement is
assigned a unique negative #next value (since -1 is return). Then
the post-checks for a given loop test for the specific codes that
refer to labels directly targetable from that block. Otherwise, the
generic

	if #next < 0 { return false }

check handles stopping the next loop to get one step closer to the label.

For example

	Top: print("start\n")
	for range f {
		for range g {
			...
			for range h {
				...
				goto Top
				...
			}
		}
	}

becomes

	Top: print("start\n")
	{
		var #next int
		var #state1 = abi.RF_READY
		f(func() {
			if #state1 != abi.RF_READY{ #state1 = abi.RF_PANIC; runtime.panicrangestate(#state1) }
			#state1 = abi.RF_PANIC
			var #state2 = abi.RF_READY
			g(func() {
				if #state2 != abi.RF_READY { #state2 = abi.RF_PANIC; runtime.panicrangestate(#state2) }
				#state2 = abi.RF_PANIC
				...
				var #state3 bool = abi.RF_READY
				h(func() {
					if #state3 != abi.RF_READY { #state3 = abi.RF_PANIC; runtime.panicrangestate(#state3) }
					#state3 = abi.RF_PANIC
					...
					{
						// goto Top
						#next = -3
						#state3 = abi.RF_DONE
						return false
					}
					...
					#state3 = abi.RF_READY
					return true
				})
				if #state3 == abi.RF_PANIC {runtime.panicrangestate(abi.RF_MISSING_PANIC)}
				#state3 = abi.RF_EXHAUSTED
				if #next < 0 {
					#state2 = abi.RF_DONE
					return false
				}
				#state2 = abi.RF_READY
				return true
			})
			if #state2 == abi.RF_PANIC {runtime.panicrangestate(abi.RF_MISSING_PANIC)}
			#state2 = abi.RF_EXHAUSTED
			if #next < 0 {
				#state1 = abi.RF_DONE
				return false
			}
			#state1 = abi.RF_READY
			return true
		})
		if #state1 == abi.RF_PANIC {runtime.panicrangestate(abi.RF_MISSING_PANIC)}
		#state1 = abi.RF_EXHAUSTED
		if #next == -3 {
			#next = 0
			goto Top
		}
	}

Labeled break/continue to non-range-over-funcs are handled the same
way as goto.

# Defers

The last wrinkle is handling defer statements. If we have

	for range f {
		defer print("A")
	}

we cannot rewrite that into

	f(func() {
		defer print("A")
	})

because the deferred code will run at the end of the iteration, not
the end of the containing function. To fix that, the runtime provides
a special hook that lets us obtain a defer "token" representing the
outer function and then use it in a later defer to attach the deferred
code to that outer function.

Normally,

	defer print("A")

compiles to

	runtime.deferproc(func() { print("A") })

This changes in a range-over-func. For example:

	for range f {
		defer print("A")
	}

compiles to

	var #defers = runtime.deferrangefunc()
	f(func() {
		runtime.deferprocat(func() { print("A") }, #defers)
	})

For this rewriting phase, we insert the explicit initialization of
#defers and then attach the #defers variable to the CallStmt
representing the defer. That variable will be propagated to the
backend and will cause the backend to compile the defer using
deferprocat instead of an ordinary deferproc.

TODO: Could call runtime.deferrangefuncend after f.
*/
package rangefunc

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

type State int

// Rewrite rewrites all the range-over-funcs in the files.
// It returns the set of function literals generated from rangefunc loop bodies.
// This allows for rangefunc loop bodies to be distingushed by debuggers.
func Rewrite(pkg *types2.Package, info *types2.Info, files []*syntax.File) map[*syntax.FuncLit]bool
