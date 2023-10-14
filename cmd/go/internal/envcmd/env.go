// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package envcmd implements the “go env” command.
package envcmd

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/cmd/go/internal/base"
	"github.com/shogo82148/std/cmd/go/internal/cfg"
)

var CmdEnv = &base.Command{
	UsageLine: "go env [-json] [-u] [-w] [var ...]",
	Short:     "print Go environment information",
	Long: `
Env prints Go environment information.

By default env prints information as a shell script
(on Windows, a batch file). If one or more variable
names is given as arguments, env prints the value of
each named variable on its own line.

The -json flag prints the environment in JSON format
instead of as a shell script.

The -u flag requires one or more arguments and unsets
the default setting for the named environment variables,
if one has been set with 'go env -w'.

The -w flag requires one or more arguments of the
form NAME=VALUE and changes the default settings
of the named environment variables to the given values.

For more about environment variables, see 'go help environment'.
	`,
}

func MkEnv() []cfg.EnvVar

// ExtraEnvVars returns environment variables that should not leak into child processes.
func ExtraEnvVars() []cfg.EnvVar

// ExtraEnvVarsCostly returns environment variables that should not leak into child processes
// but are costly to evaluate.
func ExtraEnvVarsCostly() []cfg.EnvVar

// PrintEnv prints the environment variables to w.
func PrintEnv(w io.Writer, env []cfg.EnvVar)
