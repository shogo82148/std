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

// NewFileは、指定されたファイルディスクリプタと名前で新しい [File] を返します。
// fdが有効なファイルディスクリプタでない場合、返される値はnilになります。
//
// NewFileの挙動はプラットフォームによって異なります：
//
//   - Unixでは、fdがノンブロッキングモードの場合、NewFileはポーリング可能なファイルを返そうとします。
//   - Windowsでは、fdが非同期I/O用にオープンされている場合（つまり、[syscall.FILE_FLAG_OVERLAPPED]
//     が [syscall.CreateFile] 呼び出しで指定されている場合）、NewFileはGoランタイムのI/O完了ポートに
//     fdを関連付けてポーリング可能なファイルを返そうとします。
//     関連付けに失敗した場合、I/O操作は同期的に実行されます。
//
// ポーリング可能なファイルのみが [File.SetDeadline]、[File.SetReadDeadline]、[File.SetWriteDeadline] をサポートします。
//
// fdをNewFileに渡した後は、[File.Fd] のコメントで説明されているのと同じ条件下でfdが無効になる場合があり、同じ制約が適用されます。
func NewFile(fd uintptr, name string) *File

// ReadはFileから最大len(b)バイトを読み取り、bに格納します。
// 読み取ったバイト数と発生したエラーを返します。
// ファイルの終端では、Readは0, io.EOFを返します。
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
// ファイルが [O_APPEND] フラグ付きでオープンされている場合、WriteAt はエラーを返します。
func (f *File) WriteAt(b []byte, off int64) (n int, err error)

// WriteToは、io.WriterToのWriteToメソッドを実装します。
func (f *File) WriteTo(w io.Writer) (n int64, err error)

// Seekは、次回のReadまたはWriteのためのファイルオフセットをoffsetに設定します。
// whenceの値によって解釈が異なります：0はファイルの先頭から、1は現在のオフセットから、2はファイルの末尾からの相対値です。
// 新しいオフセット値と、エラーがあればそれを返します。
// [O_APPEND] フラグ付きでオープンされたファイルに対するSeekの挙動は未定義です。
func (f *File) Seek(offset int64, whence int) (ret int64, err error)

// WriteStringはWriteと似ていますが、バイトのスライスではなく、文字列sの内容を書き込みます。
func (f *File) WriteString(s string) (n int, err error)

// Mkdirは指定された名前とパーミッションビット（umask適用前）で新しいディレクトリを作成します。
// エラーが発生した場合、それは[*PathError]型になります。
func Mkdir(name string, perm FileMode) error

// Chdirはカレントワーキングディレクトリを指定されたディレクトリに変更します。
// エラーが発生した場合、それは [*PathError] 型になります。
func Chdir(dir string) error

// Openは指定された名前のファイルを読み込み用にオープンします。
// 成功した場合、返されたファイルのメソッドで読み取りが可能です。
// 関連付けられたファイルディスクリプタは [O_RDONLY] モードです。
// エラーが発生した場合、それは [*PathError] 型になります。
func Open(name string) (*File, error)

// Createは指定された名前のファイルを作成または切り詰めます。
// ファイルが既に存在する場合は切り詰められます。存在しない場合はモード0o666（umask適用前）で作成されます。
// 成功した場合、返されたFileのメソッドでI/Oが可能です。関連付けられたファイルディスクリプタは [O_RDWR] モードです。
// ファイルを含むディレクトリは既に存在している必要があります。
// エラーが発生した場合、それは [*PathError] 型になります。
func Create(name string) (*File, error)

// OpenFileは汎用的なオープン呼び出しです。ほとんどのユーザーはOpenまたはCreateを使用します。
// 指定されたフラグ（[O_RDONLY] など）でファイルをオープンします。
// ファイルが存在せず、[O_CREATE] フラグが指定されている場合は、perm（umask適用前）で作成されます。
// ディレクトリは既に存在している必要があります。
// 成功した場合、返されたFileのメソッドでI/Oが可能です。
// エラーが発生した場合、それは [*PathError] 型になります。
func OpenFile(name string, flag int, perm FileMode) (*File, error)

// Renameはoldpathをnewpathに名前を変更（移動）します。
// newpathが既に存在していてディレクトリではない場合、Renameはそれを置き換えます。
// oldpathとnewpathが異なるディレクトリにある場合、OS固有の制限が適用される場合があります。
// 同じディレクトリ内でも、非UnixプラットフォームではRenameはアトミックな操作ではありません。
// エラーが発生した場合、それは*LinkErrorの型である可能性があります。
func Rename(oldpath, newpath string) error

