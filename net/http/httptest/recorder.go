// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httptest

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/net/http"
)

// ResponseRecorderはhttp.ResponseWriterの実装であり、テストで後で検査するためにその変更を記録します。
type ResponseRecorder struct {
	// CodeはWriteHeaderが設定したHTTPレスポンスコードです。
	//
	// HandlerがWriteHeaderやWriteを呼び出さない場合、これは暗黙のhttp.StatusOKではなく、0になることに注意してください。暗黙の値を取得するには、Resultメソッドを使用してください。
	Code int

	// HeaderMapはHandlerによって明示的に設定されたヘッダーを含んでいます。
	// これは内部の詳細です。
	//
	// Deprecated: HeaderMapは歴史的な互換性のために存在しており、使用すべきではありません。
	// ハンドラによって返されるヘッダーにアクセスするためには、Resultメソッドによって返される
	// Response.Headerマップを使用してください。
	HeaderMap http.Header

	// BodyはHandlerのWrite呼び出しで送信されるバッファです。
	// nilの場合、書き込みは黙って破棄されます。
	Body *bytes.Buffer

	// Flushed はHandlerがFlushを呼び出したかどうかを示します。
	Flushed bool

	result      *http.Response
	snapHeader  http.Header
	wroteHeader bool
}

// NewRecorderは初期化されたResponseRecorderを返します。
func NewRecorder() *ResponseRecorder

// DefaultRemoteAddrは、ResponseRecorderに明示的なDefaultRemoteAddrが設定されていない場合に、
// RemoteAddrで返すデフォルトのリモートアドレスです。
const DefaultRemoteAddr = "1.2.3.4"

// Headerはhttp.ResponseWriterを実装します。ハンドラ内で変更するためにレスポンスヘッダーを返します。ハンドラが完了した後に書き込まれたヘッダーをテストするには、Resultメソッドを使用し、返されたResponse値のHeaderを確認してください。
func (rw *ResponseRecorder) Header() http.Header

// Writeはhttp.ResponseWriterを実装します。buf内のデータは、rw.Bodyがnilでない場合にrw.Bodyに書き込まれます。
func (rw *ResponseRecorder) Write(buf []byte) (int, error)

// WriteStringはio.StringWriterを実装します。strのデータは、nilでない場合はrw.Bodyに書き込まれます。
func (rw *ResponseRecorder) WriteString(str string) (int, error)

// WriteHeaderはhttp.ResponseWriterを実装します。
func (rw *ResponseRecorder) WriteHeader(code int)

// Flushはhttp.Flusherを実装します。Flushが呼び出されたかどうかをテストするには、rw.Flushedを参照してください。
func (rw *ResponseRecorder) Flush()

// Resultはハンドラによって生成されたレスポンスを返します。
//
// 返されるレスポンスには、少なくともStatusCode、Header、Body、およびオプションでTrailerが含まれます。
// 将来的にはさらなるフィールドが追加される可能性があるため、テストでは結果をDeepEqualで比較しないよう注意してください。
//
// Response.Headerは、最初の書き込み呼び出し時またはこの呼び出し時のヘッダのスナップショットですが、ハンドラが書き込みを行っていない場合は呼び出し時のものになります。
//
// Response.Bodyは非nilであり、Body.Read呼び出しはio.EOF以外のエラーを返さないことが保証されています。
//
// Resultは、ハンドラの実行が完了した後にのみ呼び出す必要があります。
func (rw *ResponseRecorder) Result() *http.Response
