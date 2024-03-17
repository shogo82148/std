// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

<<<<<<< HEAD
// Readdirはファイルに関連付けられたディレクトリの内容を読み取り、
// ディレクトリの順序でLstatによって返される最大n個のFileInfo値のスライスを返します。
// 同じファイルに対する後続の呼び出しは、さらにFileInfosを返します。
//
// n > 0の場合、Readdirは最大n個のFileInfo構造体を返します。この場合、
// Readdirが空のスライスを返すと、非nilのエラーが返されます。
// ディレクトリの末尾では、エラーはio.EOFです。
=======
// Readdir reads the contents of the directory associated with file and
// returns a slice of up to n [FileInfo] values, as would be returned
// by [Lstat], in directory order. Subsequent calls on the same file will yield
// further FileInfos.
//
// If n > 0, Readdir returns at most n FileInfo structures. In this case, if
// Readdir returns an empty slice, it will return a non-nil error
// explaining why. At the end of a directory, the error is [io.EOF].
>>>>>>> upstream/master
//
// n <= 0の場合、Readdirはディレクトリ内のすべてのFileInfoを
// 単一のスライスで返します。この場合、Readdirが成功した場合
// （ディレクトリの終わりまで読み込む）、スライスとnilのエラーが返されます。
// ディレクトリの終わりより前にエラーが発生した場合、Readdirはその地点まで読み取った
// FileInfoと非nilのエラーを返します。
//
// パフォーマンスの向上のため、ほとんどのクライアントはより効率的なReadDirメソッドを使用することができます。
func (f *File) Readdir(n int) ([]FileInfo, error)

// Readdirnamesは、ファイルに関連付けられたディレクトリの内容を読み取り、ディレクトリ内のファイルの名前を最大n個まで含むスライスを返します。追加の呼び出しでは、さらに名前を返します。
//
<<<<<<< HEAD
// n>0の場合、Readdirnamesは最大n個の名前を返します。この場合、Readdirnamesが空のスライスを返す場合は、非nilのエラーが返され、その理由が説明されます。ディレクトリの終わりでは、エラーはio.EOFです。
=======
// If n > 0, Readdirnames returns at most n names. In this case, if
// Readdirnames returns an empty slice, it will return a non-nil error
// explaining why. At the end of a directory, the error is [io.EOF].
>>>>>>> upstream/master
//
// n <= 0の場合、Readdirnamesはディレクトリからのすべての名前を単一のスライスで返します。この場合、Readdirnamesが成功し（ディレクトリの終わりまで読み込むことができる）、スライスとnilのエラーを返します。ディレクトリの終わり前にエラーが発生した場合、Readdirnamesはその時点まで読み取られた名前と非nilのエラーを返します。
func (f *File) Readdirnames(n int) (names []string, err error)

<<<<<<< HEAD
// DirEntryはディレクトリから読み込まれたエントリです
// (ReadDir関数やファイルのReadDirメソッドを使用して読み込まれます)。
type DirEntry = fs.DirEntry

// ReadDirはファイルfに関連付けられたディレクトリの内容を読み取り、ディレクトリの順序でDirEntry値のスライスを返します。
// 同じファイルに対する後続の呼び出しは、ディレクトリ内の後続のDirEntryレコードを生成します。
//
// n > 0の場合、ReadDirは最大n個のDirEntryレコードを返します。
// この場合、ReadDirが空のスライスを返す場合、なぜかを説明するエラーが返されます。
// ディレクトリの終わりでは、エラーはio.EOFです。
=======
// A DirEntry is an entry read from a directory
// (using the [ReadDir] function or a [File.ReadDir] method).
type DirEntry = fs.DirEntry

// ReadDir reads the contents of the directory associated with the file f
// and returns a slice of [DirEntry] values in directory order.
// Subsequent calls on the same file will yield later DirEntry records in the directory.
//
// If n > 0, ReadDir returns at most n DirEntry records.
// In this case, if ReadDir returns an empty slice, it will return an error explaining why.
// At the end of a directory, the error is [io.EOF].
>>>>>>> upstream/master
//
// n <= 0の場合、ReadDirはディレクトリに残っているすべてのDirEntryレコードを返します。
// 成功した場合、nilのエラーを返します（io.EOFではありません）。
func (f *File) ReadDir(n int) ([]DirEntry, error)

// ReadDirは指定されたディレクトリを読み込み、
// ファイル名順にソートされたすべてのディレクトリエントリを返します。
// ディレクトリの読み込み中にエラーが発生した場合、
// ReadDirはエラーが発生する前に読み込むことができたエントリと共にエラーを返します。
func ReadDir(name string) ([]DirEntry, error)

// CopyFS copies the file system fsys into the directory dir,
// creating dir if necessary.
//
// Newly created directories and files have their default modes
// where any bits from the file in fsys that are not part of the
// standard read, write, and execute permissions will be zeroed
// out, and standard read and write permissions are set for owner,
// group, and others while retaining any existing execute bits from
// the file in fsys.
//
// Symbolic links in fsys are not supported, a *PathError with Err set
// to ErrInvalid is returned on symlink.
//
// Copying stops at and returns the first error encountered.
func CopyFS(dir string, fsys fs.FS) error
