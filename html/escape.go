// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// htmlパッケージは、HTMLテキストのエスケープとアンエスケープのための関数を提供します。
package html

// EscapeStringは、"<"のような特殊文字を"&lt;"にエスケープします。
// それは、<, >, &, ' および " の5つの文字だけをエスケープします。
// 常にUnescapeString(EscapeString(s)) == sが成り立ちますが、その逆は
// 常に成り立つわけではありません。
func EscapeString(s string) string

// UnescapeStringは、"&lt;"のようなエンティティを"<"にアンエスケープします。
// それはEscapeStringがエスケープするよりも広範なエンティティをアンエスケープします。
// 例えば、"&aacute;"は"á"に、"&#225;"も"&#xE1;"も"á"にアンエスケープします。
// 常にUnescapeString(EscapeString(s)) == sが成り立ちますが、その逆は
// 常に成り立つわけではありません。
func UnescapeString(s string) string
