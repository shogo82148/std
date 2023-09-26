// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

// testGOROOT is the GOROOT to use when running testgo, a cmd/go binary
// build from this process's current GOROOT, but run from a different
// (temp) directory.

// The length of an mtime tick on this system. This is an estimate of
// how long we need to sleep to ensure that the mtime of two files is
// different.
// We used to try to be clever but that didn't always work (see golang.org/issue/12205).

// Manage a single run of the testgo binary.

// The runtime version string takes one of two forms:
// "go1.X[.Y]" for Go releases, and "devel +hash" at tip.
// Determine whether we are in a released copy by
// inspecting the version.
