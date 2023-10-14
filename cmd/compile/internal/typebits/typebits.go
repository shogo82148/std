// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typebits

import (
	"github.com/shogo82148/std/cmd/compile/internal/bitvec"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// NOTE: The bitmap for a specific type t could be cached in t after
// the first run and then simply copied into bv at the correct offset
// on future calls with the same type t.
func Set(t *types.Type, off int64, bv bitvec.BitVec)

// SetNoCheck is like Set, but do not check for alignment.
func SetNoCheck(t *types.Type, off int64, bv bitvec.BitVec)
