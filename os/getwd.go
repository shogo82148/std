// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Getwdは現在のディレクトリに対応する絶対パス名を返します。
// 現在のディレクトリが（シンボリックリンクなどにより）複数のパスで到達可能な場合、
// Getwdはそのいずれかを返すことがあります。
//
// Unixプラットフォームでは、環境変数PWDが絶対パス名を提供し、
// それが現在のディレクトリ名である場合は、それが返されます。
func Getwd() (dir string, err error)
