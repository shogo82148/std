// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/syscall"
)

// IPAddrはIPエンドポイントのアドレスを表します。
type IPAddr struct {
	IP   IP
	Zone string
}

// Networkはアドレスのネットワーク名を返します。"ip"。
func (a *IPAddr) Network() string

func (a *IPAddr) String() string

// ResolveIPAddrはIPエンドポイントのアドレスを返します。
//
// ネットワークはIPネットワーク名である必要があります。
//
// アドレスパラメーターのホストがリテラルIPアドレスではない場合、
// ResolveIPAddrはIPエンドポイントのアドレスに解決します。
// そうでなければ、アドレスをリテラルのIPアドレスとして解析します。
// アドレスパラメーターはホスト名を使用することもできますが、
// これは推奨されません。なぜなら、ホスト名のIPアドレスのうち最大で1つしか返さないからです。
//
<<<<<<< HEAD
// ネットワークとアドレスパラメーターの説明については、func Dialを参照してください。
func ResolveIPAddr(network, address string) (*IPAddr, error)

// IPConnはIPネットワーク接続のConnおよびPacketConnインターフェースの実装です。
=======
// See func [Dial] for a description of the network and address
// parameters.
func ResolveIPAddr(network, address string) (*IPAddr, error)

// IPConn is the implementation of the [Conn] and [PacketConn] interfaces
// for IP network connections.
>>>>>>> upstream/master
type IPConn struct {
	conn
}

<<<<<<< HEAD
// SyscallConnは、生のネットワーク接続を返します。
// これはsyscall.Connインターフェースを実装しています。
=======
// SyscallConn returns a raw network connection.
// This implements the [syscall.Conn] interface.
>>>>>>> upstream/master
func (c *IPConn) SyscallConn() (syscall.RawConn, error)

// ReadFromIPはReadFromと同様に動作しますが、IPAddrを返します。
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)

<<<<<<< HEAD
// ReadFromはPacketConnのReadFromメソッドを実装します。
=======
// ReadFrom implements the [PacketConn] ReadFrom method.
>>>>>>> upstream/master
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)

// ReadMsgIPはcからメッセージを読み取り、ペイロードをbにコピーし、
// 関連する帯域外データをoobにコピーします。bにコピーされたバイト数、oobにコピーされたバイト数、
// メッセージに設定されたフラグ、およびメッセージの送信元アドレスを返します。
//
// パッケージgolang.org/x/net/ipv4とgolang.org/x/net/ipv6を使用して、oobに対してIPレベルのソケットオプションを操作できます。
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)

<<<<<<< HEAD
// WriteToIPはWriteToと同様の動作をするが、IPAddrを取ります。
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)

// WriteToはPacketConnのWriteToメソッドを実装します。
=======
// WriteToIP acts like [IPConn.WriteTo] but takes an [IPAddr].
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)

// WriteTo implements the [PacketConn] WriteTo method.
>>>>>>> upstream/master
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgIPは、bからペイロードを、oobから関連のオフドーバンドータをコピーし、cを経由してaddrにメッセージを送信します。送信されたペイロードとオフドーバンドズダのバイト数を返します。
//
// golang.org/x/net/ipv4とgolang.org/x/net/ipv6のパッケージを使用して、oob内のIPレベルのソケットオプションを操作することができます。
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)

<<<<<<< HEAD
// DialIPはIPネットワークに対してDialのように機能します。
=======
// DialIP acts like [Dial] for IP networks.
>>>>>>> upstream/master
//
// ネットワークはIPネットワーク名である必要があります。詳細はfunc Dialを参照してください。
//
// もしladdrがnilであれば、ローカルアドレスが自動的に選択されます。
// もしraddrのIPフィールドがnilであるか、未指定のIPアドレスである場合、
// ローカルシステムが仮定されます。
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)

<<<<<<< HEAD
// ListenIPはIPネットワーク用のListenPacketと同様に機能します。
=======
// ListenIP acts like [ListenPacket] for IP networks.
>>>>>>> upstream/master
//
// ネットワークはIPネットワーク名である必要があります。詳細についてはfunc Dialを参照してください。
//
// もしladdrのIPフィールドがnilまたは指定されていないIPアドレスである場合、
// ListenIPはローカルシステムの利用可能なすべてのIPアドレスでリッスンします
// マルチキャストIPアドレスを除く。
func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
