// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// NewSparseSet returns a sparseSet that can represent
// integers between 0 and n-1.
func NewSparseSet(n int) *SparseSet

type SparseSet struct {
	dense  []ID
	sparse []int32
}

func (s *SparseSet) Size() int

func (s *SparseSet) Contains(x ID) bool

func (s *SparseSet) Add(x ID)

func (s *SparseSet) Remove(x ID)

// Pop removes an arbitrary element from the set.
// The set must be nonempty.
func (s *SparseSet) Pop() ID

func (s *SparseSet) Clear()

func (s *SparseSet) Contents() []ID
