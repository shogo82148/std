// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse nodes.

package parse

import (
	"github.com/shogo82148/std/strings"
)

// Nodeはパースツリーの要素です。インターフェースは極めて単純です。
// インターフェースには未エクスポートのメソッドが含まれているため、
// このパッケージのローカルタイプのみがそれを満たすことができます。
type Node interface {
	Type() NodeType
	String() string
	// Copyは、Nodeとそのすべてのコンポーネントの深いコピーを行います。
	// 型アサーションを避けるため、一部のXxxNodesには、*XxxNodeを返す
	// 専用のCopyXxxメソッドもあります。
	Copy() Node
	Position() Pos
	// tree returns the containing *Tree.
	// It is unexported so all implementations of Node are in this package.
	tree() *Tree
	// writeTo writes the String output to the builder.
	writeTo(*strings.Builder)
}

// NodeTypeは、パースツリーノードのタイプを識別します。
type NodeType int

// Posは、このテンプレートがパースされた元の入力テキストのバイト位置を表します。
type Pos int

func (p Pos) Position() Pos

// Typeは自身を返し、Nodeに埋め込むための簡単なデフォルト実装を提供します。
// すべての非自明なノードに埋め込まれています。
func (t NodeType) Type() NodeType

const (
	NodeText NodeType = iota
	NodeAction
	NodeBool
	NodeChain
	NodeCommand
	NodeDot

	NodeField
	NodeIdentifier
	NodeIf
	NodeList
	NodeNil
	NodeNumber
	NodePipe
	NodeRange
	NodeString
	NodeTemplate
	NodeVariable
	NodeWith
	NodeComment
	NodeBreak
	NodeContinue
)

// ListNodeは、ノードのシーケンスを保持します。
type ListNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node
}

func (l *ListNode) String() string

func (l *ListNode) CopyList() *ListNode

func (l *ListNode) Copy() Node

// TextNodeはプレーンテキストを保持します。
type TextNode struct {
	NodeType
	Pos
	tr   *Tree
	Text []byte
}

func (t *TextNode) String() string

func (t *TextNode) Copy() Node

// CommentNode holds a comment.
type CommentNode struct {
	NodeType
	Pos
	tr   *Tree
	Text string
}

func (c *CommentNode) String() string

func (c *CommentNode) Copy() Node

// PipeNodeは、オプションの宣言を持つパイプラインを保持します。
type PipeNode struct {
	NodeType
	Pos
	tr       *Tree
	Line     int
	IsAssign bool
	Decl     []*VariableNode
	Cmds     []*CommandNode
}

func (p *PipeNode) String() string

func (p *PipeNode) CopyPipe() *PipeNode

func (p *PipeNode) Copy() Node

// ActionNodeはアクション（デリミタで区切られた何か）を保持します。
// 制御アクションはそれぞれが独自のノードを持ち、ActionNodeはフィールド評価や
// 括弧付きパイプラインのような単純なものを表します。
type ActionNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int
	Pipe *PipeNode
}

func (a *ActionNode) String() string

func (a *ActionNode) Copy() Node

// CommandNodeは、コマンド（評価アクション内のパイプライン）を保持します。
type CommandNode struct {
	NodeType
	Pos
	tr   *Tree
	Args []Node
}

func (c *CommandNode) String() string

func (c *CommandNode) Copy() Node

// CommandNodeは、コマンド（評価アクション内のパイプライン）を保持します。
type IdentifierNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident string
}

// NewIdentifierは、指定された識別子名を持つ新しいIdentifierNodeを返します。
func NewIdentifier(ident string) *IdentifierNode

// SetPosは位置を設定します。NewIdentifierは公開メソッドなので、そのシグネチャを変更することはできません。
// 便宜上チェーン化されています。
// TODO: いつか修正する？
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode

// SetTreeは、ノードの親ツリーを設定します。NewIdentifierは公開メソッドなので、そのシグネチャを変更することはできません。
// 便宜上チェーン化されています。
// TODO: いつか修正する？
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode

func (i *IdentifierNode) String() string

func (i *IdentifierNode) Copy() Node

