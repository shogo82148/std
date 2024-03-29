// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package os

// Statはファイルに関する情報を [FileInfo] 構造体で返します。
// エラーがある場合、[*PathError] 型になります。
func (f *File) Stat() (FileInfo, error)
