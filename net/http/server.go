// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP server. See RFC 7230 through 7235.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net"

	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// Errors used by the HTTP server.
var (
	// ErrBodyNotAllowed is returned by ResponseWriter.Write calls
	// when the HTTP method or response code does not permit a
	// body.
	ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")

	// ErrHijacked is returned by ResponseWriter.Write calls when
	// the underlying connection has been hijacked using the
	// Hijacker interface. A zero-byte write on a hijacked
	// connection will return ErrHijacked without any other side
	// effects.
	ErrHijacked = errors.New("http: connection has been hijacked")

	// ErrContentLength is returned by ResponseWriter.Write calls
	// when a Handler set a Content-Length response header with a
	// declared size and then attempted to write more bytes than
	// declared.
	ErrContentLength = errors.New("http: wrote more than the declared Content-Length")

	// Deprecated: ErrWriteAfterFlush is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrWriteAfterFlush = errors.New("unused")
)

// A Handler responds to an HTTP request.
//
// [Handler.ServeHTTP] should write reply headers and data to the [ResponseWriter]
// and then return. Returning signals that the request is finished; it
// is not valid to use the [ResponseWriter] or read from the
// [Request.Body] after or concurrently with the completion of the
// ServeHTTP call.
//
// Depending on the HTTP client software, HTTP protocol version, and
// any intermediaries between the client and the Go server, it may not
// be possible to read from the [Request.Body] after writing to the
// [ResponseWriter]. Cautious handlers should read the [Request.Body]
// first, and then reply.
//
// Except for reading the body, handlers should not modify the
// provided Request.
//
// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
// that the effect of the panic was isolated to the active request.
// It recovers the panic, logs a stack trace to the server error log,
// and either closes the network connection or sends an HTTP/2
// RST_STREAM, depending on the HTTP protocol. To abort a handler so
// the client sees an interrupted response but the server doesn't log
// an error, panic with the value [ErrAbortHandler].
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// A ResponseWriter interface is used by an HTTP handler to
// construct an HTTP response.
//
// A ResponseWriter may not be used after [Handler.ServeHTTP] has returned.
type ResponseWriter interface {
	Header() Header

	Write([]byte) (int, error)

	WriteHeader(statusCode int)
}

// The Flusher interface is implemented by ResponseWriters that allow
// an HTTP handler to flush buffered data to the client.
//
// The default HTTP/1.x and HTTP/2 [ResponseWriter] implementations
// support [Flusher], but ResponseWriter wrappers may not. Handlers
// should always test for this ability at runtime.
//
// Note that even for ResponseWriters that support Flush,
// if the client is connected through an HTTP proxy,
// the buffered data may not reach the client until the response
// completes.
type Flusher interface {
	Flush()
}

// The Hijacker interface is implemented by ResponseWriters that allow
// an HTTP handler to take over the connection.
//
// The default [ResponseWriter] for HTTP/1.x connections supports
// Hijacker, but HTTP/2 connections intentionally do not.
// ResponseWriter wrappers may also not support Hijacker. Handlers
// should always test for this ability at runtime.
type Hijacker interface {
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}

// The CloseNotifier interface is implemented by ResponseWriters which
// allow detecting when the underlying connection has gone away.
//
// This mechanism can be used to cancel long operations on the server
// if the client has disconnected before the response is ready.
//
// Deprecated: the CloseNotifier interface predates Go's context package.
// New code should use [Request.Context] instead.
type CloseNotifier interface {
	CloseNotify() <-chan bool
}

var (
	// ServerContextKey is a context key. It can be used in HTTP
	// handlers with Context.Value to access the server that
	// started the handler. The associated value will be of
	// type *Server.
	ServerContextKey = &contextKey{"http-server"}

	// LocalAddrContextKey is a context key. It can be used in
	// HTTP handlers with Context.Value to access the local
	// address the connection arrived on.
	// The associated value will be of type net.Addr.
	LocalAddrContextKey = &contextKey{"local-addr"}
)

// TrailerPrefix is a magic prefix for [ResponseWriter.Header] map keys
// that, if present, signals that the map entry is actually for
// the response trailers, and not the response headers. The prefix
// is stripped after the ServeHTTP call finishes and the values are
// sent in the trailers.
//
// This mechanism is intended only for trailers that are not known
// prior to the headers being written. If the set of trailers is fixed
// or known before the header is written, the normal Go trailers mechanism
// is preferred:
//
//	https://pkg.go.dev/net/http#ResponseWriter
//	https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
const TrailerPrefix = "Trailer:"

