// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージnetipは、小さな値型であるIPアドレス型を定義します。
// この[Addr]型をベースに、パッケージは [AddrPort] （IPアドレスとポート）と [Prefix] （IPアドレスとビット長のプレフィックス）も定義します。
//
// [net.IP] 型と比較して、 [Addr] 型はメモリを少なく使用し、不変であり、比較可能（==およびマップキーとしてサポート）です。
package netip

import (
	"github.com/shogo82148/std/internal/intern"
)

// Addrは、[net.IP]または[net.IPAddr]に似た、スコープ付きアドレスを持つIPv4またはIPv6アドレスを表します。
//
// [net.IP]または[net.IPAddr]とは異なり、Addrは比較可能な値型であり（==をサポートし、マップキーとして使用可能）、不変です。
//
// ゼロ値のAddrは有効なIPアドレスではありません。
// Addr{}は、0.0.0.0と::の両方とは異なります。
type Addr struct {
	// addr is the hi and lo bits of an IPv6 address. If z==z4,
	// hi and lo contain the IPv4-mapped IPv6 address.
	//
	// hi and lo are constructed by interpreting a 16-byte IPv6
	// address as a big-endian 128-bit number. The most significant
	// bits of that number go into hi, the rest into lo.
	//
	// For example, 0011:2233:4455:6677:8899:aabb:ccdd:eeff is stored as:
	//  addr.hi = 0x0011223344556677
	//  addr.lo = 0x8899aabbccddeeff
	//
	// We store IPs like this, rather than as [16]byte, because it
	// turns most operations on IPs into arithmetic and bit-twiddling
	// operations on 64-bit registers, which is much faster than
	// bytewise processing.
	addr uint128

	// z is a combination of the address family and the IPv6 zone.
	//
	// nil means invalid IP address (for a zero Addr).
	// z4 means an IPv4 address.
	// z6noz means an IPv6 address without a zone.
	//
	// Otherwise it's the interned zone name string.
	z *intern.Value
}

// IPv6LinkLocalAllNodesは、IPv6リンクローカル全ノードマルチキャストアドレスff02::1を返します。
func IPv6LinkLocalAllNodes() Addr

// IPv6LinkLocalAllRoutersは、IPv6リンクローカル全ルーターマルチキャストアドレスff02::2を返します。
func IPv6LinkLocalAllRouters() Addr

// IPv6Loopbackは、IPv6ループバックアドレス::1を返します。
func IPv6Loopback() Addr

// IPv6Unspecifiedは、IPv6未指定アドレス"::"を返します。
func IPv6Unspecified() Addr

// IPv4Unspecifiedは、IPv4未指定アドレス"0.0.0.0"を返します。
func IPv4Unspecified() Addr

// AddrFrom4は、addrのバイトで指定されたIPv4アドレスのアドレスを返します。
func AddrFrom4(addr [4]byte) Addr

// AddrFrom16は、addrのバイトで指定されたIPv6アドレスを返します。
// IPv4マップされたIPv6アドレスはIPv6アドレスのままです。
// （必要に応じて、Unmapを使用して変換してください。）
func AddrFrom16(addr [16]byte) Addr

// ParseAddrは、sをIPアドレスとして解析し、その結果を返します。
// 文字列sは、ドット付き10進表記（"192.0.2.1"）、IPv6（"2001:db8::68"）、
// またはスコープ付きアドレスゾーンを持つIPv6（"fe80::1cc0:3e8c:119f:c2e1%ens18"）のいずれかである必要があります。
func ParseAddr(s string) (Addr, error)

// / MustParseAddrは、ParseAddr(s)を呼び出し、エラーが発生した場合にパニックを引き起こします。
// ハードコードされた文字列を使用したテストで使用することを目的としています。
func MustParseAddr(s string) Addr

// AddrFromSliceは、4バイトまたは16バイトのバイトスライスをIPv4またはIPv6アドレスとして解析します。
// net.IPは、[]byte引数として直接渡すことができます。
// スライスの長さが4または16でない場合、Addr{}、falseを返します。
func AddrFromSlice(slice []byte) (ip Addr, ok bool)

