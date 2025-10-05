// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// importerパッケージは、エクスポートデータインポータへのアクセスを提供します。
//
// これらの関数は、ほとんどが非推奨であり、Go 1.11リリースでモジュールが導入される前の
// ものです。標準ライブラリのみに依存する小さなプログラムを使用するテストケースでの
// 使用を除いて、もはや依存すべきではありません。型情報の信頼性のあるモジュール対応
// ロードについては、golang.org/x/tools/go/packages のpackages.Load関数を使用してください。
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

// 新しいFileSetで [ForCompiler] を呼び出します。
//
// Deprecated: インポータによって作成されたオブジェクトの位置情報を
// FileSetに設定する [ForCompiler] を使用してください。
//
//go:fix inline
func For(compiler string, lookup Lookup) types.Importer

// Defaultは実行中のバイナリをビルドしたコンパイラ用のImporterを返します。
// 利用可能な場合、結果は [types.ImporterFrom] を実装します。
//
// Defaultは最もシンプルなケースでの使用には便利かもしれませんが、
// ほとんどのクライアントは代わりに、呼び出し元から [token.FileSet] を受け取る
// [ForCompiler] を使用すべきです；それなしでは、Importerから派生した
// すべての位置情報が不正確で誤解を招くものになります。パッケージドキュメントも参照してください。
func Default() types.Importer