// DefaultMaxHeaderBytes is the maximum permitted size of the headers
// in an HTTP request.
// This can be overridden by setting [Server.MaxHeaderBytes].
const DefaultMaxHeaderBytes = 1 << 20

// TimeFormat is the time format to use when generating times in HTTP
// headers. It is like [time.RFC1123] but hard-codes GMT as the time
// zone. The time being formatted must be in UTC for Format to
// generate the correct format.
//
// For parsing this time format, see [ParseTime].
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

var _ closeWriter = (*net.TCPConn)(nil)

// ErrAbortHandler is a sentinel panic value to abort a handler.
// While any panic from ServeHTTP aborts the response to the client,
// panicking with ErrAbortHandler also suppresses logging of a stack
// trace to the server's error log.
var ErrAbortHandler = errors.New("net/http: abort Handler")

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// [Handler] that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

// Error replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be plain text.
//
// Error deletes the Content-Length header,
// sets Content-Type to “text/plain; charset=utf-8”,
// and sets X-Content-Type-Options to “nosniff”.
// This configures the header properly for the error message,
// in case the caller had set it up expecting a successful output.
func Error(w ResponseWriter, error string, code int)

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(w ResponseWriter, r *Request)

// NotFoundHandler returns a simple request handler
// that replies to each request with a “404 page not found” reply.
func NotFoundHandler() Handler

// StripPrefix returns a handler that serves HTTP requests by removing the
// given prefix from the request URL's Path (and RawPath if set) and invoking
// the handler h. StripPrefix handles a request for a path that doesn't begin
// with prefix by replying with an HTTP 404 not found error. The prefix must
// match exactly: if the prefix in the request contains escaped characters
// the reply is also an HTTP 404 not found error.
func StripPrefix(prefix string, h Handler) Handler

// Redirect replies to the request with a redirect to url,
// which may be a path relative to the request path.
// Any non-ASCII characters in url will be percent-encoded,
// but existing percent encodings will not be changed.
//
// The provided code should be in the 3xx range and is usually
// [StatusMovedPermanently], [StatusFound] or [StatusSeeOther].
//
// If the Content-Type header has not been set, [Redirect] sets it
// to "text/html; charset=utf-8" and writes a small HTML body.
// Setting the Content-Type header to any value, including nil,
// disables that behavior.
func Redirect(w ResponseWriter, r *Request, url string, code int)

// RedirectHandler returns a request handler that redirects
// each request it receives to the given url using the given
// status code.
//
// The provided code should be in the 3xx range and is usually
// [StatusMovedPermanently], [StatusFound] or [StatusSeeOther].
func RedirectHandler(url string, code int) Handler

