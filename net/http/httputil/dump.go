// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/shogo82148/std/net/http"
)

<<<<<<< HEAD
// DumpRequestOutは、outgoingのクライアントリクエスト用のDumpRequestのようなものです。これには、標準のhttp.Transportが追加するUser-Agentなど、任意のヘッダーが含まれます。
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)

// DumpRequestは与えられたリクエストをHTTP/1.xのワイヤープレゼンテーションで返します。
// クライアントのリクエストをデバッグするために、サーバーが使用するべきです。
// 返されるプレゼンテーションは近似値です。初期リクエストの一部の詳細は、http.Requestに解析される際に失われます。
// 特にヘッダーフィールド名の順序と大文字小文字の情報が失われます。複数の値を持つヘッダーの値の順序は保持されます。
// HTTP/2のリクエストは、元のバイナリ表現ではなく、HTTP/1.xの形式でダンプされます。
//
// bodyがtrueの場合、DumpRequestはbodyも返します。そのため、req.Bodyを消費し、同じバイトを返す新しいio.ReadCloserに置き換えます。
// DumpRequestがエラーを返す場合、reqの状態は未定義です。
//
// http.Request.Writeのドキュメントには、ダンプに含まれるreqのフィールドの詳細が記載されています。
=======
// DumpRequestOut is like [DumpRequest] but for outgoing client requests. It
// includes any headers that the standard [http.Transport] adds, such as
// User-Agent.
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)

// DumpRequest returns the given request in its HTTP/1.x wire
// representation. It should only be used by servers to debug client
// requests. The returned representation is an approximation only;
// some details of the initial request are lost while parsing it into
// an [http.Request]. In particular, the order and case of header field
// names are lost. The order of values in multi-valued headers is kept
// intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
// original binary representations.
//
// If body is true, DumpRequest also returns the body. To do so, it
// consumes req.Body and then replaces it with a new [io.ReadCloser]
// that yields the same bytes. If DumpRequest returns an error,
// the state of req is undefined.
//
// The documentation for [http.Request.Write] details which fields
// of req are included in the dump.
>>>>>>> upstream/release-branch.go1.22
func DumpRequest(req *http.Request, body bool) ([]byte, error)

// DumpResponseはDumpRequestと同様ですが、レスポンスをダンプします。
func DumpResponse(resp *http.Response, body bool) ([]byte, error)
