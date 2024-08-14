// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/net/netip"
	"github.com/shogo82148/std/syscall"
)

// UDPAddrはUDPエンドポイントのアドレスを表します。
type UDPAddr struct {
	IP   IP
	Port int
	Zone string
}

// AddrPortは [UDPAddr] aを [netip.AddrPort] として返します。
//
// もしa.Portがuint16に収まらない場合、静かに切り捨てられます。
//
// もしaがnilの場合、ゼロ値が返されます。
func (a *UDPAddr) AddrPort() netip.AddrPort

// Networkはアドレスのネットワーク名、"udp"を返します。
func (a *UDPAddr) Network() string

func (a *UDPAddr) String() string

// ResolveUDPAddr はUDPのエンドポイントのアドレスを返します。
//
// ネットワークはUDPのネットワーク名である必要があります。
//
// アドレスパラメータのホストがIPアドレスのリテラルでない場合、または
// ポート番号がリテラルのポート番号でない場合、ResolveUDPAddrは
// UDPエンドポイントのアドレスに解決します。
// それ以外の場合は、アドレスをリテラルのIPアドレスとポート番号のペアとして解析します。
// アドレスパラメータはホスト名を使用することもできますが、これは
// 推奨されません。なぜなら、ホスト名のIPアドレスのいずれか一つしか返さないからです。
//
// ネットワークおよびアドレスパラメータの説明については、[Dial] 関数を参照してください。
func ResolveUDPAddr(network, address string) (*UDPAddr, error)

// UDPAddrFromAddrPortはaddrを [UDPAddr] として返します。
// もしaddr.IsValid()がfalseであれば、返されるUDPAddrにはnilのIPフィールドが含まれ、
// アドレスファミリーに依存しない未指定のアドレスを示します。
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr

// UDPConnはUDPネットワーク接続の [Conn] および [PacketConn] インターフェースの実装です。
type UDPConn struct {
	conn
}

// SyscallConnは生のネットワーク接続を返します。
// これは [syscall.Conn] インターフェースを実装しています。
func (c *UDPConn) SyscallConn() (syscall.RawConn, error)

// ReadFromUDPは [UDPConn.ReadFrom] と同様の動作をしますが、UDPAddrを返します。
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)

// ReadFrom は [PacketConn] の ReadFrom メソッドを実装します。
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)

// ReadFromUDPAddrPortはReadFromと同様の機能を提供しますが、[netip.AddrPort] を返します。
//
// cが指定されていないアドレスにバインドされている場合、返される
// netip.AddrPortのアドレスは、IPv4-mapped IPv6アドレスの可能性があります。
// IPv6のプレフィックスなしのアドレスを取得するには、[netip.Addr.Unmap] を使用してください。
func (c *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error)

// ReadMsgUDPは、cからメッセージを読み込み、ペイロードをbにコピーし、サイドバンドデータをoobにコピーします。bにコピーされたバイト数、oobにコピーされたバイト数、メッセージに設定されたフラグ、およびメッセージのソースアドレスを返します。
//
// パッケージ [golang.org/x/net/ipv4] および [golang.org/x/net/ipv6] は、oob内のIPレベルのソケットオプションを操作するために使用できます。
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)

// ReadMsgUDPAddrPortは [UDPConn.WriteTo] と同様に動作しますが、[UDPAddr] の代わりに [netip.AddrPort] を返します。
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error)

// WriteToUDPはWriteToと同様に動作しますが、[UDPAddr] を引数に取ります。
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)

// WriteToUDPAddrPortは、WriteToと同様に動作しますが、[netip.AddrPort] を受け取ります。
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error)

// WriteToは [PacketConn] のWriteToメソッドを実装します。
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgUDPは、cが接続されていない場合はcを介してaddrにメッセージを書き込み、
// cが接続されている場合はcのリモートアドレスにメッセージを書き込みます（その場合、
// addrはnilでなければなりません）。ペイロードはbからコピーされ、関連する
// フラグデータはoobからコピーされます。ペイロードとフラグデータの書き込まれた
// バイト数を返します。
//
// パッケージ [golang.org/x/net/ipv4] および [golang.org/x/net/ipv6] を使用して、
// oob内のIPレベルのソケットオプションを操作することができます。
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)

// WriteMsgUDPAddrPortは [UDPConn.WriteMsgUDP] と同様に動作しますが、[UDPAddr] の代わりに [netip.AddrPort] を受け取ります。
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error)

// DialUDPはUDPネットワークのために [Dial] と同様の機能を提供します。
//
// ネットワークはUDPネットワークの名前でなければならず、詳細については [Dial] 関数を参照してください。
//
// もしladdrがnilの場合、自動的にローカルアドレスが選択されます。
// もしraddrのIPフィールドがnilまたは未指定のIPアドレスの場合、ローカルシステムが仮定されます。
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)

// ListenUDPは、UDPネットワークに対して [ListenPacket] と同様の機能を提供します。
//
// ネットワークはUDPネットワーク名でなければなりません。詳細については、[Dial] 関数を参照してください。
//
// laddrのIPフィールドがnilまたは未指定のIPアドレスである場合、
// ListenUDPは、マルチキャストIPアドレスを除く、ローカルシステムのすべての利用可能なIPアドレスでリスンします。
// laddrのPortフィールドが0の場合、ポート番号が自動的に選択されます。
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)

// ListenMulticastUDPは、UDPネットワークに対して [ListenPacket] と同様に動作しますが、
// 特定のネットワークインターフェース上のグループアドレスを受け取ります。
//
// ネットワークはUDPネットワーク名でなければなりません。詳細については、[Dial] 関数を参照してください。
//
// ListenMulticastUDPは、グループのマルチキャストIPアドレスを含む、
// ローカルシステムのすべての利用可能なIPアドレスでリッスンします。
// ifiがnilの場合、ListenMulticastUDPはシステムが割り当てた
// マルチキャストインターフェースを使用しますが、これは推奨されません。
// 割り当てはプラットフォームに依存し、場合によってはルーティング構成が必要になることがあるためです。
// gaddrのPortフィールドが0の場合、ポート番号は自動的に選択されます。
//
// ListenMulticastUDPは、シンプルで小さなアプリケーションのための便利な関数です。
// 一般的な用途には、[golang.org/x/net/ipv4] および [golang.org/x/net/ipv6] パッケージがあります。
//
// ListenMulticastUDPは、IPPROTO_IPの下でIP_MULTICAST_LOOPソケットオプションを0に設定し、
// マルチキャストパケットのループバックを無効にすることに注意してください。
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
