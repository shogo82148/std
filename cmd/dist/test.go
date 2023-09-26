// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// tester executes cmdtest.

// work tracks command execution for a test.

// A distTest is a test run by dist test.
// Each test has a unique name and belongs to a group (heading)

// goTest represents all options to a "go test" command. The final command will
// combine configuration from goTest and tester flags.

// ranGoTest and stdMatches are state closed over by the stdlib
// testing func in registerStdTest below. The tests are run
// sequentially, so there's no need for locks.
//
// ranGoBench and benchMatches are the same, but are only used
// in -race mode.

// rtSkipFunc is a registerTest option that runs a skip check function before
// running the test.

// cgoPackages is the standard packages that use cgo.