// IsValidは、Addrが初期化されたアドレス（ゼロのAddrではない）であるかどうかを報告します。
//
// "0.0.0.0"と"::"の両方が有効な値であることに注意してください。
func (ip Addr) IsValid() bool

// BitLenは、IPアドレスのビット数を返します。
// IPv6の場合は128、IPv4の場合は32、ゼロのAddrの場合は0です。
//
// IPv4マップされたIPv6アドレスはIPv6アドレスと見なされるため、ビット長は128になります。
func (ip Addr) BitLen() int

// Zoneは、ipのIPv6スコープ付きアドレッシングゾーンを返します（存在する場合）。
func (ip Addr) Zone() string

// Compareは、2つのIPを比較して整数を返します。
// ip == ip2の場合、結果は0になります。
// ip < ip2の場合、結果は-1になります。
// ip > ip2の場合、結果は+1になります。
// "less than"の定義は、Lessメソッドと同じです。
func (ip Addr) Compare(ip2 Addr) int

// Lessは、ipがip2よりも前にソートされるかどうかを報告します。
// IPアドレスは、まず長さでソートされ、次にアドレスでソートされます。
// ゾーンを持つIPv6アドレスは、ゾーンのない同じアドレスの直後にソートされます。
func (ip Addr) Less(ip2 Addr) bool

// Is4は、ipがIPv4アドレスであるかどうかを報告します。
//
// IPv4マップされたIPv6アドレスの場合、falseを返します。Addr.Unmapを参照してください。
func (ip Addr) Is4() bool

// Is4In6は、ipがIPv4マップされたIPv6アドレスであるかどうかを報告します。
func (ip Addr) Is4In6() bool

// Is6は、IPv6アドレス、IPv4マップされたIPv6アドレスを含むかどうかを報告します。
func (ip Addr) Is6() bool

// Unmapは、IPv4マップされたIPv6アドレスのプレフィックスを削除したipを返します。
//
// つまり、ipがIPv4アドレスをラップしたIPv6アドレスである場合、
// ラップされたIPv4アドレスを返します。それ以外の場合は、ipを変更せずに返します。
func (ip Addr) Unmap() Addr

// WithZoneは、提供されたゾーンを持つipと同じIPを返します。
// zoneが空の場合、ゾーンは削除されます。
// ipがIPv4アドレスの場合、WithZoneは何も行わず、ipを変更せずに返します。
func (ip Addr) WithZone(zone string) Addr

// IsLinkLocalUnicastは、ipがリンクローカルユニキャストアドレスであるかどうかを報告します。
func (ip Addr) IsLinkLocalUnicast() bool

// IsLoopbackは、ipがループバックアドレスであるかどうかを報告します。
func (ip Addr) IsLoopback() bool

// IsMulticastは、ipがマルチキャストアドレスであるかどうかを報告します。
func (ip Addr) IsMulticast() bool

// IsInterfaceLocalMulticastは、ipがIPv6インターフェースローカルマルチキャストアドレスであるかどうかを報告します。
func (ip Addr) IsInterfaceLocalMulticast() bool

// IsLinkLocalMulticastは、ipがリンクローカルマルチキャストアドレスであるかどうかを報告します。
func (ip Addr) IsLinkLocalMulticast() bool

// IsGlobalUnicastは、ipがグローバルユニキャストアドレスであるかどうかを報告します。
//
// リンクローカルアドレススペースを除く、現在のIANA割り当て2000::/3のグローバルユニキャストスペース外にあるIPv6アドレスに対してtrueを返します。
// また、ipがIPv4プライベートアドレススペースまたはIPv6ユニークローカルアドレススペースにある場合でも、trueを返します。
// ゼロのAddrの場合はfalseを返します。
//
// 参考文献については、RFC 1122、RFC 4291、およびRFC 4632を参照してください。
func (ip Addr) IsGlobalUnicast() bool

// IsPrivateは、RFC 1918（IPv4アドレス）およびRFC 4193（IPv6アドレス）に従って、
// ipがプライベートアドレスであるかどうかを報告します。
// つまり、ipが10.0.0.0/8、172.16.0.0/12、192.168.0.0/16、またはfc00::/7のいずれかであるかどうかを報告します。
// これは、net.IP.IsPrivateと同じです。
func (ip Addr) IsPrivate() bool

