// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

// ReadMulti reads output (like [SHAKE.Read] would) from each element of s into
// the corresponding element of out. It behaves exactly like
//
//	for i := range s {
//	    s[i].Read(out[i])
//	}
//
// Depending on the platform, groups of states with the same rate, sponge phase,
// and output offset may be processed together, which can be faster than reading
// them one at a time.
//
// ReadMulti panics if len(s) != len(out). The output slices may have different
// lengths, including zero, and must not overlap each other.
func ReadMulti(s []*SHAKE, out [][]byte)
