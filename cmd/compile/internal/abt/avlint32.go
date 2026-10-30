// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abt

const (
	LEAF_HEIGHT = 1
	ZERO_HEIGHT = 0
	NOT_KEY32   = int32(-0x80000000)
)

// T is the exported applicative balanced tree data type.
// A T can be used as a value; updates to one copy of the value
// do not change other copies.
type T[D comparable] struct {
	root *node32[D]
	size int
}

// IsEmpty returns true iff t is empty.
func (t *T[D]) IsEmpty() bool

// IsSingle returns true iff t is a singleton (leaf).
func (t *T[D]) IsSingle() bool

// VisitInOrder applies f to the key and data pairs in t,
// with keys ordered from smallest to largest.
func (t *T[D]) VisitInOrder(f func(int32, D))

// Find returns the data associated with x in the tree, or
// the zero value of D if x is not in the tree.
func (t *T[D]) Find(x int32) D

// Insert either adds x to the tree if x was not previously
// a key in the tree, or updates the data for x in the tree if
// x was already a key in the tree.  The previous data associated
// with x is returned, and is the zero value of D if x was not previously a
// key in the tree.
func (t *T[D]) Insert(x int32, data D) D

func (t *T[D]) Copy() *T[D]

func (t *T[D]) Delete(x int32) (r D)

func (t *T[D]) DeleteMin() (k int32, r D)

func (t *T[D]) DeleteMax() (k int32, r D)

func (t *T[D]) Size() int

// Intersection returns the intersection of t and u, where the result
// data for any common keys is given by f(t's data, u's data) -- f need
// not be symmetric.  If f returns the zero value of D, then the key and data are not
// added to the result.  If f itself is nil, then whatever value was
// already present in the smaller set is used.
func (t *T[D]) Intersection(u *T[D], f func(x, y D) D) *T[D]

// Union returns the union of t and u, where the result data for any common keys
// is given by f(t's data, u's data) -- f need not be symmetric.  If f returns the zero value of D,
// then the key and data are not added to the result.  If f itself is nil, then
// whatever value was already present in the larger set is used.
func (t *T[D]) Union(u *T[D], f func(x, y D) D) *T[D]

// Difference returns the difference of t and u, subject to the result
// of f applied to data corresponding to equal keys.  If f returns the zero value of D
// (or if f is nil) then the key+data are excluded, as usual.  If f
// returns not-zero, then that key+data pair is inserted. instead.
func (t *T[D]) Difference(u *T[D], f func(x, y D) D) *T[D]

func (t *T[D]) Iterator() Iterator[D]

func (t *T[D]) Equals(u *T[D]) bool

func (t *T[D]) String() string

func (t *T[D]) Equiv(u *T[D], eqv func(x, y D) bool) bool

type Iterator[D comparable] struct {
	it iterator[D]
}

func (it *Iterator[D]) Next() (k int32, d D)

func (it *Iterator[D]) Done() bool

// Min returns the minimum element of t.
// If t is empty, then (NOT_KEY32, zero) is returned.
func (t *T[D]) Min() (k int32, d D)

// Max returns the maximum element of t.
// If t is empty, then (NOT_KEY32, zero) is returned.
func (t *T[D]) Max() (k int32, d D)

// Glb returns the greatest-lower-bound-exclusive of x and the associated
// data.  If x has no glb in the tree, then (NOT_KEY32, zero) is returned.
func (t *T[D]) Glb(x int32) (k int32, d D)

// GlbEq returns the greatest-lower-bound-inclusive of x and the associated
// data.  If x has no glbEQ in the tree, then (NOT_KEY32, zero) is returned.
func (t *T[D]) GlbEq(x int32) (k int32, d D)

// Lub returns the least-upper-bound-exclusive of x and the associated
// data.  If x has no lub in the tree, then (NOT_KEY32, zero) is returned.
func (t *T[D]) Lub(x int32) (k int32, d D)

// LubEq returns the least-upper-bound-inclusive of x and the associated
// data.  If x has no lubEq in the tree, then (NOT_KEY32, zero) is returned.
func (t *T[D]) LubEq(x int32) (k int32, d D)
