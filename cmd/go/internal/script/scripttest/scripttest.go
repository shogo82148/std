// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scripttest adapts the script engine for use in tests.
package scripttest

import (
	"github.com/shogo82148/std/cmd/go/internal/script"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/testing"
)

// DefaultCmds returns a set of broadly useful script commands.
//
// This set includes all of the commands in script.DefaultCmds,
// as well as a "skip" command that halts the script and causes the
// testing.TB passed to Run to be skipped.
func DefaultCmds() map[string]script.Cmd

// DefaultConds returns a set of broadly useful script conditions.
//
// This set includes all of the conditions in script.DefaultConds,
// as well as:
//
//   - Conditions of the form "exec:foo" are active when the executable "foo" is
//     found in the test process's PATH, and inactive when the executable is
//     not found.
//
//   - "short" is active when testing.Short() is true.
//
//   - "verbose" is active when testing.Verbose() is true.
func DefaultConds() map[string]script.Cond

// Run runs the script from the given filename starting at the given initial state.
// When the script completes, Run closes the state.
func Run(t testing.TB, e *script.Engine, s *script.State, filename string, testScript io.Reader)

// Skip returns a sentinel error that causes Run to mark the test as skipped.
func Skip() script.Cmd

// CachedExec returns a Condition that reports whether the PATH of the test
// binary itself (not the script's current environment) contains the named
// executable.
func CachedExec() script.Cond
