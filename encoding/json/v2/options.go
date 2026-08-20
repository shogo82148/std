// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding/json/internal/jsonopts"
)

// Optionsは [Marshal], [MarshalWrite], [MarshalEncode],
// [Unmarshal], [UnmarshalRead], [UnmarshalDecode] を特定の機能で設定します。
// 各関数は可変長のオプションリストを受け取り、後から指定されたオプションのプロパティが
// 以前に設定された値を上書きします。
//
// Options型は [encoding/json.Options] や [encoding/json/jsontext.Options] と同一です。
// 他のパッケージのOptionsも本パッケージの機能と相互利用できます。
//
// Options 値は、単一のオプションまたはオプションの集合を表します。
// これはオプションのプロパティを持つ Go のマップと考えることができます
// （ただし、実装の基盤はパフォーマンスのために Go のマップを避けています）。
//
// コンストラクタ（例: [Deterministic]）は、単一のオプション用の値を返します:
//
//	opt := Deterministic(true)
//
// これは単一エントリのマップを作ることに相当します:
//
//	opt := Options{"Deterministic": true}
//
// [JoinOptions] は複数のオプション値を1つにまとめます:
//
//	out := JoinOptions(opts...)
//
// これは新しいマップを作り、オプションをコピーすることに相当します:
//
//	out := make(Options)
//	for _, m := range opts {
//		for k, v := range m {
//			out[k] = v
//		}
//	}
//
// [GetOption] はオプションパラメータの値を取得します:
//
//	v, ok := GetOption(opts, Deterministic)
//
// これはGoのマップ検索に相当します:
//
//	v, ok := Options["Deterministic"]
//
// Options型はマーシャルとアンマーシャルの両方で使われます。
// 一部のオプションは両方に影響し、他はどちらか一方だけに影響します:
//
//   - [StringifyNumbers] はマーシャルとアンマーシャルの両方に影響します
//   - [Deterministic] はマーシャル時のみ影響します
//   - [FormatNilSliceAsNull] はマーシャル時のみ影響します
//   - [FormatNilMapAsNull] はマーシャル時のみ影響します
//   - [OmitZeroStructFields] はマーシャル時のみ影響します
//   - [MatchCaseInsensitiveNames] はマーシャルとアンマーシャルの両方に影響します
//   - [RejectUnknownMembers] はアンマーシャル時のみ影響します
//   - [WithMarshalers] はマーシャル時のみ影響します
//   - [WithUnmarshalers] はアンマーシャル時のみ影響します
//
// 特定の操作に影響しないオプションは無視されます。
type Options = jsonopts.Options

// JoinOptionsは、指定されたオプションリストを1つのOptionsにまとめます。
// 後から指定されたオプションのプロパティが、以前に設定された値を上書きします。
func JoinOptions(srcs ...Options) Options

// GetOptionは、指定されたsetterに対応する値をoptsから取得し、
// その値が存在するかどうかも返します。
// 存在しない場合、返される値は型Tのゼロ値です。
//
// 使用例:
//
//	v, ok := json.GetOption(opts, json.Deterministic)
//
// Optionsは主に、[MarshalerTo.MarshalJSONTo] や [UnmarshalerFrom.UnmarshalJSONFrom] メソッド、
// [MarshalToFunc] や [UnmarshalFromFunc] 関数のJSON表現を変更するために調査されます。
// その場合、存在ビットは通常無視されるべきです。
func GetOption[T any](opts Options, setter func(T) Options) (T, bool)

// DefaultOptionsV2は、v2のセマンティクスを定義するすべてのオプションセットです。
// これは [encoding/json.DefaultOptionsV1] のオプションセットのすべてがfalseに設定されているのと同等です。
// その他のすべてのオプションは存在しません。
func DefaultOptionsV2() Options

// StringifyNumbers は、通常は JSON 数値としてエンコードされる型が、
// 代わりに等価な JSON 数値を含む JSON 文字列としてエンコードされることを指定します。
// アンマーシャル時には、前後に空白を含まない JSON 数値を含む
// JSON 文字列から値がパースされます。
//
// Go 構造体フィールドに `string` タグオプションを指定した場合、
// このオプションはそのフィールドの最上位の JSON 値に適用されます。
// `string` タグオプション経由で適用された場合、
// StringifyNumbers オプションは JSON オブジェクトまたは配列内の
// ネストした JSON 数値には再帰的には適用されません。
//
// 他のすべてのオプションと同様に、[Marshal] や [Unmarshal] などの呼び出しで
// このオプションを明示的に指定すると、再帰的に適用されます。
//
// JSON 数値を表す独自のマーシャル/アンマーシャルを持つ Go 型は、
// StringifyNumbers オプションを尊重し、指定されている場合は
// JSON 文字列内の JSON 数値としてシリアライズすべきです。
// カスタムのマーシャル/アンマーシャルは、ネストした JSON オブジェクトに対して
// [MarshalEncode]/[UnmarshalDecode] を使用して処理する必要があり、
// これにより非再帰的な `string` タグオプションの動作が自動的に適用されます。
//
// RFC 8259 のセクション6によると、JSON実装はJSON数値の表現を
// IEEE 754 binary64値に制限する場合があります。
// これにより、デコーダがint64型やuint64型の精度を失う可能性があります。
// JSON数値をJSON文字列として引用することで、正確な精度を保持できます。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
func StringifyNumbers(v bool) Options

