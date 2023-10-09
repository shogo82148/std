// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Arrayは配列型を表します。
type Array struct {
	len  int64
	elem Type
}

// NewArrayは指定された要素の型と長さに対して新しい配列型を返します。
// 負の長さは未知の長さを示します。
func NewArray(elem Type, len int64) *Array

// Lenは配列aの長さを返します。
// 負の結果は、不明な長さを示します。
func (a *Array) Len() int64

// Elemは配列aの要素型を返します。
func (a *Array) Elem() Type

func (a *Array) Underlying() Type
func (a *Array) String() string
