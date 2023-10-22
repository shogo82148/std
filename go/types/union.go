// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Unionは、インタフェースに埋め込まれた項のユニオンを表します。
type Union struct {
	terms []*Term
}

<<<<<<< HEAD
// NewUnionは与えられた要素で新しいUnion型を返します。
// 空のUnionを作成することはエラーであり、構文的に不可能です。
=======
// NewUnion returns a new [Union] type with the given terms.
// It is an error to create an empty union; they are syntactically not possible.
>>>>>>> upstream/master
func NewUnion(terms []*Term) *Union

func (u *Union) Len() int
func (u *Union) Term(i int) *Term

func (u *Union) Underlying() Type
func (u *Union) String() string

<<<<<<< HEAD
// TermはUnion内の項を表します。
=======
// A Term represents a term in a [Union].
>>>>>>> upstream/master
type Term term

// NewTermは新しいユニオン用語を返します。
func NewTerm(tilde bool, typ Type) *Term

func (t *Term) Tilde() bool
func (t *Term) Type() Type
func (t *Term) String() string
