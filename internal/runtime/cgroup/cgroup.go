// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgroup

var (
	ErrNoCgroup error = stringError("not in a cgroup")
)

const (
	// Required amount of scratch space for CPULimit.
	//
	// TODO(prattmic): This is shockingly large (~70KiB) due to the (very
	// unlikely) combination of extremely long paths consisting mostly
	// escaped characters. The scratch buffer ends up in .bss in package
	// runtime, so it doesn't contribute to binary size and generally won't
	// be faulted in, but it would still be nice to shrink this. A more
	// complex parser that did not need to keep entire lines in memory
	// could get away with much less. Alternatively, we could do a one-off
	// mmap allocation for this buffer, which is only mapped larger if we
	// actually need the extra space.
	ScratchSize = PathSize + ParseSize

	// Required space to store a path of the cgroup in the filesystem.
	PathSize = _PATH_MAX

	// Required space to parse /proc/self/mountinfo and /proc/self/cgroup.
	// See findCPUMount and findCPURelativePath.
	ParseSize = 4 * escapedPathMax
)

// Version indicates the cgroup version.
type Version int

const (
	VersionUnknown Version = iota
	V1
	V2
)
