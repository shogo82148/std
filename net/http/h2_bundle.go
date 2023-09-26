//go:build !nethttpomithttp2
// +build !nethttpomithttp2

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.
//   $ bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

// Package http2 implements the HTTP/2 protocol.
//
// This package is low-level and intended to be used directly by very
// few people. Most users will use it indirectly through the automatic
// use by the net/http package (from Go 1.6 and later).
// For use in earlier Go versions see ConfigureServer. (Transport support
// requires Go 1.6 or later)
//
// See https://http2.github.io/ for more information on HTTP/2.
//
// See https://http2.golang.org/ for a test server running this code.
//

package http

import (
	mathrand "math/rand"
)

// ClientConnPool manages a pool of HTTP/2 client connections.

// clientConnPoolIdleCloser is the interface implemented by ClientConnPool
// implementations which can close their idle connections.

var (
	_ http2clientConnPoolIdleCloser = (*http2clientConnPool)(nil)
	_ http2clientConnPoolIdleCloser = http2noDialClientConnPool{}
)

// TODO: use singleflight for dialing and addConnCalls?

// dialCall is an in-flight Transport dial call to a host.

// noDialClientConnPool is an implementation of http2.ClientConnPool
// which never dials. We let the HTTP/1.1 client dial and use its TLS
// connection instead.

// Buffer chunks are allocated from a pool to reduce pressure on GC.
// The maximum wasted space per dataBuffer is 2x the largest size class,
// which happens when the dataBuffer has multiple chunks and there is
// one unread byte in both the first and last chunks. We use a few size
// classes to minimize overheads for servers that typically receive very
// small request bodies.
//
// TODO: Benchmark to determine if the pools are necessary. The GC may have
// improved enough that we can instead allocate chunks like this:
// make([]byte, max(16<<10, expectedBytesRemaining))

// dataBuffer is an io.ReadWriter backed by a list of data chunks.
// Each dataBuffer is used to read DATA frames on a single stream.
// The buffer is divided into chunks so the server can limit the
// total memory used by a single connection without limiting the
// request body size on any single stream.

// An ErrCode is an unsigned 32-bit error code as defined in the HTTP/2 spec.

// ConnectionError is an error that results in the termination of the
// entire connection.

// StreamError is an error that only affects one stream within an
// HTTP/2 connection.

// 6.9.1 The Flow Control Window
// "If a sender receives a WINDOW_UPDATE that causes a flow control
// window to exceed this maximum it MUST terminate either the stream
// or the connection, as appropriate. For streams, [...]; for the
// connection, a GOAWAY frame with a FLOW_CONTROL_ERROR code."

// connError represents an HTTP/2 ConnectionError error code, along
// with a string (for debugging) explaining why.
//
// Errors of this type are only returned by the frame parser functions
// and converted into ConnectionError(Code), after stashing away
// the Reason into the Framer's errDetail field, accessible via
// the (*Framer).ErrorDetail method.

// flow is the flow control window's size.

// A FrameType is a registered frame type as defined in
// http://http2.github.io/http2-spec/#rfc.section.11.2

// Flags is a bitmask of HTTP/2 flags.
// The meaning of flags varies depending on the frame type.

// Frame-specific FrameHeader flag bits.

// a frameParser parses a frame given its FrameHeader and payload
// bytes. The length of payload will always equal fh.Length (which
// might be 0).

// A FrameHeader is the 9 byte header of all HTTP/2 frames.
//
// See http://http2.github.io/http2-spec/#FrameHeader

// frame header bytes.
// Used only by ReadFrameHeader.

// A Frame is the base interface implemented by all frame types.
// Callers will generally type-assert the specific frame type:
// *HeadersFrame, *SettingsFrame, *WindowUpdateFrame, etc.
//
// Frames are only valid until the next call to Framer.ReadFrame.

// A Framer reads and writes Frames.

// ErrFrameTooLarge is returned from Framer.ReadFrame when the peer
// sends a frame that is larger than declared with SetMaxReadFrameSize.

// A DataFrame conveys arbitrary, variable-length sequences of octets
// associated with a stream.
// See http://http2.github.io/http2-spec/#rfc.section.6.1

// A SettingsFrame conveys configuration parameters that affect how
// endpoints communicate, such as preferences and constraints on peer
// behavior.
//
// See http://http2.github.io/http2-spec/#SETTINGS

// A PingFrame is a mechanism for measuring a minimal round trip time
// from the sender, as well as determining whether an idle connection
// is still functional.
// See http://http2.github.io/http2-spec/#rfc.section.6.7

