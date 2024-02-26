// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/shogo82148/std/net/http"
)

// DumpRequestOutは、outgoingのクライアントリクエスト用の [DumpRequest] のようなものです。これには、標準の [http.Transport] が追加するUser-Agentなど、任意のヘッダーが含まれます。
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)

// DumpRequestは与えられたリクエストをHTTP/1.xのワイヤープレゼンテーションで返します。
// クライアントのリクエストをデバッグするために、サーバーが使用するべきです。
// 返されるプレゼンテーションは近似値です。初期リクエストの一部の詳細は、[http.Request] に解析される際に失われます。
// 特にヘッダーフィールド名の順序と大文字小文字の情報が失われます。複数の値を持つヘッダーの値の順序は保持されます。
// HTTP/2のリクエストは、元のバイナリ表現ではなく、HTTP/1.xの形式でダンプされます。
//
// bodyがtrueの場合、DumpRequestはbodyも返します。そのため、req.Bodyを消費し、同じバイトを返す新しい [io.ReadCloser] に置き換えます。
// DumpRequestがエラーを返す場合、reqの状態は未定義です。
//
// [http.Request.Write] のドキュメントには、ダンプに含まれるreqのフィールドの詳細が記載されています。
func DumpRequest(req *http.Request, body bool) ([]byte, error)

// DumpResponseはDumpRequestと同様ですが、レスポンスをダンプします。
func DumpResponse(resp *http.Response, body bool) ([]byte, error)
