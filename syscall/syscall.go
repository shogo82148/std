// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package syscall contains an interface to the low-level operating system
// primitives. The details vary depending on the underlying system, and
// by default, godoc will display the syscall documentation for the current
// system. If you want godoc to display syscall documentation for another
// system, set $GOOS and $GOARCH to the desired system. For example, if
// you want to view documentation for freebsd/arm on linux/amd64, set $GOOS
// to freebsd and $GOARCH to arm.
// The primary use of syscall is inside other packages that provide a more
// portable interface to the system, such as "os", "time" and "net".  Use
// those packages rather than this one if you can.
// For details of the functions and data types in this package consult
// the manuals for the appropriate operating system.
// These calls return err == nil to indicate success; otherwise
// err is an operating system error describing the failure.
// On most systems, that error has type syscall.Errno.
//
// NOTE: This package is locked down. Code outside the standard
// Go repository should be migrated to use the corresponding
// package in the golang.org/x/sys repository. That is also where updates
// required by new systems or versions should be applied.
// Signal, Errno and SysProcAttr are not yet available in
// golang.org/x/sys and must still be referenced from the
// syscall package. See https://golang.org/s/go1.4-syscall
// for more information.
package syscall

// StringByteSlice converts a string to a NUL-terminated []byte,
// If s contains a NUL byte this function panics instead of
// returning an error.
//
// Deprecated: Use ByteSliceFromString instead.
func StringByteSlice(s string) []byte

// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
func ByteSliceFromString(s string) ([]byte, error)

// StringBytePtr returns a pointer to a NUL-terminated array of bytes.
// If s contains a NUL byte this function panics instead of returning
// an error.
//
// Deprecated: Use BytePtrFromString instead.
func StringBytePtr(s string) *byte

// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, EINVAL).
func BytePtrFromString(s string) (*byte, error)

// Single-word zero for use when we need a valid pointer to 0 bytes.
// See mksyscall.pl.

func (ts *Timespec) Unix() (sec int64, nsec int64)

func (tv *Timeval) Unix() (sec int64, nsec int64)

func (ts *Timespec) Nano() int64

func (tv *Timeval) Nano() int64

func Getpagesize() int
