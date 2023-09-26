// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utf8 implements functions and constants to support text encoded in
// UTF-8. It includes functions to translate between runes and UTF-8 byte sequences.
package utf8

// Numbers fundamental to the encoding.
const (
	RuneError = '\uFFFD'
	RuneSelf  = 0x80
	MaxRune   = '\U0010FFFF'
	UTFMax    = 4
)

// FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.
// An invalid encoding is considered a full Rune since it will convert as a width-1 error rune.
func FullRune(p []byte) bool

// FullRuneInString is like FullRune but its input is a string.
func FullRuneInString(s string) bool

// DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and its width in bytes.
// If the encoding is invalid, it returns (RuneError, 1), an impossible result for correct UTF-8.
func DecodeRune(p []byte) (r rune, size int)

// DecodeRuneInString is like DecodeRune but its input is a string.
// If the encoding is invalid, it returns (RuneError, 1), an impossible result for correct UTF-8.
func DecodeRuneInString(s string) (r rune, size int)

// DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and its width in bytes.
// If the encoding is invalid, it returns (RuneError, 1), an impossible result for correct UTF-8.
func DecodeLastRune(p []byte) (r rune, size int)

// DecodeLastRuneInString is like DecodeLastRune but its input is a string.
// If the encoding is invalid, it returns (RuneError, 1), an impossible result for correct UTF-8.
func DecodeLastRuneInString(s string) (r rune, size int)

// RuneLen returns the number of bytes required to encode the rune.
func RuneLen(r rune) int

// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
// It returns the number of bytes written.
func EncodeRune(p []byte, r rune) int

// RuneCount returns the number of runes in p.  Erroneous and short
// encodings are treated as single runes of width 1 byte.
func RuneCount(p []byte) int

// RuneCountInString is like RuneCount but its input is a string.
func RuneCountInString(s string) (n int)

// RuneStart reports whether the byte could be the first byte of
// an encoded rune.  Second and subsequent bytes always have the top
// two bits set to 10.
func RuneStart(b byte) bool

// Valid reports whether p consists entirely of valid UTF-8-encoded runes.
func Valid(p []byte) bool

// ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
func ValidString(s string) bool
