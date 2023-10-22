// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

// ValueConverterはConvertValueメソッドを提供するインターフェースです。
//
// ドライバーパッケージによってさまざまなValueConverterの実装が提供され、
// ドライバー間の変換の一貫性を提供します。ValueConverterにはいくつかの用途があります：
//
<<<<<<< HEAD
//   - sqlパッケージによって提供されるValue型から、
//     データベーステーブルの特定の列型に変換し、
//     特定のint64がテーブルのuint16列に合うか確認するなどの作業を行います。
//
//   - データベースから取得された値をドライバーのValue型のいずれかに変換します。
//
//   - sqlパッケージによって、スキャン中にドライバーのValue型からユーザーの型に変換します。
=======
//   - converting from the [Value] types as provided by the sql package
//     into a database table's specific column type and making sure it
//     fits, such as making sure a particular int64 fits in a
//     table's uint16 column.
//
//   - converting a value as given from the database into one of the
//     driver [Value] types.
//
//   - by the sql package, for converting from a driver's [Value] type
//     to a user's type in a scan.
>>>>>>> upstream/master
type ValueConverter interface {
	ConvertValue(v any) (Value, error)
}

// ValuerはValueメソッドを提供するインターフェースです。
//
// Valuerインターフェースを実装する型は、自分自身をドライバの値に変換できます。
type Valuer interface {
	Value() (Value, error)
}

<<<<<<< HEAD
// Boolは入力値をboolに変換するValueConverterです。
//
// 変換ルールは以下の通りです：
//   - ブール値は変更されずに返されます
//   - 整数の場合、
//     1はtrueを、
//     0はfalseを、
//     その他の整数はエラーとなります
//   - 文字列や[]byteの場合、strconv.ParseBoolと同じルールが適用されます
//   - それ以外のすべての型はエラーとなります
=======
// Bool is a [ValueConverter] that converts input values to bool.
//
// The conversion rules are:
//   - booleans are returned unchanged
//   - for integer types,
//     1 is true
//     0 is false,
//     other integers are an error
//   - for strings and []byte, same rules as [strconv.ParseBool]
//   - all other types are an error
>>>>>>> upstream/master
var Bool boolType

var _ ValueConverter = boolType{}

<<<<<<< HEAD
// Int32は、入力値をint64に変換するValueConverterであり、int32の制限を尊重します。
=======
// Int32 is a [ValueConverter] that converts input values to int64,
// respecting the limits of an int32 value.
>>>>>>> upstream/master
var Int32 int32Type

var _ ValueConverter = int32Type{}

<<<<<<< HEAD
// Stringは、入力を文字列に変換するValueConverterです。
// 値が既に文字列または[]byteの場合は変更されません。
// 値が他の型の場合は、fmt.Sprintf("%v", v)で文字列に変換されます。
var String stringType

// Nullは、nilを許可することでValueConverterを実装するタイプですが、それ以外の場合は別のValueConverterに委任します。
=======
// String is a [ValueConverter] that converts its input to a string.
// If the value is already a string or []byte, it's unchanged.
// If the value is of another type, conversion to string is done
// with fmt.Sprintf("%v", v).
var String stringType

// Null is a type that implements [ValueConverter] by allowing nil
// values but otherwise delegating to another [ValueConverter].
>>>>>>> upstream/master
type Null struct {
	Converter ValueConverter
}

func (n Null) ConvertValue(v any) (Value, error)

<<<<<<< HEAD
// NotNullは、nilを許可しないことでValueConverterを実装する型ですが、他のValueConverterに委譲します。
=======
// NotNull is a type that implements [ValueConverter] by disallowing nil
// values but otherwise delegating to another [ValueConverter].
>>>>>>> upstream/master
type NotNull struct {
	Converter ValueConverter
}

func (n NotNull) ConvertValue(v any) (Value, error)

<<<<<<< HEAD
// IsValue はvが有効なValueパラメータータイプかどうかを報告します。
=======
// IsValue reports whether v is a valid [Value] parameter type.
>>>>>>> upstream/master
func IsValue(v any) bool

// IsScanValueはIsValueと同等です。
// 互換性のために存在します。
func IsScanValue(v any) bool

<<<<<<< HEAD
// DefaultParameterConverterは、StmtがColumnConverterを実装していない場合に使用されるValueConverterのデフォルト実装です。
//
// DefaultParameterConverterは、引数がIsValue(arg)を満たす場合はその引数を直接返します。そうでない場合、引数がValuerを実装している場合はそのValueメソッドを使用してValueを返します。代替として、提供された引数の基底の型を使用してValueに変換します。
// 基底の整数型はint64に変換され、浮動小数点数はfloat64に変換され、bool型、string型、および[]byte型はそれ自体に変換されます。引数がnilポインタの場合、ConvertValueはnilのValueを返します。引数がnilではないポインタの場合、それは逆参照され、再帰的にConvertValueが呼び出されます。他の型はエラーです。
=======
// DefaultParameterConverter is the default implementation of
// [ValueConverter] that's used when a [Stmt] doesn't implement
// [ColumnConverter].
//
// DefaultParameterConverter returns its argument directly if
// IsValue(arg). Otherwise, if the argument implements [Valuer], its
// Value method is used to return a [Value]. As a fallback, the provided
// argument's underlying type is used to convert it to a [Value]:
// underlying integer types are converted to int64, floats to float64,
// bool, string, and []byte to themselves. If the argument is a nil
// pointer, [defaultConverter.ConvertValue] returns a nil [Value].
// If the argument is a non-nil pointer, it is dereferenced and
// [defaultConverter.ConvertValue] is called recursively. Other types
// are an error.
>>>>>>> upstream/master
var DefaultParameterConverter defaultConverter

var _ ValueConverter = defaultConverter{}