// IsUnspecifiedは、ipが未指定のアドレスであるかどうかを報告します。
// IPv4アドレス"0.0.0.0"またはIPv6アドレス"::"のいずれかです。
//
// ただし、ゼロのAddrは未指定のアドレスではありません。
func (ip Addr) IsUnspecified() bool

// Prefixは、IPの上位bビットのみを保持し、指定された長さのPrefixを生成します。
// ipがゼロのAddrの場合、Prefixは常にゼロのPrefixとnilエラーを返します。
// それ以外の場合、bitsが負の場合またはip.BitLen()より大きい場合、Prefixはエラーを返します。
func (ip Addr) Prefix(b int) (Prefix, error)

// As16は、IPアドレスを16バイトの表現で返します。
// IPv4アドレスはIPv4マップされたIPv6アドレスとして返されます。
// ゾーンを持つIPv6アドレスは、ゾーンを除いた形式で返されます（ゾーンを取得するにはZoneメソッドを使用してください）。
// ゼロのAddrの場合は、すべてのバイトがゼロの値を返します。
func (ip Addr) As16() (a16 [16]byte)

// As4は、IPv4またはIPv4-in-IPv6アドレスを4バイトの表現で返します。
// ipがゼロのAddrまたはIPv6アドレスの場合、As4はパニックを引き起こします。
// 0.0.0.0はゼロのAddrではないことに注意してください。
func (ip Addr) As4() (a4 [4]byte)

// AsSliceは、IPv4またはIPv6アドレスを、それぞれ4バイトまたは16バイトの表現で返します。
func (ip Addr) AsSlice() []byte

// Nextは、ipの次のアドレスを返します。
// アドレスが存在しない場合、ゼロのAddrを返します。
func (ip Addr) Next() Addr

// Prevは、ipの前のアドレスを返します。
// アドレスが存在しない場合、ゼロのAddrを返します。
func (ip Addr) Prev() Addr

// Stringは、IPアドレスipの文字列形式を返します。
// 返される形式は、次の5つのいずれかです。
//
//   - ゼロのAddrの場合は "invalid IP"
//   - IPv4ドット付き10進数表記 ("192.0.2.1")
//   - IPv6表記 ("2001:db8::1")
//   - Is4In6の場合は "::ffff:1.2.3.4"
//   - ゾーンを持つIPv6表記 ("fe80:db8::1%eth0")
//
// 注意：パッケージnetのIP.Stringメソッドとは異なり、
// IPv4マップされたIPv6アドレスは、ドット区切りの4つ組の前に"::ffff:"の接頭辞が付きます。
func (ip Addr) String() string

// AppendToは、MarshalTextによって生成されたipのテキストエンコーディングをbに追加し、拡張されたバッファを返します。
func (ip Addr) AppendTo(b []byte) []byte

// StringExpandedは、Stringと同様ですが、IPv6アドレスは先頭にゼロを付けて"::"の圧縮を行わずに展開されます。
// たとえば、"2001:db8::1"は"2001:0db8:0000:0000:0000:0000:0000:0001"になります。
func (ip Addr) StringExpanded() string

// MarshalTextは、encoding.TextMarshalerインターフェースを実装します。
// エンコーディングは、Stringが返すものと同じですが、1つの例外があります。
// ipがゼロのAddrの場合、エンコーディングは空の文字列になります。
func (ip Addr) MarshalText() ([]byte, error)

// UnmarshalTextは、encoding.TextUnmarshalerインターフェースを実装します。
// IPアドレスは、ParseAddrで受け入れられる形式で指定する必要があります。
//
// textが空の場合、UnmarshalTextは*ipをゼロのAddrに設定し、エラーを返しません。
func (ip *Addr) UnmarshalText(text []byte) error

// MarshalBinaryは、encoding.BinaryMarshalerインターフェースを実装します。
// ゼロのAddrの場合は長さ0のスライスを返し、IPv4アドレスの場合は4バイトの形式を、
// IPv6アドレスの場合はゾーンを追加した16バイトの形式を返します。
func (ip Addr) MarshalBinary() ([]byte, error)

