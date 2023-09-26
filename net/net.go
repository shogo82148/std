// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package net provides a portable interface for network I/O, including
TCP/IP, UDP, domain name resolution, and Unix domain sockets.

Although the package provides access to low-level networking
primitives, most clients will need only the basic interface provided
by the Dial, Listen, and Accept functions and the associated
Conn and Listener interfaces. The crypto/tls package uses
the same interfaces and similar Dial and Listen functions.

The Dial function connects to a server:

	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...

The Listen function creates servers:

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

# Name Resolution

The method for resolving domain names, whether indirectly with functions like Dial
or directly with functions like LookupHost and LookupAddr, varies by operating system.

On Unix systems, the resolver has two options for resolving names.
It can use a pure Go resolver that sends DNS requests directly to the servers
listed in /etc/resolv.conf, or it can use a cgo-based resolver that calls C
library routines such as getaddrinfo and getnameinfo.

By default the pure Go resolver is used, because a blocked DNS request consumes
only a goroutine, while a blocked C call consumes an operating system thread.
When cgo is available, the cgo-based resolver is used instead under a variety of
conditions: on systems that do not let programs make direct DNS requests (OS X),
when the LOCALDOMAIN environment variable is present (even if empty),
when the RES_OPTIONS or HOSTALIASES environment variable is non-empty,
when the ASR_CONFIG environment variable is non-empty (OpenBSD only),
when /etc/resolv.conf or /etc/nsswitch.conf specify the use of features that the
Go resolver does not implement, and when the name being looked up ends in .local
or is an mDNS name.

The resolver decision can be overridden by setting the netdns value of the
GODEBUG environment variable (see package runtime) to go or cgo, as in:

	export GODEBUG=netdns=go    # force pure Go resolver
	export GODEBUG=netdns=cgo   # force cgo resolver

The decision can also be forced while building the Go source tree
by setting the netgo or netcgo build tag.

A numeric netdns setting, as in GODEBUG=netdns=1, causes the resolver
to print debugging information about its decisions.
To force a particular resolver while also printing debugging information,
join the two settings by a plus sign, as in GODEBUG=netdns=go+1.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

On Windows, the resolver always uses C library functions, such as GetAddrInfo and DnsQuery.
*/
package net

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/time"
)

// netGo and netCgo contain the state of the build tags used
// to build this binary, and whether cgo is available.
// conf.go mirrors these into conf for easier testing.

// Addr represents a network end point address.
type Addr interface {
	Network() string
	String() string
}

// Conn is a generic stream-oriented network connection.
//
// Multiple goroutines may invoke methods on a Conn simultaneously.
type Conn interface {
	Read(b []byte) (n int, err error)

	Write(b []byte) (n int, err error)

	Close() error

	LocalAddr() Addr

	RemoteAddr() Addr

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}

// PacketConn is a generic packet-oriented network connection.
//
// Multiple goroutines may invoke methods on a PacketConn simultaneously.
type PacketConn interface {
	ReadFrom(b []byte) (n int, addr Addr, err error)

	WriteTo(b []byte, addr Addr) (n int, err error)

	Close() error

	LocalAddr() Addr

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}

// A Listener is a generic network listener for stream-oriented protocols.
//
// Multiple goroutines may invoke methods on a Listener simultaneously.
type Listener interface {
	Accept() (c Conn, err error)

	Close() error

	Addr() Addr
}

// An Error represents a network error.
type Error interface {
	error
	Timeout() bool
	Temporary() bool
}

// Various errors contained in OpError.
var (
	ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
)

// OpError is the error type usually returned by functions in the net
// package. It describes the operation, network type, and address of
// an error.
type OpError struct {
	Op string

	Net string

	Source Addr

	Addr Addr

	Err error
}

func (e *OpError) Error() string

func (e *OpError) Timeout() bool

func (e *OpError) Temporary() bool

// A ParseError is the error type of literal network address parsers.
type ParseError struct {
	Type string

	Text string
}

func (e *ParseError) Error() string

type AddrError struct {
	Err  string
	Addr string
}

func (e *AddrError) Error() string

func (e *AddrError) Timeout() bool
func (e *AddrError) Temporary() bool

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string
func (e UnknownNetworkError) Timeout() bool
func (e UnknownNetworkError) Temporary() bool

type InvalidAddrError string

func (e InvalidAddrError) Error() string
func (e InvalidAddrError) Timeout() bool
func (e InvalidAddrError) Temporary() bool

// DNSConfigError represents an error reading the machine's DNS configuration.
// (No longer used; kept for compatibility.)
type DNSConfigError struct {
	Err error
}

func (e *DNSConfigError) Error() string
func (e *DNSConfigError) Timeout() bool
func (e *DNSConfigError) Temporary() bool

// Various errors contained in DNSError.

// DNSError represents a DNS lookup error.
type DNSError struct {
	Err       string
	Name      string
	Server    string
	IsTimeout bool
}

func (e *DNSError) Error() string

// Timeout reports whether the DNS lookup is known to have timed out.
// This is not always known; a DNS lookup may fail due to a timeout
// and return a DNSError for which Timeout returns false.
func (e *DNSError) Timeout() bool

// Temporary reports whether the DNS error is known to be temporary.
// This is not always known; a DNS lookup may fail due to a temporary
// error and return a DNSError for which Temporary returns false.
func (e *DNSError) Temporary() bool
