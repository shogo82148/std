// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

<<<<<<< HEAD
// CommentMapはASTノードをそのノードに関連付けられたコメントグループのリストにマップします。
// 関連付けについては、NewCommentMapの説明を参照してください。
=======
// A CommentMap maps an AST node to a list of comment groups
// associated with it. See [NewCommentMap] for a description of
// the association.
>>>>>>> upstream/master
type CommentMap map[Node][]*CommentGroup

// NewCommentMapは、コメントリストのコメントグループをASTのノードと関連付けて新しいコメントマップを作成します。
// コメントグループgは、ノードnと関連付けられます。以下の条件を満たす場合です：
//   - gは、nの終了する行と同じ行で開始します。
//   - gは、nの直後の行で始まり、gと次のノードの間に少なくとも1つの空行がある場合。
//   - gは、nよりも前に開始され、前のルールを介してnの前のノードに関連付けられていない場合。
//
// NewCommentMapは、コメントグループを「最大の」ノードに関連付けようとします。たとえば、コメントが代入文の後に続く行コメントの場合、コメントは代入文全体ではなく、代入文の最後のオペランドに関連づけられます。
func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap

// Updateはコメントマップ内の古いノードを新しいノードで置き換え、新しいノードを返します。
// 古いノードに関連付けられていたコメントは、新しいノードに関連付けられます。
func (cmap CommentMap) Update(old, new Node) Node

// Filterはnodeによって指定されたASTに対応するノードが存在する場合、cmapのエントリのみで構成される新しいコメントマップを返します。
func (cmap CommentMap) Filter(node Node) CommentMap

// Commentsはコメントマップ内のコメントグループのリストを返します。
// 結果はソースの順にソートされます。
func (cmap CommentMap) Comments() []*CommentGroup

func (cmap CommentMap) String() string
