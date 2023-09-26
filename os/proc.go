// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Process etc.

package os

// Args hold the command-line arguments, starting with the program name.
var Args []string

// Getuid returns the numeric user id of the caller.
func Getuid() int

// Geteuid returns the numeric effective user id of the caller.
func Geteuid() int

// Getgid returns the numeric group id of the caller.
func Getgid() int

// Getegid returns the numeric effective group id of the caller.
func Getegid() int

// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
func Getgroups() ([]int, error)

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
func Exit(code int)
