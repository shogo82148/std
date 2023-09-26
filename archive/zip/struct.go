// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package zip provides support for reading and writing ZIP archives.

See: http://www.pkware.com/documents/casestudies/APPNOTE.TXT

This package does not support ZIP64 or disk spanning.
*/
package zip

import (
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

// Compression methods.
const (
	Store   uint16 = 0
	Deflate uint16 = 8
)

type FileHeader struct {
	Name             string
	CreatorVersion   uint16
	ReaderVersion    uint16
	Flags            uint16
	Method           uint16
	ModifiedTime     uint16
	ModifiedDate     uint16
	CRC32            uint32
	CompressedSize   uint32
	UncompressedSize uint32
	Extra            []byte
	ExternalAttrs    uint32
	Comment          string
}

// FileInfo returns an os.FileInfo for the FileHeader.
func (h *FileHeader) FileInfo() os.FileInfo

// headerFileInfo implements os.FileInfo.

// FileInfoHeader creates a partially-populated FileHeader from an
// os.FileInfo.
func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)

// ModTime returns the modification time.
// The resolution is 2s.
func (h *FileHeader) ModTime() time.Time

// SetModTime sets the ModifiedTime and ModifiedDate fields to the given time.
// The resolution is 2s.
func (h *FileHeader) SetModTime(t time.Time)

// Mode returns the permission and mode bits for the FileHeader.
func (h *FileHeader) Mode() (mode os.FileMode)

// SetMode changes the permission and mode bits for the FileHeader.
func (h *FileHeader) SetMode(mode os.FileMode)
