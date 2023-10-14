// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vet implements the “go vet” command.
package vet

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdVet = &base.Command{
	CustomFlags: true,
	UsageLine:   "go vet [build flags] [-vettool prog] [vet flags] [packages]",
	Short:       "report likely mistakes in packages",
	Long: `
Vet runs the Go vet command on the packages named by the import paths.

For more about vet and its flags, see 'go doc cmd/vet'.
For more about specifying packages, see 'go help packages'.
For a list of checkers and their flags, see 'go tool vet help'.
For details of a specific checker such as 'printf', see 'go tool vet help printf'.

The -vettool=prog flag selects a different analysis tool with alternative
or additional checks.
For example, the 'shadow' analyzer can be built and run using these commands:

  go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
  go vet -vettool=$(which shadow)

The build flags supported by go vet are those that control package resolution
and execution, such as -C, -n, -x, -v, -tags, and -toolexec.
For more about these flags, see 'go help build'.

See also: go fmt, go fix.
	`,
}
