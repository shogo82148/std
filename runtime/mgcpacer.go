// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	_ "github.com/shogo82148/std/unsafe"
)

// gcController implements the GC pacing controller that determines
// when to trigger concurrent garbage collection and how much marking
// work to do in mutator assists and background marking.
//
// It calculates the ratio between the allocation rate (in terms of CPU
// time) and the GC scan throughput to determine the heap size at which to
// trigger a GC cycle such that no GC assists are required to finish on time.
// This algorithm thus optimizes GC CPU utilization to the dedicated background
// mark utilization of 25% of GOMAXPROCS by minimizing GC assists.
// GOMAXPROCS. The high-level design of this algorithm is documented
// at https://github.com/golang/proposal/blob/master/design/44167-gc-pacer-redesign.md.
// See https://golang.org/s/go15gcpacing for additional historical context.
