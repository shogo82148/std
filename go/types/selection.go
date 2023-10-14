// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは選択操作を実装しています。

package types

// SelectionKind describes the kind of a selector expression x.f
// (excluding qualified identifiers).
//
// If x is a struct or *struct, a selector expression x.f may denote a
// sequence of selection operations x.a.b.c.f. The SelectionKind
// describes the kind of the final (explicit) operation; all the
// previous (implicit) operations are always field selections.
// Each element of Indices specifies an implicit field (a, b, c)
// by its index in the struct type of the field selection operand.
//
// For a FieldVal operation, the final selection refers to the field
// specified by Selection.Obj.
//
// For a MethodVal operation, the final selection refers to a method.
// If the "pointerness" of the method's declared receiver does not
// match that of the effective receiver after implicit field
// selection, then an & or * operation is implicitly applied to the
// receiver variable or value.
// So, x.f denotes (&x.a.b.c).f when f requires a pointer receiver but
// x.a.b.c is a non-pointer variable; and it denotes (*x.a.b.c).f when
// f requires a non-pointer receiver but x.a.b.c is a pointer value.
//
// All pointer indirections, whether due to implicit or explicit field
// selections or * operations inserted for "pointerness", panic if
// applied to a nil pointer, so a method call x.f() may panic even
// before the function call.
//
// By contrast, a MethodExpr operation T.f is essentially equivalent
// to a function literal of the form:
//
//	func(x T, args) (results) { return x.f(args) }
//
// Consequently, any implicit field selections and * operations
// inserted for "pointerness" are not evaluated until the function is
// called, so a T.f or (*T).f expression never panics.
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

// Indirect reports whether any pointer indirection was required to get from
// x to f in x.f.
//
// Beware: Indirect spuriously returns true (Go issue #8353) for a
// MethodVal selection in which the receiver argument and parameter
// both have type *T so there is no indirection.
// Unfortunately, a fix is too risky.
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
