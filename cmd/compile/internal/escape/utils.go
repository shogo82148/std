// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package escape

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// HeapAllocReason returns the reason the given Node must be heap
// allocated, or the empty string if it doesn't.
func HeapAllocReason(n ir.Node) string
