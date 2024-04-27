// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepathlite

const (
	Separator     = '/'
	ListSeparator = '\000'
)

func IsPathSeparator(c uint8) bool

// IsAbs reports whether the path is absolute.
func IsAbs(path string) bool
