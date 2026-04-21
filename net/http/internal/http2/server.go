// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: turn off the serve goroutine when idle, so
// an idle conn only has the readFrames goroutine active. (which could
// also be optimized probably to pin less memory in crypto/tls). This
// would involve tracking when the serve goroutine is active (atomic
// int32 read/CAS probably?) and starting it up when frames arrive,
// and shutting it down when all handlers exit. the occasional PING
// packets could use time.AfterFunc to call sc.wakeStartServeLoop()
// (which is a no-op if already running) and then queue the PING write
// as normal. The serve loop would then exit in most cases (if no
// Handlers running) and not be woken up again until the PING packet
// returns.

// TODO (maybe): add a mechanism for Handlers to going into
// half-closed-local mode (rw.(io.Closer) test?) but not exit their
// handler, and continue to be able to read from the
// Request.Body. This would be a somewhat semantic change from HTTP/1
// (or at least what we expose in net/http), so I'd probably want to
// add it there too. For now, this package says that returning from
// the Handler ServeHTTP function means you're both done reading and
// done writing, without a way to stop just one or the other.

package http2

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
)

// Server is an HTTP/2 server.
type Server struct {
	mu          sync.Mutex
	activeConns map[*serverConn]struct{}

	// Pool of error channels. This is per-Server rather than global
	// because channels can't be reused across synctest bubbles.
	errChanPool sync.Pool
}

func (s *Server) Configure(conf ServerConfig, tcfg *tls.Config) error

func (s *Server) GracefulShutdown()

// ServeConnOpts are options for the Server.ServeConn method.
type ServeConnOpts struct {
	// Context is the base context to use.
	// If nil, context.Background is used.
	Context context.Context

	// BaseConfig optionally sets the base configuration
	// for values. If nil, defaults are used.
	BaseConfig ServerConfig

	// Handler specifies which handler to use for processing
	// requests. If nil, BaseConfig.Handler is used. If BaseConfig
	// or BaseConfig.Handler is nil, http.DefaultServeMux is used.
	Handler Handler

	// Settings is the decoded contents of the HTTP2-Settings header
	// in an h2c upgrade request.
	Settings []byte

	UpgradeRequest *ServerRequest

	// SawClientPreface is set if the HTTP/2 connection preface
	// has already been read from the connection.
	SawClientPreface bool
}

// ServeConn serves HTTP/2 requests on the provided connection and
// blocks until the connection is no longer readable.
//
// ServeConn starts speaking HTTP/2 assuming that c has not had any
// reads or writes. It writes its initial settings frame and expects
// to be able to read the preface and settings frame from the
// client. If c has a ConnectionState method like a *tls.Conn, the
// ConnectionState is used to verify the TLS ciphersuite and to set
// the Request.TLS field in Handlers.
//
// ServeConn does not support h2c by itself. Any h2c support must be
// implemented in terms of providing a suitably-behaving net.Conn.
//
// The opts parameter is optional. If nil, default values are used.
func (s *Server) ServeConn(c net.Conn, opts *ServeConnOpts)

var (
	NewConnContextKey         = new("NewConnContextKey")
	ConnectionStateContextKey = new("ConnectionStateContextKey")
)

const DefaultMaxHeaderBytes = 1 << 20

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

// TrailerPrefix is a magic prefix for ResponseWriter.Header map keys
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
//	https://golang.org/pkg/net/http/#ResponseWriter
//	https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
const TrailerPrefix = "Trailer:"

// Push errors.
var (
	ErrRecursivePush    = errors.New("http2: recursive push not allowed")
	ErrPushLimitReached = errors.New("http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS")
)
