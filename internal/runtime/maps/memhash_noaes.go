// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(amd64 || arm64 || 386)

package maps

import (
	"github.com/shogo82148/std/unsafe"
)

func MemHash(p unsafe.Pointer, h, s uintptr) uintptr

func MemHash32(p unsafe.Pointer, h uintptr) uintptr

func MemHash64(p unsafe.Pointer, h uintptr) uintptr

func StrHash(p unsafe.Pointer, h uintptr) uintptr
