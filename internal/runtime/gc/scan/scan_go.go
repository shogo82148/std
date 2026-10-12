// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"github.com/shogo82148/std/internal/runtime/gc"
	"github.com/shogo82148/std/unsafe"
)

// ScanSpanPackedGo is an optimized pure Go implementation of ScanSpanPacked.
func ScanSpanPackedGo(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32)
