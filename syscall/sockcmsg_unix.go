// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

// Socket control messages

package syscall

// CmsgLen returns the value to store in the Len field of the Cmsghdr
// structure, taking into account any necessary alignment.
func CmsgLen(datalen int) int

// CmsgSpace returns the number of bytes an ancillary element with
// payload of the passed data length occupies.
func CmsgSpace(datalen int) int

// SocketControlMessage represents a socket control message.
type SocketControlMessage struct {
	Header Cmsghdr
	Data   []byte
}

// ParseSocketControlMessage parses b as an array of socket control
// messages.
func ParseSocketControlMessage(b []byte) ([]SocketControlMessage, error)

// UnixRights encodes a set of open file descriptors into a socket
// control message for sending to another process.
func UnixRights(fds ...int) []byte

// ParseUnixRights decodes a socket control message that contains an
// integer array of open file descriptors from another process.
func ParseUnixRights(m *SocketControlMessage) ([]int, error)
