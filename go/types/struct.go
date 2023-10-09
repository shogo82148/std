// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Structはstruct型を表します。
type Struct struct {
	fields []*Var
	tags   []string
}

// NewStructは、指定されたフィールドと対応するフィールドタグを持つ新しい構造体を返します。
// インデックスiを持つフィールドにタグがある場合、tags [i]はそのタグである必要がありますが、
// tagsの長さは、最大インデックスiのタグを保持するために必要なだけの長さである場合があります。
// したがって、フィールドにタグがない場合、tagsはnilである場合があります。
func NewStruct(fields []*Var, tags []string) *Struct

// NumFieldsは、struct内のフィールドの数を返します（空白や埋め込まれたフィールドを含む）。
func (s *Struct) NumFields() int

// Fieldは0 <= i < NumFields()という条件でi番目のフィールドを返します。
func (s *Struct) Field(i int) *Var

// Tagは0 <= i < NumFields()に対するi番目のフィールドタグを返します。
func (s *Struct) Tag(i int) string

func (t *Struct) Underlying() Type
func (t *Struct) String() string
