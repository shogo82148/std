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
// fset、pkg、およびposのパラメータの意味は [CheckExpr] と同じです。exprが正常にパースできないか、結果のexpr ASTが型チェックできない場合はエラーが返されます。
func Eval(fset *token.FileSet, pkg *Package, pos token.Pos, expr string) (_ TypeAndValue, err error)

// CheckExprは、式exprがパッケージpkgの位置posに現れたかのように型チェックを行います。
// 式についての [Type] 情報はinfoに記録されます。式は、実体化されていないジェネリック関数や型を示す識別子である場合があります。
//
// pkg == nilの場合、[Universe] のスコープが使用され、指定された位置posは無視されます。
// pkg != nilであり、posが無効な場合、パッケージのスコープが使用されます。それ以外の場合、posはパッケージに属している必要があります。
//
// posがパッケージ内にない場合や、ノードが型チェックできない場合はエラーが返されます。
//
// 注意: [Eval] とCheckExprは、型と値を計算するためにCheckを実行する代わりに使用するべきではありませんが、Checkに追加して使用する必要があります。
// これらの関数は式が使用される文脈（例: 代入）を無視するため、トップレベルの未型付き定数は対応する文脈固有の型ではなく、未型付きの型を返します。
func CheckExpr(fset *token.FileSet, pkg *Package, pos token.Pos, expr ast.Expr, info *Info) (err error)
