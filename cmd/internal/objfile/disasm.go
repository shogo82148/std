// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objfile

import (
	"github.com/shogo82148/std/container/list"
	"github.com/shogo82148/std/debug/gosym"
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/regexp"
)

// Disasm is a disassembler for a given File.
type Disasm struct {
	syms      []Sym
	pcln      Liner
	text      []byte
	textStart uint64
	textEnd   uint64
	goarch    string
	disasm    disasmFunc
	byteOrder binary.ByteOrder
}

// Disasm returns a disassembler for the file f.
func (e *Entry) Disasm() (*Disasm, error)

// CachedFile contains the content of a file split into lines.
type CachedFile struct {
	FileName string
	Lines    [][]byte
}

// FileCache is a simple LRU cache of file contents.
type FileCache struct {
	files  *list.List
	maxLen int
}

// NewFileCache returns a FileCache which can contain up to maxLen cached file contents.
func NewFileCache(maxLen int) *FileCache

// Line returns the source code line for the given file and line number.
// If the file is not already cached, reads it, inserts it into the cache,
// and removes the least recently used file if necessary.
// If the file is in cache, it is moved to the front of the list.
func (fc *FileCache) Line(filename string, line int) ([]byte, error)

// Print prints a disassembly of the file to w.
// If filter is non-nil, the disassembly only includes functions with names matching filter.
// If printCode is true, the disassembly includs corresponding source lines.
// The disassembly only includes functions that overlap the range [start, end).
func (d *Disasm) Print(w io.Writer, filter *regexp.Regexp, start, end uint64, printCode bool, gnuAsm bool)

// Decode disassembles the text segment range [start, end), calling f for each instruction.
func (d *Disasm) Decode(start, end uint64, relocs []Reloc, gnuAsm bool, f func(pc, size uint64, file string, line int, text string))

type Liner interface {
	// Given a pc, returns the corresponding file, line, and function data.
	// If unknown, returns "",0,nil.
	PCToLine(uint64) (string, int, *gosym.Func)
}
