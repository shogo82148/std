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
var DefaultTransport RoundTripper = &Transport{Proxy: ProxyFromEnvironment}

// DefaultMaxIdleConnsPerHost is the default value of Transport's
// MaxIdleConnsPerHost.
const DefaultMaxIdleConnsPerHost = 2

// Transport is an implementation of RoundTripper that supports http,
// https, and http proxies (for either http or https with CONNECT).
// Transport can also cache connections for future re-use.
type Transport struct {
	idleMu     sync.Mutex
	idleConn   map[string][]*persistConn
	idleConnCh map[string]chan *persistConn
	reqMu      sync.Mutex
	reqConn    map[*Request]*persistConn
	altMu      sync.RWMutex
	altProto   map[string]RoundTripper

	Proxy func(*Request) (*url.URL, error)

	Dial func(network, addr string) (net.Conn, error)

	TLSClientConfig *tls.Config

	DisableKeepAlives bool

	DisableCompression bool

	MaxIdleConnsPerHost int

	ResponseHeaderTimeout time.Duration
}

// ProxyFromEnvironment returns the URL of the proxy to use for a
// given request, as indicated by the environment variables
// $HTTP_PROXY and $NO_PROXY (or $http_proxy and $no_proxy).
// An error is returned if the proxy environment is invalid.
// A nil URL and nil error are returned if no proxy is defined in the
// environment, or a proxy should not be used for the given request.
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
func (t *Transport) RoundTrip(req *Request) (resp *Response, err error)

// RegisterProtocol registers a new protocol with scheme.
// The Transport will pass requests using the given scheme to rt.
// It is rt's responsibility to simulate HTTP request semantics.
//
// RegisterProtocol can be used by other packages to provide
// implementations of protocol schemes like "ftp" or "file".
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle in
// a "keep-alive" state. It does not interrupt any connections currently
// in use.
func (t *Transport) CloseIdleConnections()

// CancelRequest cancels an in-flight request by closing its
// connection.
func (t *Transport) CancelRequest(req *Request)

// connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.
//
// A connect method may be of the following types:
//
// Cache key form                Description
// -----------------             -------------------------
// ||http|foo.com                http directly to server, no proxy
// ||https|foo.com               https directly to server, no proxy
// http://proxy.com|https|foo.com  http to proxy, then CONNECT to foo.com
// http://proxy.com|http           http to proxy, http to anywhere after that
//
// Note: no support to https to the proxy yet.
//

// persistConn wraps a connection, usually a persistent one
// (but may be used for non-keep-alive requests as well)

// A writeRequest is sent by the readLoop's goroutine to the
// writeLoop's goroutine to write a request while the read loop
// concurrently waits on both the write response and the server's
// reply.

// bodyEOFSignal wraps a ReadCloser but runs fn (if non-nil) at most
// once, right before its final (error-producing) Read or Close call
// returns. If earlyCloseFn is non-nil and Close is called before
// io.EOF is seen, earlyCloseFn is called instead of fn, and its
// return value is the return value from Close.
