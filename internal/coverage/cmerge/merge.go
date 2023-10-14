// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmerge

import (
	"github.com/shogo82148/std/internal/coverage"
)

type ModeMergePolicy uint8

const (
	ModeMergeStrict ModeMergePolicy = iota
	ModeMergeRelaxed
)

// Merger provides state and methods to help manage the process of
// merging together coverage counter data for a given function, for
// tools that need to implicitly merge counter as they read multiple
// coverage counter data files.
type Merger struct {
	cmode    coverage.CounterMode
	cgran    coverage.CounterGranularity
	policy   ModeMergePolicy
	overflow bool
}

func (cm *Merger) SetModeMergePolicy(policy ModeMergePolicy)

// MergeCounters takes the counter values in 'src' and merges them
// into 'dst' according to the correct counter mode.
func (m *Merger) MergeCounters(dst, src []uint32) (error, bool)

// Saturating add does a saturating addition of 'dst' and 'src',
// returning added value or math.MaxUint32 if there is an overflow.
// Overflows are recorded in case the client needs to track them.
func (m *Merger) SaturatingAdd(dst, src uint32) uint32

// Saturating add does a saturating addition of 'dst' and 'src',
// returning added value or math.MaxUint32 plus an overflow flag.
func SaturatingAdd(dst, src uint32) (uint32, bool)

// SetModeAndGranularity records the counter mode and granularity for
// the current merge. In the specific case of merging across coverage
// data files from different binaries, where we're combining data from
// more than one meta-data file, we need to check for and resolve
// mode/granularity clashes.
func (cm *Merger) SetModeAndGranularity(mdf string, cmode coverage.CounterMode, cgran coverage.CounterGranularity) error

func (cm *Merger) ResetModeAndGranularity()

func (cm *Merger) Mode() coverage.CounterMode

func (cm *Merger) Granularity() coverage.CounterGranularity
