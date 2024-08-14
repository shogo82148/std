// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// httptraceパッケージは、HTTPクライアントリクエスト内のイベントをトレースするメカニズムを提供します。
package httptrace

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/textproto"
	"github.com/shogo82148/std/time"
)

// ContextClientTraceは与えられたコンテキストに関連付けられた [ClientTrace] を返します。関連付けられていない場合はnilを返します。
func ContextClientTrace(ctx context.Context) *ClientTrace

// WithClientTraceは提供された親コンテキストに基づいた新しいコンテキストを返します。返されたコンテキストを使用して行われるHTTPクライアントリクエストは、以前のフックに加えて、提供されたトレースフックを使用します。提供されたトレースで定義されたフックは最初に呼び出されます。
func WithClientTrace(ctx context.Context, trace *ClientTrace) context.Context

// ClientTraceは、送信中のHTTPリクエストのさまざまなステージで実行するためのフックのセットです。特定のフックはnilである場合があります。関数は複数のゴルーチンから同時に呼び出されることがあり、一部の関数はリクエストが完了または失敗した後に呼び出されることがあります。
// 現在、ClientTraceは単一のHTTPリクエストとレスポンスをトレースし、リダイレクトされたリクエストのシリーズを対象とするフックはありません。
// 詳細については、https://blog.golang.org/http-tracingを参照してください。
type ClientTrace struct {
	// GetConnは、接続が作成される前またはアイドルプールから取得される前に呼び出されます。hostPortは、対象またはプロキシの"ホスト:ポート"です。GetConnは、すでにアイドルキャッシュされた接続が利用可能である場合でも呼び出されます。
	GetConn func(hostPort string)

	// GotConnは、成功した接続後に呼び出されます。
	// 接続の取得に失敗した場合のフックはありません。代わりに、Transport.RoundTripからのエラーを使用してください。
	GotConn func(GotConnInfo)

	// PutIdleConnは接続がアイドルプールに返されたときに呼び出されます。
	// errがnilの場合、接続は正常にアイドルプールに返されました。
	// errが非nilの場合、なぜ正常に返されなかったかを説明します。
	// Transport.DisableKeepAlivesによって、接続再利用が無効化されている場合、PutIdleConnは呼び出されません。
	// PutIdleConnは呼び出し元のResponse.Body.Close呼び出しの前に呼び出されます。
	// HTTP/2では、このフックは現在使用されていません。
	PutIdleConn func(err error)

	// GotFirstResponseByteは、レスポンスヘッダーの最初のバイトが利用可能な場合に呼び出されます。
	GotFirstResponseByte func()

	// Got100Continue はサーバーが「100 Continue」の応答を返した場合に呼び出されます。
	Got100Continue func()

	// Got1xxResponseは、最終的な非1xxレスポンス前に返される1xx情報レスポンスヘッダーごとに呼び出されます。
	// Got1xxResponseは、「100 Continue」レスポンスに対しても、Got100Continueが定義されている場合でも呼び出されます。
	// エラーを返すと、そのエラー値でクライアントリクエストが中止されます。
	Got1xxResponse func(code int, header textproto.MIMEHeader) error

	// DNSStart はDNSの検索が始まった時に呼び出されます。
	DNSStart func(DNSStartInfo)

	// DNSDoneはDNSの検索が終了した時に呼び出されます。
	DNSDone func(DNSDoneInfo)

	// ConnectStart は新しい接続のダイヤルが開始された時に呼び出されます。
	// net.Dialer.DualStack (IPv6 "Happy Eyeballs") サポートが有効になっている場合、これは複数回呼び出されるかもしれません。
	ConnectStart func(network, addr string)

	// ConnectDone は新しい接続の Dial が完了すると呼び出されます。
	// 提供された err は接続が成功したかどうかを示します。
	// net.Dialer.DualStack ("Happy Eyeballs") サポートが有効な場合、これは複数回呼び出される可能性があります。
	ConnectDone func(network, addr string, err error)

	// TLSHandshakeStartはTLSハンドシェイクが開始されたときに呼び出されます。HTTPプロキシを介してHTTPSサイトに接続する場合、ハンドシェイクはプロキシによってCONNECTリクエストが処理された後に行われます。
	TLSHandshakeStart func()

	// TLSHandshakeDoneは、TLSハンドシェイクが成功した場合、またはハンドシェイクが失敗した場合に、成功したハンドシェイクの接続状態、またはハンドシェイクエラーのいずれかを受け取った後に呼び出される。
	TLSHandshakeDone func(tls.ConnectionState, error)

	// WroteHeaderField は、Transport が各リクエストヘッダーを書き込んだ後に呼び出されます。この呼び出し時点では、値はバッファリングされており、まだネットワークに書き込まれていない可能性があります。
	WroteHeaderField func(key string, value []string)

	// WroteHeaders は、Transport がすべてのリクエストヘッダを書き込んだ後に呼び出されます。
	WroteHeaders func()

	// Wait100Continue は、リクエストが "Expect: 100-continue" を指定し、
	// トランスポートがリクエストヘッダーを書き込みましたが、
	// リクエストボディを書き込む前にサーバーから "100 Continue" を待っている場合に呼び出されます。
	Wait100Continue func()

	// WroteRequestは、リクエストと任意のボディの書き込み結果が渡されたときに呼び出されます。再試行される場合には複数回呼び出されることがあります。
	WroteRequest func(WroteRequestInfo)
}

// WroteRequestInfoはWroteRequestフックに提供される情報を含んでいます。
type WroteRequestInfo struct {
	// Err はリクエストの書き込み中に遭遇したエラーです。
	Err error
}

// DNSStartInfoはDNSリクエストに関する情報を含んでいます。
type DNSStartInfo struct {
	Host string
}

// DNSDoneInfoはDNS検索の結果に関する情報を含んでいます。
type DNSDoneInfo struct {

	// AddrsにはDNSの検索で見つかったIPv4と/またはIPv6のアドレスが含まれます。
	// スライスの内容は変更しないでください。
	Addrs []net.IPAddr

	// ErrはDNSルックアップ中に発生したエラーです。
	Err error

	// Coalescedは、同時にDNSルックアップを行っていた別の呼び出し元とAddrsが共有されていたかどうかを示す。
	Coalesced bool
}

// GotConnInfoは [ClientTrace.GotConn] 関数の引数であり、
// 取得した接続に関する情報を含んでいます。
type GotConnInfo struct {

	// Connは取得された接続です。これはhttp.Transportによって所有されており、ClientTraceのユーザーは読み書きやクローズを行ってはいけません。
	Conn net.Conn

	// Reusedは、この接続が以前に別のHTTPリクエストで使用されたかどうかを示す。
	Reused bool

	// WasIdleはこのコネクションがアイドルプールから取得されたかどうかを示します。
	WasIdle bool

	// WasIdleがtrueの場合、IdleTimeは接続が前回アイドル状態だった時間を示します。
	IdleTime time.Duration
}
