// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// osパッケージは、オペレーティングシステムの機能へのプラットフォームに依存しないインターフェースを提供します。
// デザインはUnix風ですが、エラーハンドリングはGo風です。失敗した呼び出しはエラー番号ではなく、
// error型の値を返します。しばしば、エラー内にさらに詳細な情報が含まれています。
// 例えば、ファイル名を取る呼び出しが失敗した場合、[Open] や [Stat] など、エラーは印刷時に失敗したファイル名を含み、
// [*PathError] 型になります。これはさらなる情報のためにアンパックすることができます。
//
//	file, err := os.Open("file.go") // 読み込みアクセス用。
//	if err != nil {
//		log.Fatal(err)
//	}
//
// オープンが失敗した場合、エラーメッセージは自己説明的であるようになります。
//
//	open file.go: ファイルまたはディレクトリが存在しません。
//
// ファイルのデータは、バイトのスライスに読み込むことができます。ReadとWriteは、引数のスライスの長さからバイト数を取得します。
//
//	data := make([]byte, 100)
//	count, err := file.Read(data)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("read %d bytes: %q\n", count, data[:count])
//
// # Concurrency
//
// [File]のメソッドはファイルシステムの操作に対応しています。すべてが
// 同時実行に対して安全です。Fileに対する同時操作の最大数は、OSまたはシステムによって
// 制限される可能性があります。その数は高いはずですが、それを超えるとパフォーマンスが低下したり、
// 他の問題が発生する可能性があります。
package os

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// NameはOpenに渡されたファイルの名前を返します。
//
// [Close]の後にNameを呼び出しても安全です。
func (f *File) Name() string

// Stdin、Stdout、およびStderrは、標準入力、標準出力、および標準エラーファイルディスクリプタを指すオープンファイルです。
//
// Goランタイムは、パニックやクラッシュの場合には標準エラーに書き込みます。
// Stderrを閉じると、それらのメッセージは他の場所に転送される可能性があります。
// たとえば、後で開かれるファイルに転送されるかもしれません。
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

// オープンファイル時に基になるシステムのものをラップするフラグ。すべてのフラグが与えられたシステム上で実装されているわけではありません。
const (
	// O_RDONLY、O_WRONLY、またはO_RDWRのいずれかを指定する必要があります。
	O_RDONLY int = syscall.O_RDONLY
	O_WRONLY int = syscall.O_WRONLY
	O_RDWR   int = syscall.O_RDWR
	// 残りの値はOrで結合して動作を制御できます。
	O_APPEND int = syscall.O_APPEND
	O_CREATE int = syscall.O_CREAT
	O_EXCL   int = syscall.O_EXCL
	O_SYNC   int = syscall.O_SYNC
	O_TRUNC  int = syscall.O_TRUNC
)

// 値を探す。
//
// 廃止: io.SeekStart、io.SeekCurrent、io.SeekEnd を使用してください。
const (
	SEEK_SET int = 0
	SEEK_CUR int = 1
	SEEK_END int = 2
)

// LinkErrorはリンクやシンボリックリンク、リネームのシステムコール中に発生したエラーと、それによって引き起こされたパスを記録します。
type LinkError struct {
	Op  string
	Old string
	New string
	Err error
}

func (e *LinkError) Error() string

func (e *LinkError) Unwrap() error

<<<<<<< HEAD
// ReadはFileから最大len(b)バイトを読み込み、bに格納します。
// 読み込まれたバイト数とエラーがあればそれを返します。
// ファイルの末尾では、Readは0とio.EOFを返します。
=======
// NewFile returns a new [File] with the given file descriptor and name.
// The returned value will be nil if fd is not a valid file descriptor.
//
// NewFile's behavior differs on some platforms:
//
//   - On Unix, if fd is in non-blocking mode, NewFile will attempt to return a pollable file.
//   - On Windows, if fd is opened for asynchronous I/O (that is, [syscall.FILE_FLAG_OVERLAPPED]
//     has been specified in the [syscall.CreateFile] call), NewFile will attempt to return a pollable
//     file by associating fd with the Go runtime I/O completion port.
//     The I/O operations will be performed synchronously if the association fails.
//
// Only pollable files support [File.SetDeadline], [File.SetReadDeadline], and [File.SetWriteDeadline].
//
// After passing it to NewFile, fd may become invalid under the same conditions described
// in the comments of [File.Fd], and the same constraints apply.
func NewFile(fd uintptr, name string) *File

