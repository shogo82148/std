// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ast はGoのパッケージの構文木を表すために使用される型を宣言します。
package ast

import (
	"github.com/shogo82148/std/go/token"
)

// すべてのノードタイプはNodeインターフェースを実装します。
type Node interface {
	Pos() token.Pos
	End() token.Pos
}

// すべての式のノードは、Exprインターフェースを実装しています。
type Expr interface {
	Node
	exprNode()
}

// すべてのステートメントノードは、Stmtインターフェースを実装しています。
type Stmt interface {
	Node
	stmtNode()
}

// すべての宣言ノードはDeclインターフェースを実装します。
type Decl interface {
	Node
	declNode()
}

// Commentノードは、単一の//-スタイルまたは/*-スタイルのコメントを表します。
//
// Textフィールドには、ソースに存在した可能性のあるキャリッジリターン（\r）を含まないコメントテキストが含まれます。コメントの終了位置はlen（Text）を使用して計算されるため、End（）によって報告される位置は、キャリッジリターンを含むコメントの真のソース終了位置と一致しません。
type Comment struct {
	Slash token.Pos
	Text  string
}

func (c *Comment) Pos() token.Pos
func (c *Comment) End() token.Pos

// CommentGroupは、他のトークンや空の行がないコメントのシーケンスを表します。
type CommentGroup struct {
	List []*Comment
}

func (g *CommentGroup) Pos() token.Pos
func (g *CommentGroup) End() token.Pos

// Textはコメントのテキストを返します。
// コメントマーカー(//、/*、および*/)、行コメントの最初のスペース、および
// 先行および後続の空行は除去されます。
// "//line"や"//go:noinline"のようなコメントディレクティブも削除されます。
// 複数の空行は1つに減らされ、行の末尾のスペースはトリムされます。
// 結果が空でない場合、改行で終わります。
func (g *CommentGroup) Text() string

// Fieldは、struct型のフィールド宣言リスト、インタフェース型のメソッドリスト、またはシグネチャのパラメータ/結果の宣言を表します。
// Field.Namesは、無名のパラメータ（型のみを含むパラメータリスト）や埋め込まれたstructフィールドの場合はnilです。
// 後者の場合、フィールド名は型名です。
type Field struct {
	Doc     *CommentGroup
	Names   []*Ident
	Type    Expr
	Tag     *BasicLit
	Comment *CommentGroup
}

func (f *Field) Pos() token.Pos

func (f *Field) End() token.Pos

// FieldList は、かっこ、中かっこ、又は角かっこで囲まれたフィールドのリストを表します。
type FieldList struct {
	Opening token.Pos
	List    []*Field
	Closing token.Pos
}

func (f *FieldList) Pos() token.Pos

func (f *FieldList) End() token.Pos

// NumFieldsはFieldListによって表されるパラメータまたは構造体のフィールドの数を返します。
func (f *FieldList) NumFields() int

