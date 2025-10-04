// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cgrouptest provides best-effort helpers for running tests inside a
// cgroup.
package cgrouptest

import (
	"github.com/shogo82148/std/testing"
)

type CgroupV2 struct {
	orig string
	path string
}

func (c *CgroupV2) Path() string

// Path to cpu.max.
func (c *CgroupV2) CPUMaxPath() string

// Set cpu.max. Pass -1 for quota to disable the limit.
func (c *CgroupV2) SetCPUMax(quota, period int64) error

// InCgroupV2 creates a new v2 cgroup, migrates the current process into it,
// and then calls fn. When fn returns, the current process is migrated back to
// the original cgroup and the new cgroup is destroyed.
//
// If a new cgroup cannot be created, the test is skipped.
//
// This must not be used in parallel tests, as it affects the entire process.
func InCgroupV2(t *testing.T, fn func(*CgroupV2))
