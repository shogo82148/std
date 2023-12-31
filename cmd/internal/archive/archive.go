// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package archive implements reading of archive files generated by the Go
// toolchain.
package archive

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
)

// A Data is a reference to data stored in an object file.
// It records the offset and size of the data, so that a client can
// read the data only if necessary.
type Data struct {
	Offset int64
	Size   int64
}

type Archive struct {
	f       *os.File
	Entries []Entry
}

func (a *Archive) File() *os.File

type Entry struct {
	Name  string
	Type  EntryType
	Mtime int64
	Uid   int
	Gid   int
	Mode  os.FileMode
	Data
	Obj *GoObj
}

type EntryType int

const (
	EntryPkgDef EntryType = iota
	EntryGoObj
	EntryNativeObj
	EntrySentinelNonObj
)

func (e *Entry) String() string

type GoObj struct {
	TextHeader []byte
	Arch       string
	Data
}

type ErrGoObjOtherVersion struct{ magic []byte }

func (e ErrGoObjOtherVersion) Error() string

// New writes to f to make a new archive.
func New(f *os.File) (*Archive, error)

// Parse parses an object file or archive from f.
func Parse(f *os.File, verbose bool) (*Archive, error)

// AddEntry adds an entry to the end of a, with the content from r.
func (a *Archive) AddEntry(typ EntryType, name string, mtime int64, uid, gid int, mode os.FileMode, size int64, r io.Reader)

// architecture-independent object file output
const HeaderSize = 60

func ReadHeader(b *bufio.Reader, name string) int

func FormatHeader(arhdr []byte, name string, size int64)
