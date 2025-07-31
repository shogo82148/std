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
	MapGroupSlotsBits = 3

	// Number of slots in a group.
	MapGroupSlots = 1 << MapGroupSlotsBits

	// Maximum key or elem size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	MapMaxKeyBytes  = 128
	MapMaxElemBytes = 128

	// Value of control word with all empty slots.
	MapCtrlEmpty = bitsetLSB * uint64(ctrlEmpty)
)

type MapType struct {
	Type
	Key   *Type
	Elem  *Type
	Group *Type
	// function for hashing keys (ptr to key, seed) -> hash
	Hasher    func(unsafe.Pointer, uintptr) uintptr
	GroupSize uintptr
	SlotSize  uintptr
	ElemOff   uintptr
	Flags     uint32
}

// Flag values
const (
	MapNeedKeyUpdate = 1 << iota
	MapHashMightPanic
	MapIndirectKey
	MapIndirectElem
)

func (mt *MapType) NeedKeyUpdate() bool

func (mt *MapType) HashMightPanic() bool

func (mt *MapType) IndirectKey() bool

func (mt *MapType) IndirectElem() bool
