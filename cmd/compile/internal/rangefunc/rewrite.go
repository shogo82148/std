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

Return with arguments is more involved. We need somewhere to store the
arguments while we break out of f, so we add them to the var
declaration, like:

	{
		var (
			#next int
			#r1 type1
			#r2 type2
		)
		f(func(x T1) bool {
			...
			{
				// return a, b
				#r1, #r2 = a, b
				#next = -2
				return false
			}
			...
			return true
		})
		if #next == -2 { return #r1, #r2 }
	}

TODO: What about:

	func f() (x bool) {
		for range g(&x) {
			return true
		}
	}

	func g(p *bool) func(func() bool) {
		return func(yield func() bool) {
			yield()
			// Is *p true or false here?
		}
	}

With this rewrite the "return true" is not visible after yield returns,
but maybe it should be?

# Checking

To permit checking that an iterator is well-behaved -- that is, that
it does not call the loop body again after it has returned false or
after the entire loop has exited (it might retain a copy of the body
function, or pass it to another goroutine) -- each generated loop has
its own #exitK flag that is checked before each iteration, and set both
at any early exit and after the iteration completes.

For example:

	for x := range f {
		...
		if ... { break }
		...
	}

becomes

	{
		var #exit1 bool
		f(func(x T1) bool {
			if #exit1 { runtime.panicrangeexit() }
			...
			if ... { #exit1 = true ; return false }
			...
			return true
		})
		#exit1 = true
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
			#r1 type1
			#r2 type2
		)
		var #exit1 bool
		f(func() {
			if #exit1 { runtime.panicrangeexit() }
			var #exit2 bool
			g(func() {
				if #exit2 { runtime.panicrangeexit() }
				...
				{
					// return a, b
					#r1, #r2 = a, b
					#next = -2
					#exit1, #exit2 = true, true
					return false
				}
				...
				return true
			})
			#exit2 = true
			if #next < 0 {
				return false
			}
			return true
		})
		#exit1 = true
		if #next == -2 {
			return #r1, #r2
		}
	}

Note that the #next < 0 after the inner loop handles both kinds of
return with a single check.

# Labeled break/continue of range-over-func loops

For a labeled break or continue of an outer range-over-func, we
use positive #next values. Any such labeled break or continue
really means "do N breaks" or "do N breaks and 1 continue".
We encode that as perLoopStep*N or perLoopStep*N+1 respectively.

Loops that might need to propagate a labeled break or continue
add one or both of these to the #next checks:

	if #next >= 2 {
		#next -= 2
		return false
	}

	if #next == 1 {
		#next = 0
		return true
	}

For example

	F: for range f {
		for range g {
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
		var #exit1 bool
		f(func() {
			if #exit1 { runtime.panicrangeexit() }
			var #exit2 bool
			g(func() {
				if #exit2 { runtime.panicrangeexit() }
				var #exit3 bool
				h(func() {
					if #exit3 { runtime.panicrangeexit() }
					...
					{
						// break F
						#next = 4
						#exit1, #exit2, #exit3 = true, true, true
						return false
					}
					...
					{
						// continue F
						#next = 3
						#exit2, #exit3 = true, true
						return false
					}
					...
					return true
				})
				#exit3 = true
				if #next >= 2 {
					#next -= 2
					return false
				}
				return true
			})
			#exit2 = true
			if #next >= 2 {
				#next -= 2
				return false
			}
			if #next == 1 {
				#next = 0
				return true
			}
			...
			return true
		})
		#exit1 = true
	}

Note that the post-h checks only consider a break,
since no generated code tries to continue g.

# Gotos and other labeled break/continue

The final control flow translations are goto and break/continue of a
non-range-over-func statement. In both cases, we may need to break out
of one or more range-over-func loops before we can do the actual
control flow statement. Each such break/continue/goto L statement is
assigned a unique negative #next value (below -2, since -1 and -2 are
for the two kinds of return). Then the post-checks for a given loop
test for the specific codes that refer to labels directly targetable
from that block. Otherwise, the generic

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
		var #exit1 bool
		f(func() {
			if #exit1 { runtime.panicrangeexit() }
			var #exit2 bool
			g(func() {
				if #exit2 { runtime.panicrangeexit() }
				...
				var #exit3 bool
				h(func() {
				if #exit3 { runtime.panicrangeexit() }
					...
					{
						// goto Top
						#next = -3
						#exit1, #exit2, #exit3 = true, true, true
						return false
					}
					...
					return true
				})
				#exit3 = true
				if #next < 0 {
					return false
				}
				return true
			})
			#exit2 = true
			if #next < 0 {
				return false
			}
			return true
		})
		#exit1 = true
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

// Rewrite rewrites all the range-over-funcs in the files.
func Rewrite(pkg *types2.Package, info *types2.Info, files []*syntax.File)
