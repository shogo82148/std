// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package liveness

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
)

// MergeLocalsState encapsulates information about which AUTO
// (stack-allocated) variables within a function can be safely
// merged/overlapped, e.g. share a stack slot with some other auto).
// An instance of MergeLocalsState is produced by MergeLocals() below
// and then consumed in ssagen.AllocFrame. The map 'partition'
// contains entries of the form <N,SL> where N is an *ir.Name and SL
// is a slice holding the indices (within 'vars') of other variables
// that share the same slot, specifically the slot of the first
// element in the partition, which we'll call the "leader". For
// example, if a function contains five variables where v1/v2/v3 are
// safe to overlap and v4/v5 are safe to overlap, the MergeLocalsState
// content might look like
//
//	vars: [v1, v2, v3, v4, v5]
//	partition: v1 -> [1, 0, 2], v2 -> [1, 0, 2], v3 -> [1, 0, 2]
//	           v4 -> [3, 4], v5 -> [3, 4]
//
// A nil MergeLocalsState indicates that no local variables meet the
// necessary criteria for overlap.
type MergeLocalsState struct {
	// contains auto vars that participate in overlapping
	vars []*ir.Name
	// maps auto variable to overlap partition
	partition map[*ir.Name][]int
}

// MergeLocals analyzes the specified ssa function f to determine which
// of its auto variables can safely share the same stack slot, returning
// a state object that describes how the overlap should be done.
func MergeLocals(fn *ir.Func, f *ssa.Func) *MergeLocalsState

// Subsumed returns whether variable n is subsumed, e.g. appears
// in an overlap position but is not the leader in that partition.
func (mls *MergeLocalsState) Subsumed(n *ir.Name) bool

// IsLeader returns whether a variable n is the leader (first element)
// in a sharing partition.
func (mls *MergeLocalsState) IsLeader(n *ir.Name) bool

// Leader returns the leader variable for subsumed var n.
func (mls *MergeLocalsState) Leader(n *ir.Name) *ir.Name

// Followers writes a list of the followers for leader n into the slice tmp.
func (mls *MergeLocalsState) Followers(n *ir.Name, tmp []*ir.Name) []*ir.Name

// EstSavings returns the estimated reduction in stack size (number of bytes) for
// the given merge locals state via a pair of ints, the first for non-pointer types and the second for pointer types.
func (mls *MergeLocalsState) EstSavings() (int, int)

func (mls *MergeLocalsState) String() string

// for unit testing only.
func MakeMergeLocalsState(partition map[*ir.Name][]int, vars []*ir.Name) (*MergeLocalsState, error)
