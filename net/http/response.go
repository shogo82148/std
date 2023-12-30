// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP Response reading and parsing.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/url"
)

// Responseは、HTTPリクエストからのレスポンスを表します。
//
// ClientとTransportは、レスポンスヘッダが受信された後、サーバーからResponsesを返します。
// レスポンスボディは、Bodyフィールドが読み取られるたびにオンデマンドでストリーミングされます。
type Response struct {
	Status     string
	StatusCode int
	Proto      string
	ProtoMajor int
	ProtoMinor int

	// Headerは、ヘッダーキーを値にマップします。
	// レスポンスに同じキーを持つ複数のヘッダーがある場合、それらはカンマ区切りで連結される場合があります。
	// (RFC 7230 Section 3.2.2 では、複数のヘッダーがカンマ区切りのシーケンスとして意味的に等価である必要があります。)
	// Headerの値がこの構造体の他のフィールド(ContentLength、TransferEncoding、Trailerなど)によって重複する場合、フィールド値が優先されます。
	//
	// マップ内のキーは正準化されます(CanonicalHeaderKeyを参照)。
	Header Header

	// Bodyは、レスポンスボディを表します。
	//
	// レスポンスボディは、Bodyフィールドが読み取られるたびにオンデマンドでストリーミングされます。
	// ネットワーク接続が失敗した場合や、サーバーがレスポンスを終了した場合、Body.Read呼び出しはエラーを返します。
	//
	// http ClientおよびTransportは、Bodyが常にnilでないことを保証します。
	// これは、ボディのないレスポンスまたは長さがゼロのボディを持つレスポンスでも同様です。
	// Bodyを閉じるのは呼び出し側の責任です。
	// デフォルトのHTTPクライアントのTransportは、Bodyが完全に読み取られて閉じられていない場合、
	// HTTP/1.xの「keep-alive」TCP接続を再利用しない場合があります。
	//
	// サーバーが「chunked」Transfer-Encodingで返答した場合、Bodyは自動的にデチャンクされます。
	//
	// Go 1.12以降、Bodyは、WebSocketsおよびHTTP/2の「h2c」モードで使用される「101 Switching Protocols」レスポンスに成功した場合、io.Writerも実装します。
	Body io.ReadCloser

	// ContentLengthは、関連するコンテンツの長さを記録します。
	// 値-1は、長さが不明であることを示します。
	// Request.Methodが"HEAD"でない限り、0以上の値は、Bodyから指定されたバイト数を読み取ることができることを示します。
	ContentLength int64

	// TransferEncodingは、最も外側から最も内側の転送エンコーディングを含みます。
	// 値がnilの場合、"identity"エンコーディングが使用されます。
	TransferEncoding []string

	// Closeは、ヘッダーがBodyの読み取り後に接続を閉じるよう指示したかどうかを記録します。
	// この値はクライアントに対するアドバイスです。
	// ReadResponseまたはResponse.Writeは、接続を閉じません。
	Close bool

	// Uncompressedは、レスポンスが圧縮されて送信されたが、httpパッケージによって解凍されたかどうかを報告します。
	// trueの場合、Bodyから読み取ると、サーバーから実際に設定された圧縮されたコンテンツの代わりに、解凍されたコンテンツが返されます。
	// ContentLengthは-1に設定され、"Content-Length"および"Content-Encoding"フィールドはresponseHeaderから削除されます。
	// サーバーからの元のレスポンスを取得するには、Transport.DisableCompressionをtrueに設定します。
	Uncompressed bool

	// Trailerは、トレーラーキーをHeaderと同じ形式の値にマップします。
	//
	// Trailerには最初に、サーバーの"Trailer"ヘッダー値で指定されたキーごとに1つのnil値のみが含まれます。
	// これらの値はHeaderに追加されません。
	//
	// BodyのRead呼び出しと同時にTrailerにアクセスしないでください。
	//
	// Body.Readがio.EOFを返した後、Trailerにはサーバーから送信されたトレーラー値が含まれます。
	Trailer Header

	// Requestは、このResponseを取得するために送信されたリクエストです。
	// RequestのBodyはnilです(すでに消費されています)。
	// これは、クライアントリクエストに対してのみ設定されます。
	Request *Request

	// TLSは、レスポンスを受信したTLS接続に関する情報を含みます。
	// 非暗号化のレスポンスの場合、nilです。
	// このポインタはレスポンス間で共有され、変更してはいけません。
	TLS *tls.ConnectionState
}

// Cookiesは、Set-Cookieヘッダーで設定されたCookieを解析して返します。
func (r *Response) Cookies() []*Cookie

// ErrNoLocationは、LocationメソッドがLocationヘッダーが存在しない場合に返されます。
var ErrNoLocation = errors.New("http: no Location header in response")

// Locationは、レスポンスの「Location」ヘッダーのURLを返します。
// 存在する場合、相対リダイレクトはレスポンスのリクエストに対して相対的に解決されます。
// Locationヘッダーが存在しない場合、ErrNoLocationが返されます。
func (r *Response) Location() (*url.URL, error)

// ReadResponseは、rからHTTPレスポンスを読み取り、返します。
// reqパラメータは、このResponseに対応するRequestをオプションで指定します。
// nilの場合、GETリクエストが想定されます。
// クライアントは、resp.Bodyを読み取り終えたらresp.Body.Closeを呼び出す必要があります。
// その呼び出しの後、クライアントはresp.Trailerを調べて、レスポンストレーラーに含まれるキー/値ペアを見つけることができます。
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)

// ProtoAtLeastは、レスポンスで使用されるHTTPプロトコルが少なくともmajor.minorであるかどうかを報告します。
func (r *Response) ProtoAtLeast(major, minor int) bool

// Writeは、HTTP/1.xサーバーレスポンス形式でrをwに書き込みます。
// ステータス行、ヘッダー、ボディ、およびオプションのトレーラーを含みます。
//
// このメソッドは、レスポンスrの以下のフィールドを参照します。
//
//	StatusCode
//	ProtoMajor
//	ProtoMinor
//	Request.Method
//	TransferEncoding
//	Trailer
//	Body
//	ContentLength
//	Header, 非正準化キーの値は予測不可能な動作をします
//
// レスポンスボディは送信後に閉じられます。
func (r *Response) Write(w io.Writer) error
