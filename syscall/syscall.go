// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package syscall contains an interface to the low-level operating system
// primitives.  The details vary depending on the underlying system.
// Its primary use is inside other packages that provide a more portable
// interface to the system, such as "os", "time" and "net".  Use those
// packages rather than this one if you can.
// For details of the functions and data types in this package consult
// the manuals for the appropriate operating system.
// These calls return err == nil to indicate success; otherwise
// err is an operating system error describing the failure.
// On most systems, that error has type syscall.Errno.
package syscall

// StringByteSlice returns a NUL-terminated slice of bytes
// containing the text of s.
func StringByteSlice(s string) []byte

// StringBytePtr returns a pointer to a NUL-terminated array of bytes
// containing the text of s.
func StringBytePtr(s string) *byte

// Single-word zero for use when we need a valid pointer to 0 bytes.
// See mksyscall.pl.

func (ts *Timespec) Unix() (sec int64, nsec int64)

func (tv *Timeval) Unix() (sec int64, nsec int64)

func (ts *Timespec) Nano() int64

func (tv *Timeval) Nano() int64
