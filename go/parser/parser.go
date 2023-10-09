// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージparserはGoのソースファイルのパーサを実装しています。入力はさまざまな形式で提供できます（Parse*関数を参照）。出力はGoソースを表す抽象構文木（AST）です。パーサはParse*関数のいずれかを経由して呼び出されます。
// パーサは、Goの仕様で構文的に許可されていないより大きな言語を受け入れていますが、これは単純さと構文エラーの存在下での強靱性の向上のためです。たとえば、メソッド宣言では、レシーバは通常のパラメータリストのように扱われるため、複数のエントリが可能ですが、仕様では一つしか許可されていません。そのため、ASTの対応するフィールド（ast.FuncDecl.Recv）は一つに制限されていません。
package parser
