// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
netパッケージは、TCP/IP、UDP、ドメイン名の解決、およびUnixドメインソケットなど、ネットワークI/Oのためのポータブルなインターフェースを提供します。

このパッケージは、低レベルのネットワーキングプリミティブへのアクセスを提供していますが、ほとんどのクライアントは、[Dial]、[Listen]、Accept関数と関連する [Conn] と [Listener] インターフェースが提供する基本的なインターフェースだけを必要とします。crypto/tlsパッケージは、同じインターフェースと似たようなDialとListen関数を使用します。

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

ネーム解決の方法は、間接的にDialのような関数を使うか、[LookupHost] や [LookupAddr] のような関数を直接使うかによって、オペレーティングシステムによって異なります。

Unixシステムでは、名前を解決するための2つのオプションがあります。
/etc/resolv.confにリストされているサーバーに直接DNSリクエストを送信する純粋なGoリゾルバを使用するか、
getaddrinfoやgetnameinfoなどのCライブラリのルーチンを呼び出すcgoベースのリゾルバを使用するか、です。

Unixでは、ブロックされたDNSリクエストがゴルーチンだけを消費するため、
純粋なGoリゾルバがcgoリゾルバよりも優先されます。一方、ブロックされたC呼び出しは
オペレーティングシステムのスレッドを消費します。
cgoが利用可能な場合、さまざまな条件下でcgoベースのリゾルバが代わりに使用されます：
プログラムが直接DNSリクエストを行うことを許可しないシステム（OS X）上、
LOCALDOMAIN環境変数が存在する場合（空でも）、
RES_OPTIONSまたはHOSTALIASES環境変数が空でない場合、
ASR_CONFIG環境変数が空でない場合（OpenBSDのみ）、
/etc/resolv.confまたは/etc/nsswitch.confがGoリゾルバが実装していない機能の使用を指定している場合。

すべてのシステム（Plan 9を除く）で、cgoリゾルバが使用されている場合、
このパッケージは並行cgoルックアップ制限を適用して、システムがシステムスレッドを使い果たすのを防ぎます。
現在、同時ルックアップは500に制限されています。

リゾルバの決定は、GODEBUG環境変数のnetdns値をgoまたはcgoに設定することで
上書きすることができます（パッケージruntimeを参照）。

	export GODEBUG=netdns=go    # 純粋なGoリゾルバを強制する
	export GODEBUG=netdns=cgo   # ネイティブリゾルバを強制する（cgo、win32）

<<<<<<< HEAD
この判断は、Goソースツリーをビルドする際にも、netgoまたはnetcgoビルドタグを設定することで強制することができます。
=======
The decision can also be forced while building the Go source tree
by setting the netgo or netcgo build tag.
The netgo build tag disables entirely the use of the native (CGO) resolver,
meaning the Go resolver is the only one that can be used.
With the netcgo build tag the native and the pure Go resolver are compiled into the binary,
but the native (CGO) resolver is preferred over the Go resolver.
With netcgo, the Go resolver can still be forced at runtime with GODEBUG=netdns=go.
>>>>>>> upstream/release-branch.go1.25

GODEBUG=netdns=1のような数値のnetdns設定は、リゾルバが自身の判断に関するデバッグ情報を出力します。
特定のリゾルバを強制すると同時にデバッグ情報を出力するには、
2つの設定をプラス記号で結合します。例：GODEBUG=netdns=go+1。

Goリゾルバは、DNSリクエストとともにEDNS0追加ヘッダーを送信し、
より大きなDNSパケットサイズを受け入れる意思を示します。
これは、一部のモデムやルーターが運用するDNSサーバーで断続的な障害を引き起こすと報告されています。
GODEBUG=netedns0=0を設定すると、追加ヘッダーの送信が無効になります。

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
// 2つのメソッド [Addr.Network] および [Addr.String] は、通常、[Dial] の引数として渡すことができる文字列を返しますが、
// その文字列の形式や意味については実装次第です。
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
	UnwrapErr   error
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

// Unwrap returns e.UnwrapErr.
func (e *DNSError) Unwrap() error

func (e *DNSError) Error() string

// Timeoutは、DNSルックアップがタイムアウトしたことが確認されたかどうかを報告します。
// これは常に確認できるわけではありません。DNSルックアップはタイムアウトにより失敗し、
// Timeoutがfalseを返す [DNSError] を返すことがあります。
func (e *DNSError) Timeout() bool

// Temporaryは、DNSエラーが一時的であることが確認されたかどうかを報告します。
// これは常に確認できるわけではありません。DNSルックアップは一時的なエラーにより失敗し、
// Temporaryがfalseを返す [DNSError] を返すことがあります。
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
// WriteTo は、[Buffers] に [io.WriterTo] を実装します。
//
// WriteTo は、0 <= i < len(v) の範囲の v[i] およびスライス v を変更しますが、v[i][j] (i, j は任意の値) は変更しません。
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)

// バッファから読み込む。
//
// Read は [Buffers] に [io.Reader] を実装します。
//
// Read はスライス v と v[i]（ただし、0 <= i < len(v)）を変更しますが、
// v[i][j]（ただし、任意の i, j）は変更しません。
func (v *Buffers) Read(p []byte) (n int, err error)