// Read reads up to len(b) bytes from the File and stores them in b.
// It returns the number of bytes read and any error encountered.
// At end of file, Read returns 0, io.EOF.
>>>>>>> upstream/release-branch.go1.25
func (f *File) Read(b []byte) (n int, err error)

// ReadAt はオフセット off から始まる File から len(b) バイトを読み取ります。
// 読み取ったバイト数とエラー（ある場合）を返します。
// ReadAt は常に、n < len(b) の場合には非 nil のエラーを返します。
// ファイルの終端では、そのエラーは io.EOF です。
func (f *File) ReadAt(b []byte, off int64) (n int, err error)

// ReadFrom は io.ReaderFrom を実装します。
func (f *File) ReadFrom(r io.Reader) (n int64, err error)

// Writeはbからlen(b)バイトをFileに書き込みます。
// 書き込まれたバイト数とエラー（ある場合）を返します。
// n != len(b)の場合、Writeはnilでないエラーを返します。
func (f *File) Write(b []byte) (n int, err error)

// WriteAtはオフセットoffから始まるFileにlen(b)バイトを書き込みます。
// 書き込まれたバイト数とエラー（ある場合）を返します。
// n != len(b)の場合、WriteAtはnilでないエラーを返します。
//
<<<<<<< HEAD
// ファイルがO_APPENDフラグで開かれている場合、WriteAtはエラーを返します。
=======
// If file was opened with the [O_APPEND] flag, WriteAt returns an error.
>>>>>>> upstream/release-branch.go1.25
func (f *File) WriteAt(b []byte, off int64) (n int, err error)

// WriteToは、io.WriterToのWriteToメソッドを実装します。
func (f *File) WriteTo(w io.Writer) (n int64, err error)

<<<<<<< HEAD
// Seekは、オフセットをオフセットに設定します。オフセットは、whenceによって解釈されます。
// whenceの解釈は次のとおりです：0はファイルの原点に対する相対的なオフセット、1は現在のオフセットに対する相対的なオフセット、2は終端に対する相対的なオフセットを意味します。
// エラーがあれば、新しいオフセットとエラーを返します。
// O_APPENDで開かれたファイルに対するSeekの振る舞いは指定されていません。
=======
// Seek sets the offset for the next Read or Write on file to offset, interpreted
// according to whence: 0 means relative to the origin of the file, 1 means
// relative to the current offset, and 2 means relative to the end.
// It returns the new offset and an error, if any.
// The behavior of Seek on a file opened with [O_APPEND] is not specified.
>>>>>>> upstream/release-branch.go1.25
func (f *File) Seek(offset int64, whence int) (ret int64, err error)

// WriteStringはWriteと似ていますが、バイトのスライスではなく、文字列sの内容を書き込みます。
func (f *File) WriteString(s string) (n int, err error)

<<<<<<< HEAD
func Mkdir(name string, perm FileMode) error

// Chdirは現在の作業ディレクトリを指定されたディレクトリに変更します。
// エラーが発生した場合、*PathError型になります。
func Chdir(dir string) error

// Openは指定されたファイルを読み取り用に開きます。成功した場合、
// 返されたファイルのメソッドを使用して読み取りができます。
// 関連付けられたファイルディスクリプタはO_RDONLYのモードで持ちます。
// エラーが発生した場合、*PathError型のエラーが返されます。
func Open(name string) (*File, error)

// Createは指定されたファイルを作成または切り詰めます。ファイルが既に存在する場合、ファイルは切り詰められます。
// ファイルが存在しない場合、モード0o666（umaskの前）で作成されます。
// 成功した場合、返されたFileのメソッドを使用してI/Oを行うことができます。
// 関連付けられたファイルディスクリプタはO_RDWRモードになります。エラーが発生した場合、*PathError型のエラーとなります。
func Create(name string) (*File, error)

