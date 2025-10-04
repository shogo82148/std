// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

<<<<<<< HEAD
// Getwdは現在のディレクトリに対応するルート付きパス名を返します。現在のディレクトリがシンボリックリンクによって複数のパスで到達可能な場合、Getwdはそのいずれかを返すことがあります。
=======
// Getwd returns an absolute path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
//
// On Unix platforms, if the environment variable PWD
// provides an absolute name, and it is a name of the
// current directory, it is returned.
>>>>>>> upstream/release-branch.go1.25
func Getwd() (dir string, err error)
