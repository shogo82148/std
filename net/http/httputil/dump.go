// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/shogo82148/std/net/http"
)

// dumpConn is a net.Conn which writes to Writer and reads from Reader

// DumpRequestOut is like DumpRequest but for outgoing client requests. It
// includes any headers that the standard http.Transport adds, such as
// User-Agent.
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)

// delegateReader is a reader that delegates to another reader,
// once it arrives on a channel.

// DumpRequest returns the given request in its HTTP/1.x wire
// representation. It should only be used by servers to debug client
// requests. The returned representation is an approximation only;
// some details of the initial request are lost while parsing it into
// an http.Request. In particular, the order and case of header field
// names are lost. The order of values in multi-valued headers is kept
// intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
// original binary representations.
//
// If body is true, DumpRequest also returns the body. To do so, it
// consumes req.Body and then replaces it with a new io.ReadCloser
// that yields the same bytes. If DumpRequest returns an error,
// the state of req is undefined.
//
// The documentation for http.Request.Write details which fields
// of req are included in the dump.
func DumpRequest(req *http.Request, body bool) ([]byte, error)

// errNoBody is a sentinel error value used by failureToReadBody so we
// can detect that the lack of body was intentional.

// failureToReadBody is an io.ReadCloser that just returns errNoBody on
// Read. It's swapped in when we don't actually want to consume
// the body, but need a non-nil one, and want to distinguish the
// error from reading the dummy body.

// emptyBody is an instance of empty reader.

// DumpResponse is like DumpRequest but dumps a response.
func DumpResponse(resp *http.Response, body bool) ([]byte, error)