// UnmarshalBinaryは、encoding.BinaryUnmarshalerインターフェースを実装します。
// MarshalBinaryによって生成された形式のデータを想定しています。
func (ip *Addr) UnmarshalBinary(b []byte) error

// AddrPortは、IPアドレスとポート番号です。
type AddrPort struct {
	ip   Addr
	port uint16
}

// AddrPortFromは、提供されたIPとポート番号でAddrPortを返します。
// アロケーションは行いません。
func AddrPortFrom(ip Addr, port uint16) AddrPort

// Addrは、pのIPアドレスを返します。
func (p AddrPort) Addr() Addr

// Portは、pのポート番号を返します。
func (p AddrPort) Port() uint16

// ParseAddrPortは、sをAddrPortとして解析します。
//
// 名前解決は行われません。アドレスとポートの両方が数値である必要があります。
func ParseAddrPort(s string) (AddrPort, error)

// MustParseAddrPortは、ParseAddrPort(s)を呼び出し、エラーが発生した場合にパニックを引き起こします。
// テストでハードコードされた文字列を使用するために使用することを意図しています。
func MustParseAddrPort(s string) AddrPort

// IsValidは、p.Addr()が有効かどうかを報告します。
// ゼロを含むすべてのポートが有効です。
func (p AddrPort) IsValid() bool

// Compare returns an integer comparing two AddrPorts.
// The result will be 0 if p == p2, -1 if p < p2, and +1 if p > p2.
// AddrPorts sort first by IP address, then port.
func (p AddrPort) Compare(p2 AddrPort) int

func (p AddrPort) String() string

// AppendToは、MarshalTextによって生成されたpのテキストエンコーディングをbに追加し、拡張されたバッファを返します。
func (p AddrPort) AppendTo(b []byte) []byte

// MarshalTextは、encoding.TextMarshalerインターフェースを実装します。
// エンコーディングは、Stringが返すものと同じですが、1つの例外があります。
// p.Addr()がゼロのAddrの場合、エンコーディングは空の文字列になります。
func (p AddrPort) MarshalText() ([]byte, error)

// UnmarshalTextは、encoding.TextUnmarshalerインターフェースを実装します。
// AddrPortは、MarshalTextによって生成された形式のデータ、またはParseAddrPortで受け入れられる形式で指定する必要があります。
func (p *AddrPort) UnmarshalText(text []byte) error

// MarshalBinaryは、encoding.BinaryMarshalerインターフェースを実装します。
// これは、Addr.MarshalBinaryに、リトルエンディアンで表されたポートを追加したものを返します。
func (p AddrPort) MarshalBinary() ([]byte, error)

// UnmarshalBinaryは、encoding.BinaryUnmarshalerインターフェースを実装します。
// これは、MarshalBinaryによって生成された形式のデータを想定しています。
func (p *AddrPort) UnmarshalBinary(b []byte) error

// Prefixは、IPネットワークを表すIPアドレスプレフィックス（CIDR）です。
//
// Addr()の最初のBits()が指定されます。残りのビットは任意のアドレスに一致します。
// Bits()の範囲は、IPv4の場合は[0,32]、IPv6の場合は[0,128]です。
type Prefix struct {
	ip Addr

	// bitsPlusOne stores the prefix bit length plus one.
	// A Prefix is valid if and only if bitsPlusOne is non-zero.
	bitsPlusOne uint8
}

// PrefixFromは、指定されたIPアドレスとビットプレフィックス長でPrefixを返します。
//
// アロケーションは行いません。Addr.Prefixとは異なり、PrefixFromはipのホストビットをマスクしません。
//
// bitsが負の場合またはip.BitLenより大きい場合、Prefix.Bitsは無効な値-1を返します。
func PrefixFrom(ip Addr, bits int) Prefix

// Addrは、pのIPアドレスを返します。
func (p Prefix) Addr() Addr

// Bitsは、pのプレフィックス長を返します。
//
// 無効な場合は-1を報告します。
func (p Prefix) Bits() int

