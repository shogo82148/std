// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

// Errorfは書式指定子に従ってフォーマットを行い、errorを満たす値として文字列を返します。
//
// 書式指定子に%w動詞とエラー被演算子を含む場合、返されるエラーはUnwrapメソッドを実装し、被演算子を返します。
// 複数の%w動詞がある場合、返されるエラーは、引数の順序で表示されるすべての%w被演算子を含む[]errorを返します。
// エライターフェースを実装していない被演算子を%w動詞として使用することは無効です。それ以外の場合、%w動詞は%vと同義です。
func Errorf(format string, a ...any) error
