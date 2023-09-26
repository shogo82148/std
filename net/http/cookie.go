// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/time"
)

// A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
// HTTP response or the Cookie header of an HTTP request.
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string
}

// SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
func SetCookie(w ResponseWriter, cookie *Cookie)

// String returns the serialization of the cookie for use in a Cookie
// header (if only Name and Value are set) or a Set-Cookie response
// header (if other fields are set).
func (c *Cookie) String() string
