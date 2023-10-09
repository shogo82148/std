// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// UnixAddrはUnixドメインソケットエンドポイントのアドレスを表します。
type UnixAddr struct {
	Name string
	Net  string
}

// Networkはアドレスのネットワーク名を返します。"unix"、"unixgram"、または"unixpacket"です。
func (a *UnixAddr) Network() string

func (a *UnixAddr) String() string

// ResolveUnixAddrは、Unixドメインソケットエンドポイントのアドレスを返します。
//
// ネットワークはUnixのネットワーク名である必要があります。
//
// ネットワークとアドレスのパラメータについての説明は、
// func Dialを参照してください。
func ResolveUnixAddr(network, address string) (*UnixAddr, error)

// UnixConnは、Unixドメインソケットへの接続のためのConnインターフェースの実装です。
type UnixConn struct {
	conn
}

// SyscallConnは生のネットワーク接続を返します。
// これはsyscall.Connインターフェースを実装しています。
func (c *UnixConn) SyscallConn() (syscall.RawConn, error)

// CloseReadは、Unixドメイン接続の読み込み側をシャットダウンします。
// ほとんどの呼び出し元は、単にCloseを使用すべきです。
func (c *UnixConn) CloseRead() error

// CloseWriteはUnixドメイン接続の書き込み側をシャットダウンします。
// ほとんどの呼び出し元は、単にCloseを使用するだけで十分です。
func (c *UnixConn) CloseWrite() error

// ReadFromUnixは、ReadFromと同様に動作しますが、UnixAddrを返します。
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)

// ReadFromはPacketConnのReadFromメソッドを実装します。
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)

// ReadMsgUnix はcからメッセージを読み取り、そのペイロードをbに、
// 関連する帯域外データをoobにコピーします。bにコピーされたバイト数、oobに
// コピーされたバイト数、メッセージに設定されたフラグ、およびメッセージの
// 送信元アドレスを返します。
//
// なお、len(b) == 0 かつ len(oob) > 0 の場合、この関数は依然として接続から
// 1バイトを読み取り(および破棄)ます。
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)

// WriteToUnixはWriteToと同様に動作しますが、UnixAddrを取ります。
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)

// WriteToはPacketConnのWriteToメソッドを実装します。
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgUnixは、ペイロードのbと関連するオーバンドデータのoobから、cを介してaddrにメッセージを書き込みます。書き込まれたペイロードとオーバンドバイトの数を返します。
// 注意：len(b) == 0かつlen(oob) > 0の場合、この関数は依然として接続に1バイトを書き込みます。
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)

// DialUnixは、UnixネットワークにおけるDialと同様の動作をします。
//
// ネットワークはUnixネットワーク名でなければなりません。詳細についてはfunc Dialを参照してください。
//
// laddrがnilでない場合、それは接続のローカルアドレスとして使用されます。
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)

// UnixListenerはUnixドメインソケットのリスナーです。クライアントは通常、Unixドメインソケットを想定せずに、Listenerの型の変数を使用するべきです。
type UnixListener struct {
	fd         *netFD
	path       string
	unlink     bool
	unlinkOnce sync.Once
}

// SyscallConnは、生のネットワーク接続を返します。
// これはsyscall.Connインターフェースを実装します。
//
// 返されるRawConnは、Controlの呼び出しのみをサポートします。
// ReadとWriteはエラーを返します。
func (l *UnixListener) SyscallConn() (syscall.RawConn, error)

// AcceptUnixは次の着信呼び出しを受け入れ、新しい接続を返します。
func (l *UnixListener) AcceptUnix() (*UnixConn, error)

// AcceptはListenerインターフェースのAcceptメソッドを実装します。
// 返される接続は*UnixConn型です。
func (l *UnixListener) Accept() (Conn, error)

// CloseはUnixアドレス上でのリスニングを停止します。既に受け付けた接続は閉じません。
func (l *UnixListener) Close() error

// Addrはリスナーのネットワークアドレスを返します。
// 返されるAddrは、Addrのすべての呼び出しで共有されるため、
// 変更しないでください。
func (l *UnixListener) Addr() Addr

// SetDeadlineはリスナーと関連付けられた締め切りを設定します。
// ゼロの時間値は締め切りを無効にします。
func (l *UnixListener) SetDeadline(t time.Time) error

// File は基になる os.File のコピーを返します。
// 終了時には、f を閉じるのは呼び出し元の責任です。
// l を閉じても f に影響を与えず、f を閉じても l に影響を与えません。
//
// 返された os.File のファイルディスクリプタは、接続のものとは異なります。
// この複製を使用して元のもののプロパティを変更しようとしても、望ましい効果があるかどうかはわかりません。
func (l *UnixListener) File() (f *os.File, err error)

// ListenUnixはUnixネットワーク向けのListenのように機能します。
//
// ネットワークは「unix」または「unixpacket」である必要があります。
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)

// ListenUnixgramはUnixネットワーク用のListenPacketのように動作します。
//
// ネットワークは"unixgram"である必要があります。
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
