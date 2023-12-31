// "go test -run=Generate -write=all"によって生成されたコードです。編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// TypeParamは型パラメータータイプを表します。
type TypeParam struct {
	check *Checker
	id    uint64
	obj   *TypeName
	index int
	bound Type
}

// NewTypeParamは新しいTypeParamを返します。Typeパラメータは、SetTypeParamsを呼び出してNamed型またはSignature型に設定することができます。複数の型に対してタイプパラメータを設定すると、パニックが発生します。
// 制約引数はnilにすることもできますが、後でSetConstraintを介して設定することができます。制約が非nilの場合、完全に定義されている必要があります。
func NewTypeParam(obj *TypeName, constraint Type) *TypeParam

// Obj は型パラメータ t の型名を返します。
func (t *TypeParam) Obj() *TypeName

// Indexは、タイプパラメータがまだタイプにバインドされていない場合は、そのパラメータリスト内のタイプパラメータのインデックス、または-1を返します。
func (t *TypeParam) Index() int

// Constraintはtに指定された型制約を返します。
func (t *TypeParam) Constraint() Type

// SetConstraintはtの型制約を設定します。
//
// これは、boundの基礎が完全に定義された後、NewTypeParamのユーザーによって呼び出される必要があります。
// また、他のタイプを形成する以外の方法で、タイプパラメータを使用する前に呼び出す必要があります。
// SetConstraintがレシーバを返すと、tは同時使用に安全です。
func (t *TypeParam) SetConstraint(bound Type)

func (t *TypeParam) Underlying() Type

func (t *TypeParam) String() string
