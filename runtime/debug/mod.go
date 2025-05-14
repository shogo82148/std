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
	// GoVersion is the version of the Go toolchain that built the binary
	// (for example, "go1.19.2").
	GoVersion string `json:",omitempty"`

	// Path is the package path of the main package for the binary
	// (for example, "golang.org/x/tools/cmd/stringer").
	Path string `json:",omitempty"`

	// Main describes the module that contains the main package for the binary.
	Main Module `json:""`

	// Deps describes all the dependency modules, both direct and indirect,
	// that contributed packages to the build of this binary.
	Deps []*Module `json:",omitempty"`

	// Settings describes the build settings used to build the binary.
	Settings []BuildSetting `json:",omitempty"`
}

// A Module describes a single module included in a build.
type Module struct {
	Path    string  `json:",omitempty"`
	Version string  `json:",omitempty"`
	Sum     string  `json:",omitempty"`
	Replace *Module `json:",omitempty"`
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
//   - DefaultGODEBUG: the effective GODEBUG settings
//   - GOARCH: the architecture target
//   - GOAMD64/GOARM/GO386/etc: the architecture feature level for GOARCH
//   - GOOS: the operating system target
//   - GOFIPS140: the frozen FIPS 140-3 module version, if any
//   - vcs: the version control system for the source tree where the build ran
//   - vcs.revision: the revision identifier for the current commit or checkout
//   - vcs.time: the modification time associated with vcs.revision, in RFC3339 format
//   - vcs.modified: true or false indicating whether the source tree had local modifications
type BuildSetting struct {
	// Key and Value describe the build setting.
	// Key must not contain an equals sign, space, tab, or newline.
	Key string `json:",omitempty"`
	// Value must not contain newlines ('\n').
	Value string `json:",omitempty"`
}

// String returns a string representation of a [BuildInfo].
func (bi *BuildInfo) String() string

// ParseBuildInfo parses the string returned by [*BuildInfo.String],
// restoring the original BuildInfo,
// except that the GoVersion field is not set.
// Programs should normally not call this function,
// but instead call [ReadBuildInfo], [debug/buildinfo.ReadFile],
// or [debug/buildinfo.Read].
func ParseBuildInfo(data string) (bi *BuildInfo, err error)