// A GoAwayFrame informs the remote peer to stop creating streams on this connection.
// See http://http2.github.io/http2-spec/#rfc.section.6.8

// An UnknownFrame is the frame type returned when the frame type is unknown
// or no specific frame type parser exists.

// A WindowUpdateFrame is used to implement flow control.
// See http://http2.github.io/http2-spec/#rfc.section.6.9

// A HeadersFrame is used to open a stream and additionally carries a
// header block fragment.

// HeadersFrameParam are the parameters for writing a HEADERS frame.

// A PriorityFrame specifies the sender-advised priority of a stream.
// See http://http2.github.io/http2-spec/#rfc.section.6.3

// PriorityParam are the stream prioritzation parameters.

// A RSTStreamFrame allows for abnormal termination of a stream.
// See http://http2.github.io/http2-spec/#rfc.section.6.4

// A ContinuationFrame is used to continue a sequence of header block fragments.
// See http://http2.github.io/http2-spec/#rfc.section.6.10

// A PushPromiseFrame is used to initiate a server stream.
// See http://http2.github.io/http2-spec/#rfc.section.6.6

// PushPromiseParam are the parameters for writing a PUSH_PROMISE frame.

// A MetaHeadersFrame is the representation of one HEADERS frame and
// zero or more contiguous CONTINUATION frames and the decoding of
// their HPACK-encoded contents.
//
// This type of frame does not appear on the wire and is only returned
// by the Framer when Framer.ReadMetaHeaders is set.

// HTTP/2 stream states.
//
// See http://tools.ietf.org/html/rfc7540#section-5.1.
//
// For simplicity, the server code merges "reserved (local)" into
// "half-closed (remote)". This is one less state transition to track.
// The only downside is that we send PUSH_PROMISEs slightly less
// liberally than allowable. More discussion here:
// https://lists.w3.org/Archives/Public/ietf-http-wg/2016JulSep/0599.html
//
// "reserved (remote)" is omitted since the client code does not
// support server push.

// Setting is a setting parameter: which setting it is, and its value.

// A SettingID is an HTTP/2 setting as defined in
// http://http2.github.io/http2-spec/#iana-settings

// from pkg io

// A gate lets two goroutines coordinate their activities.

// A closeWaiter is like a sync.WaitGroup but only goes 1 to 0 (open to closed).

// bufferedWriter is a buffered writer that writes to w.
// Its buffered writer is lazily allocated as needed, to minimize
// idle memory usage with many connections.

// bufWriterPoolBufferSize is the size of bufio.Writer's
// buffers created using bufWriterPool.
//
// TODO: pick a less arbitrary value? this is a bit under
// (3 x typical 1500 byte MTU) at least. Other than that,
// not much thought went into it.

// incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).

// pipe is a goroutine-safe io.Reader/io.Writer pair. It's like
// io.Pipe except there are no PipeReader/PipeWriter halves, and the
// underlying buffer is an interface. (io.Pipe is always unbuffered)

// Test hooks.

// Server is an HTTP/2 server.

// ServeConnOpts are options for the Server.ServeConn method.

// stream represents a stream. This is the minimal metadata needed by
// the serve goroutine. Most of the actual stream state is owned by
// the http.Handler's goroutine in the responseWriter. Because the
// responseWriter's responseWriterState is recycled at the end of a
// handler, this struct intentionally has no pointer to the
// *responseWriter{,State} itself, as the Handler ending nils out the
// responseWriter's state field.

// frameWriteResult is the message passed from writeFrameAsync to the serve goroutine.

// Message values sent to serveMsgCh.

// errHandlerPanicked is the error given to any callers blocked in a read from
// Request.Body when the main goroutine panics. Since most handlers read in the
// main ServeHTTP goroutine, this will show up rarely.

// After sending GOAWAY, the connection will close after goAwayTimeout.
// If we close the connection immediately after sending GOAWAY, there may
// be unsent data in our kernel receive buffer, which will cause the kernel
// to send a TCP RST on close() instead of a FIN. This RST will abort the
// connection immediately, whether or not the client had received the GOAWAY.
//
// Ideally we should delay for at least 1 RTT + epsilon so the client has
// a chance to read the GOAWAY and stop sending messages. Measuring RTT
// is hard, so we approximate with 1 second. See golang.org/issue/18701.
//
// This is a var so it can be shorter in tests, where all requests uses the
// loopback interface making the expected RTT very small.
//
// TODO: configurable?

