// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// Tokenは字句的なJSONトークンを表します。次のいずれかになります:
//   - JSONリテラル（null, true, false）
//   - JSON文字列（例: "hello, world!"）
//   - JSON数値（例: 123.456）
//   - JSONオブジェクトの開始・終了デリミタ（{ または }）
//   - JSON配列の開始・終了デリミタ（[ または ]）
//
// Tokenは配列やオブジェクト全体の値は表せませんが、[Value] は表せます。
// カンマやコロンを表すTokenはありません。
// これらの構造的トークンは周囲のコンテキストから推測できます。
type Token struct {
	nonComparable

	// raw contains a reference to the raw decode buffer.
	// If non-nil, then its value takes precedence over str and num.
	// It is only valid if num == raw.previousOffsetStart().
	raw *decodeBuffer

	// str is the unescaped JSON string if num is zero.
	// Otherwise, it is "f", "i", or "u" if num should be interpreted
	// as a float64, int64, or uint64, respectively.
	str string

	// num is a float64, int64, or uint64 stored as a uint64 value.
	// It is non-zero for any JSON number in the "exact" form.
	num uint64
}

var (
	Null  Token = rawToken("null")
	False Token = rawToken("false")
	True  Token = rawToken("true")

	BeginObject Token = rawToken("{")
	EndObject   Token = rawToken("}")
	BeginArray  Token = rawToken("[")
	EndArray    Token = rawToken("]")
)

// Boolは、JSONの真偽値を表すTokenを構築します。
func Bool(b bool) Token

// Stringは、JSON文字列を表すTokenを構築します。
// 渡された文字列は有効なUTF-8である必要があります。そうでない場合、
// 不正な文字はUnicodeの置換文字として扱われることがあります。
func String(s string) Token

// Floatは、JSON数値を表すTokenを構築します。
// NaN、+Inf、-Infの値は、"NaN"、"Infinity"、"-Infinity"という値のJSON文字列として表現されます。
func Float(n float64) Token

// Intは、int64からJSON数値を表すTokenを構築します。
func Int(n int64) Token

// Uintは、uint64からJSON数値を表すTokenを構築します。
func Uint(n uint64) Token

// Cloneは、Tokenのコピーを作成し、後続の [Decoder.Read] 呼び出し後も値が有効であることを保証します。
func (t Token) Clone() Token

// Boolは、JSONの真偽値を返します。
// トークンの種類がJSONの真偽値でない場合はパニックになります。
func (t Token) Bool() bool

// Stringは、JSON文字列のエスケープされていない文字列値を返します。
// 他のJSONの種類の場合、これが生のJSON表現を返します。
func (t Token) String() string

// Floatは、JSON数値の浮動小数点値を返します。
// "NaN"、"Infinity"、"-Infinity"という値のJSON文字列に対しては、NaN、+Inf、-Infの値を返します。
// その他の場合はパニックになります。
func (t Token) Float() float64

// Intは、JSON数値の符号付き整数値を返します。
// 数値の小数部分は無視されます（ゼロ方向への切り捨て）。
// int64で表現できない数値は、最も近い表現可能な値に丸められます。
// トークンの種類がJSON数値でない場合はパニックになります。
func (t Token) Int() int64

// Uintは、JSON数値の符号なし整数値を返します。
// 数値の小数部分は無視されます（ゼロ方向への切り捨て）。
// uint64で表現できない数値は、最も近い表現可能な値に丸められます。
// トークンの種類がJSON数値でない場合はパニックになります。
func (t Token) Uint() uint64

// Kindは、トークンの種類を返します。
func (t Token) Kind() Kind

// Kindは、すべてのJSONトークン種別を1バイトで表します。
// このバイトは各種別の文法の先頭バイトであり、
// 数値は必ず '0' で表現されるという制約があります:
//
//   - 'n': null
//   - 'f': false
//   - 't': true
//   - '"': string
//   - '0': number
//   - '{': オブジェクト開始
//   - '}': オブジェクト終了
//   - '[': 配列開始
//   - ']': 配列終了
//
// 無効な種別は通常0で表されますが、
// 不正なJSONデータの場合は0以外になることもあります。
type Kind byte

// Stringは、種別を人間が読める形式で表示します。
func (k Kind) String() string
