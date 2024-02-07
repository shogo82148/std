// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package netは、TCP/IP、UDP、ドメイン名の解決、およびUnixドメインソケットなど、ネットワークI/Oのためのポータブルなインターフェースを提供します。

<<<<<<< HEAD
このパッケージは、低レベルのネットワーキングプリミティブへのアクセスを提供していますが、ほとんどのクライアントは、Dial、Listen、Accept関数と関連するConnとListenerインターフェースが提供する基本的なインターフェースだけを必要とします。crypto/tlsパッケージは、同じインターフェースと似たようなDialとListen関数を使用します。
=======
Although the package provides access to low-level networking
primitives, most clients will need only the basic interface provided
by the [Dial], [Listen], and Accept functions and the associated
[Conn] and [Listener] interfaces. The crypto/tls package uses
the same interfaces and similar Dial and Listen functions.
>>>>>>> upstream/release-branch.go1.22

Dial関数はサーバーに接続します：

	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// エラーを処理する
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...

Listen関数はサーバーを作成します：

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// エラーを処理する
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// エラーを処理する
		}
		go handleConnection(conn)
	}

# ネーム解決

<<<<<<< HEAD
ネーム解決の方法は、間接的にDialのような関数を使うか、LookupHostやLookupAddrのような関数を直接使うかによって、オペレーティングシステムによって異なります。
=======
The method for resolving domain names, whether indirectly with functions like Dial
or directly with functions like [LookupHost] and [LookupAddr], varies by operating system.
>>>>>>> upstream/release-branch.go1.22

Unixシステムでは、名前を解決するための2つのオプションがあります。
/etc/resolv.confにリストされているサーバーに直接DNSリクエストを送信する純粋なGoリゾルバを使用するか、
getaddrinfoやgetnameinfoなどのCライブラリのルーチンを呼び出すcgoベースのリゾルバを使用するか、です。

デフォルトでは、ブロックされたDNSリクエストはゴルーチンのみを消費するため、純粋なGoリゾルバが使用されますが、
ブロックされたC呼び出しはオペレーティングシステムのスレッドを消費します。
cgoが利用可能な場合、さまざまな条件下でcgoベースのリゾルバが代わりに使用されます：
直接DNSリクエストを行うことができないシステム（OS X）や、
LOCALDOMAIN環境変数が存在する（空であっても）、
RES_OPTIONSまたはHOSTALIASES環境変数が空でない、
ASR_CONFIG環境変数が空でない（OpenBSDのみ）、
/etc/resolv.confまたは/etc/nsswitch.confで、
Goリゾルバが実装していない機能の使用が指定されている場合、
検索対象となる名前が.localで終わるか、mDNS名である場合です。

リゾルバの判断は、GODEBUG環境変数（パッケージruntimeを参照）のnetdns値をgoまたはcgoに設定することで上書きすることもできます。例：

	export GODEBUG=netdns=go    # 純粋なGoリゾルバを強制する
	export GODEBUG=netdns=cgo   # ネイティブリゾルバを強制する（cgo、win32）

この判断は、Goソースツリーをビルドする際にも、netgoまたはnetcgoビルドタグを設定することで強制することができます。

GODEBUG=netdns=1のような数値のnetdns設定は、リゾルバが自身の判断に関するデバッグ情報を出力します。
特定のリゾルバを強制すると同時にデバッグ情報を出力するには、
2つの設定をプラス記号で結合します。例：GODEBUG=netdns=go+1。

macOSでは、netパッケージを使用するGoコードが-buildmode=c-archiveでビルドされる場合、
生成されたアーカイブをCプログラムにリンクする際に、Cコードをリンクするときに-lresolvを渡す必要があります。

Plan 9では、リゾルバは常に/net/csと/net/dnsにアクセスします。

Windowsでは、Go 1.18.x以前では、GetAddrInfoやDnsQueryなどのCライブラリ関数をリゾルバが常に使用していました。
*/
package net

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// Addrはネットワークのエンドポイントアドレスを表します。
//
<<<<<<< HEAD
// 2つのメソッドNetworkおよびStringは、通常、Dialの引数として渡すことができる文字列を返しますが、
// その文字列の形式や意味については実装次第です。
=======
// The two methods [Addr.Network] and [Addr.String] conventionally return strings
// that can be passed as the arguments to [Dial], but the exact form
// and meaning of the strings is up to the implementation.
>>>>>>> upstream/release-branch.go1.22
type Addr interface {
	Network() string
	String() string
}

// Connは汎用のストリーム指向のネットワーク接続です。
//
// 複数のgoroutineがConn上のメソッドを同時に呼び出すことができます。
type Conn interface {
	Read(b []byte) (n int, err error)

	Write(b []byte) (n int, err error)

	Close() error

	LocalAddr() Addr

	RemoteAddr() Addr

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}

// PacketConnは汎用のパケット指向のネットワーク接続です。
//
// 複数のゴルーチンは同時にPacketConnのメソッドを呼び出すことができます。
type PacketConn interface {
	ReadFrom(p []byte) (n int, addr Addr, err error)

	WriteTo(p []byte, addr Addr) (n int, err error)

	Close() error

	LocalAddr() Addr

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}

// リスナーはストリーム指向のプロトコルのための汎用のネットワークリスナーです。
//
// 複数のゴルーチンが同時にリスナーのメソッドを呼び出すことができます。
type Listener interface {
	Accept() (Conn, error)

	Close() error

	Addr() Addr
}

// Errorはネットワークエラーを表します。
type Error interface {
	error
	Timeout() bool

	Temporary() bool
}

