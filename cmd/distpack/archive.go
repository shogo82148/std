// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// An Archive describes an archive to write: a collection of files.
// Directories are implied by the files and not explicitly listed.
type Archive struct {
	Files []File
}

// A File describes a single file to write to an archive.
type File struct {
	Name string
	Time time.Time
	Mode fs.FileMode
	Size int64
	Src  string
}

// Info returns a FileInfo about the file, for use with tar.FileInfoHeader
// and zip.FileInfoHeader.
func (f *File) Info() fs.FileInfo

// NewArchive returns a new Archive containing all the files in the directory dir.
// The archive can be amended afterward using methods like Add and Filter.
func NewArchive(dir string) (*Archive, error)

// Add adds a file with the given name and info to the archive.
// The content of the file comes from the operating system file src.
// After a sequence of one or more calls to Add,
// the caller should invoke Sort to re-sort the archive's files.
func (a *Archive) Add(name, src string, info fs.FileInfo)

// Sort sorts the files in the archive.
// It is only necessary to call Sort after calling Add or RenameGoMod.
// NewArchive returns a sorted archive, and the other methods
// preserve the sorting of the archive.
func (a *Archive) Sort()

// Clone returns a copy of the Archive.
// Method calls like Add and Filter invoked on the copy do not affect the original,
// nor do calls on the original affect the copy.
func (a *Archive) Clone() *Archive

// AddPrefix adds a prefix to all file names in the archive.
func (a *Archive) AddPrefix(prefix string)

// Filter removes files from the archive for which keep(name) returns false.
func (a *Archive) Filter(keep func(name string) bool)

// SetMode changes the mode of every file in the archive
// to be mode(name, m), where m is the file's current mode.
func (a *Archive) SetMode(mode func(name string, m fs.FileMode) fs.FileMode)

// Remove removes files matching any of the patterns from the archive.
// The patterns use the syntax of path.Match, with an extension of allowing
// a leading **/ or trailing /**, which match any number of path elements
// (including no path elements) before or after the main match.
func (a *Archive) Remove(patterns ...string)

// SetTime sets the modification time of all files in the archive to t.
func (a *Archive) SetTime(t time.Time)

// RenameGoMod renames the go.mod files in the archive to _go.mod,
// for use with the module form, which cannot contain other go.mod files.
func (a *Archive) RenameGoMod()
