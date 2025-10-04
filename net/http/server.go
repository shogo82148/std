// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP server. See RFC 7230 through 7235.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net"

	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// HTTPサーバーで使用されるエラー。
var (
	// ErrBodyNotAllowedは、HTTPメソッドまたはレスポンスコードがボディを許可しない場合に、ResponseWriter.Write呼び出しによって返されます。
	ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")

	// ErrHijackedは、Hijackerインターフェースを使用して基礎となる接続がハイジャックされた場合に、ResponseWriter.Write呼び出しによって返されます。
	// ハイジャックされた接続でのゼロバイト書き込みは、他の副作用なしにErrHijackedを返します。
	ErrHijacked = errors.New("http: connection has been hijacked")

	// ErrContentLengthは、Handlerが宣言されたサイズを持つContent-Lengthレスポンスヘッダーを設定し、宣言されたバイト数よりも多くのバイトを書き込もうとした場合に、ResponseWriter.Write呼び出しによって返されます。
	ErrContentLength = errors.New("http: wrote more than the declared Content-Length")

	// Deprecated: ErrWriteAfterFlushは、net/httpパッケージの何も返さなくなったため、もはや返されません。
	// 呼び出し元は、この変数に対してエラーを比較するべきではありません。
	ErrWriteAfterFlush = errors.New("unused")
)

// Handlerは、HTTPリクエストに応答します。
//
// [Handler.ServeHTTP] は、応答ヘッダーとデータを [ResponseWriter] に書き込んでから返す必要があります。
// 返すことで、リクエストが完了したことを示します。
// ServeHTTPの呼び出しの完了後または同時に、 [ResponseWriter] を使用するか、[Request.Body]から読み取ることはできません。
//
// HTTPクライアントソフトウェア、HTTPプロトコルバージョン、およびクライアントとGoサーバーの間の中間者によっては、
// [ResponseWriter]に書き込んだ後に [Request.Body] から読み取ることができない場合があります。
// 注意深いハンドラーは、最初に [Request.Body] を読み取り、その後に応答する必要があります。
//
// ボディを読み取る以外の場合、ハンドラーは提供されたRequestを変更してはいけません。
//
// ServeHTTPがパニックを起こすと、サーバー（ServeHTTPの呼び出し元）は、パニックの影響がアクティブなリクエストに限定されたものであると仮定します。
// サーバーはパニックを回復し、サーバーエラーログにスタックトレースを記録し、ネットワーク接続を閉じるか、HTTP/2 RST_STREAMを送信します。
// クライアントが中断された応答を見るが、サーバーがエラーをログに記録しないように、 [ErrAbortHandler] の値でパニックを発生させることで、ハンドラーを中止できます。
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// ResponseWriterインターフェースは、HTTPハンドラーがHTTPレスポンスを構築するために使用されます。
//
// [Handler.ServeHTTP] が返された後に、ResponseWriter を使用することはできません。
type ResponseWriter interface {
	Header() Header

	Write([]byte) (int, error)

	WriteHeader(statusCode int)
}

// Flusherインターフェースは、HTTPハンドラーがバッファリングされたデータをクライアントにフラッシュすることを許可するResponseWriterによって実装されます。
//
// デフォルトのHTTP/1.xおよびHTTP/2 [ResponseWriter] 実装は [Flusher] をサポートしていますが、ResponseWriterラッパーはサポートしていない場合があります。
// ハンドラーは常にランタイムでこの機能をテストする必要があります。
//
// FlushをサポートするResponseWriterであっても、クライアントがHTTPプロキシを介して接続されている場合、
// バッファリングされたデータがレスポンスが完了するまでクライアントに到達しない場合があります。
type Flusher interface {
	Flush()
}

// Hijackerインターフェースは、HTTPハンドラーが接続を引き継ぐことを許可するResponseWriterによって実装されます。
//
// HTTP/1.x接続のデフォルト [ResponseWriter] はHijackerをサポートしていますが、HTTP/2接続は意図的にサポートしていません。
// ResponseWriterラッパーもHijackerをサポートしていない場合があります。
// ハンドラーは常にランタイムでこの機能をテストする必要があります。
type Hijacker interface {
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}

// CloseNotifierインターフェースは、基礎となる接続が切断されたときに検出できるResponseWriterによって実装されます。
//
// このメカニズムは、レスポンスが準備される前にクライアントが切断された場合、サーバー上の長時間の操作をキャンセルするために使用できます。
//
// Deprecated: CloseNotifierインターフェースは、Goのコンテキストパッケージより前に実装されました。
// 新しいコードでは、[Request.Context] を使用する必要があります。
type CloseNotifier interface {
	CloseNotify() <-chan bool
}