// 式は、以下の具体的な式ノードを1つ以上含む木で表されます。
type (

	// BadExprノードは、正しい式ノードを作成できない構文エラーを含む式のプレースホルダーです。
	BadExpr struct {
		From, To token.Pos
	}

	// Identノードは、識別子を表します。
	Ident struct {
		NamePos token.Pos
		Name    string
		Obj     *Object
	}

	// Ellipsis（省略符）ノードは、パラメータリスト内の "..." 型または配列型の "..." 長さを表します。
	Ellipsis struct {
		Ellipsis token.Pos
		Elt      Expr
	}

	// BasicLitノードは基本型のリテラルを表します。
	BasicLit struct {
		ValuePos token.Pos
		Kind     token.Token
		Value    string
	}

	// FuncLitノードは関数リテラルを表します。
	FuncLit struct {
		Type *FuncType
		Body *BlockStmt
	}

	// CompositeLitノードは複合リテラルを表します。
	CompositeLit struct {
		Type       Expr
		Lbrace     token.Pos
		Elts       []Expr
		Rbrace     token.Pos
		Incomplete bool
	}

	// ParenExprノードは、括弧で囲まれた式を表します。
	ParenExpr struct {
		Lparen token.Pos
		X      Expr
		Rparen token.Pos
	}

	// SelectorExprノードは、セレクターに続く式を表します。
	SelectorExpr struct {
		X   Expr
		Sel *Ident
	}

	// IndexExprノードは、インデックスに続く式を表します。
	IndexExpr struct {
		X      Expr
		Lbrack token.Pos
		Index  Expr
		Rbrack token.Pos
	}

	// IndexListExprノードは、複数のインデックスで続く式を表します。
	IndexListExpr struct {
		X       Expr
		Lbrack  token.Pos
		Indices []Expr
		Rbrack  token.Pos
	}

	// SliceExprノードはスライスのインデックスが続いた式を表します。
	SliceExpr struct {
		X      Expr
		Lbrack token.Pos
		Low    Expr
		High   Expr
		Max    Expr
		Slice3 bool
		Rbrack token.Pos
	}

	// TypeAssertExprノードは、式の後に型アサーションが続くことを表します。
	TypeAssertExpr struct {
		X      Expr
		Lparen token.Pos
		Type   Expr
		Rparen token.Pos
	}

	// A CallExpr node represents an expression followed by an argument list.
	// CallExprノードは、式の後に引数リストが続くことを表します。
	CallExpr struct {
		Fun      Expr
		Lparen   token.Pos
		Args     []Expr
		Ellipsis token.Pos
		Rparen   token.Pos
	}

	// StarExprノードは、"*" Expressionの形式の式を表します。
	// 意味的には、単項"*"式またはポインタータイプのいずれかになります。
	StarExpr struct {
		Star token.Pos
		X    Expr
	}

	// UnaryExprノードは単項式を表します。
	// 単項の "*" 式はStarExprノードを介して表されます。
	UnaryExpr struct {
		OpPos token.Pos
		Op    token.Token
		X     Expr
	}

	// BinaryExprノードはバイナリ式を表します。
	BinaryExpr struct {
		X     Expr
		OpPos token.Pos
		Op    token.Token
		Y     Expr
	}

	// KeyValueExprノードは、コンポジットリテラル内の(key: value)のペアを表します。
	KeyValueExpr struct {
		Key   Expr
		Colon token.Pos
		Value Expr
	}
)

type ChanDir int

const (
	SEND ChanDir = 1 << iota
	RECV
)

// 型は、次の型固有の式ノードの1つ以上からなるツリーで表されます。
type (
	// ArrayTypeノードは、配列またはスライスの型を表します。
	ArrayType struct {
		Lbrack token.Pos
		Len    Expr
		Elt    Expr
	}

	// StructTypeノードはstruct型を表します。
	StructType struct {
		Struct     token.Pos
		Fields     *FieldList
		Incomplete bool
	}

	// FuncTypeノードは関数の型を表します。
	FuncType struct {
		Func       token.Pos
		TypeParams *FieldList
		Params     *FieldList
		Results    *FieldList
	}

	// InterfaceTypeノードは、インターフェースの型を表します。
	InterfaceType struct {
		Interface  token.Pos
		Methods    *FieldList
		Incomplete bool
	}

	// MapType ノードはマップ型を表します。
	MapType struct {
		Map   token.Pos
		Key   Expr
		Value Expr
	}

	// ChanTypeノードは、チャネルの型を表します。
	ChanType struct {
		Begin token.Pos
		Arrow token.Pos
		Dir   ChanDir
		Value Expr
	}
)

func (x *BadExpr) Pos() token.Pos
func (x *Ident) Pos() token.Pos
func (x *Ellipsis) Pos() token.Pos
func (x *BasicLit) Pos() token.Pos
func (x *FuncLit) Pos() token.Pos
func (x *CompositeLit) Pos() token.Pos

func (x *ParenExpr) Pos() token.Pos
func (x *SelectorExpr) Pos() token.Pos
func (x *IndexExpr) Pos() token.Pos
func (x *IndexListExpr) Pos() token.Pos
func (x *SliceExpr) Pos() token.Pos
func (x *TypeAssertExpr) Pos() token.Pos
func (x *CallExpr) Pos() token.Pos
func (x *StarExpr) Pos() token.Pos
func (x *UnaryExpr) Pos() token.Pos
func (x *BinaryExpr) Pos() token.Pos
func (x *KeyValueExpr) Pos() token.Pos
func (x *ArrayType) Pos() token.Pos
func (x *StructType) Pos() token.Pos
func (x *FuncType) Pos() token.Pos

