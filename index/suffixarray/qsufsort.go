// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This algorithm is based on "Faster Suffix Sorting"
//   by N. Jesper Larsson and Kunihiko Sadakane
// paper: http://www.larsson.dogma.net/ssrev-tr.pdf
// code:  http://www.larsson.dogma.net/qsufsort.c

// This algorithm computes the suffix array sa by computing its inverse.
// Consecutive groups of suffixes in sa are labeled as sorted groups or
// unsorted groups. For a given pass of the sorter, all suffixes are ordered
// up to their first h characters, and sa is h-ordered. Suffixes in their
// final positions and unambiguously sorted in h-order are in a sorted group.
// Consecutive groups of suffixes with identical first h characters are an
// unsorted group. In each pass of the algorithm, unsorted groups are sorted
// according to the group number of their following suffix.

// In the implementation, if sa[i] is negative, it indicates that i is
// the first element of a sorted group of length -sa[i], and can be skipped.
// An unsorted group sa[i:k] is given the group number of the index of its
// last element, k-1. The group numbers are stored in the inverse slice (inv),
// and when all groups are sorted, this slice is the inverse suffix array.

package suffixarray
