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
// [JoinOptions]は複数のオプション値を合成します:
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
//   - [StringifyNumbers] はマーシャル・アンマーシャル両方に影響
//   - [Deterministic] はマーシャルのみ
//   - [FormatNilSliceAsNull] はマーシャルのみ
//   - [FormatNilMapAsNull] はマーシャルのみ
//   - [OmitZeroStructFields] はマーシャルのみ
//   - [MatchCaseInsensitiveNames] はマーシャル・アンマーシャル両方に影響
//   - [DiscardUnknownMembers] はマーシャルのみ
//   - [RejectUnknownMembers] はアンマーシャルのみ
//   - [WithMarshalers] はマーシャルのみ
//   - [WithUnmarshalers] はアンマーシャルのみ
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

// DefaultOptionsV2はv2セマンティクスを定義するすべてのオプションの完全なセットです。
// これは [Options]、[encoding/json.Options]、[encoding/json/jsontext.Options] の
// すべてのオプションがfalseまたはゼロ値に設定されている状態と同等ですが、
// 空白フォーマットに関連するオプションは除きます。
func DefaultOptionsV2() Options

// StringifyNumbersは、数値型のGo値を対応するJSON数値を含むJSON文字列としてマーシャルすることを指定します。
// アンマーシャル時には、数値型のGo値を、余分な空白を含まないJSON数値を含むJSON文字列からパースします。
//
// RFC 8259のセクション6によると、JSON実装はJSON数値の表現をIEEE 754 binary64値に制限する場合があります。
// これにより、int64やuint64型の精度が失われることがあります。
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

// FormatNilSliceAsNullは、nilのGoスライスをデフォルトの空JSON配列（~[]byteの場合は空JSON文字列）ではなく、JSON nullとしてマーシャルすることを指定します。
// `format:emitempty`が明示的に指定されたスライスフィールドは、引き続き空のJSON配列としてマーシャルされます。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func FormatNilSliceAsNull(v bool) Options

// FormatNilMapAsNullは、nilのGoマップをデフォルトの空JSONオブジェクトではなく、JSON nullとしてマーシャルすることを指定します。
// `format:emitempty`が明示的に指定されたマップフィールドは、引き続き空のJSONオブジェクトとしてマーシャルされます。
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

// DiscardUnknownMembersは、Go構造体の未知のJSONオブジェクトメンバーを格納する専用フィールドに保存された
// JSONオブジェクトメンバーをマーシャル時に無視することを指定します。
//
// このオプションはマーシャル時のみ影響し、アンマーシャル時は無視されます。
func DiscardUnknownMembers(v bool) Options

// RejectUnknownMembersは、未知のメンバーが格納用フィールドの有無に関わらず
// JSONオブジェクトのアンマーシャル時に拒否されるべきであることを指定します。
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
