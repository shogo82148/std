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

	conn, err := net.Dial("tcp", "golang.org:80")
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
	export GODEBUG=netdns=cgo   # force native resolver (cgo, win32)

The decision can also be forced while building the Go source tree
by setting the netgo or netcgo build tag.

A numeric netdns setting, as in GODEBUG=netdns=1, causes the resolver
to print debugging information about its decisions.
To force a particular resolver while also printing debugging information,
join the two settings by a plus sign, as in GODEBUG=netdns=go+1.

The Go resolver will send an EDNS0 additional header with a DNS request,
to signal a willingness to accept a larger DNS packet size.
This can reportedly cause sporadic failures with the DNS server run
by some modems and routers. Setting GODEBUG=netedns0=0 will disable
sending the additional header.

On macOS, if Go code that uses the net package is built with
-buildmode=c-archive, linking the resulting archive into a C program
requires passing -lresolv when linking the C code.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

On Windows, in Go 1.18.x and earlier, the resolver always used C
library functions, such as GetAddrInfo and DnsQuery.
*/
package net

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// Addr represents a network end point address.
//
// The two methods Network and String conventionally return strings
// that can be passed as the arguments to Dial, but the exact form
// and meaning of the strings is up to the implementation.
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
	ReadFrom(p []byte) (n int, addr Addr, err error)

	WriteTo(p []byte, addr Addr) (n int, err error)

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
	Accept() (Conn, error)

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
	// Op is the operation which caused the error, such as
	// "read" or "write".
	Op string

	// Net is the network type on which this error occurred,
	// such as "tcp" or "udp6".
	Net string

	// For operations involving a remote network connection, like
	// Dial, Read, or Write, Source is the corresponding local
	// network address.
	Source Addr

	// Addr is the network address for which this error occurred.
	// For local operations, like Listen or SetDeadline, Addr is
	// the address of the local endpoint being manipulated.
	// For operations involving a remote network connection, like
	// Dial, Read, or Write, Addr is the remote address of that
	// connection.
	Addr Addr

	// Err is the error that occurred during the operation.
	// The Error method panics if the error is nil.
	Err error
}

func (e *OpError) Unwrap() error

func (e *OpError) Error() string

func (e *OpError) Timeout() bool

func (e *OpError) Temporary() bool

// A ParseError is the error type of literal network address parsers.
type ParseError struct {
	// Type is the type of string that was expected, such as
	// "IP address", "CIDR address".
	Type string

	// Text is the malformed text string.
	Text string
}

func (e *ParseError) Error() string

func (e *ParseError) Timeout() bool
func (e *ParseError) Temporary() bool

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

func (e *DNSConfigError) Unwrap() error
func (e *DNSConfigError) Error() string
func (e *DNSConfigError) Timeout() bool
func (e *DNSConfigError) Temporary() bool

// DNSError represents a DNS lookup error.
type DNSError struct {
	Err         string
	Name        string
	Server      string
	IsTimeout   bool
	IsTemporary bool
	IsNotFound  bool
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

// ErrClosed is the error returned by an I/O call on a network
// connection that has already been closed, or that is closed by
// another goroutine before the I/O is completed. This may be wrapped
// in another error, and should normally be tested using
// errors.Is(err, net.ErrClosed).
var ErrClosed error = errClosed

// Buffers contains zero or more runs of bytes to write.
//
// On certain machines, for certain types of connections, this is
// optimized into an OS-specific batch write operation (such as
// "writev").
type Buffers [][]byte

var (
	_ io.WriterTo = (*Buffers)(nil)
	_ io.Reader   = (*Buffers)(nil)
)

// WriteTo writes contents of the buffers to w.
//
// WriteTo implements io.WriterTo for Buffers.
//
// WriteTo modifies the slice v as well as v[i] for 0 <= i < len(v),
// but does not modify v[i][j] for any i, j.
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)

// Read from the buffers.
//
// Read implements io.Reader for Buffers.
//
// Read modifies the slice v as well as v[i] for 0 <= i < len(v),
// but does not modify v[i][j] for any i, j.
func (v *Buffers) Read(p []byte) (n int, err error)