// OpErrorに含まれる様々なエラー。
var (
	ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
)

// OpErrorは通常、netパッケージの関数によって返されるエラータイプです。これは操作、ネットワークタイプ、およびエラーのアドレスを説明します。
type OpError struct {

	// Opはエラーの原因となった操作であり、
	// "read"または"write"などの操作です。
	Op string

	// Netはこのエラーが発生したネットワークの種類です。
	// 例えば、"tcp"や"udp6"などがあります。
	Net string

	// リモートネットワーク接続に関する操作（Dial、Read、またはWriteなど）において、ソースは対応するローカルネットワークアドレスです。
	Source Addr

	// Addrはこのエラーが発生したネットワークアドレスです。
	// ListenやSetDeadlineなどのローカルな操作の場合、Addrは操作されるローカルエンドポイントのアドレスです。
	// Dial、Read、またはWriteなどのリモートネットワーク接続に関する操作の場合、Addrはその接続のリモートアドレスです。
	Addr Addr

	// Errは操作中に発生したエラーです。
	// Errorメソッドはエラーがnilの場合にパニックを起こします。
	Err error
}

func (e *OpError) Unwrap() error

func (e *OpError) Error() string

func (e *OpError) Timeout() bool

func (e *OpError) Temporary() bool

// ParseErrorは文字列形式のネットワークアドレスパーサーのエラータイプです。
type ParseError struct {

	// Typeは期待される文字列のタイプです。例えば、「IPアドレス」、「CIDRアドレス」などです。
	Type string

	// Textは不正なテキスト文字列です。
	Text string
}

func (e *ParseError) Error() string

func (e *ParseError) Timeout() bool
func (e *ParseError) Temporary() bool

type AddrError struct {
	Err  string
	Addr string
}

func (e *AddrError) Error() string

func (e *AddrError) Timeout() bool
func (e *AddrError) Temporary() bool

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string
func (e UnknownNetworkError) Timeout() bool
func (e UnknownNetworkError) Temporary() bool

type InvalidAddrError string

func (e InvalidAddrError) Error() string
func (e InvalidAddrError) Timeout() bool
func (e InvalidAddrError) Temporary() bool

// DNSConfigErrorは、マシンのDNS設定を読み取る際のエラーを表します。
// （使用されていませんが、互換性のために保持されています。）
type DNSConfigError struct {
	Err error
}

func (e *DNSConfigError) Unwrap() error
func (e *DNSConfigError) Error() string
func (e *DNSConfigError) Timeout() bool
func (e *DNSConfigError) Temporary() bool

// DNSErrorはDNSの検索エラーを表します。
type DNSError struct {
	Err         string
	Name        string
	Server      string
	IsTimeout   bool
	IsTemporary bool

	// IsNotFound is set to true when the requested name does not
	// contain any records of the requested type (data not found),
	// or the name itself was not found (NXDOMAIN).
	IsNotFound bool
}

func (e *DNSError) Error() string

<<<<<<< HEAD
// TimeoutはDNSの検索がタイムアウトしたかどうかを報告します。
// これは常に正確にはわかりません。DNSの検索はタイムアウトにより失敗する場合があり、Timeoutがfalseを返すDNSErrorが返されることがあります。
func (e *DNSError) Timeout() bool

// Temporaryは、DNSエラーが一時的であるかどうかを示す。
// これは常にわかるわけではない。一時的なエラーによりDNS検索が失敗し、Temporaryがfalseを返す場合がある。
=======
// Timeout reports whether the DNS lookup is known to have timed out.
// This is not always known; a DNS lookup may fail due to a timeout
// and return a [DNSError] for which Timeout returns false.
func (e *DNSError) Timeout() bool

// Temporary reports whether the DNS error is known to be temporary.
// This is not always known; a DNS lookup may fail due to a temporary
// error and return a [DNSError] for which Temporary returns false.
>>>>>>> upstream/release-branch.go1.22
func (e *DNSError) Temporary() bool

// ErrClosedは、既に閉じられたネットワーク接続またはI/Oが完了する前に他のゴルーチンによって閉じられたネットワーク接続上のI/O呼び出しによって返されるエラーです。これは他のエラーに包まれる場合があり、通常はerrors.Is(err, net.ErrClosed)を使用してテストする必要があります。
var ErrClosed error = errClosed

// Buffersは、書き込むバイトのゼロ以上のランを含んでいます。
//
// ある種の接続に対して、特定のマシンでは、これはOS固有のバッチ書き込み操作（"writev"など）に最適化されます。
type Buffers [][]byte

var (
	_ io.WriterTo = (*Buffers)(nil)
	_ io.Reader   = (*Buffers)(nil)
)

// WriteTo はバッファの内容を w に書き込みます。
//
<<<<<<< HEAD
// WriteTo は、Buffeers の io.WriterTo を実装します。
=======
// WriteTo implements [io.WriterTo] for [Buffers].
>>>>>>> upstream/release-branch.go1.22
//
// WriteTo は、0 <= i < len(v) の範囲の v[i] およびスライス v を変更しますが、v[i][j] (i, j は任意の値) は変更しません。
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)

// バッファから読み込む。
//
<<<<<<< HEAD
// Read はバッファのために io.Reader を実装します。
=======
// Read implements [io.Reader] for [Buffers].
>>>>>>> upstream/release-branch.go1.22
//
// Read はスライス v と v[i]（ただし、0 <= i < len(v)）を変更しますが、
// v[i][j]（ただし、任意の i, j）は変更しません。
func (v *Buffers) Read(p []byte) (n int, err error)