var (
	// ServerContextKeyは、コンテキストキーです。
	// HTTPハンドラーでContext.Valueと一緒に使用して、ハンドラーを開始したサーバーにアクセスできます。
	// 関連する値の型は*Serverです。
	ServerContextKey = &contextKey{"http-server"}

	// LocalAddrContextKeyは、コンテキストキーです。
	// HTTPハンドラーでContext.Valueと一緒に使用して、接続が到着したローカルアドレスにアクセスできます。
	// 関連する値の型はnet.Addrです。
	LocalAddrContextKey = &contextKey{"local-addr"}
)

// TrailerPrefixは、[ResponseWriter.Header] マップのキーに対するマジックプレフィックスで、
// 存在する場合は、マップエントリが実際にはレスポンストレーラーであることを示します。
// プレフィックスは、ServeHTTP呼び出しが終了し、値がトレーラーに送信された後に削除されます。
//
// このメカニズムは、ヘッダーが書き込まれる前には不明なトレーラーにのみ使用することができます。
// トレーラーのセットが固定されている場合、またはヘッダーが書き込まれる前に既知の場合、通常のGoトレーラーメカニズムが推奨されます。
//
//	https://pkg.go.dev/net/http#ResponseWriter
//	https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
const TrailerPrefix = "Trailer:"

// DefaultMaxHeaderBytesは、HTTPリクエストのヘッダーの許容される最大サイズです。
// これは、[Server.MaxHeaderBytes] を設定することで上書きできます。
const DefaultMaxHeaderBytes = 1 << 20

// TimeFormatは、HTTPヘッダーで時間を生成するときに使用する時間形式です。
// [time.RFC1123] のようですが、タイムゾーンとしてGMTがハードコードされています。
// フォーマットされる時間はUTCである必要があります。
//
// この時間形式を解析するには、[ParseTime] を参照してください。
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

var _ closeWriter = (*net.TCPConn)(nil)

// ErrAbortHandlerは、ハンドラーを中止するためのセンチネルパニック値です。
// ServeHTTPからのパニックはすべて、クライアントへの応答を中止しますが、
// ErrAbortHandlerでパニックすると、サーバーのエラーログにスタックトレースを記録しないようにすることができます。
var ErrAbortHandler = errors.New("net/http: abort Handler")

// HandlerFunc型は、HTTPハンドラーとして通常の関数を使用できるようにするためのアダプタです。
// fが適切なシグネチャを持つ関数である場合、HandlerFunc(f)はfを呼び出す [Handler] です。
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTPは、f(w, r) を呼び出します。
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

// Errorは、指定されたエラーメッセージとHTTPコードでリクエストに応答します。
// それ以外に、リクエストを終了しません。呼び出し元は、wに対してさらに書き込みが行われないようにする必要があります。
// エラーメッセージはプレーンテキストである必要があります。
//
// ErrorはContent-Lengthヘッダーを削除し、
// Content-Typeを"text/plain; charset=utf-8"に設定し、
// X-Content-Type-Optionsを"nosniff"に設定します。
// これにより、呼び出し元が成功した出力を期待して設定していた場合でも、
// エラーメッセージ用にヘッダーが適切に設定されます。
func Error(w ResponseWriter, error string, code int)

// NotFoundは、HTTP 404 not foundエラーでリクエストに応答します。
func NotFound(w ResponseWriter, r *Request)

// NotFoundHandlerは、各リクエストに「404ページが見つかりません」という応答を返す単純なリクエストハンドラーを返します。
func NotFoundHandler() Handler

// StripPrefixは、リクエストURLのPath（および設定されている場合はRawPath）から指定された接頭辞を削除し、ハンドラーhを呼び出してHTTPリクエストを処理するハンドラーを返します。
// prefixで始まらないパスのリクエストには、HTTP 404 not foundエラーで応答します。
// 接頭辞は完全に一致する必要があります。リクエストの接頭辞にエスケープされた文字が含まれている場合、応答もHTTP 404 not foundエラーになります。
func StripPrefix(prefix string, h Handler) Handler

