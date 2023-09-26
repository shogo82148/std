// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package httptest provides utilities for HTTP testing.
package httptest

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/net/http"
)

// ResponseRecorder is an implementation of http.ResponseWriter that
// records its mutations for later inspection in tests.
type ResponseRecorder struct {
	Code      int
	HeaderMap http.Header
	Body      *bytes.Buffer
	Flushed   bool

	wroteHeader bool
}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *ResponseRecorder

// DefaultRemoteAddr is the default remote address to return in RemoteAddr if
// an explicit DefaultRemoteAddr isn't set on ResponseRecorder.
const DefaultRemoteAddr = "1.2.3.4"

// Header returns the response headers.
func (rw *ResponseRecorder) Header() http.Header

// Write always succeeds and writes to rw.Body, if not nil.
func (rw *ResponseRecorder) Write(buf []byte) (int, error)

// WriteString always succeeds and writes to rw.Body, if not nil.
func (rw *ResponseRecorder) WriteString(str string) (int, error)

// WriteHeader sets rw.Code.
func (rw *ResponseRecorder) WriteHeader(code int)

// Flush sets rw.Flushed to true.
func (rw *ResponseRecorder) Flush()
