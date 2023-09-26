// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// GOMAXPROCS sets the maximum number of CPUs that can be executing
// simultaneously and returns the previous setting.  If n < 1, it does not
// change the current setting.
// The number of logical CPUs on the local machine can be queried with NumCPU.
// This call will go away when the scheduler improves.
func GOMAXPROCS(n int) int

// NumCPU returns the number of logical CPUs usable by the current process.
func NumCPU() int

// NumCgoCall returns the number of cgo calls made by the current process.
func NumCgoCall() int64

// NumGoroutine returns the number of goroutines that currently exist.
func NumGoroutine() int
