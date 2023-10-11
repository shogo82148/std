// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// A StatFS is a file system with a Stat method.
type StatFS interface {
	FS

	Stat(name string) (FileInfo, error)
}

// Stat returns a FileInfo describing the named file from the file system.
//
// If fs implements StatFS, Stat calls fs.Stat.
// Otherwise, Stat opens the file to stat it.
func Stat(fsys FS, name string) (FileInfo, error)
