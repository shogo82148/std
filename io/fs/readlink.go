// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadLinkFS is the interface implemented by a file system
// that supports reading symbolic links.
type ReadLinkFS interface {
	FS

	ReadLink(name string) (string, error)

	Lstat(name string) (FileInfo, error)
}

// ReadLink returns the destination of the named symbolic link.
//
// If fsys does not implement [ReadLinkFS], then ReadLink returns an error.
func ReadLink(fsys FS, name string) (string, error)

// Lstat returns a [FileInfo] describing the named file.
// If the file is a symbolic link, the returned [FileInfo] describes the symbolic link.
// Lstat makes no attempt to follow the link.
//
// If fsys does not implement [ReadLinkFS], then Lstat is identical to [Stat].
func Lstat(fsys FS, name string) (FileInfo, error)
