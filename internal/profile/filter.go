// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implements methods to filter samples from profiles.

package profile

import "github.com/shogo82148/std/regexp"

// FilterSamplesByName filters the samples in a profile and only keeps
// samples where at least one frame matches focus but none match ignore.
// Returns true is the corresponding regexp matched at least one sample.
func (p *Profile) FilterSamplesByName(focus, ignore, hide *regexp.Regexp) (fm, im, hm bool)

// TagMatch selects tags for filtering
type TagMatch func(key, val string, nval int64) bool

// FilterSamplesByTag removes all samples from the profile, except
// those that match focus and do not match the ignore regular
// expression.
func (p *Profile) FilterSamplesByTag(focus, ignore TagMatch) (fm, im bool)
