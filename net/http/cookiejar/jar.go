// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cookiejar はメモリ内で RFC 6265 に準拠した http.CookieJar を実装します。
package cookiejar

import (
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/sync"
)

// PublicSuffixListはドメインの公開サフィックスを提供します。例えば：
//   - "example.com"の公開サフィックスは「com」です。
//   - "foo1.foo2.foo3.co.uk"の公開サフィックスは「co.uk」です。
//   - "bar.pvt.k12.ma.us"の公開サフィックスは「pvt.k12.ma.us」です。
//
// PublicSuffixListの実装は、複数のゴルーチンに対して安全に同時に使用できる必要があります。
//
// 常に""を返す実装は有効であり、テストには便利ですが、安全ではありません。これは、foo.comのHTTPサーバがbar.comのためにクッキーを設定できることを意味します。
//
// golang.org/x/net/publicsuffixパッケージには、公開サフィックスリストの実装があります。
type PublicSuffixList interface {
	PublicSuffix(domain string) string

	String() string
}

// Options は新しい Jar の作成オプションです。
type Options struct {

	// PublicSuffixListは、ドメインに対してHTTPサーバがクッキーを設定できるかどうかを決定する公開サフィックスリストです。
	//
	// nilの値は有効であり、テストには便利ですが、セキュリティ上の理由から使用するべきではありません。これは、foo.co.ukのHTTPサーバがbar.co.ukに対してクッキーを設定できることを意味します。
	PublicSuffixList PublicSuffixList
}

// Jarはnet/httpパッケージのhttp.CookieJarインターフェースを実装しています。
type Jar struct {
	psList PublicSuffixList

	// muは残りのフィールドをロックします。
	mu sync.Mutex

	// entriesは、eTLD+1でキー付けされ、その名前/ドメイン/パスでサブキー付けされたエントリのセットです。
	entries map[string]map[string]entry

	// nextSeqNumは新しいクッキーが作成されたSetCookiesに割り当てられる次のシーケンス番号です。
	nextSeqNum uint64
}

<<<<<<< HEAD
// Newは新しいクッキージャーを返します。nilの*OptionsはゼロのOptionsと同等です。
func New(o *Options) (*Jar, error)

// Cookiesはhttp.CookieJarインターフェースのCookiesメソッドを実装しています。
=======
// New returns a new cookie jar. A nil [*Options] is equivalent to a zero
// Options.
func New(o *Options) (*Jar, error)

// Cookies implements the Cookies method of the [http.CookieJar] interface.
>>>>>>> upstream/master
//
// URLのスキームがHTTPまたはHTTPSでない場合、空のスライスを返します。
func (j *Jar) Cookies(u *url.URL) (cookies []*http.Cookie)

<<<<<<< HEAD
// SetCookiesはhttp.CookieJarインターフェースのSetCookiesメソッドを実装します。
=======
// SetCookies implements the SetCookies method of the [http.CookieJar] interface.
>>>>>>> upstream/master
//
// URLのスキームがHTTPまたはHTTPSでない場合、何もしません。
func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie)
