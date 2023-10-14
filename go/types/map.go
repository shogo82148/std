// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Mapはマップ型を表します。
type Map struct {
	key, elem Type
}

// NewMapは与えられたキーと要素の型のための新しいマップを返します。
func NewMap(key, elem Type) *Map

// Keyはマップmのキーの型を返します。
func (m *Map) Key() Type

// Elemは、マップmの要素の型を返します。
func (m *Map) Elem() Type

func (t *Map) Underlying() Type
func (t *Map) String() string
