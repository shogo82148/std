// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

// Merge merges all the profiles in profs into a single Profile.
// Returns a new profile independent of the input profiles. The merged
// profile is compacted to eliminate unused samples, locations,
// functions and mappings. Profiles must have identical profile sample
// and period types or the merge will fail. profile.Period of the
// resulting profile will be the maximum of all profiles, and
// profile.TimeNanos will be the earliest nonzero one.
func Merge(srcs []*Profile) (*Profile, error)

// Normalize normalizes the source profile by multiplying each value in profile by the
// ratio of the sum of the base profile's values of that sample type to the sum of the
// source profile's value of that sample type.
func (p *Profile) Normalize(pb *Profile) error
