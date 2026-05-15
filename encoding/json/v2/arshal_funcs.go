// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Marshalersは、特定の型のマーシャル動作を上書きできる関数群のリストです。
// [WithMarshalers] で設定することで、[Marshal]、[MarshalWrite]、[MarshalEncode] で利用できます。
// nilの*Marshalersは空リストと同等です。
// Marshalersにはエクスポートされたフィールドやメソッドはありません。
type Marshalers = typedMarshalers

// JoinMarshalersは、マーシャル関数のフラットなリストを構築します。
// リスト内に同じ型に適用できる関数が複数ある場合、
// リストの前にある関数が後にある関数より優先されます。
// 関数が [errors.ErrUnsupported] を返した場合、
// 次に適用可能な関数が呼ばれます。
// それ以外の場合はデフォルトのマーシャル動作が使用されます。
//
// 例：
//
//	m1 := JoinMarshalers(f1, f2)
//	m2 := JoinMarshalers(f0, m1, f3)     // m3と同等
//	m3 := JoinMarshalers(f0, f1, f2, f3) // m2と同等
func JoinMarshalers(ms ...*Marshalers) *Marshalers

// Unmarshalersは、特定の型のアンマーシャル動作を上書きできる関数群のリストです。
// [WithUnmarshalers] で設定することで、[Unmarshal]、[UnmarshalRead]、[UnmarshalDecode] で利用できます。
// nilの*Unmarshalersは空リストと同等です。
// Unmarshalersにはエクスポートされたフィールドやメソッドはありません。
type Unmarshalers = typedUnmarshalers

// JoinUnmarshalersは、アンマーシャル関数のフラットなリストを構築します。
// リスト内に同じ型に適用できる関数が複数ある場合、
// リストの前にある関数が後にある関数より優先されます。
// 関数が [errors.ErrUnsupported] を返した場合、
// 次に適用可能な関数が呼ばれます。
// それ以外の場合はデフォルトのアンマーシャル動作が使用されます。
//
// 例：
//
//	u1 := JoinUnmarshalers(f1, f2)
//	u2 := JoinUnmarshalers(f0, u1, f3)     // u3と同等
//	u3 := JoinUnmarshalers(f0, f1, f2, f3) // u2と同等
func JoinUnmarshalers(us ...*Unmarshalers) *Unmarshalers

// MarshalFuncは、型T専用のマーシャル方法を指定する関数を構築します。
// Tは名前付きポインタ以外の任意の型にできます。
// Tがインターフェース型またはポインタ型の場合、関数には必ず非nilのポインタ値が渡されます。
//
// 関数はJSON値をちょうど1つマーシャルしなければなりません。
// Tの値は関数呼び出しの外部に保持してはなりません。
// [errors.ErrUnsupported] を返してはなりません。
func MarshalFunc[T any](fn func(T) ([]byte, error)) *Marshalers

// MarshalToFuncは、型T専用のマーシャル方法を指定する関数を構築します。
// Tは名前付きポインタ以外の任意の型にできます。
// Tがインターフェース型またはポインタ型の場合、関数には必ず非nilのポインタ値が渡されます。
//
// 関数は提供されたエンコーダーの書き込みメソッドを呼び出して、
// JSON値をちょうど1つマーシャルしなければなりません。
// [errors.ErrUnsupported] を返すことで、次のマーシャル関数に処理を移すことができます。
// ただし、[errors.ErrUnsupported] を返す場合、エンコーダーの変更を伴うメソッドを
// 呼び出してはなりません。
// [jsontext.Encoder] へのポインタとTの値は関数呼び出しの外部に保持してはなりません。
func MarshalToFunc[T any](fn func(*jsontext.Encoder, T) error) *Marshalers

// UnmarshalFuncは、型T専用のアンマーシャル方法を指定する関数を構築します。
// Tは名前なしポインタ型またはインターフェース型でなければなりません。
// 関数には必ず非nilのポインタ値が渡されます。
//
// 関数はJSON値をちょうど1つアンマーシャルしなければなりません。
// 入力の[]byteを変更してはなりません。
// 入力の[]byteとTの値は関数呼び出しの外部に保持してはなりません。
// [errors.ErrUnsupported] を返してはなりません。
func UnmarshalFunc[T any](fn func([]byte, T) error) *Unmarshalers

// UnmarshalFromFuncは、型T専用のアンマーシャル方法を指定する関数を構築します。
// Tは名前なしポインタ型またはインターフェース型でなければなりません。
// 関数には必ず非nilのポインタ値が渡されます。
//
// 関数は提供されたデコーダーの読み取りメソッドを呼び出して、
// JSON値をちょうど1つアンマーシャルしなければなりません。
// [errors.ErrUnsupported] を返すことで、次のアンマーシャル関数に処理を移すことができます。
// ただし、[errors.ErrUnsupported] を返す場合、デコーダーの変更を伴うメソッドを
// 呼び出してはなりません。
// [jsontext.Decoder] へのポインタとTの値は関数呼び出しの外部に保持してはなりません。
func UnmarshalFromFunc[T any](fn func(*jsontext.Decoder, T) error) *Unmarshalers
