// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.h

// Package httptrace provides mechanisms to trace the events within
// HTTP client requests.
package httptrace

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/time"
)

// unique type to prevent assignment.

// ContextClientTrace returns the ClientTrace associated with the
// provided context. If none, it returns nil.
func ContextClientTrace(ctx context.Context) *ClientTrace

// WithClientTrace returns a new context based on the provided parent
// ctx. HTTP client requests made with the returned context will use
// the provided trace hooks, in addition to any previous hooks
// registered with ctx. Any hooks defined in the provided trace will
// be called first.
func WithClientTrace(ctx context.Context, trace *ClientTrace) context.Context

// ClientTrace is a set of hooks to run at various stages of an HTTP
// client request. Any particular hook may be nil. Functions may be
// called concurrently from different goroutines, starting after the
// call to Transport.RoundTrip and ending either when RoundTrip
// returns an error, or when the Response.Body is closed.
type ClientTrace struct {
	GetConn func(hostPort string)

	GotConn func(GotConnInfo)

	PutIdleConn func(err error)

	GotFirstResponseByte func()

	Got100Continue func()

	DNSStart func(DNSStartInfo)

	DNSDone func(DNSDoneInfo)

	ConnectStart func(network, addr string)

	ConnectDone func(network, addr string, err error)

	WroteHeaders func()

	Wait100Continue func()

	WroteRequest func(WroteRequestInfo)
}

// WroteRequestInfo contains information provided to the WroteRequest
// hook.
type WroteRequestInfo struct {
	Err error
}

// DNSStartInfo contains information about a DNS request.
type DNSStartInfo struct {
	Host string
}

// DNSDoneInfo contains information about the results of a DNS lookup.
type DNSDoneInfo struct {
	Addrs []net.IPAddr

	Err error

	Coalesced bool
}

// GotConnInfo is the argument to the ClientTrace.GotConn function and
// contains information about the obtained connection.
type GotConnInfo struct {
	Conn net.Conn

	Reused bool

	WasIdle bool

	IdleTime time.Duration
}
