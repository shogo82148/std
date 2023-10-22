// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/go/token"
)

// インターフェースはインターフェース型を表します。
type Interface struct {
	check     *Checker
	methods   []*Func
	embeddeds []Type
	embedPos  *[]token.Pos
	implicit  bool
	complete  bool

	tset *_TypeSet
}

// NewInterfaceは与えられたメソッドと埋め込まれた型に対して新しいインターフェースを返します。
// NewInterfaceは提供されたメソッドの所有権を持ち、欠けているレシーバーを設定することでその型を変更する可能性があります。
//
// 廃止予定: 代わりに任意の埋め込まれた型を許可するNewInterfaceTypeを使用してください。
func NewInterface(methods []*Func, embeddeds []*Named) *Interface

// NewInterfaceTypeは、指定されたメソッドと埋め込み型に対して新しいインターフェースを返します。
// NewInterfaceTypeは、提供されたメソッドの所有権を受け取り、抜けているレシーバを設定することで型を変更する場合があります。
//
// レースコンディションを避けるために、インターフェースの型セットは、インターフェースの並行使用前にCompleteを明示的に呼び出すことで計算する必要があります。
func NewInterfaceType(methods []*Func, embeddeds []Type) *Interface

// MarkImplicitは、interface tを暗黙的にマークします。
// これは、明示的なinterfaceの埋め込みなしに、~TやA|Bなどの制約リテラルに対応することを意味します。
// MarkImplicitは、暗黙のinterfaceを同時使用する前に呼び出す必要があります。
func (t *Interface) MarkImplicit()

// NumExplicitMethods はインターフェース t の明示的に宣言されたメソッドの数を返します。
func (t *Interface) NumExplicitMethods() int

<<<<<<< HEAD
// ExplicitMethodは0 <= i < t.NumExplicitMethods()に対して、インターフェースtのi番目に明示的に宣言されたメソッドを返します。
// メソッドは一意のIdによって順序付けられます。
=======
// ExplicitMethod returns the i'th explicitly declared method of interface t for 0 <= i < t.NumExplicitMethods().
// The methods are ordered by their unique [Id].
>>>>>>> upstream/master
func (t *Interface) ExplicitMethod(i int) *Func

// NumEmbeddeds はインターフェース t 内の埋め込まれた型の数を返します。
func (t *Interface) NumEmbeddeds() int

<<<<<<< HEAD
// Embeddedは、0 <= i < t.NumEmbeddeds() の範囲でインターフェースtのi番目の埋め込まれた(*Named)型を返します。
// i番目の埋め込まれた型が定義済みの型でない場合、結果はnilです。
//
// 非推奨: 定義済みの(*Named)型に制限されないEmbeddedTypeを使用してください。
=======
// Embedded returns the i'th embedded defined (*[Named]) type of interface t for 0 <= i < t.NumEmbeddeds().
// The result is nil if the i'th embedded type is not a defined type.
//
// Deprecated: Use [Interface.EmbeddedType] which is not restricted to defined (*[Named]) types.
>>>>>>> upstream/master
func (t *Interface) Embedded(i int) *Named

// EmbeddedTypeは0 <= i < t.NumEmbeddeds()におけるインターフェースtのi番目の埋め込まれた型を返します。
func (t *Interface) EmbeddedType(i int) Type

// NumMethodsはインターフェースtのメソッドの合計数を返します。
func (t *Interface) NumMethods() int

// メソッドは、0 <= i < t.NumMethods() を満たす i の場合に、
// インターフェース t の i 番目のメソッドを返します。
// メソッドはそのユニークな ID によって順序づけされます。
func (t *Interface) Method(i int) *Func

// Emptyはtが空のインターフェースであるかどうかを報告します。
func (t *Interface) Empty() bool

// IsComparableはインターフェースtの型セットの各タイプが比較可能かどうかを報告します。
func (t *Interface) IsComparable() bool

// IsMethodSetは、インタフェースtがそのメソッドセットによって完全に記述されているかどうかを報告します。
func (t *Interface) IsMethodSet() bool

// IsImplicitは、インターフェースtが型セットリテラルのラッパーであるかどうかを報告します。
func (t *Interface) IsImplicit() bool

<<<<<<< HEAD
// Completeはインターフェースのタイプセットを計算します。これは、
// NewInterfaceTypeとNewInterfaceのユーザーによって呼び出される必要があります。インターフェースの埋め込まれた型が完全に定義され、他の型を形成する以外の方法でインターフェースのタイプを使用する前に呼び出す必要があります。インターフェースに重複するメソッドが含まれている場合、パニックが発生します。Completeはレシーバーを返します。
=======
// Complete computes the interface's type set. It must be called by users of
// [NewInterfaceType] and [NewInterface] after the interface's embedded types are
// fully defined and before using the interface type in any way other than to
// form other types. The interface must not contain duplicate methods or a
// panic occurs. Complete returns the receiver.
>>>>>>> upstream/master
//
// 完了済みのインターフェース型は、同時に使用することが安全です。
func (t *Interface) Complete() *Interface

func (t *Interface) Underlying() Type
func (t *Interface) String() string
