// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

// Readdirはファイルに関連付けられたディレクトリーの内容を読み取り、
// ディレクトリーの順序で [Lstat] によって返される最大n個の [FileInfo] 値のスライスを返します。
// 同じファイルに対する後続の呼び出しは、さらにFileInfosを返します。
//
// n > 0の場合、Readdirは最大n個のFileInfo構造体を返します。この場合、
// Readdirが空のスライスを返すと、非nilのエラーが返されます。
// ディレクトリーの末尾では、エラーは [io.EOF] です。
//
// n <= 0の場合、Readdirはディレクトリー内のすべてのFileInfoを
// 単一のスライスで返します。この場合、Readdirが成功した場合
// （ディレクトリーの終わりまで読み込む）、スライスとnilのエラーが返されます。
// ディレクトリーの終わりより前にエラーが発生した場合、Readdirはその地点まで読み取った
// FileInfoと非nilのエラーを返します。
//
// パフォーマンスの向上のため、ほとんどのクライアントはより効率的なReadDirメソッドを使用することができます。
func (f *File) Readdir(n int) ([]FileInfo, error)

// Readdirnamesは、ファイルに関連付けられたディレクトリーの内容を読み取り、ディレクトリー内のファイルの名前を最大n個まで含むスライスを返します。追加の呼び出しでは、さらに名前を返します。
//
// n>0の場合、Readdirnamesは最大n個の名前を返します。この場合、Readdirnamesが空のスライスを返す場合は、非nilのエラーが返され、その理由が説明されます。ディレクトリーの終わりでは、エラーは [io.EOF] です。
//
// n <= 0の場合、Readdirnamesはディレクトリーからのすべての名前を単一のスライスで返します。この場合、Readdirnamesが成功し（ディレクトリーの終わりまで読み込むことができる）、スライスとnilのエラーを返します。ディレクトリーの終わり前にエラーが発生した場合、Readdirnamesはその時点まで読み取られた名前と非nilのエラーを返します。
func (f *File) Readdirnames(n int) (names []string, err error)

// DirEntryはディレクトリーから読み込まれたエントリーです
// ([ReadDir] 関数やファイルの [File.ReadDir] メソッドを使用して読み込まれます)。
type DirEntry = fs.DirEntry

// ReadDirはファイルfに関連付けられたディレクトリーの内容を読み取り、ディレクトリーの順序でDirEntry値のスライスを返します。
// 同じファイルに対する後続の呼び出しは、ディレクトリー内の後続のDirEntryレコードを生成します。
//
// n > 0の場合、ReadDirは最大n個の [DirEntry] レコードを返します。
// この場合、ReadDirが空のスライスを返す場合、なぜかを説明するエラーが返されます。
// ディレクトリーの終わりでは、エラーは [io.EOF] です。
//
// n <= 0の場合、ReadDirはディレクトリーに残っているすべてのDirEntryレコードを返します。
// 成功した場合、nilのエラーを返します（io.EOFではありません）。
func (f *File) ReadDir(n int) ([]DirEntry, error)

// ReadDirは指定されたディレクトリーを読み込み、
// ファイル名順にソートされたすべてのディレクトリーエントリーを返します。
// ディレクトリーの読み込み中にエラーが発生した場合、
// ReadDirはエラーが発生する前に読み込むことができたエントリーと共にエラーを返します。
func ReadDir(name string) ([]DirEntry, error)

// CopyFSは、ファイルシステムfsysをディレクトリーdirにコピーします。
// 必要に応じてdirを作成します。
//
// ファイルは、元の実行権限を加えたモード0o666で作成され、
// ディレクトリーはモード0o777（umask適用前）で作成されます。
//
// CopyFSは既存のファイルを上書きしません。fsys内のファイル名が
// すでに宛先に存在する場合、CopyFSは errors.Is(err, fs.ErrExist) がtrueとなる
// エラーを返します。
//
// dir内のシンボリックリンクは辿られます。
//
// CopyFSの実行中にfsysに新しいファイルが追加された場合（dirがfsysのサブディレクトリーの場合も含む）
// それらがコピーされる保証はありません。
//
// コピーは最初に発生したエラーで停止し、そのエラーを返します。
func CopyFS(dir string, fsys fs.FS) error
