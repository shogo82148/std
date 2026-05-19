// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of Server

package httptest

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/internal/nettest"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/testing"
)

// A Server is an HTTP server for use in end-to-end HTTP tests.
//
// Most tests should create a server with [NewTestServer].
// The [Server.Client] method returns a client which sends requests to the test server.
//
//	// Create a test server and send a request to it.
//	server := httptest.NewTestServer(t, handler)
//	resp, err := server.Client().Get("http://www.example.com/")
//
// # Configuration
//
// Tests may change a Server's configuration prior to using it.
// The configuration must not be changed after the first call to
// [Server.Client], [Server.Start], or [Server.StartTLS].
//
//	// Configure a test server before using.
//	server := httptest.NewTestServer(t, handler)
//	server.Config.MaxHeaderBytes = 1024
//	resp, err := server.Client().Get("http://www.example.com/")
//
// # Tests
//
// Servers created with [NewTestServer] will:
//
//   - Fail the test if the server handler panics with
//     any value other than [http.ErrAbortHandler].
//   - Register a Cleanup function to shut down the server at the end of the test.
//
// Servers created in any other way must be manually shut down with [Server.Close].
//
// # In-Memory Network
//
// A Server may use an in-memory network implementation or
// listen on a local network loopback interface.
// Most tests should use the in-memory network,
// which avoids port exhaustion and other transient networking issues
// and is suitable for use with the [testing/synctest] package.
//
// To use the in-memory network, create a server with [NewTestServer].
// Do not call [Server.Start] or [Server.StartTLS].
//
// When using the in-memory network, the [http.Client] returned by [Server.Client]
// is configured to send all requests to the server.
// The client will direct HTTP and HTTPS requests,
// regardless of destination address or hostname, to the server.
// Requests do not need to use [Server.URL] as the base URL.
//
//	server := httptest.NewTestServer(t, handler)
//	client := server.Client()
//
//	// All of these requests are sent to the test server.
//	// https:// requests use TLS over the in-memory network.
//	_, _ = client.Get("http://www.example.com/")
//	_, _ = client.Get("https://go.dev/")
//	_, _ = client.Get("http://10.0.0.1/")
//
// The [Server.Listener] field is not set when using the in-memory network.
//
// # Loopback Network
//
// To listen on a loopback interface, call [Server.Start] or [Server.StartTLS].
// The server will listen on a system-chosen port.
//
// Loopback servers serve one of HTTP (when started with [Server.Start])
// or HTTPS (when started with [Server.StartTLS]).
//
// When using the loopback network, the [http.Client] returned by [Server.Client]
// is configured to send requests with a hostname of "example.com" or a subdomain
// of ".example.com" to the server.
//
// Requests may also be sent to the server's loopback address.
// The [Server.URL] field is set to a base URL containing the server's address.
//
//	server := httptest.NewTestServer(t, handler)
//	server.Start()
//	client := server.Client()
//
//	// This request is sent to the test server.
//	_, _ = server.Client().Get(server.URL + "/")
//
//	// This request (using http.DefaultClient) is also sent to the test server,
//	// since server.URL contains the server's local IP address.
//	_, _ = http.Get(server.URL + "/")
type Server struct {
	// URL is the base URL of the server, of the form http://address:port
	// with no trailing slash.
	//
	// It is set by the first call to Client, Start, or StartTLS.
	//
	// For servers listening on loopback, the address is the loopback IP address
	// of the server.
	//
	// For servers using the in-memory network, this address is "example.com".
	// Requests sent to servers using the in-memory network may use any address.
	// It is not necessary to use this base URL.
	URL string

	// Listener is the network listener for servers listening on loopback.
	// It is not set for servers using the in-memory network.
	Listener net.Listener

	// EnableHTTP2 controls whether HTTP/2 is enabled on the server.
	// It must be set before calling Client, Start, or StartTLS.
	EnableHTTP2 bool

	// TLS is the optional TLS configuration, populated with a new config
	// after TLS is started. If set on an unstarted server before StartTLS
	// is called, existing fields are copied into the new config.
	TLS *tls.Config

	// Config may be changed before calling Client, Start, or StartTLS.
	Config *http.Server

	t testing.TB

	// certificate is a parsed version of the TLS config certificate, if present.
	certificate *x509.Certificate

	// startOnce is used to start fakenet servers once.
	startOnce sync.Once

	// started indicates whether the server has been started.
	started bool

	// Fake network listeners, one for HTTP and one for HTTPS.
	fakeListener    *nettest.Listener
	fakeTLSListener *nettest.Listener

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

// NewTestServer returns a new [Server] for a test.
// The server will use an in-memory network implementation by default.
//
// If the handler is nil, the server will serve 500 responses to all requests.
// It will not use [http.DefaultServeMux].
//
// See the [Server] documentation for more details.
func NewTestServer(t testing.TB, handler http.Handler) *Server

// NewServer starts and returns a new [Server] listening on a
// local network loopback interface.
// This is equivalent to calling [NewUnstartedServer] followed by [Server.Start].
//
// The caller should call [Server.Close] when finished, to shut it down.
//
// Most users should use [NewTestServer] instead.
// See the [Server] documentation for details.
func NewServer(handler http.Handler) *Server

// NewUnstartedServer returns a new [Server] listening on a
// local network loopback interface. It does not start the server.
//
// After changing the server's configuration, the caller should
// call [Server.Start] or [Server.StartTLS].
//
// The caller should call [Server.Close] when finished, to shut it down.
//
// Most users should use [NewTestServer] instead.
// See the [Server] documentation for details.
func NewUnstartedServer(handler http.Handler) *Server

// Start starts a server on a local loopback network interface.
//
// The server should have been created by [NewTestServer] or [NewUnstartedServer].
func (s *Server) Start()

// Start starts TLS on a server on a local loopback network interface.
//
// The server should have been created by [NewTestServer] or [NewUnstartedServer].
func (s *Server) StartTLS()

// NewTLSServer starts and returns a new [Server] using TLS and listening on a
// local network loopback interface.
// This is equivalent to calling [NewUnstartedServer] followed by [Server.StartTLS].
//
// The caller should call [Server.Close] when finished, to shut it down.
//
// Most users should use [NewTestServer] instead.
// See the [Server] documentation for details.
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
func (s *Server) Client() *http.Client
