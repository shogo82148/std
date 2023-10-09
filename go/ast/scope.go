// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはスコープとそれに含まれるオブジェクトを実装しています。

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// A Scope maintains the set of named language entities declared
// in the scope and a link to the immediately surrounding (outer)
// scope.
//
// Deprecated: use the type checker [go/types] instead; see [Object].
type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}

// NewScopeは外部スコープにネストされた新しいスコープを作成します。
func NewScope(outer *Scope) *Scope

// Lookupは、与えられた名前のオブジェクトがスコープsに存在すればそのオブジェクトを返します。見つからない場合はnilを返します。外部のスコープは無視されます。
func (s *Scope) Lookup(name string) *Object

// Insertは名前付きオブジェクトobjをスコープsに挿入しようとします。
// もしスコープに同じ名前のオブジェクトaltが既に存在する場合、
// Insertはスコープを変更せずにaltを返します。そうでなければ、
// objを挿入し、nilを返します。
func (s *Scope) Insert(obj *Object) (alt *Object)

// デバッグサポート
func (s *Scope) String() string

// オブジェクトは、パッケージ、定数、型、変数、関数（メソッドを含む）、またはラベルなど、名前付きの言語エンティティを表します。
//
// データフィールドには、オブジェクト固有のデータが含まれます：
//
//	Kind    Data type         Data value
//	Pkg     *Scope            package scope
//	Con     int               iota for the respective declaration
//
// Deprecated: The relationship between Idents and Objects cannot be
// correctly computed without type information. For example, the
// expression T{K: 0} may denote a struct, map, slice, or array
// literal, depending on the type of T. If T is a struct, then K
// refers to a field of T, whereas for the other types it refers to a
// value in the environment.
//
// New programs should set the [parser.SkipObjectResolution] parser
// flag to disable syntactic object resolution (which also saves CPU
// and memory), and instead use the type checker [go/types] if object
// resolution is desired. See the Defs, Uses, and Implicits fields of
// the [types.Info] struct for details.
type Object struct {
	Kind ObjKind
	Name string
	Decl any
	Data any
	Type any
}

// NewObjは指定された種類と名前の新しいオブジェクトを作成します。
func NewObj(kind ObjKind, name string) *Object

// Posはオブジェクト名の宣言のソース位置を計算します。
// 結果は計算できない場合は無効な位置になる可能性があります
// (obj.Declがnilであるか、正しくないかもしれません)。
func (obj *Object) Pos() token.Pos

// ObjKindはオブジェクトが表すものを説明します。
type ObjKind int

// 可能なオブジェクトの種類のリスト。
const (
	Bad ObjKind = iota
	Pkg
	Con
	Typ
	Var
	Fun
	Lbl
)

func (kind ObjKind) String() string
