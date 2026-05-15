// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Marshalerは自分自身をマーシャルできる型が実装します。
// 実装が "jsontext" パッケージへの強い依存を避けたい場合を除き、[MarshalerTo] を実装することを推奨します。
//
// 実装では、呼び出し元が保持・変更しても安全なバッファを返すことを推奨します。
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// MarshalerToは自分自身をマーシャルできる型が実装します。
// より高いパフォーマンスと柔軟性のため、[Marshaler] ではなくMarshalerToを実装することが推奨されます。
// MarshalerとMarshalerToの両方を実装している場合は、MarshalerToが優先されます。
// その場合、両方の実装はデフォルトのマーシャルオプションで同等の動作を目指すべきです。
//
// 実装では、Encoderに1つのJSON値だけを書き込む必要があります。
// または、[errors.ErrUnsupported] を返してEncoderを変更せずにすることもできます。
// このメソッドを呼び出す "json" パッケージはレシーバー型の
// 次の利用可能なJSON表現を使用します。
// 実装は [jsontext.Encoder] へのポインタを保持してはいけません。
//
// 返されたエラーが [SemanticError] の場合、エラーの未設定フィールド
// は [json] によって追加コンテキストで設定されることがあります。
// 他の型のエラーは [SemanticError] でラップされます。
// ただし、IOエラーの場合は除きます。
//
// MarshalJSONToメソッドは直接呼び出すべきではありません。これは
// 特別な処理が必要なセンチネルエラーを返す場合があるためです。
// ユーザーは代わりに [MarshalEncode] を呼び出すべきです。これはそのような場合を処理します。
type MarshalerTo interface {
	MarshalJSONTo(*jsontext.Encoder) error
}

// Unmarshalerは自分自身をアンマーシャルできる型が実装します。
// 実装が "jsontext" パッケージへの強い依存を避けたい場合を除き、[UnmarshalerFrom] を実装することを推奨します。
//
// 入力はこのパッケージのアンマーシャル機能から呼び出された場合、正しいJSON値のエンコーディングであるとみなせます。
// UnmarshalJSONが返却後もJSONデータを保持する場合は、必ずコピーしてください。
// 事前に値が設定された変数にアンマーシャルする場合は、UnmarshalJSONでマージセマンティクスを実装することが推奨されます。
//
// 実装は入力の[]byteを保持したり変更したりしてはいけません。
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// UnmarshalerFromは自分自身をアンマーシャルできる型が実装します。
// より高いパフォーマンスと柔軟性のため、[Unmarshaler] ではなくUnmarshalerFromを実装することが推奨されます。
// UnmarshalerとUnmarshalerFromの両方を実装している場合は、UnmarshalerFromが優先されます。
// その場合、両方の実装はデフォルトのアンマーシャルオプションで同等の動作を目指すべきです。
//
// 実装はDecoderから1つのJSON値だけを読み込む必要があります。
// 事前に設定された値にアンマーシャルする場合は、UnmarshalJSONFromがマージセマンティクスを実装することが推奨されます。
// または、[errors.ErrUnsupported] を返してDecoderを変更せずにすることもできます。
// このメソッドを呼び出す "json" パッケージはレシーバー型の次の利用可能なJSON表現を使用します。
// 実装は [jsontext.Decoder] へのポインタを保持してはいけません。
//
// 返されたエラーが [SemanticError] の場合、エラーの未設定フィールド
// は [json] によって追加コンテキストで設定されることがあります。
// 他の型のエラーは [SemanticError] でラップされます。
// ただし、[jsontext.SyntacticError] またはIOエラーの場合は除きます。
//
// UnmarshalJSONFromメソッドは直接呼び出すべきではありません。これは
// 特別な処理が必要なセンチネルエラーを返す場合があるためです。
// ユーザーは代わりに [UnmarshalDecode] を呼び出すべきです。これはそのような場合を処理します。
type UnmarshalerFrom interface {
	UnmarshalJSONFrom(*jsontext.Decoder) error
}
