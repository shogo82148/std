// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
)

// CrossOriginProtection implements protections against [Cross-Site Request
// Forgery (CSRF)] by rejecting non-safe cross-origin browser requests.
//
// Cross-origin requests are currently detected with the [Sec-Fetch-Site]
// header, available in all browsers since 2023, or by comparing the hostname of
// the [Origin] header with the Host header.
//
// The GET, HEAD, and OPTIONS methods are [safe methods] and are always allowed.
// It's important that applications do not perform any state changing actions
// due to requests with safe methods.
//
// Requests without Sec-Fetch-Site or Origin headers are currently assumed to be
// either same-origin or non-browser requests, and are allowed.
//
// The zero value of CrossOriginProtection is valid and has no trusted origins
// or bypass patterns.
//
// [Sec-Fetch-Site]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site
// [Origin]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
// [Cross-Site Request Forgery (CSRF)]: https://developer.mozilla.org/en-US/docs/Web/Security/Attacks/CSRF
// [safe methods]: https://developer.mozilla.org/en-US/docs/Glossary/Safe/HTTP
type CrossOriginProtection struct {
	bypass    atomic.Pointer[ServeMux]
	trustedMu sync.RWMutex
	trusted   map[string]bool
	deny      atomic.Pointer[Handler]
}

// NewCrossOriginProtection returns a new [CrossOriginProtection] value.
func NewCrossOriginProtection() *CrossOriginProtection

// AddTrustedOrigin allows all requests with an [Origin] header
// which exactly matches the given value.
//
// Origin header values are of the form "scheme://host[:port]".
//
// AddTrustedOrigin can be called concurrently with other methods
// or request handling, and applies to future requests.
//
// [Origin]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
func (c *CrossOriginProtection) AddTrustedOrigin(origin string) error

// AddInsecureBypassPattern permits all requests that match the given pattern.
//
// The pattern syntax and precedence rules are the same as [ServeMux]. Only
// requests that match the pattern directly are permitted. Those that ServeMux
// would redirect to a pattern (e.g. after cleaning the path or adding a
// trailing slash) are not.
//
// AddInsecureBypassPattern panics if the pattern conflicts with one already
// registered, or if the pattern is syntactically invalid (for example, an
// improperly formed wildcard).
//
// AddInsecureBypassPattern can be called concurrently with other methods or
// request handling, and applies to future requests.
func (c *CrossOriginProtection) AddInsecureBypassPattern(pattern string)

// SetDenyHandler sets a handler to invoke when a request is rejected.
// The default error handler responds with a 403 Forbidden status.
//
// SetDenyHandler can be called concurrently with other methods
// or request handling, and applies to future requests.
//
// Check does not call the error handler.
func (c *CrossOriginProtection) SetDenyHandler(h Handler)

// Check applies cross-origin checks to a request.
// It returns an error if the request should be rejected.
func (c *CrossOriginProtection) Check(req *Request) error

// Handler returns a handler that applies cross-origin checks
// before invoking the handler h.
//
// If a request fails cross-origin checks, the request is rejected
// with a 403 Forbidden status or handled with the handler passed
// to [CrossOriginProtection.SetDenyHandler].
func (c *CrossOriginProtection) Handler(h Handler) Handler