// ServeMux is an HTTP request multiplexer.
// It matches the URL of each incoming request against a list of registered
// patterns and calls the handler for the pattern that
// most closely matches the URL.
//
// # Patterns
//
// Patterns can match the method, host and path of a request.
// Some examples:
//
//   - "/index.html" matches the path "/index.html" for any host and method.
//   - "GET /static/" matches a GET request whose path begins with "/static/".
//   - "example.com/" matches any request to the host "example.com".
//   - "example.com/{$}" matches requests with host "example.com" and path "/".
//   - "/b/{bucket}/o/{objectname...}" matches paths whose first segment is "b"
//     and whose third segment is "o". The name "bucket" denotes the second
//     segment and "objectname" denotes the remainder of the path.
//
// In general, a pattern looks like
//
//	[METHOD ][HOST]/[PATH]
//
// All three parts are optional; "/" is a valid pattern.
// If METHOD is present, it must be followed by at least one space or tab.
//
// Literal (that is, non-wildcard) parts of a pattern match
// the corresponding parts of a request case-sensitively.
//
// A pattern with no method matches every method. A pattern
// with the method GET matches both GET and HEAD requests.
// Otherwise, the method must match exactly.
//
// A pattern with no host matches every host.
// A pattern with a host matches URLs on that host only.
//
// A path can include wildcard segments of the form {NAME} or {NAME...}.
// For example, "/b/{bucket}/o/{objectname...}".
// The wildcard name must be a valid Go identifier.
// Wildcards must be full path segments: they must be preceded by a slash and followed by
// either a slash or the end of the string.
// For example, "/b_{bucket}" is not a valid pattern.
//
// Normally a wildcard matches only a single path segment,
// ending at the next literal slash (not %2F) in the request URL.
// But if the "..." is present, then the wildcard matches the remainder of the URL path, including slashes.
// (Therefore it is invalid for a "..." wildcard to appear anywhere but at the end of a pattern.)
// The match for a wildcard can be obtained by calling [Request.PathValue] with the wildcard's name.
// A trailing slash in a path acts as an anonymous "..." wildcard.
//
// The special wildcard {$} matches only the end of the URL.
// For example, the pattern "/{$}" matches only the path "/",
// whereas the pattern "/" matches every path.
//
// For matching, both pattern paths and incoming request paths are unescaped segment by segment.
// So, for example, the path "/a%2Fb/100%25" is treated as having two segments, "a/b" and "100%".
// The pattern "/a%2fb/" matches it, but the pattern "/a/b/" does not.
//
// # Precedence
//
// If two or more patterns match a request, then the most specific pattern takes precedence.
// A pattern P1 is more specific than P2 if P1 matches a strict subset of P2’s requests;
// that is, if P2 matches all the requests of P1 and more.
// If neither is more specific, then the patterns conflict.
// There is one exception to this rule, for backwards compatibility:
// if two patterns would otherwise conflict and one has a host while the other does not,
// then the pattern with the host takes precedence.
// If a pattern passed to [ServeMux.Handle] or [ServeMux.HandleFunc] conflicts with
// another pattern that is already registered, those functions panic.
//
// As an example of the general rule, "/images/thumbnails/" is more specific than "/images/",
// so both can be registered.
// The former matches paths beginning with "/images/thumbnails/"
// and the latter will match any other path in the "/images/" subtree.
//
// As another example, consider the patterns "GET /" and "/index.html":
// both match a GET request for "/index.html", but the former pattern
// matches all other GET and HEAD requests, while the latter matches any
// request for "/index.html" that uses a different method.
// The patterns conflict.
//
// # Trailing-slash redirection
//
// Consider a [ServeMux] with a handler for a subtree, registered using a trailing slash or "..." wildcard.
// If the ServeMux receives a request for the subtree root without a trailing slash,
// it redirects the request by adding the trailing slash.
// This behavior can be overridden with a separate registration for the path without
// the trailing slash or "..." wildcard. For example, registering "/images/" causes ServeMux
// to redirect a request for "/images" to "/images/", unless "/images" has
// been registered separately.
//
// # Request sanitizing
//
// ServeMux also takes care of sanitizing the URL request path and the Host
// header, stripping the port number and redirecting any request containing . or
// .. segments or repeated slashes to an equivalent, cleaner URL.
// Escaped path elements such as "%2e" for "." and "%2f" for "/" are preserved
// and aren't considered separators for request routing.
//
// # Compatibility
//
// The pattern syntax and matching behavior of ServeMux changed significantly
// in Go 1.22. To restore the old behavior, set the GODEBUG environment variable
// to "httpmuxgo121=1". This setting is read once, at program startup; changes
// during execution will be ignored.
//
// The backwards-incompatible changes include:
//   - Wildcards are just ordinary literal path segments in 1.21.
//     For example, the pattern "/{x}" will match only that path in 1.21,
//     but will match any one-segment path in 1.22.
//   - In 1.21, no pattern was rejected, unless it was empty or conflicted with an existing pattern.
//     In 1.22, syntactically invalid patterns will cause [ServeMux.Handle] and [ServeMux.HandleFunc] to panic.
//     For example, in 1.21, the patterns "/{"  and "/a{x}" match themselves,
//     but in 1.22 they are invalid and will cause a panic when registered.
//   - In 1.22, each segment of a pattern is unescaped; this was not done in 1.21.
//     For example, in 1.22 the pattern "/%61" matches the path "/a" ("%61" being the URL escape sequence for "a"),
//     but in 1.21 it would match only the path "/%2561" (where "%25" is the escape for the percent sign).
//   - When matching patterns to paths, in 1.22 each segment of the path is unescaped; in 1.21, the entire path is unescaped.
//     This change mostly affects how paths with %2F escapes adjacent to slashes are treated.
//     See https://go.dev/issue/21955 for details.
type ServeMux struct {
	mu     sync.RWMutex
	tree   routingNode
	index  routingIndex
	mux121 serveMux121
}

