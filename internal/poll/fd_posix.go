// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package poll

// Shutdown wraps syscall.Shutdown.
func (fd *FD) Shutdown(how int) error

// Fchown wraps syscall.Fchown.
func (fd *FD) Fchown(uid, gid int) error

// Ftruncate wraps syscall.Ftruncate.
func (fd *FD) Ftruncate(size int64) error

// RawControl invokes the user-defined function f for a non-IO
// operation.
func (fd *FD) RawControl(f func(uintptr)) error
