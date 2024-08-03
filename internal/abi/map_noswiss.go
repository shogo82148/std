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
	// Maximum number of key/elem pairs a bucket can hold.
	OldMapBucketCountBits = 3
	OldMapBucketCount     = 1 << OldMapBucketCountBits

	// Maximum key or elem size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	// Note: fast map functions cannot handle big elems (bigger than MapMaxElemBytes).
	OldMapMaxKeyBytes  = 128
	OldMapMaxElemBytes = 128
)

type OldMapType struct {
	Type
	Key    *Type
	Elem   *Type
	Bucket *Type
	// function for hashing keys (ptr to key, seed) -> hash
	Hasher     func(unsafe.Pointer, uintptr) uintptr
	KeySize    uint8
	ValueSize  uint8
	BucketSize uint16
	Flags      uint32
}

// Note: flag values must match those used in the TMAP case
// in ../cmd/compile/internal/reflectdata/reflect.go:writeType.
func (mt *OldMapType) IndirectKey() bool

func (mt *OldMapType) IndirectElem() bool

func (mt *OldMapType) ReflexiveKey() bool

func (mt *OldMapType) NeedKeyUpdate() bool

func (mt *OldMapType) HashMightPanic() bool
