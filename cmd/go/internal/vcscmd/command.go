// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vcscmd constructs VCS command lines, separating constant command
// words from option values and positional arguments supplied at run time.
package vcscmd

func New(bin lit, words ...lit) *Command

// Command provides methods that record intent and help
// prevent building unsafe argv for VCS sub-process calls.
type Command struct {
	argv      []string
	separated bool
	err       error
}

// Add appends constant command words.
//
// It is an error to specify "--" or "--end-of-options".
func (c *Command) Add(words ...lit) *Command

// Option appends a long option in the form --name=value.
//
// If [New] was invoked with "fossil", it also accepts "-r"
// written as "-r=value" and "-R" followed by an absolute
// repository path.
func (c *Command) Option(name lit, value string) *Command

// Operands appends args and places the correct end of arguments
// separator.
//
// For git commands that parse a revision followed by paths, use
// [Command.GitRevision].
func (c *Command) Operands(args ...string) *Command

// GitRevision allows safe specification of paths in a rev.
//
// It is an error to use GitRevision for non-git commands.
func (c *Command) GitRevision(rev string, paths ...string) *Command

// Build returns the full argv with VCS binary in argv[0].
//
// If Build returns an error, it is unsafe to use [Command].
func (c *Command) Build() ([]string, error)
