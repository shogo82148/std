// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client implementation. See RFC 2616.
//
// This is the low-level Transport implementation of RoundTripper.
// The high-level interface is in client.go.

package http

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

// DefaultTransport is the default implementation of Transport and is
// used by DefaultClient. It establishes network connections as needed
// and caches them for reuse by subsequent calls. It uses HTTP proxies
// as directed by the $HTTP_PROXY and $NO_PROXY (or $http_proxy and
// $no_proxy) environment variables.
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// DefaultMaxIdleConnsPerHost is the default value of Transport's
// MaxIdleConnsPerHost.
const DefaultMaxIdleConnsPerHost = 2

// Transport is an implementation of RoundTripper that supports HTTP,
// HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
//
// By default, Transport caches connections for future re-use.
// This may leave many open connections when accessing many hosts.
// This behavior can be managed using Transport's CloseIdleConnections method
// and the MaxIdleConnsPerHost and DisableKeepAlives fields.
//
// Transports should be reused instead of created as needed.
// Transports are safe for concurrent use by multiple goroutines.
//
// A Transport is a low-level primitive for making HTTP and HTTPS requests.
// For high-level functionality, such as cookies and redirects, see Client.
//
// Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
// for HTTPS URLs, depending on whether the server supports HTTP/2.
// See the package docs for more about HTTP/2.
type Transport struct {
	idleMu     sync.Mutex
	wantIdle   bool
	idleConn   map[connectMethodKey][]*persistConn
	idleConnCh map[connectMethodKey]chan *persistConn

	reqMu       sync.Mutex
	reqCanceler map[*Request]func()

	altMu    sync.RWMutex
	altProto map[string]RoundTripper

	Proxy func(*Request) (*url.URL, error)

	Dial func(network, addr string) (net.Conn, error)

	DialTLS func(network, addr string) (net.Conn, error)

	TLSClientConfig *tls.Config

	TLSHandshakeTimeout time.Duration

	DisableKeepAlives bool

	DisableCompression bool

	MaxIdleConnsPerHost int

	ResponseHeaderTimeout time.Duration

	ExpectContinueTimeout time.Duration

	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

	nextProtoOnce sync.Once
	h2transport   *http2Transport
}

// ProxyFromEnvironment returns the URL of the proxy to use for a
// given request, as indicated by the environment variables
// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
// thereof). HTTPS_PROXY takes precedence over HTTP_PROXY for https
// requests.
//
// The environment values may be either a complete URL or a
// "host[:port]", in which case the "http" scheme is assumed.
// An error is returned if the value is a different form.
//
// A nil URL and nil error are returned if no proxy is defined in the
// environment, or a proxy should not be used for the given request,
// as defined by NO_PROXY.
//
// As a special case, if req.URL.Host is "localhost" (with or without
// a port number), then a nil URL and nil error will be returned.
func ProxyFromEnvironment(req *Request) (*url.URL, error)

// ProxyURL returns a proxy function (for use in a Transport)
// that always returns the same URL.
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)

// transportRequest is a wrapper around a *Request that adds
// optional extra headers to write.

// RoundTrip implements the RoundTripper interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see Get, Post, and the Client type.
func (t *Transport) RoundTrip(req *Request) (*Response, error)

// ErrSkipAltProtocol is a sentinel error value defined by Transport.RegisterProtocol.
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")

// RegisterProtocol registers a new protocol with scheme.
// The Transport will pass requests using the given scheme to rt.
// It is rt's responsibility to simulate HTTP request semantics.
//
// RegisterProtocol can be used by other packages to provide
// implementations of protocol schemes like "ftp" or "file".
//
// If rt.RoundTrip returns ErrSkipAltProtocol, the Transport will
// handle the RoundTrip itself for that one request, as if the
// protocol were not registered.
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle in
// a "keep-alive" state. It does not interrupt any connections currently
// in use.
func (t *Transport) CloseIdleConnections()

// CancelRequest cancels an in-flight request by closing its connection.
// CancelRequest should only be called after RoundTrip has returned.
//
// Deprecated: Use Request.Cancel instead. CancelRequest can not cancel
// HTTP/2 requests.
func (t *Transport) CancelRequest(req *Request)

// envOnce looks up an environment variable (optionally by multiple
// names) once. It mitigates expensive lookups on some platforms
// (e.g. Windows).

// error values for debugging and testing, not seen by users.

// connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.
//
// A connect method may be of the following types:
//
// Cache key form                Description
// -----------------             -------------------------
// |http|foo.com                 http directly to server, no proxy
// |https|foo.com                https directly to server, no proxy
// http://proxy.com|https|foo.com  http to proxy, then CONNECT to foo.com
// http://proxy.com|http           http to proxy, http to anywhere after that
//
// Note: no support to https to the proxy yet.
//

// connectMethodKey is the map key version of connectMethod, with a
// stringified proxy URL (or the empty string) instead of a pointer to
// a URL.

// persistConn wraps a connection, usually a persistent one
// (but may be used for non-keep-alive requests as well)

// responseAndError is how the goroutine reading from an HTTP/1 server
// communicates with the goroutine doing the RoundTrip.

// A writeRequest is sent by the readLoop's goroutine to the
// writeLoop's goroutine to write a request while the read loop
// concurrently waits on both the write response and the server's
// reply.

// testHooks. Always non-nil.

// beforeRespHeaderError is used to indicate when an IO error has occurred before
// any header data was received.

// bodyEOFSignal wraps a ReadCloser but runs fn (if non-nil) at most
// once, right before its final (error-producing) Read or Close call
// returns. fn should return the new error to return from Read or Close.
//
// If earlyCloseFn is non-nil and Close is called before io.EOF is
// seen, earlyCloseFn is called instead of fn, and its return value is
// the return value from Close.

// gzipReader wraps a response body so it can lazily
// call gzip.NewReader on the first call to Read

// fakeLocker is a sync.Locker which does nothing. It's used to guard
// test-only fields when not under test, to avoid runtime atomic
// overhead.
