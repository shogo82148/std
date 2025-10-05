// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/encoding/json/internal/jsonopts"
)

// Optionsは [NewEncoder]、[Encoder.Reset]、[NewDecoder]、
// および [Decoder.Reset] を特定の機能で構成します。
// 各関数は可変長のオプションリストを受け取り、
// 後から指定されたオプションのプロパティが、以前に設定された値を上書きします。
//
// Options型はエンコードとデコードの両方で使用されます。
// 一部のオプションは両方の操作に影響し、他は一方の操作のみに影響します:
//
//   - [AllowDuplicateNames] はエンコードとデコードに影響します
//   - [AllowInvalidUTF8] はエンコードとデコードに影響します
//   - [EscapeForHTML] はエンコードのみに影響します
//   - [EscapeForJS] はエンコードのみに影響します
//   - [PreserveRawStrings] はエンコードのみに影響します
//   - [CanonicalizeRawInts] はエンコードのみに影響します
//   - [CanonicalizeRawFloats] はエンコードのみに影響します
//   - [ReorderRawObjects] はエンコードのみに影響します
//   - [SpaceAfterColon] はエンコードのみに影響します
//   - [SpaceAfterComma] はエンコードのみに影響します
//   - [Multiline] はエンコードのみに影響します
//   - [WithIndent] はエンコードのみに影響します
//   - [WithIndentPrefix] はエンコードのみに影響します
//
// 特定の操作に影響しないオプションは無視されます。
//
// Options型は [encoding/json.Options] および [encoding/json/v2.Options] と同一です。
// 他のパッケージのOptionsをこのパッケージの機能に渡すことはできますが、無視されます。
// このパッケージのOptionsは他のパッケージでも使用できます。
type Options = jsonopts.Options

// AllowDuplicateNamesは、JSONオブジェクトが重複したメンバー名を含むことを許可します。
// 重複名のチェックを無効にするとパフォーマンス向上の可能性がありますが、RFC 7493の2.3節に準拠しなくなります。
// 入力や出力はRFC 8259には準拠し続けますが、重複名の扱いは未定義の動作となります。
//
// このオプションはエンコードまたはデコードのいずれかに影響します。
func AllowDuplicateNames(v bool) Options

// AllowInvalidUTF8は、JSON文字列に不正なUTF-8が含まれることを許可します。
// 不正なUTF-8はUnicodeの置換文字U+FFFDとして扱われます。
// このオプションを有効にすると、エンコーダやデコーダは
// RFC 7493の2.1節およびRFC 8259の8.1節に準拠しなくなります。
//
// このオプションはエンコードまたはデコードのいずれかに影響します。
func AllowInvalidUTF8(v bool) Options

// EscapeForHTMLは、JSON文字列内の '<', '>', '&' の各文字を
// 16進数のUnicodeコードポイント（例: \u003c）としてエスケープし、
// 出力がHTML内に安全に埋め込めるようにします。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func EscapeForHTML(v bool) Options

// EscapeForJSは、JSON文字列内のU+2028およびU+2029の文字を
// 16進数のUnicodeコードポイント（例: \u2028）としてエスケープし、
// 出力がJavaScript内に安全に埋め込めるようにします。RFC 8259の12節を参照。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func EscapeForJS(v bool) Options

// PreserveRawStringsは、生のJSON文字列を [Token] や [Value] でエンコードする際、
// JSON文字列内の事前にエスケープされたシーケンスを出力にそのまま保持します。
// ただし、生の文字列でも [EscapeForHTML] や [EscapeForJS] が有効な場合は、該当する文字がエスケープされます。
// [AllowInvalidUTF8] が有効な場合、不正なUTF-8のバイトも出力に保持されます。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func PreserveRawStrings(v bool) Options

// CanonicalizeRawIntsは、生のJSON整数（小数点や指数部を持たない数値）を
// [Token] や [Value] でエンコードする際、RFC 8785の3.2.2.3節に従って正規化します。
// 特別なケースとして、-0は0として正規化されます。
//
// JSONの数値はIEEE 754倍精度浮動小数点数として扱われます。
// この形式で表現できる範囲を超える精度の数値は、正規化時に精度が失われます。
// 例えば、±2⁵³を超える整数値は精度が失われます。
// 例：1234567890123456789は1234567890123456800としてフォーマットされます。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func CanonicalizeRawInts(v bool) Options

// CanonicalizeRawFloatsは、生のJSON浮動小数点数（小数部や指数部を持つ数値）を
// [Token] や [Value] でエンコードする際、RFC 8785の3.2.2.3節に従って正規化します。
// 特別なケースとして、-0は0として正規化されます。
//
// JSONの数値はIEEE 754倍精度浮動小数点数として扱われます。
// シリアライズされた単精度数値を正規化し、再度単精度としてパースしても同じ値が得られます。
// ±1.7976931348623157e+308を超える数値は最大有限値で飽和し、その値としてフォーマットされます。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func CanonicalizeRawFloats(v bool) Options

// ReorderRawObjectsは、生のJSONオブジェクトを [Value] でエンコードする際、
// オブジェクトのメンバーをRFC 8785の3.2.3節に従って並べ替えます。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func ReorderRawObjects(v bool) Options

// SpaceAfterColonは、JSON出力で各コロン区切り（JSONオブジェクト名の後）に
// スペース文字を挿入するかどうかを指定します。
// falseの場合、コロンの後にスペース文字は挿入されません。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func SpaceAfterColon(v bool) Options

// SpaceAfterCommaは、JSON出力で各カンマ区切り（JSONオブジェクトの値や配列要素の後）に
// スペース文字を挿入するかどうかを指定します。
// falseの場合、カンマの後にスペース文字は挿入されません。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func SpaceAfterComma(v bool) Options

// Multilineは、JSON出力を複数行に展開することを指定します。
// すべてのJSONオブジェクトメンバーやJSON配列要素が、
// ネストの深さに応じて新しいインデント付きの行に表示されます。
//
// [SpaceAfterColon] が指定されていない場合、デフォルトはtrueです。
// [SpaceAfterComma] が指定されていない場合、デフォルトはfalseです。
// [WithIndent] が指定されていない場合、デフォルトは"\t"です。
//
// falseに設定すると、出力は1行のみとなり、
// 出力される空白文字は現在の [SpaceAfterColon] と [SpaceAfterComma] の値によって決まります。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
func Multiline(v bool) Options

// WithIndentは、エンコーダが複数行の出力を生成することを指定します。
// 各JSONオブジェクトや配列の要素は、新しいインデント付きの行で開始され、
// インデントプリフィックス（[WithIndentPrefix] 参照）の後に、
// ネストの深さに応じてインデント文字列が1回以上繰り返されます。
// インデント文字列はスペースまたはタブのみで構成されている必要があります。
//
// 特定のインデント文字列にこだわりがなく、インデントされた出力を生成したい場合は
// [Multiline] を使用してください。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
// このオプションを指定すると [Multiline] もtrueになります。
func WithIndent(indent string) Options

// WithIndentPrefixは、エンコーダが複数行の出力を生成する際、
// 各JSONオブジェクトや配列の要素が新しいインデント付きの行で開始され、
// インデントプリフィックスの後にインデント文字列がネストの深さに応じて1回以上繰り返されます
// （詳細は [WithIndent] を参照）。
// プリフィックスはスペースまたはタブのみで構成されている必要があります。
//
// このオプションはエンコード時のみ有効で、デコード時には無視されます。
// このオプションを指定すると [Multiline] もtrueになります。
func WithIndentPrefix(prefix string) Options
