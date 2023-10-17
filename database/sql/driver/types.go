// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

// ValueConverterはConvertValueメソッドを提供するインターフェースです。
//
// ドライバーパッケージによってさまざまなValueConverterの実装が提供され、
// ドライバー間の変換の一貫性を提供します。ValueConverterにはいくつかの用途があります：
//
//   - sqlパッケージによって提供されるValue型から、
//     データベーステーブルの特定の列型に変換し、
//     特定のint64がテーブルのuint16列に合うか確認するなどの作業を行います。
//
//   - データベースから取得された値をドライバーのValue型のいずれかに変換します。
//
//   - sqlパッケージによって、スキャン中にドライバーのValue型からユーザーの型に変換します。
type ValueConverter interface {
	// ConvertValue converts a value to a driver Value.
	ConvertValue(v any) (Value, error)
}

// ValuerはValueメソッドを提供するインターフェースです。
//
// Valuerインターフェースを実装する型は、自分自身をドライバの値に変換できます。
type Valuer interface {
	// Value returns a driver Value.
	// Value must not panic.
	Value() (Value, error)
}

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
var Bool boolType

var _ ValueConverter = boolType{}

// Int32は、入力値をint64に変換するValueConverterであり、int32の制限を尊重します。
var Int32 int32Type

var _ ValueConverter = int32Type{}

// Stringは、入力を文字列に変換するValueConverterです。
// 値が既に文字列または[]byteの場合は変更されません。
// 値が他の型の場合は、fmt.Sprintf("%v", v)で文字列に変換されます。
var String stringType

// Nullは、nilを許可することでValueConverterを実装するタイプですが、それ以外の場合は別のValueConverterに委任します。
type Null struct {
	Converter ValueConverter
}

func (n Null) ConvertValue(v any) (Value, error)

// NotNullは、nilを許可しないことでValueConverterを実装する型ですが、他のValueConverterに委譲します。
type NotNull struct {
	Converter ValueConverter
}

func (n NotNull) ConvertValue(v any) (Value, error)

// IsValue はvが有効なValueパラメータータイプかどうかを報告します。
func IsValue(v any) bool

// IsScanValueはIsValueと同等です。
// 互換性のために存在します。
func IsScanValue(v any) bool

// DefaultParameterConverterは、StmtがColumnConverterを実装していない場合に使用されるValueConverterのデフォルト実装です。
//
// DefaultParameterConverterは、引数がIsValue(arg)を満たす場合はその引数を直接返します。そうでない場合、引数がValuerを実装している場合はそのValueメソッドを使用してValueを返します。代替として、提供された引数の基底の型を使用してValueに変換します。
// 基底の整数型はint64に変換され、浮動小数点数はfloat64に変換され、bool型、string型、および[]byte型はそれ自体に変換されます。引数がnilポインタの場合、ConvertValueはnilのValueを返します。引数がnilではないポインタの場合、それは逆参照され、再帰的にConvertValueが呼び出されます。他の型はエラーです。
var DefaultParameterConverter defaultConverter

var _ ValueConverter = defaultConverter{}
