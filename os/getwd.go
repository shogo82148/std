// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Getwdは現在のディレクトリに対応するルート付きパス名を返します。現在のディレクトリがシンボリックリンクによって複数のパスで到達可能な場合、Getwdはそのいずれかを返すことがあります。
func Getwd() (dir string, err error)
