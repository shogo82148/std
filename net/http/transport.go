// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client implementation. See RFC 7230 through 7235.
//
// This is the low-level Transport implementation of RoundTripper.
// The high-level interface is in client.go.

package http

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// DefaultTransport is the default implementation of Transport and is
// used by DefaultClient. It establishes network connections as needed
// and caches them for reuse by subsequent calls. It uses HTTP proxies
// as directed by the environment variables HTTP_PROXY, HTTPS_PROXY
// and NO_PROXY (or the lowercase versions thereof).
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: defaultTransportDialContext(&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}),
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
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
// for HTTPS URLs, depending on whether the server supports HTTP/2,
// and how the Transport is configured. The DefaultTransport supports HTTP/2.
// To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2
// and call ConfigureTransport. See the package docs for more about HTTP/2.
//
// Responses with status codes in the 1xx range are either handled
// automatically (100 expect-continue) or ignored. The one
// exception is HTTP status code 101 (Switching Protocols), which is
// considered a terminal status and returned by RoundTrip. To see the
// ignored 1xx responses, use the httptrace trace package's
// ClientTrace.Got1xxResponse.
//
// Transport only retries a request upon encountering a network error
// if the request is idempotent and either has no body or has its
// Request.GetBody defined. HTTP requests are considered idempotent if
// they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or if their
// Header map contains an "Idempotency-Key" or "X-Idempotency-Key"
// entry. If the idempotency key value is a zero-length slice, the
// request is treated as idempotent but the header is not sent on the
// wire.
type Transport struct {
	idleMu       sync.Mutex
	closeIdle    bool
	idleConn     map[connectMethodKey][]*persistConn
	idleConnWait map[connectMethodKey]wantConnQueue
	idleLRU      connLRU

	reqMu       sync.Mutex
	reqCanceler map[cancelKey]func(error)

	altMu    sync.Mutex
	altProto atomic.Value

	connsPerHostMu   sync.Mutex
	connsPerHost     map[connectMethodKey]int
	connsPerHostWait map[connectMethodKey]wantConnQueue

	Proxy func(*Request) (*url.URL, error)

	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error

	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	Dial func(network, addr string) (net.Conn, error)

	DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)

	DialTLS func(network, addr string) (net.Conn, error)

	TLSClientConfig *tls.Config

	TLSHandshakeTimeout time.Duration

	DisableKeepAlives bool

	DisableCompression bool

	MaxIdleConns int

	MaxIdleConnsPerHost int

	MaxConnsPerHost int

	IdleConnTimeout time.Duration

	ResponseHeaderTimeout time.Duration

	ExpectContinueTimeout time.Duration

	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

	ProxyConnectHeader Header

	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)

	MaxResponseHeaderBytes int64

	WriteBufferSize int

	ReadBufferSize int

	nextProtoOnce      sync.Once
	h2transport        h2Transport
	tlsNextProtoWasNil bool

	ForceAttemptHTTP2 bool
}

// A cancelKey is the key of the reqCanceler map.
// We wrap the *Request in this type since we want to use the original request,
// not any transient one created by roundTrip.

// Clone returns a deep copy of t's exported fields.
func (t *Transport) Clone() *Transport

// h2Transport is the interface we expect to be able to call from
// net/http against an *http2.Transport that's either bundled into
// h2_bundle.go or supplied by the user via x/net/http2.
//
// We name it with the "h2" prefix to stay out of the "http2" prefix
// namespace used by x/tools/cmd/bundle for h2_bundle.go.

// ProxyFromEnvironment returns the URL of the proxy to use for a
// given request, as indicated by the environment variables
// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
// thereof). Requests use the proxy from the environment variable
// matching their scheme, unless excluded by NO_PROXY.
//
// The environment values may be either a complete URL or a
// "host[:port]", in which case the "http" scheme is assumed.
// The schemes "http", "https", and "socks5" are supported.
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
// optional extra headers to write and stores any error to return
// from roundTrip.

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
// Deprecated: Use Request.WithContext to create a request with a
// cancelable context instead. CancelRequest cannot cancel HTTP/2
// requests.
func (t *Transport) CancelRequest(req *Request)

