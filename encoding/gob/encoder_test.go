// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

type ET2 struct {
	X string
}

type ET1 struct {
	A    int
	Et2  *ET2
	Next *ET1
}

// Like ET1 but with a different name for a field
type ET3 struct {
	A             int
	Et2           *ET2
	DifferentNext *ET1
}

// Like ET1 but with a different type for a field
type ET4 struct {
	A    int
	Et2  float64
	Next int
}

// Types not supported by the Encoder.

type SingleTest struct {
	in  interface{}
	out interface{}
	err string
}

type Struct0 struct {
	I interface{}
}

type NewType0 struct {
	S string
}

// Another bug from golang-nuts, involving nested interfaces.
type Bug0Outer struct {
	Bug0Field interface{}
}

type Bug0Inner struct {
	A int
}

type Bug1Elem struct {
	Name string
	Id   int
}

type Bug1StructMap map[string]Bug1Elem

// Should be able to have unrepresentable fields (chan, func) as long as they
// are unexported.
type Bug2 struct {
	A int
	b chan int
}

// Mutually recursive slices of structs caused problems.
type Bug3 struct {
	Num      int
	Children []*Bug3
}

type Z struct {
}
