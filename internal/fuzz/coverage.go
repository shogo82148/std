// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuzz

// ResetCoverage sets all of the counters for each edge of the instrumented
// source code to 0.
func ResetCoverage()

// SnapshotCoverage copies the current counter values into coverageSnapshot,
// preserving them for later inspection. SnapshotCoverage also rounds each
// counter down to the nearest power of two. This lets the coordinator store
// multiple values for each counter by OR'ing them together.
func SnapshotCoverage()
