// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Marshalerは、自分自身をマーシャルできる型が実装します。
// 実装が "jsontext" パッケージへの直接依存を避けようとしていない限り、
// 型は [MarshalerTo] を実装することが推奨されます。
//
// 実装は、呼び出し元が保持して必要に応じて変更できるような
// バッファを返すべきです。
//
// 実装は [errors.ErrUnsupported] を返してはいけません。
//
// 返されたエラーが [SemanticError] の場合、未設定のフィールドには
// [json] によって追加のコンテキストが設定されることがあります。
// それ以外の型のエラーは [SemanticError] でラップされます。
//
// 実装は [Deterministic] が true であると仮定して、
// 決定的な出力を返すべきです。
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// MarshalerToは、自分自身をマーシャルできる型が実装します。
// [Marshaler] の代わりに MarshalerTo を実装することが推奨されます。
// これはより高いパフォーマンスと柔軟性を提供するためです。
// 型が Marshaler と MarshalerTo の両方を実装している場合、
// MarshalerTo が優先されます。その場合、両方の実装は
// デフォルトのマーシャルオプションに対して同等の動作を目指すべきです。
//
// 実装は Encoder に JSON 値を 1 つだけ書き込まなければなりません。
// あるいは、Encoder を変更せずに [errors.ErrUnsupported] を返しても構いません。
// そのメソッドを呼び出す "json" パッケージは、
// レシーバー型の次に利用可能なJSON表現を使います。
// これは [Marshal] で説明されている通りです。
// 実装は [jsontext.Encoder] へのポインタを保持してはいけません。
//
// 返されたエラーが [SemanticError] の場合、未設定のフィールドには
// [json] によって追加のコンテキストが設定されることがあります。
// 他の型のエラーは [SemanticError] でラップされ、
// IO エラーを除きます。
//
// MarshalJSONTo メソッドは直接呼び出してはいけません。
// 特殊な処理が必要なセンチネルエラーを返す場合があるためです。
// ユーザーはそのような場合を処理する [MarshalEncode] を代わりに呼び出すべきです。
//
// 実装は [jsontext.Encoder.Options] からマーシャルオプションを確認し、
// 必要に応じてそのオプションに従って動作を調整すべきです。
//
// 次のオプションは MarshalerTo 実装に関連している場合があります:
//
// - [Deterministic]: 実装が非決定的な出力を生成する可能性がある場合
// - [StringifyNumbers]: 型が JSON 数値として表現される場合
//
// [FormatNilSliceAsNull] のようないくつかのオプションは、
// ネイティブな Go の型にのみ適用されます。そのため、
// これらのオプションは通常 MarshalerTo 実装には直接関係しません。
// ただし、複合型を表す型は、含まれる型に対して [MarshalEncode] を使って
// マーシャルし、これらのオプションがその型に適用されるようにすべきです。
// 同様に、[WithMarshalers] は複合型内の含まれる型のマーシャルにも影響し得ます。
//
// それ以外のオプションはすべて MarshalerTo 実装の外側で自動的に処理されるため、
// 実装には関係しません。
type MarshalerTo interface {
	MarshalJSONTo(*jsontext.Encoder) error
}

// Unmarshalerは、自分自身をアンマーシャルできる型が実装します。
// 実装が "jsontext" パッケージへの直接依存を避けようとしていない限り、
// [UnmarshalerFrom] を実装することが推奨されます。
//
// このパッケージのアンマーシャル機能から呼び出された場合、
// 入力は JSON 値の有効なエンコーディングであるとみなしてよいです。
// 事前に値が設定された変数にアンマーシャルする場合は、
// [Unmarshal] で説明されているように UnmarshalJSON がマージセマンティクスを
// 実装することが推奨されます。
//
// 実装は入力の []byte を保持したり変更したりしてはいけません。
//
// 実装は [errors.ErrUnsupported] を返してはいけません。
//
// 返されたエラーが [SemanticError] の場合、未設定のフィールドには
// [json] によって追加のコンテキストが設定されることがあります。
// 他の型のエラーは [SemanticError] でラップされます。
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// UnmarshalerFromは、自分自身をアンマーシャルできる型が実装します。
// 実装が "jsontext" パッケージへの直接依存を避けようとしていない限り、
// [Unmarshaler] の代わりに UnmarshalerFrom を実装することが推奨されます。
// これはより高いパフォーマンスと柔軟性を提供するためです。
// 型が Unmarshaler と UnmarshalerFrom の両方を実装している場合、
// UnmarshalerFrom が優先されます。その場合、両方の実装は
// デフォルトのアンマーシャルオプションに対して同等の動作を目指すべきです。
//
// 実装は Decoder から JSON 値を 1 つだけ読み込まなければなりません。
// 事前に値が設定された変数にアンマーシャルする場合は、
// [Unmarshal] で説明されているように UnmarshalJSONFrom がマージセマンティクスを
// 実装することが推奨されます。
// あるいは、Decoder を変更せずに [errors.ErrUnsupported] を返しても構いません。
// そのメソッドを呼び出す "json" パッケージは、
// レシーバー型の次に利用可能なJSON表現を使います。
// 実装は [jsontext.Decoder] へのポインタを保持してはいけません。
//
// 返されたエラーが [SemanticError] の場合、未設定のフィールドには
// [json] によって追加のコンテキストが設定されることがあります。
// 他の型のエラーは [SemanticError] でラップされ、
// [jsontext.SyntacticError] と IO エラーを除きます。
//
// UnmarshalJSONFrom メソッドは直接呼び出してはいけません。
// 特殊な処理が必要なセンチネルエラーを返す場合があるためです。
// ユーザーはそのような場合を処理する [UnmarshalDecode] を代わりに呼び出すべきです。
//
// 実装は [jsontext.Decoder.Options] からアンマーシャルオプションを確認し、
// 必要に応じてそのオプションに従って動作を調整すべきです。
//
// 次のオプションは UnmarshalerFrom 実装に関連している場合があります:
//
// - [StringifyNumbers]: 型が JSON 数値として表現される場合
//
// [FormatNilSliceAsNull] のようないくつかのオプションは、
// ネイティブな Go の型にのみ適用されます。そのため、
// これらのオプションは通常 UnmarshalerFrom 実装には直接関係しません。
// ただし、複合型を表す型は、含まれる型に対して [UnmarshalDecode] を使って
// アンマーシャルし、これらのオプションがその型に適用されるようにすべきです。
// 同様に、[WithUnmarshalers] は複合型内の含まれる型のアンマーシャルにも影響し得ます。
//
// それ以外のオプションはすべて UnmarshalerFrom 実装の外側で自動的に処理されるため、
// 実装には関係しません。
type UnmarshalerFrom interface {
	UnmarshalJSONFrom(*jsontext.Decoder) error
}
