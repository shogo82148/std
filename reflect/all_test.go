// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect_test

import (
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

type Recursive struct {
	x int
	r *Recursive
}

type UnexpT struct {
	m map[int]int
}

type Point struct {
	x, y int
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

type InnerInt struct {
	X int
}

type OuterInt struct {
	Y int
	InnerInt
}

type Private struct {
	x int
	y **int
}

type Public struct {
	X int
	Y **int
}
