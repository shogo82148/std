// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/io/fs"
)

// アーカイブは、開いたアーカイブファイルを表します。バックアップを取らずに、常に開始から終了まで順にスキャンされます。
type Archive struct {
	a        *archive.Archive
	files    []string
	pad      int
	matchAll bool
}

// FileLikeは、実際のファイルを必要とせずにテストするために必要なわずかなメソッドを抽象化します。
type FileLike interface {
	Name() string
	Stat() (fs.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
