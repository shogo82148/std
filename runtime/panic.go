// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// throwType indicates the current type of ongoing throw, which affects the
// amount of detail printed to stderr. Higher values include more detail.

// Goexit terminates the goroutine that calls it. No other goroutine is affected.
// Goexit runs all deferred calls before terminating the goroutine. Because Goexit
// is not a panic, any recover calls in those deferred functions will return nil.
//
// Calling Goexit from the main goroutine terminates that goroutine
// without func main returning. Since func main has not returned,
// the program continues execution of other goroutines.
// If all other goroutines exit, the program crashes.
func Goexit()

// A PanicNilError happens when code calls panic(nil).
//
// Before Go 1.21, programs that called panic(nil) observed recover returning nil.
// Starting in Go 1.21, programs that call panic(nil) observe recover returning a *PanicNilError.
// Programs can change back to the old behavior by setting GODEBUG=panicnil=1.
type PanicNilError struct {
	_ [0]*PanicNilError
}

func (*PanicNilError) Error() string
func (*PanicNilError) RuntimeError()

// runningPanicDefers is non-zero while running deferred functions for panic.
// This is used to try hard to get a panic stack trace out when exiting.

// panicking is non-zero when crashing the program for an unrecovered panic.

// paniclk is held while printing the panic information and stack trace,
// so that two concurrent panics don't overlap their output.
