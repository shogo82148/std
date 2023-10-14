// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/shogo82148/std/flag"
)

// A StringsFlag is a command-line flag that interprets its argument
// as a space-separated list of possibly-quoted strings.
type StringsFlag []string

func (v *StringsFlag) Set(s string) error

func (v *StringsFlag) String() string

// AddBuildFlagsNX adds the -n and -x build flags to the flag set.
func AddBuildFlagsNX(flags *flag.FlagSet)

// AddChdirFlag adds the -C flag to the flag set.
func AddChdirFlag(flags *flag.FlagSet)

// AddModFlag adds the -mod build flag to the flag set.
func AddModFlag(flags *flag.FlagSet)

// AddModCommonFlags adds the module-related flags common to build commands
// and 'go mod' subcommands.
func AddModCommonFlags(flags *flag.FlagSet)

func ChdirFlag(s string) error
