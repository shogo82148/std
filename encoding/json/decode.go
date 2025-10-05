// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

//go:build !goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/reflect"
)

// Unmarshalは、JSONエンコードされたデータを解析し、結果をvが指す値に格納します。
// もしvがnilまたはポインタでない場合、Unmarshalは [InvalidUnmarshalError] を返します。
//
// Unmarshalは、[Marshal] が使用するエンコーディングの逆を使用し、
// 必要に応じてマップ、スライス、ポインタを割り当てます。
// 以下の追加ルールも適用されます：
//
// JSONをポインタにアンマーシャルするために、Unmarshalはまず
// JSONがJSONリテラルnullであるケースを処理します。その場合、Unmarshalは
// ポインタをnilに設定します。それ以外の場合、UnmarshalはJSONを
// ポインタが指す値にアンマーシャルします。もしポインタがnilなら、Unmarshalは
// それが指す新しい値を割り当てます。
//
// [Unmarshaler] インターフェースを実装する値にJSONをアンマーシャルするために、
// Unmarshalはその値の [Unmarshaler.UnmarshalJSON] メソッドを呼び出します、
// 入力がJSON nullである場合も含みます。
// それ以外の場合、もし値が [encoding.TextUnmarshaler] を実装していて、
// 入力がJSONの引用符で囲まれた文字列である場合、Unmarshalはその値の
// [encoding.TextUnmarshaler.UnmarshalText] を引用符で囲まれていない形式の文字列で呼び出します。
//
// JSONを構造体にアンマーシャルするために、Unmarshalは受信するオブジェクトキーを
// [Marshal]が使用するキー（構造体フィールド名またはそのタグ）にマッチさせ、
// 大文字小文字を無視します。複数の構造体フィールドがオブジェクトキーにマッチする場合、
// 大文字小文字を区別しないマッチよりも正確な大文字小文字マッチが優先されます。
//
// 受信するオブジェクトメンバーは、観察された順序で処理されます。オブジェクトに
// 重複したキーが含まれている場合、後の重複は以前の値を置き換えるか、マージされます。
//
// インターフェース値にJSONをアンマーシャルするために、
// Unmarshalは以下のいずれかをインターフェース値に格納します：
//
//   - bool, for JSON booleans
//   - float64, for JSON numbers
//   - string, for JSON strings
//   - []interface{}, for JSON arrays
//   - map[string]interface{}, for JSON objects
//   - nil for JSON null
//
// JSON配列をスライスにアンマーシャルするために、Unmarshalはスライスの長さを
// ゼロにリセットし、各要素をスライスに追加します。
// 特別なケースとして、空のJSON配列をスライスにアンマーシャルするために、
// Unmarshalはスライスを新しい空のスライスで置き換えます。
//
// Goの配列にJSON配列をアンマーシャルするために、Unmarshalは
// JSON配列の要素を対応するGoの配列の要素にデコードします。
// もしGoの配列がJSON配列より小さい場合、
// 追加のJSON配列の要素は破棄されます。
// もしJSON配列がGoの配列より小さい場合、
// 追加のGoの配列の要素はゼロ値に設定されます。
//
// JSONオブジェクトをマップにアンマーシャルするために、Unmarshalは最初に使用するマップを確立します。
// マップがnilの場合、Unmarshalは新しいマップを割り当てます。それ以外の場合、Unmarshalは
// 既存のマップを再利用し、既存のエントリを保持します。次に、UnmarshalはJSONオブジェクトから
// キーと値のペアをマップに格納します。マップのキーの型は、任意の文字列型、整数、
// または [encoding.TextUnmarshaler] を実装するものでなければなりません。
//
// もしJSONエンコードされたデータに構文エラーが含まれている場合、Unmarshalは [SyntaxError] を返します。
//
// もしJSON値が特定のターゲット型に適していない場合、
// またはJSON数値がターゲット型をオーバーフローする場合、Unmarshalは
// そのフィールドをスキップし、可能な限りアンマーシャルを完了します。
// もしもっと深刻なエラーが発生しなければ、Unmarshalは最初のそのようなエラーを
// 説明する [UnmarshalTypeError] を返します。いずれにせよ、問題のあるフィールドに続く
// すべてのフィールドがターゲットオブジェクトにアンマーシャルされることは保証されません。
//
// JSONのnull値は、そのGoの値をnilに設定することでインターフェース、マップ、ポインタ、スライスにアンマーシャルされます。
// nullはJSONで「存在しない」を意味することが多いため、JSONのnullを他のGoの型にアンマーシャルすると、
// 値には影響せず、エラーも発生しません。
//
// 引用符で囲まれた文字列をアンマーシャルするとき、無効なUTF-8または
// 無効なUTF-16サロゲートペアはエラーとして扱われません。
// 代わりに、それらはUnicodeの置換文字U+FFFDに置き換えられます。
func Unmarshal(data []byte, v any) error

// Unmarshalerは、自分自身のJSON記述をアンマーシャルできる型によって実装されるインターフェースです。
// 入力は、JSON値の有効なエンコーディングであると仮定できます。UnmarshalJSONは、
// 戻った後にデータを保持したい場合は、JSONデータをコピーしなければなりません。
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// UnmarshalTypeErrorは、特定のGo型の値に対して適切でないJSON値を説明します。
type UnmarshalTypeError struct {
	Value  string
	Type   reflect.Type
	Offset int64
	Struct string
	Field  string
}

func (e *UnmarshalTypeError) Error() string

// UnmarshalFieldErrorは、JSONオブジェクトキーが
// エクスポートされていない（したがって書き込み不可能な）構造体フィールドにつながることを説明します。
//
// Deprecated: もはや使用されていません。互換性のために保持されています。
type UnmarshalFieldError struct {
	Key   string
	Type  reflect.Type
	Field reflect.StructField
}

func (e *UnmarshalFieldError) Error() string

// InvalidUnmarshalErrorは、[Unmarshal] に渡された無効な引数を説明します。
// ([Unmarshal] への引数はnilでないポインタでなければなりません。)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string

// Numberは、JSONの数値リテラルを表します。
type Number string

// Stringは、数値のリテラルテキストを返します。
func (n Number) String() string

// Float64は、数値をfloat64として返します。
func (n Number) Float64() (float64, error)

// Int64は、数値をint64として返します。
func (n Number) Int64() (int64, error)
