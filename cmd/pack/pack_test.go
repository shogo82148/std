// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/fs"
)

// FakeFile implements FileLike and also fs.FileInfo.
type FakeFile struct {
	name     string
	contents string
	mode     fs.FileMode
	offset   int
}
