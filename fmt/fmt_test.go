// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt_test

import (
	"bytes"
	. "fmt"
	"math"
)

var (
	NaN = math.NaN()
)

type A struct {
	i int
	j uint
	s string
	x []int
}

type I int

type B struct {
	I I
	j int
}

type C struct {
	i int
	B
}

type F int

type G int

type S struct {
	F F
	G G
}

type SI struct {
	I interface{}
}

// P is a type with a String method with pointer receiver for testing %p.
type P int

type SE []interface{}

var _ bytes.Buffer

// PanicS is a type that panics in String.
type PanicS struct {
	message interface{}
}

// PanicGo is a type that panics in GoString.
type PanicGo struct {
	message interface{}
}

// PanicF is a type that panics in Format.
type PanicF struct {
	message interface{}
}

// recurCount tests that erroneous String routine doesn't cause fatal recursion.

type Recur struct {
	i      int
	failed *bool
}