// OpenFileは一般化されたオープンコールであり、ほとんどのユーザーは代わりにOpenまたはCreateを使用します。指定されたフラグ（O_RDONLYなど）で指定された名前のファイルを開きます。ファイルが存在しない場合、O_CREATEフラグが渡されると、モード許可（umask前）で作成されます。成功すると、返されたFileのメソッドを使用してI/Oが可能です。エラーが発生した場合、*PathErrorのタイプになります。
=======
// Mkdir creates a new directory with the specified name and permission
// bits (before umask).
// If there is an error, it will be of type [*PathError].
func Mkdir(name string, perm FileMode) error

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type [*PathError].
func Chdir(dir string) error

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode [O_RDONLY].
// If there is an error, it will be of type [*PathError].
func Open(name string) (*File, error)

// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0o666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode [O_RDWR].
// The directory containing the file must already exist.
// If there is an error, it will be of type [*PathError].
func Create(name string) (*File, error)

// OpenFile is the generalized open call; most users will use Open
// or Create instead. It opens the named file with specified flag
// ([O_RDONLY] etc.). If the file does not exist, and the [O_CREATE] flag
// is passed, it is created with mode perm (before umask);
// the containing directory must exist. If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func OpenFile(name string, flag int, perm FileMode) (*File, error)

// Renameはoldpathをnewpathに名前を変更（移動）します。
// newpathが既に存在していてディレクトリではない場合、Renameはそれを置き換えます。
// oldpathとnewpathが異なるディレクトリにある場合、OS固有の制限が適用される場合があります。
// 同じディレクトリ内でも、非UnixプラットフォームではRenameはアトミックな操作ではありません。
// エラーが発生した場合、それは*LinkErrorの型である可能性があります。
func Rename(oldpath, newpath string) error

<<<<<<< HEAD
// Readlinkは、指定されたシンボリックリンクの宛先を返します。
// エラーがある場合、そのタイプは*PathErrorになります。
=======
// Readlink returns the destination of the named symbolic link.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
//
// リンク先が相対的な場合、Readlinkはそれを絶対パスに解決せずに
// 相対パスを返します。
func Readlink(name string) (string, error)

// TempDirは一時ファイルに使用するデフォルトのディレクトリを返します。
//
// Unixシステムでは、$TMPDIRが空でない場合はそれを返し、さもなくば/tmpを返します。
// Windowsでは、GetTempPathを使用し、最初の空でない値を%TMP%、%TEMP%、%USERPROFILE%、またはWindowsディレクトリから返します。
// Plan 9では、/tmpを返します。
//
// このディレクトリは、存在することやアクセス可能な許可を持っていることが保証されていません。
func TempDir() string

// UserCacheDirは、ユーザー固有のキャッシュデータに使用するデフォルトのルートディレクトリを返します。
// ユーザーはこのディレクトリ内にアプリケーション固有のサブディレクトリを作成して使用する必要があります。
//
// Unixシステムでは、https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html で指定されている
// $XDG_CACHE_HOMEが空でない場合はその値を返し、そうでない場合は$HOME/.cacheを返します。
// Darwinでは、$HOME/Library/Cachesを返します。
// Windowsでは、%LocalAppData%を返します。
// Plan 9では、$home/lib/cacheを返します。
//
// 位置を特定できない場合（例えば、$HOMEが定義されていない場合）、エラーを返します。
func UserCacheDir() (string, error)

// UserConfigDirは、ユーザー固有の設定データに使用するデフォルトのルートディレクトリを返します。
// ユーザーはこのディレクトリ内にアプリケーション固有のサブディレクトリを作成して使用する必要があります。
//
// Unixシステムでは、https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html で指定されている
// $XDG_CONFIG_HOMEが空でない場合はその値を返し、そうでない場合は$HOME/.configを返します。
// Darwinでは、$HOME/Library/Application Supportを返します。
// Windowsでは、%AppData%を返します。
// Plan 9では、$home/libを返します。
//
// 位置を特定できない場合（例えば、$HOMEが定義されていない場合）、エラーを返します。
func UserConfigDir() (string, error)

