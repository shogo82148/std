// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// NewSparseMap returns a sparseMap that can map
// integers between 0 and n-1 to int32s.
func NewSparseMap(n int) *SparseMap

func NewSparseMapValue(n int) *SparseMapValue

type SparseMap = genericSparseMap[ID, int32]

type SparseMapValue = genericSparseMap[ID, *Value]

func (s *genericSparseMap[K, V]) Contains(k K) bool

// Get returns the value for key k, or the zero V
// if k does not appear in the map.
func (s *genericSparseMap[K, V]) Get(k K) (V, bool)

func (s *genericSparseMap[K, V]) Set(k K, v V)

func (s *genericSparseMap[K, V]) Clear()

// ClearDense zeroes out the dense slice and then truncates it. This is useful for
// making sure that maps that contain pointer data doesn't keep values live longer
// than necessary.
func (s *genericSparseMap[K, V]) ClearDense()