// NewServeMux allocates and returns a new [ServeMux].
func NewServeMux() *ServeMux

// DefaultServeMux is the default [ServeMux] used by [Serve].
var DefaultServeMux = &defaultServeMux

// Handler returns the handler to use for the given request,
// consulting r.Method, r.Host, and r.URL.Path. It always returns
// a non-nil handler. If the path is not in its canonical form, the
// handler will be an internally-generated handler that redirects
// to the canonical path. If the host contains a port, it is ignored
// when matching handlers.
//
// The path and host are used unchanged for CONNECT requests.
//
// Handler also returns the registered pattern that matches the
// request or, in the case of internally-generated redirects,
// the path that will match after following the redirect.
//
// If there is no registered handler that applies to the request,
// Handler returns a “page not found” or “method not supported”
// handler and an empty pattern.
//
// Handler does not modify its argument. In particular, it does not
// populate named path wildcards, so r.PathValue will always return
// the empty string.
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

// Handle registers the handler for the given pattern.
// If the given pattern conflicts with one that is already registered
// or if the pattern is invalid, Handle panics.
//
// See [ServeMux] for details on valid patterns and conflict rules.
func (mux *ServeMux) Handle(pattern string, handler Handler)

// HandleFunc registers the handler function for the given pattern.
// If the given pattern conflicts with one that is already registered
// or if the pattern is invalid, HandleFunc panics.
//
// See [ServeMux] for details on valid patterns and conflict rules.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Handle registers the handler for the given pattern in [DefaultServeMux].
// The documentation for [ServeMux] explains how patterns are matched.
func Handle(pattern string, handler Handler)

// HandleFunc registers the handler function for the given pattern in [DefaultServeMux].
// The documentation for [ServeMux] explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Serve accepts incoming HTTP connections on the listener l,
// creating a new service goroutine for each. The service goroutines
// read requests and then call handler to reply to them.
//
// The handler is typically nil, in which case [DefaultServeMux] is used.
//
// HTTP/2 support is only enabled if the Listener returns [*tls.Conn]
// connections and they were configured with "h2" in the TLS
// Config.NextProtos.
//
// Serve always returns a non-nil error.
func Serve(l net.Listener, handler Handler) error

// ServeTLS accepts incoming HTTPS connections on the listener l,
// creating a new service goroutine for each. The service goroutines
// read requests and then call handler to reply to them.
//
// The handler is typically nil, in which case [DefaultServeMux] is used.
//
// Additionally, files containing a certificate and matching private key
// for the server must be provided. If the certificate is signed by a
// certificate authority, the certFile should be the concatenation
// of the server's certificate, any intermediates, and the CA's certificate.
//
// ServeTLS always returns a non-nil error.
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error

// A Server defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	Handler Handler

	// DisableGeneralOptionsHandler, if true, passes "OPTIONS *" requests to the Handler,
	// otherwise responds with 200 OK and Content-Length: 0.
	DisableGeneralOptionsHandler bool

	// TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
	TLSConfig *tls.Config

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If zero, the value of
	// ReadTimeout is used. If negative, or if zero and ReadTimeout
	// is zero or negative, there is no timeout.
	ReadHeaderTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If zero, the value
	// of ReadTimeout is used. If negative, or if zero and ReadTimeout
	// is zero or negative, there is no timeout.
	IdleTimeout time.Duration

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes int

	// TLSNextProto optionally specifies a function to take over
	// ownership of the provided TLS connection when an ALPN
	// protocol upgrade has occurred. The map key is the protocol
	// name negotiated. The Handler argument should be used to
	// handle HTTP requests and will initialize the Request's TLS
	// and RemoteAddr if not already set. The connection is
	// automatically closed when the function returns.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	//
	// Historically, TLSNextProto was used to disable HTTP/2 support.
	// The Server.Protocols field now provides a simpler way to do this.
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	ConnState func(net.Conn, ConnState)

	// ErrorLog specifies an optional logger for errors accepting
	// connections, unexpected behavior from handlers, and
	// underlying FileSystem errors.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog *log.Logger

	// BaseContext optionally specifies a function that returns
	// the base context for incoming requests on this server.
	// The provided Listener is the specific Listener that's
	// about to start accepting requests.
	// If BaseContext is nil, the default is context.Background().
	// If non-nil, it must return a non-nil context.
	BaseContext func(net.Listener) context.Context

	// ConnContext optionally specifies a function that modifies
	// the context used for a new connection c. The provided ctx
	// is derived from the base context and has a ServerContextKey
	// value.
	ConnContext func(ctx context.Context, c net.Conn) context.Context

	// HTTP2 configures HTTP/2 connections.
	HTTP2 *HTTP2Config

	// Protocols is the set of protocols accepted by the server.
	//
	// If Protocols includes UnencryptedHTTP2, the server will accept
	// unencrypted HTTP/2 connections. The server can serve both
	// HTTP/1 and unencrypted HTTP/2 on the same address and port.
	//
	// If Protocols is nil, the default is usually HTTP/1 and HTTP/2.
	// If TLSNextProto is non-nil and does not contain an "h2" entry,
	// the default is HTTP/1 only.
	Protocols *Protocols

	inShutdown atomic.Bool

	disableKeepAlives atomic.Bool
	nextProtoOnce     sync.Once
	nextProtoErr      error

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	onShutdown []func()

	listenerGroup sync.WaitGroup
}

