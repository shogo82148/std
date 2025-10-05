// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/errors"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// SkipFuncは、[MarshalToFunc] および [UnmarshalFromFunc] 関数から返される場合があります。
//
// SkipFuncを返す関数は、渡された [jsontext.Encoder] や [jsontext.Decoder] に
// 観測可能な副作用を与えてはなりません。
// 例えば、[jsontext.Decoder.PeekKind] を呼び出すことは許容されますが、
// [jsontext.Decoder.ReadToken] や [jsontext.Encoder.WriteToken] のような
// 状態を変更するメソッドを呼び出すことは許容されません。
var SkipFunc = errors.New("json: skip function")

// Marshalersは、特定の型のマーシャル動作を上書きできる関数群のリストです。
// [WithMarshalers] で設定することで、[Marshal]、[MarshalWrite]、[MarshalEncode] で利用できます。
// nilの*Marshalersは空リストと同等です。
// Marshalersにはエクスポートされたフィールドやメソッドはありません。
type Marshalers = typedMarshalers

// JoinMarshalersは、マーシャル関数のフラットなリストを構築します。
// リスト内の複数の関数が特定の型の値に適用可能な場合、リストの先頭にある関数ほど優先されます。
// 関数が [SkipFunc] を返した場合は、次に適用可能な関数が呼び出され、
// それ以外の場合はデフォルトのマーシャル動作が使われます。
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
// リスト内の複数の関数が特定の型の値に適用可能な場合、リストの先頭にある関数ほど優先されます。
// 関数が [SkipFunc] を返した場合は、次に適用可能な関数が呼び出され、
// それ以外の場合はデフォルトのアンマーシャル動作が使われます。
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
// 関数は必ず1つのJSON値をマーシャルしなければなりません。
// Tの値を関数呼び出しの外部に保持してはなりません。
// [SkipFunc] を返すことはできません。
func MarshalFunc[T any](fn func(T) ([]byte, error)) *Marshalers

// MarshalToFuncは、型T専用のマーシャル方法を指定する関数を構築します。
// Tは名前付きポインタ以外の任意の型にできます。
// Tがインターフェース型またはポインタ型の場合、関数には必ず非nilのポインタ値が渡されます。
//
// 関数は必ず提供されたエンコーダのwriteメソッドを呼び出して、1つのJSON値をマーシャルしなければなりません。
// [SkipFunc] を返すことで、次のマーシャル関数に処理を移すことができますが、[SkipFunc] を返す場合はエンコーダに対して可変メソッドを呼び出してはいけません。
// [jsontext.Encoder] へのポインタやTの値は関数呼び出しの外部に保持してはいけません。
func MarshalToFunc[T any](fn func(*jsontext.Encoder, T) error) *Marshalers

// UnmarshalFuncは、型T専用のアンマーシャル方法を指定する関数を構築します。
// Tは名前なしポインタ型またはインターフェース型でなければなりません。
// 関数には必ず非nilのポインタ値が渡されます。
//
// 関数は必ず1つのJSON値をアンマーシャルしなければなりません。
// 入力の[]byteは変更してはいけません。
// 入力の[]byteやTの値は関数呼び出しの外部に保持してはいけません。
// [SkipFunc] を返すことはできません。
func UnmarshalFunc[T any](fn func([]byte, T) error) *Unmarshalers

// UnmarshalFromFuncは、型T専用のアンマーシャル方法を指定する関数を構築します。
// Tは名前なしポインタ型またはインターフェース型でなければなりません。
// 関数には必ず非nilのポインタ値が渡されます。
//
// 関数は必ず提供されたデコーダのreadメソッドを呼び出して、1つのJSON値をアンマーシャルしなければなりません。
// [SkipFunc] を返すことで、次のアンマーシャル関数に処理を移すことができますが、[SkipFunc] を返す場合はデコーダに対して可変メソッドを呼び出してはいけません。
// [jsontext.Decoder] へのポインタやTの値は関数呼び出しの外部に保持してはいけません。
func UnmarshalFromFunc[T any](fn func(*jsontext.Decoder, T) error) *Unmarshalers
