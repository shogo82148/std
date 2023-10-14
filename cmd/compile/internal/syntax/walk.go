// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements syntax tree walking.

package syntax

// Inspect traverses an AST in pre-order: it starts by calling f(root);
// root must not be nil. If f returns true, Inspect invokes f recursively
// for each of the non-nil children of root, followed by a call of f(nil).
//
// See Walk for caveats about shared nodes.
func Inspect(root Node, f func(Node) bool)

// Walk traverses an AST in pre-order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
//
// Some nodes may be shared among multiple parent nodes (e.g., types in
// field lists such as type T in "a, b, c T"). Such shared nodes are
// walked multiple times.
// TODO(gri) Revisit this design. It may make sense to walk those nodes
// only once. A place where this matters is types2.TestResolveIdents.
func Walk(root Node, v Visitor)

// A Visitor's Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(node Node) (w Visitor)
}
