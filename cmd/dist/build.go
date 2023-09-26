// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// The usual variables.

// The known architectures.

// The known operating systems.

// The old tools that no longer live in $GOBIN or $GOROOT/bin.

// Unreleased directories (relative to $GOROOT) that should
// not be in release branches.

// deptab lists changes to the default dependencies for a given prefix.
// deps ending in /* read the whole directory; deps beginning with -
// exclude files with that prefix.

// depsuffix records the allowed suffixes for source files.

// gentab records how to generate some trivial files.

// buildorder records the order of builds for the 'go bootstrap' command.
// The Go packages and commands must be in dependency order,
// maintained by hand, but the order doesn't change often.

// Copied from go/build/build.go.
// Cannot use go/build directly because cmd/dist for a new release
// builds against an old release's go/build, which may be out of sync.
