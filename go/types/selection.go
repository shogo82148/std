// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは選択操作を実装しています。

package types

// SelectionKindは、セレクタ式x.fの種類を記述します
//（修飾子を含まない）。
type SelectionKind int

const (
	FieldVal SelectionKind = iota
	MethodVal
	MethodExpr
)

// Selection（セレクション）は、セレクタ式 x.f を表します。
// 次の宣言に対して：
//
//	type T struct{ x int; E }
//	type E struct{}
//	func (e E) m() {}
//	var p *T
//
// 以下の関係が存在します：
//
//	Selector    Kind          Recv    Obj    Type       Index     Indirect
//
//	p.x         FieldVal      T       x      int        {0}       true
//	p.m         MethodVal     *T      m      func()     {1, 0}    true
//	T.m         MethodExpr    T       m      func(T)    {1, 0}    false
type Selection struct {
	kind     SelectionKind
	recv     Type
	obj      Object
	index    []int
	indirect bool
}

// Kindは選択の種類を返します。
func (s *Selection) Kind() SelectionKind

// Recvはx.fの型を返します。
func (s *Selection) Recv() Type

// Obj は x.f によって指定されたオブジェクトを返します。フィールドの選択には *Var を、それ以外の場合は *Func を返します。
func (s *Selection) Obj() Object

// Typeはx.fの型を返しますが、fの型とは異なる場合があります。
// 詳細はSelectionを参照してください。
func (s *Selection) Type() Type

// Index はxからfへのパスを記述します。
// 最後のindexエントリは、fを宣言している型のフィールドまたはメソッドのindexです。
// 以下のいずれかです:
//
//  1. 名前付き型の宣言されたメソッドのリスト
//  2. インターフェース型のメソッドのリスト
//  3. 構造体型のフィールドのリスト
//
// より早いindexエントリは、埋め込まれたフィールドのindexであり、
// xからfに移動するために（xの型から）暗黙的にトラバースされる埋め込みの深さ0から始まります。
func (s *Selection) Index() []int

// Indirectは、xからfに移動する際にポインターの間接参照が必要だったかどうかを報告します。
func (s *Selection) Indirect() bool

func (s *Selection) String() string

// SelectionStringはsの文字列形式を返します。
// Qualifierはパッケージレベルのオブジェクトの出力を制御し、nilである場合もあります。
//
// 例：
//
//	"field (T) f int"
//	"method (T) f(X) Y"
//	"method expr (T) f(X) Y"
func SelectionString(s *Selection, qf Qualifier) string