func (x *InterfaceType) Pos() token.Pos
func (x *MapType) Pos() token.Pos
func (x *ChanType) Pos() token.Pos

func (x *BadExpr) End() token.Pos
func (x *Ident) End() token.Pos
func (x *Ellipsis) End() token.Pos

func (x *BasicLit) End() token.Pos
func (x *FuncLit) End() token.Pos
func (x *CompositeLit) End() token.Pos
func (x *ParenExpr) End() token.Pos
func (x *SelectorExpr) End() token.Pos
func (x *IndexExpr) End() token.Pos
func (x *IndexListExpr) End() token.Pos
func (x *SliceExpr) End() token.Pos
func (x *TypeAssertExpr) End() token.Pos
func (x *CallExpr) End() token.Pos
func (x *StarExpr) End() token.Pos
func (x *UnaryExpr) End() token.Pos
func (x *BinaryExpr) End() token.Pos
func (x *KeyValueExpr) End() token.Pos
func (x *ArrayType) End() token.Pos
func (x *StructType) End() token.Pos
func (x *FuncType) End() token.Pos

func (x *InterfaceType) End() token.Pos
func (x *MapType) End() token.Pos
func (x *ChanType) End() token.Pos

// NewIdentは位置情報のない新しいIdentを作成します。
// Goパーサー以外のコードで生成されたASTに便利です。
func NewIdent(name string) *Ident

// IsExportedは、名前が大文字で始まるかどうかを報告します。
func IsExported(name string) bool

// IsExported は、id が大文字で始まるかどうかを報告します。
func (id *Ident) IsExported() bool

func (id *Ident) String() string

// 文は、以下の具象文ノードの1つ以上からなるツリーで表されます。
type (

	// BadStmtノードは、構文エラーを含むステートメントのプレースホルダーであり、
	// 正しいステートメントノードを作成することができません。
	BadStmt struct {
		From, To token.Pos
	}

	// DeclStmtノードは、文リスト内の宣言を表します。
	DeclStmt struct {
		Decl Decl
	}

	// EmptyStmtノードは、空の文を表します。
	// 空の文の「位置」は、直後の（明示的または暗黙の）セミコロンの位置です。
	EmptyStmt struct {
		Semicolon token.Pos
		Implicit  bool
	}

	// LabeledStmtノードは、ラベル付き文を表します。
	LabeledStmt struct {
		Label *Ident
		Colon token.Pos
		Stmt  Stmt
	}

	// ExprStmtノードは、文リストの中で単独での式を表します。
	ExprStmt struct {
		X Expr
	}

	// SendStmtノードは、送信文を表します。
	SendStmt struct {
		Chan  Expr
		Arrow token.Pos
		Value Expr
	}

	// IncDecStmtノードは、増分または減分文を表します。
	IncDecStmt struct {
		X      Expr
		TokPos token.Pos
		Tok    token.Token
	}

	// AssignStmt ノードは、代入または短い変数宣言を表します。
	AssignStmt struct {
		Lhs    []Expr
		TokPos token.Pos
		Tok    token.Token
		Rhs    []Expr
	}

	// GoStmtノードは、go文を表します。
	GoStmt struct {
		Go   token.Pos
		Call *CallExpr
	}

	// DeferStmtノードは、defer文を表します。
	DeferStmt struct {
		Defer token.Pos
		Call  *CallExpr
	}

	// ReturnStmtノードは、return文を表します。
	ReturnStmt struct {
		Return  token.Pos
		Results []Expr
	}

	// BranchStmtノードはbreak、continue、goto、またはfallthroughステートメントを表します。
	BranchStmt struct {
		TokPos token.Pos
		Tok    token.Token
		Label  *Ident
	}

	// BlockStmtノードは中括弧で囲まれた文リストを表します。
	BlockStmt struct {
		Lbrace token.Pos
		List   []Stmt
		Rbrace token.Pos
	}

	// IfStmtノードはif文を表します。
	IfStmt struct {
		If   token.Pos
		Init Stmt
		Cond Expr
		Body *BlockStmt
		Else Stmt
	}

	// CaseClauseは式や型switch文のケースを表します。
	CaseClause struct {
		Case  token.Pos
		List  []Expr
		Colon token.Pos
		Body  []Stmt
	}

	// SwitchStmtノードは、式を使ったスイッチ文を表します。
	SwitchStmt struct {
		Switch token.Pos
		Init   Stmt
		Tag    Expr
		Body   *BlockStmt
	}

	// TypeSwitchStmtノードは、型スイッチ文を表します。
	TypeSwitchStmt struct {
		Switch token.Pos
		Init   Stmt
		Assign Stmt
		Body   *BlockStmt
	}

	// CommClauseノードは、select文のcaseを表します。
	CommClause struct {
		Case  token.Pos
		Comm  Stmt
		Colon token.Pos
		Body  []Stmt
	}

	// SelectStmtノードは、select文を表します。
	SelectStmt struct {
		Select token.Pos
		Body   *BlockStmt
	}

	// ForStmt は for 文を表します。
	ForStmt struct {
		For  token.Pos
		Init Stmt
		Cond Expr
		Post Stmt
		Body *BlockStmt
	}

	// RangeStmtはrange節を持つfor文を表します。
	RangeStmt struct {
		For        token.Pos
		Key, Value Expr
		TokPos     token.Pos
		Tok        token.Token
		Range      token.Pos
		X          Expr
		Body       *BlockStmt
	}
)

