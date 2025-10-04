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
// PackageExportsは、エクスポートされた宣言があるかどうかを報告します；
// それ以外の場合はfalseを返します。
//
// Deprecated: [Package] の代わりに型チェッカー [go/types] を使用してください；
// [Object] を参照してください。または、[FileExports] を使用してください。
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
// FilterPackageは、フィルタリング後にトップレベルの宣言が
// 残っているかどうかを報告します。
//
// Deprecated: [Package] の代わりに型チェッカー [go/types] を使用してください；
// [Object] を参照してください。または、[FilterFile] を使用してください。
func FilterPackage(pkg *Package, f Filter) bool

// MergeModeフラグは [MergePackageFiles] の動作を制御します。
//
// Deprecated: [Package] の代わりに型チェッカー [go/types] を使用してください；
// [Object] を参照してください。
type MergeMode uint

// Deprecated: [Package] の代わりに型チェッカー [go/types] を使用してください；
// [Object] を参照してください。
const (
	// セットされている場合、重複する関数宣言は除外されます。
	FilterFuncDuplicates MergeMode = 1 << iota

	// セットされている場合、特定のASTノード（DocやCommentなど）に関連付けられていないコメントは除外されます。
	FilterUnassociatedComments
	// もし設定されていた場合、重複したインポート宣言は除外されます。
	FilterImportDuplicates
)

// MergePackageFilesは、パッケージに属するファイルのASTをマージして
// ファイルASTを作成します。modeフラグはマージ動作を制御します。
//
// Deprecated: この関数は仕様が不十分で修正できない
// バグがあります；また [Package] は非推奨です。
func MergePackageFiles(pkg *Package, mode MergeMode) *File
