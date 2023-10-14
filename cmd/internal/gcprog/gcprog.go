// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gcprog implements an encoder for packed GC pointer bitmaps,
// known as GC programs.
//
// # Program Format
//
// The GC program encodes a sequence of 0 and 1 bits indicating scalar or pointer words in an object.
// The encoding is a simple Lempel-Ziv program, with codes to emit literal bits and to repeat the
// last n bits c times.
//
// The possible codes are:
//
//	00000000: stop
//	0nnnnnnn: emit n bits copied from the next (n+7)/8 bytes, least significant bit first
//	10000000 n c: repeat the previous n bits c times; n, c are varints
//	1nnnnnnn c: repeat the previous n bits c times; c is a varint
//
// The numbers n and c, when they follow a code, are encoded as varints
// using the same encoding as encoding/binary's Uvarint.
package gcprog

import (
	"github.com/shogo82148/std/io"
)

// A Writer is an encoder for GC programs.
//
// The typical use of a Writer is to call Init, maybe call Debug,
// make a sequence of Ptr, Advance, Repeat, and Append calls
// to describe the data type, and then finally call End.
type Writer struct {
	writeByte func(byte)
	index     int64
	b         [progMaxLiteral]byte
	nb        int
	debug     io.Writer
	debugBuf  []byte
}

// Init initializes w to write a new GC program
// by calling writeByte for each byte in the program.
func (w *Writer) Init(writeByte func(byte))

// Debug causes the writer to print a debugging trace to out
// during future calls to methods like Ptr, Advance, and End.
// It also enables debugging checks during the encoding.
func (w *Writer) Debug(out io.Writer)

// BitIndex returns the number of bits written to the bit stream so far.
func (w *Writer) BitIndex() int64

// End marks the end of the program, writing any remaining bytes.
func (w *Writer) End()

// Ptr emits a 1 into the bit stream at the given bit index.
// that is, it records that the index'th word in the object memory is a pointer.
// Any bits between the current index and the new index
// are set to zero, meaning the corresponding words are scalars.
func (w *Writer) Ptr(index int64)

// ShouldRepeat reports whether it would be worthwhile to
// use a Repeat to describe c elements of n bits each,
// compared to just emitting c copies of the n-bit description.
func (w *Writer) ShouldRepeat(n, c int64) bool

// Repeat emits an instruction to repeat the description
// of the last n words c times (including the initial description, c+1 times in total).
func (w *Writer) Repeat(n, c int64)

// ZeroUntil adds zeros to the bit stream until reaching the given index;
// that is, it records that the words from the most recent pointer until
// the index'th word are scalars.
// ZeroUntil is usually called in preparation for a call to Repeat, Append, or End.
func (w *Writer) ZeroUntil(index int64)

// Append emits the given GC program into the current output.
// The caller asserts that the program emits n bits (describes n words),
// and Append panics if that is not true.
func (w *Writer) Append(prog []byte, n int64)
