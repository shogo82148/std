// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

import (
	"github.com/shogo82148/std/unsafe"
)

// Map constants common to several packages
// runtime/runtime-gdb.py:MapTypePrinter contains its own copy
const (
	// Number of slots in a group.
	SwissMapGroupSlots = 8
)

type SwissMapType struct {
	Type
	Key   *Type
	Elem  *Type
	Group *Type
	// function for hashing keys (ptr to key, seed) -> hash
	Hasher   func(unsafe.Pointer, uintptr) uintptr
	SlotSize uintptr
	ElemOff  uintptr
	Flags    uint32
}

// Flag values
const (
	SwissMapNeedKeyUpdate = 1 << iota
	SwissMapHashMightPanic
)

func (mt *SwissMapType) NeedKeyUpdate() bool

func (mt *SwissMapType) HashMightPanic() bool
