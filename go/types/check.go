// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは型チェックを実行するCheck関数を実装しています。

package types

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// Checkerは型チェッカーの状態を維持します。
// [NewChecker] で作成する必要があります。
type Checker struct {

	// If EnableAlias is set, alias declarations produce an Alias type.
	// Otherwise the alias information is only in the type name, which
	// points directly to the actual (aliased) type.
	enableAlias bool

	conf *Config
	ctxt *Context
	fset *token.FileSet
	pkg  *Package
	*Info
	version goVersion
	nextID  uint64
	objMap  map[Object]*declInfo
	impMap  map[importKey]*Package

	// pkgPathMapはパッケージ名をインポートパスの集合にマッピングします。
	// インポートグラフのどこかでその名前に対して見た異なるインポートパスの集合です。
	// エラーメッセージでパッケージ名を曖昧にするために使用されます。
	//
	// pkgPathMapは遅延して割り当てられており、ハッピーな経路上で構築するコストを支払いません。
	// seenPkgMapはすでに歩いたパッケージを追跡します。
	pkgPathMap map[string]map[string]bool
	seenPkgMap map[*Package]bool

	// パッケージファイルのセットの型チェック中に収集された情報
	// (Filesによって初期化され、check.Filesの間のみ有効です;
	// マップとリストは必要に応じて割り当てられます)
	files         []*ast.File
	versions      map[*ast.File]string
	imports       []*PkgName
	dotImportMap  map[dotImportKey]*PkgName
	recvTParamMap map[*ast.Ident]*TypeParam
	brokenAliases map[*TypeName]bool
	unionTypeSets map[*Union]*_TypeSet
	mono          monoGraph

	firstErr error
	methods  map[*TypeName][]*Func
	untyped  map[ast.Expr]exprInfo
	delayed  []action
	objPath  []Object
	cleaners []cleaner

	// 現在のオブジェクトが型チェックされる環境（特定のオブジェクトの型チェックの間のみ有効）
	environment

	// デバッグ中
	indent int
}

// NewCheckerは指定されたパッケージに対して新しい [Checker] インスタンスを返します。
// [Package] ファイルは、checker.Filesを通じて段階的に追加することができます。
func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker

// Filesはチェッカーのパッケージの一部として提供されたファイルをチェックします。
func (check *Checker) Files(files []*ast.File) error