// VariableNodeは、チェーンフィールドへのアクセスが可能な変数名のリストを保持します。
// ドル記号は（最初の）名前の一部です。
type VariableNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string
}

func (v *VariableNode) String() string

func (v *VariableNode) Copy() Node

// DotNodeは、特別な識別子'.'を保持します。
type DotNode struct {
	NodeType
	Pos
	tr *Tree
}

func (d *DotNode) Type() NodeType

func (d *DotNode) String() string

func (d *DotNode) Copy() Node

// NilNodeは、型指定されていないnil定数を表す特別な識別子'nil'を保持します。
type NilNode struct {
	NodeType
	Pos
	tr *Tree
}

func (n *NilNode) Type() NodeType

func (n *NilNode) String() string

func (n *NilNode) Copy() Node

// FieldNodeはフィールド（'.'で始まる識別子）を保持します。
// 名前はチェーン可能です（'.x.y'など）。
// 各識別子からピリオドは削除されます。
type FieldNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string
}

func (f *FieldNode) String() string

func (f *FieldNode) Copy() Node

// ChainNodeは、フィールドアクセスのチェーン（'.'で始まる識別子）に続く項を保持します。
// 名前はチェーン可能です（'.x.y'など）。
// 各識別子からピリオドは削除されます。
type ChainNode struct {
	NodeType
	Pos
	tr    *Tree
	Node  Node
	Field []string
}

// Addは、名前付きフィールド（ピリオドで始まるべき）をチェーンの末尾に追加します。
func (c *ChainNode) Add(field string)

func (c *ChainNode) String() string

func (c *ChainNode) Copy() Node

// BoolNodeは、ブール型の定数を保持します。
type BoolNode struct {
	NodeType
	Pos
	tr   *Tree
	True bool
}

func (b *BoolNode) String() string

func (b *BoolNode) Copy() Node

// NumberNodeは、符号付きまたは符号なしの整数、浮動小数点数、または複素数を保持します。
// 値は解析され、その値を表現できるすべての型の下に格納されます。
// これはGoの理想的な定数の振る舞いを少量のコードでシミュレートします。
type NumberNode struct {
	NodeType
	Pos
	tr         *Tree
	IsInt      bool
	IsUint     bool
	IsFloat    bool
	IsComplex  bool
	Int64      int64
	Uint64     uint64
	Float64    float64
	Complex128 complex128
	Text       string
}

func (n *NumberNode) String() string

func (n *NumberNode) Copy() Node

// StringNodeは文字列定数を保持します。値は"引用符を外され"ています。
type StringNode struct {
	NodeType
	Pos
	tr     *Tree
	Quoted string
	Text   string
}

func (s *StringNode) String() string

func (s *StringNode) Copy() Node

// BranchNodeは、if、range、およびwithの共通の表現です。
type BranchNode struct {
	NodeType
	Pos
	tr       *Tree
	Line     int
	Pipe     *PipeNode
	List     *ListNode
	ElseList *ListNode
}

func (b *BranchNode) String() string

func (b *BranchNode) Copy() Node

// IfNodeは{{if}}アクションとそのコマンドを表します。
type IfNode struct {
	BranchNode
}

func (i *IfNode) Copy() Node

// BreakNodeは{{break}}アクションを表します。
type BreakNode struct {
	tr *Tree
	NodeType
	Pos
	Line int
}

func (b *BreakNode) Copy() Node
func (b *BreakNode) String() string

// ContinueNodeは{{continue}}アクションを表します。
type ContinueNode struct {
	tr *Tree
	NodeType
	Pos
	Line int
}

func (c *ContinueNode) Copy() Node
func (c *ContinueNode) String() string

// RangeNodeは{{range}}アクションとそのコマンドを表します。
type RangeNode struct {
	BranchNode
}

func (r *RangeNode) Copy() Node

// WithNodeは{{with}}アクションとそのコマンドを表します。
type WithNode struct {
	BranchNode
}

func (w *WithNode) Copy() Node

// TemplateNodeは{{template}}アクションを表します。
type TemplateNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int
	Name string
	Pipe *PipeNode
}

func (t *TemplateNode) String() string

func (t *TemplateNode) Copy() Node
