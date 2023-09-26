// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Network interface identification

package net

// Interface represents a mapping between network interface name
// and index.  It also represents network interface facility
// information.
type Interface struct {
	Index        int
	MTU          int
	Name         string
	HardwareAddr HardwareAddr
	Flags        Flags
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func (f Flags) String() string

// Addrs returns interface addresses for a specific interface.
func (ifi *Interface) Addrs() ([]Addr, error)

// MulticastAddrs returns multicast, joined group addresses for
// a specific interface.
func (ifi *Interface) MulticastAddrs() ([]Addr, error)

// Interfaces returns a list of the system's network interfaces.
func Interfaces() ([]Interface, error)

// InterfaceAddrs returns a list of the system's network interface
// addresses.
func InterfaceAddrs() ([]Addr, error)

// InterfaceByIndex returns the interface specified by index.
func InterfaceByIndex(index int) (*Interface, error)

// InterfaceByName returns the interface specified by name.
func InterfaceByName(name string) (*Interface, error)
