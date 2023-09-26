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
// In Go 1.1 and earlier, each function (represented by a Func) had its own LineTable,
// and the line number corresponded to a numbering of all source lines in the
// program, across all files. That absolute line number would then have to be
// converted separately to a file name and line number within the file.
//
// In Go 1.2, the format of the data changed so that there is a single LineTable
// for the entire program, shared by all Funcs, and there are no absolute line
// numbers, just line numbers within specific files.
//
// For the most part, LineTable's methods should be treated as an internal
// detail of the package; callers should use the methods on Table instead.
type LineTable struct {
	Data []byte
	PC   uint64
	Line int

	mu       sync.Mutex
	go12     int
	binary   binary.ByteOrder
	quantum  uint32
	ptrsize  uint32
	functab  []byte
	nfunctab uint32
	filetab  []byte
	nfiletab uint32
	fileMap  map[string]uint32
	strings  map[uint32]string
}

// NOTE(rsc): This is wrong for GOARCH=arm, which uses a quantum of 4,
// but we have no idea whether we're using arm or not. This only
// matters in the old (pre-Go 1.2) symbol table format, so it's not worth
// fixing.

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
