// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package syscall

import (
	"github.com/shogo82148/std/unsafe"
)

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

// Errnoは、エラー状態を記述する符号なしの数値です。
// これはerrorインターフェースを実装します。ゼロのErrnoは慣例により
// エラーではないため、Errnoからerrorに変換するコードは次のようにするべきです:
//
//	err = nil
//	if errno != 0 {
//		err = errno
//	}
//
// Errnoの値は、[errors.Is] を使用してエラー値と比較できます。
// 例えば:
//
//	_, _, err := syscall.Syscall(...)
//	if errors.Is(err, fs.ErrNotExist) ...
type Errno uintptr

func (e Errno) Error() string

func (e Errno) Is(target error) bool

func (e Errno) Temporary() bool

func (e Errno) Timeout() bool

// Signalは、プロセスシグナルを記述する数値です。
// これは [os.Signal] インターフェースを実装します。
type Signal int

func (s Signal) Signal()

func (s Signal) String() string

func Read(fd int, p []byte) (n int, err error)

func Write(fd int, p []byte) (n int, err error)

func Pread(fd int, p []byte, offset int64) (n int, err error)

func Pwrite(fd int, p []byte, offset int64) (n int, err error)

// テスト用: クライアントはこのフラグを設定することで、
// IPv6ソケットの作成で強制的に [EAFNOSUPPORT] を返すようにできます。
var SocketDisableIPv6 bool

type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len _Socklen, err error)
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

type SockaddrUnix struct {
	Name string
	raw  RawSockaddrUnix
}

func Bind(fd int, sa Sockaddr) (err error)

func Connect(fd int, sa Sockaddr) (err error)

func Getpeername(fd int) (sa Sockaddr, err error)

func GetsockoptInt(fd, level, opt int) (value int, err error)

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error)

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error)

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error)

func Sendto(fd int, p []byte, flags int, to Sockaddr) (err error)

func SetsockoptByte(fd, level, opt int, value byte) (err error)

func SetsockoptInt(fd, level, opt int, value int) (err error)

func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error)

func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) (err error)

func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) (err error)

func SetsockoptICMPv6Filter(fd, level, opt int, filter *ICMPv6Filter) error

func SetsockoptLinger(fd, level, opt int, l *Linger) (err error)

func SetsockoptString(fd, level, opt int, s string) (err error)

func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (err error)

func Socket(domain, typ, proto int) (fd int, err error)

func Socketpair(domain, typ, proto int) (fd [2]int, err error)

// Sendfileは、ファイルディスクリプタinfdからファイルディスクリプタoutfdに
// 最大countバイトをコピーします。
//
// これはsendfileシステムコールをラップします。動作はオペレーティングシステムによって異なり、
// 特にoffset引数と部分的な書き込みの報告方法に関して異なります。
//
// Linuxでは、offsetがnilの場合、Sendfileはinfdの現在のファイル位置を使用して更新します。
// offsetがnilでない場合、現在のファイル位置は変更されず、offsetポインタが
// 書き込まれたバイト数を反映するように更新されます。
// nilでないエラーは通常、バイトが書き込まれなかったことを意味します。
//
// BSD派生システム（macOSを含む）では、offsetがnilの場合、Sendfileはパニックします。
// offset引数はシステムコールによって更新されないため、呼び出し元は書き込まれたバイト数を
// 使用して手動でoffsetを更新する必要があります。これらのシステムは、
// ゼロでないバイト数とエラー（例えばEAGAIN）を同時に返す場合があります。
//
// 正確なセマンティクスについては、システムのドキュメント（例えばman 2 sendfile）を参照してください。
func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
