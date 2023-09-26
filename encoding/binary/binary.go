// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binary implements translation between numbers and byte sequences
// and encoding and decoding of varints.
//
// Numbers are translated by reading and writing fixed-size values.
// A fixed-size value is either a fixed-size arithmetic
// type (int8, uint8, int16, float32, complex64, ...)
// or an array or struct containing only fixed-size values.
//
// Varints are a method of encoding integers using one or more bytes;
// numbers with smaller absolute value take a smaller number of bytes.
// For a specification, see http://code.google.com/apis/protocolbuffers/docs/encoding.html.
package binary

import (
	"github.com/shogo82148/std/io"
)

// A ByteOrder specifies how to convert byte sequences into
// 16-, 32-, or 64-bit unsigned integers.
type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	String() string
}

// LittleEndian is the little-endian implementation of ByteOrder.
var LittleEndian littleEndian

// BigEndian is the big-endian implementation of ByteOrder.
var BigEndian bigEndian

// Read reads structured binary data from r into data.
// Data must be a pointer to a fixed-size value or a slice
// of fixed-size values.
// Bytes read from r are decoded using the specified byte order
// and written to successive fields of the data.
func Read(r io.Reader, order ByteOrder, data interface{}) error

// Write writes the binary representation of data into w.
// Data must be a fixed-size value or a slice of fixed-size
// values, or a pointer to such data.
// Bytes written to w are encoded using the specified byte order
// and read from successive fields of the data.
func Write(w io.Writer, order ByteOrder, data interface{}) error

// Size returns how many bytes Write would generate to encode the value v, which
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
func Size(v interface{}) int
