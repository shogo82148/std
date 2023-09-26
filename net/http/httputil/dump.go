// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/shogo82148/std/net/http"
)

// dumpConn is a net.Conn which writes to Writer and reads from Reader

// DumpRequestOut is like DumpRequest but includes
// headers that the standard http.Transport adds,
// such as User-Agent.
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)

// delegateReader is a reader that delegates to another reader,
// once it arrives on a channel.

// DumpRequest returns the as-received wire representation of req,
// optionally including the request body, for debugging.
// DumpRequest is semantically a no-op, but in order to
// dump the body, it reads the body data into memory and
// changes req.Body to refer to the in-memory copy.
// The documentation for http.Request.Write details which fields
// of req are used.
func DumpRequest(req *http.Request, body bool) (dump []byte, err error)

// errNoBody is a sentinel error value used by failureToReadBody so we can detect
// that the lack of body was intentional.

// failureToReadBody is a io.ReadCloser that just returns errNoBody on
// Read.  It's swapped in when we don't actually want to consume the
// body, but need a non-nil one, and want to distinguish the error
// from reading the dummy body.

// DumpResponse is like DumpRequest but dumps a response.
func DumpResponse(resp *http.Response, body bool) (dump []byte, err error)
