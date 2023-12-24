// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package macho

import (
	"github.com/shogo82148/std/io"
)

// FatFileは、少なくとも1つのアーキテクチャを含むMach-Oユニバーサルバイナリです。
type FatFile struct {
	Magic  uint32
	Arches []FatArch
	closer io.Closer
}

// FatArchHeaderは、特定のイメージアーキテクチャのためのファットヘッダーを表します。
type FatArchHeader struct {
	Cpu    Cpu
	SubCpu uint32
	Offset uint32
	Size   uint32
	Align  uint32
}

// FatArchは、FatFile内のMach-Oファイルです。
type FatArch struct {
	FatArchHeader
	*File
}

<<<<<<< HEAD
// ErrNotFat is returned from [NewFatFile] or [OpenFat] when the file is not a
// universal binary but may be a thin binary, based on its magic number.
var ErrNotFat = &FormatError{0, "not a fat Mach-O file", nil}

// NewFatFile creates a new [FatFile] for accessing all the Mach-O images in a
// universal binary. The Mach-O binary is expected to start at position 0 in
// the ReaderAt.
func NewFatFile(r io.ReaderAt) (*FatFile, error)

// OpenFat opens the named file using [os.Open] and prepares it for use as a Mach-O
// universal binary.
=======
// ErrNotFatは、ファイルがユニバーサルバイナリではなく、
// マジックナンバーに基づいてシンバイナリである可能性がある場合、
// NewFatFileまたはOpenFatから返されます。
var ErrNotFat = &FormatError{0, "not a fat Mach-O file", nil}

// NewFatFileは、ユニバーサルバイナリ内のすべてのMach-Oイメージにアクセスするための新しいFatFileを作成します。
// Mach-Oバイナリは、ReaderAtの位置0で開始することが期待されています。
func NewFatFile(r io.ReaderAt) (*FatFile, error)

// OpenFatは、os.Openを使用して指定されたファイルを開き、それをMach-Oユニバーサルバイナリとして使用するための準備をします。
>>>>>>> release-branch.go1.21
func OpenFat(name string) (*FatFile, error)

func (ff *FatFile) Close() error
