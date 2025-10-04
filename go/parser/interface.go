// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルには、パーサーを呼び出すためのエクスポートされたエントリーポイントが含まれています。

package parser

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io/fs"
)

// モード値はフラグのセット（または0）です。
// これらはソースコードの解析量やその他のオプションの
// パーサー機能を制御します。
type Mode uint

const (
	PackageClauseOnly Mode = 1 << iota
	ImportsOnly
	ParseComments
	Trace
	DeclarationErrors
	SpuriousErrors
	SkipObjectResolution
	AllErrors = SpuriousErrors
)

// ParseFileは、単一のGoソースファイルのソースコードを解析し、対応する [ast.File] ノードを返します。
// ソースコードは、ソースファイルのファイル名またはsrcパラメーターを介して提供できます。
//
// src != nilの場合、ParseFileはsrcからソースを解析し、ファイル名は位置情報を記録するときにのみ使用されます。
// srcパラメーターの引数の型は、string、[]byte、または [io.Reader] である必要があります。
// src == nilの場合、ParseFileはfilenameで指定されたファイルを解析します。
//
// modeパラメーターは、解析されるソーステキストの量とその他のオプションのパーサー機能を制御します。
// [SkipObjectResolution] モードビットが設定されている場合（推奨）、解析のオブジェクト解決フェーズがスキップされ、
// File.Scope、File.Unresolved、およびすべてのIdent.Objフィールドがnilになります。
// これらのフィールドは非推奨です。詳細については、 [ast.Object] を参照してください。
//
// 位置情報は、nilであってはならないファイルセットfsetに記録されます。
//
// ソースを読み込めなかった場合、返されるASTはnilであり、エラーは特定の失敗を示します。
// ソースが読み込まれたが、構文エラーが見つかった場合、結果は部分的なAST（[ast.Bad]*ノードがエラーの断片を表す）です。
// 複数のエラーは、ソース位置でソートされたscanner.ErrorListを介して返されます。
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)

// ParseDirは、指定されたパスのディレクトリ内で拡張子が".go"で終わるすべてのファイルに対して [ParseFile] を呼び出し、
// 見つかったすべてのパッケージ名->パッケージASTのマップを返します。
//
// もしfilter != nilなら、フィルタを通過する [fs.FileInfo] エントリを持つ（かつ".go"で終わる）ファイルのみを考慮します。
// モードビットは [ParseFile] に変更なしで渡されます。
// 位置情報はfsetに記録されますが、これはnilであってはなりません。
//
// ディレクトリが読み取れなかった場合、nilマップとそれぞれのエラーが
// 返されます。解析エラーが発生した場合、nilではないが不完全なマップと
// 最初に遭遇したエラーが返されます。
//
// Deprecated: ParseDirはファイルをパッケージに関連付ける際にビルドタグを考慮しません。
// パッケージとファイルの関係に関する正確な情報については、
// golang.org/x/tools/go/packages を使用してください。
// これはオプションでファイルの解析と型チェックも行うことができます。
func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)

// ParseExprFromは式を解析するための便利な関数です。
// 引数の意味は [ParseFile] と同じですが、ソースは有効なGo（型または値）の式である必要があります。具体的には、fsetはnilであってはなりません。
//
// ソースが読み取れなかった場合、返されるASTはnilであり、エラーは特定の失敗を示します。ソースは読み取られたが構文エラーが見つかった場合、結果は部分的なAST（[ast.Bad]*ノードが誤ったソースコードの断片を表す）です。複数のエラーは、ソースの位置でソートされたscanner.ErrorListを介して返されます。
func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error)

// ParseExprは式xのASTを取得するための便利関数です。
// ASTに記録される位置情報は未定義です。エラーメッセージで使用されるファイル名は空の文字列です。
//
// 文法エラーが見つかった場合、結果は部分的なASTです（[ast.Bad]*ノードがエラーの断片を表します）。複数のエラーはソース位置でソートされたscanner.ErrorListを介して返されます。
func ParseExpr(x string) (ast.Expr, error)
