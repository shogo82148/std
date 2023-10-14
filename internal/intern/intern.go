// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package intern lets you make smaller comparable values by boxing
// a larger comparable value (such as a 16 byte string header) down
// into a globally unique 8 byte pointer.
//
// The globally unique pointers are garbage collected with weak
// references and finalizers. This package hides that.
package intern

// A Value pointer is the handle to an underlying comparable value.
// See func Get for how Value pointers may be used.
type Value struct {
	_      [0]func()
	cmpVal any
	// resurrected is guarded by mu (for all instances of Value).
	// It is set true whenever v is synthesized from a uintptr.
	resurrected bool
}

// Get returns the comparable value passed to the Get func
// that returned v.
func (v *Value) Get() any

// Get returns a pointer representing the comparable value cmpVal.
//
// The returned pointer will be the same for Get(v) and Get(v2)
// if and only if v == v2, and can be used as a map key.
func Get(cmpVal any) *Value

// GetByString is identical to Get, except that it is specialized for strings.
// This avoids an allocation from putting a string into an interface{}
// to pass as an argument to Get.
func GetByString(s string) *Value
