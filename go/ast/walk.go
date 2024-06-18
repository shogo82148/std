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

// InspectはASTを深さ優先順で走査します：まずf(node)を呼び出します。nodeはnilであってはなりません。fがtrueを返す場合、Inspectはnodeの非nilな子のそれぞれに対して再帰的にfを呼び出し、その後にf(nil)を呼び出します。
func Inspect(node Node, f func(Node) bool)

// Preorderは、指定されたルート以下（ルートを含む）の構文木のすべてのノードに対するイテレータを返します。
// これは深さ優先のプレオーダーで行われます。
//
// 各サブツリーの走査をより細かく制御するには、[Inspect] を使用します。
func Preorder(root Node) iter.Seq[Node]
