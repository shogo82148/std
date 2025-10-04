// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

// FileExportsは、GoのソースファイルのASTを現在の場所でトリムします。
// エクスポートされたノードのみが残り、エクスポートされていないトップレベルの識別子とそれに関連する情報
// （型、初期値、または関数本体など）は削除されます。エクスポートされた型の非エクスポートフィールドとメソッドも剥ぎ取られます。
// [File.Comments] リストは変更されません。
//
// FileExportsは、エクスポートされた宣言があるかどうかを報告します。
func FileExports(src *File) bool

// PackageExportsは、GoパッケージのASTを変更して、エクスポートされたノードのみが残るようにします。pkg.Filesリストは変更されず、ファイル名とトップレベルのパッケージコメントは失われません。
//
<<<<<<< HEAD
// PackageExportsは、エクスポートされた宣言があるかどうかを報告します。エクスポートされた宣言がない場合、falseを返します。
=======
// PackageExports reports whether there are exported declarations;
// it returns false otherwise.
//
// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object]. Alternatively, use [FileExports].
>>>>>>> upstream/release-branch.go1.25
func PackageExports(pkg *Package) bool

type Filter func(string) bool

// FilterDeclはGoの宣言のASTを変更して、フィルターfを通過しない名前（構造体フィールドやインタフェースメソッドの名前を含むが、パラメーターリストからは除外）を削除します。
//
// FilterDeclは、フィルタリング後に残された宣言された名前があるかどうかを報告します。
func FilterDecl(decl Decl, f Filter) bool

// FilterFileは、フィルタfを通過しない（構造体のフィールドやインターフェースのメソッド名を含むが、パラメータリストからは含まれない）トップレベルの宣言からすべての名前を削除することで、GoファイルのASTを修正します。もし宣言が空になった場合、宣言はASTから削除されます。Import宣言は必ず削除されます。[File.Comments] のリストは変更されません。
// FilterFileは、フィルタリング後にトップレベルの宣言が残っているかどうかを報告します。
func FilterFile(src *File, f Filter) bool

// FilterPackageは、フィルターfを通過しない（構造体フィールドやインターフェースメソッド名を含むが、パラメータリストからは除かれない）トップレベル宣言のすべての名前を削除することにより、GoパッケージのASTを修正します。 宣言がその後空になった場合、宣言はASTから削除されます。 pkg.Filesリストは変更されないため、ファイル名やトップレベルのパッケージコメントが失われることはありません。
//
<<<<<<< HEAD
// FilterPackageは、フィルタリング後にトップレベルの宣言が残っているかどうかを報告します。
func FilterPackage(pkg *Package, f Filter) bool

// [MergePackageFiles] の動作を制御するMergeModeフラグ。
=======
// FilterPackage reports whether there are any top-level declarations
// left after filtering.
//
// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object]. Alternatively, use [FilterFile].
func FilterPackage(pkg *Package, f Filter) bool

// The MergeMode flags control the behavior of [MergePackageFiles].
//
// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object].
>>>>>>> upstream/release-branch.go1.25
type MergeMode uint

// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object].
const (
	// セットされている場合、重複する関数宣言は除外されます。
	FilterFuncDuplicates MergeMode = 1 << iota

	// セットされている場合、特定のASTノード（DocやCommentなど）に関連付けられていないコメントは除外されます。
	FilterUnassociatedComments
	// もし設定されていた場合、重複したインポート宣言は除外されます。
	FilterImportDuplicates
)

<<<<<<< HEAD
// MergePackageFilesはパッケージに所属するファイルのASTをマージしてファイルASTを作成します。モードフラグはマージの動作を制御します。
=======
// MergePackageFiles creates a file AST by merging the ASTs of the
// files belonging to a package. The mode flags control merging behavior.
//
// Deprecated: this function is poorly specified and has unfixable
// bugs; also [Package] is deprecated.
>>>>>>> upstream/release-branch.go1.25
func MergePackageFiles(pkg *Package, mode MergeMode) *File
