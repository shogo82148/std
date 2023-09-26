// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	Name string
	Net  string
}

// Network returns the address's network name, "unix", "unixgram" or
// "unixpacket".
func (a *UnixAddr) Network() string

func (a *UnixAddr) String() string

// ResolveUnixAddr parses addr as a Unix domain socket address.
// The string net gives the network name, "unix", "unixgram" or
// "unixpacket".
func ResolveUnixAddr(net, addr string) (*UnixAddr, error)
