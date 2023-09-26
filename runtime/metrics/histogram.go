// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

// Float64Histogram represents a distribution of float64 values.
type Float64Histogram struct {
	Counts []uint64

	Buckets []float64
}
