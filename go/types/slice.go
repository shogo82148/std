// "go test -run=Generate -write=all" によって生成されたコードです。 編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Sliceはスライス型を表します。
type Slice struct {
	elem Type
}

// NewSliceは与えられた要素タイプ用の新しいスライスタイプを返します。
func NewSlice(elem Type) *Slice

// Elemはスライスsの要素の型を返します。
func (s *Slice) Elem() Type

func (s *Slice) Underlying() Type
func (s *Slice) String() string
