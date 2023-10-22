// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

// [Walk] によって遭遇したノードごとに Visitor の Visit メソッドが呼び出されます。
// もし結果の visitor w が nil でない場合、[Walk] はノードの各子ノードを visitor w と共に訪問し、その後に w.Visit(nil) の呼び出しを行います。
type Visitor interface {
	Visit(node Node) (w Visitor)
}

// WalkはASTを深さ優先でトラバースします。最初にv.Visit(node)を呼び出します。nodeはnilであってはいけません。v.Visit(node)から返されるビジターwがnilでない場合、Walkはnodeのnilでない子要素ごとに再帰的にビジターwを用いて呼び出され、その後にw.Visit(nil)が呼び出されます。
func Walk(v Visitor, node Node)

// InspectはASTを深さ優先順で走査します：まずf(node)を呼び出します。nodeはnilであってはなりません。fがtrueを返す場合、Inspectはnodeの非nilな子のそれぞれに対して再帰的にfを呼び出し、その後にf(nil)を呼び出します。
func Inspect(node Node, f func(Node) bool)
