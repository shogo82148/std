// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP Response reading and parsing.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/url"
)

// Response represents the response from an HTTP request.
type Response struct {
	Status     string
	StatusCode int
	Proto      string
	ProtoMajor int
	ProtoMinor int

	Header Header

	Body io.ReadCloser

	ContentLength int64

	TransferEncoding []string

	Close bool

	Trailer Header

	Request *Request
}

// Cookies parses and returns the cookies set in the Set-Cookie headers.
func (r *Response) Cookies() []*Cookie

var ErrNoLocation = errors.New("http: no Location header in response")

// Location returns the URL of the response's "Location" header,
// if present.  Relative redirects are resolved relative to
// the Response's Request.  ErrNoLocation is returned if no
// Location header is present.
func (r *Response) Location() (*url.URL, error)

// ReadResponse reads and returns an HTTP response from r.  The
// req parameter specifies the Request that corresponds to
// this Response.  Clients must call resp.Body.Close when finished
// reading resp.Body.  After that call, clients can inspect
// resp.Trailer to find key/value pairs included in the response
// trailer.
func ReadResponse(r *bufio.Reader, req *Request) (resp *Response, err error)

// ProtoAtLeast returns whether the HTTP protocol used
// in the response is at least major.minor.
func (r *Response) ProtoAtLeast(major, minor int) bool

// Writes the response (header, body and trailer) in wire format. This method
// consults the following fields of the response:
//
//	StatusCode
//	ProtoMajor
//	ProtoMinor
//	RequestMethod
//	TransferEncoding
//	Trailer
//	Body
//	ContentLength
//	Header, values for non-canonical keys will have unpredictable behavior
func (r *Response) Write(w io.Writer) error
