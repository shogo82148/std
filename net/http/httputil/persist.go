// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/textproto"
	"github.com/shogo82148/std/sync"
)

var (
	// Deprecated: 使用されていません。
	ErrPersistEOF = &http.ProtocolError{ErrorString: "persistent connection closed"}

	// Deprecated: 使用されていません。
	ErrClosed = &http.ProtocolError{ErrorString: "connection closed by user"}

	// Deprecated: 使用されていません。
	ErrPipeline = &http.ProtocolError{ErrorString: "pipeline error"}
)

// ServerConnはGoの初期のHTTP実装の遺物です。
// それは低レベルで古く、Goの現在のHTTPスタックでは使用されていません。
// Go 1よりも前に削除すべきでした。
//
// Deprecated: 代わりに [net/http] パッケージのServerを使用してください。
type ServerConn struct {
	mu              sync.Mutex
	c               net.Conn
	r               *bufio.Reader
	re, we          error
	lastbody        io.ReadCloser
	nread, nwritten int
	pipereq         map[*http.Request]uint

	pipe textproto.Pipeline
}

// NewServerConnはGoの初期のHTTP実装の遺物です。
// これは低レベルで古いものであり、現行のGoのHTTPスタックでは使用されていません。
// Go 1より前に削除すべきでした。
//
// Deprecated: 代わりに package [net/http] のServerを使用してください。
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn

// Hijackは [ServerConn] を切り離し、基礎となる接続と、残っているデータを持つ読み込み側のbufioを返します。 HijackはReadがkeep-aliveロジックの終了を示すまえに呼び出される可能性があります。ユーザーは [ServerConn.Read] や [ServerConn.Write] が進行中の間にHijackを呼び出すべきではありません。
func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)

// Closeによって [ServerConn.Hijack] され、その後基礎となる接続も閉じます。
func (sc *ServerConn) Close() error

// Readはワイヤ上の次のリクエストを返します。[ErrPersistEOF] は、優雅にもうリクエストがないことが確定した場合に返されます（例えば、HTTP/1.0接続の最初のリクエスト後、またはHTTP/1.1接続のConnection: close後など）。
func (sc *ServerConn) Read() (*http.Request, error)

// Pendingは、接続されたリクエストの未回答数を返します。
func (sc *ServerConn) Pending() int

// Writeはreqに応じたrespを書き込みます。接続を正常に終了させるためには、
// Response.Closeフィールドをtrueに設定してください。Writeは、[ServerConn.Read] 側で返される
// すべてのエラーに関係なく、エラーが返されるまで操作可能であると見なされるべきです。
func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error

// ClientConnはGoの初期のHTTP実装の遺物です。
// 低レベルで古く、現在のGoのHTTPスタックでは使用されていません。
// Go 1の前に削除すべきでした。
//
// Deprecated: 代わりにpackage [net/http] のClientまたはTransportを使用してください。
type ClientConn struct {
	mu              sync.Mutex
	c               net.Conn
	r               *bufio.Reader
	re, we          error
	lastbody        io.ReadCloser
	nread, nwritten int
	pipereq         map[*http.Request]uint

	pipe     textproto.Pipeline
	writeReq func(*http.Request, io.Writer) error
}

// NewClientConnはGoの初期のHTTP実装の遺産です。
// これは低レベルで古く、現在のGoのHTTPスタックでは使用されていません。
// Go 1より前に削除すべきでした。
//
// Deprecated: 代わりにパッケージ [net/http] のClientまたはTransportを使用してください。
func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// NewProxyClientConn はGo言語の初期のHTTP実装の遺物です。
// これは低レベルで古く、現在のGoのHTTPスタックでは使用されていません。
// Go 1 より前に削除する必要があります。
//
// Deprecated: 代わりに package [net/http] の Client または Transport を使用してください。
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// Hijackは [ClientConn] を切り離し、基礎となる接続と読み込み側のbufioを返します。
// また、左にデータが残っているかもしれない読み込み側のbufioも返します。
// HijackはユーザーまたはReadがkeep-aliveロジックの終了をシグナルした前に呼び出すことができます。
// ユーザーは、[ClientConn.Read] または ClientConn.Writeが進行中の間にHijackを呼び出さないでください。
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)

// Closeは [ClientConn.Hijack] を呼び出し、その後下層の接続も閉じます。
func (cc *ClientConn) Close() error

// Writeはリクエストを書き込みます。もしHTTP keep-aliveによって接続が閉じられた場合、[ErrPersistEOF] エラーが返されます。もしreq.Closeがtrueの場合、このリクエストの後にkeep-alive接続が論理的に閉じられ、対向サーバーに通知されます。ErrUnexpectedEOFは、リモートが基礎となるTCP接続を閉じたことを示しており、通常は正常な終了と見なされます。
func (cc *ClientConn) Write(req *http.Request) error

// Pendingは、接続に送信された未応答のリクエストの数を返します。
func (cc *ClientConn) Pending() int

// Readはワイヤから次のレスポンスを読み込みます。有効なレスポンスは [ErrPersistEOF] と一緒に返される場合があります。これはリモートがこれがサービスされる最後のリクエストであることを要求したことを意味します。Readは [ClientConn.Write] と同時に呼び出すことができますが、他のReadと同時には呼び出すことはできません。
func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)

// Doはリクエストを書き込み、レスポンスを読み取る便利なメソッドです。
func (cc *ClientConn) Do(req *http.Request) (*http.Response, error)
