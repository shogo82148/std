// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"math/rand"
	. "runtime"
)

type HashSet struct {
	m map[uintptr]struct{}
	n int
}

type Key interface {
	clear()
	random(r *rand.Rand)
	bits() int
	flipBit(i int)
	hash() uintptr
	name() string
}

type BytesKey struct {
	b []byte
}

type Int32Key struct {
	i uint32
}

type Int64Key struct {
	i uint64
}

// size of the hash output (32 or 64 bits)
