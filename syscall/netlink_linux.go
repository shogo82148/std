// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Netlink sockets and messages

package syscall

// NetlinkRouteRequest represents the request message to receive
// routing and link states from the kernel.
type NetlinkRouteRequest struct {
	Header NlMsghdr
	Data   RtGenmsg
}

// NetlinkRIB returns routing information base, as known as RIB,
// which consists of network facility information, states and
// parameters.
func NetlinkRIB(proto, family int) ([]byte, error)

// NetlinkMessage represents the netlink message.
type NetlinkMessage struct {
	Header NlMsghdr
	Data   []byte
}

// ParseNetlinkMessage parses buf as netlink messages and returns
// the slice containing the NetlinkMessage structs.
func ParseNetlinkMessage(buf []byte) ([]NetlinkMessage, error)

// NetlinkRouteAttr represents the netlink route attribute.
type NetlinkRouteAttr struct {
	Attr  RtAttr
	Value []byte
}

// ParseNetlinkRouteAttr parses msg's payload as netlink route
// attributes and returns the slice containing the NetlinkRouteAttr
// structs.
func ParseNetlinkRouteAttr(msg *NetlinkMessage) ([]NetlinkRouteAttr, error)
