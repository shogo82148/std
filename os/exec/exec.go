// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exec runs external commands. It wraps os.StartProcess to make it
// easier to remap stdin and stdout, connect I/O with pipes, and do other
// adjustments.
package exec

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/syscall"
)

// Error records the name of a binary that failed to be executed
// and the reason it failed.
type Error struct {
	Name string
	Err  error
}

func (e *Error) Error() string

// Cmd represents an external command being prepared or run.
type Cmd struct {
	Path string

	Args []string

	Env []string

	Dir string

	Stdin io.Reader

	Stdout io.Writer
	Stderr io.Writer

	ExtraFiles []*os.File

	SysProcAttr *syscall.SysProcAttr

	Process *os.Process

	ProcessState *os.ProcessState

	err             error
	finished        bool
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error
}

// Command returns the Cmd struct to execute the named program with
// the given arguments.
//
// It sets Path and Args in the returned structure and zeroes the
// other fields.
//
// If name contains no path separators, Command uses LookPath to
// resolve the path to a complete name if possible. Otherwise it uses
// name directly.
//
// The returned Cmd's Args field is constructed from the command name
// followed by the elements of arg, so arg should not include the
// command name itself. For example, Command("echo", "hello")
func Command(name string, arg ...string) *Cmd

// Run starts the specified command and waits for it to complete.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
//
// If the command fails to run or doesn't complete successfully, the
// error is of type *ExitError. Other error types may be
// returned for I/O problems.
func (c *Cmd) Run() error

// Start starts the specified command but does not wait for it to complete.
func (c *Cmd) Start() error

// An ExitError reports an unsuccessful exit by a command.
type ExitError struct {
	*os.ProcessState
}

func (e *ExitError) Error() string

// Wait waits for the command to exit.
// It must have been started by Start.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
//
// If the command fails to run or doesn't complete successfully, the
// error is of type *ExitError. Other error types may be
// returned for I/O problems.
func (c *Cmd) Wait() error

// Output runs the command and returns its standard output.
func (c *Cmd) Output() ([]byte, error)

// CombinedOutput runs the command and returns its combined standard
// output and standard error.
func (c *Cmd) CombinedOutput() ([]byte, error)

// StdinPipe returns a pipe that will be connected to the command's
// standard input when the command starts.
func (c *Cmd) StdinPipe() (io.WriteCloser, error)

// StdoutPipe returns a pipe that will be connected to the command's
// standard output when the command starts.
// The pipe will be closed automatically after Wait sees the command exit.
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)

// StderrPipe returns a pipe that will be connected to the command's
// standard error when the command starts.
// The pipe will be closed automatically after Wait sees the command exit.
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
