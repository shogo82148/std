// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

// Unwrapは、errの型にUnwrapメソッドが含まれている場合、
// errのUnwrapメソッドを呼び出した結果を返します。
// それ以外の場合、Unwrapはnilを返します。
//
// Unwrapは、"Unwrap() error"形式のメソッドのみを呼び出します。
// 特に、Unwrapは [Join] によって返されたエラーをアンラップしません。
func Unwrap(err error) error

// Isは、errのツリー内の任意のエラーがtargetに一致するかどうかを報告します。
//
<<<<<<< HEAD
// ツリーは、err自体に続いて、 [Unwrap] を繰り返し呼び出して得られたエラーで構成されています。
// errが複数のエラーをラップしている場合、Isは、errに続いてその子の深さ優先のトラバースを行います。
=======
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling its Unwrap() error or Unwrap() []error method. When err wraps multiple
// errors, Is examines err followed by a depth-first traversal of its children.
>>>>>>> upstream/master
//
// ターゲットに一致するエラーは、そのターゲットに等しい場合、または
// Is(error) boolというメソッドを実装している場合、Is(target)がtrueを返す場合です。
//
// エラータイプは、既存のエラーと同等に扱うためにIsメソッドを提供する場合があります。
// たとえば、MyErrorが次のように定義されている場合、
//
//	func (m MyError) Is(target error) bool { return target == fs.ErrExist }
//
// 例えば、Is(MyError{}, fs.ErrExist)はtrueを返します。
// 標準ライブラリの例については、 [syscall.Errno.Is] を参照してください。
// Isメソッドは、errとターゲットを浅く比較し、[Unwrap] を呼び出さないようにする必要があります。
func Is(err, target error) bool

// Asは、errのツリー内で最初にtargetに一致するエラーを検索し、
// 一致するエラーが見つかった場合、targetをそのエラー値に設定してtrueを返します。
// それ以外の場合、falseを返します。
//
<<<<<<< HEAD
// ツリーは、err自体に続いて、[Unwrap] を繰り返し呼び出して得られたエラーで構成されています。
// errが複数のエラーをラップしている場合、Asは、errに続いてその子の深さ優先のトラバースを行います。
=======
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling its Unwrap() error or Unwrap() []error method. When err wraps multiple
// errors, As examines err followed by a depth-first traversal of its children.
>>>>>>> upstream/master
//
// エラーがターゲットに一致する場合、エラーの具体的な値がtargetが指す値に代入可能であるか、
// またはエラーがAs(interface{}) boolというメソッドを持ち、As(target)がtrueを返す場合です。
// 後者の場合、Asメソッドはtargetを設定する責任があります。
//
// エラータイプは、異なるエラータイプであるかのように扱うことができるように、Asメソッドを提供する場合があります。
//
// Asは、targetがエラーを実装する型または任意のインターフェース型の、
// 非nilポインタでない場合にパニックを引き起こします。
func As(err error, target any) bool
