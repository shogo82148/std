// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tar implements access to tar archives.
//
// Tape archives (tar) are a file format for storing a sequence of files that
// can be read and written in a streaming manner.
// This package aims to cover most variations of the format,
// including those produced by GNU and BSD tar tools.
package tar

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

var (
	ErrHeader          = errors.New("archive/tar: invalid tar header")
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")
	ErrWriteAfterClose = errors.New("archive/tar: write after close")
)

// Type flags for Header.Typeflag.
const (
	// Type '0' indicates a regular file.
	TypeReg  = '0'
	TypeRegA = '\x00'

	// Type '1' to '6' are header-only flags and may not have a data body.
	TypeLink    = '1'
	TypeSymlink = '2'
	TypeChar    = '3'
	TypeBlock   = '4'
	TypeDir     = '5'
	TypeFifo    = '6'

	// Type '7' is reserved.
	TypeCont = '7'

	// Type 'x' is used by the PAX format to store key-value records that
	// are only relevant to the next file.
	// This package transparently handles these types.
	TypeXHeader = 'x'

	// Type 'g' is used by the PAX format to store key-value records that
	// are relevant to all subsequent files.
	// This package only supports parsing and composing such headers,
	// but does not currently support persisting the global state across files.
	TypeXGlobalHeader = 'g'

	// Type 'S' indicates a sparse file in the GNU format.
	TypeGNUSparse = 'S'

	// Types 'L' and 'K' are used by the GNU format for a meta file
	// used to store the path or link name for the next file.
	// This package transparently handles these types.
	TypeGNULongName = 'L'
	TypeGNULongLink = 'K'
)

// Keywords for PAX extended header records.

// basicKeys is a set of the PAX keys for which we have built-in support.
// This does not contain "charset" or "comment", which are both PAX-specific,
// so adding them as first-class features of Header is unlikely.
// Users can use the PAXRecords field to set it themselves.

// A Header represents a single header in a tar archive.
// Some fields may not be populated.
//
// For forward compatibility, users that retrieve a Header from Reader.Next,
// mutate it in some ways, and then pass it back to Writer.WriteHeader
// should do so by creating a new Header and copying the fields
// that they are interested in preserving.
type Header struct {
	Typeflag byte

	Name     string
	Linkname string

	Size  int64
	Mode  int64
	Uid   int
	Gid   int
	Uname string
	Gname string

	ModTime    time.Time
	AccessTime time.Time
	ChangeTime time.Time

	Devmajor int64
	Devminor int64

	Xattrs map[string]string

	PAXRecords map[string]string

	Format Format
}

// sparseEntry represents a Length-sized fragment at Offset in the file.

// A sparse file can be represented as either a sparseDatas or a sparseHoles.
// As long as the total size is known, they are equivalent and one can be
// converted to the other form and back. The various tar formats with sparse
// file support represent sparse files in the sparseDatas form. That is, they
// specify the fragments in the file that has data, and treat everything else as
// having zero bytes. As such, the encoding and decoding logic in this package
// deals with sparseDatas.
//
// However, the external API uses sparseHoles instead of sparseDatas because the
// zero value of sparseHoles logically represents a normal file (i.e., there are
// no holes in it). On the other hand, the zero value of sparseDatas implies
// that the file has no data in it, which is rather odd.
//
// As an example, if the underlying raw file contains the 10-byte data:
//	var compactFile = "abcdefgh"
//
// And the sparse map has the following entries:
//	var spd sparseDatas = []sparseEntry{
//		{Offset: 2,  Length: 5},  // Data fragment for 2..6
//		{Offset: 18, Length: 3},  // Data fragment for 18..20
//	}
//	var sph sparseHoles = []sparseEntry{
//		{Offset: 0,  Length: 2},  // Hole fragment for 0..1
//		{Offset: 7,  Length: 11}, // Hole fragment for 7..17
//		{Offset: 21, Length: 4},  // Hole fragment for 21..24
//	}
//
// Then the content of the resulting sparse file with a Header.Size of 25 is:
//	var sparseFile = "\x00"*2 + "abcde" + "\x00"*11 + "fgh" + "\x00"*4

// fileState tracks the number of logical (includes sparse holes) and physical
// (actual in tar archive) bytes remaining for the current file.
//
// Invariant: LogicalRemaining >= PhysicalRemaining

// FileInfo returns an os.FileInfo for the Header.
func (h *Header) FileInfo() os.FileInfo

// headerFileInfo implements os.FileInfo.

// sysStat, if non-nil, populates h from system-dependent fields of fi.

// FileInfoHeader creates a partially-populated Header from fi.
// If fi describes a symlink, FileInfoHeader records link as the link target.
// If fi describes a directory, a slash is appended to the name.
//
// Since os.FileInfo's Name method only returns the base name of
// the file it describes, it may be necessary to modify Header.Name
// to provide the full path name of the file.
func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
