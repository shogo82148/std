// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type *PathError.
func (file *File) Stat() (fi FileInfo, err error)

// Stat returns a FileInfo structure describing the named file.
// If there is an error, it will be of type *PathError.
func Stat(name string) (fi FileInfo, err error)

// Lstat returns the FileInfo structure describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link.  Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
func Lstat(name string) (fi FileInfo, err error)
