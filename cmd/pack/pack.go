// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/os"
)

// The unusual ancestry means the arguments are not Go-standard.
// These variables hold the decoded operation specified by the first argument.
// op holds the operation we are doing (prtx).
// verbose tells whether the 'v' option was specified.

// An Archive represents an open archive file. It is always scanned sequentially
// from start to end, without backing up.
type Archive struct {
	fd       *os.File
	files    []string
	pad      int
	matchAll bool
}

// An Entry is the internal representation of the per-file header information of one entry in the archive.
type Entry struct {
	name  string
	mtime int64
	uid   int
	gid   int
	mode  os.FileMode
	size  int64
}

func (e *Entry) String() string

// FileLike abstracts the few methods we need, so we can test without needing real files.
type FileLike interface {
	Name() string
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

// can be modified for testing.
