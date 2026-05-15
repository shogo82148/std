// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// uuid パッケージはUUIDの生成と操作のサポートを提供します。
//
// 詳細は [RFC 9562] を参照してください。
//
// 新しいUUIDのランダムコンポーネントは、
// 暗号化して安全な乱数生成器で生成されます。
//
// UUIDはさまざまなアルゴリズムを使用して生成できます。
// [New] 関数は、ほとんどの目的に適したアルゴリズムを使用して
// 新しいUUIDを返します。
//
// [RFC 9562]: https://www.rfc-editor.org/rfc/rfc9562.html
package uuid

// UUIDはRFC 9562で指定された汎用一意識別子です。
//
// UUIDは比較可能で、== 演算子で比較できます。
type UUID [16]byte

// ParseはsでUUIDを返します。
//
// 次の形式での文字列を受け付けます。
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6
//	{f81d4fae-7dec-11d0-a765-00a0c91e6bf6}
//	urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6
//	f81d4fae7dec11d0a76500a0c91e6bf6
//
// 二進法のアルファベット文字は大小文字で使用できます。
func Parse(s string) (UUID, error)

// MustParseはsで表されたUUIDを返します。
//
// sが [Parse] で定義されているUUIDの有効な文字列表記でない場合、パニックを引き起こします。
func MustParse(s string) UUID

// Newは新しいUUIDを返します。
//
// 特定のUUID生成アルゴリズムを必要としないプログラムはNewを使用してください。
// 現在のNewは [NewV4] と等侧です。
func New() UUID

// NilはNil UUID 00000000-0000-0000-0000-000000000000を返します。
//
// Nil UUIDは [RFC 9562のセクション5.9] で定群されています。
// Goのnil値とは異なることに注意してください。
//
// [RFC 9562のセクション5.9]: https://www.rfc-editor.org/rfc/rfc9562#section-5.9
func Nil() UUID

// MaxはMax UUID ffffffff-ffff-ffff-ffff-ffffffffffffを返します。
//
// Max UUIDは [RFC 9562のセクション5.10] で定群されています。
//
// [RFC 9562のセクション5.10]: https://www.rfc-editor.org/rfc/rfc9562#section-5.10
func Max() UUID

// Stringはuの文字列表記を返します。
//
// RFC 9562で定義されている小文字の16進法とダッシュ表記を使用します。
func (u UUID) String() string

// MarshalTextは [encoding.TextMarshaler] インターフェースを実装します。
// エンコードは [UUID.String] で返されるものと同じです
func (u UUID) MarshalText() ([]byte, error)

// AppendTextは [encoding.TextAppender] インターフェースを実装します。
// エンコードは [UUID.String] で返されるものと同じです
func (u UUID) AppendText(b []byte) ([]byte, error)

// UnmarshalTextは [encoding.TextUnmarshaler] インターフェースを実装します。
// UUIDは [Parse] が受け付ける形式であることを予想します。
func (u *UUID) UnmarshalText(b []byte) error

// Compareはuとvを比較します。
// uがvより前にある場合、-1を返します。
// uがvより後ろにある場合、+1を返します。
// 同じ場合は0を返します。
//
// Compareはソート用に [RFC 9562のセクション6.11] で定義されている
// ビッグエンディアンバイト位置を使用します。
//
// [RFC 9562のセクション6.11]: https://www.rfc-editor.org/rfc/rfc9562#section-6.11
func (u UUID) Compare(v UUID) int

// NewV4は新しいバージョン4 UUIDを返します。
//
// バージョン4 UUIDは122ビットのランダムデータを含みます。
func NewV4() UUID

// NewV7は新しいバージョン7 UUIDを返します。
//
// バージョン7 UUIDは最も重要な48ビットにタイムスタンプを含み、
// 最低でも62ビット以上のランダムデータを含みます。
//
// NewV7はシステムクロックが逆方向に動いた時を除いて、
// 常に昇順にソートされたUUIDを返します。
func NewV7() UUID
