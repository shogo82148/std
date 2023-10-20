// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ast declares the types used to represent syntax trees for Go
// packages.
package ast

import (
	"github.com/shogo82148/std/go/token"
)

// All node types implement the Node interface.
type Node interface {
	Pos() token.Pos
	End() token.Pos
}

// All expression nodes implement the Expr interface.
type Expr interface {
	Node
	exprNode()
}

// All statement nodes implement the Stmt interface.
type Stmt interface {
	Node
	stmtNode()
}

// All declaration nodes implement the Decl interface.
type Decl interface {
	Node
	declNode()
}

// A Comment node represents a single //-style or /*-style comment.
//
// The Text field contains the comment text without carriage returns (\r) that
// may have been present in the source. Because a comment's end position is
// computed using len(Text), the position reported by [Comment.End] does not match the
// true source end position for comments containing carriage returns.
type Comment struct {
	Slash token.Pos
	Text  string
}

func (c *Comment) Pos() token.Pos
func (c *Comment) End() token.Pos

// A CommentGroup represents a sequence of comments
// with no other tokens and no empty lines between.
type CommentGroup struct {
	List []*Comment
}

func (g *CommentGroup) Pos() token.Pos
func (g *CommentGroup) End() token.Pos

// Text returns the text of the comment.
// Comment markers (//, /*, and */), the first space of a line comment, and
// leading and trailing empty lines are removed.
// Comment directives like "//line" and "//go:noinline" are also removed.
// Multiple empty lines are reduced to one, and trailing space on lines is trimmed.
// Unless the result is empty, it is newline-terminated.
func (g *CommentGroup) Text() string

// A Field represents a Field declaration list in a struct type,
// a method list in an interface type, or a parameter/result declaration
// in a signature.
// [Field.Names] is nil for unnamed parameters (parameter lists which only contain types)
// and embedded struct fields. In the latter case, the field name is the type name.
type Field struct {
	Doc     *CommentGroup
	Names   []*Ident
	Type    Expr
	Tag     *BasicLit
	Comment *CommentGroup
}

func (f *Field) Pos() token.Pos

func (f *Field) End() token.Pos

// A FieldList represents a list of Fields, enclosed by parentheses,
// curly braces, or square brackets.
type FieldList struct {
	Opening token.Pos
	List    []*Field
	Closing token.Pos
}

func (f *FieldList) Pos() token.Pos

func (f *FieldList) End() token.Pos

// NumFields returns the number of parameters or struct fields represented by a [FieldList].
func (f *FieldList) NumFields() int

// An expression is represented by a tree consisting of one
// or more of the following concrete expression nodes.
type (
	// A BadExpr node is a placeholder for an expression containing
	// syntax errors for which a correct expression node cannot be
	// created.
	//
	BadExpr struct {
		From, To token.Pos
	}

	// An Ident node represents an identifier.
	Ident struct {
		NamePos token.Pos
		Name    string
		Obj     *Object
	}

	// An Ellipsis node stands for the "..." type in a
	// parameter list or the "..." length in an array type.
	//
	Ellipsis struct {
		Ellipsis token.Pos
		Elt      Expr
	}

	// A BasicLit node represents a literal of basic type.
	BasicLit struct {
		ValuePos token.Pos
		Kind     token.Token
		Value    string
	}

	// A FuncLit node represents a function literal.
	FuncLit struct {
		Type *FuncType
		Body *BlockStmt
	}

	// A CompositeLit node represents a composite literal.
	CompositeLit struct {
		Type       Expr
		Lbrace     token.Pos
		Elts       []Expr
		Rbrace     token.Pos
		Incomplete bool
	}

	// A ParenExpr node represents a parenthesized expression.
	ParenExpr struct {
		Lparen token.Pos
		X      Expr
		Rparen token.Pos
	}

	// A SelectorExpr node represents an expression followed by a selector.
	SelectorExpr struct {
		X   Expr
		Sel *Ident
	}

	// An IndexExpr node represents an expression followed by an index.
	IndexExpr struct {
		X      Expr
		Lbrack token.Pos
		Index  Expr
		Rbrack token.Pos
	}

	// An IndexListExpr node represents an expression followed by multiple
	// indices.
	IndexListExpr struct {
		X       Expr
		Lbrack  token.Pos
		Indices []Expr
		Rbrack  token.Pos
	}

	// A SliceExpr node represents an expression followed by slice indices.
	SliceExpr struct {
		X      Expr
		Lbrack token.Pos
		Low    Expr
		High   Expr
		Max    Expr
		Slice3 bool
		Rbrack token.Pos
	}

	// A TypeAssertExpr node represents an expression followed by a
	// type assertion.
	//
	TypeAssertExpr struct {
		X      Expr
		Lparen token.Pos
		Type   Expr
		Rparen token.Pos
	}

	// A CallExpr node represents an expression followed by an argument list.
	CallExpr struct {
		Fun      Expr
		Lparen   token.Pos
		Args     []Expr
		Ellipsis token.Pos
		Rparen   token.Pos
	}

	// A StarExpr node represents an expression of the form "*" Expression.
	// Semantically it could be a unary "*" expression, or a pointer type.
	//
	StarExpr struct {
		Star token.Pos
		X    Expr
	}

	// A UnaryExpr node represents a unary expression.
	// Unary "*" expressions are represented via StarExpr nodes.
	//
	UnaryExpr struct {
		OpPos token.Pos
		Op    token.Token
		X     Expr
	}

	// A BinaryExpr node represents a binary expression.
	BinaryExpr struct {
		X     Expr
		OpPos token.Pos
		Op    token.Token
		Y     Expr
	}

	// A KeyValueExpr node represents (key : value) pairs
	// in composite literals.
	//
	KeyValueExpr struct {
		Key   Expr
		Colon token.Pos
		Value Expr
	}
)

