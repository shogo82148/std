// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

package os

const (
	PathSeparator     = '/'
	PathListSeparator = ':'
)

// IsPathSeparator returns true if c is a directory separator character.
func IsPathSeparator(c uint8) bool
