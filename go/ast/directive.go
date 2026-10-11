// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// Directiveは次の形式のコメントです:
//
//	//tool:name args
//
// 例えば、次のディレクティブ:
//
//	//go:generate stringer -type Op -trimprefix Op
//
// はTool "go"、Name "generate"、Args "stringer -type Op -trimprefix Op"を持ちます。
//
// Argsには厳密な構文はありませんが、慣例として、引用符で囲まれていない単語、
// '"'で囲まれたGo文字列、または'`'で囲まれた生文字列をスペースで区切った並びです。
//
// 仕様については https://go.dev/doc/comment#directives を参照してください。
type Directive struct {
	Tool string
	Name string
	Args string

	// Slashは、ディレクティブの先頭にある"//"の位置です。
	Slash token.Pos

	// ArgsPosは、ParseDirectiveに渡された位置に基づく、Argsが始まる位置です。
	ArgsPos token.Pos
}

// ParseDirectiveは、ディレクティブコメントの単一のコメント行をパースします。
//
// 行がディレクティブコメントでない場合、falseを返します。
//
// 指定されるテキストは1行でなければならず、先頭の"//"を含む必要があります。
// テキストが"//"で始まらない場合、falseを返します。
//
// 呼び出し元はcの先頭のファイル位置を指定できます。これは引数の位置を
// 追跡するために使用されます。これは [Comment.Slash]、呼び出し元が合成した値、
// または単に0にすることができます。呼び出し元が0を渡した場合、位置は実質的に
// 文字列cへのバイトオフセットになります。
func ParseDirective(pos token.Pos, c string) (Directive, bool)

func (d *Directive) Pos() token.Pos
func (d *Directive) End() token.Pos

// DirectiveArgは、ディレクティブコメントの引数です。
type DirectiveArg struct {
	// Argはパースされた引数文字列です。引数が引用符で囲まれた文字列だった場合、
	// これは引用符を外した形式です。
	Arg string
	// Posは、この引数の最初の文字の位置です。
	Pos token.Pos
}

// ParseArgsは、標準の慣例を使用して [Directive] の引数をパースします。
// 標準の慣例とはトークンの並びであり、各トークンは裸の単語、
// ダブルクォートで囲まれたGo文字列、またはバッククォートで囲まれたGoの生文字列です。
// 各トークンは1つ以上のUnicodeスペースで区切られていなければなりません。
//
// 引数がこの構文に従っていない場合、エラーを返します。
func (d *Directive) ParseArgs() ([]DirectiveArg, error)
