// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Goroutine preemption
//
// A goroutine can be preempted at any safe-point. Currently, there
// are a few categories of safe-points:
//
// 1. A blocked safe-point occurs for the duration that a goroutine is
//    descheduled, blocked on synchronization, or in a system call.
//
// 2. Synchronous safe-points occur when a running goroutine checks
//    for a preemption request.
//
// 3. Asynchronous safe-points occur at any instruction in user code
//    where the goroutine can be safely paused and a conservative
//    stack and register scan can find stack roots. The runtime can
//    stop a goroutine at an async safe-point using a signal.
//
// At both blocked and synchronous safe-points, a goroutine's CPU
// state is minimal and the garbage collector has complete information
// about its entire stack. This makes it possible to deschedule a
// goroutine with minimal space, and to precisely scan a goroutine's
// stack.
//
// Synchronous safe-points are implemented by overloading the stack
// bound check in function prologues. To preempt a goroutine at the
// next synchronous safe-point, the runtime poisons the goroutine's
// stack bound to a value that will cause the next stack bound check
// to fail and enter the stack growth implementation, which will
// detect that it was actually a preemption and redirect to preemption
// handling.
//
// Preemption at asynchronous safe-points is implemented by suspending
// the thread using an OS mechanism (e.g., signals) and inspecting its
// state to determine if the goroutine was at an asynchronous
// safe-point. Since the thread suspension itself is generally
// asynchronous, it also checks if the running goroutine wants to be
// preempted, since this could have changed. If all conditions are
// satisfied, it adjusts the signal context to make it look like the
// signaled thread just called asyncPreempt and resumes the thread.
// asyncPreempt spills all registers and enters the scheduler.
//
// (An alternative would be to preempt in the signal handler itself.
// This would let the OS save and restore the register state and the
// runtime would only need to know how to extract potentially
// pointer-containing registers from the signal context. However, this
// would consume an M for every preempted G, and the scheduler itself
// is not designed to run from a signal handler, as it tends to
// allocate memory and start threads in the preemption path.)

package runtime