// A bodyReadMsg tells the server loop that the http.Handler read n
// bytes of the DATA from the client on the given stream.

// requestBody is the Handler's Request.Body type.
// Read and Close may be called concurrently.

// responseWriter is the http.ResponseWriter implementation. It's
// intentionally small (1 pointer wide) to minimize garbage. The
// responseWriterState pointer inside is zeroed at the end of a
// request (in handlerDone) and calls on the responseWriter thereafter
// simply crash (caller's mistake), but the much larger responseWriterState
// and buffers are reused between multiple requests.

// Optional http.ResponseWriter interfaces implemented.
var (
	_ CloseNotifier     = (*http2responseWriter)(nil)
	_ Flusher           = (*http2responseWriter)(nil)
	_ http2stringWriter = (*http2responseWriter)(nil)
)

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
//    https://golang.org/pkg/net/http/#ResponseWriter
//    https://golang.org/pkg/net/http/#example_ResponseWriter_trailers

// Push errors.

var _ Pusher = (*http2responseWriter)(nil)

// From http://httpwg.org/specs/rfc7540.html#rfc.section.8.1.2.2

// Transport is an HTTP/2 Transport.
//
// A Transport internally caches connections to servers. It is safe
// for concurrent use by multiple goroutines.

// ClientConn is the state of a single HTTP/2 client connection to an
// HTTP/2 server.

// clientStream is the state for a single HTTP/2 stream. One of these
// is created for each Transport.RoundTrip call.

// noCachedConnError is the concrete type of ErrNoCachedConn, which
// needs to be detected by net/http regardless of whether it's its
// bundled version (in h2_bundle.go with a rewritten type name) or
// from a user's x/net/http2. As such, as it has a unique method name
// (IsHTTP2NoCachedConnError) that net/http sniffs for via func
// isNoCachedConnError.

// RoundTripOpt are options for the Transport.RoundTripOpt method.

// clientConnIdleState describes the suitability of a client
// connection to initiate a new RoundTrip request.

// errRequestCanceled is a copy of net/http's errRequestCanceled because it's not
// exported. At least they'll be DeepEqual for h1-vs-h2 comparisons tests.

// internal error values; they don't escape to callers

// clientConnReadLoop is the state owned by the clientConn's frame-reading readLoop.

// GoAwayError is returned by the Transport when the server closes the
// TCP connection after sending a GOAWAY frame.

// transportResponseBody is the concrete type of Transport.RoundTrip's
// Response.Body. It is an io.ReadCloser. On Read, it reads from cs.body.
// On Close it sends RST_STREAM if EOF wasn't already seen.

// gzipReader wraps a response body so it can lazily
// call gzip.NewReader on the first call to Read

// bodyWriterState encapsulates various state around the Transport's writing
// of the request body, particularly regarding doing delayed writes of the body
// when the request contains "Expect: 100-continue".

// noDialH2RoundTripper is a RoundTripper which only tries to complete the request
// if there's already has a cached connection to the host.
// (The field is exported so it can be accessed via reflect from net/http; tested
// by TestNoDialH2RoundTripperType)

// writeFramer is implemented by any type that is used to write frames.

// writeContext is the interface needed by the various frame writer
// types below. All the writeFrame methods below are scheduled via the
// frame writing scheduler (see writeScheduler in writesched.go).
//
// This interface is implemented by *serverConn.
//
// TODO: decide whether to a) use this in the client code (which didn't
// end up using this yet, because it has a simpler design, not
// currently implementing priorities), or b) delete this and
// make the server code a bit more concrete.

// handlerPanicRST is the message sent from handler goroutines when
// the handler panics.

// writeResHeaders is a request to write a HEADERS and 0+ CONTINUATION frames
// for HTTP response headers or trailers from a server handler.

// writePushPromise is a request to write a PUSH_PROMISE and 0+ CONTINUATION frames.

// WriteScheduler is the interface implemented by HTTP/2 write schedulers.
// Methods are never called concurrently.

// OpenStreamOptions specifies extra options for WriteScheduler.OpenStream.

// FrameWriteRequest is a request to write a frame.

// writeQueue is used by implementations of WriteScheduler.

// RFC 7540, Section 5.3.5: the default weight is 16.

// PriorityWriteSchedulerConfig configures a priorityWriteScheduler.

// priorityNode is a node in an HTTP/2 priority tree.
// Each node is associated with a single stream ID.
// See RFC 7540, Section 5.3.
