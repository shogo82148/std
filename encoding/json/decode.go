// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

package json

import (
	"github.com/shogo82148/std/reflect"
)

<<<<<<< HEAD
// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an [InvalidUnmarshalError].
//
// Unmarshal uses the inverse of the encodings that
// [Marshal] uses, allocating maps, slices, and pointers as necessary,
// with the following additional rules:
=======
// Unmarshalは、JSONエンコードされたデータを解析し、結果をvが指す値に格納します。
// もしvがnilまたはポインタでない場合、UnmarshalはInvalidUnmarshalErrorを返します。
//
// Unmarshalは、Marshalが使用するエンコーディングの逆を使用し、
// 必要に応じてマップ、スライス、ポインタを割り当てます。
// 以下の追加ルールも適用されます：
>>>>>>> release-branch.go1.21
//
// JSONをポインタにアンマーシャルするために、Unmarshalはまず
// JSONがJSONリテラルnullであるケースを処理します。その場合、Unmarshalは
// ポインタをnilに設定します。それ以外の場合、UnmarshalはJSONを
// ポインタが指す値にアンマーシャルします。もしポインタがnilなら、Unmarshalは
// それが指す新しい値を割り当てます。
//
<<<<<<< HEAD
// To unmarshal JSON into a value implementing [Unmarshaler],
// Unmarshal calls that value's [Unmarshaler.UnmarshalJSON] method, including
// when the input is a JSON null.
// Otherwise, if the value implements [encoding.TextUnmarshaler]
// and the input is a JSON quoted string, Unmarshal calls
// [encoding.TextUnmarshaler.UnmarshalText] with the unquoted form of the string.
//
// To unmarshal JSON into a struct, Unmarshal matches incoming object
// keys to the keys used by [Marshal] (either the struct field name or its tag),
// preferring an exact match but also accepting a case-insensitive match. By
// default, object keys which don't have a corresponding struct field are
// ignored (see [Decoder.DisallowUnknownFields] for an alternative).
=======
// Unmarshalerインターフェースを実装する値にJSONをアンマーシャルするために、
// Unmarshalはその値のUnmarshalJSONメソッドを呼び出します、
// 入力がJSON nullである場合も含みます。
// それ以外の場合、もし値がencoding.TextUnmarshalerを実装していて、
// 入力がJSONの引用符で囲まれた文字列である場合、Unmarshalはその値の
// UnmarshalTextメソッドを引用符で囲まれていない形式の文字列で呼び出します。
//
// JSONを構造体にアンマーシャルするために、Unmarshalは受信したオブジェクトの
// キーをMarshalが使用するキー（構造体のフィールド名またはそのタグ）と一致させます。
// これは完全一致を優先しますが、大文字小文字を区別しない一致も受け入れます。
// デフォルトでは、対応する構造体のフィールドがないオブジェクトのキーは無視されます
// （代替としてDecoder.DisallowUnknownFieldsを参照してください）。
>>>>>>> release-branch.go1.21
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
<<<<<<< HEAD
// To unmarshal a JSON object into a map, Unmarshal first establishes a map to
// use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal
// reuses the existing map, keeping existing entries. Unmarshal then stores
// key-value pairs from the JSON object into the map. The map's key type must
// either be any string type, an integer, implement [json.Unmarshaler], or
// implement [encoding.TextUnmarshaler].
//
// If the JSON-encoded data contain a syntax error, Unmarshal returns a [SyntaxError].
//
// If a JSON value is not appropriate for a given target type,
// or if a JSON number overflows the target type, Unmarshal
// skips that field and completes the unmarshaling as best it can.
// If no more serious errors are encountered, Unmarshal returns
// an [UnmarshalTypeError] describing the earliest such error. In any
// case, it's not guaranteed that all the remaining fields following
// the problematic one will be unmarshaled into the target object.
=======
// JSONオブジェクトをマップにアンマーシャルするために、Unmarshalは最初に使用するマップを確立します。
// マップがnilの場合、Unmarshalは新しいマップを割り当てます。それ以外の場合、Unmarshalは
// 既存のマップを再利用し、既存のエントリを保持します。次に、UnmarshalはJSONオブジェクトから
// キーと値のペアをマップに格納します。マップのキーの型は、任意の文字列型、整数、
// json.Unmarshalerを実装するもの、またはencoding.TextUnmarshalerを実装するものでなければなりません。
//
// もしJSONエンコードされたデータに構文エラーが含まれている場合、UnmarshalはSyntaxErrorを返します。
//
// もしJSON値が特定のターゲット型に適していない場合、
// またはJSON数値がターゲット型をオーバーフローする場合、Unmarshalは
// そのフィールドをスキップし、可能な限りアンマーシャルを完了します。
// もしもっと深刻なエラーが発生しなければ、Unmarshalは最初のそのようなエラーを
// 説明するUnmarshalTypeErrorを返します。いずれにせよ、問題のあるフィールドに続く
// すべてのフィールドがターゲットオブジェクトにアンマーシャルされることは保証されません。
>>>>>>> release-branch.go1.21
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
// 入力は、JSON値の有効なエンコーディングであると想定できます。
// UnmarshalJSONは、戻り値後にデータを保持したい場合、JSONデータをコピーする必要があります。
//
<<<<<<< HEAD
// By convention, to approximate the behavior of [Unmarshal] itself,
// Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
=======
// 慣習的に、Unmarshal自体の振る舞いを近似するために、
// UnmarshalersはUnmarshalJSON([]byte("null"))を何もしない操作として実装します。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// An InvalidUnmarshalError describes an invalid argument passed to [Unmarshal].
// (The argument to [Unmarshal] must be a non-nil pointer.)
=======
// InvalidUnmarshalErrorは、Unmarshalに渡された無効な引数を説明します。
// (Unmarshalへの引数はnilでないポインタでなければなりません。)
>>>>>>> release-branch.go1.21
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
