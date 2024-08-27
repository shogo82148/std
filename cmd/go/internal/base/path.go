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
// There are rare cases where the path produced by ShortPath could be incorrect
// so it should only be used when formatting paths for error messages, not to read
// a file.
func ShortPath(path string) string

// ShortPathConservative is similar to ShortPath, but returns the input if the result of ShortPath
// would meet conditions that could make it invalid. If the short path would reach into a
// parent directory and the base path contains a symlink, a ".." component can
// cross a symlink boundary. That could be a problem because the symlinks could be evaluated,
// changing the relative location of the boundary, before the ".." terms are applied to
// go to parents. The check here is a little more conservative: it checks
// whether the path starts with a ../ or ..\ component, and if any of the parent directories
// of the working directory are symlinks.
// See #68383 for a case where this could happen.
func ShortPathConservative(path string) string

// RelPaths returns a copy of paths with absolute paths
// made relative to the current directory if they would be shorter.
func RelPaths(paths []string) []string

// IsTestFile reports whether the source file is a set of tests and should therefore
// be excluded from coverage analysis.
func IsTestFile(file string) bool

// IsNull reports whether the path is a common name for the null device.
// It returns true for /dev/null on Unix, or NUL (case-insensitive) on Windows.
func IsNull(path string) bool
