// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

type Bar struct {
	X string
}

// This structure has pointers and refers to itself, making it a good test case.
type Foo struct {
	A int
	B int32
	C string
	D []byte
	E *float64
	F ****float64
	G *Bar
	H *Bar
	I *Foo
}

type N1 struct{}
type N2 struct{}
