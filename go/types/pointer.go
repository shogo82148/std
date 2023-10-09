// 「go test -run=Generate -write=all」によって生成されたコードです。編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// ポインタはポインタ型を表します。
type Pointer struct {
	base Type
}

// NewPointerは、指定された要素（ベース）タイプの新しいポインタータイプを返します。
func NewPointer(elem Type) *Pointer

// Elemは与えられたポインタpの要素の型を返します。
func (p *Pointer) Elem() Type

func (p *Pointer) Underlying() Type
func (p *Pointer) String() string
