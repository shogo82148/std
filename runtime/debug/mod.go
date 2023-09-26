// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

// ReadBuildInfo returns the build information embedded
// in the running binary. The information is available only
// in binaries built with module support.
func ReadBuildInfo() (info *BuildInfo, ok bool)

// BuildInfo represents the build information read from
// the running binary.
type BuildInfo struct {
	Path string
	Main Module
	Deps []*Module
}

// Module represents a module.
type Module struct {
	Path    string
	Version string
	Sum     string
	Replace *Module
}