<<<<<<< HEAD
// Redirectは、リクエストに対してurlにリダイレクトする応答を返します。
// urlは、リクエストパスに対する相対パスである場合があります。
=======
// Redirect replies to the request with a redirect to url,
// which may be a path relative to the request path.
// Any non-ASCII characters in url will be percent-encoded,
// but existing percent encodings will not be changed.
>>>>>>> upstream/release-branch.go1.25
//
// 提供されたコードは通常、[StatusMovedPermanently]、[StatusFound]、または [StatusSeeOther] の3xx範囲にあります。
//
// Content-Typeヘッダーが設定されていない場合、[Redirect] はそれを"text/html; charset=utf-8"に設定し、小さなHTML本文を書き込みます。
// Content-Typeヘッダーを任意の値、nilを含む任意の値に設定すると、その動作が無効になります。
func Redirect(w ResponseWriter, r *Request, url string, code int)

// RedirectHandlerは、受信した各リクエストを、指定されたステータスコードを使用して、指定されたURLにリダイレクトするリクエストハンドラーを返します。
//
// 提供されたコードは通常、[StatusMovedPermanently]、[StatusFound]、または [StatusSeeOther] の3xx範囲にあります。
func RedirectHandler(url string, code int) Handler

// ServeMuxは、HTTPリクエストマルチプレクサーです。
// それは、登録されたパターンのリストに対して、各受信リクエストのURLを一致させ、
// URLに最も近いパターンのハンドラーを呼び出します。
//
// # Patterns
//
// パターンは、リクエストのメソッド、ホスト、およびパスに一致することができます。
// いくつかの例：
//
//   - "/index.html" は、任意のホストとメソッドに対してパス "/index.html" に一致します。
//   - "GET /static/" は、"/static/" で始まるGETリクエストに一致します。
//   - "example.com/" は、ホスト "example.com" への任意のリクエストに一致します。
//   - "example.com/{$}" は、ホストが "example.com" でパスが "/" のリクエストに一致します。
//   - "/b/{bucket}/o/{objectname...}" は、最初のセグメントが "b" で、3番目のセグメントが "o" のパスに一致します。
//     "bucket" は2番目のセグメントを示し、"objectname" はパスの残りを示します。
//
// 一般的に、パターンは以下のようになります。
//
//	[METHOD][HOST]/[PATH]
//
// すべての3つの部分はオプションです。"/" は有効なパターンです。
// METHODが存在する場合、後に単一のスペースもしくはタブが続く必要があります。
//
// パターンのリテラル（ワイルドカードでない）部分は、リクエストの対応する部分と大文字小文字を区別して一致します。
//
// メソッドのないパターンはすべてのメソッドに一致します。メソッドがGETの場合、GETとHEADの両方のリクエストに一致します。
// それ以外の場合、メソッドは完全に一致する必要があります。
//
// ホストのないパターンはすべてのホストに一致します。ホストがあるパターンは、そのホストのURLにのみ一致します。
//
// パスには、{NAME}または{NAME...}のワイルドカードセグメントを含めることができます。
// 例えば、"/b/{bucket}/o/{objectname...}" です。
// ワイルドカード名は有効なGo識別子でなければなりません。
// ワイルドカードは完全なパスセグメントでなければなりません。つまり、スラッシュに続き、スラッシュまたは文字列の終わりに続く必要があります。
// 例えば、"/b_{bucket}" は有効なパターンではありません。
//
// 通常、ワイルドカードはリクエストURLの次のリテラルスラッシュ（%2Fではない）で終わる、単一のパスセグメントにのみ一致します。
// ただし、"..."が存在する場合、ワイルドカードは、スラッシュを含むURLパスの残り全体に一致します。
// （したがって、"..."ワイルドカードはパターンの末尾以外に現れることはできません。）
// ワイルドカードの一致は、ワイルドカードの名前を指定して [Request.PathValue] を呼び出すことで取得できます。
// パスの末尾にスラッシュがある場合、匿名の "..." ワイルドカードとして機能します。
//
// 特別なワイルドカード {$} は、URLの末尾にのみ一致します。
// 例えば、パターン "/{$}" はパス "/" にのみ一致し、パターン "/" はすべてのパスに一致します。
//
// 一致には、パターンパスと受信リクエストパスの両方が、セグメントごとにエスケープされていない状態で使用されます。
// したがって、パス "/a%2Fb/100%25" は、2つのセグメント "a/b" と "100%" を持つと見なされます。
// パターン "/a%2fb/" はそれに一致しますが、パターン "/a/b/" は一致しません。
//
// # Precedence
//
// 2つ以上のパターンがリクエストに一致する場合、最も具体的なパターンが優先されます。
// パターンP1がP2よりも具体的であるとは、P1がP2のリクエストの厳密なサブセットに一致する場合を指します。
// つまり、P2がP1のすべてのリクエストに一致し、それ以上に一致する場合です。
// もし、どちらも具体的でない場合、そのパターンは競合します。
// このルールには、後方互換性のための1つの例外があります：
// 2つのパターンがそれ以外の場合に競合し、1つはホストを持ち、もう1つは持っていない場合、
// ホストを持つパターンが優先されます。
// [ServeMux.Handle]または[ServeMux.HandleFunc]に渡されるパターンが
// すでに登録されている他のパターンと競合する場合、それらの関数はパニックを引き起こします。
//
// 一般的なルールの例として、"/images/thumbnails/"は"/images/"よりも具体的であり、両方とも登録できます。
// 前者は"/images/thumbnails/"で始まるパスに一致し、後者は"/images/"サブツリー内の他のパスに一致します。
//
// 別の例として、パターン"GET /"と"/index.html"を考えてみてください。
// 両方が"/index.html"のGETリクエストに一致しますが、前者のパターンはすべての他のGETおよびHEADリクエストに一致し、
// 後者のパターンは異なるメソッドを使用する"/index.html"のすべてのリクエストに一致します。
// パターンは競合します。
//
// # Trailing-slash redirection
//
// 末尾スラッシュまたは "..." ワイルドカードを使用して登録されたサブツリーのハンドラを持つ [ServeMux] を考えてみてください。
// ServeMuxが末尾スラッシュのないサブツリールートのリクエストを受信した場合、
// 末尾スラッシュを追加してリクエストをリダイレクトします。
// この動作は、末尾スラッシュまたは "..." ワイルドカードを使用しないパスの別個の登録によって上書きできます。
// 例えば、"/images/"を登録すると、ServeMuxは"/images"のリクエストを"/images/"にリダイレクトします。
// "/images"が別途登録されていない限りです。
//
// # Request sanitizing
//
<<<<<<< HEAD
// ServeMuxは、URLリクエストパスとHostヘッダーをサニタイズし、ポート番号を削除し、.または..セグメントまたは重複したスラッシュを含むリクエストを同等のクリーンなURLにリダイレクトします。
=======
// ServeMux also takes care of sanitizing the URL request path and the Host
// header, stripping the port number and redirecting any request containing . or
// .. segments or repeated slashes to an equivalent, cleaner URL.
// Escaped path elements such as "%2e" for "." and "%2f" for "/" are preserved
// and aren't considered separators for request routing.
>>>>>>> upstream/release-branch.go1.25
//
// # Compatibility
//
// ServeMuxのパターン構文と一致動作は、Go 1.22で大幅に変更されました。
// 古い動作を復元するには、GODEBUG環境変数を "httpmuxgo121=1" に設定します。
// この設定は、プログラムの起動時に1回だけ読み取られます。実行中の変更は無視されます。
//
// 互換性のない変更には以下が含まれます。
//   - ワイルドカードは1.21では通常のリテラルパスセグメントでした。
//     例えば、パターン "/{x}" は1.21ではそのパスのみに一致しますが、1.22では1つのセグメントのパスに一致します。
//   - 1.21では、既存のパターンと競合しない限り、パターンは拒否されませんでした。
//     1.22では、構文的に無効なパターンは [ServeMux.Handle] および [ServeMux.HandleFunc] でパニックを引き起こします。
//     例えば、1.21では、パターン "/{" と "/a{x}" はそれ自身に一致しますが、1.22では無効であり、登録時にパニックを引き起こします。
//   - 1.22では、パターンの各セグメントがエスケープ解除されますが、1.21ではそうではありませんでした。
//     例えば、1.22ではパターン "/%61" はパス "/a" ("%61"は "a"のURLエスケープシーケンス) に一致しますが、
//     1.21ではパス "/%2561" のみに一致します（"%25"はパーセント記号のエスケープです）。
//   - パターンをパスに一致させる場合、1.22ではパスの各セグメントがエスケープ解除されますが、1.21ではパス全体がエスケープ解除されます。
//     この変更は、スラッシュに隣接する%2Fエスケープを持つパスがどのように処理されるかに影響します。
//     詳細については、https://go.dev/issue/21955 を参照してください。
type ServeMux struct {
	mu     sync.RWMutex
	tree   routingNode
	index  routingIndex
	mux121 serveMux121
}