// Close immediately closes all active net.Listeners and any
// connections in state [StateNew], [StateActive], or [StateIdle]. For a
// graceful shutdown, use [Server.Shutdown].
//
// Close does not attempt to close (and does not even know about)
// any hijacked connections, such as WebSockets.
//
// Close returns any error returned from closing the [Server]'s
// underlying Listener(s).
func (s *Server) Close() error

// Shutdown gracefully shuts down the server without interrupting any
// active connections. Shutdown works by first closing all open
// listeners, then closing all idle connections, and then waiting
// indefinitely for connections to return to idle and then shut down.
// If the provided context expires before the shutdown is complete,
// Shutdown returns the context's error, otherwise it returns any
// error returned from closing the [Server]'s underlying Listener(s).
//
// When Shutdown is called, [Serve], [ServeTLS], [ListenAndServe], and
// [ListenAndServeTLS] immediately return [ErrServerClosed]. Make sure the
// program doesn't exit and waits instead for Shutdown to return.
//
// Shutdown does not attempt to close nor wait for hijacked
// connections such as WebSockets. The caller of Shutdown should
// separately notify such long-lived connections of shutdown and wait
// for them to close, if desired. See [Server.RegisterOnShutdown] for a way to
// register shutdown notification functions.
//
// Once Shutdown has been called on a server, it may not be reused;
// future calls to methods such as Serve will return ErrServerClosed.
func (s *Server) Shutdown(ctx context.Context) error

// RegisterOnShutdown registers a function to call on [Server.Shutdown].
// This can be used to gracefully shutdown connections that have
// undergone ALPN protocol upgrade or that have been hijacked.
// This function should start protocol-specific graceful shutdown,
// but should not wait for shutdown to complete.
func (s *Server) RegisterOnShutdown(f func())

// A ConnState represents the state of a client connection to a server.
// It's used by the optional [Server.ConnState] hook.
type ConnState int

const (
	// StateNew represents a new connection that is expected to
	// send a request immediately. Connections begin at this
	// state and then transition to either StateActive or
	// StateClosed.
	StateNew ConnState = iota

	// StateActive represents a connection that has read 1 or more
	// bytes of a request. The Server.ConnState hook for
	// StateActive fires before the request has entered a handler
	// and doesn't fire again until the request has been
	// handled. After the request is handled, the state
	// transitions to StateClosed, StateHijacked, or StateIdle.
	// For HTTP/2, StateActive fires on the transition from zero
	// to one active request, and only transitions away once all
	// active requests are complete. That means that ConnState
	// cannot be used to do per-request work; ConnState only notes
	// the overall state of the connection.
	StateActive

	// StateIdle represents a connection that has finished
	// handling a request and is in the keep-alive state, waiting
	// for a new request. Connections transition from StateIdle
	// to either StateActive or StateClosed.
	StateIdle

	// StateHijacked represents a hijacked connection.
	// This is a terminal state. It does not transition to StateClosed.
	StateHijacked

	// StateClosed represents a closed connection.
	// This is a terminal state. Hijacked connections do not
	// transition to StateClosed.
	StateClosed
)

func (c ConnState) String() string

// AllowQuerySemicolons returns a handler that serves requests by converting any
// unescaped semicolons in the URL query to ampersands, and invoking the handler h.
//
// This restores the pre-Go 1.17 behavior of splitting query parameters on both
// semicolons and ampersands. (See golang.org/issue/25192). Note that this
// behavior doesn't match that of many proxies, and the mismatch can lead to
// security issues.
//
// AllowQuerySemicolons should be invoked before [Request.ParseForm] is called.
func AllowQuerySemicolons(h Handler) Handler

