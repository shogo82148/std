// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package os

const (
	PathSeparator     = '/'
	PathListSeparator = ':'
)

// IsPathSeparator は c がディレクトリの区切り文字であるかどうかを報告します。
func IsPathSeparator(c uint8) bool
