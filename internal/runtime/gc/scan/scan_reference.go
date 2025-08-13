// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"github.com/shogo82148/std/internal/runtime/gc"
	"github.com/shogo82148/std/unsafe"
)

// ScanSpanPackedReference is the reference implementation of ScanScanPacked. It prioritizes clarity over performance.
//
// Concretely, ScanScanPacked functions read pointers from mem, assumed to be gc.PageSize-aligned and gc.PageSize in size,
// and writes them to bufp, which is large enough to guarantee that even if pointer-word of mem is a pointer, it will fit.
// Therefore bufp, is always at least gc.PageSize in size.
//
// ScanSpanPacked is supposed to identify pointers by first filtering words by objMarks, where each bit of the mask
// represents gc.SizeClassToSize[sizeClass] bytes of memory, and then filtering again by the bits in ptrMask.
func ScanSpanPackedReference(mem unsafe.Pointer, bufp *uintptr, objMarks *gc.ObjMask, sizeClass uintptr, ptrMask *gc.PtrMask) (count int32)
