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
	Path      string
	Main      Module
	Deps      []*Module
	Settings  []BuildSetting
}

// Module represents a module.
type Module struct {
	Path    string
	Version string
	Sum     string
	Replace *Module
}

// BuildSetting describes a setting that may be used to understand how the
// binary was built. For example, VCS commit and dirty status is stored here.
type BuildSetting struct {
	Key, Value string
}

func (bi *BuildInfo) String() string

func ParseBuildInfo(data string) (bi *BuildInfo, err error)
