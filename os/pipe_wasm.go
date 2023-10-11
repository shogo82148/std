// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

package os

// Pipe は接続された2つのファイルを返します。r から読み込まれたバイトは w に書き込まれます。
// ファイルとエラー（あれば）を返します。
func Pipe() (r *File, w *File, err error)
