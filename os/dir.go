// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

// Readdirはファイルに関連付けられたディレクトリの内容を読み取り、
// ディレクトリの順序で [Lstat] によって返される最大n個の [FileInfo] 値のスライスを返します。
// 同じファイルに対する後続の呼び出しは、さらにFileInfosを返します。
//
// n > 0の場合、Readdirは最大n個のFileInfo構造体を返します。この場合、
// Readdirが空のスライスを返すと、非nilのエラーが返されます。
// ディレクトリの末尾では、エラーは [io.EOF] です。
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
// n>0の場合、Readdirnamesは最大n個の名前を返します。この場合、Readdirnamesが空のスライスを返す場合は、非nilのエラーが返され、その理由が説明されます。ディレクトリの終わりでは、エラーは [io.EOF] です。
//
// n <= 0の場合、Readdirnamesはディレクトリからのすべての名前を単一のスライスで返します。この場合、Readdirnamesが成功し（ディレクトリの終わりまで読み込むことができる）、スライスとnilのエラーを返します。ディレクトリの終わり前にエラーが発生した場合、Readdirnamesはその時点まで読み取られた名前と非nilのエラーを返します。
func (f *File) Readdirnames(n int) (names []string, err error)

// DirEntryはディレクトリから読み込まれたエントリです
// ([ReadDir] 関数やファイルの [File.ReadDir] メソッドを使用して読み込まれます)。
type DirEntry = fs.DirEntry

// ReadDirはファイルfに関連付けられたディレクトリの内容を読み取り、ディレクトリの順序でDirEntry値のスライスを返します。
// 同じファイルに対する後続の呼び出しは、ディレクトリ内の後続のDirEntryレコードを生成します。
//
// n > 0の場合、ReadDirは最大n個の [DirEntry] レコードを返します。
// この場合、ReadDirが空のスライスを返す場合、なぜかを説明するエラーが返されます。
// ディレクトリの終わりでは、エラーは [io.EOF] です。
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
// Files are created with mode 0o666 plus any execute permissions
// from the source, and directories are created with mode 0o777
// (before umask).
//
// CopyFS will not overwrite existing files. If a file name in fsys
// already exists in the destination, CopyFS will return an error
// such that errors.Is(err, fs.ErrExist) will be true.
<<<<<<< HEAD
//
// Symbolic links in fsys are not supported. A *PathError with Err set
// to ErrInvalid is returned when copying from a symbolic link.
=======
>>>>>>> upstream/release-branch.go1.25
//
// Symbolic links in dir are followed.
//
// New files added to fsys (including if dir is a subdirectory of fsys)
// while CopyFS is running are not guaranteed to be copied.
//
// Copying stops at and returns the first error encountered.
func CopyFS(dir string, fsys fs.FS) error
