// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージasn1は、ITU-T Rec X.690で定義されたDERエンコードされたASN.1データ構造の解析を実装します。
// また、「ASN.1、BER、およびDERのサブセットの素人向けガイド」も参照してください。
// http://luca.ntop.org/Teaching/Appunti/asn1.html。
package asn1

// StructuralErrorは、ASN.1データが有効であることを示していますが、それを受け取るGoの型が一致していません。
type StructuralError struct {
	Msg string
}

func (e StructuralError) Error() string

// SyntaxErrorは、ASN.1データが無効であることを示唆しています。
type SyntaxError struct {
	Msg string
}

func (e SyntaxError) Error() string

// BitStringは、ASN.1 BIT STRINGタイプを使用したい場合に使用する構造体です。ビット文字列は、メモリ上で最も近いバイトまでパディングされ、有効なビット数が記録されます。パディングビットはゼロになります。
type BitString struct {
	Bytes     []byte
	BitLength int
}

// Atは、指定されたインデックスのビットを返します。インデックスが範囲外の場合は0を返します。
func (b BitString) At(i int) int

// RightAlign はパディングビットが先頭にあるスライスを返します。スライスは BitString とメモリを共有する場合があります。
func (b BitString) RightAlign() []byte

// NullRawValue is a [RawValue] with its Tag set to the ASN.1 NULL type tag (5).
var NullRawValue = RawValue{Tag: TagNull}

// NullBytesには、DERエンコードされたASN.1 NULLタイプを表すバイトが含まれています。
var NullBytes = []byte{TagNull, 0}

// ObjectIdentifierは、ASN.1オブジェクト識別子を表します。
type ObjectIdentifier []int

// Equalはoiとotherが同じ識別子を表しているかどうかを報告します。
func (oi ObjectIdentifier) Equal(other ObjectIdentifier) bool

func (oi ObjectIdentifier) String() string

// Enumerated（列挙型）はプレーンなintで表されます。
type Enumerated int

// フラグは任意のデータを受け入れ、存在する場合にはtrueに設定されます。
type Flag bool

// RawValueは、復号化されていないASN.1オブジェクトを表します。
type RawValue struct {
	Class, Tag int
	IsCompound bool
	Bytes      []byte
	FullBytes  []byte
}

// RawContentは、未デコードのDERデータが構造体にとって保存される必要があることを示すために使用されます。使用するには、構造体の最初のフィールドはこの型でなければなりません。他のフィールドがこの型であることはエラーです。
type RawContent []byte

// UnmarshalはDER形式のASN.1データ構造bを解析し、reflectパッケージを使用してvalで指定された任意の値を埋める。
// Unmarshalはreflectパッケージを使用するため、書き込まれる構造体は大文字のフィールド名を使用する必要がある。
// valがnilまたはポインタでない場合、Unmarshalはエラーを返す。
//
// bを解析した後、valに埋めるために使用されなかったバイトはrestとして返される。
// 構造体へのSEQUENCEの解析時、valにマッチするフィールドを持たないトレーリング要素は、
// トレーリングデータではなくSEQUENCEの有効な要素と見なされないため、restには含まれません。
//
//   - An ASN.1 INTEGER can be written to an int, int32, int64,
//     or *[big.Int].
//     If the encoded value does not fit in the Go type,
//     Unmarshal returns a parse error.
//
//   - An ASN.1 BIT STRING can be written to a [BitString].
//
//   - An ASN.1 OCTET STRING can be written to a []byte.
//
//   - An ASN.1 OBJECT IDENTIFIER can be written to an [ObjectIdentifier].
//
//   - An ASN.1 ENUMERATED can be written to an [Enumerated].
//
//   - An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a [time.Time].
//
//   - An ASN.1 PrintableString, IA5String, or NumericString can be written to a string.
//
//   - Any of the above ASN.1 values can be written to an interface{}.
//     The value stored in the interface has the corresponding Go type.
//     For integers, that type is int64.
//
//   - An ASN.1 SEQUENCE OF x or SET OF x can be written
//     to a slice if an x can be written to the slice's element type.
//
//   - An ASN.1 SEQUENCE or SET can be written to a struct
//     if each of the elements in the sequence can be
//     written to the corresponding element in the struct.
//
// 構造体フィールドに対する以下のタグにはUnmarshalに特別な意味があります。
//
//	applicationはAPPLICATIONタグが使用されていることを指定します
//	privateはPRIVATEタグが使用されていることを指定します
//	default:xはオプションの整数フィールドのデフォルト値を設定します（オプションも指定されている場合のみ使用）
//	explicitは暗黙のタグを追加の明示的なタグでラップすることを指定します
//	optionalはフィールドをASN.1 OPTIONALとしてマークします
//	setはSEQUENCEではなくSET型を期待します
//	tag:xはASN.1タグ番号を指定します。これはASN.1 CONTEXT SPECIFICであるということを意味します。
//
// IMPLICITタグを持つASN.1値を文字列フィールドにデコードする場合、
// UnmarshalはデフォルトでPrintableStringになります。これは'@'や'&'などの文字をサポートしません。
// 他のエンコーディングを強制するには、次のタグを使用します:
//
//	ia5は文字列をASN.1 IA5String値として復元します
//	numericは文字列をASN.1 NumericString値として復元します
//	utf8は文字列をASN.1 UTF8String値として復元します
//
// 構造体の最初のフィールドの型がRawContentの場合、構造体の生のASN1コンテンツがそれに保存されます。
//
// スライスの型名が"SET"で終わる場合、これは"set"タグが設定されたように扱われます。これにより、
// タイプがSEQUENCEではなくSET OF xと解釈されます。これは、
// 構造体タグが付けられないネストしたスライスで使用することができます。
//
// 他のASN.1の型はサポートされていません; 遭遇すると、
// Unmarshalは解析エラーを返します。
func Unmarshal(b []byte, val any) (rest []byte, err error)

// UnmarshalWithParamsでは、トップレベルの要素にフィールドパラメータを指定することができます。パラメータの形式は、フィールドタグと同じです。
func UnmarshalWithParams(b []byte, val any, params string) (rest []byte, err error)
