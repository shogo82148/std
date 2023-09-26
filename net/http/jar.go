// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/net/url"
)

// A CookieJar manages storage and use of cookies in HTTP requests.
//
// Implementations of CookieJar must be safe for concurrent use by multiple
// goroutines.
type CookieJar interface {
	SetCookies(u *url.URL, cookies []*Cookie)

	Cookies(u *url.URL) []*Cookie
}
