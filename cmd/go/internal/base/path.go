// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

// UncachedCwd returns the current working directory.
// Most callers should use Cwd, which caches the result for future use.
// UncachedCwd is appropriate to call early in program startup before flag parsing,
// because the -C flag may change the current directory.
func UncachedCwd() string

// Cwd returns the current working directory at the time of the first call.
func Cwd() string

// ShortPath returns an absolute or relative name for path, whatever is shorter.
func ShortPath(path string) string

// RelPaths returns a copy of paths with absolute paths
// made relative to the current directory if they would be shorter.
func RelPaths(paths []string) []string

// IsTestFile reports whether the source file is a set of tests and should therefore
// be excluded from coverage analysis.
func IsTestFile(file string) bool

// IsNull reports whether the path is a common name for the null device.
// It returns true for /dev/null on Unix, or NUL (case-insensitive) on Windows.
func IsNull(path string) bool
