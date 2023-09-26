// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package zip provides support for reading and writing ZIP archives.

See: http://www.pkware.com/documents/casestudies/APPNOTE.TXT

This package does not support disk spanning.

A note about ZIP64:

To be backwards compatible the FileHeader has both 32 and 64 bit Size
fields. The 64 bit fields will always contain the correct value and
for normal archives both fields will be the same. For files requiring
the ZIP64 format the 32 bit fields will be 0xffffffff and the 64 bit
fields must be used instead.
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

// FileHeader describes a file within a zip file.
// See the zip spec for details.
type FileHeader struct {
	Name string

	CreatorVersion     uint16
	ReaderVersion      uint16
	Flags              uint16
	Method             uint16
	ModifiedTime       uint16
	ModifiedDate       uint16
	CRC32              uint32
	CompressedSize     uint32
	UncompressedSize   uint32
	CompressedSize64   uint64
	UncompressedSize64 uint64
	Extra              []byte
	ExternalAttrs      uint32
	Comment            string
}

// FileInfo returns an os.FileInfo for the FileHeader.
func (h *FileHeader) FileInfo() os.FileInfo

// headerFileInfo implements os.FileInfo.

// FileInfoHeader creates a partially-populated FileHeader from an
// os.FileInfo.
// Because os.FileInfo's Name method returns only the base name of
// the file it describes, it may be necessary to modify the Name field
// of the returned header to provide the full path name of the file.
func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)

// ModTime returns the modification time in UTC.
// The resolution is 2s.
func (h *FileHeader) ModTime() time.Time

// SetModTime sets the ModifiedTime and ModifiedDate fields to the given time in UTC.
// The resolution is 2s.
func (h *FileHeader) SetModTime(t time.Time)

// Mode returns the permission and mode bits for the FileHeader.
func (h *FileHeader) Mode() (mode os.FileMode)

// SetMode changes the permission and mode bits for the FileHeader.
func (h *FileHeader) SetMode(mode os.FileMode)
