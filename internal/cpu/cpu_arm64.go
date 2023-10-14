// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// CacheLinePadSize is used to prevent false sharing of cache lines.
// We choose 128 because Apple Silicon, a.k.a. M1, has 128-byte cache line size.
// It doesn't cost much and is much more future-proof.
const CacheLinePadSize = 128
