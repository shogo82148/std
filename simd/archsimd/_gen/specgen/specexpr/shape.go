// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpr

import (
	"github.com/shogo82148/std/fmt"
)

// MinWidth is the minimum vector width, in bits.
const MinWidth = Int(128)

type Type interface {
	isType()
	String() string
}

type Basic struct {
	Base string
	Bits Int
}

var MakeBasic = MakeFunc2("Basic", func(base string, bits Int) (any, error) {

	bits = max(8, bits)
	return Basic{Base: base, Bits: bits}, nil
})

func (t Basic) String() string

type Vector struct {
	Elem  Basic
	Width Num
}

var MakeVector = MakeFunc2("VectorW", func(elem Basic, w Num) (any, error) {
	if !w.ValidWidth() {
		return nil, fmt.Errorf("invalid width %s", w)
	}
	return Vector{Elem: elem, Width: w}, nil
})

func (t Vector) String() string

func (t Vector) Scalable() bool

type Pointer struct {
	Elem Type
}

var MakePointer = MakeFunc1("Pointer", func(elem Type) (any, error) {
	return Pointer{elem}, nil
})

func (t Pointer) String() string

type Array struct {
	Elem Type
	Len  Int
}

var MakeArray = MakeFunc2("Array", func(elem Type, len Int) (any, error) {
	return Array{elem, len}, nil
})

func (t Array) String() string

type Slice struct {
	Elem Type
}

var MakeSlice = MakeFunc1("Slice", func(elem Type) (any, error) {
	return Slice{elem}, nil
})

func (t Slice) String() string
