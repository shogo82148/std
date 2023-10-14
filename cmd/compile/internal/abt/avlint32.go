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
type T struct {
	root *node32
	size int
}

// IsEmpty returns true iff t is empty.
func (t *T) IsEmpty() bool

// IsSingle returns true iff t is a singleton (leaf).
func (t *T) IsSingle() bool

// VisitInOrder applies f to the key and data pairs in t,
// with keys ordered from smallest to largest.
func (t *T) VisitInOrder(f func(int32, interface{}))

// Find returns the data associated with x in the tree, or
// nil if x is not in the tree.
func (t *T) Find(x int32) interface{}

// Insert either adds x to the tree if x was not previously
// a key in the tree, or updates the data for x in the tree if
// x was already a key in the tree.  The previous data associated
// with x is returned, and is nil if x was not previously a
// key in the tree.
func (t *T) Insert(x int32, data interface{}) interface{}

func (t *T) Copy() *T

func (t *T) Delete(x int32) interface{}

func (t *T) DeleteMin() (int32, interface{})

func (t *T) DeleteMax() (int32, interface{})

func (t *T) Size() int

// Intersection returns the intersection of t and u, where the result
// data for any common keys is given by f(t's data, u's data) -- f need
// not be symmetric.  If f returns nil, then the key and data are not
// added to the result.  If f itself is nil, then whatever value was
// already present in the smaller set is used.
func (t *T) Intersection(u *T, f func(x, y interface{}) interface{}) *T

// Union returns the union of t and u, where the result data for any common keys
// is given by f(t's data, u's data) -- f need not be symmetric.  If f returns nil,
// then the key and data are not added to the result.  If f itself is nil, then
// whatever value was already present in the larger set is used.
func (t *T) Union(u *T, f func(x, y interface{}) interface{}) *T

// Difference returns the difference of t and u, subject to the result
// of f applied to data corresponding to equal keys.  If f returns nil
// (or if f is nil) then the key+data are excluded, as usual.  If f
// returns not-nil, then that key+data pair is inserted. instead.
func (t *T) Difference(u *T, f func(x, y interface{}) interface{}) *T

func (t *T) Iterator() Iterator

func (t *T) Equals(u *T) bool

func (t *T) String() string

func (t *T) Equiv(u *T, eqv func(x, y interface{}) bool) bool

type Iterator struct {
	it iterator
}

func (it *Iterator) Next() (int32, interface{})

func (it *Iterator) Done() bool

// Min returns the minimum element of t.
// If t is empty, then (NOT_KEY32, nil) is returned.
func (t *T) Min() (k int32, d interface{})

// Max returns the maximum element of t.
// If t is empty, then (NOT_KEY32, nil) is returned.
func (t *T) Max() (k int32, d interface{})

// Glb returns the greatest-lower-bound-exclusive of x and the associated
// data.  If x has no glb in the tree, then (NOT_KEY32, nil) is returned.
func (t *T) Glb(x int32) (k int32, d interface{})

// GlbEq returns the greatest-lower-bound-inclusive of x and the associated
// data.  If x has no glbEQ in the tree, then (NOT_KEY32, nil) is returned.
func (t *T) GlbEq(x int32) (k int32, d interface{})

// Lub returns the least-upper-bound-exclusive of x and the associated
// data.  If x has no lub in the tree, then (NOT_KEY32, nil) is returned.
func (t *T) Lub(x int32) (k int32, d interface{})

// LubEq returns the least-upper-bound-inclusive of x and the associated
// data.  If x has no lubEq in the tree, then (NOT_KEY32, nil) is returned.
func (t *T) LubEq(x int32) (k int32, d interface{})
