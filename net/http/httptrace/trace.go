// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package httptrace provides mechanisms to trace the events within
// HTTP client requests.
package httptrace

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/textproto"
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

// ClientTrace is a set of hooks to run at various stages of an outgoing
// HTTP request. Any particular hook may be nil. Functions may be
// called concurrently from different goroutines and some may be called
// after the request has completed or failed.
//
// ClientTrace currently traces a single HTTP request & response
// during a single round trip and has no hooks that span a series
// of redirected requests.
//
// See https://blog.golang.org/http-tracing for more.
type ClientTrace struct {
	GetConn func(hostPort string)

	GotConn func(GotConnInfo)

	PutIdleConn func(err error)

	GotFirstResponseByte func()

	Got100Continue func()

	Got1xxResponse func(code int, header textproto.MIMEHeader) error

	DNSStart func(DNSStartInfo)

	DNSDone func(DNSDoneInfo)

	ConnectStart func(network, addr string)

	ConnectDone func(network, addr string, err error)

	TLSHandshakeStart func()

	TLSHandshakeDone func(tls.ConnectionState, error)

	WroteHeaderField func(key string, value []string)

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
