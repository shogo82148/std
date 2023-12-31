// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// StatはファイルについてのFileInfo構造体を返します。
// エラーがある場合は*PathError型になります。
func (file *File) Stat() (FileInfo, error)
