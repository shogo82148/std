// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements parsers to convert legacy profiles into the
// profile.proto format.

package profile

import (
	"github.com/shogo82148/std/io"
)

var (

	// LegacyHeapAllocated instructs the heapz parsers to use the
	// allocated memory stats instead of the default in-use memory. Note
	// that tcmalloc doesn't provide all allocated memory, only in-use
	// stats.
	LegacyHeapAllocated bool
)

// ParseTracebacks parses a set of tracebacks and returns a newly
// populated profile. It will accept any text file and generate a
// Profile out of it with any hex addresses it can identify, including
// a process map if it can recognize one. Each sample will include a
// tag "source" with the addresses recognized in string format.
func ParseTracebacks(b []byte) (*Profile, error)

// ParseMemoryMap parses a memory map in the format of
// /proc/self/maps, and overrides the mappings in the current profile.
// It renumbers the samples and locations in the profile correspondingly.
func (p *Profile) ParseMemoryMap(rd io.Reader) error