// NewServeMuxは、新しい [ServeMux] を割り当てて返します。
func NewServeMux() *ServeMux

// DefaultServeMuxは、[Serve] によって使用されるデフォルトの [ServeMux] です。
var DefaultServeMux = &defaultServeMux

// Handlerは、r.Method、r.Host、およびr.URL.Pathを参照して、
// 指定されたリクエストに使用するハンドラーを返します。
// 常にnilでないハンドラーを返します。
// パスが正規形式でない場合、ハンドラーは正規パスにリダイレクトする内部生成ハンドラーになります。
// ホストにポートが含まれている場合、ハンドラーの一致時には無視されます。
//
// CONNECTリクエストでは、パスとホストは変更されずに使用されます。
//
// Handlerは、リクエストに一致する登録済みのパターン、または内部で生成されたリダイレクトの場合はリダイレクトをたどった後に一致するパスを返します。
//
<<<<<<< HEAD
// リクエストに適用される登録済みハンドラーがない場合、
// Handlerは「ページが見つかりません」というハンドラーと空のパターンを返します。
=======
// If there is no registered handler that applies to the request,
// Handler returns a “page not found” or “method not supported”
// handler and an empty pattern.
//
// Handler does not modify its argument. In particular, it does not
// populate named path wildcards, so r.PathValue will always return
// the empty string.
>>>>>>> upstream/release-branch.go1.25
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)

