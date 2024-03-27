// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package script

import (
	"github.com/shogo82148/std/os/exec"
	"github.com/shogo82148/std/time"
)

// DefaultCmds returns a set of broadly useful script commands.
//
// Run the 'help' command within a script engine to view a list of the available
// commands.
func DefaultCmds() map[string]Cmd

// Command returns a new Cmd with a Usage method that returns a copy of the
// given CmdUsage and a Run method calls the given function.
func Command(usage CmdUsage, run func(*State, ...string) (WaitFunc, error)) Cmd

// Cat writes the concatenated contents of the named file(s) to the script's
// stdout buffer.
func Cat() Cmd

// Cd changes the current working directory.
func Cd() Cmd

// Chmod changes the permissions of a file or a directory..
func Chmod() Cmd

// Cmp compares the contents of two files, or the contents of either the
// "stdout" or "stderr" buffer and a file, returning a non-nil error if the
// contents differ.
func Cmp() Cmd

// Cmpenv is like Compare, but also performs environment substitutions
// on the contents of both arguments.
func Cmpenv() Cmd

// Cp copies one or more files to a new location.
func Cp() Cmd

// Echo writes its arguments to stdout, followed by a newline.
func Echo() Cmd

// Env sets or logs the values of environment variables.
//
// With no arguments, Env reports all variables in the environment.
// "key=value" arguments set variables, and arguments without "="
// cause the corresponding value to be printed to the stdout buffer.
func Env() Cmd

// Exec runs an arbitrary executable as a subprocess.
//
// When the Script's context is canceled, Exec sends the interrupt signal, then
// waits for up to the given delay for the subprocess to flush output before
// terminating it with os.Kill.
func Exec(cancel func(*exec.Cmd) error, waitDelay time.Duration) Cmd

// Exists checks that the named file(s) exist.
func Exists() Cmd

// Grep checks that file content matches a regexp.
// Like stdout/stderr and unlike Unix grep, it accepts Go regexp syntax.
//
// Grep does not modify the State's stdout or stderr buffers.
// (Its output goes to the script log, not stdout.)
func Grep() Cmd

// Help writes command documentation to the script log.
func Help() Cmd

// Mkdir creates a directory and any needed parent directories.
func Mkdir() Cmd

// Mv renames an existing file or directory to a new path.
func Mv() Cmd

// Program returns a new command that runs the named program, found from the
// host process's PATH (not looked up in the script's PATH).
func Program(name string, cancel func(*exec.Cmd) error, waitDelay time.Duration) Cmd

// Replace replaces all occurrences of a string in a file with another string.
func Replace() Cmd

// Rm removes a file or directory.
//
// If a directory, Rm also recursively removes that directory's
// contents.
func Rm() Cmd

// Sleep sleeps for the given Go duration or until the script's context is
// canceled, whichever happens first.
func Sleep() Cmd

// Stderr searches for a regular expression in the stderr buffer.
func Stderr() Cmd

// Stdout searches for a regular expression in the stdout buffer.
func Stdout() Cmd

// Stop returns a sentinel error that causes script execution to halt
// and s.Execute to return with a nil error.
func Stop() Cmd

// Symlink creates a symbolic link.
func Symlink() Cmd

// Wait waits for the completion of background commands.
//
// When Wait returns, the stdout and stderr buffers contain the concatenation of
// the background commands' respective outputs in the order in which those
// commands were started.
func Wait() Cmd
