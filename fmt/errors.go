// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

// Errorfは書式指定子に従ってフォーマットを行い、errorを満たす値として文字列を返します。
//
// 書式指定子にerrorオペランドを持つ%w動詞が含まれている場合、
// 返されたエラーはそのオペランドを返すUnwrapメソッドを実装します。
// %w動詞が複数ある場合、返されたエラーは引数に現れる順序で
// すべての%wオペランドを含む[]errorを返すUnwrapメソッドを実装します。
// errorインターフェースを実装していないオペランドを%w動詞に指定するのは無効です。
// %w動詞は、それ以外の場合は%vの同義語です。
func Errorf(format string, a ...any) (err error)
