// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// AppendFormatは、src内のJSON値を指定されたオプションに従ってフォーマットし、
// dstに追加して返します。
// フォーマット動作の詳細については [Value.Format] を参照してください。
//
// dstとsrcは重複していても構いません。
// エラーが報告された場合は、src全体がdstに追加されます。
func AppendFormat(dst, src []byte, opts ...Options) ([]byte, error)

// Valueは1つの生のJSON値を表します。次のいずれかになります:
//   - JSONリテラル（null, true, false）
//   - JSON文字列（例: "hello, world!"）
//   - JSON数値（例: 123.456）
//   - JSONオブジェクト全体（例: {"fizz":"buzz"} ）
//   - JSON配列全体（例: [1,2,3] ）
//
// Valueは配列やオブジェクト全体の値を表せますが、[Token] は表せません。
// Valueは前後に空白を含む場合があります。
type Value []byte

// Cloneは、vのコピーを返します。
func (v Value) Clone() Value

// Stringは、vの文字列表現を返します。
func (v Value) String() string

// IsValidは、生のJSON値が指定されたオプションに従って構文的に有効かどうかを報告します。
//
// デフォルト（オプションを指定しない場合）はRFC 7493に従って検証します。
// 入力が正しくUTF-8でエンコードされているか、
// 文字列内のエスケープシーケンスが有効なUnicodeコードポイントにデコードされるか、
// 各オブジェクト内の名前がすべて一意であるかを検証します。
// 数値が一般的な数値型（例: float64, int64, uint64）の範囲内で表現可能かどうかは検証しません。
//
// 関連するオプション:
//   - [AllowDuplicateNames]
//   - [AllowInvalidUTF8]
//
// その他のオプションは無視されます。
func (v Value) IsValid(opts ...Options) bool

// Formatは、生のJSON値をその場でフォーマットします。
//
// デフォルト（オプションを指定しない場合）はRFC 7493に従って検証し、
// 最小限のJSON表現を生成します。すべての空白は除去され、
// JSON文字列は最短のエンコーディングで表現されます。
//
// 関連するオプション:
//   - [AllowDuplicateNames]
//   - [AllowInvalidUTF8]
//   - [EscapeForHTML]
//   - [EscapeForJS]
//   - [PreserveRawStrings]
//   - [CanonicalizeRawInts]
//   - [CanonicalizeRawFloats]
//   - [ReorderRawObjects]
//   - [SpaceAfterColon]
//   - [SpaceAfterComma]
//   - [Multiline]
//   - [WithIndent]
//   - [WithIndentPrefix]
//
// その他のオプションは無視されます。
//
// 同じオプションで値が有効な場合は必ず成功します。
// すでにフォーマット済みの場合、バッファは変更されません。
func (v *Value) Format(opts ...Options) error

// Compactは、生のJSON値からすべての空白を除去します。
//
// JSON文字列や数値の表現は変更しません。
// フォーマット可能なJSON値の集合を最大化するため、
// 重複した名前や不正なUTF-8を含む値も許容します。
//
// Compactは、以下のオプションを指定して[Value.Format]を呼び出すのと同等です:
//   - [AllowDuplicateNames](true)
//   - [AllowInvalidUTF8](true)
//   - [PreserveRawStrings](true)
//
// 呼び出し元が指定したオプションは初期セットの後に適用され、
// 意図的に前のオプションを上書きすることもできます。
func (v *Value) Compact(opts ...Options) error

// Indentは、生のJSON値内の空白を再フォーマットし、
// JSONオブジェクトや配列の各要素がネストに応じてインデントされた行で始まるようにします。
//
// JSON文字列や数値の表現は変更しません。
// フォーマット可能なJSON値の集合を最大化するため、
// 重複した名前や不正なUTF-8を含む値も許容します。
//
// Indentは、以下のオプションを指定して[Value.Format]を呼び出すのと同等です:
//   - [AllowDuplicateNames](true)
//   - [AllowInvalidUTF8](true)
//   - [PreserveRawStrings](true)
//   - [Multiline](true)
//
// 呼び出し元が指定したオプションは初期セットの後に適用され、
// 意図的に前のオプションを上書きすることもできます。
func (v *Value) Indent(opts ...Options) error

// Canonicalizeは、RFC 8785で定義されたJSON正規化スキーム（JCS）に従って
// 生のJSON値を正規化し、安定した表現を生成します。
//
// JSON文字列は最小表現でフォーマットされ、
// JSON数値は安定したシリアライズアルゴリズムに従い倍精度でフォーマットされます。
// JSONオブジェクトのメンバーは名前順に昇順でソートされます。
// すべての空白は除去されます。
//
// 出力の安定性はアプリケーションデータの安定性に依存します
// （RFC 8785の付録E参照）。本質的に不安定な入力からは
// 安定した出力を生成できません。例えば、JSON値に
// 一時的なデータ（頻繁に変化するタイムスタンプなど）が含まれている場合、
// この関数を呼び出しても値は不安定なままです。
//
// Canonicalizeは、以下のオプションを指定して [Value.Format] を呼び出すのと同等です:
//   - [CanonicalizeRawInts](true)
//   - [CanonicalizeRawFloats](true)
//   - [ReorderRawObjects](true)
//
// 呼び出し元が指定したオプションは初期セットの後に適用され、
// 意図的に前のオプションを上書きすることもできます。
//
// JCSはすべてのJSON数値をIEEE 754倍精度数値として扱うことに注意してください。
// この形式で表現できない精度を持つ数値は、正規化時に精度を失います。
// 例えば、±2⁵³を超える整数値は精度を失います。
// 元のJSON整数表現を保持したい場合は、追加で [CanonicalizeRawInts] をfalseに設定してください:
//
//	v.Canonicalize(jsontext.CanonicalizeRawInts(false))
func (v *Value) Canonicalize(opts ...Options) error

// MarshalJSONは、vをJSONエンコーディングとして返します。
// 検証を行わずに、格納されている値を生のJSON出力として返します。
// vがnilの場合は、JSONのnullを返します。
func (v Value) MarshalJSON() ([]byte, error)

// UnmarshalJSONは、bをJSONエンコーディングとしてvに設定します。
// 検証を行わずに、提供された生のJSON入力のコピーを格納します。
func (v *Value) UnmarshalJSON(b []byte) error

// Kindは、開始トークン種別を返します。
// 有効な値の場合、'}' や ']' を含むことはありません。
func (v Value) Kind() Kind
