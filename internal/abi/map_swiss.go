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
	// Number of bits in the group.slot count.
	SwissMapGroupSlotsBits = 3

	// Number of slots in a group.
	SwissMapGroupSlots = 1 << SwissMapGroupSlotsBits

	// Maximum key or elem size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	SwissMapMaxKeyBytes  = 128
	SwissMapMaxElemBytes = 128

	// Value of control word with all empty slots.
	SwissMapCtrlEmpty = bitsetLSB * uint64(ctrlEmpty)
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
	SwissMapIndirectKey
	SwissMapIndirectElem
)

func (mt *SwissMapType) NeedKeyUpdate() bool

func (mt *SwissMapType) HashMightPanic() bool

func (mt *SwissMapType) IndirectKey() bool

func (mt *SwissMapType) IndirectElem() bool
