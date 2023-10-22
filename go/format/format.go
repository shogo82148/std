// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージformatはGoソースコードの標準的なフォーマットを実装します。
//
// Goソースコードのフォーマットは時間とともに変化するため、
// 一貫したフォーマットに依存するツールは、このパッケージを使う代わりに特定のバージョンのgofmtバイナリを実行する必要があります。
// その方法で、フォーマットが安定し、ツールをgoftmtの変更ごとに再コンパイルする必要がなくなります。
//
// たとえば、このパッケージを直接使用するプレサブミットチェックは、
// 開発者が使用しているGoのバージョンによって異なる動作をするため、不安定になる可能性があります。
package format

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// Nodeはソースコードを標準的なgofmtスタイルに整形し、結果をdstに書き込みます。
//
<<<<<<< HEAD
// nodeの型は*ast.File、*printer.CommentedNode、[]ast.Decl、[]ast.Stmtのいずれかである必要があります。
// もしくはast.Expr、ast.Decl、ast.Spec、ast.Stmtと互換性のある代入可能な型である必要があります。
// Nodeはnodeを変更しません。部分的なソースファイルを表すノードの場合、
// （例えば、nodeが*ast.Fileでない場合や*printer.CommentedNodeが*ast.Fileを包んでいない場合）インポートはソートされません。
=======
// The node type must be *[ast.File], *[printer.CommentedNode], [][ast.Decl],
// [][ast.Stmt], or assignment-compatible to [ast.Expr], [ast.Decl], [ast.Spec],
// or [ast.Stmt]. Node does not modify node. Imports are not sorted for
// nodes representing partial source files (for instance, if the node is
// not an *[ast.File] or a *[printer.CommentedNode] not wrapping an *[ast.File]).
>>>>>>> upstream/master
//
// 関数は早期に（結果が完全に書き込まれる前に）戻って、
// 正しくないASTのためにフォーマットエラーを返す場合があります。
func Node(dst io.Writer, fset *token.FileSet, node any) error

// この関数は、ソースコードが正規のgofmtスタイルで書かれていると仮定して、
// ソースコードを変換し結果を返します。エラーが発生した場合はそれも返されます。
// srcは構文的に正しいGoのソースファイル、またはGoの宣言または文のリストであることが期待されます。
//
// srcが部分的なソースファイルの場合、srcの先頭と末尾のスペースが結果に適用されます
// (つまり、先頭と末尾のスペースがsrcと同じになるように)。
// また、結果はsrcのコードを含む最初の行と同じだけインデントされます。
// 部分的なソースファイルでは、インポートステートメントはソートされません。
func Source(src []byte) ([]byte, error)
