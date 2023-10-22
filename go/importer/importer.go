// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージインポーターは、エクスポートデータのインポートを提供します。
package importer

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/io"
)

// Lookup 関数は、与えられたインポートパスに対してパッケージデータにアクセスするためのリーダー、または一致するパッケージが見つからない場合はエラーを返します。
type Lookup func(path string) (io.ReadCloser, error)

// ForCompilerは、インストールされたパッケージからインポートするためのImporterを返します。
// コンパイラの引数が「gc」または「gccgo」の場合、またはコンパイラの引数が「source」の場合はソースから直接インポートするためのものです。
// この後者の場合、純粋なGoソースコードに完全に定義されていない場合には、インポートが失敗することがあります
// (パッケージAPIがcgoで定義されたエンティティに依存する場合、型チェッカはそれにアクセスできません)。
// 結果のImporterがインポートパスの解決に必要な場合、lookup関数が呼び出されます。
// このモードでは、インポータには正規のインポートパス(相対的または絶対のものではない)でのみ呼び出されるものとします。
// 正規のインポートパスへの変換は、インポータのクライアントによって行われているものとします。
// 正しいモジュール対応動作のためには、lookup関数を提供する必要があります。
// Deprecated: もしlookupがnilの場合、後方互換性のためにインポータは$GOPATHワークスペースでインポートを解決しようとします。
func ForCompiler(fset *token.FileSet, compiler string, lookup Lookup) types.Importer

<<<<<<< HEAD
// 新しいFileSetでForCompilerを呼び出します。
//
// Deprecated:  importerによって作成されたオブジェクトの位置を
// FileSetで設定するためにForCompilerを使用してください。
func For(compiler string, lookup Lookup) types.Importer

// Defaultは実行バイナリをビルドしたコンパイラのためのImporterを返します。
// もし利用可能であれば、結果はtypes.ImporterFromを実装します。
=======
// For calls [ForCompiler] with a new FileSet.
//
// Deprecated: Use [ForCompiler], which populates a FileSet
// with the positions of objects created by the importer.
func For(compiler string, lookup Lookup) types.Importer

// Default returns an Importer for the compiler that built the running binary.
// If available, the result implements [types.ImporterFrom].
>>>>>>> upstream/master
func Default() types.Importer
