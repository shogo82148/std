// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは型チェックを実行するCheck関数を実装しています。

package types

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

<<<<<<< HEAD
// Checkerは型チェッカーの状態を維持します。
// NewCheckerで作成する必要があります。
=======
// A Checker maintains the state of the type checker.
// It must be created with [NewChecker].
>>>>>>> upstream/master
type Checker struct {

	// パッケージの情報
	// (NewCheckerによって初期化され、checkerの寿命の間有効)
	conf *Config
	ctxt *Context
	fset *token.FileSet
	pkg  *Package
	*Info
	version version
	posVers map[token.Pos]version
	nextID  uint64
	objMap  map[Object]*declInfo
	impMap  map[importKey]*Package
	valids  instanceLookup

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

<<<<<<< HEAD
// NewCheckerは指定されたパッケージに対して新しいCheckerインスタンスを返します。
// パッケージファイルは、checker.Filesを通じて段階的に追加することができます。
=======
// NewChecker returns a new [Checker] instance for a given package.
// [Package] files may be added incrementally via checker.Files.
>>>>>>> upstream/master
func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker

// Filesはチェッカーのパッケージの一部として提供されたファイルをチェックします。
func (check *Checker) Files(files []*ast.File) error
