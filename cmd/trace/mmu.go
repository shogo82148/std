// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Minimum mutator utilization (MMU) graphing.

// TODO:
//
// In worst window list, show break-down of GC utilization sources
// (STW, assist, etc). Probably requires a different MutatorUtil
// representation.
//
// When a window size is selected, show a second plot of the mutator
// utilization distribution for that window size.
//
// Render plot progressively so rough outline is visible quickly even
// for very complex MUTs. Start by computing just a few window sizes
// and then add more window sizes.
//
// Consider using sampling to compute an approximate MUT. This would
// work by sampling the mutator utilization at randomly selected
// points in time in the trace to build an empirical distribution. We
// could potentially put confidence intervals on these estimates and
// render this progressively as we refine the distributions.

package main
