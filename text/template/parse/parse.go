// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parse builds parse trees for templates as defined by text/template
// and html/template. Clients should use those packages to construct templates
// rather than this one, which provides shared internal data structures not
// intended for general use.
package parse

// Tree is the representation of a single parsed template.
type Tree struct {
	Name      string
	ParseName string
	Root      *ListNode
	Mode      Mode
	text      string

	funcs      []map[string]interface{}
	lex        *lexer
	token      [3]item
	peekCount  int
	vars       []string
	treeSet    map[string]*Tree
	actionLine int
	mode       Mode
}

// A mode value is a set of flags (or 0). Modes control parser behavior.
type Mode uint

const (
	ParseComments Mode = 1 << iota
	SkipFuncCheck
)

// Copy returns a copy of the Tree. Any parsing state is discarded.
func (t *Tree) Copy() *Tree

// Parse returns a map from template name to parse.Tree, created by parsing the
// templates described in the argument string. The top-level template will be
// given the specified name. If an error is encountered, parsing stops and an
// empty map is returned with the error.
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]interface{}) (map[string]*Tree, error)

// New allocates a new parse tree with the given name.
func New(name string, funcs ...map[string]interface{}) *Tree

// ErrorContext returns a textual representation of the location of the node in the input text.
// The receiver is only used when the node does not have a pointer to the tree inside,
// which can occur in old code.
func (t *Tree) ErrorContext(n Node) (location, context string)

// Parse parses the template definition string to construct a representation of
// the template for execution. If either action delimiter string is empty, the
// default ("{{" or "}}") is used. Embedded template definitions are added to
// the treeSet map.
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]interface{}) (tree *Tree, err error)

// IsEmptyTree reports whether this tree (node) is empty of everything but space or comments.
func IsEmptyTree(n Node) bool
