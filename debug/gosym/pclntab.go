// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Line tables
 */

package gosym

import (
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/sync"
)

// A LineTable is a data structure mapping program counters to line numbers.
//
// In Go 1.1 and earlier, each function (represented by a [Func]) had its own LineTable,
// and the line number corresponded to a numbering of all source lines in the
// program, across all files. That absolute line number would then have to be
// converted separately to a file name and line number within the file.
//
// In Go 1.2, the format of the data changed so that there is a single LineTable
// for the entire program, shared by all Funcs, and there are no absolute line
// numbers, just line numbers within specific files.
//
// For the most part, LineTable's methods should be treated as an internal
// detail of the package; callers should use the methods on [Table] instead.
type LineTable struct {
	Data []byte
	PC   uint64
	Line int

	// This mutex is used to keep parsing of pclntab synchronous.
	mu sync.Mutex

	// Contains the version of the pclntab section.
	version version

	// Go 1.2/1.16/1.18 state
	binary      binary.ByteOrder
	quantum     uint32
	ptrsize     uint32
	textStart   uint64
	funcnametab []byte
	cutab       []byte
	funcdata    []byte
	functab     []byte
	nfunctab    uint32
	filetab     []byte
	pctab       []byte
	nfiletab    uint32
	funcNames   map[uint32]string
	strings     map[uint32]string
	// fileMap varies depending on the version of the object file.
	// For ver12, it maps the name to the index in the file table.
	// For ver116, it maps the name to the offset in filetab.
	fileMap map[string]uint32
}

// PCToLine returns the line number for the given program counter.
//
// Deprecated: Use Table's PCToLine method instead.
func (t *LineTable) PCToLine(pc uint64) int

// LineToPC returns the program counter for the given line number,
// considering only program counters before maxpc.
//
// Deprecated: Use Table's LineToPC method instead.
func (t *LineTable) LineToPC(line int, maxpc uint64) uint64

// NewLineTable returns a new PC/line table
// corresponding to the encoded data.
// Text must be the start address of the
// corresponding text segment.
func NewLineTable(data []byte, text uint64) *LineTable
