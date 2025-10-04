// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// jsonパッケージは、RFC 7159で定義されているJSONのエンコーディングとデコーディングを実装します。
// JSONとGoの値の間のマッピングは、Marshal関数とUnmarshal関数のドキュメンテーションで説明されています。
=======
//go:build !goexperiment.jsonv2

// Package json implements encoding and decoding of JSON as defined in RFC 7159.
// The mapping between JSON and Go values is described in the documentation for
// the Marshal and Unmarshal functions.
>>>>>>> upstream/release-branch.go1.25
//
// このパッケージの紹介については、「JSONとGo」を参照してください：
// https://golang.org/doc/articles/json_and_go.html
//
// # Security Considerations
//
// The JSON standard (RFC 7159) is lax in its definition of a number of parser
// behaviors. As such, many JSON parsers behave differently in various
// scenarios. These differences in parsers mean that systems that use multiple
// independent JSON parser implementations may parse the same JSON object in
// differing ways.
//
// Systems that rely on a JSON object being parsed consistently for security
// purposes should be careful to understand the behaviors of this parser, as
// well as how these behaviors may cause interoperability issues with other
// parser implementations.
//
// Due to the Go Backwards Compatibility promise (https://go.dev/doc/go1compat)
// there are a number of behaviors this package exhibits that may cause
// interopability issues, but cannot be changed. In particular the following
// parsing behaviors may cause issues:
//
//   - If a JSON object contains duplicate keys, keys are processed in the order
//     they are observed, meaning later values will replace or be merged into
//     prior values, depending on the field type (in particular maps and structs
//     will have values merged, while other types have values replaced).
//   - When parsing a JSON object into a Go struct, keys are considered in a
//     case-insensitive fashion.
//   - When parsing a JSON object into a Go struct, unknown keys in the JSON
//     object are ignored (unless a [Decoder] is used and
//     [Decoder.DisallowUnknownFields] has been called).
//   - Invalid UTF-8 bytes in JSON strings are replaced by the Unicode
//     replacement character.
//   - Large JSON number integers will lose precision when unmarshaled into
//     floating-point types.
package json

import (
	"github.com/shogo82148/std/reflect"
)

