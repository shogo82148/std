// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package doc はGo ASTからソースコードのドキュメンテーションを取得します。
package doc

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/doc/comment"
	"github.com/shogo82148/std/go/token"
)

// Packageはパッケージ全体のドキュメントです。
type Package struct {
	Doc        string
	Name       string
	ImportPath string
	Imports    []string
	Filenames  []string
	Notes      map[string][]*Note

	// 廃止予定: 後方互換性のためにBugsは引き続き使用されますが、新しいコードでは代わりにNotesを使用する必要があります。
	Bugs []string

	// 宣言
	Consts []*Value
	Types  []*Type
	Vars   []*Value
	Funcs  []*Func

	// Examplesはパッケージに関連付けられた例のソートされたリストです。
	// これらの例はNewFromFilesに提供される_test.goファイルから抽出されます。
	Examples []*Example

	importByName map[string]string
	syms         map[string]bool
}

// Valueは（おそらくグループ化された）varまたはconstの宣言のためのドキュメントです。
type Value struct {
	Doc   string
	Names []string
	Decl  *ast.GenDecl

	order int
}

// Typeは型宣言のためのドキュメントです。
type Type struct {
	Doc  string
	Name string
	Decl *ast.GenDecl

	// 関連する宣言
	Consts  []*Value
	Vars    []*Value
	Funcs   []*Func
	Methods []*Func

	// Examplesは、この型に関連付けられた例のソートされたリストです。例は、NewFromFilesに提供される_test.goファイルから抽出されます。
	Examples []*Example
}

// Funcはfunc宣言のためのドキュメンテーションです。
type Func struct {
	Doc  string
	Name string
	Decl *ast.FuncDecl

	// メソッド
	// (関数の場合、これらのフィールドはそれぞれのゼロ値を持ちます)
	Recv  string
	Orig  string
	Level int

	// Examplesはこの関数またはメソッドに関連付けられた並べ替えられた例のリストです。例は、NewFromFilesに提供される_test.goファイルから抽出されます。
	Examples []*Example
}

// Noteは"MARKER（uid）：ノートの本文"で始まるマークされたコメントを表します。
// マーカーが2つ以上の大文字[A-Z]とuidが少なくとも1つの文字で構成されるノートは認識されます。
// uidの後ろの":"はオプションです。
// ノートはPackage.Notesマップに、ノートのマーカーをインデックスとして収集されます。
type Note struct {
	Pos, End token.Pos
	UID      string
	Body     string
}

// Modeの値は、[New] と [NewFromFiles] の動作を制御します。
type Mode int

const (

	// AllDeclsは、公開されているものだけでなく、すべてのパッケージレベルの宣言のドキュメントを抽出するよう指示します。
	AllDecls Mode = 1 << iota

	// AllMethodsは、見えない（非公開）無名フィールドのみでなく、すべての埋め込まれたメソッドを表示するように指定します。
	AllMethods

	// PreserveASTは、ASTを変更せずにそのまま保持することを指定します。もともと、godocでは、関数の本体などのASTの一部がnilになってメモリを節約していましたが、すべてのプログラムがその動作を望むわけではありません。
	PreserveAST
)

// Newは指定されたパッケージASTのパッケージドキュメントを計算します。
// NewはAST pkgを所有し、編集または上書きすることができます。
// [Examples] フィールドが入力されている場合は、[NewFromFiles] を使用して
// パッケージの_test.goファイルを含めてください。
func New(pkg *ast.Package, importPath string, mode Mode) *Package

// NewFromFilesはパッケージのドキュメントを計算します。
//
// パッケージは*ast.Filesのリストと対応するファイルセットで指定されます。
// ファイルセットはnilであってはなりません。
// NewFromFilesはドキュメントを計算する際に提供されたすべてのファイルを使用しますので、
// 呼び出し側は必要なビルドコンテキストに一致するファイルのみを提供する責任があります。
// ファイルがビルドコンテキストに一致するかどうかを判断するためには、
// "go/build".Context.MatchFileを使用できます。
// この関数は、望ましいGOOSおよびGOARCHの値と他のビルド制約と一致するかどうかを判断します。
// パッケージのインポートパスはimportPathで指定されます。
//
// _test.goファイルに見つかった例は、それらの名前に基づいて対応する型、関数、メソッド、またはパッケージに関連付けられます。
// もし例の名前に接尾辞がある場合、それは [Example.Suffix] フィールドに設定されます。
// 名前が正しくない [Example] はスキップされます。
//
// オプションとして、抽出の振る舞いの低レベルな側面を制御するために [Mode] 型の単一の追加引数を指定することができます。
//
// NewFromFilesはASTファイルの所有権を持ち、それらを編集する場合があります。
// ただし、PreserveASTモードビットがオンになっている場合は、編集しません。
func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*Package, error)

// Parserは、パッケージpからドキュメントコメントを解析するために設定されたドキュメントコメントパーサーを返します。
// 各呼び出しは新しいパーサーを返すため、呼び出し元は使用前にカスタマイズすることができます。
func (p *Package) Parser() *comment.Parser

// Printerは、パッケージpからドキュメントコメントの印刷に設定されたドキュメントコメントプリンターを返します。
// 各呼び出しは、新しいプリンターを返すため、呼び出し元は使用前にカスタマイズすることができます。
func (p *Package) Printer() *comment.Printer

// HTMLは、ドキュメントコメントテキストのフォーマットされたHTMLを返します。
//
// HTMLの詳細をカスタマイズするには、[Package.Printer]を使用して[comment.Printer]を取得し、そのHTMLメソッドを呼び出す前に設定します。
func (p *Package) HTML(text string) []byte

// MarkdownはドキュメントコメントテキストのフォーマットされたMarkdownを返します。
//
// Markdownの詳細をカスタマイズするには、[Package.Printer]を使用して[comment.Printer]を取得し、
// そのMarkdownメソッドを呼び出す前に設定してください。
func (p *Package) Markdown(text string) []byte

// Textは、ドキュメントコメントのテキストを80のユニコードコードポイントに折り返し、コードブロックのインデントにはタブを使用したフォーマット済みのテキストを返します。
//
// フォーマットの詳細をカスタマイズするには、[Package.Printer]を使用して[comment.Printer]を取得し、そのTextメソッドを呼び出す前に設定してください。
func (p *Package) Text(text string) []byte
