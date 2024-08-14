// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/internal/txtar"
	"github.com/shogo82148/std/io"
)

// A State encapsulates the current state of a running script engine,
// including the script environment and any running background commands.
type State struct {
	engine *Engine

	ctx    context.Context
	cancel context.CancelFunc
	file   string
	log    bytes.Buffer

	workdir string
	pwd     string
	env     []string
	envMap  map[string]string
	stdout  string
	stderr  string

	background []backgroundCmd
}

// NewState returns a new State permanently associated with ctx, with its
// initial working directory in workdir and its initial environment set to
// initialEnv (or os.Environ(), if initialEnv is nil).
//
// The new State also contains pseudo-environment-variables for
// ${/} and ${:} (for the platform's path and list separators respectively),
// but does not pass those to subprocesses.
func NewState(ctx context.Context, workdir string, initialEnv []string) (*State, error)

// CloseAndWait cancels the State's Context and waits for any background commands to
// finish. If any remaining background command ended in an unexpected state,
// Close returns a non-nil error.
func (s *State) CloseAndWait(log io.Writer) error

// Chdir changes the State's working directory to the given path.
func (s *State) Chdir(path string) error

// Context returns the Context with which the State was created.
func (s *State) Context() context.Context

// Environ returns a copy of the current script environment,
// in the form "key=value".
func (s *State) Environ() []string

// ExpandEnv replaces ${var} or $var in the string according to the values of
// the environment variables in s. References to undefined variables are
// replaced by the empty string.
func (s *State) ExpandEnv(str string, inRegexp bool) string

// ExtractFiles extracts the files in ar to the state's current directory,
// expanding any environment variables within each name.
//
// The files must reside within the working directory with which the State was
// originally created.
func (s *State) ExtractFiles(ar *txtar.Archive) error

// Getwd returns the directory in which to run the next script command.
func (s *State) Getwd() string

// Logf writes output to the script's log without updating its stdout or stderr
// buffers. (The output log functions as a kind of meta-stderr.)
func (s *State) Logf(format string, args ...any)

// LookupEnv retrieves the value of the environment variable in s named by the key.
func (s *State) LookupEnv(key string) (string, bool)

// Path returns the absolute path in the host operating system for a
// script-based (generally slash-separated and relative) path.
func (s *State) Path(path string) string

// Setenv sets the value of the environment variable in s named by the key.
func (s *State) Setenv(key, value string) error

// Stdout returns the stdout output of the last command run,
// or the empty string if no command has been run.
func (s *State) Stdout() string

// Stderr returns the stderr output of the last command run,
// or the empty string if no command has been run.
func (s *State) Stderr() string
