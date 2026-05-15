// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transport code.

package http2

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"

	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"

	"golang.org/x/net/http2/hpack"
)

// Transport is an HTTP/2 Transport.
//
// A Transport internally caches connections to servers. It is safe
// for concurrent use by multiple goroutines.
type Transport struct {
	t1       TransportConfig
	connPool noDialClientConnPool
	*transportTestHooks
}

func NewTransport(t1 TransportConfig) *Transport

func (t *Transport) AddConn(scheme, authority string, c net.Conn) error

// ClientConn is the state of a single HTTP/2 client connection to an
// HTTP/2 server.
type ClientConn struct {
	t             *Transport
	tconn         net.Conn
	tlsState      *tls.ConnectionState
	atomicReused  uint32
	singleUse     bool
	getConnCalled bool

	// readLoop goroutine fields:
	readerDone chan struct{}
	readerErr  error

	idleTimeout time.Duration
	idleTimer   *time.Timer

	mu               sync.Mutex
	cond             *sync.Cond
	flow             outflow
	inflow           inflow
	doNotReuse       bool
	closing          bool
	closed           bool
	closedOnIdle     bool
	seenSettings     bool
	seenSettingsChan chan struct{}
	wantSettingsAck  bool
	goAway           *GoAwayFrame
	goAwayDebug      string
	streams          map[uint32]*clientStream
	streamsReserved  int
	nextStreamID     uint32
	pendingRequests  int
	pings            map[[8]byte]chan struct{}
	br               *bufio.Reader
	lastActive       time.Time
	lastIdle         time.Time
	// Settings from peer: (also guarded by wmu)
	maxFrameSize                uint32
	maxConcurrentStreams        uint32
	peerMaxHeaderListSize       uint64
	peerMaxHeaderTableSize      uint32
	initialWindowSize           uint32
	initialStreamRecvWindowSize int32
	readIdleTimeout             time.Duration
	pingTimeout                 time.Duration
	extendedConnectAllowed      bool
	strictMaxConcurrentStreams  bool

	// rstStreamPingsBlocked works around an unfortunate gRPC behavior.
	// gRPC strictly limits the number of PING frames that it will receive.
	// The default is two pings per two hours, but the limit resets every time
	// the gRPC endpoint sends a HEADERS or DATA frame. See golang/go#70575.
	//
	// rstStreamPingsBlocked is set after receiving a response to a PING frame
	// bundled with an RST_STREAM (see pendingResets below), and cleared after
	// receiving a HEADERS or DATA frame.
	rstStreamPingsBlocked bool

	// pendingResets is the number of RST_STREAM frames we have sent to the peer,
	// without confirming that the peer has received them. When we send a RST_STREAM,
	// we bundle it with a PING frame, unless a PING is already in flight. We count
	// the reset stream against the connection's concurrency limit until we get
	// a PING response. This limits the number of requests we'll try to send to a
	// completely unresponsive connection.
	pendingResets int

	// readBeforeStreamID is the smallest stream ID that has not been followed by
	// a frame read from the peer. We use this to determine when a request may
	// have been sent to a completely unresponsive connection:
	// If the request ID is less than readBeforeStreamID, then we have had some
	// indication of life on the connection since sending the request.
	readBeforeStreamID uint32

	// reqHeaderMu is a 1-element semaphore channel controlling access to sending new requests.
	// Write to reqHeaderMu to lock it, read from it to unlock.
	// Lock reqmu BEFORE mu or wmu.
	reqHeaderMu chan struct{}

	// internalStateHook reports state changes back to the net/http.ClientConn.
	// Note that this is different from the user state hook registered by
	// net/http.ClientConn.SetStateHook: The internal hook calls ClientConn,
	// which calls the user hook.
	internalStateHook func()

	// wmu is held while writing.
	// Acquire BEFORE mu when holding both, to avoid blocking mu on network writes.
	// Only acquire both at the same time when changing peer settings.
	wmu  sync.Mutex
	bw   *bufio.Writer
	fr   *Framer
	werr error
	hbuf bytes.Buffer
	henc *hpack.Encoder
}

