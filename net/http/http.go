// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

package http

import (
	"github.com/shogo82148/std/io"
)

// incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).

// maxInt64 is the effective "infinite" value for the Server and
// Transport's byte-limiting readers.

// aLongTimeAgo is a non-zero time, far in the past, used for
// immediate cancellation of network operations.

// omitBundledHTTP2 is set by omithttp2.go when the nethttpomithttp2
// build tag is set. That means h2_bundle.go isn't compiled in and we
// shouldn't try to use it.

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.

// NoBodyは、バイトを持たないio.ReadCloserです。Readは常にEOFを返し、Closeは常にnilを返します。
// リクエストに0バイトがあることを明示的に示すために、送信クライアントのリクエストで使用できます。
// ただし、Request.Bodyをnilに設定する代替手段もあります。
var NoBody = noBody{}

var (
	// verify that an io.Copy from NoBody won't require a buffer:
	_ io.WriterTo   = NoBody
	_ io.ReadCloser = NoBody
)

// PushOptionsは、Pusher.Pushのオプションを記述します。
type PushOptions struct {
	Method string

	Header Header
}

// Pusherは、HTTP/2サーバープッシュをサポートするResponseWritersによって実装されるインターフェースです。
// 詳細については、 https://tools.ietf.org/html/rfc7540#section-8.2 を参照してください。
type Pusher interface {
	Push(target string, opts *PushOptions) error
}
