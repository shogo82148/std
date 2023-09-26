// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Process etc.

package os

// Args hold the command-line arguments, starting with the program name.
var Args []string

// Getuid returns the numeric user id of the caller.
//
// On Windows, it returns -1.
func Getuid() int

// Geteuid returns the numeric effective user id of the caller.
//
// On Windows, it returns -1.
func Geteuid() int

// Getgid returns the numeric group id of the caller.
//
// On Windows, it returns -1.
func Getgid() int

// Getegid returns the numeric effective group id of the caller.
//
// On Windows, it returns -1.
func Getegid() int

// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
//
// On Windows, it returns syscall.EWINDOWS. See the os/user package
// for a possible alternative.
func Getgroups() ([]int, error)

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
//
// For portability, the status code should be in the range [0, 125].
func Exit(code int)