var ErrNoCachedConn error = noCachedConnError{}

// RoundTripOpt are options for the Transport.RoundTripOpt method.
type RoundTripOpt struct {
	// OnlyCachedConn controls whether RoundTripOpt may
	// create a new TCP connection. If set true and
	// no cached connection is available, RoundTripOpt
	// will return ErrNoCachedConn.
	OnlyCachedConn bool
}

func (t *Transport) RoundTrip(req *ClientRequest) (*ClientResponse, error)

// RoundTripOpt is like RoundTrip, but takes options.
func (t *Transport) RoundTripOpt(req *ClientRequest, opt RoundTripOpt) (*ClientResponse, error)

func (t *Transport) IdleConnStrsForTesting() []string

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle.
// It does not interrupt any connections currently in use.
func (t *Transport) CloseIdleConnections()

func (t *Transport) NewClientConn(c net.Conn, internalStateHook func()) (NetHTTPClientConn, error)

// SetDoNotReuse marks cc as not reusable for future HTTP requests.
func (cc *ClientConn) SetDoNotReuse()

// CanTakeNewRequest reports whether the connection can take a new request,
// meaning it has not been closed or received or sent a GOAWAY.
//
// If the caller is going to immediately make a new request on this
// connection, use ReserveNewRequest instead.
func (cc *ClientConn) CanTakeNewRequest() bool

// ReserveNewRequest is like CanTakeNewRequest but also reserves a
// concurrent stream in cc. The reservation is decremented on the
// next call to RoundTrip.
func (cc *ClientConn) ReserveNewRequest() bool

// ClientConnState describes the state of a ClientConn.
type ClientConnState struct {
	// Closed is whether the connection is closed.
	Closed bool

	// Closing is whether the connection is in the process of
	// closing. It may be closing due to shutdown, being a
	// single-use connection, being marked as DoNotReuse, or
	// having received a GOAWAY frame.
	Closing bool

	// StreamsActive is how many streams are active.
	StreamsActive int

	// StreamsReserved is how many streams have been reserved via
	// ClientConn.ReserveNewRequest.
	StreamsReserved int

	// StreamsPending is how many requests have been sent in excess
	// of the peer's advertised MaxConcurrentStreams setting and
	// are waiting for other streams to complete.
	StreamsPending int

	// MaxConcurrentStreams is how many concurrent streams the
	// peer advertised as acceptable. Zero means no SETTINGS
	// frame has been received yet.
	MaxConcurrentStreams uint32

	// LastIdle, if non-zero, is when the connection last
	// transitioned to idle state.
	LastIdle time.Time
}

// State returns a snapshot of cc's state.
func (cc *ClientConn) State() ClientConnState

// Shutdown gracefully closes the client connection, waiting for running streams to complete.
func (cc *ClientConn) Shutdown(ctx context.Context) error

// Close closes the client connection immediately.
//
// In-flight requests are interrupted. For a graceful shutdown, use Shutdown instead.
func (cc *ClientConn) Close() error

func (cc *ClientConn) RoundTrip(req *ClientRequest) (*ClientResponse, error)

// GoAwayError is returned by the Transport when the server closes the
// TCP connection after sending a GOAWAY frame.
type GoAwayError struct {
	LastStreamID uint32
	ErrCode      ErrCode
	DebugData    string
}

func (e GoAwayError) Error() string

// Ping sends a PING frame to the server and waits for the ack.
func (cc *ClientConn) Ping(ctx context.Context) error

// netHTTPClientConn wraps ClientConn and implements the interface net/http expects from
// the RoundTripper returned by NewClientConn.
type NetHTTPClientConn struct {
	cc *ClientConn
}

func (cc NetHTTPClientConn) RoundTrip(req *ClientRequest) (*ClientResponse, error)

func (cc NetHTTPClientConn) Close() error

func (cc NetHTTPClientConn) Err() error

func (cc NetHTTPClientConn) Reserve() error

func (cc NetHTTPClientConn) Release()

func (cc NetHTTPClientConn) Available() int

func (cc NetHTTPClientConn) InFlight() int

func (cc NetHTTPClientConn) Ping(ctx context.Context) error
