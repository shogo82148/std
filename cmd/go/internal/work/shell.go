// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package work

import (
	"github.com/shogo82148/std/io/fs"
)

// A Shell runs shell commands and performs shell-like file system operations.
//
// Shell tracks context related to running commands, and form a tree much like
// context.Context.
type Shell struct {
	action *Action
	*shellShared
}

// NewShell returns a new Shell.
//
// Shell will internally serialize calls to the print function.
// If print is nil, it defaults to printing to stderr.
func NewShell(workDir string, print func(a ...any) (int, error)) *Shell

// Print emits a to this Shell's output stream, formatting it like fmt.Print.
// It is safe to call concurrently.
func (sh *Shell) Print(a ...any)

// WithAction returns a Shell identical to sh, but bound to Action a.
func (sh *Shell) WithAction(a *Action) *Shell

// Shell returns a shell for running commands on behalf of Action a.
func (b *Builder) Shell(a *Action) *Shell

// BackgroundShell returns a Builder-wide Shell that's not bound to any Action.
// Try not to use this unless there's really no sensible Action available.
func (b *Builder) BackgroundShell() *Shell

// copyFile is like 'cp src dst'.
func (sh *Shell) CopyFile(dst, src string, perm fs.FileMode, force bool) error

// Mkdir makes the named directory.
func (sh *Shell) Mkdir(dir string) error

// RemoveAll is like 'rm -rf'. It attempts to remove all paths even if there's
// an error, and returns the first error.
func (sh *Shell) RemoveAll(paths ...string) error

// Symlink creates a symlink newname -> oldname.
func (sh *Shell) Symlink(oldname, newname string) error

// ShowCmd prints the given command to standard output
// for the implementation of -n or -x.
//
// ShowCmd also replaces the name of the current script directory with dot (.)
// but only when it is at the beginning of a space-separated token.
//
// If dir is not "" or "/" and not the current script directory, ShowCmd first
// prints a "cd" command to switch to dir and updates the script directory.
func (sh *Shell) ShowCmd(dir string, format string, args ...any)
