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
*/
package net

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/time"
)

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

// An Error represents a network error.
type Error interface {
	error
	Timeout() bool
	Temporary() bool
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

	Addr Addr

	Err error
}

func (e *OpError) Error() string

func (e *OpError) Temporary() bool

func (e *OpError) Timeout() bool

type AddrError struct {
	Err  string
	Addr string
}

func (e *AddrError) Error() string

func (e *AddrError) Temporary() bool

func (e *AddrError) Timeout() bool

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string
func (e UnknownNetworkError) Temporary() bool
func (e UnknownNetworkError) Timeout() bool

type InvalidAddrError string

func (e InvalidAddrError) Error() string
func (e InvalidAddrError) Timeout() bool
func (e InvalidAddrError) Temporary() bool

// DNSConfigError represents an error reading the machine's DNS configuration.
type DNSConfigError struct {
	Err error
}

func (e *DNSConfigError) Error() string

func (e *DNSConfigError) Timeout() bool
func (e *DNSConfigError) Temporary() bool
