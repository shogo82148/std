// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fmtcmd implements the “go fmt” command.
package fmtcmd

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdFmt = &base.Command{
	Run:       runFmt,
	UsageLine: "go fmt [-n] [-x] [packages]",
	Short:     "gofmt (reformat) package sources",
	Long: `
Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths. It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

The -mod flag's value sets which module download mode
to use: readonly or vendor. See 'go help modules' for more.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
	`,
}
