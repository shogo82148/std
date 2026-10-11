// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows system calls.

package syscall

import (
	"github.com/shogo82148/std/unsafe"
)

type Handle uintptr

const InvalidHandle = ^Handle(0)

// StringToUTF16は、UTF-8文字列sのUTF-16エンコーディングに終端のNULを追加したものを返します。
// sにNULバイトが含まれている場合、この関数はエラーを返す代わりにパニックします。
//
// Deprecated: 代わりに[UTF16FromString]を使用してください。
func StringToUTF16(s string) []uint16

// UTF16FromStringは、UTF-8文字列sのUTF-16エンコーディングに終端のNULを追加したものを返します。
// sのいずれかの位置にNULバイトが含まれている場合、(nil, [EINVAL])を返します。
// 対になっていないサロゲートはWTF-8を使用してエンコードされます。
func UTF16FromString(s string) ([]uint16, error)

// UTF16ToStringは、UTF-16シーケンスsのUTF-8エンコーディングから終端のNULを
// 取り除いたものを返します。対になっていないサロゲートは、UTF-8の代わりに
// WTF-8を使用してデコードされます。
func UTF16ToString(s []uint16) string

// StringToUTF16Ptrは、UTF-8文字列sのUTF-16エンコーディングに終端のNULを
// 追加したものへのポインタを返します。sにNULバイトが含まれている場合、
// この関数はエラーを返す代わりにパニックします。
//
// Deprecated: 代わりに[UTF16PtrFromString]を使用してください。
func StringToUTF16Ptr(s string) *uint16

// UTF16PtrFromStringは、UTF-8文字列sのUTF-16エンコーディングに終端のNULを
// 追加したものへのポインタを返します。sのいずれかの位置にNULバイトが
// 含まれている場合、(nil, EINVAL)を返します。
// 対になっていないサロゲートはWTF-8を使用してエンコードされます。
func UTF16PtrFromString(s string) (*uint16, error)

// ErrnoはWindowsのエラー番号です。
//
// Errnoの値は、[errors.Is]を使用してエラー値と比較できます。
// 例えば:
//
//	_, _, err := syscall.Syscall(...)
//	if errors.Is(err, fs.ErrNotExist) ...
type Errno uintptr

// FormatMessageは非推奨です（msgsrcはuint32ではなくuintptrであるべきですが、
// Go 1の互換性保証により変更できません）。
//
// Deprecated: 代わりにgolang.org/x/sys/windowsのFormatMessageを使用してください。
func FormatMessage(flags uint32, msgsrc uint32, msgid uint32, langid uint32, buf []uint16, args *byte) (n uint32, err error)

func (e Errno) Error() string

func (e Errno) Is(target error) bool

func (e Errno) Temporary() bool

func (e Errno) Timeout() bool

// NewCallbackは、Goの関数をstdcall呼び出し規約に準拠した関数ポインタに変換します。
// これは、コールバックを必要とするWindowsのコードと相互運用する場合に便利です。
// 引数は、uintptrサイズの結果を1つ持つ関数であることが期待されます。関数はuintptrのサイズより大きいサイズの引数を持ってはいけません。
// 単一のGoプロセスで作成できるコールバックの数は限られており、これらのコールバックのために
// 割り当てられたメモリが解放されることはありません。
// NewCallbackとNewCallbackCDeclを合わせて、少なくとも1024個のコールバックを常に作成できます。
func NewCallback(fn any) uintptr

// NewCallbackCDeclは、Goの関数をcdecl呼び出し規約に準拠した関数ポインタに変換します。
// これは、コールバックを必要とするWindowsのコードと相互運用する場合に便利です。
// 引数は、uintptrサイズの結果を1つ持つ関数であることが期待されます。関数はuintptrのサイズより大きいサイズの引数を持ってはいけません。
// 単一のGoプロセスで作成できるコールバックの数は限られており、これらのコールバックのために
// 割り当てられたメモリが解放されることはありません。
// NewCallbackとNewCallbackCDeclを合わせて、少なくとも1024個のコールバックを常に作成できます。
func NewCallbackCDecl(fn any) uintptr

func Open(name string, flag int, perm uint32) (fd Handle, err error)

func Read(fd Handle, p []byte) (n int, err error)

func Write(fd Handle, p []byte) (n int, err error)

