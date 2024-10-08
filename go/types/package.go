// Code generated by "go test -run=Generate -write=all"; DO NOT EDIT.
// Source: ../../cmd/compile/internal/types2/package.go

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// パッケージはGoのパッケージを説明します。
type Package struct {
	path      string
	name      string
	scope     *Scope
	imports   []*Package
	complete  bool
	fake      bool
	cgo       bool
	goVersion string
}

// NewPackageは、指定されたパッケージのパスと名前に対して新しいPackageを返します。
// このパッケージは完全ではなく、明示的なインポートは含まれていません。
func NewPackage(path, name string) *Package

// Pathはパッケージのパスを返します。
func (pkg *Package) Path() string

// Nameはパッケージ名を返します。
func (pkg *Package) Name() string

// SetNameはパッケージ名を設定します。
func (pkg *Package) SetName(name string)

// GoVersionはこのパッケージで必要な最小のGoバージョンを返します。
// もし最小バージョンが分からない場合、GoVersionは空文字列を返します。
// 各ソースファイルは、[go/ast.File.GoVersion]フィールドに報告されるよう、異なる最小Goバージョンを指定することができます。
func (pkg *Package) GoVersion() string

// Scopeは、パッケージレベルで宣言されたオブジェクト（TypeNames、Consts、Vars、およびFuncs）を保持する（完全または不完全な）パッケージスコープを返します。
// nilのpkgレシーバーの場合、ScopeはUniverseスコープを返します。
func (pkg *Package) Scope() *Scope

// パッケージは、そのスコープに（少なくとも）すべての公開オブジェクトが含まれている場合は完全であり、それ以外の場合は不完全です。
func (pkg *Package) Complete() bool

// MarkCompleteはパッケージを完了としてマークします。
func (pkg *Package) MarkComplete()

// Importsは、pkgによって直接インポートされたパッケージのリストを返します。リストはソース順に並んでいます。
// pkgがエクスポートデータからロードされた場合、Importsにはpkgが参照しているパッケージレベルのオブジェクトを提供するパッケージが含まれます。これは、pkgのソースコードに直接インポートされたパッケージのセットよりも多くまたは少ない場合があります。
// pkgがcgoを使用し、FakeImportC構成オプションが有効になっている場合、インポートリストには偽の「C」パッケージが含まれている可能性があります。
func (pkg *Package) Imports() []*Package

// SetImportsは、明示的にインポートされるパッケージのリストを設定します。
// リストの要素が一意であることは、呼び出し元の責任です。
func (pkg *Package) SetImports(list []*Package)

func (pkg *Package) String() string
