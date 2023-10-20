// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// A CommentMap maps an AST node to a list of comment groups
// associated with it. See [NewCommentMap] for a description of
// the association.
type CommentMap map[Node][]*CommentGroup

// NewCommentMap creates a new comment map by associating comment groups
// of the comments list with the nodes of the AST specified by node.
//
// A comment group g is associated with a node n if:
//
//   - g starts on the same line as n ends
//   - g starts on the line immediately following n, and there is
//     at least one empty line after g and before the next node
//   - g starts before n and is not associated to the node before n
//     via the previous rules
//
// NewCommentMap tries to associate a comment group to the "largest"
// node possible: For instance, if the comment is a line comment
// trailing an assignment, the comment is associated with the entire
// assignment rather than just the last operand in the assignment.
func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap

// Update replaces an old node in the comment map with the new node
// and returns the new node. Comments that were associated with the
// old node are associated with the new node.
func (cmap CommentMap) Update(old, new Node) Node

// Filter returns a new comment map consisting of only those
// entries of cmap for which a corresponding node exists in
// the AST specified by node.
func (cmap CommentMap) Filter(node Node) CommentMap

// Comments returns the list of comment groups in the comment map.
// The result is sorted in source order.
func (cmap CommentMap) Comments() []*CommentGroup

func (cmap CommentMap) String() string
