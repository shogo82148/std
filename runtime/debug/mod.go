// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

// ReadBuildInfo returns the build information embedded
// in the running binary. The information is available only
// in binaries built with module support.
func ReadBuildInfo() (info *BuildInfo, ok bool)

// BuildInfo represents the build information read from a Go binary.
type BuildInfo struct {
	GoVersion string

	Path string

	Main Module

	Deps []*Module

	Settings []BuildSetting
}

// A Module describes a single module included in a build.
type Module struct {
	Path    string
	Version string
	Sum     string
	Replace *Module
}

// A BuildSetting is a key-value pair describing one setting that influenced a build.
//
// Defined keys include:
//
//   - -buildmode: the buildmode flag used (typically "exe")
//   - -compiler: the compiler toolchain flag used (typically "gc")
//   - CGO_ENABLED: the effective CGO_ENABLED environment variable
//   - CGO_CFLAGS: the effective CGO_CFLAGS environment variable
//   - CGO_CPPFLAGS: the effective CGO_CPPFLAGS environment variable
//   - CGO_CXXFLAGS:  the effective CGO_CXXFLAGS environment variable
//   - CGO_LDFLAGS: the effective CGO_LDFLAGS environment variable
//   - GOARCH: the architecture target
//   - GOAMD64/GOARM/GO386/etc: the architecture feature level for GOARCH
//   - GOOS: the operating system target
//   - vcs: the version control system for the source tree where the build ran
//   - vcs.revision: the revision identifier for the current commit or checkout
//   - vcs.time: the modification time associated with vcs.revision, in RFC3339 format
//   - vcs.modified: true or false indicating whether the source tree had local modifications
type BuildSetting struct {
	Key, Value string
}

func (bi *BuildInfo) String() string

func ParseBuildInfo(data string) (bi *BuildInfo, err error)
