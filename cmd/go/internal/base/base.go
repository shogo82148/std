// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base defines shared basic pieces of the go command,
// in particular logging and the Command structure.
package base

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/flag"
)

// A Command is an implementation of a go command
// like go build or go fix.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(ctx context.Context, cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The words between "go" and the first flag or argument in the line are taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool

	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	Commands []*Command
}

var Go = &Command{
	UsageLine: "go",
	Long:      `Go is a tool for managing Go source code.`,
}

// Lookup returns the subcommand with the given name, if any.
// Otherwise it returns nil.
//
// Lookup ignores subcommands that have len(c.Commands) == 0 and c.Run == nil.
// Such subcommands are only for use as arguments to "help".
func (c *Command) Lookup(name string) *Command

// LongName returns the command's long name: all the words in the usage line between "go" and a flag or argument,
func (c *Command) LongName() string

// Name returns the command's short name: the last word in the usage line before a flag or argument.
func (c *Command) Name() string

func (c *Command) Usage()

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool

func AtExit(f func())

func Exit()

func Fatalf(format string, args ...any)

func Errorf(format string, args ...any)

func ExitIfErrors()

func Error(err error)

func Fatal(err error)

func SetExitStatus(n int)

func GetExitStatus() int

// Run runs the command, with stdout and stderr
// connected to the go command's own stdout and stderr.
// If the command fails, Run reports the error using Errorf.
func Run(cmdargs ...any)

// RunStdin is like run but connects Stdin. It retries if it encounters an ETXTBSY.
func RunStdin(cmdline []string)

// Usage is the usage-reporting function, filled in by package main
// but here for reference by other packages.
var Usage func()
