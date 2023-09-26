// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris

package os

const (
	PathSeparator     = '/'
	PathListSeparator = ':'
)

// IsPathSeparator reports whether c is a directory separator character.
func IsPathSeparator(c uint8) bool
