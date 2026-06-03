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
// Optionsは単一のオプションまたはオプションの集合を表します。
// 機能的にはGoのオプションプロパティのマップのように考えられます
// （実装上はパフォーマンスのためGoのmapは使っていません）。
//
// コンストラクタ（例: [Deterministic]）は単一のオプション値を返します:
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

// GetOptionは、optsに格納されたsetterで指定された値を返し、
// その値が存在するかどうかを報告します。
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

// StringifyNumbersは、通常JSON数値としてエンコードされる型を、
// 同等のJSON数値を含むJSON文字列としてエンコードすることを指定します。
// アンマーシャル時には、前後に空白を含まないJSON数値を含む
// JSON文字列から値をパースします。
//
// Go構造体フィールドに `string` タグオプションが指定されている場合、
// このオプションはそのフィールドの最上位JSON値に適用されます。
// StringifyNumbers がグローバルに適用されていない限り、
// JSONオブジェクトまたは配列内にネストされたJSON数値には
// 再帰的には適用されません。
// JSON数値を表現するカスタムのマーシャル/アンマーシャルを持つGo型は、
// StringifyNumbers オプションを尊重し、指定されている場合は
// JSON文字列内のJSON数値としてシリアライズすべきです。
//
// RFC 8259 のセクション6によると、JSON実装はJSON数値の表現を
// IEEE 754 binary64値に制限する場合があります。
// これにより、デコーダがint64型やuint64型の精度を失う可能性があります。
// JSON数値をJSON文字列として引用することで、正確な精度を保持できます。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
func StringifyNumbers(v bool) Options

// Deterministicは、同じ入力値が常に同じ出力バイト列としてシリアライズされることを指定します。
// 同じプログラムの異なるプロセスは、等価な値を同じバイト列にシリアライズしますが、
// プログラムのバージョンが異なる場合は、必ずしも同じバイト列になるとは限りません。
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

// OmitZeroStructFieldsは、Go構造体のゼロ値フィールドをマーシャル出力から省略することを指定します。
// ゼロ値かどうかは "IsZero() bool" メソッドがあればその結果で、なければGoのゼロ値かどうかで判定します。
// これはすべてのフィールドに `omitzero` タグオプションを指定するのと同等です。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func OmitZeroStructFields(v bool) Options

// MatchCaseInsensitiveNamesは、JSONオブジェクトのメンバー名をGo構造体のフィールド名と
// 大文字・小文字を区別せずに一致させることを指定します。
// Go構造体フィールドに `case:strict` や `case:ignore` が明示的に指定されている場合は、
// このオプションの値に関わらず、それぞれ大文字・小文字を区別する（または区別しない）一致が使われます。
//
// このオプションはマーシャル・アンマーシャルの両方に影響します。
// マーシャル時には、インラインフィールドから宣言済みフィールドと一致する場合、
// （[jsontext.AllowDuplicateNames]がfalseの場合）重複名の検出方法が変わることがあります。
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
