// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements int-to-string conversion functions.

package big

import (
	"github.com/shogo82148/std/fmt"
)

// Text returns the string representation of x in the given base.
// Base must be between 2 and 36, inclusive. The result uses the
// lower-case letters 'a' to 'z' for digit values >= 10. No base
// prefix (such as "0x") is added to the string.
func (x *Int) Text(base int) string

// Append appends the string representation of x, as generated by
// x.Text(base), to buf and returns the extended buffer.
func (x *Int) Append(buf []byte, base int) []byte

func (x *Int) String() string

// Format implements fmt.Formatter. It accepts the formats
// 'b' (binary), 'o' (octal), 'd' (decimal), 'x' (lowercase
// hexadecimal), and 'X' (uppercase hexadecimal).
// Also supported are the full suite of package fmt's format
// flags for integral types, including '+' and ' ' for sign
// control, '#' for leading zero in octal and for hexadecimal,
// a leading "0x" or "0X" for "%#x" and "%#X" respectively,
// specification of minimum digits precision, output field
// width, space or zero padding, and '-' for left or right
// justification.
func (x *Int) Format(s fmt.State, ch rune)

// byteReader is a local wrapper around fmt.ScanState;
// it implements the ByteReader interface.

// Scan is a support routine for fmt.Scanner; it sets z to the value of
// the scanned number. It accepts the formats 'b' (binary), 'o' (octal),
// 'd' (decimal), 'x' (lowercase hexadecimal), and 'X' (uppercase hexadecimal).
func (z *Int) Scan(s fmt.ScanState, ch rune) error
