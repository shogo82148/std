// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはメソッドセットを実装します。

package types

// MethodSetは具体的または抽象（インターフェース）メソッドの順序付けられたセットです。
// メソッドはMethodVal選択であり、m.Obj().Id()によって昇順に並べられます。
// MethodSetのゼロ値は使用準備完了の空のメソッドセットです。
type MethodSet struct {
	list []*Selection
}

func (s *MethodSet) String() string

// Lenはsのメソッドの数を返します。
func (s *MethodSet) Len() int

// Atは、0 <= i < s.Len()に対してsのi番目のメソッドを返します。
func (s *MethodSet) At(i int) *Selection

// Lookupはパッケージと名前が一致するメソッドを返します。見つからない場合はnilを返します。
func (s *MethodSet) Lookup(pkg *Package, name string) *Selection

// NewMethodSetは、指定された型Tのメソッドセットを返します。
// 空であっても、必ず非nilのメソッドセットを返します。
func NewMethodSet(T Type) *MethodSet
