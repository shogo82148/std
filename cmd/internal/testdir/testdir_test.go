// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Individual test cases in GOROOT/test are already responsible for setting -race flag
// and other flags to suit their needs, and that flag doesn't automatically propagate
// to the individual test cases anyway.
//
// So, avoid repeating these tests in race mode to avoid doing duplicate work.
// In particular, this saves time on the race builders.
//
//go:build !race

// Package testdir_test runs tests in the GOROOT/test directory.
package testdir_test
