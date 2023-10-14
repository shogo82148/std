// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

// WorkingDir returns the current working directory
// (or "/???" if the directory cannot be identified),
// with "/" as separator.
func WorkingDir() string

// AbsFile returns the absolute filename for file in the given directory,
// as rewritten by the rewrites argument.
// For unrewritten paths, AbsFile rewrites a leading $GOROOT prefix to the literal "$GOROOT".
// If the resulting path is the empty string, the result is "??".
//
// The rewrites argument is a ;-separated list of rewrites.
// Each rewrite is of the form "prefix" or "prefix=>replace",
// where prefix must match a leading sequence of path elements
// and is either removed entirely or replaced by the replacement.
func AbsFile(dir, file, rewrites string) string

// ApplyRewrites returns the filename for file in the given directory,
// as rewritten by the rewrites argument.
//
// The rewrites argument is a ;-separated list of rewrites.
// Each rewrite is of the form "prefix" or "prefix=>replace",
// where prefix must match a leading sequence of path elements
// and is either removed entirely or replaced by the replacement.
func ApplyRewrites(file, rewrites string) (string, bool)
