// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

// Package jsonwire implements stateless functionality for handling JSON text.
package jsonwire

import (
	"github.com/shogo82148/std/errors"
)

// TrimSuffixWhitespace trims JSON from the end of b.
func TrimSuffixWhitespace(b []byte) []byte

// TrimSuffixString trims a valid JSON string at the end of b.
// The behavior is undefined if there is not a valid JSON string present.
func TrimSuffixString(b []byte) []byte

// HasSuffixByte reports whether b ends with c.
func HasSuffixByte(b []byte, c byte) bool

// TrimSuffixByte removes c from the end of b if it is present.
func TrimSuffixByte(b []byte, c byte) []byte

// QuoteRune quotes the first rune in the input.
func QuoteRune[Bytes ~[]byte | ~string](b Bytes) string

// CompareUTF16 lexicographically compares x to y according
// to the UTF-16 codepoints of the UTF-8 encoded input strings.
// This implements the ordering specified in RFC 8785, section 3.2.3.
func CompareUTF16[Bytes ~[]byte | ~string](x, y Bytes) int

// TODO(https://go.dev/issue/70547): Use utf8.ErrInvalid instead.
var ErrInvalidUTF8 = errors.New("invalid UTF-8")

func NewInvalidCharacterError[Bytes ~[]byte | ~string](prefix Bytes, where string) error

func NewInvalidEscapeSequenceError[Bytes ~[]byte | ~string](what Bytes) error

// TruncatePointer optionally truncates the JSON pointer,
// enforcing that the length roughly does not exceed n.
func TruncatePointer(s string, n int) string
