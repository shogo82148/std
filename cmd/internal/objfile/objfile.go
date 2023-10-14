// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package objfile implements portable access to OS-specific executable files.
package objfile

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/os"
)

// A File is an opened executable file.
type File struct {
	r       *os.File
	entries []*Entry
}

type Entry struct {
	name string
	raw  rawFile
}

// A Sym is a symbol defined in an executable file.
type Sym struct {
	Name   string
	Addr   uint64
	Size   int64
	Code   rune
	Type   string
	Relocs []Reloc
}

type Reloc struct {
	Addr     uint64
	Size     uint64
	Stringer RelocStringer
}

type RelocStringer interface {
	// insnOffset is the offset of the instruction containing the relocation
	// from the start of the symbol containing the relocation.
	String(insnOffset uint64) string
}

// Open opens the named file.
// The caller must call f.Close when the file is no longer needed.
func Open(name string) (*File, error)

func (f *File) Close() error

func (f *File) Entries() []*Entry

func (f *File) Symbols() ([]Sym, error)

func (f *File) PCLineTable() (Liner, error)

func (f *File) Text() (uint64, []byte, error)

func (f *File) GOARCH() string

func (f *File) LoadAddress() (uint64, error)

func (f *File) DWARF() (*dwarf.Data, error)

func (f *File) Disasm() (*Disasm, error)

func (e *Entry) Name() string

func (e *Entry) Symbols() ([]Sym, error)

func (e *Entry) PCLineTable() (Liner, error)

func (e *Entry) Text() (uint64, []byte, error)

func (e *Entry) GOARCH() string

// LoadAddress returns the expected load address of the file.
// This differs from the actual load address for a position-independent
// executable.
func (e *Entry) LoadAddress() (uint64, error)

// DWARF returns DWARF debug data for the file, if any.
// This is for cmd/pprof to locate cgo functions.
func (e *Entry) DWARF() (*dwarf.Data, error)
