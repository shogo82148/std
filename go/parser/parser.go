// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// parserパッケージは、Goソースファイル用のパーサーを実装します。
//
// [ParseFile] 関数は、string、[]byte、またはio.Readerからファイル入力を読み取り、
// ファイルの完全な抽象構文木を表す [ast.File] を生成します。
//
// [ParseExprFrom] 関数は、単一のソースレベル式を読み取り、
// その式の構文木である [ast.Expr] を生成します。
//
// パーサーは、簡単性と構文エラーの存在下での堅牢性の向上のために、
// Go仕様で構文的に許可されているよりも大きな言語を受け入れます。
// 例えば、メソッド宣言では、レシーバーは通常のパラメータリストのように扱われるため、
// 仕様では正確に1つが許可されている場所で複数のエントリを含むことができます。
// その結果、ASTの対応するフィールド（ast.FuncDecl.Recv）フィールドは1つのエントリに制限されません。
//
// Goソースコードの1つ以上の完全なパッケージを解析する必要があるアプリケーションは、
// パーサーと直接やり取りするのではなく、代わりに
// [golang.org/x/tools/go/packages] パッケージのLoad関数を使用する方が便利かもしれません。
package parser
