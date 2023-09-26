// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	cmdgo "cmd/go"
)

// netTestSem is a semaphore limiting the number of tests that may use the
// external network in parallel. If non-nil, it contains one buffer slot per
// test (send to acquire), with a low enough limit that the overall number of
// connections (summed across subprocesses) stays at or below base.NetLimit.

// testGOROOT is the GOROOT to use when running testgo, a cmd/go binary
// build from this process's current GOROOT, but run from a different
// (temp) directory.

// testGOROOT_FINAL is the GOROOT_FINAL with which the test binary is assumed to
// have been built.

// The length of an mtime tick on this system. This is an estimate of
// how long we need to sleep to ensure that the mtime of two files is
// different.
// We used to try to be clever but that didn't always work (see golang.org/issue/12205).

// Manage a single run of the testgo binary.

// If -testwork is specified, the test prints the name of the temp directory
// and does not remove it when done, so that a programmer can
// poke at the test file tree afterward.
