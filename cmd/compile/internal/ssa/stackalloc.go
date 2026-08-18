// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: live at start of block instead?

package ssa

func NewStackAllocState(f *Func) *StackAllocState

func PutStackAllocState(s *StackAllocState)

type StackAllocState struct {
	f *Func

	// Live is the output of stackalloc.
	// Live[b.id] = Live values at the end of block b.
	Live [][]ID

	// The following slices are reused across multiple users
	// of stackAllocState.
	values    []stackValState
	interfere [][]ID
	names     []LocalSlot

	NArgSlot,
	NNotNeed,
	NNamedSlot,
	NReuse,
	NAuto,
	NSelfInterfere int32
}

func (s *StackAllocState) Init(f *Func, spillLive [][]ID)

func (s *StackAllocState) Stackalloc()

func (f *Func) GetHome(vid ID) Location

func (f *Func) SetHome(v *Value, loc Location)
