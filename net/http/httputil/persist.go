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
	ErrPersistEOF = &http.ProtocolError{ErrorString: "persistent connection closed"}
	ErrClosed     = &http.ProtocolError{ErrorString: "connection closed by user"}
	ErrPipeline   = &http.ProtocolError{ErrorString: "pipeline error"}
)

// This is an API usage error - the local side is closed.
// ErrPersistEOF (above) reports that the remote side is closed.

// A ServerConn reads requests and sends responses over an underlying
// connection, until the HTTP keepalive logic commands an end. ServerConn
// also allows hijacking the underlying connection by calling Hijack
// to regain control over the connection. ServerConn supports pipe-lining,
// i.e. requests can be read out of sync (but in the same order) while the
// respective responses are sent.
//
// ServerConn is low-level and old. Applications should instead use Server
// in the net/http package.
type ServerConn struct {
	lk              sync.Mutex
	c               net.Conn
	r               *bufio.Reader
	re, we          error
	lastbody        io.ReadCloser
	nread, nwritten int
	pipereq         map[*http.Request]uint

	pipe textproto.Pipeline
}

// NewServerConn returns a new ServerConn reading and writing c. If r is not
// nil, it is the buffer to use when reading c.
//
// ServerConn is low-level and old. Applications should instead use Server
// in the net/http package.
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn

// Hijack detaches the ServerConn and returns the underlying connection as well
// as the read-side bufio which may have some left over data. Hijack may be
// called before Read has signaled the end of the keep-alive logic. The user
// should not call Hijack while Read or Write is in progress.
func (sc *ServerConn) Hijack() (c net.Conn, r *bufio.Reader)

// Close calls Hijack and then also closes the underlying connection
func (sc *ServerConn) Close() error

// Read returns the next request on the wire. An ErrPersistEOF is returned if
// it is gracefully determined that there are no more requests (e.g. after the
// first request on an HTTP/1.0 connection, or after a Connection:close on a
// HTTP/1.1 connection).
func (sc *ServerConn) Read() (req *http.Request, err error)

// Pending returns the number of unanswered requests
// that have been received on the connection.
func (sc *ServerConn) Pending() int

// Write writes resp in response to req. To close the connection gracefully, set the
// Response.Close field to true. Write should be considered operational until
// it returns an error, regardless of any errors returned on the Read side.
func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error

// A ClientConn sends request and receives headers over an underlying
// connection, while respecting the HTTP keepalive logic. ClientConn
// supports hijacking the connection calling Hijack to
// regain control of the underlying net.Conn and deal with it as desired.
//
// ClientConn is low-level and old. Applications should instead use
// Client or Transport in the net/http package.
type ClientConn struct {
	lk              sync.Mutex
	c               net.Conn
	r               *bufio.Reader
	re, we          error
	lastbody        io.ReadCloser
	nread, nwritten int
	pipereq         map[*http.Request]uint

	pipe     textproto.Pipeline
	writeReq func(*http.Request, io.Writer) error
}

// NewClientConn returns a new ClientConn reading and writing c.  If r is not
// nil, it is the buffer to use when reading c.
//
// ClientConn is low-level and old. Applications should use Client or
// Transport in the net/http package.
func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// NewProxyClientConn works like NewClientConn but writes Requests
// using Request's WriteProxy method.
//
// New code should not use NewProxyClientConn. See Client or
// Transport in the net/http package instead.
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// Hijack detaches the ClientConn and returns the underlying connection as well
// as the read-side bufio which may have some left over data. Hijack may be
// called before the user or Read have signaled the end of the keep-alive
// logic. The user should not call Hijack while Read or Write is in progress.
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)

// Close calls Hijack and then also closes the underlying connection
func (cc *ClientConn) Close() error

// Write writes a request. An ErrPersistEOF error is returned if the connection
// has been closed in an HTTP keepalive sense. If req.Close equals true, the
// keepalive connection is logically closed after this request and the opposing
// server is informed. An ErrUnexpectedEOF indicates the remote closed the
// underlying TCP connection, which is usually considered as graceful close.
func (cc *ClientConn) Write(req *http.Request) (err error)

// Pending returns the number of unanswered requests
// that have been sent on the connection.
func (cc *ClientConn) Pending() int

// Read reads the next response from the wire. A valid response might be
// returned together with an ErrPersistEOF, which means that the remote
// requested that this be the last request serviced. Read can be called
// concurrently with Write, but not with another Read.
func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)

// Do is convenience method that writes a request and reads a response.
func (cc *ClientConn) Do(req *http.Request) (resp *http.Response, err error)
