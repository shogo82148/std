// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utf16 implements encoding and decoding of UTF-16 sequences.
package utf16

// IsSurrogate returns true if the specified Unicode code point
// can appear in a surrogate pair.
func IsSurrogate(r rune) bool

// DecodeRune returns the UTF-16 decoding of a surrogate pair.
// If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
// the Unicode replacement code point U+FFFD.
func DecodeRune(r1, r2 rune) rune

// EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune.
// If the rune is not a valid Unicode code point or does not need encoding,
// EncodeRune returns U+FFFD, U+FFFD.
func EncodeRune(r rune) (r1, r2 rune)

// Encode returns the UTF-16 encoding of the Unicode code point sequence s.
func Encode(s []rune) []uint16

// Decode returns the Unicode code point sequence represented
// by the UTF-16 encoding s.
func Decode(s []uint16) []rune
