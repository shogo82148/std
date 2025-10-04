// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package os

<<<<<<< HEAD
// Fdは、開いているファイルを参照する整数型のUnixファイルディスクリプタを返します。
// もしfが閉じている場合、ファイルディスクリプタは無効になります。
// もしfがガベージコレクションされる場合、ファイナライザーがファイルディスクリプタを閉じる可能性があり、
// それにより無効になります。ファイナライザーがいつ実行されるかの詳細については、
// [runtime.SetFinalizer] を参照してください。Unixシステムでは、これにより [File.SetDeadline]
// メソッドが動作しなくなります。
// ファイルディスクリプタは再利用可能であるため、返されるファイルディスクリプタは、
// fの [File.Close] メソッドを通じて、またはガベージコレクション中のそのファイナライザーによってのみ閉じることができます。
// それ以外の場合、ガベージコレクション中にファイナライザーが同じ（再利用された）番号の無関係なファイルディスクリプタを閉じる可能性があります。
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
=======
// DevNull is the name of the operating system's “null device.”
// On Unix-like systems, it is "/dev/null"; on Windows, "NUL".
const DevNull = "/dev/null"

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
// If there is an error, it will be of type [*PathError].
func Truncate(name string, size int64) error

// Remove removes the named file or (empty) directory.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func Remove(name string) error

// Link は newname を oldname ファイルのハードリンクとして作成します。
// エラーがある場合、*LinkError 型になります。
func Link(oldname, newname string) error

// Symlinkはnewnameをoldnameへのシンボリックリンクとして作成します。
// Windowsでは、存在しないoldnameへのシンボリックリンクはファイルのシンボリックリンクとして作成されます。
// もしoldnameが後でディレクトリとして作成された場合、シンボリックリンクは機能しません。
// エラーが発生した場合は、*LinkError型のエラーになります。
func Symlink(oldname, newname string) error
