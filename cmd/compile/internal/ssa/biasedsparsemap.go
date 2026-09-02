// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// A BiasedSparseMap is a sparseMap for integers between J and K inclusive,
// where J might be somewhat larger than zero (and K-J is probably much smaller than J).
// (The motivating use case is the line numbers of statements for a single function.)
// Not all features of a SparseMap are exported, and it is also easy to treat a
// BiasedSparseMap like a SparseSet.
type BiasedSparseMap struct {
	s     *SparseMap
	first int
}
