// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwarf

import (
	"github.com/shogo82148/std/errors"
)

// A LineReader reads a sequence of LineEntry structures from a DWARF
// "line" section for a single compilation unit. LineEntries occur in
// order of increasing PC and each LineEntry gives metadata for the
// instructions from that LineEntry's PC to just before the next
// LineEntry's PC. The last entry will have its EndSequence field set.
type LineReader struct {
	buf buf

	section []byte

	version              uint16
	minInstructionLength int
	maxOpsPerInstruction int
	defaultIsStmt        bool
	lineBase             int
	lineRange            int
	opcodeBase           int
	opcodeLengths        []int
	directories          []string
	fileEntries          []*LineFile

	programOffset Offset
	endOffset     Offset

	initialFileEntries int

	state     LineEntry
	fileIndex int
}

// A LineEntry is a row in a DWARF line table.
type LineEntry struct {
	Address uint64

	OpIndex int

	File *LineFile

	Line int

	Column int

	IsStmt bool

	BasicBlock bool

	PrologueEnd bool

	EpilogueBegin bool

	ISA int

	Discriminator int

	EndSequence bool
}

// A LineFile is a source file referenced by a DWARF line table entry.
type LineFile struct {
	Name   string
	Mtime  uint64
	Length int
}

// LineReader returns a new reader for the line table of compilation
// unit cu, which must be an Entry with tag TagCompileUnit.
//
// If this compilation unit has no line table, it returns nil, nil.
func (d *Data) LineReader(cu *Entry) (*LineReader, error)

// Next sets *entry to the next row in this line table and moves to
// the next row. If there are no more entries and the line table is
// properly terminated, it returns io.EOF.
//
// Rows are always in order of increasing entry.Address, but
// entry.Line may go forward or backward.
func (r *LineReader) Next(entry *LineEntry) error

// knownOpcodeLengths gives the opcode lengths (in varint arguments)
// of known standard opcodes.

// A LineReaderPos represents a position in a line table.
type LineReaderPos struct {
	off Offset

	numFileEntries int

	state     LineEntry
	fileIndex int
}

// Tell returns the current position in the line table.
func (r *LineReader) Tell() LineReaderPos

// Seek restores the line table reader to a position returned by Tell.
//
// The argument pos must have been returned by a call to Tell on this
// line table.
func (r *LineReader) Seek(pos LineReaderPos)

// Reset repositions the line table reader at the beginning of the
// line table.
func (r *LineReader) Reset()

// ErrUnknownPC is the error returned by LineReader.ScanPC when the
// seek PC is not covered by any entry in the line table.
var ErrUnknownPC = errors.New("ErrUnknownPC")

// SeekPC sets *entry to the LineEntry that includes pc and positions
// the reader on the next entry in the line table. If necessary, this
// will seek backwards to find pc.
//
// If pc is not covered by any entry in this line table, SeekPC
// returns ErrUnknownPC. In this case, *entry and the final seek
// position are unspecified.
//
// Note that DWARF line tables only permit sequential, forward scans.
// Hence, in the worst case, this takes time linear in the size of the
// line table. If the caller wishes to do repeated fast PC lookups, it
// should build an appropriate index of the line table.
func (r *LineReader) SeekPC(pc uint64, entry *LineEntry) error
