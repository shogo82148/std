// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// The usual variables.

// The known architectures.

// The known operating systems.

// clangos lists the operating systems where we prefer clang to gcc.

// The old tools that no longer live in $GOBIN or $GOROOT/bin.

// Unreleased directories (relative to $GOROOT) that should
// not be in release branches.

// depsuffix records the allowed suffixes for source files.

// gentab records how to generate some trivial files.
// Files listed here should also be listed in ../distpack/pack.go's srcArch.Remove list.

// installed maps from a dir name (as given to install) to a chan
// closed when the dir's package is installed.

// unixOS is the set of GOOS values matched by the "unix" build tag.
// This is the same list as in go/build/syslist.go and
// cmd/go/internal/imports/build.go.

// Cannot use go/build directly because cmd/dist for a new release
// builds against an old release's go/build, which may be out of sync.
// To reduce duplication, we generate the list for go/build from this.
//
// We list all supported platforms in this list, so that this is the
// single point of truth for supported platforms. This list is used
// by 'go tool dist list'.

// List of platforms that are marked as broken ports.
// These require -force flag to build, and also
// get filtered out of cgoEnabled for 'dist list'.
// See go.dev/issue/56679.

// List of platforms which are first class ports. See go.dev/issue/38874.
