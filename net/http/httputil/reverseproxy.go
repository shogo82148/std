// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP reverse proxy handler

package httputil

import (
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"
)

// onExitFlushLoop is a callback set by tests to detect the state of the
// flushLoop() goroutine.

// ReverseProxy is an HTTP Handler that takes an incoming request and
// sends it to another server, proxying the response back to the
// client.
type ReverseProxy struct {
	Director func(*http.Request)

	Transport http.RoundTripper

	FlushInterval time.Duration
}

// NewSingleHostReverseProxy returns a new ReverseProxy that rewrites
// URLs to the scheme, host, and base path provided in target. If the
// target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)
