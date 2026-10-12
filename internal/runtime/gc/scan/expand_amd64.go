// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import "github.com/shogo82148/std/internal/runtime/gc"

// ExpandAVX512 expands each bit in packed into f consecutive bits in unpacked,
// where f is the word size of objects in sizeClass.
//
// This is a testing entrypoint to the expanders used by scanSpanPacked*.
//
//go:noescape
func ExpandAVX512(sizeClass int, packed *gc.ObjMask, unpacked *gc.PtrMask)
