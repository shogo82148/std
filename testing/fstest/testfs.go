// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fstest implements support for testing implementations and users of file systems.
package fstest

import (
	"github.com/shogo82148/std/io/fs"
)

// TestFS tests a file system implementation.
// It walks the entire tree of files in fsys,
// opening and checking that each file behaves correctly.
// Symbolic links are not followed,
// but their Lstat values are checked
// if the file system implements [fs.ReadLinkFS].
// It also checks that the file system contains at least the expected files.
// As a special case, if no expected files are listed, fsys must be empty.
// Otherwise, fsys must contain at least the listed files; it can also contain others.
// The contents of fsys must not change concurrently with TestFS.
//
// If TestFS finds any misbehaviors, it returns either the first error or a
// list of errors. Use [errors.Is] or [errors.As] to inspect.
//
// Typical usage inside a test is:
//
//	if err := fstest.TestFS(myFS, "file/that/should/be/present"); err != nil {
//		t.Fatal(err)
//	}
func TestFS(fsys fs.FS, expected ...string) error
