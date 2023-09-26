// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse nodes.

package parse

import (
	"github.com/shogo82148/std/strings"
)

// A Node is an element in the parse tree. The interface is trivial.
// The interface contains an unexported method so that only
// types local to this package can satisfy it.
type Node interface {
	Type() NodeType
	String() string

	Copy() Node
	Position() Pos

	tree() *Tree

	writeTo(*strings.Builder)
}

// NodeType identifies the type of a parse tree node.
type NodeType int

// Pos represents a byte position in the original input text from which
// this template was parsed.
type Pos int

func (p Pos) Position() Pos

// Type returns itself and provides an easy default implementation
// for embedding in a Node. Embedded in all non-trivial Nodes.
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
)

// ListNode holds a sequence of nodes.
type ListNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node
}

func (l *ListNode) String() string

func (l *ListNode) CopyList() *ListNode

func (l *ListNode) Copy() Node

// TextNode holds plain text.
type TextNode struct {
	NodeType
	Pos
	tr   *Tree
	Text []byte
}

func (t *TextNode) String() string

func (t *TextNode) Copy() Node

// PipeNode holds a pipeline with optional declaration
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

// ActionNode holds an action (something bounded by delimiters).
// Control actions have their own nodes; ActionNode represents simple
// ones such as field evaluations and parenthesized pipelines.
type ActionNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int
	Pipe *PipeNode
}

func (a *ActionNode) String() string

func (a *ActionNode) Copy() Node

// CommandNode holds a command (a pipeline inside an evaluating action).
type CommandNode struct {
	NodeType
	Pos
	tr   *Tree
	Args []Node
}

func (c *CommandNode) String() string

func (c *CommandNode) Copy() Node

// IdentifierNode holds an identifier.
type IdentifierNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident string
}

// NewIdentifier returns a new IdentifierNode with the given identifier name.
func NewIdentifier(ident string) *IdentifierNode

// SetPos sets the position. NewIdentifier is a public method so we can't modify its signature.
// Chained for convenience.
// TODO: fix one day?
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode

// SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature.
// Chained for convenience.
// TODO: fix one day?
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode

func (i *IdentifierNode) String() string

func (i *IdentifierNode) Copy() Node

// VariableNode holds a list of variable names, possibly with chained field
// accesses. The dollar sign is part of the (first) name.
type VariableNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string
}

func (v *VariableNode) String() string

func (v *VariableNode) Copy() Node

// DotNode holds the special identifier '.'.
type DotNode struct {
	NodeType
	Pos
	tr *Tree
}

func (d *DotNode) Type() NodeType

func (d *DotNode) String() string

func (d *DotNode) Copy() Node

// NilNode holds the special identifier 'nil' representing an untyped nil constant.
type NilNode struct {
	NodeType
	Pos
	tr *Tree
}

func (n *NilNode) Type() NodeType

func (n *NilNode) String() string

func (n *NilNode) Copy() Node

// FieldNode holds a field (identifier starting with '.').
// The names may be chained ('.x.y').
// The period is dropped from each ident.
type FieldNode struct {
	NodeType
	Pos
	tr    *Tree
	Ident []string
}

func (f *FieldNode) String() string

func (f *FieldNode) Copy() Node

// ChainNode holds a term followed by a chain of field accesses (identifier starting with '.').
// The names may be chained ('.x.y').
// The periods are dropped from each ident.
type ChainNode struct {
	NodeType
	Pos
	tr    *Tree
	Node  Node
	Field []string
}

// Add adds the named field (which should start with a period) to the end of the chain.
func (c *ChainNode) Add(field string)

func (c *ChainNode) String() string

func (c *ChainNode) Copy() Node

// BoolNode holds a boolean constant.
type BoolNode struct {
	NodeType
	Pos
	tr   *Tree
	True bool
}

func (b *BoolNode) String() string

func (b *BoolNode) Copy() Node

// NumberNode holds a number: signed or unsigned integer, float, or complex.
// The value is parsed and stored under all the types that can represent the value.
// This simulates in a small amount of code the behavior of Go's ideal constants.
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

// StringNode holds a string constant. The value has been "unquoted".
type StringNode struct {
	NodeType
	Pos
	tr     *Tree
	Quoted string
	Text   string
}

func (s *StringNode) String() string

func (s *StringNode) Copy() Node

// endNode represents an {{end}} action.
// It does not appear in the final parse tree.

// elseNode represents an {{else}} action. Does not appear in the final tree.

// BranchNode is the common representation of if, range, and with.
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

// IfNode represents an {{if}} action and its commands.
type IfNode struct {
	BranchNode
}

func (i *IfNode) Copy() Node

// RangeNode represents a {{range}} action and its commands.
type RangeNode struct {
	BranchNode
}

func (r *RangeNode) Copy() Node

// WithNode represents a {{with}} action and its commands.
type WithNode struct {
	BranchNode
}

func (w *WithNode) Copy() Node

// TemplateNode represents a {{template}} action.
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
