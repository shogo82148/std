// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/time"
)

var ErrPlan9 = errors.New("unimplemented on Plan 9")

// File represents an open file descriptor.
type File struct {
	*file
}

// file is the real representation of *File.
// The extra level of indirection ensures that no clients of os
// can overwrite this data, which could cause the finalizer
// to close the wrong file descriptor.

// Fd returns the integer Unix file descriptor referencing the open file.
func (f *File) Fd() uintptr

// NewFile returns a new File with the given file descriptor and name.
func NewFile(fd uintptr, name string) *File

// Auxiliary information if the File describes a directory

// DevNull is the name of the operating system's “null device.”
// On Unix-like systems, it is "/dev/null"; on Windows, "NUL".
const DevNull = "/dev/null"

// OpenFile is the generalized open call; most users will use Open
// or Create instead.  It opens the named file with specified flag
// (O_RDONLY etc.) and perm, (0666 etc.) if applicable.  If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

// Close closes the File, rendering it unusable for I/O.
// It returns an error, if any.
func (file *File) Close() error

// Stat returns the FileInfo structure describing file.
// It returns the FileInfo and an error, if any.
func (f *File) Stat() (FileInfo, error)

// Truncate changes the size of the file.
// It does not change the I/O offset.
func (f *File) Truncate(size int64) error

// Chmod changes the mode of the file to mode.
// If there is an error, it will be of type *PathError.
func (f *File) Chmod(mode FileMode) error

// Sync commits the current contents of the file to stable storage.
// Typically, this means flushing the file system's in-memory copy
// of recently written data to disk.
func (f *File) Sync() (err error)

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
// If there is an error, it will be of type *PathError.
func Truncate(name string, size int64) error

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func Remove(name string) error

// Rename renames a file.
func Rename(oldname, newname string) error

// Chmod changes the mode of the named file to mode.
// If there is an error, it will be of type *PathError.
func Chmod(name string, mode FileMode) error

// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
func Chtimes(name string, atime time.Time, mtime time.Time) error

func Pipe() (r *File, w *File, err error)

// Link creates a hard link.
// If there is an error, it will be of type *LinkError.
func Link(oldname, newname string) error

// Symlink creates newname as a symbolic link to oldname.
// If there is an error, it will be of type *LinkError.
func Symlink(oldname, newname string) error

func Readlink(name string) (string, error)

func Chown(name string, uid, gid int) error

func Lchown(name string, uid, gid int) error

func (f *File) Chown(uid, gid int) error

// TempDir returns the default directory to use for temporary files.
func TempDir() string
