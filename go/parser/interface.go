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

<<<<<<< HEAD
// ParseFile parses the source code of a single Go source file and returns
// the corresponding ast.File node. The source code may be provided via
// the filename of the source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is
// only used when recording position information. The type of the argument
// for the src parameter must be string, []byte, or io.Reader.
// If src == nil, ParseFile parses the file specified by filename.
//
// The mode parameter controls the amount of source text parsed and
// other optional parser functionality. If the SkipObjectResolution
// mode bit is set (recommended), the object resolution phase of
// parsing will be skipped, causing File.Scope, File.Unresolved, and
// all Ident.Obj fields to be nil. Those fields are deprecated; see
// [ast.Object] for details.
//
// Position information is recorded in the file set fset, which must not be
// nil.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure. If the source was read but syntax
// errors were found, the result is a partial AST (with ast.Bad* nodes
// representing the fragments of erroneous source code). Multiple errors
// are returned via a scanner.ErrorList which is sorted by source position.
=======
// ParseFileは単一のGoソースファイルのソースコードを解析し、対応するast.Fileノードを返します。ソースコードはソースファイルのファイル名またはsrcパラメータを介して指定できます。
// src != nilの場合、ParseFileはsrcからソースを解析し、ファイル名は位置情報の記録にのみ使用されます。srcパラメータの引数のタイプはstring、[]byte、またはio.Readerである必要があります。src == nilの場合、ParseFileはfilenameで指定されたファイルを解析します。
// modeパラメータはソーステキストの解析量とその他のオプションのパーサ機能を制御します。SkipObjectResolutionモードビットがセットされている場合、解析のオブジェクト解決フェーズがスキップされ、File.Scope、File.Unresolved、およびすべてのIdent.Objフィールドはnilになります。
// 位置情報はnilではないファイルセットfsetに記録されます。
// ソースを読み取ることができなかった場合、返されるASTはnilであり、エラーには具体的な失敗が示されます。ソースが読み取られたが構文エラーが見つかった場合、結果は部分的なAST（ast.Bad*ノードがエラーソースコードのフラグメントを表す）です。複数のエラーはソース位置でソートされたscanner.ErrorListを介して返されます。
>>>>>>> release-branch.go1.21
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)

// ParseDirは、指定されたパスのディレクトリ内で拡張子が".go"で終わるすべてのファイルに対してParseFileを呼び出し、
// 見つかったすべてのパッケージ名->パッケージASTのマップを返します。
//
// もしfilter != nilなら、フィルタを通過するfs.FileInfoエントリを持つ（かつ".go"で終わる）ファイルのみを考慮します。
// モードビットはParseFileに変更なしで渡されます。
// 位置情報はfsetに記録されますが、これはnilであってはなりません。
//
// ディレクトリが読み込めなかった場合、nilのマップと対応するエラーが返されます。
// パースエラーが発生した場合、非nilで不完全なマップと最初に遭遇したエラーが返されます。
func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)

// ParseExprFromは式を解析するための便利な関数です。
// 引数の意味はParseFileと同じですが、ソースは有効なGo（型または値）の式である必要があります。具体的には、fsetはnilであってはなりません。
//
// ソースが読み取れなかった場合、返されるASTはnilであり、エラーは特定の失敗を示します。ソースは読み取られたが構文エラーが見つかった場合、結果は部分的なAST（ast.Bad*ノードが誤ったソースコードの断片を表す）です。複数のエラーは、ソースの位置でソートされたscanner.ErrorListを介して返されます。
func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error)

// ParseExprは式xのASTを取得するための便利関数です。
// ASTに記録される位置情報は未定義です。エラーメッセージで使用されるファイル名は空の文字列です。
//
// 文法エラーが見つかった場合、結果は部分的なASTです（ast.Bad*ノードがエラーの断片を表します）。複数のエラーはソース位置でソートされたscanner.ErrorListを介して返されます。
func ParseExpr(x string) (ast.Expr, error)
