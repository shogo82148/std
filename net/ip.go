// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IPアドレスの操作
//
// IPv4アドレスは4バイトで、IPv6アドレスは16バイトです。
// IPv4アドレスは正規のプレフィックス（10個のゼロ、2つの0xFF）を追加することで、IPv6アドレスに変換できます。
// このライブラリはどちらのバイトスライスのサイズも受け付けますが、常に16バイトのアドレスを返します。

package net

// IPアドレスの長さ（バイト単位）。
const (
	IPv4len = 4
	IPv6len = 16
)

// IPは単一のIPアドレス、バイトのスライスです。
// このパッケージの関数は、4バイト（IPv4）または16バイト（IPv6）のスライスを入力として受け付けます。
//
// このドキュメントでは、IPアドレスをIPv4アドレスまたはIPv6アドレスと呼ぶことは、アドレスの意味的な属性であり、単にバイトスライスの長さだけではありません：16バイトスライスもIPv4アドレスである可能性があります。
type IP []byte

// IPMaskは、IPアドレスのアドレッシングとルーティングに使用できる
// ビットマスクです。
//
// 詳細については、型 [IPNet] と関数 [ParseCIDR] を参照してください。
type IPMask []byte

// IPNetはIPネットワークを表します。
type IPNet struct {
	IP   IP
	Mask IPMask
}

// IPv4は、IPv4アドレスa.b.c.dのIPアドレス（16バイト形式）を返します。
func IPv4(a, b, c, d byte) IP

// IPv4Maskは、IPv4マスクa.b.c.dのIPマスク（4バイト形式）を返します。
func IPv4Mask(a, b, c, d byte) IPMask

// CIDRMaskは、'ones'個の1ビットで構成された [IPMask] を返します。
// その後、0ビットが 'bits'ビットの総長になるまで続きます。
// この形式のマスクに対して、CIDRMaskは [IPMask.Size] の逆です。
func CIDRMask(ones, bits int) IPMask

// 有名なIPv4アドレス
var (
	IPv4bcast     = IPv4(255, 255, 255, 255)
	IPv4allsys    = IPv4(224, 0, 0, 1)
	IPv4allrouter = IPv4(224, 0, 0, 2)
	IPv4zero      = IPv4(0, 0, 0, 0)
)

// 有名なIPv6アドレス
var (
	IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)

// IsUnspecified は、ipが特定されていないアドレスであるかどうかを報告します。
// つまり、IPv4アドレスの "0.0.0.0" またはIPv6アドレス "::" です。
func (ip IP) IsUnspecified() bool

// IsLoopbackはipがループバックアドレスであるかどうかを報告します。
func (ip IP) IsLoopback() bool

// IsPrivateは、RFC 1918（IPv4アドレス）およびRFC 4193（IPv6アドレス）に基づいて、IPがプライベートアドレスかどうかを報告します。
func (ip IP) IsPrivate() bool

// IsMulticastは、ipがマルチキャストアドレスかどうかを報告します。
func (ip IP) IsMulticast() bool

// IsInterfaceLocalMulticastは、ipがインターフェースローカルなマルチキャストアドレスかどうかを報告します。
func (ip IP) IsInterfaceLocalMulticast() bool

// IsLinkLocalMulticast は、与えられた IP アドレスがリンクローカルマルチキャストアドレスかどうかを報告します。
func (ip IP) IsLinkLocalMulticast() bool

// IsLinkLocalUnicast は、ip がリンクローカルユニキャストアドレスであるかどうかを報告します。
func (ip IP) IsLinkLocalUnicast() bool

// IsGlobalUnicastは、ipがグローバルユニキャストアドレスであるかどうかを報告します。
//
// グローバルユニキャストアドレスの識別は、RFC 1122、RFC 4632、RFC 4291で定義されたアドレスタイプの識別を使用しますが、
// IPv4の指示ブロードキャストアドレスは除外します。
// ipがIPv4のプライベートアドレススペースまたはローカルIPv6ユニキャストアドレススペースにある場合でも、trueを返します。
func (ip IP) IsGlobalUnicast() bool

