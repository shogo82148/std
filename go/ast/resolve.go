// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはNewPackageを実装しています。

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// Importerは、インポートパスをパッケージオブジェクトに変換します。
// importsマップは、既にインポートされたパッケージを記録し、パッケージID（正規のインポートパス）で索引付けされます。
// Importerは正規のインポートパスを決定し、importsマップに既に存在するかどうかを確認する必要があります。
// もし存在する場合、Importerはマップのエントリを返すことができます。
// そうでない場合、Importerは指定されたパスのパッケージデータを新しい*Object（pkg）に読み込み、
// pkgをimportsマップに記録した後、pkgを返します。
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)

// NewPackageは、一連のFileノードから新しいPackageノードを作成します。ファイル間の未解決の識別子を解決し、各ファイルの未解決リストを更新します。非nilのimporterとuniverseスコープが指定されている場合、パッケージファイルで宣言されていない識別子を解決するために使用されます。残りの未解決の識別子は宣言されていないとして報告されます。ファイルが異なるパッケージに属する場合、パッケージ名が選択され、パッケージ名が異なるファイルが報告された後、無視されます。
// 結果は、パッケージノードとscanner.ErrorListです。エラーがある場合にのみ返されます。
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)
