// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "github.com/shogo82148/std/cmd/internal/src"

type SparseEntryPos struct {
	Key ID
	Val int32
	Pos src.XPos
}

type SparseMapPos struct {
	dense  []SparseEntryPos
	sparse []int32
}

func (s *SparseMapPos) Size() int

func (s *SparseMapPos) Contains(k ID) bool

// Get returns the value for key k, or -1 if k does
// not appear in the map.
func (s *SparseMapPos) Get(k ID) int32

func (s *SparseMapPos) Set(k ID, v int32, a src.XPos)

func (s *SparseMapPos) Remove(k ID)

func (s *SparseMapPos) Clear()

func (s *SparseMapPos) Contents() []SparseEntryPos
