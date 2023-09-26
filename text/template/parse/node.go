// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse nodes.

package parse

// A node is an element in the parse tree. The interface is trivial.
type Node interface {
	Type() NodeType
	String() string

	Copy() Node
}

// NodeType identifies the type of a parse tree node.
type NodeType int

// Type returns itself and provides an easy default implementation
// for embedding in a Node. Embedded in all non-trivial Nodes.
func (t NodeType) Type() NodeType

const (
	NodeText    NodeType = iota
	NodeAction
	NodeBool
	NodeCommand
	NodeDot

	NodeField
	NodeIdentifier
	NodeIf
	NodeList
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
	Nodes []Node
}

func (l *ListNode) String() string

func (l *ListNode) CopyList() *ListNode

func (l *ListNode) Copy() Node

// TextNode holds plain text.
type TextNode struct {
	NodeType
	Text []byte
}

func (t *TextNode) String() string

func (t *TextNode) Copy() Node

// PipeNode holds a pipeline with optional declaration
type PipeNode struct {
	NodeType
	Line int
	Decl []*VariableNode
	Cmds []*CommandNode
}

func (p *PipeNode) String() string

func (p *PipeNode) CopyPipe() *PipeNode

func (p *PipeNode) Copy() Node

// ActionNode holds an action (something bounded by delimiters).
// Control actions have their own nodes; ActionNode represents simple
// ones such as field evaluations.
type ActionNode struct {
	NodeType
	Line int
	Pipe *PipeNode
}

func (a *ActionNode) String() string

func (a *ActionNode) Copy() Node

// CommandNode holds a command (a pipeline inside an evaluating action).
type CommandNode struct {
	NodeType
	Args []Node
}

func (c *CommandNode) String() string

func (c *CommandNode) Copy() Node

// IdentifierNode holds an identifier.
type IdentifierNode struct {
	NodeType
	Ident string
}

// NewIdentifier returns a new IdentifierNode with the given identifier name.
func NewIdentifier(ident string) *IdentifierNode

func (i *IdentifierNode) String() string

func (i *IdentifierNode) Copy() Node

// VariableNode holds a list of variable names. The dollar sign is
// part of the name.
type VariableNode struct {
	NodeType
	Ident []string
}

func (v *VariableNode) String() string

func (v *VariableNode) Copy() Node

// DotNode holds the special identifier '.'. It is represented by a nil pointer.
type DotNode bool

func (d *DotNode) Type() NodeType

func (d *DotNode) String() string

func (d *DotNode) Copy() Node

// FieldNode holds a field (identifier starting with '.').
// The names may be chained ('.x.y').
// The period is dropped from each ident.
type FieldNode struct {
	NodeType
	Ident []string
}

func (f *FieldNode) String() string

func (f *FieldNode) Copy() Node

// BoolNode holds a boolean constant.
type BoolNode struct {
	NodeType
	True bool
}

func (b *BoolNode) String() string

func (b *BoolNode) Copy() Node

// NumberNode holds a number: signed or unsigned integer, float, or complex.
// The value is parsed and stored under all the types that can represent the value.
// This simulates in a small amount of code the behavior of Go's ideal constants.
type NumberNode struct {
	NodeType
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
	Quoted string
	Text   string
}

func (s *StringNode) String() string

func (s *StringNode) Copy() Node

// endNode represents an {{end}} action. It is represented by a nil pointer.
// It does not appear in the final parse tree.

// elseNode represents an {{else}} action. Does not appear in the final tree.

// BranchNode is the common representation of if, range, and with.
type BranchNode struct {
	NodeType
	Line     int
	Pipe     *PipeNode
	List     *ListNode
	ElseList *ListNode
}

func (b *BranchNode) String() string

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
	Line int
	Name string
	Pipe *PipeNode
}

func (t *TemplateNode) String() string

func (t *TemplateNode) Copy() Node