// The direction of a channel type is indicated by a bit
// mask including one or both of the following constants.
type ChanDir int

const (
	SEND ChanDir = 1 << iota
	RECV
)

// A type is represented by a tree consisting of one
// or more of the following type-specific expression
// nodes.
type (
	// An ArrayType node represents an array or slice type.
	ArrayType struct {
		Lbrack token.Pos
		Len    Expr
		Elt    Expr
	}

	// A StructType node represents a struct type.
	StructType struct {
		Struct     token.Pos
		Fields     *FieldList
		Incomplete bool
	}

	// A FuncType node represents a function type.
	FuncType struct {
		Func       token.Pos
		TypeParams *FieldList
		Params     *FieldList
		Results    *FieldList
	}

	// An InterfaceType node represents an interface type.
	InterfaceType struct {
		Interface  token.Pos
		Methods    *FieldList
		Incomplete bool
	}

	// A MapType node represents a map type.
	MapType struct {
		Map   token.Pos
		Key   Expr
		Value Expr
	}

	// A ChanType node represents a channel type.
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

// NewIdent creates a new [Ident] without position.
// Useful for ASTs generated by code other than the Go parser.
func NewIdent(name string) *Ident

// IsExported reports whether name starts with an upper-case letter.
func IsExported(name string) bool

// IsExported reports whether id starts with an upper-case letter.
func (id *Ident) IsExported() bool

func (id *Ident) String() string

// A statement is represented by a tree consisting of one
// or more of the following concrete statement nodes.
type (
	// A BadStmt node is a placeholder for statements containing
	// syntax errors for which no correct statement nodes can be
	// created.
	//
	BadStmt struct {
		From, To token.Pos
	}

	// A DeclStmt node represents a declaration in a statement list.
	DeclStmt struct {
		Decl Decl
	}

	// An EmptyStmt node represents an empty statement.
	// The "position" of the empty statement is the position
	// of the immediately following (explicit or implicit) semicolon.
	//
	EmptyStmt struct {
		Semicolon token.Pos
		Implicit  bool
	}

	// A LabeledStmt node represents a labeled statement.
	LabeledStmt struct {
		Label *Ident
		Colon token.Pos
		Stmt  Stmt
	}

	// An ExprStmt node represents a (stand-alone) expression
	// in a statement list.
	//
	ExprStmt struct {
		X Expr
	}

	// A SendStmt node represents a send statement.
	SendStmt struct {
		Chan  Expr
		Arrow token.Pos
		Value Expr
	}

	// An IncDecStmt node represents an increment or decrement statement.
	IncDecStmt struct {
		X      Expr
		TokPos token.Pos
		Tok    token.Token
	}

	// An AssignStmt node represents an assignment or
	// a short variable declaration.
	//
	AssignStmt struct {
		Lhs    []Expr
		TokPos token.Pos
		Tok    token.Token
		Rhs    []Expr
	}

	// A GoStmt node represents a go statement.
	GoStmt struct {
		Go   token.Pos
		Call *CallExpr
	}

	// A DeferStmt node represents a defer statement.
	DeferStmt struct {
		Defer token.Pos
		Call  *CallExpr
	}

	// A ReturnStmt node represents a return statement.
	ReturnStmt struct {
		Return  token.Pos
		Results []Expr
	}

	// A BranchStmt node represents a break, continue, goto,
	// or fallthrough statement.
	//
	BranchStmt struct {
		TokPos token.Pos
		Tok    token.Token
		Label  *Ident
	}

	// A BlockStmt node represents a braced statement list.
	BlockStmt struct {
		Lbrace token.Pos
		List   []Stmt
		Rbrace token.Pos
	}

	// An IfStmt node represents an if statement.
	IfStmt struct {
		If   token.Pos
		Init Stmt
		Cond Expr
		Body *BlockStmt
		Else Stmt
	}

	// A CaseClause represents a case of an expression or type switch statement.
	CaseClause struct {
		Case  token.Pos
		List  []Expr
		Colon token.Pos
		Body  []Stmt
	}

	// A SwitchStmt node represents an expression switch statement.
	SwitchStmt struct {
		Switch token.Pos
		Init   Stmt
		Tag    Expr
		Body   *BlockStmt
	}

	// A TypeSwitchStmt node represents a type switch statement.
	TypeSwitchStmt struct {
		Switch token.Pos
		Init   Stmt
		Assign Stmt
		Body   *BlockStmt
	}

	// A CommClause node represents a case of a select statement.
	CommClause struct {
		Case  token.Pos
		Comm  Stmt
		Colon token.Pos
		Body  []Stmt
	}

	// A SelectStmt node represents a select statement.
	SelectStmt struct {
		Select token.Pos
		Body   *BlockStmt
	}

	// A ForStmt represents a for statement.
	ForStmt struct {
		For  token.Pos
		Init Stmt
		Cond Expr
		Post Stmt
		Body *BlockStmt
	}

	// A RangeStmt represents a for statement with a range clause.
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

// A Spec node represents a single (non-parenthesized) import,
// constant, type, or variable declaration.
type (
	// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.
	Spec interface {
		Node
		specNode()
	}

	// An ImportSpec node represents a single package import.
	ImportSpec struct {
		Doc     *CommentGroup
		Name    *Ident
		Path    *BasicLit
		Comment *CommentGroup
		EndPos  token.Pos
	}

	// A ValueSpec node represents a constant or variable declaration
	// (ConstSpec or VarSpec production).
	//
	ValueSpec struct {
		Doc     *CommentGroup
		Names   []*Ident
		Type    Expr
		Values  []Expr
		Comment *CommentGroup
	}

	// A TypeSpec node represents a type declaration (TypeSpec production).
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

// A declaration is represented by one of the following declaration nodes.
type (
	// A BadDecl node is a placeholder for a declaration containing
	// syntax errors for which a correct declaration node cannot be
	// created.
	//
	BadDecl struct {
		From, To token.Pos
	}

	// A GenDecl node (generic declaration node) represents an import,
	// constant, type or variable declaration. A valid Lparen position
	// (Lparen.IsValid()) indicates a parenthesized declaration.
	//
	// Relationship between Tok value and Specs element type:
	//
	//	token.IMPORT  *ImportSpec
	//	token.CONST   *ValueSpec
	//	token.TYPE    *TypeSpec
	//	token.VAR     *ValueSpec
	//
	GenDecl struct {
		Doc    *CommentGroup
		TokPos token.Pos
		Tok    token.Token
		Lparen token.Pos
		Specs  []Spec
		Rparen token.Pos
	}

	// A FuncDecl node represents a function declaration.
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

// A File node represents a Go source file.
//
// The Comments list contains all comments in the source file in order of
// appearance, including the comments that are pointed to from other nodes
// via Doc and Comment fields.
//
// For correct printing of source code containing comments (using packages
// go/format and go/printer), special care must be taken to update comments
// when a File's syntax tree is modified: For printing, comments are interspersed
// between tokens based on their position. If syntax tree nodes are
// removed or moved, relevant comments in their vicinity must also be removed
// (from the [File.Comments] list) or moved accordingly (by updating their
// positions). A [CommentMap] may be used to facilitate some of these operations.
//
// Whether and how a comment is associated with a node depends on the
// interpretation of the syntax tree by the manipulating program: Except for Doc
// and [Comment] comments directly associated with nodes, the remaining comments
// are "free-floating" (see also issues #18593, #20744).
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

// Pos returns the position of the package declaration.
// (Use FileStart for the start of the entire file.)
func (f *File) Pos() token.Pos

// End returns the end of the last declaration in the file.
// (Use FileEnd for the end of the entire file.)
func (f *File) End() token.Pos

// A Package node represents a set of source files
// collectively building a Go package.
//
// Deprecated: use the type checker [go/types] instead; see [Object].
type Package struct {
	Name    string
	Scope   *Scope
	Imports map[string]*Object
	Files   map[string]*File
}

func (p *Package) Pos() token.Pos
func (p *Package) End() token.Pos

// IsGenerated reports whether the file was generated by a program,
// not handwritten, by detecting the special comment described
// at https://go.dev/s/generatedcode.
//
// The syntax tree must have been parsed with the ParseComments flag.
// Example:
//
//	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments|parser.PackageClauseOnly)
//	if err != nil { ... }
//	gen := ast.IsGenerated(f)
func IsGenerated(file *File) bool

// Unparen returns the expression with any enclosing parentheses removed.
func Unparen(e Expr) Expr
