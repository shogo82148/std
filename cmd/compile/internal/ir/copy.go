// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Node may implement the Orig and SetOrig method to
// maintain a pointer to the "unrewritten" form of a Node.
// If a Node does not implement OrigNode, it is its own Orig.
//
// Note that both SepCopy and Copy have definitions compatible
// with a Node that does not implement OrigNode: such a Node
// is its own Orig, and in that case, that's what both want to return
// anyway (SepCopy unconditionally, and Copy only when the input
// is its own Orig as well, but if the output does not implement
// OrigNode, then neither does the input, making the condition true).
type OrigNode interface {
	Node
	Orig() Node
	SetOrig(Node)
}

// Orig returns the “original” node for n.
// If n implements OrigNode, Orig returns n.Orig().
// Otherwise Orig returns n itself.
func Orig(n Node) Node

// SepCopy returns a separate shallow copy of n,
// breaking any Orig link to any other nodes.
func SepCopy(n Node) Node

// Copy returns a shallow copy of n.
// If Orig(n) == n, then Orig(Copy(n)) == the copy.
// Otherwise the Orig link is preserved as well.
//
// The specific semantics surrounding Orig are subtle but right for most uses.
// See issues #26855 and #27765 for pitfalls.
func Copy(n Node) Node

// DeepCopy returns a “deep” copy of n, with its entire structure copied
// (except for shared nodes like ONAME, ONONAME, OLITERAL, and OTYPE).
// If pos.IsKnown(), it sets the source position of newly allocated Nodes to pos.
func DeepCopy(pos src.XPos, n Node) Node

// DeepCopyList returns a list of deep copies (using DeepCopy) of the nodes in list.
func DeepCopyList(pos src.XPos, list []Node) []Node
