// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xcoff

import (
	"github.com/shogo82148/std/io"
)

const (
	SAIAMAG   = 0x8
	AIAFMAG   = "`\n"
	AIAMAG    = "<aiaff>\n"
	AIAMAGBIG = "<bigaf>\n"

	// Sizeof
	FL_HSZ_BIG = 0x80
	AR_HSZ_BIG = 0x70
)

// Archive represents an open AIX big archive.
type Archive struct {
	ArchiveHeader
	Members []*Member

	closer io.Closer
}

// ArchiveHeader holds information about a big archive file header
type ArchiveHeader struct {
	magic string
}

// Member represents a member of an AIX big archive.
type Member struct {
	MemberHeader
	sr *io.SectionReader
}

// MemberHeader holds information about a big archive member
type MemberHeader struct {
	Name string
	Size uint64
}

// OpenArchive opens the named archive using os.Open and prepares it for use
// as an AIX big archive.
func OpenArchive(name string) (*Archive, error)

// Close closes the Archive.
// If the Archive was created using NewArchive directly instead of OpenArchive,
// Close has no effect.
func (a *Archive) Close() error

// NewArchive creates a new Archive for accessing an AIX big archive in an underlying reader.
func NewArchive(r io.ReaderAt) (*Archive, error)

// GetFile returns the XCOFF file defined by member name.
// FIXME: This doesn't work if an archive has two members with the same
// name which can occur if an archive has both 32-bits and 64-bits files.
func (arch *Archive) GetFile(name string) (*File, error)
