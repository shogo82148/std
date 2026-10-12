// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"github.com/shogo82148/std/internal/runtime/gc"
)

// ExpandReference is a reference implementation of an expander function
// that translates object mark bits into a bitmap of one bit per word of
// marked object, assuming the object is of the provided size class.
func ExpandReference(sizeClass int, packed *gc.ObjMask, unpacked *gc.PtrMask)
