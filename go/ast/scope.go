// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはスコープとそれに含まれるオブジェクトを実装しています。

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// Scopeは、スコープ内で宣言された名前付き言語エンティティのセットと、
// 直接囲む（外側の）スコープへのリンクを維持します。
//
// Deprecated: 代わりに型チェッカー [go/types] を使用してください。詳細は [Object] を参照してください。
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
// Deprecated: IdentsとObjectsの関係は、型情報なしでは正しく計算できません。
// 例えば、式T{K: 0}は、Tの型によって、構造体、マップ、スライス、または配列リテラルを表す可能性があります。
// Tが構造体の場合、KはTのフィールドを参照しますが、他の型では環境内の値を参照します。
//
// 新しいプログラムは、[parser.SkipObjectResolution] パーサーフラグを設定して、
// 構文的なオブジェクト解決を無効にするべきです（これによりCPUとメモリも節約されます）。
// そして、オブジェクト解決が必要な場合は代わりに型チェッカー [go/types] を使用します。
// 詳細は、[types.Info] 構造体のDefs、Uses、およびImplicitsフィールドを参照してください。
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

// ObjKindは [Object] が表すものを説明します。
type ObjKind int

// 可能な [Object] の種類のリスト。
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
