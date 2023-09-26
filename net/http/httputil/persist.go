// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/textproto"
	"github.com/shogo82148/std/sync"
)

var (
	// Deprecated: No longer used.
	ErrPersistEOF = &http.ProtocolError{ErrorString: "persistent connection closed"}

	// Deprecated: No longer used.
	ErrClosed = &http.ProtocolError{ErrorString: "connection closed by user"}

	// Deprecated: No longer used.
	ErrPipeline = &http.ProtocolError{ErrorString: "pipeline error"}
)

// This is an API usage error - the local side is closed.
// ErrPersistEOF (above) reports that the remote side is closed.

// ServerConn is an artifact of Go's early HTTP implementation.
// It is low-level, old, and unused by Go's current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Server in package net/http instead.
type ServerConn struct {
	mu              sync.Mutex
	c               net.Conn
	r               *bufio.Reader
	re, we          error
	lastbody        io.ReadCloser
	nread, nwritten int
	pipereq         map[*http.Request]uint

	pipe textproto.Pipeline
}

// NewServerConn is an artifact of Go's early HTTP implementation.
// It is low-level, old, and unused by Go's current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Server in package net/http instead.
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn

// Hijack detaches the ServerConn and returns the underlying connection as well
// as the read-side bufio which may have some left over data. Hijack may be
// called before Read has signaled the end of the keep-alive logic. The user
// should not call Hijack while Read or Write is in progress.
func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)

// Close calls Hijack and then also closes the underlying connection.
func (sc *ServerConn) Close() error

// Read returns the next request on the wire. An ErrPersistEOF is returned if
// it is gracefully determined that there are no more requests (e.g. after the
// first request on an HTTP/1.0 connection, or after a Connection:close on a
// HTTP/1.1 connection).
func (sc *ServerConn) Read() (*http.Request, error)

// Pending returns the number of unanswered requests
// that have been received on the connection.
func (sc *ServerConn) Pending() int

// Write writes resp in response to req. To close the connection gracefully, set the
// Response.Close field to true. Write should be considered operational until
// it returns an error, regardless of any errors returned on the Read side.
func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error

// ClientConn is an artifact of Go's early HTTP implementation.
// It is low-level, old, and unused by Go's current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use Client or Transport in package net/http instead.
type ClientConn struct {
	mu              sync.Mutex
	c               net.Conn
	r               *bufio.Reader
	re, we          error
	lastbody        io.ReadCloser
	nread, nwritten int
	pipereq         map[*http.Request]uint

	pipe     textproto.Pipeline
	writeReq func(*http.Request, io.Writer) error
}

// NewClientConn is an artifact of Go's early HTTP implementation.
// It is low-level, old, and unused by Go's current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Client or Transport in package net/http instead.
func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// NewProxyClientConn is an artifact of Go's early HTTP implementation.
// It is low-level, old, and unused by Go's current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Client or Transport in package net/http instead.
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// Hijack detaches the ClientConn and returns the underlying connection as well
// as the read-side bufio which may have some left over data. Hijack may be
// called before the user or Read have signaled the end of the keep-alive
// logic. The user should not call Hijack while Read or Write is in progress.
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)

// Close calls Hijack and then also closes the underlying connection.
func (cc *ClientConn) Close() error

// Write writes a request. An ErrPersistEOF error is returned if the connection
// has been closed in an HTTP keep-alive sense. If req.Close equals true, the
// keep-alive connection is logically closed after this request and the opposing
// server is informed. An ErrUnexpectedEOF indicates the remote closed the
// underlying TCP connection, which is usually considered as graceful close.
func (cc *ClientConn) Write(req *http.Request) error

// Pending returns the number of unanswered requests
// that have been sent on the connection.
func (cc *ClientConn) Pending() int

// Read reads the next response from the wire. A valid response might be
// returned together with an ErrPersistEOF, which means that the remote
// requested that this be the last request serviced. Read can be called
// concurrently with Write, but not with another Read.
func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)

// Do is convenience method that writes a request and reads a response.
func (cc *ClientConn) Do(req *http.Request) (*http.Response, error)
