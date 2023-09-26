// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

// The function is used to test preemption at split stack checks.
// Declaring a var avoids inlining at the call site.

// sysNanosleep is defined by OS-specific files (such as runtime_linux_test.go)
// to sleep for the given duration. If nil, dependent tests are skipped.
// The implementation should invoke a blocking system call and not
// call time.Sleep, which would deschedule the goroutine.

type Matrix [][]float64
