// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadLinkFSは、シンボリックリンクの読み取りをサポートするファイルシステムによって
// 実装されるインターフェースです。
type ReadLinkFS interface {
	FS

	ReadLink(name string) (string, error)

	Lstat(name string) (FileInfo, error)
}

// ReadLinkは、指定されたシンボリックリンクのリンク先を返します。
//
// fsysが[ReadLinkFS]を実装していない場合、ReadLinkはエラーを返します。
func ReadLink(fsys FS, name string) (string, error)

// Lstatは、指定されたファイルを記述する[FileInfo]を返します。
// ファイルがシンボリックリンクの場合、返される[FileInfo]はシンボリックリンクを記述します。
// Lstatはリンクをたどろうとはしません。
//
// fsysが[ReadLinkFS]を実装していない場合、Lstatは[Stat]と同じです。
func Lstat(fsys FS, name string) (FileInfo, error)