// ServeHTTPは、リクエストURLに最も近いパターンを持つハンドラにリクエストをディスパッチします。
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

<<<<<<< HEAD
// Handleは、指定されたパターンのハンドラを登録します。
// 登録済みのパターンと競合する場合、Handleはパニックを発生させます。
func (mux *ServeMux) Handle(pattern string, handler Handler)

// HandleFuncは、指定されたパターンのハンドラ関数を登録します。
// 登録済みのパターンと競合する場合、HandleFuncはパニックを発生させます。
=======
// Handle registers the handler for the given pattern.
// If the given pattern conflicts with one that is already registered, Handle
// panics.
func (mux *ServeMux) Handle(pattern string, handler Handler)

// HandleFunc registers the handler function for the given pattern.
// If the given pattern conflicts with one that is already registered, HandleFunc
// panics.
>>>>>>> upstream/release-branch.go1.25
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Handleは、[DefaultServeMux]に指定されたパターンのハンドラを登録します。
// [ServeMux]のドキュメントには、パターンの一致方法が説明されています。
func Handle(pattern string, handler Handler)

// HandleFuncは、[DefaultServeMux]に指定されたパターンのハンドラ関数を登録します。
// [ServeMux]のドキュメントには、パターンの一致方法が説明されています。
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Serveは、リスナーlに対して着信HTTP接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはリクエストを読み取り、その後handlerを呼び出して応答します。
//
// ハンドラは通常nilであり、その場合は [DefaultServeMux] が使用されます。
//
// TLS Config.NextProtosで"h2"が設定された [*tls.Conn] 接続を返すリスナーがある場合、HTTP/2サポートが有効になります。
//
// Serveは常にnilでないエラーを返します。
func Serve(l net.Listener, handler Handler) error

// ServeTLSは、リスナーlに対して着信HTTPS接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはリクエストを読み取り、その後handlerを呼び出して応答します。
//
// ハンドラは通常nilであり、その場合は [DefaultServeMux] が使用されます。
//
// さらに、サーバーの証明書と対応する秘密鍵を含むファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
//
// ServeTLSは常にnilでないエラーを返します。
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error

