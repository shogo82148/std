// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/iter"
)

// A Visitor's Visit method is invoked for each node encountered by [Walk].
// If the result visitor w is not nil, [Walk] visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(node Node) (w Visitor)
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
func Walk(v Visitor, node Node)

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of node, followed by a
// call of f(nil).
//
// In many cases it may be more convenient to use [Preorder], which
// returns an iterator over the sequence of nodes, or [PreorderStack],
// which (like [Inspect]) provides control over descent into subtrees,
// but additionally reports the stack of enclosing nodes.
func Inspect(node Node, f func(Node) bool)

// Preorder returns an iterator over all the nodes of the syntax tree
// beneath (and including) the specified root, in depth-first
// preorder.
//
// For greater control over the traversal of each subtree, use
// [Inspect] or [PreorderStack].
func Preorder(root Node) iter.Seq[Node]

// PreorderStack traverses the tree rooted at root,
// calling f before visiting each node.
//
// Each call to f provides the current node and traversal stack,
// consisting of the original value of stack appended with all nodes
// from root to n, excluding n itself. (This design allows calls
// to PreorderStack to be nested without double counting.)
//
// If f returns false, the traversal skips over that subtree. Unlike
// [Inspect], no second call to f is made after visiting node n.
// (In practice, the second call is nearly always used only to pop the
// stack, and it is surprisingly tricky to do this correctly.)
func PreorderStack(root Node, stack []Node, f func(n Node, stack []Node) bool)
