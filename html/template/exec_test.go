// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests for template execution, copied from text/template.

package template

import (
	"fmt"
)

// T has lots of interesting pieces to use to test execution.
type T struct {
	True        bool
	I           int
	U16         uint16
	X, S        string
	FloatZero   float64
	ComplexZero complex128

	U *U

	V0     V
	V1, V2 *V

	W0     W
	W1, W2 *W

	SI      []int
	SICap   []int
	SIEmpty []int
	SB      []bool

	AI [3]int

	MSI      map[string]int
	MSIone   map[string]int
	MSIEmpty map[string]int
	MXI      map[interface{}]int
	MII      map[int]int
	MI32S    map[int32]string
	MI64S    map[int64]string
	MUI32S   map[uint32]string
	MUI64S   map[uint64]string
	MI8S     map[int8]string
	MUI8S    map[uint8]string
	SMSI     []map[string]int

	Empty0 interface{}
	Empty1 interface{}
	Empty2 interface{}
	Empty3 interface{}
	Empty4 interface{}

	NonEmptyInterface         I
	NonEmptyInterfacePtS      *I
	NonEmptyInterfaceNil      I
	NonEmptyInterfaceTypedNil I

	Str fmt.Stringer
	Err error

	PI  *int
	PS  *string
	PSI *[]int
	NIL *int

	BinaryFunc      func(string, string) string
	VariadicFunc    func(...string) string
	VariadicFuncInt func(int, ...string) string
	NilOKFunc       func(*int) bool
	ErrFunc         func() (string, error)
	PanicFunc       func() string

	Tmpl *Template

	unexported int
}

type S []string

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
// Also test the trimming of leading and trailing spaces.

type ErrorWriter int

// recursiveInvoker is for TestRecursiveExecuteViaMethod.
