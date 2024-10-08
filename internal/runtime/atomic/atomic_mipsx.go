// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips || mipsle

// Export some functions via linkname to assembly in sync/atomic.
//
//go:linkname Xadd64
//go:linkname Xchg64
//go:linkname Cas64
//go:linkname Load64
//go:linkname Store64
//go:linkname Or64
//go:linkname And64

package atomic

import (
	"github.com/shogo82148/std/unsafe"
)

//go:nosplit
func Xadd64(addr *uint64, delta int64) (new uint64)

//go:nosplit
func Xchg64(addr *uint64, new uint64) (old uint64)

//go:nosplit
func Cas64(addr *uint64, old, new uint64) (swapped bool)

//go:nosplit
func Load64(addr *uint64) (val uint64)

//go:nosplit
func Store64(addr *uint64, val uint64)

//go:nosplit
func Or64(addr *uint64, val uint64) (old uint64)

//go:nosplit
func And64(addr *uint64, val uint64) (old uint64)

//go:noescape
func Xadd(ptr *uint32, delta int32) uint32

//go:noescape
func Xadduintptr(ptr *uintptr, delta uintptr) uintptr

//go:noescape
func Xchg(ptr *uint32, new uint32) uint32

//go:noescape
func Xchguintptr(ptr *uintptr, new uintptr) uintptr

//go:noescape
func Load(ptr *uint32) uint32

//go:noescape
func Load8(ptr *uint8) uint8

// NO go:noescape annotation; *ptr escapes if result escapes (#31525)
func Loadp(ptr unsafe.Pointer) unsafe.Pointer

//go:noescape
func LoadAcq(ptr *uint32) uint32

//go:noescape
func LoadAcquintptr(ptr *uintptr) uintptr

//go:noescape
func And8(ptr *uint8, val uint8)

//go:noescape
func Or8(ptr *uint8, val uint8)

//go:noescape
func And(ptr *uint32, val uint32)

//go:noescape
func Or(ptr *uint32, val uint32)

//go:noescape
func And32(ptr *uint32, val uint32) uint32

//go:noescape
func Or32(ptr *uint32, val uint32) uint32

//go:noescape
func Anduintptr(ptr *uintptr, val uintptr) uintptr

//go:noescape
func Oruintptr(ptr *uintptr, val uintptr) uintptr

//go:noescape
func Store(ptr *uint32, val uint32)

//go:noescape
func Store8(ptr *uint8, val uint8)

// NO go:noescape annotation; see atomic_pointer.go.
func StorepNoWB(ptr unsafe.Pointer, val unsafe.Pointer)

//go:noescape
func StoreRel(ptr *uint32, val uint32)

//go:noescape
func StoreReluintptr(ptr *uintptr, val uintptr)

//go:noescape
func CasRel(addr *uint32, old, new uint32) bool
