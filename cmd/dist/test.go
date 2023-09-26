// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// tester executes cmdtest.

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

// rtSequential is a registerTest option that causes the registered test to run
// sequentially.

// rtPreFunc is a registerTest option that runs a pre function before running
// the test.

// rtHostTest is a registerTest option that indicates this is a host test that
// should be run using goTest.runHostTest. It implies rtSequential.

// cgoPackages is the standard packages that use cgo.
