// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

// jsonパッケージはRFC 7159で定義されたJSONのエンコードとデコードを実装します。
// JSONとGoの値の対応関係については、MarshalおよびUnmarshal関数のドキュメントを参照してください。
//
// このパッケージの概要については「JSON and Go」を参照してください:
// https://golang.org/doc/articles/json_and_go.html
//
// # Security Considerations
//
// [encoding/json/v2] の「Security Considerations」セクションを参照してください。
//
// 歴史的な理由により、v1 [encoding/json] のデフォルト動作は
// 残念ながらセキュリティ面で安全性が低い設定となっています。
// Goで新しくJSONを利用する場合は [encoding/json/v2] の使用を推奨します。
package json

import (
	"github.com/shogo82148/std/reflect"
)

// MarshalはvのJSONエンコーディングを返します。
//
// Marshalは値vを再帰的に走査します。
// 走査中の値が [Marshaler] を実装していてnilポインタでない場合、Marshalは [Marshaler.MarshalJSON] を呼び出してJSONを生成します。
// [Marshaler.MarshalJSON] メソッドが存在しないが値が [encoding.TextMarshaler] を実装している場合、Marshalは [encoding.TextMarshaler.MarshalText] を呼び出し、その結果をJSON文字列としてエンコードします。
// nilポインタの例外は厳密には必要ありませんが、[Unmarshaler.UnmarshalJSON] の動作にある同様の必要な例外を模倣しています。
//
// それ以外の場合、Marshalは型ごとのデフォルトエンコーディングを使用します：
//
// ブール値はJSONの真偽値としてエンコードされます。
//
// 浮動小数点、整数、および [Number] 値はJSONの数値としてエンコードされます。
// NaNや±Inf値は [UnsupportedValueError] を返します。
//
// 文字列値は有効なUTF-8に強制されてJSON文字列としてエンコードされ、無効なバイトはUnicodeの置換文字に置き換えられます。
// JSONをHTMLの<script>タグ内に安全に埋め込むため、文字列は [HTMLEscape] を使ってエンコードされ、"<", ">", "&", U+2028, U+2029はそれぞれ"\u003c","\u003e", "\u0026", "\u2028", "\u2029"にエスケープされます。
// この置換は [Encoder.SetEscapeHTML](false) を呼び出すことで無効化できます。
//
// 配列とスライス値はJSON配列としてエンコードされます。ただし、[]byteはbase64エンコードされた文字列として、nilスライスはnull JSON値としてエンコードされます。
//
// 構造体値はJSONオブジェクトとしてエンコードされます。
// エクスポートされた各構造体フィールドはオブジェクトのメンバーとなり、フィールド名がオブジェクトキーになります（下記の理由で省略される場合を除く）。
//
// 各構造体フィールドのエンコーディングは、フィールドタグの"json"キーに格納されたフォーマット文字列でカスタマイズできます。
// フォーマット文字列はフィールド名と、カンマ区切りのオプションリストを指定できます。名前が空の場合はデフォルトのフィールド名を上書きせずにオプションのみ指定できます。
//
// "omitempty"オプションは、フィールド値が空の場合（false, 0, nilポインタ, nilインターフェース値、長さ0の配列・スライス・マップ・文字列）、エンコーディングからフィールドを省略します。
//
// 特別なケースとして、フィールドタグが"-"の場合は常に省略されます。
// カンマや引用符を含むJSON名、""や"-"と同じ名前は、シングルクォート文字列リテラルで指定できます。構文はGoのダブルクォート文字列リテラルと同じですが、区切りがシングルクォートです。
//
// 構造体フィールドタグの例と意味：
//
//	// フィールドはJSONでキー"myName"として現れます。
//	Field int `json:"myName"`
//
//	// フィールドはJSONでキー"myName"として現れ、値が空の場合はオブジェクトから省略されます。
//	Field int `json:"myName,omitempty"`
//
//	// フィールドはJSONでキー"Field"（デフォルト）として現れますが、値が空の場合は省略されます。
//	// 先頭のカンマに注意。
//	Field int `json:",omitempty"`
//
//	// このパッケージではフィールドは無視されます。
//	Field int `json:"-"`
//
//	// フィールドはJSONでキー"-"として現れます。
//	Field int `json:"'-'"`
//
// "omitzero"オプションは、フィールド値がゼロ値の場合にエンコーディングから省略します。判定ルール：
//
// 1) フィールド型に"IsZero() bool"メソッドがあれば、それでゼロ値か判定します。
// 2) それ以外は型のゼロ値ならゼロ値とみなします。
//
// "omitempty"と"omitzero"両方指定した場合、値が空またはゼロ値（または両方）なら省略されます。
//
// "string"オプションは、フィールドをJSONエンコードされた文字列内に格納することを示します。文字列、浮動小数点、整数、ブール型フィールドにのみ適用されます。
// この追加のエンコーディングはJavaScriptプログラムとの通信時などに使われます：
//
//	Int64String int64 `json:",string"`
//
// キー名は、空でなく、Unicodeの文字・数字・ASCII句読点（引用符、バックスラッシュ、カンマを除く）のみからなる場合に使用されます。
//
// 埋め込み構造体フィールドは、通常、内部のエクスポートされたフィールドが外側の構造体のフィールドとしてマーシャルされます（Goの可視性ルールに従う。ただし次の段落で修正あり）。
// JSONタグで名前が指定された匿名構造体フィールドは匿名ではなくその名前を持つものとして扱われます。
// インターフェース型の匿名構造体フィールドも、その型名を持つものとして扱われます。
//
// Goの構造体フィールドの可視性ルールは、JSONのマーシャル・アンマーシャル時に修正されます。
// 同じレベルに複数のフィールドがあり、そのレベルが最も浅い（通常のGoルールで選択されるネストレベル）場合、次の追加ルールが適用されます：
//
// 1) そのフィールドのうちJSONタグ付きがあれば、タグ付きのみが対象となります（タグなしフィールドが複数あっても競合は無視）。
// 2) 1つだけ（タグ付きか否かに関わらず）ならそれが選択されます。
// 3) それ以外は複数ある場合、すべて無視され、エラーは発生しません。
//
// 匿名構造体フィールドの扱いはGo 1.1で新しくなりました。
// Go 1.1以前は匿名構造体フィールドは無視されていました。両方のバージョンで匿名構造体フィールドを無視したい場合はJSONタグ"-"を付けてください。
//
// マップ値はJSONオブジェクトとしてエンコードされます。マップのキー型は文字列型、整数型、または [encoding.TextMarshaler] を実装している必要があります。
// マップキーはソートされ、JSONオブジェクトキーとして使われます。文字列値のUTF-8強制に関するルールに従います：
//   - 文字列型キーはそのまま使用
//   - [encoding.TextMarshaler] を実装するキーはマーシャルされる
//   - 整数キーは文字列に変換される
//
// ポインタ値は指す値としてエンコードされます。
// nilポインタはnull JSON値としてエンコードされます。
//
// インターフェース値はインターフェースに格納された値としてエンコードされます。
// nilインターフェース値はnull JSON値としてエンコードされます。
//
// チャネル、複素数、関数値はJSONでエンコードできません。
// これらの値をエンコードしようとするとMarshalは [UnsupportedTypeError] を返します。
//
// JSONは循環データ構造を表現できず、Marshalも対応しません。循環構造をMarshalに渡すとエラーになります。
func Marshal(v any) ([]byte, error)