// UserHomeDirは現在のユーザーのホームディレクトリを返します。
//
// Unix（macOSを含む）では、$HOME環境変数を返します。
// Windowsでは、%USERPROFILE%を返します。
// Plan 9では、$home環境変数を返します。
//
// 環境変数に期待される変数が設定されていない場合、UserHomeDir
// は、プラットフォーム固有のデフォルト値または非nilのエラーを返します。
func UserHomeDir() (string, error)

<<<<<<< HEAD
// Chmodは指定されたファイルのモードを変更します。
// もしファイルがシンボリックリンクであれば、リンクのターゲットのモードを変更します。
// エラーが発生した場合は、*PathError型になります。
=======
// Chmod changes the mode of the named file to mode.
// If the file is a symbolic link, it changes the mode of the link's target.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
//
// オペレーティングシステムによって使用されるモードビットのサブセットが異なります。
//
<<<<<<< HEAD
// Unixでは、モードのパーミッションビットであるModeSetuid、ModeSetgid、およびModeStickyが使用されます。
=======
// On Unix, the mode's permission bits, [ModeSetuid], [ModeSetgid], and
// [ModeSticky] are used.
>>>>>>> upstream/release-branch.go1.25
//
// Windowsでは、モードの0o200ビット（所有者書き込み可能）のみが使用されます。
// これにより、ファイルの読み取り専用属性が設定されるかクリアされるかが制御されます。
// その他のビットは現在未使用です。
// Go 1.12以前との互換性を保つために、ゼロ以外のモードを使用してください。
// 読み取り専用ファイルにはモード0o400、読み書き可能なファイルにはモード0o600を使用します。
//
<<<<<<< HEAD
// Plan 9では、モードのパーミッションビットであるModeAppend、ModeExclusive、およびModeTemporaryが使用されます。
func Chmod(name string, mode FileMode) error

// Chmodはファイルのモードをmodeに変更します。
// エラーが発生した場合、それは*PathError型です。
=======
// On Plan 9, the mode's permission bits, [ModeAppend], [ModeExclusive],
// and [ModeTemporary] are used.
func Chmod(name string, mode FileMode) error

// Chmod changes the mode of the file to mode.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func (f *File) Chmod(mode FileMode) error

// SetDeadlineは、ファイルの読み取りと書き込みのデッドラインを設定します。
// SetReadDeadlineおよびSetWriteDeadlineの両方を呼び出すのと同等です。
//
// デッドラインを設定できるファイルの種類には制限があります。デッドラインをサポートしないファイルにSetDeadlineを呼び出すと、ErrNoDeadlineが返されます。
// ほとんどのシステムでは、通常のファイルはデッドラインをサポートしませんが、パイプはサポートします。
//
// デッドラインは、I/Oのブロックではなく、エラーとなる絶対時刻です。デッドラインは、未来および保留中のすべてのI/Oに適用されます。ただし、即座に次のReadまたはWrite呼び出しに適用されるわけではありません。
// デッドラインが超過された後は、将来のデッドラインを設定することで、接続をリフレッシュできます。
//
// デッドラインが超過されると、ReadまたはWriteまたは他のI/Oメソッドの呼び出しは、ErrDeadlineExceededをラップしたエラーを返します。
// これは、errors.Is(err, os.ErrDeadlineExceeded)を使用してテストできます。
// そのエラーにはTimeoutメソッドが実装されており、Timeoutメソッドを呼び出すとtrueが返ります。ただし、デッドラインが超過されていなくても、Timeoutがtrueを返す可能性がある他のエラーもあります。
//
// 成功したReadまたはWrite呼び出しの後、デッドラインを繰り返し延長することでアイドルタイムアウトを実装できます。
//
// tのゼロ値は、I/O操作がタイムアウトしないことを意味します。
func (f *File) SetDeadline(t time.Time) error

// SetReadDeadlineは、将来のRead呼び出しと現在ブロックされているRead呼び出しの締め切りを設定します。
// tのゼロ値は、Readがタイムアウトしないことを意味します。
// すべてのファイルが締め切りを設定できるわけではありません。SetDeadlineを参照してください。
func (f *File) SetReadDeadline(t time.Time) error

