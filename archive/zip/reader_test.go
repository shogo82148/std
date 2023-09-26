// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"io"
	"os"
)

type ZipTest struct {
	Name    string
	Source  func() (r io.ReaderAt, size int64)
	Comment string
	File    []ZipTestFile
	Error   error
}

type ZipTestFile struct {
	Name  string
	Mode  os.FileMode
	Mtime string

	ContentErr error
	Content    []byte
	File       string
	Size       uint64
}
