// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/iter"
)

// VisitorのVisitメソッドは、[Walk] によって遭遇した各ノードに対して呼び出されます。
// 結果のビジターwがnilでない場合、[Walk] はノードの各子に対してビジターwで訪問し、
// その後にw.Visit(nil)を呼び出します。
type Visitor interface {
	Visit(node Node) (w Visitor)
}

// WalkはASTを深さ優先でトラバースします。最初にv.Visit(node)を呼び出します。nodeはnilであってはいけません。v.Visit(node)から返されるビジターwがnilでない場合、Walkはnodeのnilでない子要素ごとに再帰的にビジターwを用いて呼び出され、その後にw.Visit(nil)が呼び出されます。
func Walk(v Visitor, node Node)

<<<<<<< HEAD
// InspectはASTを深さ優先順で走査します：まずf(node)を呼び出します。nodeはnilであってはなりません。fがtrueを返す場合、Inspectはnodeの非nilな子のそれぞれに対して再帰的にfを呼び出し、その後にf(nil)を呼び出します。
=======
// Inspect traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of node, followed by a
// call of f(nil).
//
// In many cases it may be more convenient to use [Preorder], which
// returns an iterator over the sqeuence of nodes, or [PreorderStack],
// which (like [Inspect]) provides control over descent into subtrees,
// but additionally reports the stack of enclosing nodes.
>>>>>>> upstream/release-branch.go1.25
func Inspect(node Node, f func(Node) bool)

// Preorderは、指定されたルート以下（ルートを含む）の構文木のすべてのノードに対するイテレータを返します。
// これは深さ優先のプレオーダーで行われます。
//
<<<<<<< HEAD
// 各サブツリーの走査をより細かく制御するには、[Inspect] を使用します。
=======
// For greater control over the traversal of each subtree, use
// [Inspect] or [PreorderStack].
>>>>>>> upstream/release-branch.go1.25
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
