// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/net/url"
)

// CookieJarは、HTTPリクエストでのCookieのストレージと使用を管理します。
//
// CookieJarの実装は、複数のゴルーチンによる同時使用に対して安全である必要があります。
//
// net/http/cookiejarパッケージはCookieJarの実装を提供します。
type CookieJar interface {
	SetCookies(u *url.URL, cookies []*Cookie)

	Cookies(u *url.URL) []*Cookie
}