// IsValidは、p.Addr()に対してp.Bits()が有効な範囲であるかどうかを報告します。
// p.Addr()がゼロのAddrの場合、IsValidはfalseを返します。
// pがゼロのPrefixの場合、p.IsValid() == falseになることに注意してください。
func (p Prefix) IsValid() bool

// IsSingleIPは、pが正確に1つのIPを含むかどうかを報告します。
func (p Prefix) IsSingleIP() bool

// Compareは、2つのプレフィックスを比較して整数を返します。
// p == p2の場合、結果は0になります。p < p2の場合は-1、p > p2の場合は+1になります。
// プレフィックスは、まず有効性（有効でない場合は有効な前）、次にアドレスファミリ（IPv4はIPv6の前）、
// 次にプレフィックス長、最後にアドレスでソートされます。
func (p Prefix) Compare(p2 Prefix) int

// ParsePrefixは、sをIPアドレスプレフィックスとして解析します。
// 文字列は、RFC 4632およびRFC 4291で定義されたCIDR表記形式である場合があります。
// 文字列は、"192.168.1.0/24"または"2001:db8::/32"のようになります。
// IPv6ゾーンはプレフィックスで許可されていません。ゾーンが存在する場合は、エラーが返されます。
//
// マスクされたアドレスビットはゼロになりません。そのため、Maskedを使用してください。
func ParsePrefix(s string) (Prefix, error)

// MustParsePrefixは、ParsePrefix(s)を呼び出し、エラーが発生した場合にパニックを引き起こします。
// テストでハードコードされた文字列を使用するために使用することを意図しています。
func MustParsePrefix(s string) Prefix

// Maskedは、pを正規形式で返します。p.Addr()の高位p.Bits()ビット以外はすべてマスクされます。
//
// pがゼロまたは無効な場合、MaskedはゼロのPrefixを返します。
func (p Prefix) Masked() Prefix

// Containsは、ネットワークpがipを含むかどうかを報告します。
//
// IPv4アドレスはIPv6プレフィックスに一致しません。
// IPv4マップされたIPv6アドレスはIPv4プレフィックスに一致しません。
// ゼロ値のIPはどのプレフィックスにも一致しません。
// ipがIPv6ゾーンを持つ場合、Prefixesはゾーンを削除するため、Containsはfalseを返します。
func (p Prefix) Contains(ip Addr) bool

// Overlapsは、pとoが共通のIPアドレスを含むかどうかを報告します。
//
// pとoが異なるアドレスファミリであるか、どちらかがゼロのIPである場合、falseを報告します。
// Containsメソッドと同様に、IPv4マップされたIPv6アドレスを持つプレフィックスは、IPv6マスクとして扱われます。
func (p Prefix) Overlaps(o Prefix) bool

// AppendToは、MarshalTextによって生成されたpのテキストエンコーディングをbに追加し、拡張されたバッファを返します。
func (p Prefix) AppendTo(b []byte) []byte

// MarshalTextは、encoding.TextMarshalerインターフェースを実装します。
// エンコーディングは、Stringが返すものと同じですが、1つの例外があります。
// pがゼロ値の場合、エンコーディングは空の文字列になります。
func (p Prefix) MarshalText() ([]byte, error)

// UnmarshalTextは、encoding.TextUnmarshalerインターフェースを実装します。
// IPアドレスは、ParsePrefixで受け入れられる形式で指定する必要があります。
// または、MarshalTextによって生成された形式である必要があります。
func (p *Prefix) UnmarshalText(text []byte) error

// MarshalBinaryは、encoding.BinaryMarshalerインターフェースを実装します。
// これは、Addr.MarshalBinaryに、プレフィックスビットを表す追加のバイトを追加したものを返します。
func (p Prefix) MarshalBinary() ([]byte, error)

// UnmarshalBinaryは、encoding.BinaryUnmarshalerインターフェースを実装します。
// これは、MarshalBinaryによって生成された形式のデータを想定しています。
func (p *Prefix) UnmarshalBinary(b []byte) error

// Stringは、pのCIDR表記を返します: "<ip>/<bits>"。
func (p Prefix) String() string
