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
	// SetCookiesは、指定されたURLの応答でのクッキーの受け取りを処理します。
	// ジャーのポリシーと実装により、クッキーを保存するかどうかを選択するかもしれませんし、しないかもしれません。
	SetCookies(u *url.URL, cookies []*Cookie)

	// Cookiesは、指定されたURLのリクエストで送信するクッキーを返します。
	// 標準的なクッキー使用制限（RFC 6265など）を尊重するかどうかは、実装次第です。
	Cookies(u *url.URL) []*Cookie
}
