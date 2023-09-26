// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP reverse proxy handler

package httputil

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"
)

// ReverseProxy is an HTTP Handler that takes an incoming request and
// sends it to another server, proxying the response back to the
// client.
//
// ReverseProxy by default sets the client IP as the value of the
// X-Forwarded-For header.
//
// If an X-Forwarded-For header already exists, the client IP is
// appended to the existing values. As a special case, if the header
// exists in the Request.Header map but has a nil value (such as when
// set by the Director func), the X-Forwarded-For header is
// not modified.
//
// To prevent IP spoofing, be sure to delete any pre-existing
// X-Forwarded-For header coming from the client or
// an untrusted proxy.
type ReverseProxy struct {
	Director func(*http.Request)

	Transport http.RoundTripper

	FlushInterval time.Duration

	ErrorLog *log.Logger

	BufferPool BufferPool

	ModifyResponse func(*http.Response) error

	ErrorHandler func(http.ResponseWriter, *http.Request, error)
}

// A BufferPool is an interface for getting and returning temporary
// byte slices for use by io.CopyBuffer.
type BufferPool interface {
	Get() []byte
	Put([]byte)
}

// NewSingleHostReverseProxy returns a new ReverseProxy that routes
// URLs to the scheme, host, and base path provided in target. If the
// target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
// NewSingleHostReverseProxy does not rewrite the Host header.
// To rewrite Host headers, use ReverseProxy directly with a custom
// Director policy.
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy

// Hop-by-hop headers. These are removed when sent to the backend.
// As of RFC 7230, hop-by-hop headers are required to appear in the
// Connection header field. These are the headers defined by the
// obsoleted RFC 2616 (section 13.5.1) and are used for backward
// compatibility.

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)

// switchProtocolCopier exists so goroutines proxying data back and
// forth have nice names in stacks.
