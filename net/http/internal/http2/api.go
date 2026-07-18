// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/mime/multipart"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http/internal"
	"github.com/shogo82148/std/net/textproto"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"
)

// Variables defined in net/http and initialized by an init func in that package.
//
// NoBody and LocalAddrContextKey have concrete types in net/http,
// and therefore can't be moved into a common package without introducing
// a dependency cycle.
var (
	NoBody              io.ReadCloser
	LocalAddrContextKey any
)

var (
	ErrAbortHandler    = internal.ErrAbortHandler
	ErrBodyNotAllowed  = internal.ErrBodyNotAllowed
	ErrNotSupported    = errors.ErrUnsupported
	ErrSkipAltProtocol = internal.ErrSkipAltProtocol
)

// A ClientRequest is a Request used by the HTTP/2 client (Transport).
type ClientRequest struct {
	Context       context.Context
	Method        string
	URL           *url.URL
	Header        Header
	Trailer       Header
	Body          io.ReadCloser
	Host          string
	GetBody       func() (io.ReadCloser, error)
	ContentLength int64
	Cancel        <-chan struct{}
	Close         bool
	ResTrailer    *Header

	// Include the per-request stream in the ClientRequest to avoid an allocation.
	stream clientStream
}

// Clone makes a shallow copy of ClientRequest.
//
// Clone is only used in shouldRetryRequest.
// We can drop it if we ever get rid of or rework that function.
func (req *ClientRequest) Clone() *ClientRequest

// A ClientResponse is a Request used by the HTTP/2 client (Transport).
type ClientResponse struct {
	Status        string
	StatusCode    int
	ContentLength int64
	Uncompressed  bool
	Header        Header
	Trailer       Header
	Body          io.ReadCloser
	TLS           *tls.ConnectionState
}

type Header = textproto.MIMEHeader

// TransportConfig is configuration from an http.Transport.
type TransportConfig interface {
	MaxHeaderListSize() int64
	MaxResponseHeaderBytes() int64
	DisableCompression() bool
	DisableKeepAlives() bool
	ExpectContinueTimeout() time.Duration
	ResponseHeaderTimeout() time.Duration
	IdleConnTimeout() time.Duration
	HTTP2Config() Config
}

// ServerConfig is configuration from an http.Server.
type ServerConfig interface {
	MaxHeaderBytes() int
	MaxHeaderValueCount() int
	ConnState(net.Conn, ConnState)
	DoKeepAlives() bool
	WriteTimeout() time.Duration
	SendPingTimeout() time.Duration
	ErrorLog() *log.Logger
	ReadTimeout() time.Duration
	HTTP2Config() Config
	DisableClientPriority() bool
	IdleTimeout() time.Duration
}

type Handler interface {
	ServeHTTP(*ResponseWriter, *ServerRequest)
}

type ResponseWriter = responseWriter

type PushOptions struct {
	Method string
	Header Header
}

// A ServerRequest is a Request used by the HTTP/2 server.
type ServerRequest struct {
	Context       context.Context
	Proto         string
	ProtoMajor    int
	ProtoMinor    int
	Method        string
	URL           *url.URL
	Header        Header
	Trailer       Header
	Body          io.ReadCloser
	Host          string
	ContentLength int64
	RemoteAddr    string
	RequestURI    string
	TLS           *tls.ConnectionState
	MultipartForm *multipart.Form
}

// ConnState is identical to net/http.ConnState.
type ConnState int

const (
	ConnStateNew ConnState = iota
	ConnStateActive
	ConnStateIdle
	ConnStateHijacked
	ConnStateClosed
)
