// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp

// A onePassProg is a compiled one-pass regular expression program.
// It is the same as syntax.Prog except for the use of onePassInst.

// A onePassInst is a single instruction in a one-pass regular expression program.
// It is the same as syntax.Inst except for the new 'Next' field.

// Sparse Array implementation is used as a queueOnePass.

// mergeRuneSets merges two non-intersecting runesets, and returns the merged result,
// and a NextIp array. The idea is that if a rune matches the OnePassRunes at index
// i, NextIp[i/2] is the target. If the input sets intersect, an empty runeset and a
// NextIp array with the single element mergeFailed is returned.
// The code assumes that both inputs contain ordered and non-intersecting rune pairs.

// runeSlice exists to permit sorting the case-folded rune sets.
