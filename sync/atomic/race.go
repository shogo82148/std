// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race
// +build race

package atomic

import (
	"github.com/shogo82148/std/unsafe"
)

func CompareAndSwapInt32(val *int32, old, new int32) bool

func CompareAndSwapUint32(val *uint32, old, new uint32) (swapped bool)

func CompareAndSwapInt64(val *int64, old, new int64) bool

func CompareAndSwapUint64(val *uint64, old, new uint64) (swapped bool)

func CompareAndSwapPointer(val *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

func CompareAndSwapUintptr(val *uintptr, old, new uintptr) (swapped bool)

func AddInt32(val *int32, delta int32) int32

func AddUint32(val *uint32, delta uint32) (new uint32)

func AddInt64(val *int64, delta int64) int64

func AddUint64(val *uint64, delta uint64) (new uint64)

func AddUintptr(val *uintptr, delta uintptr) (new uintptr)

func LoadInt32(addr *int32) int32

func LoadUint32(addr *uint32) (val uint32)

func LoadInt64(addr *int64) int64

func LoadUint64(addr *uint64) (val uint64)

func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

func LoadUintptr(addr *uintptr) (val uintptr)

func StoreInt32(addr *int32, val int32)

func StoreUint32(addr *uint32, val uint32)

func StoreInt64(addr *int64, val int64)

func StoreUint64(addr *uint64, val uint64)

func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)

func StoreUintptr(addr *uintptr, val uintptr)
