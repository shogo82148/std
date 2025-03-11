// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package is a lightly modified version of the mmap code
// in github.com/google/codesearch/index.

// The mmap package provides an abstraction for memory mapping files
// on different platforms.
package mmap

import (
	"github.com/shogo82148/std/os"
)

// Data is mmap'ed read-only data from a file.
// The backing file is never closed, so Data
// remains valid for the lifetime of the process.
type Data struct {
	f    *os.File
	Data []byte
}

// Mmap maps the given file into memory.
// The boolean result indicates whether the file was opened.
// If it is true, the caller should avoid attempting
// to write to the file on Windows, because Windows locks
// the open file, and writes to it will fail.
func Mmap(file string) (Data, bool, error)