// Marshalは、vのJSONエンコーディングを返します。
//
// Marshalは、値vを再帰的に走査します。
// もし遭遇した値が [Marshaler] を実装しており、
// それがnilポインタでない場合、Marshalは [Marshaler.MarshalJSON] を呼び出して
// JSONを生成します。[Marshaler.MarshalJSON] メソッドが存在しないが、
// その値が代わりに [encoding.TextMarshaler] を実装している場合、Marshalは
// [encoding.TextMarshaler.MarshalText] を呼び出し、その結果をJSON文字列としてエンコードします。
// nilポインタの例外は厳密には必要ではありませんが、
// [Unmarshaler.UnmarshalJSON] の振る舞いにおける同様の、必要な例外を模倣します。
//
// それ以外の場合、Marshalは以下の型依存のデフォルトエンコーディングを使用します：
//
// ブール値はJSONのブール値としてエンコードされます。
//
// 浮動小数点数、整数、および [Number] の値はJSONの数値としてエンコードされます。
// NaNおよび+/-Infの値は [UnsupportedValueError] を返します。
//
// 文字列の値は、無効なバイトをUnicodeの置換文字に置き換えて、
// 有効なUTF-8に強制されたJSON文字列としてエンコードされます。
// JSONがHTMLの<script>タグ内に埋め込んでも安全であるように、
// 文字列は [HTMLEscape] を使用してエンコードされ、
// "<", ">", "&", U+2028, および U+2029 が "\u003c","\u003e", "\u0026", "\u2028", および "\u2029" にエスケープされます。
// この置換は、[Encoder] を使用している場合、[Encoder.SetEscapeHTML](false)を呼び出すことで無効にできます。
//
// 配列とスライスの値はJSON配列としてエンコードされますが、
// []byteはbase64エンコードされた文字列としてエンコードされ、
// nilスライスはnullのJSON値としてエンコードされます。
//
// 構造体の値はJSONオブジェクトとしてエンコードされます。
// エクスポートされた各構造体フィールドは、オブジェクトのメンバーとなり、
// フィールド名がオブジェクトキーとして使用されます。ただし、以下に示す理由のいずれかでフィールドが省略される場合があります。
//
// 各構造体フィールドのエンコーディングは、構造体フィールドのタグの"json"キーの下に格納された
// フォーマット文字列によってカスタマイズできます。
// フォーマット文字列はフィールド名を指定し、それに続いてカンマで区切られたオプションのリストが続く可能性があります。
// デフォルトのフィールド名を上書きせずにオプションを指定するために、名前は空にすることができます。
//
<<<<<<< HEAD
// "omitempty"オプションは、フィールドが空の値を持つ場合、
// エンコーディングからそのフィールドを省略することを指定します。
// 空の値とは、false、0、nilポインタ、nilインターフェース値、
// そして任意の空の配列、スライス、マップ、または文字列を指します。
=======
// The "omitempty" option specifies that the field should be omitted
// from the encoding if the field has an empty value, defined as
// false, 0, a nil pointer, a nil interface value, and any array,
// slice, map, or string of length zero.
>>>>>>> upstream/release-branch.go1.25
//
// 特別なケースとして、フィールドタグが"-"の場合、フィールドは常に省略されます。
// フィールド名が"-"のフィールドでも、タグ"-,"を使用して生成することができることに注意してください。
//
// 構造体フィールドタグの例とその意味：
//
//	// フィールドはJSONでキー"myName"として現れます。
//	Field int `json:"myName"`
//
//	// フィールドはJSONでキー"myName"として現れ、
//	// フィールドの値が空の場合、オブジェクトから省略されます。
//	// 上記で定義されているように。
//	Field int `json:"myName,omitempty"`
//
//	// フィールドはJSONでキー"Field"（デフォルト）として現れますが、
//	// フィールドが空の場合はスキップされます。
//	// 先頭のカンマに注意してください。
//	Field int `json:",omitempty"`
//
//	// フィールドはこのパッケージによって無視されます。
//	Field int `json:"-"`
//
//	// フィールドはJSONでキー"-"として現れます。
//	Field int `json:"-,"`
//
<<<<<<< HEAD
// "string"オプションは、フィールドがJSONエンコードされた文字列内にJSONとして格納されることを示します。
// これは、文字列、浮動小数点数、整数、またはブール型のフィールドにのみ適用されます。
// この追加のエンコーディングレベルは、JavaScriptプログラムと通信する際に時々使用されます：
=======
// The "omitzero" option specifies that the field should be omitted
// from the encoding if the field has a zero value, according to rules:
//
// 1) If the field type has an "IsZero() bool" method, that will be used to
// determine whether the value is zero.
//
// 2) Otherwise, the value is zero if it is the zero value for its type.
//
// If both "omitempty" and "omitzero" are specified, the field will be omitted
// if the value is either empty or zero (or both).
//
// The "string" option signals that a field is stored as JSON inside a
// JSON-encoded string. It applies only to fields of string, floating point,
// integer, or boolean types. This extra level of encoding is sometimes used
// when communicating with JavaScript programs:
>>>>>>> upstream/release-branch.go1.25
//
//	Int64String int64 `json:",string"`
//
// キー名は、Unicodeの文字、数字、および引用符、バックスラッシュ、カンマを除くASCIIの句読点のみで構成される
// 空でない文字列の場合に使用されます。
//
// 埋め込まれた構造体のフィールドは、通常、その内部のエクスポートされたフィールドが
// 外部の構造体のフィールドであるかのようにマーシャルされます。これは、次の段落で説明されるように
// 通常のGoの可視性ルールを修正したものに従います。
// JSONタグで名前が指定された匿名の構造体フィールドは、匿名ではなく、その名前を持つものとして扱われます。
// インターフェース型の匿名の構造体フィールドは、匿名ではなく、その型の名前を持つものとして同様に扱われます。
//
// 構造体フィールドのマーシャルまたはアンマーシャルを決定する際に、
// JSONに対してGoの可視性ルールが修正されます。
// 同じレベルに複数のフィールドが存在し、そのレベルが最もネストが少ない
// （したがって、通常のGoのルールによって選択されるネストレベル）場合、
// 次の追加のルールが適用されます：
//
// 1) それらのフィールドの中で、JSONタグが付けられているものがある場合、
// それ以外に競合する未タグのフィールドが複数あっても、タグ付きのフィールドのみが考慮されます。
//
// 2) フィールドが1つだけ（最初のルールに従ってタグ付けされているかどうか）存在する場合、それが選択されます。
//
// 3) それ以外の場合、複数のフィールドが存在し、すべてが無視されます。エラーは発生しません。
//
// 匿名の構造体フィールドの扱いはGo 1.1で新しくなりました。
// Go 1.1より前では、匿名の構造体フィールドは無視されていました。現在のバージョンと以前のバージョンの両方で
// 匿名の構造体フィールドを強制的に無視するには、フィールドにJSONタグ "-" を付けてください。
//
// マップの値はJSONオブジェクトとしてエンコードされます。マップのキーの型は、
// 文字列、整数型、または [encoding.TextMarshaler] を実装する必要があります。マップのキーは
// ソートされ、上記の文字列値に対するUTF-8の強制に従って、以下のルールを適用して
// JSONオブジェクトのキーとして使用されます：
//   - 任意の文字列型のキーは直接使用されます
//   - [encoding.TextMarshalers] を実装しているキーはマーシャルされます
//   - 整数キーは文字列に変換されます
//
// ポインタ値は指している値としてエンコードされます。
// nilポインタはnullのJSON値としてエンコードされます。
//
// インターフェースの値は、インターフェースに含まれる値としてエンコードされます。
// nilのインターフェース値は、nullのJSON値としてエンコードされます。
//
// チャネル、複素数、および関数の値はJSONでエンコードすることはできません。
// そのような値をエンコードしようとすると、Marshalは
// [UnsupportedTypeError] を返します。
//
// JSONは循環データ構造を表現することはできませんし、Marshalはそれらを処理しません。
// 循環構造をMarshalに渡すとエラーが発生します。
func Marshal(v any) ([]byte, error)

