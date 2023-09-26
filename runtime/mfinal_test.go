// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

type Tintptr *int
type Tint int

type Tinter interface {
	m()
}

// One chunk must be exactly one sizeclass in size.
// It should be a sizeclass not used much by others, so we
// have a greater chance of finding adjacent ones.
// size class 19: 320 byte objects, 25 per page, 1 page alloc at a time

type Object1 struct {
	Something []byte
}

type Object2 struct {
	Something byte
}

var (
	Foo2 = &Object2{}
	Foo1 = &Object1{}
)
