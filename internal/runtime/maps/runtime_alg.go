// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maps

// MinAeshashSize is the smallest key size hashed with the AES-based
// implementation. Selecting hash is a size comparison against this value.
// Setting this to MaxUintptr disables AES altogether.
//
// When support is detected, the threshold is lowered to select between the
// scalar-based fallback hash and the vector-based AES hash. This value is
// selected on a per-platform basis based on what value produces the best
// benchmark results.
//
// Scalar hashes are faster on small values because it avoids taking a trip
// into the vector unit, which hurts latency (and for very small values,
// throughput).
var MinAeshashSize uintptr = ^uintptr(0)

// AeshashEnabled reports whether this machine hashes any sizes with AES.
//
// Test-only; compare against MinAeshashSize in non-test code to fuse this
// comparison with the MinAeshashSize check.
func AeshashEnabled() bool

func AlgInit()
