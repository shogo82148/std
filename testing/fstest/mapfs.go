// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fstest

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// A MapFS is a simple in-memory file system for use in tests,
// represented as a map from path names (arguments to Open)
// to information about the files or directories they represent.
//
// The map need not include parent directories for files contained
// in the map; those will be synthesized if needed.
// But a directory can still be included by setting the MapFile.Mode's ModeDir bit;
// this may be necessary for detailed control over the directory's FileInfo
// or to create an empty directory.
//
// File system operations read directly from the map,
// so that the file system can be changed by editing the map as needed.
// An implication is that file system operations must not run concurrently
// with changes to the map, which would be a race.
// Another implication is that opening or reading a directory requires
// iterating over the entire map, so a MapFS should typically be used with not more
// than a few hundred entries or directory reads.
type MapFS map[string]*MapFile

// A MapFile describes a single file in a MapFS.
type MapFile struct {
	Data    []byte
	Mode    fs.FileMode
	ModTime time.Time
	Sys     any
}

var _ fs.FS = MapFS(nil)
var _ fs.File = (*openMapFile)(nil)

// Open opens the named file.
func (fsys MapFS) Open(name string) (fs.File, error)

func (fsys MapFS) ReadFile(name string) ([]byte, error)

func (fsys MapFS) Stat(name string) (fs.FileInfo, error)

func (fsys MapFS) ReadDir(name string) ([]fs.DirEntry, error)

func (fsys MapFS) Glob(pattern string) ([]string, error)

func (fsys MapFS) Sub(dir string) (fs.FS, error)
