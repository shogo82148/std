// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Tupleは変数の順序付きリストを表します。nil *Tupleは有効な（空の）タプルです。
// Tupleはシグネチャの構成要素や複数の代入の型を表すために使用されますが、Goのファーストクラスの型ではありません。
type Tuple struct {
	vars []*Var
}

// NewTupleは指定された変数に対して新しいタプルを返します。
func NewTuple(x ...*Var) *Tuple

// Lenはタプルtの変数の数を返します。
func (t *Tuple) Len() int

// Atはタプルtのi番目の変数を返します。
func (t *Tuple) At(i int) *Var

func (t *Tuple) Underlying() Type
func (t *Tuple) String() string