// MarshalIndentは[Marshal]と同様ですが、[Indent] を適用して出力を整形します。
// 出力される各JSON要素は新しい行で始まり、prefixの後に
// インデントのネストに応じてindentが1回以上繰り返されます。
func MarshalIndent(v any, prefix, indent string) ([]byte, error)

// Marshalerは、自身を有効なJSONにマーシャルできる型が実装するインターフェースです。
type Marshaler = jsonv2.Marshaler

// UnsupportedTypeErrorは、サポートされていない値型をエンコードしようとした際に [Marshal] から返されます。
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string

// UnsupportedValueErrorは、サポートされていない値をエンコードしようとした際に [Marshal] から返されます。
type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}

func (e *UnsupportedValueError) Error() string

// Go 1.2以前は、無効なUTF-8シーケンスを含む文字列値をエンコードしようとした場合、[Marshal] はInvalidUTF8Errorを返していました。
// Go 1.2以降は、[Marshal] は無効なバイトをUnicodeの置換文字U+FFFDに置き換えて、文字列を有効なUTF-8に強制します。
//
// Deprecated: 互換性維持のために残されていますが、現在は使用されていません。
type InvalidUTF8Error struct {
	S string
}

func (e *InvalidUTF8Error) Error() string

// MarshalerErrorは、[Marshaler.MarshalJSON] または [encoding.TextMarshaler.MarshalText] メソッドの呼び出しによるエラーを表します。
type MarshalerError struct {
	Type       reflect.Type
	Err        error
	sourceFunc string
}

func (e *MarshalerError) Error() string

// Unwrapは元となるエラーを返します。
func (e *MarshalerError) Unwrap() error
