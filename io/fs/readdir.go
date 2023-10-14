// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

<<<<<<< HEAD
// ReadDirFS is the interface implemented by a file system
// that provides an optimized implementation of [ReadDir].
=======
// ReadDirFSは、ReadDirの最適化された実装を提供するファイルシステムで実装されるインターフェースです。
>>>>>>> release-branch.go1.21
type ReadDirFS interface {
	FS

	ReadDir(name string) ([]DirEntry, error)
}

// ReadDirは指定されたディレクトリを読み取り、
// ファイル名でソートされたディレクトリエントリのリストを返します。
//
<<<<<<< HEAD
// If fs implements [ReadDirFS], ReadDir calls fs.ReadDir.
// Otherwise ReadDir calls fs.Open and uses ReadDir and Close
// on the returned file.
func ReadDir(fsys FS, name string) ([]DirEntry, error)

// FileInfoToDirEntry returns a [DirEntry] that returns information from info.
// If info is nil, FileInfoToDirEntry returns nil.
=======
// fsがReadDirFSを実装している場合、ReadDirはfs.ReadDirを呼び出します。
// そうでない場合、ReadDirはfs.Openを呼び出し、返されたファイルでReadDirとCloseを使用します。
func ReadDir(fsys FS, name string) ([]DirEntry, error)

// FileInfoToDirEntryは、infoから情報を返すDirEntryを返します。
// もしinfoがnilの場合、FileInfoToDirEntryはnilを返します。
>>>>>>> release-branch.go1.21
func FileInfoToDirEntry(info FileInfo) DirEntry