func ReadFile(fd Handle, p []byte, done *uint32, overlapped *Overlapped) error

func WriteFile(fd Handle, p []byte, done *uint32, overlapped *Overlapped) error

func Seek(fd Handle, offset int64, whence int) (newoffset int64, err error)

func Close(fd Handle) (err error)

var (
	Stdin  = getStdHandle(STD_INPUT_HANDLE)
	Stdout = getStdHandle(STD_OUTPUT_HANDLE)
	Stderr = getStdHandle(STD_ERROR_HANDLE)
)

const ImplementsGetwd = true

func Getwd() (wd string, err error)

func Chdir(path string) (err error)

func Mkdir(path string, mode uint32) (err error)

func Rmdir(path string) (err error)

func Unlink(path string) (err error)

func Rename(oldpath, newpath string) (err error)

func ComputerName() (name string, err error)

func Ftruncate(fd Handle, length int64) (err error)

func Gettimeofday(tv *Timeval) (err error)

func Pipe(p []Handle) (err error)

func Utimes(path string, tv []Timeval) (err error)

func UtimesNano(path string, ts []Timespec) (err error)

func Fsync(fd Handle) (err error)

func Chmod(path string, mode uint32) (err error)

func LoadCancelIoEx() error

func LoadSetFileCompletionNotificationModes() error

// テスト用: クライアントはこのフラグを設定することで、
// IPv6ソケットの作成で強制的に[EAFNOSUPPORT]を返すようにできます。
var SocketDisableIPv6 bool

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte
	Scope_id uint32
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [100]int8
}

type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len int32, err error)
}

type SockaddrInet4 struct {
	Port int
	Addr [4]byte
	raw  RawSockaddrInet4
}

type SockaddrInet6 struct {
	Port   int
	ZoneId uint32
	Addr   [16]byte
	raw    RawSockaddrInet6
}

type RawSockaddrUnix struct {
	Family uint16
	Path   [UNIX_PATH_MAX]int8
}

type SockaddrUnix struct {
	Name string
	raw  RawSockaddrUnix
}

func (rsa *RawSockaddrAny) Sockaddr() (Sockaddr, error)

func Socket(domain, typ, proto int) (fd Handle, err error)

func SetsockoptInt(fd Handle, level, opt int, value int) (err error)

func Bind(fd Handle, sa Sockaddr) (err error)

func Connect(fd Handle, sa Sockaddr) (err error)

func Getsockname(fd Handle) (sa Sockaddr, err error)

func Getpeername(fd Handle) (sa Sockaddr, err error)

func Listen(s Handle, n int) (err error)

func Shutdown(fd Handle, how int) (err error)

func WSASendto(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to Sockaddr, overlapped *Overlapped, croutine *byte) (err error)

func LoadGetAddrInfo() error

func LoadConnectEx() error

func ConnectEx(fd Handle, sa Sockaddr, sendBuf *byte, sendDataLen uint32, bytesSent *uint32, overlapped *Overlapped) error

// osパッケージが期待するものをサポートするために作られた構造体。
type Rusage struct {
	CreationTime Filetime
	ExitTime     Filetime
	KernelTime   Filetime
	UserTime     Filetime
}

type WaitStatus struct {
	ExitCode uint32
}

func (w WaitStatus) Exited() bool

func (w WaitStatus) ExitStatus() int

func (w WaitStatus) Signal() Signal

func (w WaitStatus) CoreDump() bool

func (w WaitStatus) Stopped() bool

func (w WaitStatus) Continued() bool

func (w WaitStatus) StopSignal() Signal

func (w WaitStatus) Signaled() bool

func (w WaitStatus) TrapCause() int

// TimespecはWindowsでは作られた構造体ですが、他のオペレーティングシステム向けの
// syscallパッケージとの一貫性のためにここに存在します。
type Timespec struct {
	Sec  int64
	Nsec int64
}

func TimespecToNsec(ts Timespec) int64

func NsecToTimespec(nsec int64) (ts Timespec)

func Accept(fd Handle) (nfd Handle, sa Sockaddr, err error)
func Recvfrom(fd Handle, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd Handle, p []byte, flags int, to Sockaddr) (err error)
func SetsockoptTimeval(fd Handle, level, opt int, tv *Timeval) (err error)

