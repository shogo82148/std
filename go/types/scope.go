// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはスコープを実装しています。

package types

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// スコープはオブジェクトのセットを保持し、それが含まれる（親）スコープと含まれる（子）スコープへのリンクを維持します。オブジェクトは名前で挿入および検索することができます。Scopeのゼロ値は使用可能な空のスコープです。
type Scope struct {
	parent   *Scope
	children []*Scope
	number   int
	elems    map[string]Object
	pos, end token.Pos
	comment  string
	isFunc   bool
}

// NewScopeは、指定された親スコープに含まれる新しい空のスコープを返します（存在する場合）。コメントはデバッグ用です。
func NewScope(parent *Scope, pos, end token.Pos, comment string) *Scope

// Parentはスコープの含まれる（親）スコープを返します。
func (s *Scope) Parent() *Scope

// Lenはスコープ要素の数を返します。
func (s *Scope) Len() int

// Namesはスコープ内の要素名をソートされた順序で返します。
func (s *Scope) Names() []string

// NumChildrenはsにネストされたスコープの数を返します。
func (s *Scope) NumChildren() int

// Childは0 <= i < NumChildren()の範囲でi番目の子スコープを返します。
func (s *Scope) Child(i int) *Scope

// Lookupは、名前が与えられたスコープs内のオブジェクトを返します。
// オブジェクトが存在する場合はそのオブジェクトを返し、存在しない場合はnilを返します。
func (s *Scope) Lookup(name string) Object

// LookupParentは、sから始まるスコープの親チェーンをたどり、Lookup（name）が非nilのオブジェクトを返すスコープを見つけ、そのスコープとオブジェクトを返します。有効な位置posが指定された場合、posで宣言されたかそれ以前のオブジェクトのみが考慮されます。このようなスコープとオブジェクトが存在しない場合、結果は（nil、nil）です。
// obj.Parent()は、オブジェクトがスコープに挿入され、すでにその時点で親があった場合に、返されたスコープと異なるかもしれないことに注意してください（Insertを参照）。これは、スコープがそれらをエクスポートしたパッケージのスコープであるドットインポートされたオブジェクトにのみ起こり得ます。
func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)

// Insertはオブジェクトobjをスコープsに挿入します。
// もしsが同じ名前の代替オブジェクトaltを既に含んでいる場合、Insertはsを変更せずにaltを返します。
// そうでなければ、objを挿入し、オブジェクトの親スコープを設定し、nilを返します。
func (s *Scope) Insert(obj Object) Object

// Pos and End describe the scope's source code extent [pos, end).
// The results are guaranteed to be valid only if the type-checked
// AST has complete position information. The extent is undefined
// for Universe and package scopes.
func (s *Scope) Pos() token.Pos
func (s *Scope) End() token.Pos

// Containsは、posがスコープの範囲内にあるかどうかを報告します。
// 結果は、型チェック済みのASTに完全な位置情報がある場合にのみ有効です。
func (s *Scope) Contains(pos token.Pos) bool

// Innermostは、posを含む最も内側（子）のスコープを返します。
// posがどのスコープにも含まれていない場合、結果はnilになります。
// Universeスコープの場合も結果はnilです。
// 結果は、型チェック済みのASTに完全な位置情報がある場合にのみ有効です。
func (s *Scope) Innermost(pos token.Pos) *Scope

// WriteToは、スコープの文字列表現をwに書き込みます。
// スコープ要素は名前でソートされます。
// インデントのレベルはn >= 0で制御され、
// インデントなしの場合はn == 0です。
// recurseが設定されている場合は、ネストされた（子）スコープも書き込みます。
func (s *Scope) WriteTo(w io.Writer, n int, recurse bool)

// Stringはデバッグ用のスコープの文字列表現を返します。
func (s *Scope) String() string
