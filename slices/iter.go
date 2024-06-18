// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import (
	"github.com/shogo82148/std/cmp"
	"github.com/shogo82148/std/iter"
)

// All returns an iterator over index-value pairs in the slice.
// The indexes range in the usual order, from 0 through len(s)-1.
func All[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]

// Backward returns an iterator over index-value pairs in the slice,
// traversing it backward. The indexes range from len(s)-1 down to 0.
func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]

// Values returns an iterator over the slice elements.
// starting with s[0].
func Values[Slice ~[]E, E any](s Slice) iter.Seq[E]

// AppendSeq appends the values from seq to the slice and
// returns the extended slice.
func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice

// Collect collects values from seq into a new slice and returns it.
func Collect[E any](seq iter.Seq[E]) []E

// Sorted collects values from seq into a new slice, sorts the slice,
// and returns it.
func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E

// SortedFunc collects values from seq into a new slice, sorts the slice
// using the comparison function, and returns it.
func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E

// SortedStableFunc collects values from seq into a new slice.
// It then sorts the slice while keeping the original order of equal elements,
// using the comparison function to compare elements.
// It returns the new slice.
func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E

// Chunk returns an iterator over consecutive sub-slices of up to n elements of s.
// All but the last sub-slice will have size n.
// All sub-slices are clipped to have no capacity beyond the length.
// If s is empty, the sequence is empty: there is no empty slice in the sequence.
// Chunk panics if n is less than 1.
func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice]
