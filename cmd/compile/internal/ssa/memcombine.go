// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// A BaseAddress represents the address ptr+idx, where
// ptr is a pointer type and idx is an integer type.
// idx may be nil, in which case it is treated as 0.
type BaseAddress struct {
	ptr *Value
	idx Index
}

// Index represents an address index in the form exp<<shift.
//
// The shift is typically introduced by slice indexing (log2(element size)),
// but may also originate from shifts in the source expression.
type Index struct {
	exp   *Value
	shift int64
}
