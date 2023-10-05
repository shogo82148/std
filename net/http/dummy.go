package http

import (
	"errors"

	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
)

const defaultUserAgent = "Go-http-client/1.1"

var newLoggingConn any
var appendTime any
var refererForURL any

func (s *Server) newConn(c net.Conn)

type conn struct{}

var errRequestCanceled = errors.New("net/http: request canceled")
var errRequestCanceledConn = errors.New("net/http: request canceled while waiting for connection")
var errServerClosedIdle = errors.New("net/http: server closed idle connection")

func serveFile(w ResponseWriter, r *Request, fsys FileSystem, name string)
func scanETag(etag string) (string, bool)
func http2ConfigureServer()
func shouldCopyHeaderOnRedirect(header string) bool
func writeStatusLine(w ResponseWriter, proto, status string)
func is408Message(status string) bool

var maxWriteWaitBeforeConnReuse any

type hookSetter any

var testHookEnterRoundTrip any
var testHookRoundTripRetried any

type http2clientConnPoolIdleCloser any
type http2clientConnPool struct{}
type http2noDialClientConnPool struct{}

type http2responseWriter struct{}

func (*http2responseWriter) CloseNotify() <-chan bool

func (w *http2responseWriter) Flush()

func (w *http2responseWriter) Push(target string, opts *PushOptions) error

type http2stringWriter any

type noBody struct{}

func (noBody) WriteTo(w io.Writer) (n int64, err error)

func (noBody) Read(p []byte) (n int, err error)

func (noBody) Close() error

type httpRange struct{ a, b int }

type contextKey struct{ name string }

type closeWriter any

type muxEntry struct{}

var defaultServeMux ServeMux

type timeoutWriter struct{}

func (*timeoutWriter) Push(target string, opts *PushOptions) error

func defaultTransportDialContext(*net.Dialer) func(ctx context.Context, network string, addr string) (net.Conn, error)

type connectMethodKey struct{}

type persistConn struct{}

type wantConnQueue struct{}

type connLRU struct{}

type cancelKey struct{}

type h2Transport struct{}

type persistConnWriter struct{}

func (*persistConnWriter) ReadFrom(r io.Reader) (n int64, err error)
