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

// ForCompilerは、コンパイラ「gc」および「gccgo」のインストールされたパッケージからインポートするための
// Importerを返します。または、コンパイラ引数が「source」の場合はソースから直接インポートします。
// 後者の場合、エクスポートされたAPIが純粋なGoソースコードで完全に定義されていない状況では、
// インポートが失敗する可能性があります（パッケージAPIがcgo定義のエンティティに依存している場合、
// 型チェッカーはそれらにアクセスできません）。
//
// lookup関数は、結果のインポータがインポートパスを解決する必要があるたびに呼び出されます。
// このモードでは、インポータは正規のインポートパスでのみ呼び出すことができます
// （相対パスまたは絶対パスではなく）。インポートパスの正規形への変換がインポータの
// クライアントによって行われていると想定されます。
//
// 正しいモジュール対応操作のため、lookup関数を提供する必要があります。lookupにnil値を
// 提供することは非推奨ですが、後方互換性のため、インポータはこの場合$GOPATH
// ワークスペース内でインポートを解決しようとします。
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
