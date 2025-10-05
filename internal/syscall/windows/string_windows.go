// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

// NTUnicodeString is a UTF-16 string for NT native APIs, corresponding to UNICODE_STRING.
type NTUnicodeString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

// NewNTUnicodeString returns a new NTUnicodeString structure for use with native
// NT APIs that work over the NTUnicodeString type. Note that most Windows APIs
// do not use NTUnicodeString, and instead UTF16PtrFromString should be used for
// the more common *uint16 string type.
func NewNTUnicodeString(s string) (*NTUnicodeString, error)
