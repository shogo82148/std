// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(amd64 || arm64 || 386)

package maps

import (
	"github.com/shogo82148/std/unsafe"
)

func MemHash(p unsafe.Pointer, h, s uintptr) uintptr

func MemHash32(k uint32, h uintptr) uintptr

func MemHash64(k uint64, h uintptr) uintptr

func StrHash(s string, h uintptr) uintptr
