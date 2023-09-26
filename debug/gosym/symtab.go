// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gosym implements access to the Go symbol
// and line number tables embedded in Go binaries generated
// by the gc compilers.
package gosym

// A Sym represents a single symbol table entry.
type Sym struct {
	Value  uint64
	Type   byte
	Name   string
	GoType uint64

	Func *Func
}

// Static reports whether this symbol is static (not visible outside its file).
func (s *Sym) Static() bool

// PackageName returns the package part of the symbol name,
// or the empty string if there is none.
func (s *Sym) PackageName() string

// ReceiverName returns the receiver type name of this symbol,
// or the empty string if there is none.
func (s *Sym) ReceiverName() string

// BaseName returns the symbol name without the package or receiver name.
func (s *Sym) BaseName() string

// A Func collects information about a single function.
type Func struct {
	Entry uint64
	*Sym
	End       uint64
	Params    []*Sym
	Locals    []*Sym
	FrameSize int
	LineTable *LineTable
	Obj       *Obj
}

// An Obj represents a collection of functions in a symbol table.
//
// The exact method of division of a binary into separate Objs is an internal detail
// of the symbol table format.
//
// In early versions of Go each source file became a different Obj.
//
// In Go 1 and Go 1.1, each package produced one Obj for all Go sources
// and one Obj per C source file.
//
// In Go 1.2, there is a single Obj for the entire program.
type Obj struct {
	Funcs []Func

	Paths []Sym
}

// Table represents a Go symbol table.  It stores all of the
// symbols decoded from the program and provides methods to translate
// between symbols, names, and addresses.
type Table struct {
	Syms  []Sym
	Funcs []Func
	Files map[string]*Obj
	Objs  []Obj

	go12line *LineTable
}

// NewTable decodes the Go symbol table in data,
// returning an in-memory representation.
func NewTable(symtab []byte, pcln *LineTable) (*Table, error)

// PCToFunc returns the function containing the program counter pc,
// or nil if there is no such function.
func (t *Table) PCToFunc(pc uint64) *Func

// PCToLine looks up line number information for a program counter.
// If there is no information, it returns fn == nil.
func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func)

// LineToPC looks up the first program counter on the given line in
// the named file.  It returns UnknownPathError or UnknownLineError if
// there is an error looking up this line.
func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err error)

// LookupSym returns the text, data, or bss symbol with the given name,
// or nil if no such symbol is found.
func (t *Table) LookupSym(name string) *Sym

// LookupFunc returns the text, data, or bss symbol with the given name,
// or nil if no such symbol is found.
func (t *Table) LookupFunc(name string) *Func

// SymByAddr returns the text, data, or bss symbol starting at the given address.
func (t *Table) SymByAddr(addr uint64) *Sym

// UnknownFileError represents a failure to find the specific file in
// the symbol table.
type UnknownFileError string

func (e UnknownFileError) Error() string

// UnknownLineError represents a failure to map a line to a program
// counter, either because the line is beyond the bounds of the file
// or because there is no code on the given line.
type UnknownLineError struct {
	File string
	Line int
}

func (e *UnknownLineError) Error() string

// DecodingError represents an error during the decoding of
// the symbol table.
type DecodingError struct {
	off int
	msg string
	val interface{}
}

func (e *DecodingError) Error() string
