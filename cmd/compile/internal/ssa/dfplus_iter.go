// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/iter"
)

// IterDomFrontierPlus iterates the DF+ of seeds: every block at which
// a phi may need to be placed if a variable were defined in the seed
// blocks. Blocks are yielded at most once, in a deterministic order;
// an early break stops the walk. Seed blocks themselves are not
// yielded as such, but a seed that is also a merge point (e.g. a loop
// header) is.
// seeds iterator is consumed in full before the walk starts (the current
// algorithm has to walk deeper roots first).
// CFG must not change while iteration is in progress; inserting
// values (like phis) is fine.
func (f *Func) IterDomFrontierPlus(seeds iter.Seq[*Block]) iter.Seq[*Block]
