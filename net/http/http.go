// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

package http

import (
	"github.com/shogo82148/std/io"
)

// NoBody is an io.ReadCloser with no bytes. Read always returns EOF
// and Close always returns nil. It can be used in an outgoing client
// request to explicitly signal that a request has zero bytes.
// An alternative, however, is to simply set Request.Body to nil.
var NoBody = noBody{}

var (
	// verify that an io.Copy from NoBody won't require a buffer:
	_ io.WriterTo   = NoBody
	_ io.ReadCloser = NoBody
)

// PushOptionsは、Pusher.Pushのオプションを記述します。
type PushOptions struct {
	// Method specifies the HTTP method for the promised request.
	// If set, it must be "GET" or "HEAD". Empty means "GET".
	Method string

	// Header specifies additional promised request headers. This cannot
	// include HTTP/2 pseudo header fields like ":path" and ":scheme",
	// which will be added automatically.
	Header Header
}

// Pusherは、HTTP/2サーバープッシュをサポートするResponseWritersによって実装されるインターフェースです。
// 詳細については、 https://tools.ietf.org/html/rfc7540#section-8.2 を参照してください。
type Pusher interface {
	Push(target string, opts *PushOptions) error
}
