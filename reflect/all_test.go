// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect_test

import (
	"io"
	. "reflect"
)

type T struct {
	a int
	b float64
	c string
	d *int
}

type Basic struct {
	x int
	y float32
}

type NotBasic Basic

type DeepEqualTest struct {
	a, b interface{}
	eq   bool
}

// Simple functions for DeepEqual tests.

type Loop *Loop
type Loopy interface{}

type Recursive struct {
	x int
	r *Recursive
}

type UnexpT struct {
	m map[int]int
}

// caseInfo describes a single case in a select test.

// selectWatch and the selectWatcher are a watchdog mechanism for running Select.
// If the selectWatcher notices that the select has been blocked for >1 second, it prints
// an error describing the select and panics the entire test binary.

type Point struct {
	x, y int
}

type Tinter interface {
	M(int, byte) (byte, int)
}

type Tsmallv byte

type Tsmallp byte

type Twordv uintptr

type Twordp uintptr

type Tbigv [2]uintptr

type Tbigp [2]uintptr

type Tm1 struct {
	Tm2
}

type Tm2 struct {
	*Tm3
}

type Tm3 struct {
	*Tm4
}

type Tm4 struct {
}

type T1 struct {
	a string
	int
}

type FTest struct {
	s     interface{}
	name  string
	index []int
	value int
}

type D1 struct {
	d int
}
type D2 struct {
	d int
}

type S0 struct {
	A, B, C int
	D1
	D2
}

type S1 struct {
	B int
	S0
}

type S2 struct {
	A int
	*S1
}

type S1x struct {
	S1
}

type S1y struct {
	S1
}

type S3 struct {
	S1x
	S2
	D, E int
	*S1y
}

type S4 struct {
	*S4
	A int
}

// The X in S6 and S7 annihilate, but they also block the X in S8.S9.
type S5 struct {
	S6
	S7
	S8
}

type S6 struct {
	X int
}

type S7 S6

type S8 struct {
	S9
}

type S9 struct {
	X int
	Y int
}

// The X in S11.S6 and S12.S6 annihilate, but they also block the X in S13.S8.S9.
type S10 struct {
	S11
	S12
	S13
}

type S11 struct {
	S6
}

type S12 struct {
	S6
}

type S13 struct {
	S8
}

// The X in S15.S11.S1 and S16.S11.S1 annihilate.
type S14 struct {
	S15
	S16
}

type S15 struct {
	S11
}

type S16 struct {
	S11
}

type InnerInt struct {
	X int
}

type OuterInt struct {
	Y int
	InnerInt
}

type FuncDDD func(...interface{}) error

type Private struct {
	x int
	y **int
	Z int
}

type Public struct {
	X int
	Y **int
	private
}

var V = ValueOf

type Empty struct{}
type MyStruct struct {
	x int `some:"tag"`
}
type MyString string
type MyBytes []byte
type MyRunes []int32
type MyFunc func()
type MyByte byte

type ComparableStruct struct {
	X int
}

type NonComparableStruct struct {
	X int
	Y map[string]int
}

type StructI int

type StructIPtr int

type B1 struct {
	X int
	Y int
	Z int
}

type R0 struct {
	*R1
	*R2
	*R3
	*R4
}

type R1 struct {
	*R5
	*R6
	*R7
	*R8
}

type R2 R1
type R3 R1
type R4 R1

type R5 struct {
	*R9
	*R10
	*R11
	*R12
}

type R6 R5
type R7 R5
type R8 R5

type R9 struct {
	*R13
	*R14
	*R15
	*R16
}

type R10 R9
type R11 R9
type R12 R9

type R13 struct {
	*R17
	*R18
	*R19
	*R20
}

type R14 R13
type R15 R13
type R16 R13

type R17 struct {
	*R21
	*R22
	*R23
	*R24
}

type R18 R17
type R19 R17
type R20 R17

type R21 struct {
	X int
}

type R22 R21
type R23 R21
type R24 R21

type S struct {
	i1 int64
	i2 int64
}

// An exhaustive is a mechanism for writing exhaustive or stochastic tests.
// The basic usage is:
//
//	for x.Next() {
//		... code using x.Maybe() or x.Choice(n) to create test cases ...
//	}
//
// Each iteration of the loop returns a different set of results, until all
// possible result sets have been explored. It is okay for different code paths
// to make different method call sequences on x, but there must be no
// other source of non-determinism in the call sequences.
//
// When faced with a new decision, x chooses randomly. Future explorations
// of that path will choose successive values for the result. Thus, stopping
// the loop after a fixed number of iterations gives somewhat stochastic
// testing.
//
// Example:
//
//	for x.Next() {
//		v := make([]bool, x.Choose(4))
//		for i := range v {
//			v[i] = x.Maybe()
//		}
//		fmt.Println(v)
//	}
//
// prints (in some order):
//
//	[]
//	[false]
//	[true]
//	[false false]
//	[false true]
//	...
//	[true true]
//	[false false false]
//	...
//	[true true true]
//	[false false false false]
//	...
//	[true true true true]
//

type Outer struct {
	*Inner
	R io.Reader
}

type Inner struct {
	X  *Outer
	P1 uintptr
	P2 uintptr
}

type Impl struct{}

// Issue 18635 (method version).
type KeepMethodLive struct{}

type XM struct{ _ bool }

type TheNameOfThisTypeIsExactly255BytesLongSoWhenTheCompilerPrependsTheReflectTestPackageNameAndExtraStarTheLinkerRuntimeAndReflectPackagesWillHaveToCorrectlyDecodeTheSecondLengthByte0123456789_0123456789_0123456789_0123456789_0123456789_012345678 int

type Tint int

type Tint2 = Tint

type Talias1 struct {
	byte
	uint8
	int
	int32
	rune
}

type Talias2 struct {
	Tint
	Tint2
}
