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

	// EnableHTTP2 controls whether HTTP/2 is enabled
	// on the server. It must be set between calling
	// NewUnstartedServer and calling Server.StartTLS.
	EnableHTTP2 bool

	// TLS is the optional TLS configuration, populated with a new config
	// after TLS is started. If set on an unstarted server before StartTLS
	// is called, existing fields are copied into the new config.
	TLS *tls.Config

	// Config may be changed after calling NewUnstartedServer and
	// before Start or StartTLS.
	Config *http.Server

	// certificate is a parsed version of the TLS config certificate, if present.
	certificate *x509.Certificate

	// wg counts the number of outstanding HTTP requests on this server.
	// Close blocks until all requests are finished.
	wg sync.WaitGroup

	mu     sync.Mutex
	closed bool
	conns  map[net.Conn]http.ConnState

	// client is configured for use with the server.
	// Its transport is automatically closed when Close is called.
	client *http.Client
}

// NewServer starts and returns a new [Server].
// The caller should call Close when finished, to shut it down.
func NewServer(handler http.Handler) *Server

// NewUnstartedServer returns a new [Server] but doesn't start it.
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

// NewTLSServer starts and returns a new [Server] using TLS.
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
// close its idle connections on [Server.Close].
// Use Server.URL as the base URL to send requests to the server.
// The returned client will also redirect any requests to "example.com"
// or its subdomains to the server.
func (s *Server) Client() *http.Client
