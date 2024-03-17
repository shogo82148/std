// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package syscall

// RouteRIB returns routing information base, as known as RIB,
// which consists of network facility information, states and
// parameters.
//
// Deprecated: Use golang.org/x/net/route instead.
func RouteRIB(facility, param int) ([]byte, error)

// RoutingMessage represents a routing message.
//
// Deprecated: Use golang.org/x/net/route instead.
type RoutingMessage interface {
	sockaddr() ([]Sockaddr, error)
}

// RouteMessage represents a routing message containing routing
// entries.
//
// Deprecated: Use golang.org/x/net/route instead.
type RouteMessage struct {
	Header RtMsghdr
	Data   []byte
}

// InterfaceMessage represents a routing message containing
// network interface entries.
//
// Deprecated: Use golang.org/x/net/route instead.
type InterfaceMessage struct {
	Header IfMsghdr
	Data   []byte
}

// InterfaceAddrMessage represents a routing message containing
// network interface address entries.
//
// Deprecated: Use golang.org/x/net/route instead.
type InterfaceAddrMessage struct {
	Header IfaMsghdr
	Data   []byte
}

// ParseRoutingMessage parses b as routing messages and returns the
// slice containing the [RoutingMessage] interfaces.
//
// Deprecated: Use golang.org/x/net/route instead.
func ParseRoutingMessage(b []byte) (msgs []RoutingMessage, err error)

// ParseRoutingSockaddr parses msg's payload as raw sockaddrs and
// returns the slice containing the [Sockaddr] interfaces.
//
// Deprecated: Use golang.org/x/net/route instead.
func ParseRoutingSockaddr(msg RoutingMessage) ([]Sockaddr, error)
