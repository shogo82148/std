// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// htmlパッケージは、HTMLテキストのエスケープとアンエスケープのための関数を提供します。
package html

<<<<<<< HEAD
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
=======
// EscapeString escapes special characters like "<" to become "&lt;". It
// escapes only five such characters: <, >, &, ' and ".
// [UnescapeString](EscapeString(s)) == s always holds, but the converse isn't
// always true.
func EscapeString(s string) string

// UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a
// larger range of entities than [EscapeString] escapes. For example, "&aacute;"
// unescapes to "á", as does "&#225;" and "&#xE1;".
// UnescapeString([EscapeString](s)) == s always holds, but the converse isn't
// always true.
>>>>>>> upstream/release-branch.go1.25
func UnescapeString(s string) string
