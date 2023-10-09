// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはNewPackageを実装しています。

package ast

import (
	"github.com/shogo82148/std/go/token"
)

<<<<<<< HEAD
// An Importer resolves import paths to package Objects.
// The imports map records the packages already imported,
// indexed by package id (canonical import path).
// An Importer must determine the canonical import path and
// check the map to see if it is already present in the imports map.
// If so, the Importer can return the map entry. Otherwise, the
// Importer should load the package data for the given path into
// a new *Object (pkg), record pkg in the imports map, and then
// return pkg.
//
// Deprecated: use the type checker [go/types] instead; see [Object].
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)

// NewPackage creates a new Package node from a set of File nodes. It resolves
// unresolved identifiers across files and updates each file's Unresolved list
// accordingly. If a non-nil importer and universe scope are provided, they are
// used to resolve identifiers not declared in any of the package files. Any
// remaining unresolved identifiers are reported as undeclared. If the files
// belong to different packages, one package name is selected and files with
// different package names are reported and then ignored.
// The result is a package node and a scanner.ErrorList if there were errors.
//
// Deprecated: use the type checker [go/types] instead; see [Object].
=======
// Importerは、インポートパスをパッケージオブジェクトに変換します。
// importsマップは、既にインポートされたパッケージを記録し、パッケージID（正規のインポートパス）で索引付けされます。
// Importerは正規のインポートパスを決定し、importsマップに既に存在するかどうかを確認する必要があります。
// もし存在する場合、Importerはマップのエントリを返すことができます。
// そうでない場合、Importerは指定されたパスのパッケージデータを新しい*Object（pkg）に読み込み、
// pkgをimportsマップに記録した後、pkgを返します。
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)

// NewPackageは、一連のFileノードから新しいPackageノードを作成します。ファイル間の未解決の識別子を解決し、各ファイルの未解決リストを更新します。非nilのimporterとuniverseスコープが指定されている場合、パッケージファイルで宣言されていない識別子を解決するために使用されます。残りの未解決の識別子は宣言されていないとして報告されます。ファイルが異なるパッケージに属する場合、パッケージ名が選択され、パッケージ名が異なるファイルが報告された後、無視されます。
// 結果は、パッケージノードとscanner.ErrorListです。エラーがある場合にのみ返されます。
>>>>>>> release-branch.go1.21
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)
