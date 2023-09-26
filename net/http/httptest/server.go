// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of Server

package httptest

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
)

// A Server is an HTTP server listening on a system-chosen port on the
// local loopback interface, for use in end-to-end HTTP tests.
type Server struct {
	URL      string
	Listener net.Listener

	EnableHTTP2 bool

	TLS *tls.Config

	Config *http.Server

	certificate *x509.Certificate

	wg sync.WaitGroup

	mu     sync.Mutex
	closed bool
	conns  map[net.Conn]http.ConnState

	client *http.Client
}

// When debugging a particular http server-based test,
// this flag lets you run
//
//	go test -run=BrokenTest -httptest.serve=127.0.0.1:8000
//
// to start the broken server so you can interact with it manually.
// We only register this flag if it looks like the caller knows about it
// and is trying to use it as we don't want to pollute flags and this
// isn't really part of our API. Don't depend on this.

// NewServer starts and returns a new Server.
// The caller should call Close when finished, to shut it down.
func NewServer(handler http.Handler) *Server

// NewUnstartedServer returns a new Server but doesn't start it.
//
// After changing its configuration, the caller should call Start or
// StartTLS.
//
// The caller should call Close when finished, to shut it down.
func NewUnstartedServer(handler http.Handler) *Server

// Start starts a server from NewUnstartedServer.
func (s *Server) Start()

// StartTLS starts TLS on a server from NewUnstartedServer.
func (s *Server) StartTLS()

// NewTLSServer starts and returns a new Server using TLS.
// The caller should call Close when finished, to shut it down.
func NewTLSServer(handler http.Handler) *Server

// Close shuts down the server and blocks until all outstanding
// requests on this server have completed.
func (s *Server) Close()

// CloseClientConnections closes any open HTTP connections to the test Server.
func (s *Server) CloseClientConnections()

// Certificate returns the certificate used by the server, or nil if
// the server doesn't use TLS.
func (s *Server) Certificate() *x509.Certificate

// Client returns an HTTP client configured for making requests to the server.
// It is configured to trust the server's TLS test certificate and will
// close its idle connections on Server.Close.
func (s *Server) Client() *http.Client
