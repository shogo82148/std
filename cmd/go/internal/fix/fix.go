// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix implements the “go fix” command.
package fix

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdFix = &base.Command{
	UsageLine: "go fix [-fix list] [packages]",
	Short:     "update packages to use new APIs",
	Long: `
Fix runs the Go fix command on the packages named by the import paths.

The -fix flag sets a comma-separated list of fixes to run.
The default is all known fixes.
(Its value is passed to 'go tool fix -r'.)

For more about fix, see 'go doc cmd/fix'.
For more about specifying packages, see 'go help packages'.

To run fix with other options, run 'go tool fix'.

See also: go fmt, go vet.
	`,
}
