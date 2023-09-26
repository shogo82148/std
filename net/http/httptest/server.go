// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of Server

package httptest

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
)

// A Server is an HTTP server listening on a system-chosen port on the
// local loopback interface, for use in end-to-end HTTP tests.
type Server struct {
	URL      string
	Listener net.Listener

	TLS *tls.Config

	Config *http.Server

	wg sync.WaitGroup
}

// historyListener keeps track of all connections that it's ever
// accepted.

// When debugging a particular http server-based test,
// this flag lets you run
//	go test -run=BrokenTest -httptest.serve=127.0.0.1:8000
// to start the broken server so you can interact with it manually.

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

// CloseClientConnections closes any currently open HTTP connections
// to the test Server.
func (s *Server) CloseClientConnections()

// waitGroupHandler wraps a handler, incrementing and decrementing a
// sync.WaitGroup on each request, to enable Server.Close to block
// until outstanding requests are finished.

// localhostCert is a PEM-encoded TLS cert with SAN IPs
// "127.0.0.1" and "[::1]", expiring at the last second of 2049 (the end
// of ASN.1 time).
// generated from src/crypto/tls:
// go run generate_cert.go  --rsa-bits 512 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h

// localhostKey is the private key for localhostCert.
