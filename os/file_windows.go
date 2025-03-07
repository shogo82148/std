// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// NewFile returns a new File with the given file descriptor and
// name. The returned value will be nil if fd is not a valid file
// descriptor.
func NewFile(fd uintptr, name string) *File

// DevNull is the name of the operating system's “null device.”
// On Unix-like systems, it is "/dev/null"; on Windows, "NUL".
const DevNull = "NUL"

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
func Truncate(name string, size int64) error

// Remove removes the named file or directory.
// If there is an error, it will be of type [*PathError].
func Remove(name string) error

// Pipe returns a connected pair of Files; reads from r return bytes written to w.
// It returns the files and an error, if any. The Windows handles underlying
// the returned files are marked as inheritable by child processes.
func Pipe() (r *File, w *File, err error)

// Link creates newname as a hard link to the oldname file.
// If there is an error, it will be of type *LinkError.
func Link(oldname, newname string) error

// Symlink creates newname as a symbolic link to oldname.
// On Windows, a symlink to a non-existent oldname creates a file symlink;
// if oldname is later created as a directory the symlink will not work.
// If there is an error, it will be of type *LinkError.
func Symlink(oldname, newname string) error
