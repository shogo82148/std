// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージhttptestは、HTTPテストのためのユーティリティを提供します。
package httptest

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/http"
)

<<<<<<< HEAD
// NewRequestはテスト用に、[http.Handler] に渡すことができる新しい受信サーバーリクエストを返します。
=======
// NewRequest wraps NewRequestWithContext using context.Background.
func NewRequest(method, target string, body io.Reader) *http.Request

// NewRequestWithContext returns a new incoming server Request, suitable
// for passing to an [http.Handler] for testing.
>>>>>>> upstream/master
//
// targetはRFC 7230の「要求ターゲット」です。パスまたは絶対URLのいずれかを使用できます。targetが絶対URLの場合、URLからホスト名が使用されます。それ以外の場合は、"example.com"が使用されます。
//
// targetのスキームが「https」の場合、TLSフィールドは非nilのダミー値に設定されます。
//
// Request.Protoは常にHTTP/1.1です。
//
// 空のメソッドは「GET」を意味します。
//
// 指定されたbodyはnilである場合があります。bodyが*bytes.Reader、*strings.Reader、または*bytes.Bufferの型の場合、Request.ContentLengthが設定されます。
//
// NewRequestはエラー時にパニックを発生させます。テストではパニックが許容されるため、使用の便宜上です。
//
<<<<<<< HEAD
// サーバーリクエストではなく、クライアントのHTTPリクエストを生成するには、net/httpパッケージのNewRequest関数を参照してください。
func NewRequest(method, target string, body io.Reader) *http.Request
=======
// To generate a client HTTP request instead of a server request, see
// the NewRequest function in the net/http package.
func NewRequestWithContext(ctx context.Context, method, target string, body io.Reader) *http.Request
>>>>>>> upstream/master
