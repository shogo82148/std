// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package debug contains facilities for programs to debug themselves while
// they are running.
package debug

import (
	"github.com/shogo82148/std/os"
)

// PrintStack prints to standard error the stack trace returned by runtime.Stack.
func PrintStack()

// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls [runtime.Stack] with a large enough buffer to capture the entire trace.
func Stack() []byte

// SetCrashOutput configures a single additional file where unhandled
// panics and other fatal errors are printed, in addition to standard error.
// There is only one additional file: calling SetCrashOutput again
// overrides any earlier call; it does not close the previous file.
// SetCrashOutput(nil) disables the use of any additional file.
func SetCrashOutput(f *os.File) error
