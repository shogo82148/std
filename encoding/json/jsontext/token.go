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
// Token は配列やオブジェクト全体の値を表せませんが、[Value] は表せます。
// カンマやコロンを表す Token はありません。
// これらの構造トークンは周囲のコンテキストから推測できるためです。
//
// Token は、次の 2 つの形式のいずれかでデータを保持します。
//
//   - 生の JSON テキストとして: [Decoder] の内部バッファを参照しており、
//     [Decoder.ReadToken] によってのみ生成されます。
//     このようなトークンは、その [Decoder] に対する次の任意のメソッド呼び出し
//     （たとえば [Decoder.PeekKind]、[Decoder.ReadToken]、[Decoder.ReadValue]、
//     または [Decoder.SkipValue]）までしか有効ではありません。
//     [Token.Clone] を呼び出すと、生テキストを独立した割り当て先へコピーでき、
//     後続の [Decoder] 呼び出し後も保持されます。
//
//   - 型付きの Go 値として: コンストラクタ関数（たとえば [String]、[Int]、
//     [Uint]、[Float]）によって生成される自己完結した表現です。
//     このようなトークンは無期限に有効であり、複製する必要はありません。
type Token struct {
	nonComparable

	// raw contains a reference to the raw decode buffer.
	// If non-nil, then its value takes precedence over str and num.
	// It is only valid if num == raw.previousOffsetStart().
	raw *decodeBuffer

	// str is the unescaped JSON string if num is zero.
	// Otherwise, it is "F", "f", "i", or "u" if num should be interpreted
	// as a float32, float64, int64, or uint64, respectively.
	str string

	// num is a float32, float64, int64, or uint64 stored as a uint64 value.
	// For floating-point values, it stores the raw IEEE-754 bit-pattern.
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

// Float32 は、JSON 数値を 32 ビット浮動小数点数として表す Token を構築します。
// 形式は ECMA-262 第 10 版の 7.1.12.1 節に従いますが、
// -0 は引き続き -0 として形式化されます。
// NaN、+Inf、-Inf の値は、それぞれ "NaN"、"Infinity"、"-Infinity" の
// JSON 文字列として表されます。
//
// ほとんどの JSON ライブラリと標準は、JSON 数値が 64 ビット浮動小数点数であることを
// 前提としています。32 ビット精度の使用は、対応するデコーダーがこの JSON 数値トークンが
// 32 ビット精度のみを持つことを想定している場合に限るべきです。
// それ以外の状況では、代わりに [Float] コンストラクタを使うことを推奨します。
func Float32(n float32) Token

// Float は、JSON 数値を 64 ビット浮動小数点数として表す Token を構築します。
// 形式は ECMA-262 第 10 版の 7.1.12.1 節および RFC 8785 の 3.2.2.3 節に従いますが、
// -0 は引き続き -0 として形式化されます。
// NaN、+Inf、-Inf の値は、それぞれ "NaN"、"Infinity"、"-Infinity" の
// JSON 文字列として表されます。
func Float(n float64) Token

// Intは、int64からJSON数値を表すTokenを構築します。
func Int(n int64) Token

// Uintは、uint64からJSON数値を表すTokenを構築します。
func Uint(n uint64) Token

// Clone は、[Decoder] バッファに裏打ちされていない値を持つトークンのコピーを返します。
// そのため、後続の [Decoder] 呼び出し後も有効なままです。
// コンストラクタ関数によって生成されたトークンは、すでに自己完結しているため、
// このメソッドを呼んでも影響はありません。
func (t Token) Clone() Token

// Boolは、JSONの真偽値を返します。
// トークンの種類がJSONの真偽値でない場合はパニックになります。
func (t Token) Bool() bool

// Stringは、JSON文字列のエスケープされていない文字列値を返します。
// 他のJSONの種類の場合、これが生のJSON表現を返します。
func (t Token) String() string

// Float32 は、32 ビット精度として解釈された JSON 数値の浮動小数点値を返します。
//
// JSON 数値が float32 の表現可能範囲外である場合、
// +Inf または -Inf を返し、あわせて [errors.Is] に従って
// [strconv.ErrRange] に一致するエラーを返します。
//
// 値が "NaN"、"Infinity"、"-Infinity" の JSON 文字列に対しては、
// NaN、+Inf、-Inf の値を返します。
//
// トークン種別が JSON 数値でも、前述の値を持つ JSON 文字列でもない場合は
// パニックになります。
//
// ほとんどの JSON ライブラリと標準は、JSON 数値が 64 ビット浮動小数点数であることを
// 前提としています。このメソッドは、呼び出し元が別の文脈から、このトークンが
// 32 ビット精度だけで形式化された JSON 数値であることを知っている場合にのみ
// 使用するべきです（たとえば [Float32] コンストラクタでエンコードされた場合など）。
// それ以外の状況では、代わりに [Token.Float] アクセサを使うことを推奨します。
func (t Token) Float32() (float32, error)

// Float は、64 ビット精度として解釈された JSON 数値の浮動小数点値を返します。
//
// JSON 数値が float64 の表現可能範囲外である場合、
// +Inf または -Inf を返し、あわせて [errors.Is] に従って
// [strconv.ErrRange] に一致するエラーを返します。
//
// 値が "NaN"、"Infinity"、"-Infinity" の JSON 文字列に対しては、
// NaN、+Inf、-Inf の値を返します。
//
// トークン種別が JSON 数値でも、前述の値を持つ JSON 文字列でもない場合は
// パニックになります。
func (t Token) Float() (float64, error)

// Int は、JSON 数値の符号付き整数値を返します。
//
// JSON 数値が符号付き整数だけから成る制限された文法に一致しない場合、
// [errors.Is] に従って [strconv.ErrSyntax] に一致するエラーを報告します。
// JSON 数値が符号付き整数であっても int64 の範囲外である場合、
// [errors.Is] に従って [strconv.ErrRange] に一致するエラーを報告します。
// エラーが報告された場合でも、妥当な値は返されます。
// 数値の小数部分は無視されます（ゼロ方向への切り捨て）。
// int64 の表現を超える数値は、最も近い表現可能な値に飽和されます。
//
// トークン種別が JSON 数値でない場合はパニックになります。
func (t Token) Int() (int64, error)

// Uint は、JSON 数値の符号なし整数値を返します。
//
// JSON数値が符号なし整数だけから成る制限された文法に一致しない場合、
// [errors.Is] に従って [strconv.ErrSyntax] に一致するエラーを報告します。
// JSON数値が符号なし整数であっても uint64 の表現可能範囲外である場合、
// [errors.Is] に従って [strconv.ErrRange] に一致するエラーを報告します。
// エラーが報告された場合でも、妥当な値は返されます。
// 数値の小数部分は無視されます（ゼロ方向への切り捨て）。
// uint64の表現を超える数値は、最も近い表現可能な値に飽和されます。
//
// トークン種別が JSON 数値でない場合はパニックになります。
func (t Token) Uint() (uint64, error)

// Kindは、トークンの種類を返します。
func (t Token) Kind() Kind

// KindはJSONトークンの種類を表します。
//
// Kindは各JSONトークンの種類を1バイトで表現し、便利なことにその種類の文法の
// 最初のバイトになっています。ただし、数値は常に'0'で表現されるという制限があります。
type Kind byte

const (
	KindInvalid     Kind = 0
	KindNull        Kind = 'n'
	KindFalse       Kind = 'f'
	KindTrue        Kind = 't'
	KindString      Kind = '"'
	KindNumber      Kind = '0'
	KindBeginObject Kind = '{'
	KindEndObject   Kind = '}'
	KindBeginArray  Kind = '['
	KindEndArray    Kind = ']'
)

// Stringは種類を人間が読みやすい形で出力します。
func (k Kind) String() string