// To4はIPv4アドレスを4バイトの表現に変換します。
// もしipがIPv4アドレスでない場合、To4はnilを返します。
func (ip IP) To4() IP

// To16はIPアドレスipを16バイトの表現に変換します。
// ipがIPアドレスでない場合（長さが正しくない場合）、To16はnilを返します。
func (ip IP) To16() IP

// DefaultMaskはIPアドレスipのデフォルトのマスクを返します。
// IPv4アドレスのみがデフォルトマスクを持ちます。ipが有効なIPv4アドレスでない場合、DefaultMaskは
// nilを返します。
func (ip IP) DefaultMask() IPMask

// MaskはIPアドレスipをmaskでマスクした結果を返します。
func (ip IP) Mask(mask IPMask) IP

// StringはIPアドレスipの文字列形式を返します。
// それは以下の4つの形式のいずれかを返します:
//   - "<nil>", ipの長さが0の場合
//   - ドット付きの10進表現 ("192.0.2.1"), ipがIPv4またはIP4-mapped IPv6アドレスの場合
//   - RFC 5952に準拠したIPv6形式 ("2001:db8::1"), ipが有効なIPv6アドレスの場合
//   - 上記の条件に当てはまらない場合は、ipの句読点を除いた16進数形式
func (ip IP) String() string

// MarshalTextは [encoding.TextMarshaler] インターフェースを実装します。
// エンコードは [IP.String] で返されるものと同じですが、1つ例外があります：
// len(ip)がゼロの場合、空のスライスを返します。
func (ip IP) MarshalText() ([]byte, error)

// UnmarshalTextは [encoding.TextUnmarshaler] インターフェースを実装します。
// IPアドレスは [ParseIP] で受け入れられる形式で指定することが期待されています。
func (ip *IP) UnmarshalText(text []byte) error

// Equalは、ipとxが同じIPアドレスであるかどうかを報告します。
// IPv4アドレスと同じアドレスを持つIPv6形式は
// 同じものと見なされます。
func (ip IP) Equal(x IP) bool

// Sizeはマスクの先頭の1の数と合計のビット数を返します。
// マスクが正規の形式でない場合、つまり、1が0に続く形式でない場合は、
// Sizeは0, 0を返します。
func (m IPMask) Size() (ones, bits int)

// Stringは、句読点なしのmの16進数形式を返します。
func (m IPMask) String() string

// Containsは、ネットワークが指定したIPを含んでいるかどうかを報告します。
func (n *IPNet) Contains(ip IP) bool

// Networkはアドレスのネットワーク名、"ip+net"を返します。
func (n *IPNet) Network() string

// Stringは、CIDR表記であるnの文字列を返します。例えば「192.0.2.0/24」やRFC 4632およびRFC 4291で定義されている「2001:db8::/48」です。
// もしマスクが正規形式でない場合、IPアドレスに続いてスラッシュ文字とパンクチュエーションを含まない16進形式のマスクで表された文字列を返します。例えば「198.51.100.0/c000ff00」。
func (n *IPNet) String() string

// ParseIPは文字列sをIPアドレスと解釈し、結果を返します。
// 文字列sはIPv4点区切りの10進法（「192.0.2.1」）、IPv6（「2001:db8::68」）、またはIPv4-mapped IPv6形式（「::ffff:192.0.2.1」）で書かれている必要があります。
// もしsが有効なIPアドレスのテキスト表現ではない場合、ParseIPはnilを返します。
// 返されるアドレスは常に16バイトであり、
// IPv4アドレスはIPv4マップドIPv6形式で返されます。
func ParseIP(s string) IP

// ParseCIDRはCIDR表記のIPアドレスとプレフィックス長を含むsを解析します。
// 例えば、"192.0.2.0/24"や"2001:db8::/32"のようなものです。
// RFC 4632とRFC 4291で定義されています。
//
// IPアドレスとプレフィックス長によって暗示されるIPとネットワークを返します。
// 例えば、ParseCIDR("192.0.2.1/24")はIPアドレスが
// 192.0.2.1でネットワークが192.0.2.0/24を返します。
func ParseCIDR(s string) (IP, *IPNet, error)
