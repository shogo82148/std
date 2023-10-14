// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

// HasPathPrefix reports whether the slash-separated path s
// begins with the elements in prefix.
func HasPathPrefix(s, prefix string) bool

// HasFilePathPrefix reports whether the filesystem path s
// begins with the elements in prefix.
//
// HasFilePathPrefix is case-sensitive (except for volume names) even if the
// filesystem is not, does not apply Unicode normalization even if the
// filesystem does, and assumes that all path separators are canonicalized to
// filepath.Separator (as returned by filepath.Clean).
func HasFilePathPrefix(s, prefix string) bool

// TrimFilePathPrefix returns s without the leading path elements in prefix,
// such that joining the string to prefix produces s.
//
// If s does not start with prefix (HasFilePathPrefix with the same arguments
// returns false), TrimFilePathPrefix returns s. If s equals prefix,
// TrimFilePathPrefix returns "".
func TrimFilePathPrefix(s, prefix string) string

// WithFilePathSeparator returns s with a trailing path separator, or the empty
// string if s is empty.
func WithFilePathSeparator(s string) string

// QuoteGlob returns s with all Glob metacharacters quoted.
// We don't try to handle backslash here, as that can appear in a
// file path on Windows.
func QuoteGlob(s string) string