func (s *BadStmt) Pos() token.Pos
func (s *DeclStmt) Pos() token.Pos
func (s *EmptyStmt) Pos() token.Pos
func (s *LabeledStmt) Pos() token.Pos
func (s *ExprStmt) Pos() token.Pos
func (s *SendStmt) Pos() token.Pos
func (s *IncDecStmt) Pos() token.Pos
func (s *AssignStmt) Pos() token.Pos
func (s *GoStmt) Pos() token.Pos
func (s *DeferStmt) Pos() token.Pos
func (s *ReturnStmt) Pos() token.Pos
func (s *BranchStmt) Pos() token.Pos
func (s *BlockStmt) Pos() token.Pos
func (s *IfStmt) Pos() token.Pos
func (s *CaseClause) Pos() token.Pos
func (s *SwitchStmt) Pos() token.Pos
func (s *TypeSwitchStmt) Pos() token.Pos
func (s *CommClause) Pos() token.Pos
func (s *SelectStmt) Pos() token.Pos
func (s *ForStmt) Pos() token.Pos
func (s *RangeStmt) Pos() token.Pos

func (s *BadStmt) End() token.Pos
func (s *DeclStmt) End() token.Pos
func (s *EmptyStmt) End() token.Pos

func (s *LabeledStmt) End() token.Pos
func (s *ExprStmt) End() token.Pos
func (s *SendStmt) End() token.Pos
func (s *IncDecStmt) End() token.Pos

func (s *AssignStmt) End() token.Pos
func (s *GoStmt) End() token.Pos
func (s *DeferStmt) End() token.Pos
func (s *ReturnStmt) End() token.Pos

func (s *BranchStmt) End() token.Pos

func (s *BlockStmt) End() token.Pos

func (s *IfStmt) End() token.Pos

func (s *CaseClause) End() token.Pos

func (s *SwitchStmt) End() token.Pos
func (s *TypeSwitchStmt) End() token.Pos
func (s *CommClause) End() token.Pos

func (s *SelectStmt) End() token.Pos
func (s *ForStmt) End() token.Pos
func (s *RangeStmt) End() token.Pos

// Spec ノードは、単一の（括弧で囲まれていない）import、定数、型、または変数の宣言を表します。
type (
	// Spec型は、*ImportSpec、*ValueSpec、および*TypeSpecのいずれかを表します。
	Spec interface {
		Node
		specNode()
	}

	// ImportSpecノードは1つのパッケージのインポートを表します。
	ImportSpec struct {
		Doc     *CommentGroup
		Name    *Ident
		Path    *BasicLit
		Comment *CommentGroup
		EndPos  token.Pos
	}

	// ValueSpecノードは定数または変数宣言を表します。
	// (ConstSpecまたはVarSpecプロダクション)。
	ValueSpec struct {
		Doc     *CommentGroup
		Names   []*Ident
		Type    Expr
		Values  []Expr
		Comment *CommentGroup
	}

	// TypeSpecノードは、型の宣言を表します (TypeSpecの生成)。
	TypeSpec struct {
		Doc        *CommentGroup
		Name       *Ident
		TypeParams *FieldList
		Assign     token.Pos
		Type       Expr
		Comment    *CommentGroup
	}
)

