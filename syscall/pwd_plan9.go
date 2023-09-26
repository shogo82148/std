// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The working directory in Plan 9 is effectively per P, so different
// goroutines and even the same goroutine as it's rescheduled on
// different Ps can see different working directories.
//
// Instead, track a Go process-wide intent of the current working directory,
// and switch to it at important points.

package syscall

// Ensure current working directory seen by this goroutine matches
// the most recent Chdir called in any goroutine. It's called internally
// before executing any syscall which uses a relative pathname. Must
// be called with the goroutine locked to the OS thread, to prevent
// rescheduling on a different thread (potentially with a different
// working directory) before the syscall is executed.
func Fixwd()

func Getwd() (wd string, err error)

func Chdir(path string) error
