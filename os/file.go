// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// Package osは、オペレーティングシステムの機能に対するプラットフォーム非依存のインターフェースを提供します。
// 設計はUnixライクですが、エラーハンドリングはGoのようです。失敗する呼び出しは、エラーナンバーではなくエラー型の値を返します。
// エラーには、より詳細な情報が含まれることがよくあります。たとえば、ファイル名を受け取る呼び出し（OpenやStatなど）が失敗する場合、
// エラーメッセージには失敗したファイル名が含まれ、その型は*PathErrorで、さらなる情報を抽出できます。
// osインターフェースは、すべてのオペレーティングシステムで統一されたものとすることを意図しています。
// 一般的に利用できない機能は、システム固有のパッケージsyscallに現れます。
// 以下に、ファイルを開いて一部を読み込む簡単な例を示します。
=======
// Package os provides a platform-independent interface to operating system
// functionality. The design is Unix-like, although the error handling is
// Go-like; failing calls return values of type error rather than error numbers.
// Often, more information is available within the error. For example,
// if a call that takes a file name fails, such as [Open] or [Stat], the error
// will include the failing file name when printed and will be of type
// [*PathError], which may be unpacked for more information.
>>>>>>> upstream/master
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
// 注意: File上の同時操作の最大数は、OSまたはシステムによって制限される場合があります。数は大きくするべきですが、それを超えるとパフォーマンスが低下したり他の問題が発生する可能性があります。
package os

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// NameはOpenに渡されたファイルの名前を返します。
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

// ReadはFileから最大len(b)バイトを読み込み、bに格納します。
// 読み込まれたバイト数とエラーがあればそれを返します。
// ファイルの末尾では、Readは0とio.EOFを返します。
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
// ファイルがO_APPENDフラグで開かれている場合、WriteAtはエラーを返します。
func (f *File) WriteAt(b []byte, off int64) (n int, err error)

// WriteToは、io.WriterToのWriteToメソッドを実装します。
func (f *File) WriteTo(w io.Writer) (n int64, err error)

// Seekは、オフセットをオフセットに設定します。オフセットは、whenceによって解釈されます。
// whenceの解釈は次のとおりです：0はファイルの原点に対する相対的なオフセット、1は現在のオフセットに対する相対的なオフセット、2は終端に対する相対的なオフセットを意味します。
// エラーがあれば、新しいオフセットとエラーを返します。
// O_APPENDで開かれたファイルに対するSeekの振る舞いは指定されていません。
func (f *File) Seek(offset int64, whence int) (ret int64, err error)

// WriteStringはWriteと似ていますが、バイトのスライスではなく、文字列sの内容を書き込みます。
func (f *File) WriteString(s string) (n int, err error)

func Mkdir(name string, perm FileMode) error

// Chdirは現在の作業ディレクトリを指定されたディレクトリに変更します。
// エラーが発生した場合、*PathError型になります。
func Chdir(dir string) error

// Openは指定されたファイルを読み取り用に開きます。成功した場合、
// 返されたファイルのメソッドを使用して読み取りができます。
// 関連付けられたファイルディスクリプタはO_RDONLYのモードで持ちます。
// エラーが発生した場合、*PathError型のエラーが返されます。
func Open(name string) (*File, error)

// Createは指定されたファイルを作成または切り詰めます。ファイルが既に存在する場合、ファイルは切り詰められます。ファイルが存在しない場合、モード0666（umaskの前）で作成されます。成功した場合、返されたFileのメソッドを使用してI/Oを行うことができます。関連付けられたファイルディスクリプタはO_RDWRモードになります。エラーが発生した場合、*PathError型のエラーとなります。
func Create(name string) (*File, error)

// OpenFileは一般化されたオープンコールであり、ほとんどのユーザーは代わりにOpenまたはCreateを使用します。指定されたフラグ（O_RDONLYなど）で指定された名前のファイルを開きます。ファイルが存在しない場合、O_CREATEフラグが渡されると、モード許可（umask前）で作成されます。成功すると、返されたFileのメソッドを使用してI/Oが可能です。エラーが発生した場合、*PathErrorのタイプになります。
func OpenFile(name string, flag int, perm FileMode) (*File, error)

// Renameはoldpathをnewpathに名前を変更（移動）します。
// newpathが既に存在していてディレクトリではない場合、Renameはそれを置き換えます。
// oldpathとnewpathが異なるディレクトリにある場合、OS固有の制限が適用される場合があります。
// 同じディレクトリ内でも、非UnixプラットフォームではRenameはアトミックな操作ではありません。
// エラーが発生した場合、それは*LinkErrorの型である可能性があります。
func Rename(oldpath, newpath string) error

