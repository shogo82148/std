// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

// Map constants common to several packages
// runtime/runtime-gdb.py:MapTypePrinter contains its own copy
const (
	// Maximum number of key/elem pairs a bucket can hold.
	MapBucketCountBits = 3
	MapBucketCount     = 1 << MapBucketCountBits

	// Maximum key or elem size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	// Note: fast map functions cannot handle big elems (bigger than MapMaxElemBytes).
	MapMaxKeyBytes  = 128
	MapMaxElemBytes = 128
)
<<<<<<< HEAD
=======

type MapType struct {
	Type
	Key   *Type
	Elem  *Type
	Group *Type
	// function for hashing keys (ptr to key, seed) -> hash
	Hasher    func(unsafe.Pointer, uintptr) uintptr
	GroupSize uintptr
	// These fields describe how to access keys and elems within a group.
	// The formulas key(i) = KeysOff + i*KeyStride and
	// elem(i) = ElemsOff + i*ElemStride work for both group layouts:
	//
	// With GOEXPERIMENT=mapsplitgroup (split arrays KKKKVVVV):
	//   KeysOff    = offset of keys array in group
	//   KeyStride  = size of a single key
	//   ElemsOff   = offset of elems array in group
	//   ElemStride = size of a single elem
	//
	// Without (interleaved slots KVKVKVKV):
	//   KeysOff    = offset of slots array in group
	//   KeyStride  = size of a key/elem slot (stride between keys)
	//   ElemsOff   = offset of first elem (slots offset + elem offset within slot)
	//   ElemStride = size of a key/elem slot (stride between elems)
	KeysOff    uintptr
	KeyStride  uintptr
	ElemsOff   uintptr
	ElemStride uintptr
	ElemOff    uintptr
	Flags      uint32
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
>>>>>>> af828fe07e833d4b09d20a822ebe5fb87440e115
