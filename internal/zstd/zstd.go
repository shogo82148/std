// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package zstd provides a decompressor for zstd streams,
// described in RFC 8878. It does not support dictionaries.
package zstd

import (
	"github.com/shogo82148/std/io"
)

// Reader implements [io.Reader] to read a zstd compressed stream.
type Reader struct {
	// The underlying Reader.
	r io.Reader

	// Whether we have read the frame header.
	// This is of interest when buffer is empty.
	// If true we expect to see a new block.
	sawFrameHeader bool

	// Whether the current frame expects a checksum.
	hasChecksum bool

	// Whether we have read at least one frame.
	readOneFrame bool

	// True if the frame size is not known.
	frameSizeUnknown bool

	// The number of uncompressed bytes remaining in the current frame.
	// If frameSizeUnknown is true, this is not valid.
	remainingFrameSize uint64

	// The number of bytes read from r up to the start of the current
	// block, for error reporting.
	blockOffset int64

	// Buffered decompressed data.
	buffer []byte
	// Current read offset in buffer.
	off int

	// The current repeated offsets.
	repeatedOffset1 uint32
	repeatedOffset2 uint32
	repeatedOffset3 uint32

	// The current Huffman tree used for compressing literals.
	huffmanTable     []uint16
	huffmanTableBits int

	// The window for back references.
	windowSize int
	window     []byte

	// A buffer available to hold a compressed block.
	compressedBuf []byte

	// A buffer for literals.
	literals []byte

	// Sequence decode FSE tables.
	seqTables    [3][]fseBaselineEntry
	seqTableBits [3]uint8

	// Buffers for sequence decode FSE tables.
	seqTableBuffers [3][]fseBaselineEntry

	// Scratch space used for small reads, to avoid allocation.
	scratch [16]byte

	// A scratch table for reading an FSE. Only temporarily valid.
	fseScratch []fseEntry

	// For checksum computation.
	checksum xxhash64
}

// NewReader creates a new Reader that decompresses data from the given reader.
func NewReader(input io.Reader) *Reader

// Reset discards the current state and starts reading a new stream from r.
// This permits reusing a Reader rather than allocating a new one.
func (r *Reader) Reset(input io.Reader)

// Read implements [io.Reader].
func (r *Reader) Read(p []byte) (int, error)

// ReadByte implements [io.ByteReader].
func (r *Reader) ReadByte() (byte, error)
