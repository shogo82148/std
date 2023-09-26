// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/flag"
)

// A Command is an implementation of a go command
// like go build or go fix.
type Command struct {
	Run func(cmd *Command, args []string)

	UsageLine string

	Short string

	Long string

	Flag flag.FlagSet

	CustomFlags bool
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string

func (c *Command) Usage()

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'go help'.
