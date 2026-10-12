// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64

package scan

import (
	"github.com/shogo82148/std/internal/runtime/gc"
	"github.com/shogo82148/std/unsafe"
)

func HasFastScanSpanPacked() bool

func ScanSpanPacked(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32)
