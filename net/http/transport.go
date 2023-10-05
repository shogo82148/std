// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client implementation. See RFC 7230 through 7235.
//
// This is the low-level Transport implementation of RoundTripper.
// The high-level interface is in client.go.

package http

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// DefaultTransportはTransportのデフォルト実装であり、DefaultClientによって使用されます。
// 必要に応じてネットワーク接続を確立し、後続の呼び出しで再利用するためにキャッシュします。
// 環境変数HTTP_PROXY、HTTPS_PROXY、およびNO_PROXY（またはその小文字バージョン）によって指示されたように、HTTPプロキシを使用します。
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: defaultTransportDialContext(&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}),
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// DefaultMaxIdleConnsPerHostは、TransportのMaxIdleConnsPerHostのデフォルト値です。
const DefaultMaxIdleConnsPerHost = 2

// Transportは、HTTP、HTTPS、およびHTTPプロキシ（HTTPまたはHTTPS with CONNECTのいずれか）をサポートするRoundTripperの実装です。
//
// デフォルトでは、Transportは将来の再利用のために接続をキャッシュします。
// これにより、多くのホストにアクセスする場合に多数のオープンな接続が残る可能性があります。
// この動作は、TransportのCloseIdleConnectionsメソッドとMaxIdleConnsPerHostおよびDisableKeepAlivesフィールドを使用して管理できます。
//
// Transportは必要に応じて作成するのではなく、再利用する必要があります。
// Transportは、複数のgoroutineによる同時使用に対して安全です。
//
// Transportは、HTTPおよびHTTPSリクエストを行うための低レベルのプリミティブです。
// クッキーやリダイレクトなどの高レベルの機能については、Clientを参照してください。
//
// Transportは、HTTP URLではHTTP/1.1を、HTTPS URLではHTTP/1.1またはHTTP/2を使用します。
// これは、サーバーがHTTP/2をサポートしているかどうか、およびTransportの構成によって異なります。
// DefaultTransportはHTTP/2をサポートしています。
// Transportで明示的にHTTP/2を有効にするには、golang.org/x/net/http2を使用してConfigureTransportを呼び出します。
// HTTP/2についての詳細については、パッケージのドキュメントを参照してください。
//
// ステータスコードが1xx範囲にあるレスポンスは、自動的に処理されます（100 expect-continue）。
// ただし、HTTPステータスコード101（Switching Protocols）は、終端ステータスと見なされ、RoundTripによって返されます。
// 無視された1xxレスポンスを表示するには、httptraceトレースパッケージのClientTrace.Got1xxResponseを使用します。
//
// Transportは、ネットワークエラーに遭遇した場合にのみ、接続がすでに正常に使用されており、
// リクエストが冪等であり、ボディがないか、またはRequest.GetBodyが定義されている場合に、
// リクエストを再試行します。HTTPリクエストは、HTTPメソッドがGET、HEAD、OPTIONS、またはTRACEである場合、
// またはHeaderマップに「Idempotency-Key」または「X-Idempotency-Key」エントリが含まれている場合、冪等と見なされます。
// 冪等性キーの値がゼロ長のスライスの場合、リクエストは冪等と見なされますが、ヘッダーはワイヤーに送信されません。
type Transport struct {
	idleMu       sync.Mutex
	closeIdle    bool
	idleConn     map[connectMethodKey][]*persistConn
	idleConnWait map[connectMethodKey]wantConnQueue
	idleLRU      connLRU

	reqMu       sync.Mutex
	reqCanceler map[cancelKey]func(error)

	altMu    sync.Mutex
	altProto atomic.Value

	connsPerHostMu   sync.Mutex
	connsPerHost     map[connectMethodKey]int
	connsPerHostWait map[connectMethodKey]wantConnQueue

	// Proxyは、指定されたRequestに対するプロキシを返す関数を指定します。
	// 関数が非nilのエラーを返す場合、リクエストは提供されたエラーで中止されます。
	//
	// プロキシのタイプは、URLスキームによって決定されます。
	// "http"、"https"、および"socks5"がサポートされています。
	// スキームが空の場合、"http"が想定されます。
	//
<<<<<<< HEAD
	// If the proxy URL contains a userinfo subcomponent,
	// the proxy request will pass the username and password
	// in a Proxy-Authorization header.
	//
	// If Proxy is nil or returns a nil *URL, no proxy is used.
=======
	// Proxyがnilであるか、nilの*URLを返す場合、プロキシは使用されません。
>>>>>>> release-branch.go1.21
	Proxy func(*Request) (*url.URL, error)

	// OnProxyConnectResponseは、TransportがCONNECTリクエストのプロキシからHTTPレスポンスを受信したときに呼び出されます。
	// これは、200 OKレスポンスのチェックの前に呼び出されます。
	// エラーを返すと、リクエストはそのエラーで失敗します。
	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error

	// DialContextは、暗号化されていないTCP接続を作成するためのダイアル関数を指定します。
	// DialContextがnilである場合（および下記の非推奨のDialもnilである場合）、
	// トランスポートはnetパッケージを使用してダイアルします。
	//
	// DialContextは、RoundTripの呼び出しと並行して実行されます。
	// ダイアルを開始するRoundTrip呼び出しが、後のDialContextが完了する前に
	// 以前にダイアルされた接続を使用する場合があります。
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// Dialは、暗号化されていないTCP接続を作成するためのダイアル関数を指定します。
	//
	// Dialは、RoundTripの呼び出しと並行して実行されます。
	// 以前にダイアルされた接続が後でアイドル状態になる場合、
	// 後のDialが完了する前に、ダイアルを開始するRoundTrip呼び出しが以前にダイアルされた接続を使用する場合があります。
	//
	// Deprecated: 代わりにDialContextを使用してください。これにより、トランスポートはダイアルが不要になった直後にキャンセルできます。
	// 両方が設定されている場合、DialContextが優先されます。
	Dial func(network, addr string) (net.Conn, error)

	// DialTLSContextは、プロキシを使用しないHTTPSリクエストのためのTLS接続を作成するためのオプションのダイアル関数を指定します。
	//
	// DialTLSContextがnilである場合（および下記の非推奨のDialTLSもnilである場合）、
	// DialContextとTLSClientConfigが使用されます。
	//
	// DialTLSContextが設定されている場合、HTTPSリクエストに対してDialおよびDialContextフックは使用されず、
	// TLSClientConfigおよびTLSHandshakeTimeoutは無視されます。
	// 返されたnet.Connは、すでにTLSハンドシェイクを完了しているものと見なされます。
	DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// DialTLSは、プロキシを使用しないHTTPSリクエストのためのTLS接続を作成するためのオプションのダイアル関数を指定します。
	//
	// Deprecated: 代わりにDialTLSContextを使用してください。これにより、トランスポートはダイアルが不要になった直後にキャンセルできます。
	// 両方が設定されている場合、DialTLSContextが優先されます。
	DialTLS func(network, addr string) (net.Conn, error)

	// TLSClientConfigは、tls.Clientで使用するTLS構成を指定します。
	// nilの場合、デフォルトの構成が使用されます。
	// nil以外の場合、HTTP/2サポートがデフォルトで有効になっていない場合があります。
	TLSClientConfig *tls.Config

	// TLSHandshakeTimeoutは、TLSハンドシェイクを待機する最大時間を指定します。
	// ゼロの場合、タイムアウトはありません。
	TLSHandshakeTimeout time.Duration

	// DisableKeepAlivesがtrueの場合、HTTP keep-alivesが無効になり、
	// サーバーへの接続は単一のHTTPリクエストにのみ使用されます。
	//
	// これは、同様に名前が付けられたTCP keep-alivesとは無関係です。
	DisableKeepAlives bool

	// DisableCompressionがtrueの場合、Transportは、Requestに既存のAccept-Encoding値がない場合に、
	// "Accept-Encoding: gzip"リクエストヘッダーで圧縮を要求しません。
	// Transportが自動的にgzipを要求し、gzipされたレスポンスを受け取った場合、Response.Bodyで透過的にデコードされます。
	// ただし、ユーザーが明示的にgzipを要求した場合は、自動的に解凍されません。
	DisableCompression bool

	// MaxIdleConnsは、すべてのホストをまたいでアイドル（keep-alive）接続の最大数を制御します。
	// ゼロの場合、制限はありません。
	MaxIdleConns int

	// MaxIdleConnsPerHostがゼロでない場合、ホストごとに保持する最大アイドル（keep-alive）接続数を制御します。
	// ゼロの場合、DefaultMaxIdleConnsPerHostが使用されます。
	MaxIdleConnsPerHost int

	// MaxConnsPerHostは、ダイアル、アクティブ、およびアイドル状態の接続を含む、ホストごとの総接続数をオプションで制限します。
	// 制限を超えると、ダイアルはブロックされます。
	//
	// ゼロは制限がないことを意味します。
	MaxConnsPerHost int

	// IdleConnTimeoutは、アイドル（keep-alive）接続が自己クローズする前にアイドル状態になる最大時間です。
	// ゼロは制限がないことを意味します。
	IdleConnTimeout time.Duration

	// ResponseHeaderTimeoutがゼロでない場合、リクエスト（ボディがある場合はそれも含む）を完全に書き込んだ後、
	// サーバーのレスポンスヘッダーを待機する時間を指定します。
	// この時間には、レスポンスボディを読み取る時間は含まれません。
	ResponseHeaderTimeout time.Duration

	// ExpectContinueTimeoutがゼロでない場合、リクエストに"Expect: 100-continue"ヘッダーがある場合、
	// リクエストヘッダーを完全に書き込んだ後、サーバーの最初のレスポンスヘッダーを待機する時間を指定します。
	// ゼロはタイムアウトがないことを意味し、サーバーの承認を待たずに、すぐにボディを送信します。
	// この時間には、リクエストヘッダーを送信する時間は含まれません。
	ExpectContinueTimeout time.Duration

	// TLSNextProtoは、TLS ALPNプロトコルネゴシエーション後にTransportが代替プロトコル（HTTP/2など）に切り替える方法を指定します。
	// Transportがプロトコル名が空でないTLS接続をダイアルし、TLSNextProtoにそのキーのマップエントリが含まれている場合（"h2"など）、
	// リクエストの権限（"example.com"または"example.com:1234"など）とTLS接続でfuncが呼び出されます。
	// この関数は、その後リクエストを処理するRoundTripperを返さなければなりません。
	// TLSNextProtoがnilでない場合、HTTP/2サポートは自動的に有効になりません。
	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

	// ProxyConnectHeaderは、CONNECTリクエスト中にプロキシに送信するヘッダーをオプションで指定します。
	// ヘッダーを動的に設定するには、GetProxyConnectHeaderを参照してください。
	ProxyConnectHeader Header

	// GetProxyConnectHeaderは、ip:portターゲットへのCONNECTリクエスト中にproxyURLに送信するヘッダーを返すためのオプションの関数を指定します。
	// エラーを返すと、TransportのRoundTripはそのエラーで失敗します。
	// ヘッダーを追加しない場合は、(nil, nil)を返すことができます。
	// GetProxyConnectHeaderが非nilの場合、ProxyConnectHeaderは無視されます。
	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)

	// MaxResponseHeaderBytesは、サーバーのレスポンスヘッダーに許可されるレスポンスバイト数の制限を指定します。
	//
	// ゼロは、デフォルトの制限を使用することを意味します。
	MaxResponseHeaderBytes int64

	// WriteBufferSizeは、トランスポートへの書き込み時に使用される書き込みバッファのサイズを指定します。
	// ゼロの場合、デフォルト値（現在は4KB）が使用されます。
	WriteBufferSize int

	// ReadBufferSizeは、トランスポートから読み取るときに使用される読み取りバッファのサイズを指定します。
	// ゼロの場合、デフォルト値（現在は4KB）が使用されます。
	ReadBufferSize int

	// nextProtoOnce guards initialization of TLSNextProto and
	// h2transport (via onceSetNextProtoDefaults)
	nextProtoOnce      sync.Once
	h2transport        h2Transport
	tlsNextProtoWasNil bool

	// ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
	// Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
	// By default, use of any those fields conservatively disables HTTP/2.
	// To use a custom dialer or TLS config and still attempt HTTP/2
	// upgrades, set this to true.
	ForceAttemptHTTP2 bool
}