// SetWriteDeadlineは、将来のWrite呼び出しや現在ブロックされているWrite呼び出しの締め切りを設定します。
// Writeがタイムアウトしても、n>0が返される場合があります。
// これは、一部のデータが正常に書き込まれたことを示します。
// tのゼロ値は、Writeがタイムアウトしないことを意味します。
// すべてのファイルが締め切りを設定できるわけではありません。SetDeadlineを参照してください。
func (f *File) SetWriteDeadline(t time.Time) error

// SyscallConnは生のファイルを返します。
// これはsyscall.Connインターフェースを実装しています。
func (f *File) SyscallConn() (syscall.RawConn, error)

<<<<<<< HEAD
// DirFSはディレクトリdirをルートとするファイルツリーのファイルシステム（fs.FS）を返します。
=======
// Fd returns the system file descriptor or handle referencing the open file.
// If f is closed, the descriptor becomes invalid.
// If f is garbage collected, a finalizer may close the descriptor,
// making it invalid; see [runtime.SetFinalizer] for more information on when
// a finalizer might be run.
//
// Do not close the returned descriptor; that could cause a later
// close of f to close an unrelated descriptor.
//
// Fd's behavior differs on some platforms:
//
//   - On Unix and Windows, [File.SetDeadline] methods will stop working.
//   - On Windows, the file descriptor will be disassociated from the
//     Go runtime I/O completion port if there are no concurrent I/O
//     operations on the file.
//
// For most uses prefer the f.SyscallConn method.
func (f *File) Fd() uintptr

// DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.
>>>>>>> upstream/release-branch.go1.25
//
// ただし、DirFS("/prefix")は、オペレーティングシステムへのOpen呼び出しが常に"/prefix"で始まることを保証するだけです。
// つまり、DirFS("/prefix").Open("file")はos.Open("/prefix/file")と同じです。
// よって、/prefix/fileが/prefixツリーの外部を指すシンボリックリンクである場合、DirFSを使用してもos.Openを使用してもアクセスが止まるわけではありません。
// また、相対パスの場合、fs.FSのルート（DirFS("prefix")で返されるもの）は、後続のChdir呼び出しの影響を受けます。
// したがって、ディレクトリツリーに任意のコンテンツが含まれる場合、DirFSは一般的なchrootスタイルのセキュリティメカニズムの代替ではありません。
//
<<<<<<< HEAD
// ディレクトリdirは空ではありません。
//
// 結果は[io/fs.StatFS]、[io/fs.ReadFileFS]、[io/fs.ReadDirFS]を実装しています。
func DirFS(dir string) fs.FS

// ReadFileは指定されたファイルを読み込み、その内容を返します。
// 成功した呼び出しはerr == nilを返します。 err == EOFではありません。
// ReadFileはファイル全体を読み込むため、ReadからのEOFをエラーとして報告しません。
=======
// Use [Root.FS] to obtain a fs.FS that prevents escapes from the tree via symbolic links.
//
// The directory dir must not be "".
//
// The result implements [io/fs.StatFS], [io/fs.ReadFileFS], [io/fs.ReadDirFS], and
// [io/fs.ReadLinkFS].
func DirFS(dir string) fs.FS

var _ fs.StatFS = dirFS("")
var _ fs.ReadFileFS = dirFS("")
var _ fs.ReadDirFS = dirFS("")
var _ fs.ReadLinkFS = dirFS("")

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
>>>>>>> upstream/release-branch.go1.25
func ReadFile(name string) ([]byte, error)

// WriteFileはデータを指定されたファイルに書き込みます。必要に応じて新規作成されます。
// ファイルが存在しない場合、WriteFileはパーミッションperm（umaskの前に）で作成します。
// ファイルが存在する場合、WriteFileは書き込み前にファイルを切り詰め、パーミッションは変更しません。
// WriteFileは複数のシステムコールが必要なため、途中で失敗するとファイルは一部だけ書き込まれた状態になる可能性があります。
func WriteFile(name string, data []byte, perm FileMode) error
