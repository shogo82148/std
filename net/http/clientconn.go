// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/sync"
)

// A ClientConn is a client connection to an HTTP server.
//
// Unlike a [Transport], a ClientConn represents a single connection.
// Most users should use a Transport rather than creating client connections directly.
type ClientConn struct {
	cc genericClientConn

	stateHookMu      sync.Mutex
	userStateHook    func(*ClientConn)
	stateHookRunning bool
	lastAvailable    int
	lastInFlight     int
	lastClosed       bool
}

// NewClientConn creates a new client connection to the given address.
//
// If scheme is "http", the connection is unencrypted.
// If scheme is "https", the connection uses TLS.
//
// The protocol used for the new connection is determined by the scheme,
// Transport.Protocols configuration field, and protocols supported by the
// server. See Transport.Protocols for more details.
//
// If Transport.Proxy is set and indicates that a request sent to the given
// address should use a proxy, the new connection uses that proxy.
//
// NewClientConn always creates a new connection,
// even if the Transport has an existing cached connection to the given host.
//
// The new connection is not added to the Transport's connection cache,
// and will not be used by [Transport.RoundTrip].
// It does not count against the MaxIdleConns and MaxConnsPerHost limits.
//
// The caller is responsible for closing the new connection.
func (t *Transport) NewClientConn(ctx context.Context, scheme, address string) (*ClientConn, error)

// Close closes the connection.
// Outstanding RoundTrip calls are interrupted.
func (cc *ClientConn) Close() error

// Err reports any fatal connection errors.
// It returns nil if the connection is usable.
// If it returns non-nil, the connection can no longer be used.
func (cc *ClientConn) Err() error

// RoundTrip implements the [RoundTripper] interface.
//
// The request is sent on the client connection,
// regardless of the URL being requested or any proxy settings.
//
// If the connection is at its concurrency limit,
// RoundTrip waits for the connection to become available
// before sending the request.
func (cc *ClientConn) RoundTrip(req *Request) (*Response, error)

// Available reports the number of requests that may be sent
// to the connection without blocking.
// It returns 0 if the connection is closed.
func (cc *ClientConn) Available() int

// InFlight reports the number of requests in flight,
// including reserved requests.
// It returns 0 if the connection is closed.
func (cc *ClientConn) InFlight() int

// Reserve reserves a concurrency slot on the connection.
// If Reserve returns nil, one additional RoundTrip call may be made
// without waiting for an existing request to complete.
//
// The reserved concurrency slot is accounted as an in-flight request.
// A successful call to RoundTrip will decrement the Available count
// and increment the InFlight count.
//
// Each successful call to Reserve should be followed by exactly one call
// to RoundTrip or Release, which will consume or release the reservation.
//
// If the connection is closed or at its concurrency limit,
// Reserve returns an error.
func (cc *ClientConn) Reserve() error

// Release releases an unused concurrency slot reserved by Reserve.
// If there are no reserved concurrency slots, it has no effect.
func (cc *ClientConn) Release()

// SetStateHook arranges for f to be called when the state of the connection changes.
// At most one call to f is made at a time.
// If the connection's state has changed since it was created,
// f is called immediately in a separate goroutine.
// f may be called synchronously from RoundTrip or Response.Body.Close.
//
// If SetStateHook is called multiple times, the new hook replaces the old one.
// If f is nil, no further calls will be made to f after SetStateHook returns.
//
// f is called when Available increases (more requests may be sent on the connection),
// InFlight decreases (existing requests complete), or Err begins returning non-nil
// (the connection is no longer usable).
func (cc *ClientConn) SetStateHook(f func(*ClientConn))
