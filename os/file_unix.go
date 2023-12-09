// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package os

// Fdメソッドは、オープンされたファイルを参照する整数型のUnixファイルディスクリプタを返します。
// fがクローズされると、ファイルディスクリプタは無効になります。
// fがガベージコレクトされると、ファイナライザによってファイルディスクリプタがクローズされることがあります。
// これにより、SetDeadlineメソッドが動作しなくなる可能性があります。
// ファイルディスクリプタは再利用される可能性があるため、返されたファイルディスクリプタは
// fのCloseメソッドを通じてのみクローズするか、ガベージコレクション時のファイナライザによってクローズされます。
// そうでない場合、ガベージコレクション時にはファイナライザが同じ（再利用された）番号を持つ他のファイルディスクリプタをクローズする可能性があります。
//
// 代替として、f.SyscallConnメソッドを参照してください。
func (f *File) Fd() uintptr

// NewFileは指定されたファイルディスクリプタと名前で新しいFileを返します。
// fdが有効なファイルディスクリプタでない場合、返される値はnilになります。
// Unixシステムでは、ファイルディスクリプタが非同期モードの場合、NewFileはSetDeadlineメソッドが動作するポーラブルなFileを返そうとします。
//
// NewFileに渡した後、fdはFdメソッドのコメントで説明されている条件の下で無効になり、同じ制約が適用されます。
func NewFile(fd uintptr, name string) *File

// DevNullはオペレーティングシステムの「nullデバイス」の名前です。
// Unix-likeなシステムでは、"/dev/null"です。Windowsでは、"NUL"です。
const DevNull = "/dev/null"

func Truncate(name string, size int64) error

// Removeは指定したファイルまたは(空の)ディレクトリを削除します。
// エラーが発生した場合、*PathError型のエラーとなります。
func Remove(name string) error

// Link は newname を oldname ファイルのハードリンクとして作成します。
// エラーがある場合、*LinkError 型になります。
func Link(oldname, newname string) error

// Symlinkはnewnameをoldnameへのシンボリックリンクとして作成します。
// Windowsでは、存在しないoldnameへのシンボリックリンクはファイルのシンボリックリンクとして作成されます。
// もしoldnameが後でディレクトリとして作成された場合、シンボリックリンクは機能しません。
// エラーが発生した場合は、*LinkError型のエラーになります。
func Symlink(oldname, newname string) error
<<<<<<< HEAD

// Readlinkは指定されたシンボリックリンクの宛先を返します。
// エラーがある場合、*PathError型で返されます。
func Readlink(name string) (string, error)
=======
>>>>>>> upstream/master
