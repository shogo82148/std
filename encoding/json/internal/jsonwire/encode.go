// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsonwire

import (
	"github.com/shogo82148/std/encoding/json/internal/jsonflags"
)

// NeedEscape reports whether src needs escaping of any characters.
// It conservatively assumes EscapeForHTML and EscapeForJS.
// It reports true for inputs with invalid UTF-8.
func NeedEscape[Bytes ~[]byte | ~string](src Bytes) bool

// AppendQuote appends src to dst as a JSON string per RFC 7159, section 7.
//
// It takes in flags and respects the following:
//   - EscapeForHTML escapes '<', '>', and '&'.
//   - EscapeForJS escapes '\u2028' and '\u2029'.
//   - AllowInvalidUTF8 avoids reporting an error for invalid UTF-8.
//
// Regardless of whether AllowInvalidUTF8 is specified,
// invalid bytes are replaced with the Unicode replacement character ('\ufffd').
// If no escape flags are set, then the shortest representable form is used,
// which is also the canonical form for strings (RFC 8785, section 3.2.2.2).
func AppendQuote[Bytes ~[]byte | ~string](dst []byte, src Bytes, flags *jsonflags.Flags) ([]byte, error)

// ReformatString consumes a JSON string from src and appends it to dst,
// reformatting it if necessary according to the specified flags.
// It returns the appended output and the number of consumed input bytes.
func ReformatString(dst, src []byte, flags *jsonflags.Flags) ([]byte, int, error)

// AppendFloat appends src to dst as a JSON number per RFC 7159, section 6.
// It formats numbers similar to the ES6 number-to-string conversion.
// See https://go.dev/issue/14135.
//
// The output is identical to ECMA-262, 6th edition, section 7.1.12.1 and with
// RFC 8785, section 3.2.2.3 for 64-bit floating-point numbers except for -0,
// which is formatted as -0 instead of just 0.
//
// For 32-bit floating-point numbers,
// the output is a 32-bit equivalent of the algorithm.
// Note that ECMA-262 specifies no algorithm for 32-bit numbers.
func AppendFloat(dst []byte, src float64, bits int) []byte

// ReformatNumber consumes a JSON string from src and appends it to dst,
// canonicalizing it if specified.
// It returns the appended output and the number of consumed input bytes.
func ReformatNumber(dst, src []byte, flags *jsonflags.Flags) ([]byte, int, error)
