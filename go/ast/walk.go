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

// InspectはASTを深さ優先の順序でトラバースします：最初にf(node)を呼び出すことから始まります；
// nodeはnilであってはいけません。fがtrueを返す場合、Inspectはnodeのnilでない子要素それぞれに対して
// 再帰的にfを呼び出し、その後にf(nil)を呼び出します。
//
// 多くの場合、ノードの列に対するイテレータを返す [Preorder]、または
// （ [Inspect] のように）サブツリーへの降下の制御を提供しつつ、さらに囲むノードの
// スタックを報告する [PreorderStack] を使用する方が便利な場合があります。
func Inspect(node Node, f func(Node) bool)

// Preorderは、指定されたルート以下（ルートを含む）の構文木のすべてのノードに対するイテレータを返します。
// これは深さ優先のプレオーダーで行われます。
//
// 各サブツリーのトラバースをより詳細に制御するには、
// [Inspect] または [PreorderStack] を使用してください。
func Preorder(root Node) iter.Seq[Node]

// PreorderStackは、rootをルートとするツリーをトラバースし、
// 各ノードを訪問する前にfを呼び出します。
//
// fの各呼び出しでは、現在のノードとトラバースのスタックが提供されます。
// スタックは、stackの元の値に、rootからnまでのすべてのノード（n自体は除く）を
// 追加したものです。（この設計により、PreorderStackの呼び出しを
// 二重カウントなしでネストできます。）
//
// fがfalseを返した場合、トラバースはそのサブツリーをスキップします。
// [Inspect] とは異なり、ノードnを訪問した後にfへの2回目の呼び出しは行われません。
// （実際には、2回目の呼び出しはほぼ常にスタックをポップするためだけに使用され、
// これを正しく行うのは驚くほど難しいです。）
func PreorderStack(root Node, stack []Node, f func(n Node, stack []Node) bool)
