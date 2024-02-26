// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	Destroy func()

	// deadlines
	rmu       sync.Mutex
	wmu       sync.Mutex
	raio      *asyncIO
	waio      *asyncIO
	rtimer    *time.Timer
	wtimer    *time.Timer
	rtimedout bool
	wtimedout bool

	// Whether this is a normal file.
	// On Plan 9 we do not use this package for ordinary files,
	// so this is always false, but the field is present because
	// shared code in fd_mutex.go checks it.
	isFile bool
}

// Close handles the locking for closing an FD. The real operation
// is in the net package.
func (fd *FD) Close() error

// Read implements io.Reader.
func (fd *FD) Read(fn func([]byte) (int, error), b []byte) (int, error)

// Write implements io.Writer.
func (fd *FD) Write(fn func([]byte) (int, error), b []byte) (int, error)

// SetDeadline sets the read and write deadlines associated with fd.
func (fd *FD) SetDeadline(t time.Time) error

// SetReadDeadline sets the read deadline associated with fd.
func (fd *FD) SetReadDeadline(t time.Time) error

// SetWriteDeadline sets the write deadline associated with fd.
func (fd *FD) SetWriteDeadline(t time.Time) error

// ReadLock wraps FD.readLock.
func (fd *FD) ReadLock() error

// ReadUnlock wraps FD.readUnlock.
func (fd *FD) ReadUnlock()

// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
func IsPollDescriptor(fd uintptr) bool

// RawControl invokes the user-defined function f for a non-IO
// operation.
func (fd *FD) RawControl(f func(uintptr)) error

// RawRead invokes the user-defined function f for a read operation.
func (fd *FD) RawRead(f func(uintptr) bool) error

// RawWrite invokes the user-defined function f for a write operation.
func (fd *FD) RawWrite(f func(uintptr) bool) error

func DupCloseOnExec(fd int) (int, string, error)
