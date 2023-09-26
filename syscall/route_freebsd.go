// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Routing sockets and messages for FreeBSD

package syscall

// InterfaceMulticastAddrMessage represents a routing message
// containing network interface address entries.
type InterfaceMulticastAddrMessage struct {
	Header IfmaMsghdr
	Data   []byte
}
