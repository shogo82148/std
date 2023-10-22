// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// Evalは、パッケージpkg内の位置posで評価された式exprの型と、定数であれば値を返します。
// これは、提供されたファイルセットに対して完全な位置情報を持つASTを型チェックして導出したpkgに関連している必要があります。
//
<<<<<<< HEAD
// fset、pkg、およびposのパラメータの意味はCheckExprと同じです。exprが正常にパースできないか、結果のexpr ASTが型チェックできない場合はエラーが返されます。
func Eval(fset *token.FileSet, pkg *Package, pos token.Pos, expr string) (_ TypeAndValue, err error)

// CheckExprは、式exprがパッケージpkgの位置posに現れたかのように型チェックを行います。
// 式についての型情報はinfoに記録されます。式は、実体化されていないジェネリック関数や型を示す識別子である場合があります。
//
// pkg == nilの場合、ユニバースのスコープが使用され、指定された位置posは無視されます。
// pkg != nilであり、posが無効な場合、パッケージのスコープが使用されます。それ以外の場合、posはパッケージに属している必要があります。
=======
// The meaning of the parameters fset, pkg, and pos is the
// same as in [CheckExpr]. An error is returned if expr cannot
// be parsed successfully, or the resulting expr AST cannot be
// type-checked.
func Eval(fset *token.FileSet, pkg *Package, pos token.Pos, expr string) (_ TypeAndValue, err error)

// CheckExpr type checks the expression expr as if it had appeared at position
// pos of package pkg. [Type] information about the expression is recorded in
// info. The expression may be an identifier denoting an uninstantiated generic
// function or type.
//
// If pkg == nil, the [Universe] scope is used and the provided
// position pos is ignored. If pkg != nil, and pos is invalid,
// the package scope is used. Otherwise, pos must belong to the
// package.
>>>>>>> upstream/master
//
// posがパッケージ内にない場合や、ノードが型チェックできない場合はエラーが返されます。
//
<<<<<<< HEAD
// 注意: EvalとCheckExprは、型と値を計算するためにCheckを実行する代わりに使用するべきではありませんが、Checkに追加して使用する必要があります。
// これらの関数は式が使用される文脈（例: 代入）を無視するため、トップレベルの未型付き定数は対応する文脈固有の型ではなく、未型付きの型を返します。
=======
// Note: [Eval] and CheckExpr should not be used instead of running Check
// to compute types and values, but in addition to Check, as these
// functions ignore the context in which an expression is used (e.g., an
// assignment). Thus, top-level untyped constants will return an
// untyped type rather than the respective context-specific type.
>>>>>>> upstream/master
func CheckExpr(fset *token.FileSet, pkg *Package, pos token.Pos, expr ast.Expr, info *Info) (err error)
