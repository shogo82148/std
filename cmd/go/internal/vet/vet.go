// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vet implements the “go vet” and “go fix” commands.
package vet

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdVet = &base.Command{
	CustomFlags: true,
	UsageLine:   "go vet [build flags] [-vettool prog] [vet flags] [packages]",
	Short:       "report likely mistakes in packages",
	Long: `
Vet runs the Go vet tool (cmd/vet) on the named packages
and reports diagnostics.

It supports these flags:

  -c int
	display offending line with this many lines of context (default -1)
  -json
	emit JSON output
  -fix
	instead of printing each diagnostic, apply its first fix (if any)
  -diff
	instead of applying each fix, print the patch as a unified diff

The -vettool=prog flag selects a different analysis tool with
alternative or additional checks. For example, the 'shadow' analyzer
can be built and run using these commands:

  go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
  go vet -vettool=$(which shadow)

Alternative vet tools should be built atop golang.org/x/tools/go/analysis/unitchecker,
which handles the interaction with go vet.

The default vet tool is 'go tool vet' or cmd/vet.
For help on its checkers and their flags, run 'go tool vet help'.
For details of a specific checker such as 'printf', see 'go tool vet help printf'.

For more about specifying packages, see 'go help packages'.

The build flags supported by go vet are those that control package resolution
and execution, such as -C, -n, -x, -v, -tags, and -toolexec.
For more about these flags, see 'go help build'.

See also: go fmt, go fix.
	`,
}

var CmdFix = &base.Command{
	CustomFlags: true,
	UsageLine:   "go fix [build flags] [-fixtool prog] [fix flags] [packages]",
	Short:       "apply fixes suggested by static checkers",
	Long: `
Fix runs the Go fix tool (cmd/fix) on the named packages
and applies suggested fixes.

It supports these flags:

  -diff
	instead of applying each fix, print the patch as a unified diff

The -fixtool=prog flag selects a different analysis tool with
alternative or additional fixers; see the documentation for go vet's
-vettool flag for details.

The default fix tool is 'go tool fix' or cmd/fix.
For help on its fixers and their flags, run 'go tool fix help'.
For details of a specific fixer such as 'hostport', see 'go tool fix help hostport'.

For more about specifying packages, see 'go help packages'.

The build flags supported by go fix are those that control package resolution
and execution, such as -C, -n, -x, -v, -tags, and -toolexec.
For more about these flags, see 'go help build'.

See also: go fmt, go vet.
	`,
}
