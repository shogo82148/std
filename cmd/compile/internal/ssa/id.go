// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

type ID int32

// IDAlloc provides an allocator for unique integers.
type IDAlloc struct {
	last ID
}

// Get allocates an ID and returns it. IDs are always > 0.
func (a *IDAlloc) Get() ID

// Num returns the maximum ID ever returned + 1.
func (a *IDAlloc) Num() int
