// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(neelance): implement with actual atomic operations as soon as threads are available
// See https://github.com/WebAssembly/design/issues/1073

// Export some functions via linkname to assembly in sync/atomic.
//
//go:linkname Load
//go:linkname Loadp
//go:linkname Load64
//go:linkname Loadint32
//go:linkname Loadint64
//go:linkname Loaduintptr
//go:linkname Xadd
//go:linkname Xaddint32
//go:linkname Xaddint64
//go:linkname Xadd64
//go:linkname Xadduintptr
//go:linkname Xchg
//go:linkname Xchg64
//go:linkname Xchgint32
//go:linkname Xchgint64
//go:linkname Xchguintptr
//go:linkname Cas
//go:linkname Cas64
//go:linkname Casint32
//go:linkname Casint64
//go:linkname Casuintptr
//go:linkname Store
//go:linkname Store64
//go:linkname Storeint32
//go:linkname Storeint64
//go:linkname Storeuintptr

package atomic

import "github.com/shogo82148/std/unsafe"

//go:nosplit
//go:noinline
func Load(ptr *uint32) uint32

//go:nosplit
//go:noinline
func Loadp(ptr unsafe.Pointer) unsafe.Pointer

//go:nosplit
//go:noinline
func LoadAcq(ptr *uint32) uint32

//go:nosplit
//go:noinline
func LoadAcq64(ptr *uint64) uint64

//go:nosplit
//go:noinline
func LoadAcquintptr(ptr *uintptr) uintptr

//go:nosplit
//go:noinline
func Load8(ptr *uint8) uint8

//go:nosplit
//go:noinline
func Load64(ptr *uint64) uint64

//go:nosplit
//go:noinline
func Xadd(ptr *uint32, delta int32) uint32

//go:nosplit
//go:noinline
func Xadd64(ptr *uint64, delta int64) uint64

//go:nosplit
//go:noinline
func Xadduintptr(ptr *uintptr, delta uintptr) uintptr

//go:nosplit
//go:noinline
func Xchg(ptr *uint32, new uint32) uint32

//go:nosplit
//go:noinline
func Xchg64(ptr *uint64, new uint64) uint64

//go:nosplit
//go:noinline
func Xchgint32(ptr *int32, new int32) int32

//go:nosplit
//go:noinline
func Xchgint64(ptr *int64, new int64) int64

//go:nosplit
//go:noinline
func Xchguintptr(ptr *uintptr, new uintptr) uintptr

//go:nosplit
//go:noinline
func And8(ptr *uint8, val uint8)

//go:nosplit
//go:noinline
func Or8(ptr *uint8, val uint8)

//go:nosplit
//go:noinline
func And(ptr *uint32, val uint32)

//go:nosplit
//go:noinline
func Or(ptr *uint32, val uint32)

//go:nosplit
//go:noinline
func Cas64(ptr *uint64, old, new uint64) bool

//go:nosplit
//go:noinline
func Store(ptr *uint32, val uint32)

//go:nosplit
//go:noinline
func StoreRel(ptr *uint32, val uint32)

//go:nosplit
//go:noinline
func StoreRel64(ptr *uint64, val uint64)

//go:nosplit
//go:noinline
func StoreReluintptr(ptr *uintptr, val uintptr)

//go:nosplit
//go:noinline
func Store8(ptr *uint8, val uint8)

//go:nosplit
//go:noinline
func Store64(ptr *uint64, val uint64)

// StorepNoWB performs *ptr = val atomically and without a write
// barrier.
//
// NO go:noescape annotation; see atomic_pointer.go.
func StorepNoWB(ptr unsafe.Pointer, val unsafe.Pointer)

//go:nosplit
//go:noinline
func Casint32(ptr *int32, old, new int32) bool

//go:nosplit
//go:noinline
func Casint64(ptr *int64, old, new int64) bool

//go:nosplit
//go:noinline
func Cas(ptr *uint32, old, new uint32) bool

//go:nosplit
//go:noinline
func Casp1(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool

//go:nosplit
//go:noinline
func Casuintptr(ptr *uintptr, old, new uintptr) bool

//go:nosplit
//go:noinline
func CasRel(ptr *uint32, old, new uint32) bool

//go:nosplit
//go:noinline
func Storeint32(ptr *int32, new int32)

//go:nosplit
//go:noinline
func Storeint64(ptr *int64, new int64)

//go:nosplit
//go:noinline
func Storeuintptr(ptr *uintptr, new uintptr)

//go:nosplit
//go:noinline
func Loaduintptr(ptr *uintptr) uintptr

//go:nosplit
//go:noinline
func Loaduint(ptr *uint) uint

//go:nosplit
//go:noinline
func Loadint32(ptr *int32) int32

//go:nosplit
//go:noinline
func Loadint64(ptr *int64) int64

//go:nosplit
//go:noinline
func Xaddint32(ptr *int32, delta int32) int32

//go:nosplit
//go:noinline
func Xaddint64(ptr *int64, delta int64) int64

//go:nosplit
//go:noinline
func And32(ptr *uint32, val uint32) uint32

//go:nosplit
//go:noinline
func And64(ptr *uint64, val uint64) uint64

//go:nosplit
//go:noinline
func Anduintptr(ptr *uintptr, val uintptr) uintptr

//go:nosplit
//go:noinline
func Or32(ptr *uint32, val uint32) uint32

//go:nosplit
//go:noinline
func Or64(ptr *uint64, val uint64) uint64

//go:nosplit
//go:noinline
func Oruintptr(ptr *uintptr, val uintptr) uintptr
