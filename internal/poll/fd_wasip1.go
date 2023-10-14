// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import (
	"github.com/shogo82148/std/syscall"
)

type SysFile struct {
	// RefCountPtr is a pointer to the reference count of Sysfd.
	//
	// WASI preview 1 lacks a dup(2) system call. When the os and net packages
	// need to share a file/socket, instead of duplicating the underlying file
	// descriptor, we instead provide a way to copy FD instances and manage the
	// underlying file descriptor with reference counting.
	RefCountPtr *int32

	// RefCount is the reference count of Sysfd. When a copy of an FD is made,
	// it points to the reference count of the original FD instance.
	RefCount int32

	// Cache for the file type, lazily initialized when Seek is called.
	Filetype uint32

	// If the file represents a directory, this field contains the current
	// readdir position. It is reset to zero if the program calls Seek(0, 0).
	Dircookie uint64

	// Absolute path of the file, as returned by syscall.PathOpen;
	// this is used by Fchdir to emulate setting the current directory
	// to an open file descriptor.
	Path string
}

// Copy creates a copy of the FD.
//
// The FD instance points to the same underlying file descriptor. The file
// descriptor isn't closed until all FD instances that refer to it have been
// closed/destroyed.
func (fd *FD) Copy() FD

// Fchdir wraps syscall.Fchdir.
func (fd *FD) Fchdir() error

// ReadDir wraps syscall.ReadDir.
// We treat this like an ordinary system call rather than a call
// that tries to fill the buffer.
func (fd *FD) ReadDir(buf []byte, cookie syscall.Dircookie) (int, error)

func (fd *FD) ReadDirent(buf []byte) (int, error)

// Seek wraps syscall.Seek.
func (fd *FD) Seek(offset int64, whence int) (int64, error)
