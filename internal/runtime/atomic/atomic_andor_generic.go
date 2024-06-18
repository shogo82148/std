// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm || wasm

// Export some functions via linkname to assembly in sync/atomic.
//
//go:linkname And32
//go:linkname Or32
//go:linkname And64
//go:linkname Or64
//go:linkname Anduintptr
//go:linkname Oruintptr

package atomic

//go:nosplit
func And32(ptr *uint32, val uint32) uint32

//go:nosplit
func Or32(ptr *uint32, val uint32) uint32

//go:nosplit
func And64(ptr *uint64, val uint64) uint64

//go:nosplit
func Or64(ptr *uint64, val uint64) uint64

//go:nosplit
func Anduintptr(ptr *uintptr, val uintptr) uintptr

//go:nosplit
func Oruintptr(ptr *uintptr, val uintptr) uintptr
