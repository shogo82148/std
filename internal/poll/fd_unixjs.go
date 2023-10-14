// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm)

package poll

import "github.com/shogo82148/std/syscall"

type SysFile struct {
	// Writev cache.
	iovecs *[]syscall.Iovec
}

// Fchdir wraps syscall.Fchdir.
func (fd *FD) Fchdir() error

// ReadDirent wraps syscall.ReadDirent.
// We treat this like an ordinary system call rather than a call
// that tries to fill the buffer.
func (fd *FD) ReadDirent(buf []byte) (int, error)

// Seek wraps syscall.Seek.
func (fd *FD) Seek(offset int64, whence int) (int64, error)
