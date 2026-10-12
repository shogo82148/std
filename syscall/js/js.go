// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

// jsパッケージは、js/wasmアーキテクチャを使用する場合にWebAssemblyのホスト環境へのアクセスを提供します。
// そのAPIはJavaScriptのセマンティクスに基づいています。
//
// このパッケージは実験的なものです。現在の範囲はテストを実行できるようにすることだけであり、
// ユーザー向けの包括的なAPIを提供するものではまだありません。Goの互換性の約束の対象外です。
package js

// Valueは、JavaScriptの値を表します。ゼロ値はJavaScriptの値"undefined"です。
// 値はEqualメソッドで等価性を確認できます。
type Value struct {
	_     [0]func()
	ref   ref
	gcPtr *ref
}

// Errorは、JavaScriptのエラーをラップします。
type Error struct {
	// Valueは、基礎となるJavaScriptのエラー値です。
	Value
}

// Errorはerrorインターフェースを実装します。
func (e Error) Error() string

// Equalは、JavaScriptの===演算子に従ってvとwが等しいかどうかを報告します。
func (v Value) Equal(w Value) bool

// Undefinedは、JavaScriptの値"undefined"を返します。
func Undefined() Value

// IsUndefinedは、vがJavaScriptの値"undefined"であるかどうかを報告します。
func (v Value) IsUndefined() bool

// Nullは、JavaScriptの値"null"を返します。
func Null() Value

// IsNullは、vがJavaScriptの値"null"であるかどうかを報告します。
func (v Value) IsNull() bool

// IsNaNは、vがJavaScriptの値"NaN"であるかどうかを報告します。
func (v Value) IsNaN() bool

// Globalは、JavaScriptのグローバルオブジェクト（通常は"window"または"global"）を返します。
func Global() Value

// ValueOfは、xをJavaScriptの値として返します:
//
//	| Go                     | JavaScript             |
//	| ---------------------- | ---------------------- |
//	| js.Value               | [its value]            |
//	| js.Func                | function               |
//	| nil                    | null                   |
//	| bool                   | boolean                |
//	| integers and floats    | number                 |
//	| string                 | string                 |
//	| []interface{}          | new array              |
//	| map[string]interface{} | new object             |
//
// xが期待される型のいずれでもない場合はパニックします。
func ValueOf(x any) Value

// Typeは、ValueのJavaScriptの型を表します。
type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string

// Typeは、値vのJavaScriptの型を返します。JavaScriptのtypeof演算子と似ていますが、
// nullに対してTypeObjectではなくTypeNullを返す点が異なります。
func (v Value) Type() Type

// Getは、値vのJavaScriptのプロパティーpを返します。
// vがJavaScriptのオブジェクトでない場合はパニックします。
func (v Value) Get(p string) Value

// Setは、値vのJavaScriptのプロパティーpをValueOf(x)に設定します。
// vがJavaScriptのオブジェクトでない場合はパニックします。
func (v Value) Set(p string, x any)

// Deleteは、値vのJavaScriptのプロパティーpを削除します。
// vがJavaScriptのオブジェクトでない場合はパニックします。
func (v Value) Delete(p string)

// Indexは、値vのJavaScriptのインデックスiを返します。
// vがJavaScriptのオブジェクトでない場合はパニックします。
func (v Value) Index(i int) Value

// SetIndexは、値vのJavaScriptのインデックスiをValueOf(x)に設定します。
// vがJavaScriptのオブジェクトでない場合はパニックします。
func (v Value) SetIndex(i int, x any)

// Lengthは、vのJavaScriptのプロパティー"length"を返します。
// vがJavaScriptのオブジェクトでない場合はパニックします。
func (v Value) Length() int

// Callは、指定された引数で値vのメソッドmをJavaScriptで呼び出します。
// vがメソッドmを持たない場合はパニックします。
// 引数は、ValueOf関数に従ってJavaScriptの値にマッピングされます。
func (v Value) Call(m string, args ...any) Value

// Invokeは、指定された引数で値vをJavaScriptで呼び出します。
// vがJavaScriptの関数でない場合はパニックします。
// 引数は、ValueOf関数に従ってJavaScriptの値にマッピングされます。
func (v Value) Invoke(args ...any) Value

// Newは、値vをコンストラクタとし、指定された引数でJavaScriptの"new"演算子を使用します。
// vがJavaScriptの関数でない場合はパニックします。
// 引数は、ValueOf関数に従ってJavaScriptの値にマッピングされます。
func (v Value) New(args ...any) Value

// Floatは、値vをfloat64として返します。
// vがJavaScriptの数値でない場合はパニックします。
func (v Value) Float() float64

// Intは、値vをintに切り捨てて返します。
// vがJavaScriptの数値でない場合はパニックします。
func (v Value) Int() int

// Boolは、値vをboolとして返します。
// vがJavaScriptの真偽値でない場合はパニックします。
func (v Value) Bool() bool

// Truthyは、値vのJavaScriptの「truthiness」を返します。JavaScriptでは、
// false、0、""、null、undefined、NaNは「falsy」であり、それ以外はすべて
// 「truthy」です。https://developer.mozilla.org/en-US/docs/Glossary/Truthy を参照してください。
func (v Value) Truthy() bool

// Stringは、値vを文字列として返します。
// GoのStringメソッドの慣例のため、Stringは特別なケースです。他のゲッターとは異なり、
// vのTypeがTypeStringでない場合でもパニックしません。代わりに、"<T>"または
// "<T: V>"の形式の文字列を返します。ここでTはvの型、Vはvの値の文字列表現です。
func (v Value) String() string

// InstanceOfは、JavaScriptのinstanceof演算子に従ってvが型tのインスタンスであるかどうかを報告します。
func (v Value) InstanceOf(t Value) bool

// ValueErrorは、ValueメソッドがそれをサポートしていないValueに対して
// 呼び出された場合に発生します。そのようなケースは各メソッドの説明に記載されています。
type ValueError struct {
	Method string
	Type   Type
}

func (e *ValueError) Error() string

// CopyBytesToGoは、srcからdstにバイトをコピーします。
// srcがUint8ArrayまたはUint8ClampedArrayでない場合はパニックします。
// コピーされたバイト数を返します。これはsrcとdstの長さの最小値になります。
func CopyBytesToGo(dst []byte, src Value) int

// CopyBytesToJSは、srcからdstにバイトをコピーします。
// dstがUint8ArrayまたはUint8ClampedArrayでない場合はパニックします。
// コピーされたバイト数を返します。これはsrcとdstの長さの最小値になります。
func CopyBytesToJS(dst Value, src []byte) int
