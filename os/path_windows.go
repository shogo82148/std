// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

const (
	PathSeparator     = '\\'
	PathListSeparator = ';'
)

// IsPathSeparatorは、cがディレクトリの区切り文字であるかどうかを報告します。
func IsPathSeparator(c uint8) bool
