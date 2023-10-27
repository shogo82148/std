// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

// ValueConverterはConvertValueメソッドを提供するインターフェースです。
//
// ドライバーパッケージによってさまざまなValueConverterの実装が提供され、
// ドライバー間の変換の一貫性を提供します。ValueConverterにはいくつかの用途があります：
//
//   - sqlパッケージによって提供される [Value] 型から、
//     データベーステーブルの特定の列型に変換し、
//     特定のint64がテーブルのuint16列に合うか確認するなどの作業を行います。
//
//   - データベースから取得された値をドライバーの [Value] 型のいずれかに変換します。
//
<<<<<<< HEAD
//   - sqlパッケージによって、スキャン中にドライバーの [Value] 型からユーザーの型に変換します。
=======
//   - by the [database/sql] package, for converting from a driver's [Value] type
//     to a user's type in a scan.
>>>>>>> upstream/master
type ValueConverter interface {
	ConvertValue(v any) (Value, error)
}

// ValuerはValueメソッドを提供するインターフェースです。
//
<<<<<<< HEAD
// Valuerインターフェースを実装する型は、自分自身をドライバの値に変換できます。
=======
// Types implementing Valuer interface are able to convert
// themselves to a driver [Value].
>>>>>>> upstream/master
type Valuer interface {
	Value() (Value, error)
}

// Boolは入力値をboolに変換する [ValueConverter] です。
//
// 変換ルールは以下の通りです：
//   - ブール値は変更されずに返されます
//   - 整数の場合、
//     1はtrueを、
//     0はfalseを、
//     その他の整数はエラーとなります
//   - 文字列や[]byteの場合、 [strconv.ParseBool] と同じルールが適用されます
//   - それ以外のすべての型はエラーとなります
var Bool boolType

var _ ValueConverter = boolType{}

// Int32は、入力値をint64に変換する [ValueConverter] であり、int32の制限を尊重します。
var Int32 int32Type

var _ ValueConverter = int32Type{}

// Stringは、入力を文字列に変換する [ValueConverter] です。
// 値が既に文字列または[]byteの場合は変更されません。
// 値が他の型の場合は、fmt.Sprintf("%v", v)で文字列に変換されます。
var String stringType

// Nullは、nilを許可することでValueConverterを実装するタイプですが、それ以外の場合は別のValueConverterに委任します。
type Null struct {
	Converter ValueConverter
}

func (n Null) ConvertValue(v any) (Value, error)

// NotNullは、nilを許可しないことで [ValueConverter] を実装する型ですが、他の [ValueConverter] に委譲します。
type NotNull struct {
	Converter ValueConverter
}

func (n NotNull) ConvertValue(v any) (Value, error)

// IsValue はvが有効な [Value] パラメータータイプかどうかを報告します。
func IsValue(v any) bool

<<<<<<< HEAD
// IsScanValueはIsValueと同等です。
// 互換性のために存在します。
=======
// IsScanValue is equivalent to [IsValue].
// It exists for compatibility.
>>>>>>> upstream/master
func IsScanValue(v any) bool

// DefaultParameterConverterは、 [Stmt] が [ColumnConverter] を実装していない場合に使用される [ValueConverter] のデフォルト実装です。
//
<<<<<<< HEAD
// DefaultParameterConverterは、引数がIsValue(arg)を満たす場合はその引数を直接返します。そうでない場合、引数が [Valuer] を実装している場合はそのValueメソッドを使用して [Value] を返します。代替として、提供された引数の基底の型を使用して [Value] に変換します。
// 基底の整数型はint64に変換され、浮動小数点数はfloat64に変換され、bool型、string型、および[]byte型はそれ自体に変換されます。引数がnilポインタの場合、 [defaultConverter.ConvertValue] はnilのValueを返します。引数がnilではないポインタの場合、それは逆参照され、再帰的に [defaultConverter.ConvertValue] が呼び出されます。他の型はエラーです。
=======
// DefaultParameterConverter returns its argument directly if
// IsValue(arg). Otherwise, if the argument implements [Valuer], its
// Value method is used to return a [Value]. As a fallback, the provided
// argument's underlying type is used to convert it to a [Value]:
// underlying integer types are converted to int64, floats to float64,
// bool, string, and []byte to themselves. If the argument is a nil
// pointer, defaultConverter.ConvertValue returns a nil [Value].
// If the argument is a non-nil pointer, it is dereferenced and
// defaultConverter.ConvertValue is called recursively. Other types
// are an error.
>>>>>>> upstream/master
var DefaultParameterConverter defaultConverter

var _ ValueConverter = defaultConverter{}
