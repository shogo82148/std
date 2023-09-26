// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Goexit terminates the goroutine that calls it. No other goroutine is affected.
// Goexit runs all deferred calls before terminating the goroutine. Because Goexit
// is not a panic, any recover calls in those deferred functions will return nil.
//
// Calling Goexit from the main goroutine terminates that goroutine
// without func main returning. Since func main has not returned,
// the program continues execution of other goroutines.
// If all other goroutines exit, the program crashes.
func Goexit()

// runningPanicDefers is non-zero while running deferred functions for panic.
// runningPanicDefers is incremented and decremented atomically.
// This is used to try hard to get a panic stack trace out when exiting.

// panicking is non-zero when crashing the program for an unrecovered panic.
// panicking is incremented and decremented atomically.

// paniclk is held while printing the panic information and stack trace,
// so that two concurrent panics don't overlap their output.

// throwReportQuirk, if non-nil, is called by throw after dumping the stacks.
//
// TODO(austin): Remove this after Go 1.15 when we remove the
// mlockGsignal workaround.
