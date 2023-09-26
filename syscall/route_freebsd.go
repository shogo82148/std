// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// InterfaceAnnounceMessage represents a routing message containing
// network interface arrival and departure information.
//
// Deprecated: Use golang.org/x/net/route instead.
type InterfaceAnnounceMessage struct {
	Header IfAnnounceMsghdr
}

// InterfaceMulticastAddrMessage represents a routing message
// containing network interface address entries.
//
// Deprecated: Use golang.org/x/net/route instead.
type InterfaceMulticastAddrMessage struct {
	Header IfmaMsghdr
	Data   []byte
}
