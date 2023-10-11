// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

const (
	PathSeparator     = '/'
	PathListSeparator = '\000'
)

// IsPathSeparatorは、文字cがディレクトリセパレータ文字かどうかを報告します。
func IsPathSeparator(c uint8) bool
