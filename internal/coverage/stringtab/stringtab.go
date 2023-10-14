// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringtab

import (
	"github.com/shogo82148/std/internal/coverage/slicereader"
	"github.com/shogo82148/std/io"
)

// Writer implements a string table writing utility.
type Writer struct {
	stab   map[string]uint32
	strs   []string
	tmp    []byte
	frozen bool
}

// InitWriter initializes a stringtab.Writer.
func (stw *Writer) InitWriter()

// Nentries returns the number of strings interned so far.
func (stw *Writer) Nentries() uint32

// Lookup looks up string 's' in the writer's table, adding
// a new entry if need be, and returning an index into the table.
func (stw *Writer) Lookup(s string) uint32

// Size computes the memory in bytes needed for the serialized
// version of a stringtab.Writer.
func (stw *Writer) Size() uint32

// Write writes the string table in serialized form to the specified
// io.Writer.
func (stw *Writer) Write(w io.Writer) error

// Freeze sends a signal to the writer that no more additions are
// allowed, only lookups of existing strings (if a lookup triggers
// addition, a panic will result). Useful as a mechanism for
// "finalizing" a string table prior to writing it out.
func (stw *Writer) Freeze()

// Reader is a helper for reading a string table previously
// serialized by a Writer.Write call.
type Reader struct {
	r    *slicereader.Reader
	strs []string
}

// NewReader creates a stringtab.Reader to read the contents
// of a string table from 'r'.
func NewReader(r *slicereader.Reader) *Reader

// Read reads/decodes a string table using the reader provided.
func (str *Reader) Read()

// Entries returns the number of decoded entries in a string table.
func (str *Reader) Entries() int

// Get returns string 'idx' within the string table.
func (str *Reader) Get(idx uint32) string
