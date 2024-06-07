// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fsys is an abstraction for reading files that
// allows for virtual overlays on top of the files on disk.
package fsys

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/os"

	"github.com/shogo82148/std/path/filepath"
)

// Trace emits a trace event for the operation and file path to the trace log,
// but only when $GODEBUG contains gofsystrace=1.
// The traces are appended to the file named by the $GODEBUG setting gofsystracelog, or else standard error.
// For debugging, if the $GODEBUG setting gofsystracestack is non-empty, then trace events for paths
// matching that glob pattern (using path.Match) will be followed by a full stack trace.
func Trace(op, path string)

// OverlayFile is the path to a text file in the OverlayJSON format.
// It is the value of the -overlay flag.
var OverlayFile string

// OverlayJSON is the format overlay files are expected to be in.
// The Replace map maps from overlaid paths to replacement paths:
// the Go command will forward all reads trying to open
// each overlaid path to its replacement path, or consider the overlaid
// path not to exist if the replacement path is empty.
type OverlayJSON struct {
	Replace map[string]string
}

// Init initializes the overlay, if one is being used.
func Init(wd string) error

// IsDir returns true if path is a directory on disk or in the
// overlay.
func IsDir(path string) (bool, error)

// ReadDir provides a slice of fs.FileInfo entries corresponding
// to the overlaid files in the directory.
func ReadDir(dir string) ([]fs.FileInfo, error)

// OverlayPath returns the path to the overlaid contents of the
// file, the empty string if the overlay deletes the file, or path
// itself if the file is not in the overlay, the file is a directory
// in the overlay, or there is no overlay.
// It returns true if the path is overlaid with a regular file
// or deleted, and false otherwise.
func OverlayPath(path string) (string, bool)

// Open opens the file at or overlaid on the given path.
func Open(path string) (*os.File, error)

// ReadFile reads the file at or overlaid on the given path.
func ReadFile(path string) ([]byte, error)

// IsDirWithGoFiles reports whether dir is a directory containing Go files
// either on disk or in the overlay.
func IsDirWithGoFiles(dir string) (bool, error)

// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root.
func Walk(root string, walkFn filepath.WalkFunc) error

// Lstat implements a version of os.Lstat that operates on the overlay filesystem.
func Lstat(path string) (fs.FileInfo, error)

// Stat implements a version of os.Stat that operates on the overlay filesystem.
func Stat(path string) (fs.FileInfo, error)

// Glob is like filepath.Glob but uses the overlay file system.
func Glob(pattern string) (matches []string, err error)
