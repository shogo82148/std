// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stats provides a convenient way to collect compiler statistics.
package stats

type Stats struct {
	stats []*PrefixStats
}

func (s *Stats) Merge(other *Stats)

func (s *Stats) Print()

func (s *Stats) NewPrefixStat(p string) *PrefixStats

// Stats holds a collection of statistics.
type PrefixStats struct {
	prefix string
	stats  map[string]int64
}

// Record records a value for a given statistic.
func (s *PrefixStats) Record(name string, value int64)

// Merge merges another Stats object into this one.
func (s *PrefixStats) Merge(other *PrefixStats)

// Print prints the collected statistics to the console.
func (s *PrefixStats) Print()
