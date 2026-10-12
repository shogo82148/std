// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

// Package macos provides cgo-less wrappers for Core Foundation and
// Security.framework, similarly to how package syscall provides access to
// libSystem.dylib.
package macos

import (
	"github.com/shogo82148/std/time"
)

// CFRef is an opaque reference to a Core Foundation object. It is a pointer,
// but to memory not owned by Go, so not an unsafe.Pointer.
type CFRef uintptr

// CFDataToSlice returns a copy of the contents of data as a bytes slice.
func CFDataToSlice(data CFRef) []byte

// CFStringToString returns a Go string representation of the passed
// in CFString, or an empty string if it's invalid.
func CFStringToString(ref CFRef) string

// TimeToCFDateRef converts a time.Time into an apple CFDateRef.
func TimeToCFDateRef(t time.Time) CFRef

type CFString CFRef

func BytesToCFData(b []byte) CFRef

// StringToCFString returns a copy of the UTF-8 contents of s as a new CFString.
func StringToCFString(s string) CFString

func CFDictionaryGetValueIfPresent(dict CFRef, key CFString) (value CFRef, ok bool)

func CFNumberGetValue(num CFRef) (int32, error)

func CFDataGetLength(data CFRef) int

func CFDataGetBytePtr(data CFRef) uintptr

func CFArrayGetCount(array CFRef) int

func CFArrayGetValueAtIndex(array CFRef, index int) CFRef

func CFEqual(a, b CFRef) bool

func CFRelease(ref CFRef)

func CFArrayCreateMutable() CFRef

func CFArrayAppendValue(array CFRef, val CFRef)

func CFDateCreate(seconds float64) CFRef

func CFErrorCopyDescription(errRef CFRef) CFRef

func CFErrorGetCode(errRef CFRef) int

func CFStringCreateExternalRepresentation(strRef CFRef) (CFRef, error)

// ReleaseCFArray iterates through an array, releasing its contents, and then
// releases the array itself. This is necessary because we cannot, easily, set the
// CFArrayCallBacks argument when creating CFArrays.
func ReleaseCFArray(array CFRef)
