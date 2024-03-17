// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package os

<<<<<<< HEAD
// Fdメソッドは、オープンされたファイルを参照する整数型のUnixファイルディスクリプタを返します。
// fがクローズされると、ファイルディスクリプタは無効になります。
// fがガベージコレクトされると、ファイナライザによってファイルディスクリプタがクローズされることがあります。
// これにより、SetDeadlineメソッドが動作しなくなる可能性があります。
// ファイルディスクリプタは再利用される可能性があるため、返されたファイルディスクリプタは
// fのCloseメソッドを通じてのみクローズするか、ガベージコレクション時のファイナライザによってクローズされます。
// そうでない場合、ガベージコレクション時にはファイナライザが同じ（再利用された）番号を持つ他のファイルディスクリプタをクローズする可能性があります。
=======
// Fd returns the integer Unix file descriptor referencing the open file.
// If f is closed, the file descriptor becomes invalid.
// If f is garbage collected, a finalizer may close the file descriptor,
// making it invalid; see [runtime.SetFinalizer] for more information on when
// a finalizer might be run. On Unix systems this will cause the [File.SetDeadline]
// methods to stop working.
// Because file descriptors can be reused, the returned file descriptor may
// only be closed through the [File.Close] method of f, or by its finalizer during
// garbage collection. Otherwise, during garbage collection the finalizer
// may close an unrelated file descriptor with the same (reused) number.
>>>>>>> upstream/master
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
