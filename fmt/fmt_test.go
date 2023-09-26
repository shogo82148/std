// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt_test

import (
	"bytes"
	. "fmt"
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

// A type with a String method with pointer receiver for testing %p
type P int

var _ bytes.Buffer

// A type that panics in String.
type Panic struct {
	message interface{}
}

// A type that panics in Format.
type PanicF struct {
	message interface{}
}

// Test that erroneous String routine doesn't cause fatal recursion.

type Recur struct {
	i      int
	failed *bool
}