// MarshalIndentは [Marshal] と同様ですが、出力のフォーマットに [Indent] を適用します。
// 出力の各JSON要素は、インデントのネストに従ってprefixで始まり、
// その後にindentの1つ以上のコピーが続く新しい行で始まります。
func MarshalIndent(v any, prefix, indent string) ([]byte, error)

// Marshalerは、自身を有効なJSONにマーシャルできる型が実装するインターフェースです。
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// UnsupportedTypeErrorは、サポートされていない値の型をエンコードしようとしたときに
// [Marshal] によって返されます。
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string

// UnsupportedValueErrorは、サポートされていない値をエンコードしようとしたときに
// [Marshal] によって返されます。
type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}

func (e *UnsupportedValueError) Error() string

// Go 1.2より前では、InvalidUTF8Errorは、無効なUTF-8シーケンスを含む文字列値をエンコードしようとしたときに
// [Marshal] によって返されました。Go 1.2以降では、[Marshal] は代わりに無効なバイトをUnicodeの置換ルーンU+FFFDで
// 置き換えることにより、文字列を有効なUTF-8に強制します。
//
// Deprecated: もう使用されていません。互換性のために保持されています。
type InvalidUTF8Error struct {
	S string
}

func (e *InvalidUTF8Error) Error() string

// MarshalerErrorは、[Marshaler.MarshalJSON] または [encoding.TextMarshaler.MarshalText] メソッドを呼び出す際のエラーを表します。
type MarshalerError struct {
	Type       reflect.Type
	Err        error
	sourceFunc string
}

func (e *MarshalerError) Error() string

// Unwrapは基礎となるエラーを返します。
func (e *MarshalerError) Unwrap() error
