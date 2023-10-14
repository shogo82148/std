// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package web defines minimal helper routines for accessing HTTP/HTTPS
// resources without requiring external dependencies on the net package.
//
// If the cmd_go_bootstrap build tag is present, web avoids the use of the net
// package and returns errors for all network operations.
package web

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/url"
)

// SecurityMode specifies whether a function should make network
// calls using insecure transports (eg, plain text HTTP).
// The zero value is "secure".
type SecurityMode int

const (
	SecureOnly SecurityMode = iota
	DefaultSecurity
	Insecure
)

// An HTTPError describes an HTTP error response (non-200 result).
type HTTPError struct {
	URL        string
	Status     string
	StatusCode int
	Err        error
	Detail     string
}

func (e *HTTPError) Error() string

func (e *HTTPError) Is(target error) bool

func (e *HTTPError) Unwrap() error

// GetBytes returns the body of the requested resource, or an error if the
// response status was not http.StatusOK.
//
// GetBytes is a convenience wrapper around Get and Response.Err.
func GetBytes(u *url.URL) ([]byte, error)

type Response struct {
	URL        string
	Status     string
	StatusCode int
	Header     map[string][]string
	Body       io.ReadCloser

	fileErr     error
	errorDetail errorDetailBuffer
}

// Err returns an *HTTPError corresponding to the response r.
// If the response r has StatusCode 200 or 0 (unset), Err returns nil.
// Otherwise, Err may read from r.Body in order to extract relevant error detail.
func (r *Response) Err() error

// Get returns the body of the HTTP or HTTPS resource specified at the given URL.
//
// If the URL does not include an explicit scheme, Get first tries "https".
// If the server does not respond under that scheme and the security mode is
// Insecure, Get then tries "http".
// The URL included in the response indicates which scheme was actually used,
// and it is a redacted URL suitable for use in error messages.
//
// For the "https" scheme only, credentials are attached using the
// cmd/go/internal/auth package. If the URL itself includes a username and
// password, it will not be attempted under the "http" scheme unless the
// security mode is Insecure.
//
// Get returns a non-nil error only if the request did not receive a response
// under any applicable scheme. (A non-2xx response does not cause an error.)
func Get(security SecurityMode, u *url.URL) (*Response, error)

// OpenBrowser attempts to open the requested URL in a web browser.
func OpenBrowser(url string) (opened bool)

// Join returns the result of adding the slash-separated
// path elements to the end of u's path.
func Join(u *url.URL, path string) *url.URL

// IsLocalHost reports whether the given URL refers to a local
// (loopback) host, such as "localhost" or "127.0.0.1:8080".
func IsLocalHost(u *url.URL) bool
