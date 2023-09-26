// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"fmt"
)

// T has lots of interesting pieces to use to test execution.
type T struct {
	True        bool
	I           int
	U16         uint16
	X           string
	FloatZero   float64
	ComplexZero float64

	U *U

	V0     V
	V1, V2 *V

	W0     W
	W1, W2 *W

	SI      []int
	SIEmpty []int
	SB      []bool

	MSI      map[string]int
	MSIone   map[string]int
	MSIEmpty map[string]int
	MXI      map[interface{}]int
	MII      map[int]int
	SMSI     []map[string]int

	Empty0 interface{}
	Empty1 interface{}
	Empty2 interface{}
	Empty3 interface{}
	Empty4 interface{}

	NonEmptyInterface I

	Str fmt.Stringer
	Err error

	PI  *int
	PSI *[]int
	NIL *int

	BinaryFunc      func(string, string) string
	VariadicFunc    func(...string) string
	VariadicFuncInt func(int, ...string) string
	NilOKFunc       func(*int) bool

	Tmpl *Template

	unexported int
}

type U struct {
	V string
}

type V struct {
	j int
}

type W struct {
	k int
}

// A non-empty interface.
type I interface {
	Method0() string
}

// bigInt and bigUint are hex string representing numbers either side
// of the max int boundary.
// We do it this way so the test doesn't depend on ints being 32 bits.

type Tree struct {
	Val         int
	Left, Right *Tree
}

// Use different delimiters to test Set.Delims.
