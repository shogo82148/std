// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

// Joinは、指定されたエラーをラップするエラーを返します。
// nilエラー値は破棄されます。
// errsのすべての値がnilの場合、Joinはnilを返します。
// エラーは、errsの各要素のErrorメソッドを呼び出して得られた文字列を連結したもので、
// 各文字列の間に改行が挿入されたものとしてフォーマットされます。
//
// Joinによって返されるnilでないエラーは、Unwrap() []errorメソッドを実装します。
func Join(errs ...error) error
