// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadDirFS is the interface implemented by a file system
// that provides an optimized implementation of ReadDir.
type ReadDirFS interface {
	FS

	ReadDir(name string) ([]DirEntry, error)
}

// ReadDir reads the named directory
// and returns a list of directory entries sorted by filename.
//
// If fs implements ReadDirFS, ReadDir calls fs.ReadDir.
// Otherwise ReadDir calls fs.Open and uses ReadDir and Close
// on the returned file.
func ReadDir(fsys FS, name string) ([]DirEntry, error)

// dirInfo is a DirEntry based on a FileInfo.

// FileInfoToDirEntry returns a DirEntry that returns information from info.
// If info is nil, FileInfoToDirEntry returns nil.
func FileInfoToDirEntry(info FileInfo) DirEntry
