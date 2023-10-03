package http

import (
	"errors"

	"github.com/shogo82148/std/net"
)

const defaultUserAgent = "Go-http-client/1.1"

var newLoggingConn any
var appendTime any
var refererForURL any

func (s *Server) newConn(c net.Conn)

type connType struct{}

var conn *connType

func (connType) closeWriteAndWait()

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
