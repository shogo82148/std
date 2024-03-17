// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// MkdirAllは、パスという名前のディレクトリと、必要な親ディレクトリを作成し、nilを返します。
// それ以外の場合はエラーを返します。
// MkdirAllが作成するすべてのディレクトリには、パーミッションビットperm（umaskの前）が使用されます。
// パスが既にディレクトリである場合、MkdirAllは何もせずにnilを返します。
func MkdirAll(path string, perm FileMode) error

<<<<<<< HEAD
// RemoveAllはpathとその中に含まれるすべての子要素を削除します。
// 削除できる範囲で削除を実行しますが、最初に出会ったエラーを返します。
// パスが存在しない場合、RemoveAllはnil（エラーなし）を返します。
// エラーがある場合、それは*PathError型のエラーです。
=======
// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters. If the path does not exist, RemoveAll
// returns nil (no error).
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/master
func RemoveAll(path string) error