// Deterministic は、同じ入力値をマーシャルした場合に、
// 常に同じ出力バイト列になることを指定します。
//
// たとえば、Go のマップはキー順にソートされてマーシャルされます。
//
// ネイティブな Go 型については、同一のバイナリの異なるインスタンス間では
// 決定性が保証されますが、同一のプログラムの異なるビルド間では
// 保証されません（たとえば、異なるソースやツールチェインのバージョン、
// 異なる GOOS/GOARCH、異なるビルドフラグなど）。
//
// カスタムのマーシャラーを持つ Go 型は、Deterministic オプションを
// 尊重し、true の場合は決定的にシリアライズすべきです。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func Deterministic(v bool) Options

// FormatNilSliceAsNullは、nilのGoスライスをJSON nullとしてマーシャルすることを指定し、
// デフォルトの空のJSON配列表現（または~[]byteの場合は空のJSON文字列）の代わりに使用されます。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func FormatNilSliceAsNull(v bool) Options

// FormatNilMapAsNullは、nilのGoマップをJSON nullとしてマーシャルすることを指定し、
// デフォルトの空のJSONオブジェクト表現の代わりに使用されます。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func FormatNilMapAsNull(v bool) Options

// OmitZeroStructFields は、Go 構造体のゼロ値フィールドを
// マーシャルされた出力から省略することを指定します。
// 値がゼロであるとみなされるのは、その型が "IsZero() bool" メソッドを持ち、
// その結果が true を返す場合、またはそのメソッドを持たず、かつ値が Go のゼロ値である場合です。
// このオプションは、Go 構造体のすべてのフィールドに `omitzero` タグオプションを
// 指定するのと同等です。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func OmitZeroStructFields(v bool) Options

// MatchCaseInsensitiveNames は、JSON オブジェクトのメンバー名を
// Go 構造体フィールドに対して大文字と小文字を区別しない名前一致で
// 対応付けることを指定します。
// ある名前が複数のフィールドに一致する場合、
// その名前と完全一致するフィールドが選ばれます。
// 完全一致するフィールドがない場合は、エラーが報告されます。
// `case:strict` または `case:ignore` が明示されている Go 構造体フィールドは、
// このオプションの値に関係なく、常に大文字・小文字を区別する
// （または区別しない）名前一致を使用します。
//
// このオプションはマーシャルまたはアンマーシャルのいずれかに影響します。
//
// 名前を大文字・小文字を区別せずに一致させることは、
// 重複名の検出にも影響します（[jsontext.AllowDuplicateNames] が false と仮定）。
// これは、同じ名前の揺れが同じ Go 構造体フィールドに一致しうるためです。
// たとえばアンマーシャル時には、"foo" と "Foo" の両方が同じ
// Go 構造体フィールドに一致し、そのため重複名とみなされることがあります。
// マーシャル時には通常、2 つの Go 構造体フィールドが、
// それらをアンマーシャルしたときに同じ Go 構造体フィールドに入るような形で
// シリアライズされることはありません。
// ただし、埋め込みフォールバック内に、Go 構造体フィールド名にも一致する名前が含まれている場合、
// その結果として重複名エラーになる可能性があります。
func MatchCaseInsensitiveNames(v bool) Options

// RejectUnknownMembersは、JSONオブジェクトのアンマーシャル時に
// 未知のメンバーを拒否することを指定します。
//
// このオプションはアンマーシャル時のみ影響し、マーシャル時は無視されます。
func RejectUnknownMembers(v bool) Options

// WithMarshalersは、型ごとのマシャラーのリストを指定します。
// これにより、特定の型の値に対するデフォルトのマーシャル動作を上書きできます。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func WithMarshalers(v *Marshalers) Options

// WithUnmarshalersは、型ごとのアンマーシャラーのリストを指定します。
// これにより、特定の型の値に対するデフォルトのアンマーシャル動作を上書きできます。
//
// このオプションはアンマーシャル時のみ影響し、マーシャル時は無視されます。
func WithUnmarshalers(v *Unmarshalers) Options
