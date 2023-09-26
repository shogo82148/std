// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64,
// checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for
// information.
package crc64

import (
	"github.com/shogo82148/std/hash"
)

// The size of a CRC-64 checksum in bytes.
const Size = 8

// Predefined polynomials.
const (
	// The ISO polynomial, defined in ISO 3309 and used in HDLC.
	ISO = 0xD800000000000000

	// The ECMA polynomial, defined in ECMA 182.
	ECMA = 0xC96C5795D7870F42
)

// Table is a 256-word table representing the polynomial for efficient processing.
type Table [256]uint64

// MakeTable returns a Table constructed from the specified polynomial.
// The contents of this Table must not be modified.
func MakeTable(poly uint64) *Table

// digest represents the partial evaluation of a checksum.

// New creates a new hash.Hash64 computing the CRC-64 checksum using the
// polynomial represented by the Table. Its Sum method will lay the
// value out in big-endian byte order. The returned Hash64 also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
func New(tab *Table) hash.Hash64

// Update returns the result of adding the bytes in p to the crc.
func Update(crc uint64, tab *Table, p []byte) uint64

// Checksum returns the CRC-64 checksum of data
// using the polynomial represented by the Table.
func Checksum(data []byte, tab *Table) uint64
