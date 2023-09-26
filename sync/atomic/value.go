// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

// A Value provides an atomic load and store of a consistently typed value.
// Values can be created as part of other data structures.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value struct {
	noCopy noCopy

	v interface{}
}

// ifaceWords is interface{} internal representation.

// Load returns the value set by the most recent Store.
// It returns nil if there has been no call to Store for this Value.
func (v *Value) Load() (x interface{})

// Store sets the value of the Value to x.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value) Store(x interface{})

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://github.com/golang/go/issues/8005#issuecomment-190753527
// for details.
