// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package os

// Pipe は接続された一対のファイルを返します。r からの読み取りは w に書き込まれます。
// ファイルとエラーを返します。
func Pipe() (r *File, w *File, err error)
