// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// importerパッケージは、エクスポートデータのインポートを提供します。
=======
// Package importer provides access to export data importers.
//
// These functions, which are mostly deprecated, date from before the
// introduction of modules in release Go 1.11. They should no longer
// be relied on except for use in test cases using small programs that
// depend only on the standard library. For reliable module-aware
// loading of type information, use the packages.Load function from
// golang.org/x/tools/go/packages.
>>>>>>> upstream/release-branch.go1.25
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
<<<<<<< HEAD
// Deprecated:  importerによって作成されたオブジェクトの位置を
// FileSetで設定するために [ForCompiler] を使用してください。
func For(compiler string, lookup Lookup) types.Importer

// Defaultは実行バイナリをビルドしたコンパイラのためのImporterを返します。
// もし利用可能であれば、結果は [types.ImporterFrom] を実装します。
=======
// Deprecated: Use [ForCompiler], which populates a FileSet
// with the positions of objects created by the importer.
//
//go:fix inline
func For(compiler string, lookup Lookup) types.Importer

// Default returns an Importer for the compiler that built the running binary.
// If available, the result implements [types.ImporterFrom].
//
// Default may be convenient for use in the simplest of cases, but
// most clients should instead use [ForCompiler], which accepts a
// [token.FileSet] from the caller; without it, all position
// information derived from the Importer will be incorrect and
// misleading. See also the package documentation.
>>>>>>> upstream/release-branch.go1.25
func Default() types.Importer
