// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binary implements simple translation between numbers and byte
// sequences and encoding and decoding of varints.
//
// Numbers are translated by reading and writing fixed-size values.
// A fixed-size value is either a fixed-size arithmetic
// type (bool, int8, uint8, int16, float32, complex64, ...)
// or an array or struct containing only fixed-size values.
//
// The varint functions encode and decode single integer values using
// a variable-length encoding; smaller values require fewer bytes.
// For a specification, see
// https://developers.google.com/protocol-buffers/docs/encoding.
//
// This package favors simplicity over efficiency. Clients that require
// high-performance serialization, especially for large data structures,
// should look at more advanced solutions such as the [encoding/gob]
// package or [google.golang.org/protobuf] for protocol buffers.
package binary

import (
	"github.com/shogo82148/std/io"
)

// A ByteOrder specifies how to convert byte slices into
// 16-, 32-, or 64-bit unsigned integers.
//
// It is implemented by [LittleEndian], [BigEndian], and [NativeEndian].
type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	String() string
}

// AppendByteOrder specifies how to append 16-, 32-, or 64-bit unsigned integers
// into a byte slice.
//
// It is implemented by [LittleEndian], [BigEndian], and [NativeEndian].
type AppendByteOrder interface {
	AppendUint16([]byte, uint16) []byte
	AppendUint32([]byte, uint32) []byte
	AppendUint64([]byte, uint64) []byte
	String() string
}

// LittleEndian is the little-endian implementation of [ByteOrder] and [AppendByteOrder].
var LittleEndian littleEndian

// BigEndian is the big-endian implementation of [ByteOrder] and [AppendByteOrder].
var BigEndian bigEndian

// Read reads structured binary data from r into data.
// Data must be a pointer to a fixed-size value or a slice
// of fixed-size values.
// Bytes read from r are decoded using the specified byte order
// and written to successive fields of the data.
// When decoding boolean values, a zero byte is decoded as false, and
// any other non-zero byte is decoded as true.
// When reading into structs, the field data for fields with
// blank (_) field names is skipped; i.e., blank field names
// may be used for padding.
// When reading into a struct, all non-blank fields must be exported
// or Read may panic.
//
// The error is [io.EOF] only if no bytes were read.
// If an [io.EOF] happens after reading some but not all the bytes,
// Read returns [io.ErrUnexpectedEOF].
func Read(r io.Reader, order ByteOrder, data any) error

// Decode binary data from buf into data according to the given byte order.
//
// Returns an error if buf is too small, otherwise the number of
// bytes consumed from buf.
func Decode(buf []byte, order ByteOrder, data any) (int, error)

// Write writes the binary representation of data into w.
// Data must be a fixed-size value or a slice of fixed-size
// values, or a pointer to such data.
// Boolean values encode as one byte: 1 for true, and 0 for false.
// Bytes written to w are encoded using the specified byte order
// and read from successive fields of the data.
// When writing structs, zero values are written for fields
// with blank (_) field names.
func Write(w io.Writer, order ByteOrder, data any) error

// Encode the binary representation of data into buf according to the given byte order.
//
// Returns an error if the buffer is too short, otherwise the number of bytes
// written into buf.
func Encode(buf []byte, order ByteOrder, data any) (int, error)

// Append the binary representation of data to buf.
//
// buf may be nil, in which case a new buffer will be allocated.
// See [Write] on which data are acceptable.
//
// Returns the (possibily extended) buffer containing data or an error.
func Append(buf []byte, order ByteOrder, data any) ([]byte, error)

// Size returns how many bytes [Write] would generate to encode the value v, which
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
// If v is neither of these, Size returns -1.
func Size(v any) int
