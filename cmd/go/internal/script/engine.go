// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package script implements a small, customizable, platform-agnostic scripting
// language.
//
// Scripts are run by an [Engine] configured with a set of available commands
// and conditions that guard those commands. Each script has an associated
// working directory and environment, along with a buffer containing the stdout
// and stderr output of a prior command, tracked in a [State] that commands can
// inspect and modify.
//
// The default commands configured by [NewEngine] resemble a simplified Unix
// shell.
//
// # Script Language
//
// Each line of a script is parsed into a sequence of space-separated command
// words, with environment variable expansion within each word and # marking an
// end-of-line comment. Additional variables named ':' and '/' are expanded
// within script arguments (expanding to the value of os.PathListSeparator and
// os.PathSeparator respectively) but are not inherited in subprocess
// environments.
//
// Adding single quotes around text keeps spaces in that text from being treated
// as word separators and also disables environment variable expansion.
// Inside a single-quoted block of text, a repeated single quote indicates
// a literal single quote, as in:
//
//	'Don''t communicate by sharing memory.'
//
// A line beginning with # is a comment and conventionally explains what is
// being done or tested at the start of a new section of the script.
//
// Commands are executed one at a time, and errors are checked for each command;
// if any command fails unexpectedly, no subsequent commands in the script are
// executed. The command prefix ! indicates that the command on the rest of the
// line (typically go or a matching predicate) must fail instead of succeeding.
// The command prefix ? indicates that the command may or may not succeed, but
// the script should continue regardless.
//
// The command prefix [cond] indicates that the command on the rest of the line
// should only run when the condition is satisfied.
//
// A condition can be negated: [!root] means to run the rest of the line only if
// the user is not root. Multiple conditions may be given for a single command,
// for example, '[linux] [amd64] skip'. The command will run if all conditions
// are satisfied.
package script

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// An Engine stores the configuration for executing a set of scripts.
//
// The same Engine may execute multiple scripts concurrently.
type Engine struct {
	Cmds  map[string]Cmd
	Conds map[string]Cond

	// If Quiet is true, Execute deletes log prints from the previous
	// section when starting a new section.
	Quiet bool
}

// NewEngine returns an Engine configured with a basic set of commands and conditions.
func NewEngine() *Engine

// A Cmd is a command that is available to a script.
type Cmd interface {
	Run(s *State, args ...string) (WaitFunc, error)

	Usage() *CmdUsage
}

// A WaitFunc is a function called to retrieve the results of a Cmd.
type WaitFunc func(*State) (stdout, stderr string, err error)

// A CmdUsage describes the usage of a Cmd, independent of its name
// (which can change based on its registration).
type CmdUsage struct {
	Summary string
	Args    string
	Detail  []string

	// If Async is true, the Cmd is meaningful to run in the background, and its
	// Run method must return either a non-nil WaitFunc or a non-nil error.
	Async bool

	// RegexpArgs reports which arguments, if any, should be treated as regular
	// expressions. It takes as input the raw, unexpanded arguments and returns
	// the list of argument indices that will be interpreted as regular
	// expressions.
	//
	// If RegexpArgs is nil, all arguments are assumed not to be regular
	// expressions.
	RegexpArgs func(rawArgs ...string) []int
}

// A Cond is a condition deciding whether a command should be run.
type Cond interface {
	Eval(s *State, suffix string) (bool, error)

	Usage() *CondUsage
}

// A CondUsage describes the usage of a Cond, independent of its name
// (which can change based on its registration).
type CondUsage struct {
	Summary string

	// If Prefix is true, the condition is a prefix and requires a
	// colon-separated suffix (like "[GOOS:linux]" for the "GOOS" condition).
	// The suffix may be the empty string (like "[prefix:]").
	Prefix bool
}

// Execute reads and executes script, writing the output to log.
//
// Execute stops and returns an error at the first command that does not succeed.
// The returned error's text begins with "file:line: ".
//
// If the script runs to completion or ends by a 'stop' command,
// Execute returns nil.
//
// Execute does not stop background commands started by the script
// before returning. To stop those, use [State.CloseAndWait] or the
// [Wait] command.
func (e *Engine) Execute(s *State, file string, script *bufio.Reader, log io.Writer) (err error)

// ListCmds prints to w a list of the named commands,
// annotating each with its arguments and a short usage summary.
// If verbose is true, ListCmds prints full details for each command.
//
// Each of the name arguments should be a command name.
// If no names are passed as arguments, ListCmds lists all the
// commands registered in e.
func (e *Engine) ListCmds(w io.Writer, verbose bool, names ...string) error

// ListConds prints to w a list of conditions, one per line,
// annotating each with a description and whether the condition
// is true in the state s (if s is non-nil).
//
// Each of the tag arguments should be a condition string of
// the form "name" or "name:suffix". If no tags are passed as
// arguments, ListConds lists all conditions registered in
// the engine e.
func (e *Engine) ListConds(w io.Writer, s *State, tags ...string) error
