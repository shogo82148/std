// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package poll

import (
	"github.com/shogo82148/std/syscall"
)

// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// System file descriptor. Immutable until Close.
	Sysfd int

	// Platform dependent state of the file descriptor.
	SysFile

	// I/O poller.
	pd pollDesc

	// Semaphore signaled when file is closed.
	csema uint32

	// Non-zero if this file has been set to blocking mode.
	isBlocking uint32

	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
	IsStream bool

	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	ZeroReadIsEOF bool

	// Whether this is a file rather than a network socket.
	isFile bool
}

// Init initializes the FD. The Sysfd field should already be set.
// This can be called multiple times on a single FD.
// The net argument is a network name from the net package (e.g., "tcp"),
// or "file".
// Set pollable to true if fd should be managed by runtime netpoll.
func (fd *FD) Init(net string, pollable bool) error

// Close closes the FD. The underlying file descriptor is closed by the
// destroy method when there are no remaining references.
func (fd *FD) Close() error

// SetBlocking puts the file into blocking mode.
func (fd *FD) SetBlocking() error

// Read implements io.Reader.
func (fd *FD) Read(p []byte) (int, error)

// Pread wraps the pread system call.
func (fd *FD) Pread(p []byte, off int64) (int, error)

// ReadFrom wraps the recvfrom network call.
func (fd *FD) ReadFrom(p []byte) (int, syscall.Sockaddr, error)

// ReadFromInet4 wraps the recvfrom network call for IPv4.
func (fd *FD) ReadFromInet4(p []byte, from *syscall.SockaddrInet4) (int, error)

// ReadFromInet6 wraps the recvfrom network call for IPv6.
func (fd *FD) ReadFromInet6(p []byte, from *syscall.SockaddrInet6) (int, error)

// ReadMsg wraps the recvmsg network call.
func (fd *FD) ReadMsg(p []byte, oob []byte, flags int) (int, int, int, syscall.Sockaddr, error)

// ReadMsgInet4 is ReadMsg, but specialized for syscall.SockaddrInet4.
func (fd *FD) ReadMsgInet4(p []byte, oob []byte, flags int, sa4 *syscall.SockaddrInet4) (int, int, int, error)

// ReadMsgInet6 is ReadMsg, but specialized for syscall.SockaddrInet6.
func (fd *FD) ReadMsgInet6(p []byte, oob []byte, flags int, sa6 *syscall.SockaddrInet6) (int, int, int, error)

// Write implements io.Writer.
func (fd *FD) Write(p []byte) (int, error)

// Pwrite wraps the pwrite system call.
func (fd *FD) Pwrite(p []byte, off int64) (int, error)

// WriteToInet4 wraps the sendto network call for IPv4 addresses.
func (fd *FD) WriteToInet4(p []byte, sa *syscall.SockaddrInet4) (int, error)

// WriteToInet6 wraps the sendto network call for IPv6 addresses.
func (fd *FD) WriteToInet6(p []byte, sa *syscall.SockaddrInet6) (int, error)

// WriteTo wraps the sendto network call.
func (fd *FD) WriteTo(p []byte, sa syscall.Sockaddr) (int, error)

// WriteMsg wraps the sendmsg network call.
func (fd *FD) WriteMsg(p []byte, oob []byte, sa syscall.Sockaddr) (int, int, error)

// WriteMsgInet4 is WriteMsg specialized for syscall.SockaddrInet4.
func (fd *FD) WriteMsgInet4(p []byte, oob []byte, sa *syscall.SockaddrInet4) (int, int, error)

// WriteMsgInet6 is WriteMsg specialized for syscall.SockaddrInet6.
func (fd *FD) WriteMsgInet6(p []byte, oob []byte, sa *syscall.SockaddrInet6) (int, int, error)

// Accept wraps the accept network call.
func (fd *FD) Accept() (int, syscall.Sockaddr, string, error)

// Fchmod wraps syscall.Fchmod.
func (fd *FD) Fchmod(mode uint32) error

// Fstat wraps syscall.Fstat
func (fd *FD) Fstat(s *syscall.Stat_t) error

// DupCloseOnExec dups fd and marks it close-on-exec.
func DupCloseOnExec(fd int) (int, string, error)

// Dup duplicates the file descriptor.
func (fd *FD) Dup() (int, string, error)

// WaitWrite waits until data can be written to fd.
func (fd *FD) WaitWrite() error

// WriteOnce is for testing only. It makes a single write call.
func (fd *FD) WriteOnce(p []byte) (int, error)

// RawRead invokes the user-defined function f for a read operation.
func (fd *FD) RawRead(f func(uintptr) bool) error

// RawWrite invokes the user-defined function f for a write operation.
func (fd *FD) RawWrite(f func(uintptr) bool) error
