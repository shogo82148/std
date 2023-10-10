// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// ReadDirFSは、ReadDirの最適化された実装を提供するファイルシステムで実装されるインターフェースです。
type ReadDirFS interface {
	FS

	// ReadDir reads the named directory
	// and returns a list of directory entries sorted by filename.
	ReadDir(name string) ([]DirEntry, error)
}

// ReadDirは指定されたディレクトリを読み取り、
// ファイル名でソートされたディレクトリエントリのリストを返します。
//
// fsがReadDirFSを実装している場合、ReadDirはfs.ReadDirを呼び出します。
// そうでない場合、ReadDirはfs.Openを呼び出し、返されたファイルでReadDirとCloseを使用します。
func ReadDir(fsys FS, name string) ([]DirEntry, error)

// FileInfoToDirEntryは、infoから情報を返すDirEntryを返します。
// もしinfoがnilの場合、FileInfoToDirEntryはnilを返します。
func FileInfoToDirEntry(info FileInfo) DirEntry