// error values for debugging and testing, not seen by users.

// transportReadFromServerError is used by Transport.readLoop when the
// 1 byte peek read fails and we're actually anticipating a response.
// Usually this is just due to the inherent keep-alive shut down race,
// where the server closed the connection at the same time the client
// wrote. The underlying err field is usually io.EOF or some
// ECONNRESET sort of thing which varies by platform. But it might be
// the user's custom net.Conn.Read error too, so we carry it along for
// them to return from Transport.RoundTrip.

// A wantConn records state about a wanted connection
// (that is, an active call to getConn).
// The conn may be gotten by dialing or by finding an idle connection,
// or a cancellation may make the conn no longer wanted.
// These three options are racing against each other and use
// wantConn to coordinate and agree about the winning outcome.

// A wantConnQueue is a queue of wantConns.

// persistConnWriter is the io.Writer written to by pc.bw.
// It accumulates the number of bytes written to the underlying conn,
// so the retry logic can determine whether any bytes made it across
// the wire.
// This is exactly 1 pointer field wide so it can go into an interface
// without allocation.

var _ io.ReaderFrom = (*persistConnWriter)(nil)

// connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.
//
// A connect method may be of the following types:
//
//	connectMethod.key().String()      Description
//	------------------------------    -------------------------
//	|http|foo.com                     http directly to server, no proxy
//	|https|foo.com                    https directly to server, no proxy
//	|https,h1|foo.com                 https directly to server w/o HTTP/2, no proxy
//	http://proxy.com|https|foo.com    http to proxy, then CONNECT to foo.com
//	http://proxy.com|http             http to proxy, http to anywhere after that
//	socks5://proxy.com|http|foo.com   socks5 to proxy, then http to foo.com
//	socks5://proxy.com|https|foo.com  socks5 to proxy, then https to foo.com
//	https://proxy.com|https|foo.com   https to proxy, then CONNECT to foo.com
//	https://proxy.com|http            https to proxy, http to anywhere after that

// connectMethodKey is the map key version of connectMethod, with a
// stringified proxy URL (or the empty string) instead of a pointer to
// a URL.

// persistConn wraps a connection, usually a persistent one
// (but may be used for non-keep-alive requests as well)

// errCallerOwnsConn is an internal sentinel error used when we hand
// off a writable response.Body to the caller. We use this to prevent
// closing a net.Conn that is now owned by the caller.

// readWriteCloserBody is the Response.Body type used when we want to
// give users write access to the Body through the underlying
// connection (TCP, unless using custom dialers). This is then
// the concrete type for a Response.Body on the 101 Switching
// Protocols response, as used by WebSockets, h2c, etc.

// nothingWrittenError wraps a write errors which ended up writing zero bytes.

// maxWriteWaitBeforeConnReuse is how long the a Transport RoundTrip
// will wait to see the Request's Body.Write result after getting a
// response from the server. See comments in (*persistConn).wroteRequest.

// responseAndError is how the goroutine reading from an HTTP/1 server
// communicates with the goroutine doing the RoundTrip.

// A writeRequest is sent by the caller's goroutine to the
// writeLoop's goroutine to write a request while the read loop
// concurrently waits on both the write response and the server's
// reply.

// errRequestCanceled is set to be identical to the one from h2 to facilitate
// testing.

// testHooks. Always non-nil.

// tLogKey is a context WithValue key for test debugging contexts containing
// a t.Logf func. See export_test.go's Request.WithT method.

// bodyEOFSignal is used by the HTTP/1 transport when reading response
// bodies to make sure we see the end of a response body before
// proceeding and reading on the connection again.
//
// It wraps a ReadCloser but runs fn (if non-nil) at most
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