func (s *ImportSpec) Pos() token.Pos

func (s *ValueSpec) Pos() token.Pos
func (s *TypeSpec) Pos() token.Pos

func (s *ImportSpec) End() token.Pos

func (s *ValueSpec) End() token.Pos

func (s *TypeSpec) End() token.Pos

// 宣言は次の宣言ノードのいずれかによって表されます。
type (

	// BadDeclノードは、正しい宣言ノードを作成できない構文エラーを含む宣言のプレースホルダです。
	BadDecl struct {
		From, To token.Pos
	}

	// GenDeclノード（ジェネリック宣言ノード）は、import、constant、type、またはvariableの宣言を表します。有効なLparenの位置（Lparen.IsValid()）は、括弧で囲まれた宣言を示します。
	//
	// Tokの値とSpecs要素の型の関係：
	//
	//	token.IMPORT  *ImportSpec
	//	token.CONST   *ValueSpec
	//	token.TYPE    *TypeSpec
	//	token.VAR     *ValueSpec
	GenDecl struct {
		Doc    *CommentGroup
		TokPos token.Pos
		Tok    token.Token
		Lparen token.Pos
		Specs  []Spec
		Rparen token.Pos
	}

	// FuncDeclノードは関数宣言を表します。
	FuncDecl struct {
		Doc  *CommentGroup
		Recv *FieldList
		Name *Ident
		Type *FuncType
		Body *BlockStmt
	}
)

func (d *BadDecl) Pos() token.Pos
func (d *GenDecl) Pos() token.Pos
func (d *FuncDecl) Pos() token.Pos

func (d *BadDecl) End() token.Pos
func (d *GenDecl) End() token.Pos

func (d *FuncDecl) End() token.Pos

// FileノードはGoのソースファイルを表します。
//
// Commentsリストには、出現順にソースファイル内のすべてのコメントが含まれており、
// DocとCommentフィールドを介して他のノードから指し示されるコメントも含まれます。
//
// コメントを含むソースコードを正しく出力するために（パッケージgo/formatとgo/printerを使用して）特別な注意が必要です：
// コメントは、位置に基づいてトークンの間に挿入されます。構文木ノードが削除または移動される場合、
// その近くにある関連するコメントも削除（File.Commentsリストから）またはそれらの位置を更新して移動しなければなりません。
// これらの操作の一部を容易にするために、CommentMapを使用することもできます。
//
// コメントがノードとどのように関連付けられるかは、操作するプログラムによる構文木の解釈に依存します：
// DocとCommentコメント以外の残りのコメントは、「free-floating」です（#18593号、#20744号も参照）。
type File struct {
	Doc     *CommentGroup
	Package token.Pos
	Name    *Ident
	Decls   []Decl

	FileStart, FileEnd token.Pos
	Scope              *Scope
	Imports            []*ImportSpec
	Unresolved         []*Ident
	Comments           []*CommentGroup
	GoVersion          string
}

// Posはパッケージ宣言の位置を返します。
// （ファイル全体の開始位置にはFileStartを使用してください。）
func (f *File) Pos() token.Pos

// Endはファイル中の最後の宣言の終了位置を返します。
// （ファイル全体の終了位置にはFileEndを使用してください。）
func (f *File) End() token.Pos

// パッケージノードは、Goパッケージを構築するために共に使用される
// 一連のソースファイルを表します。
type Package struct {
	Name    string
	Scope   *Scope
	Imports map[string]*Object
	Files   map[string]*File
}

func (p *Package) Pos() token.Pos
func (p *Package) End() token.Pos

// IsGeneratedは、プログラムによって生成されたファイルか、手書きではないかを報告します。
// https://go.dev/s/generatedcodeに記載されている特殊コメントを検出します。
//
// 構文木は、ParseCommentsフラグを使用して解析されている必要があります。
// 例：
//
//	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments|parser.PackageClauseOnly)
//	if err != nil { ... }
//	gen := ast.IsGenerated(f)
func IsGenerated(file *File) bool
