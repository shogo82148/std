// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Value interface {
	kind() *kind
	getOp() *op
}

type Word interface {
	Value
	isWord()
}

type Ptr[T Value] struct {
	valGP
}

// Ptr is a Word
var _ Word = Ptr[void]{}

func (x Ptr[T]) AddConst(off int) (y Ptr[T])

func Deref[W wrap[T], T Value](ptr Ptr[W]) T

type Array[T Value] struct {
	valAny
}

func ConstArray[T Value](vals []T, name string) (y Array[T])
