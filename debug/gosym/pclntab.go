// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Line tables
 */

package gosym

type LineTable struct {
	Data []byte
	PC   uint64
	Line int
}

// TODO(rsc): Need to pull in quantum from architecture definition.

func (t *LineTable) PCToLine(pc uint64) int

func (t *LineTable) LineToPC(line int, maxpc uint64) uint64

// NewLineTable returns a new PC/line table
// corresponding to the encoded data.
// Text must be the start address of the
// corresponding text segment.
func NewLineTable(data []byte, text uint64) *LineTable
