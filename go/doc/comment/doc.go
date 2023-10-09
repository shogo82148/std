// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージコメントは、Goのドキュメントコメント（ドキュメンテーションコメント）を解析および再フォーマットするためのものです。
パッケージ、const、func、type、またはvarのトップレベルの宣言の直前にあるコメントを指します。

Goのドキュメントコメントの構文は、リンク、見出し、段落、リスト（ネストなし）、および整形済みのテキストブロックをサポートする、
Markdownの簡略化されたサブセットです。構文の詳細は、https://go.dev/doc/commentで文書化されています。

（コメントマーカーを削除した後の）ドキュメントコメントに関連付けられたテキストを解析するには、[Parser]を使用します：

	var p comment.Parser
	doc := p.Parse(text)

結果は、[*Doc]です。
ドキュメントコメント、HTML、Markdown、またはプレーンテキストとして再フォーマットするには、[Printer]を使用します：

	var pr comment.Printer
	os.Stdout.Write(pr.Text(doc))

[Parser]と[Printer]の型は、その操作をカスタマイズするために変更できる構造体です。
詳細については、それらの型のドキュメントを参照してください。

再フォーマットに追加の制御が必要な使用例では、解析された構文自体を検査することで独自のロジックを実装できます。
概要および追加の型へのリンクについては、[Doc]、[Block]、[Text]のドキュメントを参照してください。
*/
package comment
