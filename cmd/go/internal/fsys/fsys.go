// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fsys implements a virtual file system that the go command
// uses to read source file trees. The virtual file system redirects some
// OS file paths to other OS file paths, according to an overlay file.
// Editors can use this overlay support to invoke the go command on
// temporary files that have been edited but not yet saved into their
// final locations.
package fsys

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/os"
)

// Trace emits a trace event for the operation and file path to the trace log,
// but only when $GODEBUG contains gofsystrace=1.
// The traces are appended to the file named by the $GODEBUG setting gofsystracelog, or else standard error.
// For debugging, if the $GODEBUG setting gofsystracestack is non-empty, then trace events for paths
// matching that glob pattern (using path.Match) will be followed by a full stack trace.
func Trace(op, path string)

// OverlayFile is the -overlay flag value.
// It names a file containing the JSON for an overlayJSON struct.
var OverlayFile string

// Bind makes the virtual file system use dir as if it were mounted at mtpt,
// like Plan 9's “bind” or Linux's “mount --bind”, or like os.Symlink
// but without the symbolic link.
//
// For now, the behavior of using Bind on multiple overlapping
// mountpoints (for example Bind("x", "/a") and Bind("y", "/a/b"))
// is undefined.
func Bind(dir, mtpt string)

// Init initializes the overlay, if one is being used.
func Init() error

// IsDir returns true if path is a directory on disk or in the
// overlay.
func IsDir(path string) (bool, error)

// ReadDir reads the named directory in the virtual file system.
func ReadDir(name string) ([]fs.DirEntry, error)

// Actual returns the actual file system path for the named file.
// It returns the empty string if name has been deleted in the virtual file system.
func Actual(name string) string

// Replaced reports whether the named file has been modified
// in the virtual file system compared to the OS file system.
func Replaced(name string) bool

// Open opens the named file in the virtual file system.
// It must be an ordinary file, not a directory.
func Open(name string) (*os.File, error)

// ReadFile reads the named file from the virtual file system
// and returns the contents.
func ReadFile(name string) ([]byte, error)

// IsGoDir reports whether the named directory in the virtual file system
// is a directory containing one or more Go source files.
func IsGoDir(name string) (bool, error)

// Lstat returns a FileInfo describing the named file in the virtual file system.
// It does not follow symbolic links
func Lstat(name string) (fs.FileInfo, error)

// Stat returns a FileInfo describing the named file in the virtual file system.
// It follows symbolic links.
func Stat(name string) (fs.FileInfo, error)
