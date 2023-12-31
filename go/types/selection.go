// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは選択操作を実装しています。

package types

// SelectionKindは、セレクタ式x.fの種類を記述します
// （修飾識別子は除く）。
//
// xがstructまたは*structの場合、セレクタ式x.fは
// 一連の選択操作x.a.b.c.fを表す可能性があります。SelectionKindは
// 最終的な（明示的な）操作の種類を記述します。すべての
// 以前の（暗黙的な）操作は常にフィールド選択です。
// Indicesの各要素は、暗黙的なフィールド（a、b、c）を
// フィールド選択オペランドのstruct型のインデックスで指定します。
//
// FieldVal操作の場合、最終的な選択はSelection.Objで指定されたフィールドを参照します。
//
// MethodVal操作の場合、最終的な選択はメソッドを参照します。
// メソッドの宣言されたレシーバの"ポインタ性"が、暗黙のフィールド
// 選択後の実効レシーバと一致しない場合、&または*操作が暗黙的に
// レシーバ変数または値に適用されます。
// したがって、fがポインタレシーバを必要とするがx.a.b.cが非ポインタ変数である場合、
// x.fは(&x.a.b.c).fを表します。また、fが非ポインタレシーバを必要とするが
// x.a.b.cがポインタ値である場合、x.fは(*x.a.b.c).fを表します。
//
// 暗黙的または明示的なフィールド選択、または"ポインタ性"のために挿入された*操作による
// すべてのポインタ間接参照は、nilポインタに適用されるとパニックを引き起こします。
// したがって、メソッド呼び出しx.f()は、関数呼び出しの前にパニックを引き起こす可能性があります。
//
// 対照的に、MethodExpr操作T.fは基本的に以下の形式の関数リテラルと等価です：
//
//	func(x T, args) (results) { return x.f(args) }
//
// その結果、"ポインタ性"のために挿入された任意の暗黙的なフィールド選択と*操作は、
// 関数が呼び出されるまで評価されません。したがって、T.fまたは(*T).fの式は決してパニックしません。
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

// Indirectは、x.fでxからfへ移動するためにポインタ間接参照が必要だったかどうかを報告します。
//
// 注意：レシーバ引数とパラメータの両方が型*Tを持つMethodVal選択で、
// 間接参照がないにもかかわらず、Indirectが誤ってtrueを返す（Go issue #8353）ことがあります。
// 残念ながら、修正はリスクが高すぎます。
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
