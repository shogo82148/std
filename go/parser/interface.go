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
// ParseFileは、単一のGoソースファイルのソースコードを解析し、対応するast.Fileノードを返します。
// ソースコードは、ソースファイルのファイル名またはsrcパラメーターを介して提供できます。
//
// src != nilの場合、ParseFileはsrcからソースを解析し、ファイル名は位置情報を記録するときにのみ使用されます。
// srcパラメーターの引数の型は、string、[]byte、またはio.Readerである必要があります。
// src == nilの場合、ParseFileはfilenameで指定されたファイルを解析します。
//
// modeパラメーターは、解析されるソーステキストの量とその他のオプションのパーサー機能を制御します。
// SkipObjectResolutionモードビットが設定されている場合（推奨）、解析のオブジェクト解決フェーズがスキップされ、
// File.Scope、File.Unresolved、およびすべてのIdent.Objフィールドがnilになります。
// これらのフィールドは非推奨です。詳細については、 [ast.Object] を参照してください。
=======
// ParseFile parses the source code of a single Go source file and returns
// the corresponding [ast.File] node. The source code may be provided via
// the filename of the source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is
// only used when recording position information. The type of the argument
// for the src parameter must be string, []byte, or [io.Reader].
// If src == nil, ParseFile parses the file specified by filename.
//
// The mode parameter controls the amount of source text parsed and
// other optional parser functionality. If the [SkipObjectResolution]
// mode bit is set (recommended), the object resolution phase of
// parsing will be skipped, causing File.Scope, File.Unresolved, and
// all Ident.Obj fields to be nil. Those fields are deprecated; see
// [ast.Object] for details.
>>>>>>> upstream/master
//
// 位置情報は、nilであってはならないファイルセットfsetに記録されます。
//
<<<<<<< HEAD
// ソースを読み込めなかった場合、返されるASTはnilであり、エラーは特定の失敗を示します。
// ソースが読み込まれたが、構文エラーが見つかった場合、結果は部分的なAST（ast.Bad*ノードがエラーの断片を表す）です。
// 複数のエラーは、ソース位置でソートされたscanner.ErrorListを介して返されます。
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)

// ParseDirは、指定されたパスのディレクトリ内で拡張子が".go"で終わるすべてのファイルに対してParseFileを呼び出し、
// 見つかったすべてのパッケージ名->パッケージASTのマップを返します。
//
// もしfilter != nilなら、フィルタを通過するfs.FileInfoエントリを持つ（かつ".go"で終わる）ファイルのみを考慮します。
// モードビットはParseFileに変更なしで渡されます。
// 位置情報はfsetに記録されますが、これはnilであってはなりません。
=======
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure. If the source was read but syntax
// errors were found, the result is a partial AST (with [ast.Bad]* nodes
// representing the fragments of erroneous source code). Multiple errors
// are returned via a scanner.ErrorList which is sorted by source position.
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)

// ParseDir calls [ParseFile] for all files with names ending in ".go" in the
// directory specified by path and returns a map of package name -> package
// AST with all the packages found.
//
// If filter != nil, only the files with [fs.FileInfo] entries passing through
// the filter (and ending in ".go") are considered. The mode bits are passed
// to [ParseFile] unchanged. Position information is recorded in fset, which
// must not be nil.
>>>>>>> upstream/master
//
// ディレクトリが読み込めなかった場合、nilのマップと対応するエラーが返されます。
// パースエラーが発生した場合、非nilで不完全なマップと最初に遭遇したエラーが返されます。
func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)

<<<<<<< HEAD
// ParseExprFromは式を解析するための便利な関数です。
// 引数の意味はParseFileと同じですが、ソースは有効なGo（型または値）の式である必要があります。具体的には、fsetはnilであってはなりません。
//
// ソースが読み取れなかった場合、返されるASTはnilであり、エラーは特定の失敗を示します。ソースは読み取られたが構文エラーが見つかった場合、結果は部分的なAST（ast.Bad*ノードが誤ったソースコードの断片を表す）です。複数のエラーは、ソースの位置でソートされたscanner.ErrorListを介して返されます。
=======
// ParseExprFrom is a convenience function for parsing an expression.
// The arguments have the same meaning as for [ParseFile], but the source must
// be a valid Go (type or value) expression. Specifically, fset must not
// be nil.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure. If the source was read but syntax
// errors were found, the result is a partial AST (with [ast.Bad]* nodes
// representing the fragments of erroneous source code). Multiple errors
// are returned via a scanner.ErrorList which is sorted by source position.
>>>>>>> upstream/master
func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error)

// ParseExprは式xのASTを取得するための便利関数です。
// ASTに記録される位置情報は未定義です。エラーメッセージで使用されるファイル名は空の文字列です。
//
<<<<<<< HEAD
// 文法エラーが見つかった場合、結果は部分的なASTです（ast.Bad*ノードがエラーの断片を表します）。複数のエラーはソース位置でソートされたscanner.ErrorListを介して返されます。
=======
// If syntax errors were found, the result is a partial AST (with [ast.Bad]* nodes
// representing the fragments of erroneous source code). Multiple errors are
// returned via a scanner.ErrorList which is sorted by source position.
>>>>>>> upstream/master
func ParseExpr(x string) (ast.Expr, error)
