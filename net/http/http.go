// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

package http

import (
	"github.com/shogo82148/std/io"
)

// NoBodyはバイトを持たない [io.ReadCloser] です。Readは常にEOFを返し、
// Closeは常にnilを返します。これは、リクエストがゼロバイトであることを
// 明示的に示すために、送信元クライアントのリクエストで使用することができます。
// ただし、代わりに [Request.Body] をnilに設定することもできます。
var NoBody = noBody{}

var (
	// NoBodyからのio.Copyがバッファを必要としないことを検証する
	_ io.WriterTo   = NoBody
	_ io.ReadCloser = NoBody
)

// PushOptionsは、[Pusher.Push] のオプションを記述します。
type PushOptions struct {

	// Methodは要求されたリクエストのHTTPメソッドを指定します。
	// 設定する場合、"GET"または"HEAD"でなければなりません。空は"GET"を意味します。
	Method string

	// Headerは追加の約束されたリクエストヘッダーを指定します。これには":path"や":scheme"などのHTTP/2疑似ヘッダーフィールドは含めることができませんが、これらは自動的に追加されます。
	Header Header
}

// Pusherは、HTTP/2サーバープッシュをサポートするResponseWritersによって実装されるインターフェースです。
// 詳細については、 https://tools.ietf.org/html/rfc7540#section-8.2 を参照してください。
type Pusher interface {
	Push(target string, opts *PushOptions) error
}
