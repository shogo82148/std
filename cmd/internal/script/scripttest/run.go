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

// RunToolScriptTest kicks off a set of script tests runs for
// a tool of some sort (compiler, linker, etc). The expectation
// is that we'll be called from the top level cmd/X dir for tool X,
// and that instead of executing the install tool X we'll use the
// test binary instead.
func RunToolScriptTest(t *testing.T, repls []ToolReplacement, pattern string)

// RunTests kicks off one or more script-based tests using the
// specified engine, running all test files that match pattern.
// This function adapted from Russ's rsc.io/script/scripttest#Run
// function, which was in turn forked off cmd/go's runner.
func RunTests(t *testing.T, ctx context.Context, engine *script.Engine, env []string, pattern string)
