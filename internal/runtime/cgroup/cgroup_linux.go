// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgroup

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

// FindCPUCgroup finds the path to the CPU cgroup that this process is a member of
// and places it in out. scratch is a scratch buffer for internal use.
//
// out must have length PathSize. scratch must have length ParseSize.
//
// Returns the number of bytes written to out and the cgroup version (1 or 2).
//
// Returns ErrNoCgroup if the process is not in a CPU cgroup.
func FindCPUCgroup(out []byte, scratch []byte) (int, Version, error)

// FindCPUMountPoint finds the mount point containing the specified cgroup and
// version with cpu controller, and compose the full path to the cgroup in out.
// scratch is a scratch buffer for internal use.
//
// out must have length PathSize, may overlap with cgroup.
// scratch must have length ParseSize.
//
// Returns the number of bytes written to out.
//
// Returns ErrNoCgroup if no matching mount point is found.
func FindCPUMountPoint(out, cgroup []byte, version Version, scratch []byte) (int, error)