// Cloneは、tのエクスポートされたフィールドのディープコピーを返します。
func (t *Transport) Clone() *Transport

// ProxyFromEnvironmentは、環境変数HTTP_PROXY、HTTPS_PROXY、およびNO_PROXY（またはそれらの小文字バージョン）によって示されるように、
// 指定されたリクエストに使用するプロキシのURLを返します。
// リクエストは、NO_PROXYによって除外されていない限り、スキームに一致する環境変数からプロキシを使用します。
//
// 環境値は、完全なURLまたは"host[:port]"のいずれかである場合があります。この場合、"http"スキームが想定されます。
// スキーム"http"、"https"、および"socks5"がサポートされています。
// 値が異なる形式の場合は、エラーが返されます。
//
// 環境変数でプロキシが定義されていない場合、またはNO_PROXYによって指定されたリクエストにプロキシを使用しない場合、
// nilのURLとnilのエラーが返されます。
//
// 特別な場合として、req.URL.Hostが"localhost"（ポート番号ありまたはなし）の場合、nilのURLとnilのエラーが返されます。
func ProxyFromEnvironment(req *Request) (*url.URL, error)

// ProxyURLは、常に同じURLを返すプロキシ関数（Transportで使用するため）を返します。
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)

// ErrSkipAltProtocolは、Transport.RegisterProtocolによって定義されたセンチネルエラー値です。
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")

