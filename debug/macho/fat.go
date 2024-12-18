// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package macho

import (
	"github.com/shogo82148/std/io"
)

// A FatFile is a Mach-O universal binary that contains at least one architecture.
type FatFile struct {
	Magic  uint32
	Arches []FatArch
	closer io.Closer
}

// A FatArchHeader represents a fat header for a specific image architecture.
type FatArchHeader struct {
	Cpu    Cpu
	SubCpu uint32
	Offset uint32
	Size   uint32
	Align  uint32
}

// A FatArch is a Mach-O File inside a FatFile.
type FatArch struct {
	FatArchHeader
	*File
}

// ErrNotFat is returned from [NewFatFile] or [OpenFat] when the file is not a
// universal binary but may be a thin binary, based on its magic number.
var ErrNotFat = &FormatError{0, "not a fat Mach-O file", nil}

// NewFatFile creates a new [FatFile] for accessing all the Mach-O images in a
// universal binary. The Mach-O binary is expected to start at position 0 in
// the ReaderAt.
func NewFatFile(r io.ReaderAt) (*FatFile, error)

// OpenFat opens the named file using [os.Open] and prepares it for use as a Mach-O
// universal binary.
func OpenFat(name string) (*FatFile, error)

func (ff *FatFile) Close() error