// ListenAndServe listens on the TCP network address s.Addr and then
// calls [Serve] to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// If s.Addr is blank, ":http" is used.
//
// ListenAndServe always returns a non-nil error. After [Server.Shutdown] or [Server.Close],
// the returned error is [ErrServerClosed].
func (s *Server) ListenAndServe() error

// ErrServerClosed is returned by the [Server.Serve], [ServeTLS], [ListenAndServe],
// and [ListenAndServeTLS] methods after a call to [Server.Shutdown] or [Server.Close].
var ErrServerClosed = errors.New("http: Server closed")

// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each. The service goroutines read requests and
// then call s.Handler to reply to them.
//
// HTTP/2 support is only enabled if the Listener returns [*tls.Conn]
// connections and they were configured with "h2" in the TLS
// Config.NextProtos.
//
// Serve always returns a non-nil error and closes l.
// After [Server.Shutdown] or [Server.Close], the returned error is [ErrServerClosed].
func (s *Server) Serve(l net.Listener) error

// ServeTLS accepts incoming connections on the Listener l, creating a
// new service goroutine for each. The service goroutines perform TLS
// setup and then read requests, calling s.Handler to reply to them.
//
// Files containing a certificate and matching private key for the
// server must be provided if neither the [Server]'s
// TLSConfig.Certificates, TLSConfig.GetCertificate nor
// config.GetConfigForClient are populated.
// If the certificate is signed by a certificate authority, the
// certFile should be the concatenation of the server's certificate,
// any intermediates, and the CA's certificate.
//
// ServeTLS always returns a non-nil error. After [Server.Shutdown] or [Server.Close], the
// returned error is [ErrServerClosed].
func (s *Server) ServeTLS(l net.Listener, certFile, keyFile string) error

// SetKeepAlivesEnabled controls whether HTTP keep-alives are enabled.
// By default, keep-alives are always enabled. Only very
// resource-constrained environments or servers in the process of
// shutting down should disable them.
func (s *Server) SetKeepAlivesEnabled(v bool)

// ListenAndServe listens on the TCP network address addr and then calls
// [Serve] with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case [DefaultServeMux] is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error

// ListenAndServeTLS acts identically to [ListenAndServe], except that it
// expects HTTPS connections. Additionally, files containing a certificate and
// matching private key for the server must be provided. If the certificate
// is signed by a certificate authority, the certFile should be the concatenation
// of the server's certificate, any intermediates, and the CA's certificate.
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error

// ListenAndServeTLS listens on the TCP network address s.Addr and
// then calls [ServeTLS] to handle requests on incoming TLS connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// Filenames containing a certificate and matching private key for the
// server must be provided if neither the [Server]'s TLSConfig.Certificates
// nor TLSConfig.GetCertificate are populated. If the certificate is
// signed by a certificate authority, the certFile should be the
// concatenation of the server's certificate, any intermediates, and
// the CA's certificate.
//
// If s.Addr is blank, ":https" is used.
//
// ListenAndServeTLS always returns a non-nil error. After [Server.Shutdown] or
// [Server.Close], the returned error is [ErrServerClosed].
func (s *Server) ListenAndServeTLS(certFile, keyFile string) error

// TimeoutHandler returns a [Handler] that runs h with the given time limit.
//
// The new Handler calls h.ServeHTTP to handle each request, but if a
// call runs for longer than its time limit, the handler responds with
// a 503 Service Unavailable error and the given message in its body.
// (If msg is empty, a suitable default message will be sent.)
// After such a timeout, writes by h to its [ResponseWriter] will return
// [ErrHandlerTimeout].
//
// TimeoutHandler supports the [Pusher] interface but does not support
// the [Hijacker] or [Flusher] interfaces.
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

// ErrHandlerTimeout is returned on [ResponseWriter] Write calls
// in handlers which have timed out.
var ErrHandlerTimeout = errors.New("http: Handler timeout")

var _ Pusher = (*timeoutWriter)(nil)

// MaxBytesHandler returns a [Handler] that runs h with its [ResponseWriter] and [Request.Body] wrapped by a MaxBytesReader.
func MaxBytesHandler(h Handler, n int64) Handler
