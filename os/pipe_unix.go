// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin

package os

// Pipe は接続された一対のファイルを返します。rからの読み取りはwに書き込まれたバイトを返します。
// ファイルとエラー（ある場合）を返します。
func Pipe() (r *File, w *File, err error)
