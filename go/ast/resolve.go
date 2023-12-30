// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはNewPackageを実装しています。

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// Importerは、インポートパスをパッケージオブジェクトに解決します。
// importsマップは、すでにインポートされたパッケージを記録し、
// パッケージID（正規のインポートパス）でインデックス付けします。
// Importerは、正規のインポートパスを決定し、
// それがすでにimportsマップに存在するかどうかを確認する必要があります。
// もしそうであれば、Importerはマップエントリを返すことができます。そうでなければ、
// Importerは指定されたパスのパッケージデータを新しい *[Object](pkg)にロードし、
// pkgをimportsマップに記録し、その後pkgを返すべきです。
//
// Deprecated: 代わりに型チェッカー [go/types] を使用してください。詳細は [Object] を参照してください。
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)

// NewPackageは、[File] ノードのセットから新しい [Package] ノードを作成します。それは
// ファイル間の未解決の識別子を解決し、それぞれのファイルのUnresolvedリストを
// それに応じて更新します。非nilのimporterとuniverseスコープが提供されている場合、
// パッケージファイルのいずれにも宣言されていない識別子を解決するために使用されます。任意の
// 残りの未解決の識別子は未宣言として報告されます。ファイルが
// 異なるパッケージに属している場合、一つのパッケージ名が選択され、
// 異なるパッケージ名を持つファイルが報告され、その後無視されます。
// 結果はパッケージノードと、エラーがあった場合は [scanner.ErrorList] です。
//
// Deprecated: 代わりに型チェッカー [go/types] を使用してください。詳細は [Object] を参照してください。
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)
