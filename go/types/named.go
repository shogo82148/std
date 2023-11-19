// 「go test -run = Generate -write = all」によって生成されたコードです。編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/sync"
)

// Namedは名前付き（定義された）型を表します。
type Named struct {
	check *Checker
	obj   *TypeName

	// fromRHS はこの *Named 型が派生元となる宣言の右辺値の型（サイクルの報告用）を保持します。
	// validType のみで使用されるため、同期化は必要ありません。
	fromRHS Type

	// インスタンス化された型に関する情報; それ以外はnil
	inst *instance

	mu         sync.Mutex
	state_     uint32
	underlying Type
	tparams    *TypeParamList

	// この型に宣言されたメソッド（この型のメソッドセットではない）
	// シグネチャは遅延してチェックされます。
	// インスタンス化されていない型の場合、これは完全なメソッドのリストです。インスタンス化された型の場合、
	// メソッドは最初にアクセスされたときに個別に展開されます。
	methods []*Func

	// loaderは型パラメータ、基底型、およびメソッドを遅延読み込みするために提供されるかもしれません。
	loader func(*Named) (tparams []*TypeParam, underlying Type, methods []*Func)
}

// NewNamedは、与えられた型名、基礎型、および関連するメソッドに対して新しい名前付き型を返します。
// もし与えられた型名のオブジェクトがまだ型を持っていない場合、その型は返された名前付き型に設定されます。
// 基礎型は*Namedではない必要があります。
func NewNamed(obj *TypeName, underlying Type, methods []*Func) *Named

// Objは、名前tで定義された宣言の型名を返します。インスタンス化された型の場合、これは元の型の型名と同じです。
func (t *Named) Obj() *TypeName

// Originは、指定された型tがインスタンス化されたジェネリック型を返します。
// tがインスタンス化されていない型の場合は、結果はtとなります。
func (t *Named) Origin() *Named

// TypeParams は、名前付きの型 t の型パラメーターを返します。返り値は nil です。
// (元々) ジェネリック型である場合、インスタンス化されているかどうかに関わらず、結果は非 nil です。
func (t *Named) TypeParams() *TypeParamList

// SetTypeParamsは名前付き型tの型パラメータを設定します。
// tには型引数を持っていてはいけません。
func (t *Named) SetTypeParams(tparams []*TypeParam)

// TypeArgsは、名前付きの型tをインスタンス化するために使用される型引数を返します。
func (t *Named) TypeArgs() *TypeList

// NumMethodsはtに定義された明示的なメソッドの数を返します。
func (t *Named) NumMethods() int

// Method returns the i'th method of named type t for 0 <= i < t.NumMethods().
//
// For an ordinary or instantiated type t, the receiver base type of this
// method is the named type t. For an uninstantiated generic type t, each
// method receiver is instantiated with its receiver type parameters.
func (t *Named) Method(i int) *Func

// SetUnderlyingは基本型を設定し、tを完全なものとしてマークします。
// tには型引数を持っていてはいけません。
func (t *Named) SetUnderlying(underlying Type)

// AddMethodは、メソッドリスト内にメソッドmがない場合に追加します。
// tには型引数が含まれていてはいけません。
func (t *Named) AddMethod(m *Func)

// TODO(gri) Investigate if Unalias can be moved to where underlying is set.
func (t *Named) Underlying() Type
func (t *Named) String() string
