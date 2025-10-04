// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// htmlパッケージは、HTMLテキストのエスケープとアンエスケープのための関数を提供します。
package html

// EscapeStringは"<"のような特殊文字を"&lt;"にエスケープします。
// エスケープするのは5つの文字のみです：<、>、&、'、"。
// [UnescapeString](EscapeString(s)) == sは常に成り立ちますが、逆は常に真ではありません。
func EscapeString(s string) string

// UnescapeStringは"&lt;"のようなエンティティを"<"にアンエスケープします。
// [EscapeString]がエスケープするよりも広範囲のエンティティをアンエスケープします。
// たとえば、"&aacute;"は"á"にアンエスケープされ、"&#225;"や"&#xE1;"も同様です。
// UnescapeString([EscapeString](s)) == sは常に成り立ちますが、逆は常に真ではありません。
func UnescapeString(s string) string
