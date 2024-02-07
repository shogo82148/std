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
<<<<<<< HEAD
// Deprecated: 代わりにnet/httpパッケージのServerを使用してください。
=======
// Deprecated: Use the Server in package [net/http] instead.
>>>>>>> upstream/release-branch.go1.22
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
<<<<<<< HEAD
// Deprecated: 代わりに package net/http のServerを使用してください。
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn

// HijackはServerConnを切り離し、基礎となる接続と、残っているデータを持つ読み込み側のbufioを返します。 HijackはReadがkeep-aliveロジックの終了を示すまえに呼び出される可能性があります。ユーザーはReadやWriteが進行中の間にHijackを呼び出すべきではありません。
func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)

// Closeによってヒジャックされ、その後基礎となる接続も閉じます。
func (sc *ServerConn) Close() error

// Readはワイヤ上の次のリクエストを返します。ErrPersistEOFは、優雅にもうリクエストがないことが確定した場合に返されます（例えば、HTTP/1.0接続の最初のリクエスト後、またはHTTP/1.1接続のConnection: close後など）。
=======
// Deprecated: Use the Server in package [net/http] instead.
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn

// Hijack detaches the [ServerConn] and returns the underlying connection as well
// as the read-side bufio which may have some left over data. Hijack may be
// called before Read has signaled the end of the keep-alive logic. The user
// should not call Hijack while [ServerConn.Read] or [ServerConn.Write] is in progress.
func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)

// Close calls [ServerConn.Hijack] and then also closes the underlying connection.
func (sc *ServerConn) Close() error

// Read returns the next request on the wire. An [ErrPersistEOF] is returned if
// it is gracefully determined that there are no more requests (e.g. after the
// first request on an HTTP/1.0 connection, or after a Connection:close on a
// HTTP/1.1 connection).
>>>>>>> upstream/release-branch.go1.22
func (sc *ServerConn) Read() (*http.Request, error)

// Pendingは、接続されたリクエストの未回答数を返します。
func (sc *ServerConn) Pending() int

<<<<<<< HEAD
// Writeはreqに応じたrespを書き込みます。接続を正常に終了させるためには、
// Response.Closeフィールドをtrueに設定してください。Writeは、Read側で返される
// すべてのエラーに関係なく、エラーが返されるまで操作可能であると見なされるべきです。
=======
// Write writes resp in response to req. To close the connection gracefully, set the
// Response.Close field to true. Write should be considered operational until
// it returns an error, regardless of any errors returned on the [ServerConn.Read] side.
>>>>>>> upstream/release-branch.go1.22
func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error

// ClientConnはGoの初期のHTTP実装の遺物です。
// 低レベルで古く、現在のGoのHTTPスタックでは使用されていません。
// Go 1の前に削除すべきでした。
//
<<<<<<< HEAD
// Deprecated: 代わりにpackage net/httpのClientまたはTransportを使用してください。
=======
// Deprecated: Use Client or Transport in package [net/http] instead.
>>>>>>> upstream/release-branch.go1.22
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
<<<<<<< HEAD
// 廃止予定: 代わりにパッケージnet/httpのClientまたはTransportを使用してください。
=======
// Deprecated: Use the Client or Transport in package [net/http] instead.
>>>>>>> upstream/release-branch.go1.22
func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// NewProxyClientConn はGo言語の初期のHTTP実装の遺物です。
// これは低レベルで古く、現在のGoのHTTPスタックでは使用されていません。
// Go 1 より前に削除する必要があります。
//
<<<<<<< HEAD
// 廃止予定: 代わりに package net/http の Client または Transport を使用してください。
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// HijackはClientConnを切り離し、基礎となる接続と読み込み側のbufioを返します。
// また、左にデータが残っているかもしれない読み込み側のbufioも返します。
// HijackはユーザーまたはReadがkeep-aliveロジックの終了をシグナルした前に呼び出すことができます。
// ユーザーは、ReadまたはWriteが進行中の間にHijackを呼び出さないでください。
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)

// Close calls Hijack and then also closes the underlying connection.
// CloseはHijackを呼び出し、その後下層の接続も閉じます。
func (cc *ClientConn) Close() error

// Writeはリクエストを書き込みます。もしHTTP keep-aliveによって接続が閉じられた場合、ErrPersistEOFエラーが返されます。もしreq.Closeがtrueの場合、このリクエストの後にkeep-alive接続が論理的に閉じられ、対向サーバーに通知されます。ErrUnexpectedEOFは、リモートが基礎となるTCP接続を閉じたことを示しており、通常は正常な終了と見なされます。
=======
// Deprecated: Use the Client or Transport in package [net/http] instead.
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn

// Hijack detaches the [ClientConn] and returns the underlying connection as well
// as the read-side bufio which may have some left over data. Hijack may be
// called before the user or Read have signaled the end of the keep-alive
// logic. The user should not call Hijack while [ClientConn.Read] or ClientConn.Write is in progress.
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)

// Close calls [ClientConn.Hijack] and then also closes the underlying connection.
func (cc *ClientConn) Close() error

// Write writes a request. An [ErrPersistEOF] error is returned if the connection
// has been closed in an HTTP keep-alive sense. If req.Close equals true, the
// keep-alive connection is logically closed after this request and the opposing
// server is informed. An ErrUnexpectedEOF indicates the remote closed the
// underlying TCP connection, which is usually considered as graceful close.
>>>>>>> upstream/release-branch.go1.22
func (cc *ClientConn) Write(req *http.Request) error

// Pendingは、接続に送信された未応答のリクエストの数を返します。
func (cc *ClientConn) Pending() int

<<<<<<< HEAD
// Readはワイヤから次のレスポンスを読み込みます。有効なレスポンスはErrPersistEOFと一緒に返される場合があります。これはリモートがこれがサービスされる最後のリクエストであることを要求したことを意味します。ReadはWriteと同時に呼び出すことができますが、他のReadと同時には呼び出すことはできません。
=======
// Read reads the next response from the wire. A valid response might be
// returned together with an [ErrPersistEOF], which means that the remote
// requested that this be the last request serviced. Read can be called
// concurrently with [ClientConn.Write], but not with another Read.
>>>>>>> upstream/release-branch.go1.22
func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)

// Doはリクエストを書き込み、レスポンスを読み取る便利なメソッドです。
func (cc *ClientConn) Do(req *http.Request) (*http.Response, error)
