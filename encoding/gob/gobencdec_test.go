// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests of the GobEncoder/GobDecoder support.

package gob

type ByteStruct struct {
	a byte
}

type StringStruct struct {
	s string
}

type ArrayStruct struct {
	a [8192]byte
}

type Gobber int

type ValueGobber string

type BinaryGobber int

type BinaryValueGobber string

type TextGobber int

type TextValueGobber string

type GobTest0 struct {
	X int
	G *ByteStruct
}

type GobTest1 struct {
	X int
	G *StringStruct
}

type GobTest2 struct {
	X int
	G string
}

type GobTest3 struct {
	X int
	G *Gobber
	B *BinaryGobber
	T *TextGobber
}

type GobTest4 struct {
	X  int
	V  ValueGobber
	BV BinaryValueGobber
	TV TextValueGobber
}

type GobTest5 struct {
	X  int
	V  *ValueGobber
	BV *BinaryValueGobber
	TV *TextValueGobber
}

type GobTest6 struct {
	X  int
	V  ValueGobber
	W  *ValueGobber
	BV BinaryValueGobber
	BW *BinaryValueGobber
	TV TextValueGobber
	TW *TextValueGobber
}

type GobTest7 struct {
	X  int
	V  *ValueGobber
	W  ValueGobber
	BV *BinaryValueGobber
	BW BinaryValueGobber
	TV *TextValueGobber
	TW TextValueGobber
}

type GobTestIgnoreEncoder struct {
	X int
}

type GobTestValueEncDec struct {
	X int
	G StringStruct
}

type GobTestIndirectEncDec struct {
	X int
	G ***StringStruct
}

type GobTestArrayEncDec struct {
	X int
	A ArrayStruct
}

type GobTestIndirectArrayEncDec struct {
	X int
	A ***ArrayStruct
}
