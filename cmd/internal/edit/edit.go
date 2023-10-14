// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package edit implements buffered position-based editing of byte slices.
package edit

// A Buffer is a queue of edits to apply to a given byte slice.
type Buffer struct {
	old []byte
	q   edits
}

// NewBuffer returns a new buffer to accumulate changes to an initial data slice.
// The returned buffer maintains a reference to the data, so the caller must ensure
// the data is not modified until after the Buffer is done being used.
func NewBuffer(data []byte) *Buffer

func (b *Buffer) Insert(pos int, new string)

func (b *Buffer) Delete(start, end int)

func (b *Buffer) Replace(start, end int, new string)

// Bytes returns a new byte slice containing the original data
// with the queued edits applied.
func (b *Buffer) Bytes() []byte

// String returns a string containing the original data
// with the queued edits applied.
func (b *Buffer) String() string
