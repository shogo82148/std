// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package buildinfo provides access to information embedded in a Go binary
// about how it was built. This includes the Go toolchain version, and the
// set of modules used (for binaries built in module mode).
//
// Build information is available for the currently running binary in
// runtime/debug.ReadBuildInfo.
package buildinfo

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/runtime/debug"
)

// Type alias for build info. We cannot move the types here, since
// runtime/debug would need to import this package, which would make it
// a much larger dependency.
type BuildInfo = debug.BuildInfo

// ReadFile returns build information embedded in a Go binary
// file at the given path. Most information is only available for binaries built
// with module support.
func ReadFile(name string) (info *BuildInfo, err error)

// Read returns build information embedded in a Go binary file
// accessed through the given ReaderAt. Most information is only available for
// binaries built with module support.
func Read(r io.ReaderAt) (*BuildInfo, error)

// elfExe is the ELF implementation of the exe interface.

// peExe is the PE (Windows Portable Executable) implementation of the exe interface.

// machoExe is the Mach-O (Apple macOS/iOS) implementation of the exe interface.

// xcoffExe is the XCOFF (AIX eXtended COFF) implementation of the exe interface.
