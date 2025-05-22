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

// CPU owns the FDs required to read the CPU limit from a cgroup.
type CPU struct {
	version Version

	// For cgroup v1, this is cpu.cfs_quota_us.
	// For cgroup v2, this is cpu.max.
	quotaFD int

	// For cgroup v1, this is cpu.cfs_period_us.
	// For cgroup v2, this is unused.
	periodFD int
}

func (c CPU) Close()

// OpenCPU returns a CPU for the CPU cgroup containing the current process, or
// ErrNoCgroup if the process is not in a CPU cgroup.
//
// scratch must have length ScratchSize.
func OpenCPU(scratch []byte) (CPU, error)

// Returns average CPU throughput limit from the cgroup, or ok false if there
// is no limit.
func ReadCPULimit(c CPU) (float64, bool, error)

// FindCPU finds the path to the CPU cgroup that this process is a member of
// and places it in out. scratch is a scratch buffer for internal use.
//
// out must have length PathSize. scratch must have length ParseSize.
//
// Returns the number of bytes written to out and the cgroup version (1 or 2).
//
// Returns ErrNoCgroup if the process is not in a CPU cgroup.
func FindCPU(out []byte, scratch []byte) (int, Version, error)

// FindCPURelativePath finds the path to the CPU cgroup that this process is a member of
// relative to the root of the cgroup mount and places it in out. scratch is a
// scratch buffer for internal use.
//
// out must have length PathSize minus the size of the cgroup mount root (if
// known). scratch must have length ParseSize.
//
// Returns the number of bytes written to out and the cgroup version (1 or 2).
//
// Returns ErrNoCgroup if the process is not in a CPU cgroup.
func FindCPURelativePath(out []byte, scratch []byte) (int, Version, error)

// FindCPUMountPoint finds the root of the CPU cgroup mount places it in out.
// scratch is a scratch buffer for internal use.
//
// out must have length PathSize. scratch must have length ParseSize.
//
// Returns the number of bytes written to out.
//
// Returns ErrNoCgroup if the process is not in a CPU cgroup.
func FindCPUMountPoint(out []byte, scratch []byte) (int, error)
