// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm

package atomic

import (
	"github.com/shogo82148/std/unsafe"
)

// Atomic add and return new value.
//
//go:nosplit
func Xadd(val *uint32, delta int32) uint32

//go:noescape
func Xadduintptr(ptr *uintptr, delta uintptr) uintptr

//go:nosplit
func Xchg(addr *uint32, v uint32) uint32

//go:nosplit
func Xchguintptr(addr *uintptr, v uintptr) uintptr

// Not noescape -- it installs a pointer to addr.
func StorepNoWB(addr unsafe.Pointer, v unsafe.Pointer)

//go:noescape
func Store(addr *uint32, v uint32)

//go:noescape
func StoreRel(addr *uint32, v uint32)

//go:noescape
func StoreReluintptr(addr *uintptr, v uintptr)

//go:noescape
func Or8(addr *uint8, v uint8)

//go:noescape
func And8(addr *uint8, v uint8)

//go:nosplit
func Or(addr *uint32, v uint32)

//go:nosplit
func And(addr *uint32, v uint32)

//go:noescape
func Load(addr *uint32) uint32

// NO go:noescape annotation; *addr escapes if result escapes (#31525)
func Loadp(addr unsafe.Pointer) unsafe.Pointer

//go:noescape
func Load8(addr *uint8) uint8

//go:noescape
func LoadAcq(addr *uint32) uint32

//go:noescape
func LoadAcquintptr(ptr *uintptr) uintptr

//go:noescape
func Cas64(addr *uint64, old, new uint64) bool

//go:noescape
func CasRel(addr *uint32, old, new uint32) bool

//go:noescape
func Xadd64(addr *uint64, delta int64) uint64

//go:noescape
func Xchg64(addr *uint64, v uint64) uint64

//go:noescape
func Load64(addr *uint64) uint64

//go:noescape
func Store8(addr *uint8, v uint8)

//go:noescape
func Store64(addr *uint64, v uint64)
