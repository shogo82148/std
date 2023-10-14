// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package version implements the “go version” command.
package version

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdVersion = &base.Command{
	UsageLine: "go version [-m] [-v] [file ...]",
	Short:     "print Go version",
	Long: `Version prints the build information for Go binary files.

Go version reports the Go version used to build each of the named files.

If no files are named on the command line, go version prints its own
version information.

If a directory is named, go version walks that directory, recursively,
looking for recognized Go binaries and reporting their versions.
By default, go version does not report unrecognized files found
during a directory scan. The -v flag causes it to report unrecognized files.

The -m flag causes go version to print each file's embedded
module version information, when available. In the output, the module
information consists of multiple lines following the version line, each
indented by a leading tab character.

See also: go doc runtime/debug.BuildInfo.
`,
}
