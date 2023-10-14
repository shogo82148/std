// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// A GlobFS is a file system with a Glob method.
type GlobFS interface {
	FS

	// Glob returns the names of all files matching pattern,
	// providing an implementation of the top-level
	// Glob function.
	Glob(pattern string) ([]string, error)
}

// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in path.Match. The pattern may describe hierarchical names such as
// usr/*/bin/ed.
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is path.ErrBadPattern, reporting that
// the pattern is malformed.
//
// If fs implements GlobFS, Glob calls fs.Glob.
// Otherwise, Glob uses ReadDir to traverse the directory tree
// and look for matches for the pattern.
func Glob(fsys FS, pattern string) (matches []string, err error)
