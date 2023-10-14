// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// A Cache holds reusable compiler state.
// It is intended to be re-used for multiple Func compilations.
type Cache struct {
	// Storage for low-numbered values and blocks.
	values [2000]Value
	blocks [200]Block
	locs   [2000]Location

	// Reusable stackAllocState.
	// See stackalloc.go's {new,put}StackAllocState.
	stackAllocState *stackAllocState

	scrPoset []*poset

	// Reusable regalloc state.
	regallocValues []valState

	ValueToProgAfter []*obj.Prog
	debugState       debugState

	Liveness interface{}

	// Free "headers" for use by the allocators in allocators.go.
	// Used to put slices in sync.Pools without allocation.
	hdrValueSlice []*[]*Value
	hdrInt64Slice []*[]int64
}

func (c *Cache) Reset()