// Readlinkは指定されたシンボリックリンクのリンク先を返します。
// エラーが発生した場合、それは [*PathError] 型になります。
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

// Chmodは指定されたファイルのモードをmodeに変更します。
// ファイルがシンボリックリンクの場合は、リンク先のモードを変更します。
// エラーが発生した場合、それは [*PathError] 型になります。
//
// オペレーティングシステムによって使用されるモードビットのサブセットが異なります。
//
// Unixでは、モードのパーミッションビット [ModeSetuid]、[ModeSetgid]、および [ModeSticky] が使用されます。
//
// Windowsでは、モードの0o200ビット（所有者書き込み可能）のみが使用されます。
// これにより、ファイルの読み取り専用属性が設定されるかクリアされるかが制御されます。
// その他のビットは現在未使用です。
// Go 1.12以前との互換性を保つために、ゼロ以外のモードを使用してください。
// 読み取り専用ファイルにはモード0o400、読み書き可能なファイルにはモード0o600を使用します。
//
// Plan 9では、モードのパーミッションビット [ModeAppend]、[ModeExclusive]、および [ModeTemporary] が使用されます。
func Chmod(name string, mode FileMode) error

// Chmodはファイルのモードをmodeに変更します。
// エラーが発生した場合、それは [*PathError] 型になります。
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

// Fdは、オープンされているファイルを参照するシステムのファイルディスクリプタまたはハンドルを返します。
// fがクローズされている場合、ディスクリプタは無効になります。
// fがガベージコレクトされた場合、ファイナライザによってディスクリプタがクローズされ、無効になることがあります。
// ファイナライザがいつ実行されるかについては [runtime.SetFinalizer] を参照してください。
//
// 返されたディスクリプタをクローズしないでください。後でfをクローズした際に、無関係なディスクリプタがクローズされる可能性があります。
//
// Fdの挙動はプラットフォームによって異なります：
//
//   - UnixおよびWindowsでは、[File.SetDeadline] メソッドが動作しなくなります。
//   - Windowsでは、ファイルディスクリプタはGoランタイムのI/O完了ポートから切り離されます。
//     これはファイルに対して同時I/O操作がない場合に発生します。
//
// ほとんどの場合、f.SyscallConnメソッドの使用を推奨します。
func (f *File) Fd() uintptr

// DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.
//
// ただし、DirFS("/prefix")は、オペレーティングシステムへのOpen呼び出しが常に"/prefix"で始まることを保証するだけです。
// つまり、DirFS("/prefix").Open("file")はos.Open("/prefix/file")と同じです。
// よって、/prefix/fileが/prefixツリーの外部を指すシンボリックリンクである場合、DirFSを使用してもos.Openを使用してもアクセスが止まるわけではありません。
// また、相対パスの場合、fs.FSのルート（DirFS("prefix")で返されるもの）は、後続のChdir呼び出しの影響を受けます。
// したがって、ディレクトリツリーに任意のコンテンツが含まれる場合、DirFSは一般的なchrootスタイルのセキュリティメカニズムの代替ではありません。
//
// [Root.FS] を使用して、シンボリックリンクによるツリーからの脱出を防ぐ fs.FS を取得してください。
//
// ディレクトリ dir は "" であってはなりません。
//
// 結果は [io/fs.StatFS]、[io/fs.ReadFileFS]、[io/fs.ReadDirFS]、および [io/fs.ReadLinkFS] を実装します。
func DirFS(dir string) fs.FS

var _ fs.StatFS = dirFS("")
var _ fs.ReadFileFS = dirFS("")
var _ fs.ReadDirFS = dirFS("")
var _ fs.ReadLinkFS = dirFS("")

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
func ReadFile(name string) ([]byte, error)

// WriteFileはデータを指定されたファイルに書き込みます。必要に応じて新規作成されます。
// ファイルが存在しない場合、WriteFileはパーミッションperm（umaskの前に）で作成します。
// ファイルが存在する場合、WriteFileは書き込み前にファイルを切り詰め、パーミッションは変更しません。
// WriteFileは複数のシステムコールが必要なため、途中で失敗するとファイルは一部だけ書き込まれた状態になる可能性があります。
func WriteFile(name string, data []byte, perm FileMode) error
