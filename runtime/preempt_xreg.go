// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64

// This provides common support for architectures that use extended register
// state in asynchronous preemption.
//
// While asynchronous preemption stores general-purpose (GP) registers on the
// preempted goroutine's own stack, extended register state can be used to save
// non-GP state off the stack. In particular, this is meant for large vector
// register files. Currently, we assume this contains only scalar data, though
// we could change this constraint by conservatively scanning this memory.
//
// For an architecture to support extended register state, it must provide a Go
// definition of an xRegState type for storing the state, and its asyncPreempt
// implementation must write this register state to p.xRegs.scratch.

package runtime
