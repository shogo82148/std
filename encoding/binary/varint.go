// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

import (
	"github.com/shogo82148/std/io"
)

// MaxVarintLenN is the maximum length of a varint-encoded N-bit integer.
const (
	MaxVarintLen16 = 3
	MaxVarintLen32 = 5
	MaxVarintLen64 = 10
)

// PutUvarint encodes a uint64 into buf and returns the number of bytes written.
// If the buffer is too small, PutUvarint will panic.
func PutUvarint(buf []byte, x uint64) int

// Uvarint decodes a uint64 from buf and returns that value and the
// number of bytes read (> 0). If an error occurred, the value is 0
// and the number of bytes n is <= 0 meaning:
//
//		n == 0: buf too small
//		n  < 0: value larger than 64 bits (overflow)
//	             and -n is the number of bytes read
func Uvarint(buf []byte) (uint64, int)

// PutVarint encodes an int64 into buf and returns the number of bytes written.
// If the buffer is too small, PutVarint will panic.
func PutVarint(buf []byte, x int64) int

// Varint decodes an int64 from buf and returns that value and the
// number of bytes read (> 0). If an error occurred, the value is 0
// and the number of bytes n is <= 0 with the following meaning:
//
//		n == 0: buf too small
//		n  < 0: value larger than 64 bits (overflow)
//	             and -n is the number of bytes read
func Varint(buf []byte) (int64, int)

// ReadUvarint reads an encoded unsigned integer from r and returns it as a uint64.
func ReadUvarint(r io.ByteReader) (uint64, error)

// ReadVarint reads an encoded unsigned integer from r and returns it as a uint64.
func ReadVarint(r io.ByteReader) (int64, error)