// Serverは、HTTPサーバーを実行するためのパラメータを定義します。
// Serverのゼロ値は有効な構成です。
type Server struct {
	// Addrは、サーバーがリッスンするためのTCPアドレスをオプションで指定します。
	// 形式は「host:port」です。空の場合、「:http」（ポート80）が使用されます。
	// サービス名はRFC 6335で定義され、IANAによって割り当てられます。
	// アドレスの形式の詳細については、net.Dialを参照してください。
	Addr string

	Handler Handler

	// DisableGeneralOptionsHandlerがtrueの場合、"OPTIONS *"リクエストをHandlerに渡します。
	// それ以外の場合、200 OKとContent-Length: 0で応答します。
	DisableGeneralOptionsHandler bool

	// TLSConfigは、ServeTLSとListenAndServeTLSで使用するためのTLS構成をオプションで提供します。
	// この値はServeTLSとListenAndServeTLSによってクローンされるため、tls.Config.SetSessionTicketKeysなどのメソッドを使用して構成を変更することはできません。
	// SetSessionTicketKeysを使用するには、TLSリスナーを使用してServer.Serveを使用します。
	TLSConfig *tls.Config

	// ReadTimeoutは、ボディを含むリクエスト全体を読み取るための最大時間です。
	// ゼロまたは負の値はタイムアウトがないことを意味します。
	//
	// ReadTimeoutは、ハンドラが各リクエストボディの許容可能な締め切りまたはアップロードレートを決定することを許可しないため、
	// 大抵のユーザーはReadHeaderTimeoutを使用することを好むでしょう。
	// 両方を使用することもできます。
	ReadTimeout time.Duration

	// ReadHeaderTimeoutは、リクエストヘッダを読み取るために許される時間です。
	// ヘッダを読み取った後、接続の読み取りデッドラインはリセットされ、
	// ハンドラは本文にとって何が遅すぎると考えられるかを決定できます。
	// もしゼロならば、ReadTimeoutの値が使用されます。
	// もし負の値、またはゼロでReadTimeoutがゼロまたは負の値の場合、タイムアウトはありません。
	ReadHeaderTimeout time.Duration

	// WriteTimeoutは、レスポンスの書き込みがタイムアウトする前の最大時間です。
	// 新しいリクエストヘッダーが読み取られるたびにリセットされます。
	// ReadTimeoutと同様に、ハンドラがリクエストごとに決定を下すことを許可しません。
	// ゼロまたは負の値はタイムアウトがないことを意味します。
	WriteTimeout time.Duration

	// IdleTimeoutは、keep-alivesが有効な場合に次のリクエストを待つための最大時間です。
	// もしゼロならば、ReadTimeoutの値が使用されます。
	// もし負の値、またはゼロでReadTimeoutがゼロまたは負の値の場合、タイムアウトはありません。
	IdleTimeout time.Duration

	// MaxHeaderBytesは、リクエストヘッダーのキーと値、およびリクエストラインを解析するためにサーバーが読み取ることができる最大バイト数を制御します。
	// リクエストボディのサイズには影響しません。
	// ゼロの場合、DefaultMaxHeaderBytesが使用されます。
	MaxHeaderBytes int

	// TLSNextProtoは、ALPNプロトコルアップグレードが発生した場合に提供されたTLS接続の所有権を引き継ぐための関数をオプションで指定します。
	// マップキーはネゴシエートされたプロトコル名です。
	// Handler引数はHTTPリクエストを処理するために使用され、RequestのTLSとRemoteAddrを初期化します（設定されていない場合）。
	// 関数が返されると、接続は自動的に閉じられます。
	// TLSNextProtoがnilでない場合、HTTP/2サポートは自動的に有効になりません。
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	// ConnStateは、クライアント接続の状態が変化したときに呼び出されるオプションのコールバック関数を指定します。
	// 詳細については、ConnState型と関連する定数を参照してください。
	ConnState func(net.Conn, ConnState)

	// ErrorLogは、接続の受け入れ時のエラー、ハンドラの予期しない動作、および
	// FileSystemの基礎となるエラーに対するオプションのロガーを指定します。
	// nilの場合、ログはlogパッケージの標準ロガーを使用して行われます。
	ErrorLog *log.Logger

	// BaseContextは、このサーバーの着信リクエストのベースコンテキストを返すオプションの関数を指定します。
	// 提供されたListenerは、リクエストを開始する特定のリスナーです。
	// BaseContextがnilの場合、デフォルトはcontext.Background()です。
	// nilでない場合、非nilのコンテキストを返す必要があります。
	BaseContext func(net.Listener) context.Context

	// ConnContextは、新しい接続cに使用されるコンテキストを変更するオプションの関数を指定します。
	// 提供されたctxはBaseContextから派生し、ServerContextKeyの値を持ちます。
	ConnContext func(ctx context.Context, c net.Conn) context.Context

	// HTTP2 configures HTTP/2 connections.
	//
	// This field does not yet have any effect.
	// See https://go.dev/issue/67813.
	HTTP2 *HTTP2Config

	// Protocols is the set of protocols accepted by the server.
	//
	// If Protocols includes UnencryptedHTTP2, the server will accept
	// unencrypted HTTP/2 connections. The server can serve both
	// HTTP/1 and unencrypted HTTP/2 on the same address and port.
	//
	// If Protocols is nil, the default is usually HTTP/1 and HTTP/2.
	// If TLSNextProto is non-nil and does not contain an "h2" entry,
	// the default is HTTP/1 only.
	Protocols *Protocols

	inShutdown atomic.Bool

	disableKeepAlives atomic.Bool
	nextProtoOnce     sync.Once
	nextProtoErr      error

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	onShutdown []func()

	listenerGroup sync.WaitGroup
}

