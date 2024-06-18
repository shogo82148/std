// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/time"
)

// Cookieは、HTTP応答のSet-CookieヘッダーまたはHTTPリクエストのCookieヘッダーで送信されるHTTPクッキーを表します。
//
// 詳細については、https://tools.ietf.org/html/rfc6265 を参照してください。
type Cookie struct {
	Name  string
	Value string

	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0は 'Max-Age'属性が指定されていないことを意味します。
	// MaxAge<0は今すぐクッキーを削除することを意味し、'Max-Age: 0'と同等です。
	// MaxAge>0はMax-Age属性が存在し、秒で指定されていることを意味します。
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string
}

// SameSiteは、サーバーがクッキー属性を定義して、ブラウザがクロスサイトリクエストと一緒にこのクッキーを送信できなくすることを可能にします。主な目的は、クロスオリジン情報漏洩のリスクを軽減し、クロスサイトリクエスト偽造攻撃に対する保護を提供することです。
//
// 詳細については、https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 を参照してください。
type SameSite int

const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)

// ParseCookieはCookieヘッダーの値を解析し、それに設定されていたすべてのクッキーを返します。
// 同じクッキー名が複数回現れることがあるため、返されるValuesには、特定のキーに対して複数の値が含まれることがあります。
func ParseCookie(line string) ([]*Cookie, error)

// ParseSetCookieはSet-Cookieヘッダーの値を解析し、クッキーを返します。
// 構文エラーがある場合はエラーを返します。
func ParseSetCookie(line string) (*Cookie, error)

// SetCookieは、提供された [ResponseWriter] のヘッダーにSet-Cookieヘッダーを追加します。
// 提供されたクッキーには有効な名前が必要です。無効なクッキーは黙って破棄される場合があります。
func SetCookie(w ResponseWriter, cookie *Cookie)

// Stringは、[Cookie] ヘッダー（NameとValueのみが設定されている場合）またはSet-Cookie応答ヘッダー（他のフィールドが設定されている場合）で使用するためのクッキーのシリアル化を返します。
// cがnilであるか、c.Nameが無効な場合、空の文字列が返されます。
func (c *Cookie) String() string

// Validは、クッキーが有効かどうかを報告します。
func (c *Cookie) Valid() error
