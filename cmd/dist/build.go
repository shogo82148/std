// Copyright 2012 The Go Authors. All rights reserved.
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
// Note that this table applies only to the build of cmd/go,
// after the main compiler bootstrap.

// depsuffix records the allowed suffixes for source files.

// gentab records how to generate some trivial files.

// installed maps from a dir name (as given to install) to a chan
// closed when the dir's package is installed.

// cleanlist is a list of packages with generated files and commands.

// Cannot use go/build directly because cmd/dist for a new release
// builds against an old release's go/build, which may be out of sync.
// To reduce duplication, we generate the list for go/build from this.
//
// We list all supported platforms in this list, so that this is the
// single point of truth for supported platforms. This list is used
// by 'go tool dist list'.
