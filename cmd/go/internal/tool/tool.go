// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tool implements the “go tool” command.
package tool

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdTool = &base.Command{
	Run:       runTool,
	UsageLine: "go tool [-n] command [args...]",
	Short:     "run specified go tool",
	Long: `
Tool runs the go tool command identified by the arguments.

Go ships with a number of builtin tools, and additional tools
may be defined in the go.mod of the current module. 'go get -tool'
can be used to define additional tools in the current module's
go.mod file. See 'go help get' for more information.

The command can be specified using the full package path to the tool declared with
a tool directive. The default binary name of the tool, which is the last component of
the package path, excluding the major version suffix, can also be used if it is unique
among declared tools.

With no arguments it prints the list of known tools.

The -n flag causes tool to print the command that would be
executed but not execute it.

The -modfile=file.mod build flag causes tool to use an alternate file
instead of the go.mod in the module root directory.

Tool also provides the -C, -overlay, and -modcacherw build flags.

The go command places $GOROOT/bin at the beginning of $PATH in the
environment of commands run via tool directives, so that they use the
same 'go' as the parent 'go tool'.

For more about build flags, see 'go help build'.

For more about each builtin tool command, see 'go doc cmd/<command>'.
`,
}
