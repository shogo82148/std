// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exithook provides limited support for on-exit cleanup.
//
// CAREFUL! The expectation is that Add should only be called
// from a safe context (e.g. not an error/panic path or signal
// handler, preemption enabled, allocation allowed, write barriers
// allowed, etc), and that the exit function F will be invoked under
// similar circumstances. That is to say, we are expecting that F
// uses normal / high-level Go code as opposed to one of the more
// restricted dialects used for the trickier parts of the runtime.
package exithook

// A Hook is a function to be run at program termination
// (when someone invokes os.Exit, or when main.main returns).
// Hooks are run in reverse order of registration:
// the first hook added is the last one run.
type Hook struct {
	F            func()
	RunOnFailure bool
}

var (

	// runtime sets these for us
	Gosched func()
	Goid    func() uint64
	Throw   func(string)
)

// Add adds a new exit hook.
func Add(h Hook)

// Run runs the exit hooks.
//
// If an exit hook panics, Run will throw with the panic on the stack.
// If an exit hook invokes exit in the same goroutine, the goroutine will throw.
// If an exit hook invokes exit in another goroutine, that exit will block.
func Run(code int)
