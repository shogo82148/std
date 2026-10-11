// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// Simplifyは、reと同等で、回数指定の繰り返しを含まず、/(?:a+)+/を/a+/に
// 書き換えるなどのさまざまな簡略化を施した正規表現を返します。
// 結果の正規表現は正しく実行されますが、キャプチャの括弧が複製または削除されている
// 可能性があるため、その文字列表現は同じパースツリーを生成しません。
// 例えば、/(x){1,2}/の簡略化形式は/(x)(x)?/ですが、両方の括弧が$1としてキャプチャします。
// 返される正規表現は元の正規表現と構造を共有するか、元の正規表現そのものである場合があります。
func (re *Regexp) Simplify() *Regexp