// RegisterProtocolは、新しいプロトコルをスキームとともに登録します。
// Transportは、指定されたスキームを使用してリクエストをrtに渡します。
// HTTPリクエストのセマンティクスをシミュレートする責任は、rtにあります。
//
// RegisterProtocolは、他のパッケージが"ftp"や"file"などのプロトコルスキームの実装を提供するために使用できます。
//
// rt.RoundTripがErrSkipAltProtocolを返す場合、Transportは、
// 登録されたプロトコルのように扱わずに、その1つのリクエストに対して自身でRoundTripを処理します。
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)

// CloseIdleConnectionsは、以前のリクエストから接続されていたが、現在はアイドル状態になっている"keep-alive"状態の接続を閉じます。
// 現在使用中の接続は中断しません。
func (t *Transport) CloseIdleConnections()

// CancelRequestは、その接続を閉じることにより、進行中のリクエストをキャンセルします。
// CancelRequestは、RoundTripが返された後にのみ呼び出す必要があります。
//
// Deprecated: 代わりに、キャンセル可能なコンテキストを持つリクエストを作成するためにRequest.WithContextを使用してください。
// CancelRequestは、HTTP/2リクエストをキャンセルできません。
func (t *Transport) CancelRequest(req *Request)

var _ io.ReaderFrom = (*persistConnWriter)(nil)