// Closeは、すべてのアクティブなnet.Listenerと、[StateNew]、[StateActive]、または [StateIdle] の状態にある接続をすぐに閉じます。
// 優雅なシャットダウンには、[Server.Shutdown] を使用してください。
//
// Closeは、WebSocketsなどのハイジャックされた接続を閉じようとはせず（そしてそれらについては何も知りません）、試みません。
//
// Closeは、[Server] の基礎となるListener(s)を閉じる際に返される任意のエラーを返します。
func (srv *Server) Close() error

// Shutdownは、アクティブな接続を中断することなく、サーバーを正常にシャットダウンします。
// Shutdownは、まずすべてのオープンなリスナーを閉じ、次にすべてのアイドル状態の接続を閉じ、
// そして接続がアイドル状態に戻ってから無期限に待機し、その後シャットダウンします。
// 提供されたコンテキストがシャットダウンが完了する前に期限切れになった場合、
// Shutdownはコンテキストのエラーを返します。それ以外の場合は、[Server] の基礎となる
// Listener(s)を閉じる際に返される任意のエラーを返します。
//
<<<<<<< HEAD
// Shutdownが呼び出されると、[Serve]、[ListenAndServe]、および
// [ListenAndServeTLS] はすぐに [ErrServerClosed] を返します。プログラムが
// 終了せず、代わりにShutdownが返るのを待つことを確認してください。
=======
// When Shutdown is called, [Serve], [ServeTLS], [ListenAndServe], and
// [ListenAndServeTLS] immediately return [ErrServerClosed]. Make sure the
// program doesn't exit and waits instead for Shutdown to return.
>>>>>>> upstream/release-branch.go1.25
//
// Shutdownは、WebSocketsのようなハイジャックされた接続を閉じたり、それらを待つことは試みません。
// Shutdownの呼び出し元は、長時間稼働する接続に対してシャットダウンを別途通知し、
// 必要に応じてそれらが閉じるのを待つべきです。シャットダウン通知関数を登録する方法については、
// [Server.RegisterOnShutdown] を参照してください。
//
// Shutdownを呼び出した後、サーバーを再利用することはできません。Serveなどのメソッドを呼び出すと、ErrServerClosedが返されます。
func (srv *Server) Shutdown(ctx context.Context) error

// RegisterOnShutdownは、[Server.Shutdown] 時に呼び出す関数を登録します。
// これは、ALPNプロトコルアップグレードを受けた接続やハイジャックされた接続を優雅にシャットダウンするために使用できます。
// この関数は、プロトコル固有の優雅なシャットダウンを開始する必要がありますが、シャットダウンが完了するのを待つ必要はありません。
func (srv *Server) RegisterOnShutdown(f func())

// ConnStateは、サーバーへのクライアント接続の状態を表します。
// これは、オプションの [Server.ConnState] フックによって使用されます。
type ConnState int

const (
	// StateNewは、すぐにリクエストを送信することが期待される新しい接続を表します。
	// 接続はこの状態で開始し、StateActiveまたはStateClosedに移行します。
	StateNew ConnState = iota

	// StateActiveは、1バイト以上のリクエストを読み取った接続を表します。
	// StateActiveのServer.ConnStateフックは、リクエストがハンドラに入る前に発生し、
	// リクエストが処理されるまで再び発生しません。
	// リクエストが処理された後、状態はStateClosed、StateHijacked、またはStateIdleに移行します。
	// HTTP/2の場合、StateActiveはゼロから1つのアクティブなリクエストに移行するときに発生し、
	// すべてのアクティブなリクエストが完了するまでにしか移行しません。
	// つまり、ConnStateはリクエストごとの作業に使用できません。
	// ConnStateは接続の全体的な状態のみを示します。
	StateActive

	// StateIdleは、リクエストの処理が完了し、新しいリクエストを待機しているkeep-alive状態の接続を表します。
	// 接続はStateIdleからStateActiveまたはStateClosedに移行します。
	StateIdle

	// StateHijackedは、ハイジャックされた接続を表します。
	// これは終端状態です。StateClosedに移行しません。
	StateHijacked

	// StateClosedは、閉じられた接続を表します。
	// これは終端状態です。ハイジャックされた接続はStateClosedに移行しません。
	StateClosed
)

func (c ConnState) String() string

// AllowQuerySemicolonsは、URLクエリ内のエスケープされていないセミコロンをアンパサンドに変換し、ハンドラhを呼び出すハンドラを返します。
//
// これにより、Go 1.17以前のクエリパラメータをセミコロンとアンパサンドの両方で分割する動作が復元されます（golang.org/issue/25192 を参照）。
// ただし、この動作は多くのプロキシと一致せず、不一致がセキュリティ上の問題を引き起こす可能性があります。
//
// AllowQuerySemicolonsは、[Request.ParseForm] が呼び出される前に呼び出す必要があります。
func AllowQuerySemicolons(h Handler) Handler

