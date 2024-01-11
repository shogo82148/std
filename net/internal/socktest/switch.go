// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package socktest provides utilities for socket testing.
package socktest

import (
	"github.com/shogo82148/std/sync"
)

// A Switch represents a callpath point switch for socket system
// calls.
type Switch struct {
	once sync.Once

	fmu   sync.RWMutex
	fltab map[FilterType]Filter

	smu   sync.RWMutex
	sotab Sockets
	stats stats
}

// Stats returns a list of per-cookie socket statistics.
func (sw *Switch) Stats() []Stat

// Sockets returns mappings of socket descriptor to socket status.
func (sw *Switch) Sockets() Sockets

// A Cookie represents a 3-tuple of a socket; address family, socket
// type and protocol number.
type Cookie uint64

// Family returns an address family.
func (c Cookie) Family() int

// Type returns a socket type.
func (c Cookie) Type() int

// Protocol returns a protocol number.
func (c Cookie) Protocol() int

// A Status represents the status of a socket.
type Status struct {
	Cookie    Cookie
	Err       error
	SocketErr error
}

func (so Status) String() string

// A Stat represents a per-cookie socket statistics.
type Stat struct {
	Family   int
	Type     int
	Protocol int

	Opened    uint64
	Connected uint64
	Listened  uint64
	Accepted  uint64
	Closed    uint64

	OpenFailed    uint64
	ConnectFailed uint64
	ListenFailed  uint64
	AcceptFailed  uint64
	CloseFailed   uint64
}

func (st Stat) String() string

// A FilterType represents a filter type.
type FilterType int

const (
	FilterSocket FilterType = iota
	FilterConnect
	FilterListen
	FilterAccept
	FilterGetsockoptInt
	FilterClose
)

// A Filter represents a socket system call filter.
//
// It will only be executed before a system call for a socket that has
// an entry in internal table.
// If the filter returns a non-nil error, the execution of system call
// will be canceled and the system call function returns the non-nil
// error.
// It can return a non-nil [AfterFilter] for filtering after the
// execution of the system call.
type Filter func(*Status) (AfterFilter, error)

// An AfterFilter represents a socket system call filter after an
// execution of a system call.
//
// It will only be executed after a system call for a socket that has
// an entry in internal table.
// If the filter returns a non-nil error, the system call function
// returns the non-nil error.
type AfterFilter func(*Status) error

// Set deploys the socket system call filter f for the filter type t.
func (sw *Switch) Set(t FilterType, f Filter)
