// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm || s390x || mips || mipsle || mips64 || mips64le || wasm

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
