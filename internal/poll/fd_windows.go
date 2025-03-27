// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/syscall"
)

// InitWSA initiates the use of the Winsock DLL by the current process.
// It is called from the net package at init time to avoid
// loading ws2_32.dll when net is not used.
var InitWSA = sync.OnceFunc(func() {
	var d syscall.WSAData
	e := syscall.WSAStartup(uint32(0x202), &d)
	if e != nil {
		initErr = e
	}
	checkSetFileCompletionNotificationModes()
})

// FD is a file descriptor. The net and os packages embed this type in
// a larger type representing a network connection or OS file.
type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// System file descriptor. Immutable until Close.
	Sysfd syscall.Handle

	// Read operation.
	rop operation
	// Write operation.
	wop operation

	// I/O poller.
	pd pollDesc

	// Used to implement pread/pwrite.
	l sync.Mutex

	// The file offset for the next read or write.
	// Overlapped IO operations don't use the real file pointer,
	// so we need to keep track of the offset ourselves.
	offset int64

	// For console I/O.
	lastbits       []byte
	readuint16     []uint16
	readbyte       []byte
	readbyteOffset int

	// Semaphore signaled when file is closed.
	csema uint32

	skipSyncNotif bool

	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket.
	IsStream bool

	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	ZeroReadIsEOF bool

	// Whether this is a file rather than a network socket.
	isFile bool

	// The kind of this file.
	kind fileKind
}

// Init initializes the FD. The Sysfd field should already be set.
// This can be called multiple times on a single FD.
// The net argument is a network name from the net package (e.g., "tcp"),
// or "file" or "console" or "dir".
// Set pollable to true if fd should be managed by runtime netpoll.
func (fd *FD) Init(net string, pollable bool) error

// Close closes the FD. The underlying file descriptor is closed by
// the destroy method when there are no remaining references.
func (fd *FD) Close() error

// Read implements io.Reader.
func (fd *FD) Read(buf []byte) (int, error)

var ReadConsole = syscall.ReadConsole

// Pread emulates the Unix pread system call.
func (fd *FD) Pread(b []byte, off int64) (int, error)

// ReadFrom wraps the recvfrom network call.
func (fd *FD) ReadFrom(buf []byte) (int, syscall.Sockaddr, error)

// ReadFromInet4 wraps the recvfrom network call for IPv4.
func (fd *FD) ReadFromInet4(buf []byte, sa4 *syscall.SockaddrInet4) (int, error)

// ReadFromInet6 wraps the recvfrom network call for IPv6.
func (fd *FD) ReadFromInet6(buf []byte, sa6 *syscall.SockaddrInet6) (int, error)

// Write implements io.Writer.
func (fd *FD) Write(buf []byte) (int, error)

// Pwrite emulates the Unix pwrite system call.
func (fd *FD) Pwrite(buf []byte, off int64) (int, error)

// Writev emulates the Unix writev system call.
func (fd *FD) Writev(buf *[][]byte) (int64, error)

// WriteTo wraps the sendto network call.
func (fd *FD) WriteTo(buf []byte, sa syscall.Sockaddr) (int, error)

// WriteToInet4 is WriteTo, specialized for syscall.SockaddrInet4.
func (fd *FD) WriteToInet4(buf []byte, sa4 *syscall.SockaddrInet4) (int, error)

// WriteToInet6 is WriteTo, specialized for syscall.SockaddrInet6.
func (fd *FD) WriteToInet6(buf []byte, sa6 *syscall.SockaddrInet6) (int, error)

// Call ConnectEx. This doesn't need any locking, since it is only
// called when the descriptor is first created. This is here rather
// than in the net package so that it can use fd.wop.
func (fd *FD) ConnectEx(ra syscall.Sockaddr) error

// Accept handles accepting a socket. The sysSocket parameter is used
// to allocate the net socket.
func (fd *FD) Accept(sysSocket func() (syscall.Handle, error)) (syscall.Handle, []syscall.RawSockaddrAny, uint32, string, error)

// Seek wraps syscall.Seek.
func (fd *FD) Seek(offset int64, whence int) (int64, error)

// Fchmod updates syscall.ByHandleFileInformation.Fileattributes when needed.
func (fd *FD) Fchmod(mode uint32) error

// Fchdir wraps syscall.Fchdir.
func (fd *FD) Fchdir() error

// GetFileType wraps syscall.GetFileType.
func (fd *FD) GetFileType() (uint32, error)

// GetFileInformationByHandle wraps GetFileInformationByHandle.
func (fd *FD) GetFileInformationByHandle(data *syscall.ByHandleFileInformation) error

// RawRead invokes the user-defined function f for a read operation.
func (fd *FD) RawRead(f func(uintptr) bool) error

// RawWrite invokes the user-defined function f for a write operation.
func (fd *FD) RawWrite(f func(uintptr) bool) error

// ReadMsg wraps the WSARecvMsg network call.
func (fd *FD) ReadMsg(p []byte, oob []byte, flags int) (int, int, int, syscall.Sockaddr, error)

// ReadMsgInet4 is ReadMsg, but specialized to return a syscall.SockaddrInet4.
func (fd *FD) ReadMsgInet4(p []byte, oob []byte, flags int, sa4 *syscall.SockaddrInet4) (int, int, int, error)

// ReadMsgInet6 is ReadMsg, but specialized to return a syscall.SockaddrInet6.
func (fd *FD) ReadMsgInet6(p []byte, oob []byte, flags int, sa6 *syscall.SockaddrInet6) (int, int, int, error)

// WriteMsg wraps the WSASendMsg network call.
func (fd *FD) WriteMsg(p []byte, oob []byte, sa syscall.Sockaddr) (int, int, error)

// WriteMsgInet4 is WriteMsg specialized for syscall.SockaddrInet4.
func (fd *FD) WriteMsgInet4(p []byte, oob []byte, sa *syscall.SockaddrInet4) (int, int, error)

// WriteMsgInet6 is WriteMsg specialized for syscall.SockaddrInet6.
func (fd *FD) WriteMsgInet6(p []byte, oob []byte, sa *syscall.SockaddrInet6) (int, int, error)

func DupCloseOnExec(fd int) (int, string, error)
