// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

const (
	PathSeparator     = '\\'
	PathListSeparator = ';'
)

// IsPathSeparator reports whether c is a directory separator character.
func IsPathSeparator(c uint8) bool

// This is set via go:linkname on runtime.canUseLongPaths, and is true when the OS
// supports opting into proper long path handling without the need for fixups.