type Linger struct {
	Onoff  int32
	Linger int32
}

type IPMreq struct {
	Multiaddr [4]byte
	Interface [4]byte
}

type IPv6Mreq struct {
	Multiaddr [16]byte
	Interface uint32
}

func GetsockoptInt(fd Handle, level, opt int) (int, error)

func SetsockoptLinger(fd Handle, level, opt int, l *Linger) (err error)

func SetsockoptInet4Addr(fd Handle, level, opt int, value [4]byte) (err error)

func SetsockoptIPMreq(fd Handle, level, opt int, mreq *IPMreq) (err error)

func SetsockoptIPv6Mreq(fd Handle, level, opt int, mreq *IPv6Mreq) (err error)

func Getpid() (pid int)

func FindFirstFile(name *uint16, data *Win32finddata) (handle Handle, err error)

func FindNextFile(handle Handle, data *Win32finddata) (err error)

func Getppid() (ppid int)

func Fchdir(fd Handle) (err error)

// TODO(brainman): fix all needed for os
func Link(oldpath, newpath string) (err error)
func Symlink(path, link string) (err error)

func Fchmod(fd Handle, mode uint32) (err error)
func Chown(path string, uid int, gid int) (err error)
func Lchown(path string, uid int, gid int) (err error)
func Fchown(fd Handle, uid int, gid int) (err error)

func Getuid() (uid int)
func Geteuid() (euid int)
func Getgid() (gid int)
func Getegid() (egid int)
func Getgroups() (gids []int, err error)

type Signal int

func (s Signal) Signal()

func (s Signal) String() string

func LoadCreateSymbolicLink() error

// Readlinkは、指定されたシンボリックリンクのリンク先を返します。
func Readlink(path string, buf []byte) (n int, err error)

// Deprecated: CreateIoCompletionPortは関数シグネチャが誤っています。x/sys/windows.CreateIoCompletionPortを使用してください。
func CreateIoCompletionPort(filehandle Handle, cphandle Handle, key uint32, threadcnt uint32) (Handle, error)

// Deprecated: GetQueuedCompletionStatusは関数シグネチャが誤っています。x/sys/windows.GetQueuedCompletionStatusを使用してください。
func GetQueuedCompletionStatus(cphandle Handle, qty *uint32, key *uint32, overlapped **Overlapped, timeout uint32) error

// Deprecated: PostQueuedCompletionStatusは関数シグネチャが誤っています。x/sys/windows.PostQueuedCompletionStatusを使用してください。
func PostQueuedCompletionStatus(cphandle Handle, qty uint32, key uint32, overlapped *Overlapped) error

// RegEnumKeyExは、開いているレジストリキーのサブキーを列挙します。
// 各呼び出しは1つのサブキーに関する情報を取得します。nameは、
// サブキーの名前とNUL終端文字を保持するのに十分な大きさのバッファであるべきです。
// nameLenはその長さです。戻り時に、nameLenにはサブキーの実際の長さが含まれます。
//
// nameがサブキーを保持するのに十分な大きさでない場合、この関数は
// ERROR_MORE_DATAを返し、適切なサイズのバッファで再度呼び出す必要があります。
//
// reservedはnilでなければなりません。classとclassLenは、nameとnameLenと同様に
// 動作しますが、サブキーのクラスに対するものであり、省略可能です。
// lastWriteTimeがnilでない場合、サブキーが最後に書き込まれた時刻が設定されます。
//
// 呼び出し元はすべてのサブキーを順番に列挙しなければなりません。つまり、
// RegEnumKeyExは、インデックス0から開始して関数がERROR_NO_MORE_ITEMSを返すまで
// インデックスを増やしながら呼び出すか、最後のサブキーのインデックス
// （RegQueryInfoKeyから取得可能）から開始してインデックス0が列挙されるまで
// 減らしながら呼び出さなければなりません。
//
// このAPIの連続した呼び出しは同じOSスレッド上で行われなければならないため、
// この関数を呼び出す前に[runtime.LockOSThread]を呼び出してください。
func RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error)

func GetStartupInfo(startupInfo *StartupInfo) error

func CreateFile(name *uint16, access uint32, mode uint32, sa *SecurityAttributes, createmode uint32, attrs uint32, templatefile int32) (handle Handle, err error)
