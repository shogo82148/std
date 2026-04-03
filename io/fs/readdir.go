// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadDirFS is the interface implemented by a file system
// that provides an optimized implementation of [ReadDir].
type ReadDirFS interface {
	FS

	ReadDir(name string) ([]DirEntry, error)
}

// ReadDir reads the named directory
// and returns a list of directory entries sorted by filename.
//
// If fsys implements [ReadDirFS], ReadDir calls fsys.ReadDir.
// Otherwise ReadDir calls fsys.Open and uses ReadDir and Close
// on the returned [ReadDirFile].
func ReadDir(fsys FS, name string) ([]DirEntry, error)

// FileInfoToDirEntry returns a [DirEntry] that returns information from info.
// If info is nil, FileInfoToDirEntry returns nil.
func FileInfoToDirEntry(info FileInfo) DirEntry
