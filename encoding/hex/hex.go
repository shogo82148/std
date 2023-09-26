// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hex implements hexadecimal encoding and decoding.
package hex

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// EncodedLen returns the length of an encoding of n source bytes.
func EncodedLen(n int) int

// Encode encodes src into EncodedLen(len(src))
// bytes of dst.  As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements hexadecimal encoding.
func Encode(dst, src []byte) int

// ErrLength results from decoding an odd length slice.
var ErrLength = errors.New("encoding/hex: odd length hex string")

// InvalidByteError values describe errors resulting from an invalid byte in a hex string.
type InvalidByteError byte

func (e InvalidByteError) Error() string

func DecodedLen(x int) int

// Decode decodes src into DecodedLen(len(src)) bytes, returning the actual
// number of bytes written to dst.
//
// If Decode encounters invalid input, it returns an error describing the failure.
func Decode(dst, src []byte) (int, error)

// EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string

// DecodeString returns the bytes represented by the hexadecimal string s.
func DecodeString(s string) ([]byte, error)

// Dump returns a string that contains a hex dump of the given data. The format
// of the hex dump matches the output of `hexdump -C` on the command line.
func Dump(data []byte) string

// Dumper returns a WriteCloser that writes a hex dump of all written data to
// w. The format of the dump matches the output of `hexdump -C` on the command
// line.
func Dumper(w io.Writer) io.WriteCloser