// ListenAndServeは、TCPネットワークアドレスs.Addrでリッスンし、
// [Serve] を呼び出して着信接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// s.Addrが空白の場合、":http"が使用されます。
//
// ListenAndServeは常に非nilのエラーを返します。[Server.Shutdown] または [Server.Close] の後、
// 返されるエラーは [ErrServerClosed] です。
func (srv *Server) ListenAndServe() error

// ErrServerClosedは、[Server.Shutdown] または [Server.Close] の呼び出し後、[Server.Serve]、[ServeTLS]、[ListenAndServe]、および [ListenAndServeTLS] メソッドによって返されます。
var ErrServerClosed = errors.New("http: Server closed")

// Serveは、Listener lで着信接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはリクエストを読み取り、srv.Handlerを呼び出してそれに応答します。
//
// Listenerが [*tls.Conn] 接続を返し、TLS Config.NextProtosで「h2」が構成されている場合、HTTP/2サポートが有効になります。
//
// Serveは常に非nilのエラーを返し、lを閉じます。
// [Server.Shutdown] または [Server.Close] の後、返されるエラーは [ErrServerClosed] です。
func (srv *Server) Serve(l net.Listener) error

// ServeTLSは、Listener lで着信接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはTLSのセットアップを実行し、リクエストを読み取り、s.Handlerを呼び出してそれに応答します。
//
// サーバー用の証明書と一致する秘密鍵を含むファイルを提供する必要があります。これは、[Server] の
// TLSConfig.Certificates、TLSConfig.GetCertificate、または
// config.GetConfigForClientが設定されていない場合に必要です。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
//
// ServeTLSは常に非nilのエラーを返し、lを閉じます。
// [Server.Shutdown] または [Server.Close] の後、返されるエラーは [ErrServerClosed] です。
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error

// SetKeepAlivesEnabledは、HTTP keep-alivesが有効かどうかを制御します。
// デフォルトでは、keep-alivesは常に有効になっています。非常にリソースが制限された環境またはシャットダウン中のサーバーのみ、それらを無効にする必要があります。
func (srv *Server) SetKeepAlivesEnabled(v bool)

// ListenAndServeは、TCPネットワークアドレスaddrでリッスンし、
// [Serve] を呼び出して着信接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// ハンドラは通常nilであり、その場合は [DefaultServeMux] が使用されます。
//
// ListenAndServeは常に非nilのエラーを返します。
func ListenAndServe(addr string, handler Handler) error

// ListenAndServeTLSは、[ListenAndServe] と同じように動作しますが、HTTPS接続を想定しています。
// さらに、サーバーの証明書と一致する秘密鍵を含むファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error

// ListenAndServeTLSは、TCPネットワークアドレスsrv.Addrでリッスンし、
// [ServeTLS] を呼び出して着信TLS接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// [Server] のTLSConfig.CertificatesまたはTLSConfig.GetCertificateがどちらも設定されていない場合、
// サーバーの証明書と対応する秘密鍵が含まれるファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
//
// srv.Addrが空白の場合、":https"が使用されます。
//
// ListenAndServeTLSは常に非nilのエラーを返します。[Server.Shutdown] または
// [Server.Close] の後、返されるエラーは [ErrServerClosed] です。
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error

// TimeoutHandlerは、指定された時間制限でhを実行する [Handler] を返します。
//
// 新しいHandlerは、各リクエストを処理するためにh.ServeHTTPを呼び出しますが、
// 呼び出しがその時間制限を超えて実行されると、ハンドラは503 Service Unavailableエラーと
// そのボディ内の指定されたメッセージで応答します。
// （もしmsgが空であれば、適切なデフォルトメッセージが送信されます。）
// そのようなタイムアウトの後、hによるその [ResponseWriter] への書き込みは
// [ErrHandlerTimeout] を返します。
//
// TimeoutHandlerは [Pusher] インターフェースをサポートしますが、
// [Hijacker] または [Flusher] インターフェースはサポートしません。
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

// ErrHandlerTimeout is returned on [ResponseWriter] Write calls
// in handlers which have timed out.
var ErrHandlerTimeout = errors.New("http: Handler timeout")

var _ Pusher = (*timeoutWriter)(nil)

// MaxBytesHandlerは、[ResponseWriter] と [Request.Body] をMaxBytesReaderでラップしてhを実行する [Handler] を返します。
func MaxBytesHandler(h Handler, n int64) Handler
