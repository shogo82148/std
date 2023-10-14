// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

package http

import (
	"github.com/shogo82148/std/io"
)

// NoBodyはバイトを持たないio.ReadCloserです。Readは常にEOFを返し、
// Closeは常にnilを返します。これは、リクエストがゼロバイトであることを
// 明示的に示すために、送信元クライアントのリクエストで使用することができます。
// ただし、代わりにRequest.Bodyをnilに設定することもできます。
var NoBody = noBody{}

var (
	// NoBodyからのio.Copyがバッファを必要としないことを検証する
	_ io.WriterTo   = NoBody
	_ io.ReadCloser = NoBody
)

// PushOptionsは、Pusher.Pushのオプションを記述します。
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
	// Push initiates an HTTP/2 server push. This constructs a synthetic
	// request using the given target and options, serializes that request
	// into a PUSH_PROMISE frame, then dispatches that request using the
	// server's request handler. If opts is nil, default options are used.
	//
	// The target must either be an absolute path (like "/path") or an absolute
	// URL that contains a valid host and the same scheme as the parent request.
	// If the target is a path, it will inherit the scheme and host of the
	// parent request.
	//
	// The HTTP/2 spec disallows recursive pushes and cross-authority pushes.
	// Push may or may not detect these invalid pushes; however, invalid
	// pushes will be detected and canceled by conforming clients.
	//
	// Handlers that wish to push URL X should call Push before sending any
	// data that may trigger a request for URL X. This avoids a race where the
	// client issues requests for X before receiving the PUSH_PROMISE for X.
	//
	// Push will run in a separate goroutine making the order of arrival
	// non-deterministic. Any required synchronization needs to be implemented
	// by the caller.
	//
	// Push returns ErrNotSupported if the client has disabled push or if push
	// is not supported on the underlying connection.
	Push(target string, opts *PushOptions) error
}
