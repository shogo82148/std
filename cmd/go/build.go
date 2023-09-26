// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Flags set by multiple commands.

// buildCompiler implements flag.Var.
// It implements Set by updating both
// buildToolchain and buildContext.Compiler.

// Global build parameters (used during package load)

// A builder holds global state about a build.
// It does not hold per-package state, because we
// build packages in parallel, and the builder is shared.

// An action represents a single action in the action graph.

// cacheKey is the key for the action cache.

// buildMode specifies the build mode:
// are we just building things or also installing the results?

// errPrintedOutput is a special error indicating that a command failed
// but that it generated output as well, and that output has already
// been printed, so there's no point showing 'exit status 1' or whatever
// the wait status was. The main executor, builder.do, knows not to
// print this error.

// The Go toolchain.

// The Gccgo toolchain.

// Make sure SWIG is new enough.

// Find the value to pass for the -intgosize option to swig.

// This code fails to build if sizeof(int) <= 32

// An actionQueue is a priority queue of actions.
