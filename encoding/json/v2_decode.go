// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

package json

import (
	"github.com/shogo82148/std/reflect"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// UnmarshalはJSONエンコードされたデータを解析し、結果をvが指す値に格納します。
// vがnilまたはポインタでない場合、Unmarshalは [InvalidUnmarshalError] を返します。
//
// Unmarshalは [Marshal] が使用するエンコーディングの逆を使い、必要に応じてマップ、スライス、ポインタを割り当てます。
// 追加のルールは以下の通りです：
//
// ポインタ型にJSONをアンマーシャルする場合、まずJSONリテラルnullの場合を処理します。
// この場合、Unmarshalはポインタをnilに設定します。それ以外の場合、ポインタが指す値にアンマーシャルします。
// ポインタがnilの場合は新しい値を割り当てて指すようにします。
//
// [Unmarshaler] を実装する値にJSONをアンマーシャルする場合、Unmarshalはその値の [Unmarshaler.UnmarshalJSON] メソッドを呼び出します。
// 入力がJSON nullの場合も含みます。
// それ以外の場合、値が [encoding.TextUnmarshaler] を実装していて入力がJSONの引用符付き文字列なら、Unmarshalは
// [encoding.TextUnmarshaler.UnmarshalText] を文字列のアンエスケープ版で呼び出します。
//
// 構造体にJSONをアンマーシャルする場合、Unmarshalは受信したオブジェクトのキーを [Marshal] で使われるキー（フィールド名またはタグ）に一致させます。
// 完全一致を優先しますが、大文字小文字を無視した一致も受け入れます。
// デフォルトでは、対応するフィールドがないオブジェクトキーは無視されます（代替として [Decoder.DisallowUnknownFields] を参照）。
//
// インターフェース値にJSONをアンマーシャルする場合、Unmarshalは以下のいずれかをインターフェース値に格納します：
//
//   - bool（JSONの真偽値）
//   - float64（JSONの数値）
//   - string（JSONの文字列）
//   - []any（JSONの配列）
//   - map[string]any（JSONのオブジェクト）
//   - nil（JSONのnull）
//
// JSON配列をスライスにアンマーシャルする場合、Unmarshalはスライスの長さをゼロにリセットし、各要素をスライスに追加します。
// 特別なケースとして、空のJSON配列をスライスにアンマーシャルする場合、Unmarshalは新しい空のスライスに置き換えます。
//
// JSON配列をGo配列にアンマーシャルする場合、UnmarshalはJSON配列の要素を対応するGo配列の要素にデコードします。
// Go配列がJSON配列より小さい場合、余分なJSON配列要素は破棄されます。
// JSON配列がGo配列より小さい場合、余分なGo配列要素はゼロ値に設定されます。
//
// JSONオブジェクトをマップにアンマーシャルする場合、Unmarshalはまず使用するマップを決定します。
// マップがnilの場合は新しいマップを割り当てます。そうでなければ既存のマップを再利用し、既存のエントリを保持します。
// その後、JSONオブジェクトのキーと値のペアをマップに格納します。
// マップのキー型は任意の文字列型、整数型、または [encoding.TextUnmarshaler] を実装している必要があります。
//
// JSONエンコードされたデータに構文エラーが含まれている場合、Unmarshalは [SyntaxError] を返します。
//
// JSON値がターゲット型に適していない場合や、JSON数値がターゲット型でオーバーフローする場合、Unmarshalはそのフィールドをスキップし、可能な限りアンマーシャル処理を続行します。
// より重大なエラーがなければ、Unmarshalは最初に発生したエラーを説明する [UnmarshalTypeError] を返します。
// いずれにせよ、問題のあるフィールド以降の残りのフィールドがターゲットオブジェクトにアンマーシャルされる保証はありません。
//
// JSONのnull値は、インターフェース、マップ、ポインタ、スライスにアンマーシャルする場合、Go値をnilに設定します。
// nullはJSONで「存在しない」を意味することが多いため、他のGo型にアンマーシャルする場合は値に影響せず、エラーも発生しません。
//
// 引用符付き文字列をアンマーシャルする際、無効なUTF-8や無効なUTF-16サロゲートペアはエラーとして扱われません。
// 代わりにUnicodeの置換文字U+FFFDに置き換えられます。
func Unmarshal(data []byte, v any) error

// Unmarshalerは、自身のJSON記述をアンマーシャルできる型が実装するインターフェースです。
// 入力は有効なJSON値のエンコーディングであるとみなせます。
// UnmarshalJSONは、戻り値の後もデータを保持したい場合はJSONデータをコピーする必要があります。
type Unmarshaler = jsonv2.Unmarshaler

// UnmarshalTypeErrorは、特定のGo型の値として不適切なJSON値を説明します。
type UnmarshalTypeError struct {
	Value  string
	Type   reflect.Type
	Offset int64
	Struct string
	Field  string
	Err    error
}

func (e *UnmarshalTypeError) Error() string

func (e *UnmarshalTypeError) Unwrap() error

// UnmarshalFieldErrorは、JSONオブジェクトのキーが
// エクスポートされていない（そのため書き込み不可な）構造体フィールドに
// 対応していたことを説明します。
//
// Deprecated: 互換性維持のために残されていますが、現在は使用されていません。
type UnmarshalFieldError struct {
	Key   string
	Type  reflect.Type
	Field reflect.StructField
}

func (e *UnmarshalFieldError) Error() string

// InvalidUnmarshalErrorは[Unmarshal]に渡された無効な引数を説明します。
// （[Unmarshal]への引数はnilでないポインタでなければなりません。）
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string

// NumberはJSONの数値リテラルを表します。
type Number string

// Stringは数値のリテラルテキストを返します。
func (n Number) String() string

// Float64は数値をfloat64として返します。
func (n Number) Float64() (float64, error)

// Int64は数値をint64として返します。
func (n Number) Int64() (int64, error)

// MarshalJSONToは [jsonv2.MarshalerTo] を実装します。
func (n Number) MarshalJSONTo(enc *jsontext.Encoder) error

// UnmarshalJSONFromは [jsonv2.UnmarshalerFrom] を実装します。
func (n *Number) UnmarshalJSONFrom(dec *jsontext.Decoder) error
