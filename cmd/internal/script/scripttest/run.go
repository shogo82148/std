// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scripttest adapts the script engine for use in tests.
package scripttest

import (
	"github.com/shogo82148/std/cmd/internal/script"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/testing"
)

// ToolReplacement records the name of a tool to replace
// within a given GOROOT for script testing purposes.
type ToolReplacement struct {
	ToolName        string
	ReplacementPath string
	EnvVar          string
}

// NewEngine constructs a new [script.Engine] and environment to be used with
// [RunTests].
func NewEngine(t *testing.T, repls []ToolReplacement) (*script.Engine, []string)

// RunToolScriptTest kicks off a set of script tests runs for
// a tool of some sort (compiler, linker, etc). The expectation
// is that we'll be called from the top level cmd/X dir for tool X,
// and that instead of executing the install tool X we'll use the
// test binary instead.
func RunToolScriptTest(t *testing.T, repls []ToolReplacement, scriptsdir string, fixReadme bool)

// ScriptTestContext returns a context with a grace period for cleaning up
// subprocesses of a script test.
//
// When we run commands that execute subprocesses, we want to reserve two grace
// periods to clean up. We will send the first termination signal when the
// context expires, then wait one grace period for the process to produce
// whatever useful output it can (such as a stack trace). After the first grace
// period expires, we'll escalate to os.Kill, leaving the second grace period
// for the test function to record its output before the test process itself
// terminates.
//
// The grace period is 100ms or 5% of the time remaining until
// [testing.T.Deadline], whichever is greater.
func ScriptTestContext(t *testing.T, ctx context.Context) context.Context

// RunTests kicks off one or more script-based tests using the
// specified engine, running all test files that match pattern.
// This function adapted from Russ's rsc.io/script/scripttest#Run
// function, which was in turn forked off cmd/go's runner.
func RunTests(t *testing.T, ctx context.Context, engine *script.Engine, env []string, pattern string)

// InitScriptDirs sets up directories for executing a script test.
//
//   - WORK (env var) is set to the current working directory.
//   - TMPDIR (env var; TMP on Windows) is set to $WORK/tmp.
//   - $TMPDIR is created.
func InitScriptDirs(t testing.TB, s *script.State)