// Readlinkは、指定されたシンボリックリンクの宛先を返します。
// エラーがある場合、そのタイプは*PathErrorになります。
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

// UserCacheDirは、ユーザー固有のキャッシュデータのデフォルトのルートディレクトリを返します。ユーザーは、このディレクトリ内に独自のアプリケーション固有のサブディレクトリを作成し、それを使用する必要があります。
// Unixシステムでは、これは$XDG_CACHE_HOME（https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.htmlで指定されています）が空でない場合には、$HOME/.cacheを返します。
// Darwinでは、これは$HOME/Library/Cachesを返します。
// Windowsでは、これは%LocalAppData%を返します。
// Plan 9では、これは$home/lib/cacheを返します。
// 位置を特定できない場合（たとえば、$HOMEが定義されていない場合）は、エラーが返されます。
func UserCacheDir() (string, error)

// UserConfigDirは、ユーザー固有の設定データに使用するデフォルトのルートディレクトリを返します。ユーザーは、このディレクトリ内に自分自身のアプリケーション固有のサブディレクトリを作成し、それを使用するべきです。
// Unixシステムでは、$XDG_CONFIG_HOMEが空でない場合は、https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.htmlで指定されているようにそれを返し、それ以外の場合は$HOME/.configを返します。
// Darwinでは、$HOME/Library/Application Supportを返します。
// Windowsでは、%AppData%を返します。
// Plan 9では、$home/libを返します。
// 場所を特定できない場合（たとえば、$HOMEが定義されていない場合）は、エラーを返します。
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

// Chmodは指定されたファイルのモードを変更します。
// もしファイルがシンボリックリンクであれば、リンクのターゲットのモードを変更します。
// エラーが発生した場合は、*PathError型になります。
//
// オペレーティングシステムによって使用されるモードビットのサブセットが異なります。
//
// Unixでは、モードのパーミッションビットであるModeSetuid、ModeSetgid、およびModeStickyが使用されます。
//
// Windowsでは、モードの0200ビット（所有者書き込み可能）のみが使用されます。これにより、ファイルの読み取り専用属性が設定されるかクリアされるかが制御されます。
// その他のビットは現在未使用です。Go 1.12以前との互換性を保つために、ゼロ以外のモードを使用してください。読み取り専用ファイルにはモード0400、読み書き可能なファイルにはモード0600を使用します。
//
// Plan 9では、モードのパーミッションビットであるModeAppend、ModeExclusive、およびModeTemporaryが使用されます。
func Chmod(name string, mode FileMode) error

// Chmodはファイルのモードをmodeに変更します。
// エラーが発生した場合、それは*PathError型です。
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

// DirFSはディレクトリdirをルートとするファイルツリーのファイルシステム（fs.FS）を返します。
//
// ただし、DirFS("/prefix")は、オペレーティングシステムへのOpen呼び出しが常に"/prefix"で始まることを保証するだけです。
// つまり、DirFS("/prefix").Open("file")はos.Open("/prefix/file")と同じです。
// よって、/prefix/fileが/prefixツリーの外部を指すシンボリックリンクである場合、DirFSを使用してもos.Openを使用してもアクセスが止まるわけではありません。
// また、相対パスの場合、fs.FSのルート（DirFS("prefix")で返されるもの）は、後続のChdir呼び出しの影響を受けます。
// したがって、ディレクトリツリーに任意のコンテンツが含まれる場合、DirFSは一般的なchrootスタイルのセキュリティメカニズムの代替ではありません。
//
// ディレクトリdirは空ではありません。
//
// 結果は[io/fs.StatFS]、[io/fs.ReadFileFS]、[io/fs.ReadDirFS]を実装しています。
func DirFS(dir string) fs.FS

// ReadFileは指定されたファイルを読み込み、その内容を返します。
// 成功した呼び出しはerr == nilを返します。 err == EOFではありません。
// ReadFileはファイル全体を読み込むため、ReadからのEOFをエラーとして報告しません。
func ReadFile(name string) ([]byte, error)

// WriteFileはデータを指定されたファイルに書き込みます。必要に応じて新規作成されます。
// ファイルが存在しない場合、WriteFileはパーミッションperm（umaskの前に）で作成します。
// ファイルが存在する場合、WriteFileは書き込み前にファイルを切り詰め、パーミッションは変更しません。
// WriteFileは複数のシステムコールが必要なため、途中で失敗するとファイルは一部だけ書き込まれた状態になる可能性があります。
func WriteFile(name string, data []byte, perm FileMode) error
