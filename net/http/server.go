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
// ServeHTTPは、応答ヘッダーとデータをResponseWriterに書き込んでから返す必要があります。
// 返すことで、リクエストが完了したことを示します。
// ServeHTTP呼び出しの完了後または同時に、ResponseWriterを使用するか、Request.Bodyから読み取ることは無効です。
//
// HTTPクライアントソフトウェア、HTTPプロトコルバージョン、およびクライアントとGoサーバーの間の中間者によっては、
// ResponseWriterに書き込んだ後にRequest.Bodyから読み取ることができない場合があります。
// 慎重なハンドラーは、最初にRequest.Bodyを読み取り、その後に返信する必要があります。
//
// ボディを読み取る以外の場合、ハンドラーは提供されたRequestを変更してはいけません。
//
// ServeHTTPがパニックを起こすと、サーバー（ServeHTTPの呼び出し元）は、パニックの影響がアクティブなリクエストに限定されたと仮定します。
// サーバーはパニックを回復し、サーバーエラーログにスタックトレースを記録し、ネットワーク接続を閉じるか、HTTP/2 RST_STREAMを送信します。
// クライアントが中断された応答を見るが、サーバーがエラーをログに記録しないように、ErrAbortHandlerの値でパニックを発生させることで、ハンドラーを中止できます。
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// ResponseWriterインターフェースは、HTTPハンドラーがHTTPレスポンスを構築するために使用されます。
//
// Handler.ServeHTTPメソッドが返された後に、ResponseWriterを使用することはできません。
type ResponseWriter interface {
	// Headerは、WriteHeaderによって送信されるヘッダーマップを返します。
	// ヘッダーマップはまた、ハンドラがHTTPトレイラーを設定するメカニズムでもあります。
	//
	// WriteHeader（またはWrite）の呼び出し後にヘッダーマップを変更しても、
	// HTTPステータスコードが1xxクラスであった場合、または変更されたヘッダーがトレイラーであった場合を除き、
	// 影響はありません。
	//
	// トレイラーを設定する方法は2つあります。推奨される方法は、
	// ヘッダーで後で送信するトレイラーを事前に宣言することです。
	// これは、"Trailer"ヘッダーを後で来るトレイラーキーの名前に設定することで行います。
	// この場合、ヘッダーマップのキーはトレイラーであるかのように扱われます。例を参照してください。
	// 2つ目の方法は、最初のWriteの後までハンドラーには未知のトレイラーキーについて、
	// ヘッダーマップのキーにTrailerPrefix定数値をプレフィックスとして付けることです。
	// TrailerPrefixを参照してください。
	//
	// 自動的に設定されるレスポンスヘッダー（例えば "Date"）を抑制するには、
	// その値をnilに設定します。
	Header() Header

	// Writeは、HTTP応答の一部としてデータを接続に書き込みます。
	//
	// WriteHeaderがまだ呼び出されていない場合、Writeはデータを書き込む前に
	// WriteHeader(http.StatusOK)を呼び出します。ヘッダーにContent-Type行が含まれていない場合、
	// Writeは書き込まれたデータの最初の512バイトをDetectContentTypeに渡す結果に設定されたContent-Typeを追加します。
	// さらに、書き込まれたすべてのデータの合計サイズが数KB以下でFlush呼び出しがない場合、
	// Content-Lengthヘッダーが自動的に追加されます。
	//
	// HTTPプロトコルのバージョンとクライアントによっては、
	// WriteまたはWriteHeaderを呼び出すと、今後のRequest.Bodyの読み取りが防止される場合があります。
	// HTTP/1.xのリクエストに対しては、ハンドラはレスポンスを書き込む前に必要なリクエストボディデータを読み取るべきです。
	// ヘッダがフラッシュされると（明示的なFlusher.Flush呼び出しによるものか、フラッシュをトリガーするだけのデータを書き込むことによるものかは問わず）、
	// リクエストボディは利用できなくなる可能性があります。HTTP/2のリクエストに対しては、GoのHTTPサーバーはハンドラがレスポンスを
	// 並行して書き込みながらリクエストボディを読み続けることを許可します。しかし、そのような動作はすべてのHTTP/2クライアントでサポートされているわけではありません。
	// 可能な限り互換性を最大化するために、ハンドラは書き込む前に読み取るべきです。
	Write([]byte) (int, error)

	// WriteHeaderは、提供されたステータスコードでHTTPレスポンスヘッダーを送信します。
	//
	// WriteHeaderが明示的に呼び出されない場合、最初のWriteの呼び出しは
	// 暗黙的なWriteHeader(http.StatusOK)をトリガーします。
	// したがって、WriteHeaderへの明示的な呼び出しは主に、
	// エラーコードや1xx情報応答を送信するために使用されます。
	//
	// 提供されるコードは、有効なHTTP 1xx-5xxのステータスコードでなければなりません。
	// 任意の数の1xxヘッダーを書き込むことができ、その後に最大で
	// 一つの2xx-5xxヘッダーが続きます。1xxヘッダーはすぐに送信されますが、2xx-5xx
	// ヘッダーはバッファリングされる可能性があります。バッファリングされたデータを送信するには
	// Flusherインターフェースを使用します。2xx-5xxヘッダーが
	// 送信されるとヘッダーマップはクリアされますが、1xxヘッダーではクリアされません。
	//
	// サーバーは、リクエストが "Expect: 100-continue" ヘッダーを持っている場合、
	// リクエストボディからの最初の読み取り時に自動的に100（Continue）ヘッダーを送信します。
	WriteHeader(statusCode int)
}

// Flusherインターフェースは、HTTPハンドラーがバッファリングされたデータをクライアントにフラッシュすることを許可するResponseWriterによって実装されます。
//
// デフォルトのHTTP/1.xおよびHTTP/2 ResponseWriter実装はFlusherをサポートしていますが、ResponseWriterラッパーはサポートしていない場合があります。
// ハンドラーは常にランタイムでこの機能をテストする必要があります。
//
// FlushをサポートするResponseWriterであっても、クライアントがHTTPプロキシを介して接続されている場合、
// バッファリングされたデータがレスポンスが完了するまでクライアントに到達しない場合があります。
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}

// Hijackerインターフェースは、HTTPハンドラーが接続を引き継ぐことを許可するResponseWriterによって実装されます。
//
// HTTP/1.x接続のデフォルトResponseWriterはHijackerをサポートしていますが、HTTP/2接続は意図的にサポートしていません。
// ResponseWriterラッパーもHijackerをサポートしていない場合があります。
// ハンドラーは常にランタイムでこの機能をテストする必要があります。
type Hijacker interface {
	// Hijack lets the caller take over the connection.
	// After a call to Hijack the HTTP server library
	// will not do anything else with the connection.
	//
	// It becomes the caller's responsibility to manage
	// and close the connection.
	//
	// The returned net.Conn may have read or write deadlines
	// already set, depending on the configuration of the
	// Server. It is the caller's responsibility to set
	// or clear those deadlines as needed.
	//
	// The returned bufio.Reader may contain unprocessed buffered
	// data from the client.
	//
	// After a call to Hijack, the original Request.Body must not
	// be used. The original Request's Context remains valid and
	// is not canceled until the Request's ServeHTTP method
	// returns.
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}

// CloseNotifierインターフェースは、基礎となる接続が切断されたときに検出できるResponseWriterによって実装されます。
//
// このメカニズムは、レスポンスが準備される前にクライアントが切断された場合、サーバー上の長時間の操作をキャンセルするために使用できます。
//
// 廃止予定: CloseNotifierインターフェースは、Goのコンテキストパッケージより前に実装されました。
// 新しいコードでは、Request.Contextを使用する必要があります。
type CloseNotifier interface {
	// CloseNotify returns a channel that receives at most a
	// single value (true) when the client connection has gone
	// away.
	//
	// CloseNotify may wait to notify until Request.Body has been
	// fully read.
	//
	// After the Handler has returned, there is no guarantee
	// that the channel receives a value.
	//
	// If the protocol is HTTP/1.1 and CloseNotify is called while
	// processing an idempotent request (such a GET) while
	// HTTP/1.1 pipelining is in use, the arrival of a subsequent
	// pipelined request may cause a value to be sent on the
	// returned channel. In practice HTTP/1.1 pipelining is not
	// enabled in browsers and not seen often in the wild. If this
	// is a problem, use HTTP/2 or only use CloseNotify on methods
	// such as POST.
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

// TrailerPrefixは、ResponseWriter.Headerマップのキーに対するマジックプレフィックスで、
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
// これは、Server.MaxHeaderBytesを設定することで上書きできます。
const DefaultMaxHeaderBytes = 1 << 20

// TimeFormatは、HTTPヘッダーで時間を生成するときに使用する時間形式です。
// time.RFC1123のようですが、タイムゾーンとしてGMTがハードコードされています。
// フォーマットされる時間はUTCである必要があります。
//
// この時間形式を解析するには、ParseTimeを参照してください。
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

var _ closeWriter = (*net.TCPConn)(nil)

// ErrAbortHandlerは、ハンドラーを中止するためのセンチネルパニック値です。
// ServeHTTPからのパニックはすべて、クライアントへの応答を中止しますが、
// ErrAbortHandlerでパニックすると、サーバーのエラーログにスタックトレースを記録しないようにすることができます。
var ErrAbortHandler = errors.New("net/http: abort Handler")

// HandlerFunc型は、HTTPハンドラーとして通常の関数を使用できるようにするためのアダプタです。
// fが適切なシグネチャを持つ関数である場合、HandlerFunc(f)はfを呼び出すHandlerです。
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTPは、f(w, r) を呼び出します。
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

// Errorは、指定されたエラーメッセージとHTTPコードでリクエストに応答します。
// それ以外に、リクエストを終了しません。呼び出し元は、wに対してさらに書き込みが行われないようにする必要があります。
// エラーメッセージはプレーンテキストである必要があります。
func Error(w ResponseWriter, error string, code int)

// NotFoundは、HTTP 404 not foundエラーでリクエストに応答します。
func NotFound(w ResponseWriter, r *Request)

// NotFoundHandlerは、各リクエストに「404ページが見つかりません」という応答を返す単純なリクエストハンドラーを返します。
func NotFoundHandler() Handler

// StripPrefixは、リクエストURLのPath（および設定されている場合はRawPath）から指定された接頭辞を削除し、ハンドラーhを呼び出してHTTPリクエストを処理するハンドラーを返します。
// prefixで始まらないパスのリクエストには、HTTP 404 not foundエラーで応答します。
// 接頭辞は完全に一致する必要があります。リクエストの接頭辞にエスケープされた文字が含まれている場合、応答もHTTP 404 not foundエラーになります。
func StripPrefix(prefix string, h Handler) Handler

// Redirectは、リクエストに対してurlにリダイレクトする応答を返します。
// urlは、リクエストパスに対する相対パスである場合があります。
//
// 提供されたコードは通常、StatusMovedPermanently、StatusFound、またはStatusSeeOtherの3xx範囲にあります。
//
// Content-Typeヘッダーが設定されていない場合、Redirectはそれを"text/html; charset=utf-8"に設定し、小さなHTML本文を書き込みます。
// Content-Typeヘッダーを任意の値、nilを含む任意の値に設定すると、その動作が無効になります。
func Redirect(w ResponseWriter, r *Request, url string, code int)

// RedirectHandlerは、受信した各リクエストを、指定されたステータスコードを使用して、指定されたURLにリダイレクトするリクエストハンドラーを返します。
//
// 提供されたコードは通常、StatusMovedPermanently、StatusFound、またはStatusSeeOtherの3xx範囲にあります。
func RedirectHandler(url string, code int) Handler

// ServeMuxは、HTTPリクエストマルチプレクサーです。
// それは、登録されたパターンのリストに対して、各受信リクエストのURLを一致させ、
// URLに最も近いパターンのハンドラーを呼び出します。
//
// パターンは、"/favicon.ico"のような固定されたルートパス、または"/images/"のようなルートサブツリーの名前を付けます（末尾のスラッシュに注意）。
// より長いパターンが優先されるため、"/images/"と"/images/thumbnails/"の両方にハンドラーが登録されている場合、後者のハンドラーは"/images/thumbnails/"で始まるパスに対して呼び出され、前者は"/images/"サブツリー内の他のパスに対してリクエストを受け取ります。
//
// スラッシュで終わるパターンは、ルートサブツリーを名前付けるため、注意が必要です。
// パターン"/"は、他の登録されたパターンに一致しないすべてのパス（Path == "/"のURLだけでなく）に一致します。
//
// サブツリーが登録され、トレーリングスラッシュなしでサブツリールートを指定するリクエストが受信された場合、ServeMuxはそのリクエストをサブツリールートにリダイレクトします（トレーリングスラッシュを追加）。
// この動作は、トレーリングスラッシュなしのパスに対する別個の登録でオーバーライドできます。たとえば、"/images/"を登録すると、ServeMuxは"/images"のリクエストを"/images/"にリダイレクトしますが、"/images"が別個に登録されている場合は、リダイレクトは行われません。
//
// パターンは、ホスト名で始まることがあり、そのホスト上のURLにのみ一致するように制限できます。ホスト固有のパターンは、一般的なパターンより優先されるため、ハンドラーは"/codesearch"と"codesearch.google.com/"の2つのパターンに登録でき、"http://www.google.com/"のリクエストを引き継ぐことはありません。
//
// ServeMuxは、URLリクエストパスとHostヘッダーをサニタイズし、ポート番号を削除し、.または..要素または重複したスラッシュを含むリクエストを同等のクリーンなURLにリダイレクトします。
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry
	hosts bool
}

// NewServeMuxは、新しいServeMuxを割り当てて返します。
func NewServeMux() *ServeMux

// DefaultServeMuxは、Serveによって使用されるデフォルトのServeMuxです。
var DefaultServeMux = &defaultServeMux

// Handlerは、r.Method、r.Host、およびr.URL.Pathを参照して、
// 指定されたリクエストに使用するハンドラーを返します。
// 常にnilでないハンドラーを返します。
// パスが正規形式でない場合、ハンドラーは正規パスにリダイレクトする内部生成ハンドラーになります。
// ホストにポートが含まれている場合、ハンドラーの一致時には無視されます。
//
// CONNECTリクエストでは、パスとホストは変更されずに使用されます。
//
// Handlerは、リクエストに一致する登録済みパターンと、
// 内部生成されたリダイレクトの場合は、リダイレクトをたどった後に一致するパターンを返します。
//
// リクエストに適用される登録済みハンドラーがない場合、
// Handlerは「ページが見つかりません」というハンドラーと空のパターンを返します。
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)

// ServeHTTPは、リクエストURLに最も近いパターンを持つハンドラにリクエストをディスパッチします。
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

// Handleは、指定されたパターンのハンドラを登録します。
// パターンに対するハンドラがすでに存在する場合、Handleはパニックを発生させます。
func (mux *ServeMux) Handle(pattern string, handler Handler)

// HandleFuncは、指定されたパターンのハンドラ関数を登録します。
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Handleは、DefaultServeMuxに指定されたパターンのハンドラを登録します。
// ServeMuxのドキュメントには、パターンがどのようにマッチするかが説明されています。
func Handle(pattern string, handler Handler)

// HandleFuncは、DefaultServeMuxに指定されたパターンのハンドラ関数を登録します。
// ServeMuxのドキュメントには、パターンがどのようにマッチするかが説明されています。
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// Serveは、リスナーlに対して着信HTTP接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはリクエストを読み取り、その後handlerを呼び出して応答します。
//
// handlerは通常nilであり、その場合はDefaultServeMuxが使用されます。
//
// TLS Config.NextProtosで"h2"が設定された *tls.Conn 接続を返すリスナーがある場合、HTTP / 2サポートが有効になります。
//
// Serveは常にnilでないエラーを返します。
func Serve(l net.Listener, handler Handler) error

// ServeTLSは、リスナーlに対して着信HTTPS接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはリクエストを読み取り、その後handlerを呼び出して応答します。
//
// handlerは通常nilであり、その場合はDefaultServeMuxが使用されます。
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

	// ReadHeaderTimeoutは、リクエストヘッダーを読み取るために許可される時間です。
	// ヘッダーを読み取った後、接続の読み取り期限がリセットされ、Handlerはボディに対して何が遅すぎるかを決定できます。
	// ReadHeaderTimeoutがゼロの場合、ReadTimeoutの値が使用されます。
	// 両方がゼロの場合、タイムアウトはありません。
	ReadHeaderTimeout time.Duration

	// WriteTimeoutは、レスポンスの書き込みがタイムアウトする前の最大時間です。
	// 新しいリクエストヘッダーが読み取られるたびにリセットされます。
	// ReadTimeoutと同様に、ハンドラがリクエストごとに決定を下すことを許可しません。
	// ゼロまたは負の値はタイムアウトがないことを意味します。
	WriteTimeout time.Duration

	// IdleTimeoutは、keep-aliveが有効な場合に次のリクエストを待機する最大時間です。
	// IdleTimeoutがゼロの場合、ReadTimeoutの値が使用されます。
	// 両方がゼロの場合、タイムアウトはありません。
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

// Closeは、すべてのアクティブなnet.Listenerと、StateNew、StateActive、またはStateIdleの状態にある接続をすぐに閉じます。
// 優雅なシャットダウンには、Shutdownを使用してください。
//
// Closeは、WebSocketsなどのハイジャックされた接続を閉じようとはせず（そしてそれらについては何も知りません）、試みません。
//
// Closeは、Serverの基礎となるListenerの閉じる際に返されるエラーを返します。
func (srv *Server) Close() error

// Shutdownは、アクティブな接続を中断することなく、サーバーを優雅にシャットダウンします。
// Shutdownは、まずすべてのオープンリスナーを閉じ、次にすべてのアイドル接続を閉じ、接続がアイドル状態に戻ってからシャットダウンするまで無期限に待機します。
// 提供されたコンテキストがシャットダウンが完了する前に期限切れになった場合、Shutdownはコンテキストのエラーを返します。それ以外の場合、Serverの基礎となるListenerの閉じる際に返されるエラーを返します。
//
// Shutdownが呼び出されると、Serve、ListenAndServe、およびListenAndServeTLSはすぐにErrServerClosedを返します。プログラムが終了せずにShutdownが返るのを待つようにしてください。
//
// Shutdownは、WebSocketsなどのハイジャックされた接続を閉じようとはせず（そしてそれらについては何も知りません）、試みません。
//
// Shutdownを呼び出した後、サーバーを再利用することはできません。Serveなどのメソッドを呼び出すと、ErrServerClosedが返されます。
func (srv *Server) Shutdown(ctx context.Context) error

// RegisterOnShutdownは、Shutdown時に呼び出す関数を登録します。
// これは、ALPNプロトコルアップグレードを受けた接続やハイジャックされた接続を優雅にシャットダウンするために使用できます。
// この関数は、プロトコル固有の優雅なシャットダウンを開始する必要がありますが、シャットダウンが完了するのを待つ必要はありません。
func (srv *Server) RegisterOnShutdown(f func())

// ConnStateは、サーバーへのクライアント接続の状態を表します。
// これは、オプションのServer.ConnStateフックによって使用されます。
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
// AllowQuerySemicolonsは、Request.ParseFormが呼び出される前に呼び出す必要があります。
func AllowQuerySemicolons(h Handler) Handler

// ListenAndServeは、TCPネットワークアドレスsrv.Addrでリッスンし、
// Serveを呼び出して着信接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// srv.Addrが空白の場合、":http"が使用されます。
//
// ListenAndServeは常に非nilのエラーを返します。ShutdownまたはCloseの後、
// 返されるエラーはErrServerClosedです。
func (srv *Server) ListenAndServe() error

// ErrServerClosedは、ShutdownまたはCloseの呼び出し後、ServerのServe、ServeTLS、ListenAndServe、およびListenAndServeTLSメソッドによって返されます。
var ErrServerClosed = errors.New("http: Server closed")

// Serveは、Listener lで着信接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはリクエストを読み取り、srv.Handlerを呼び出してそれに応答します。
//
// Listenerが* tls.Conn接続を返し、TLS Config.NextProtosで「h2」が構成されている場合、HTTP/2サポートが有効になります。
//
// Serveは常に非nilのエラーを返し、lを閉じます。
// ShutdownまたはCloseの後、返されるエラーはErrServerClosedです。
func (srv *Server) Serve(l net.Listener) error

// ServeTLSは、Listener lで着信接続を受け入れ、それぞれに新しいサービスgoroutineを作成します。
// サービスgoroutineはTLSのセットアップを実行し、リクエストを読み取り、srv.Handlerを呼び出してそれに応答します。
//
// サーバーのTLSConfig.CertificatesまたはTLSConfig.GetCertificateがどちらも設定されていない場合、
// サーバーの証明書と対応する秘密鍵が含まれるファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
//
// ServeTLSは常に非nilのエラーを返し、lを閉じます。
// ShutdownまたはCloseの後、返されるエラーはErrServerClosedです。
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error

// SetKeepAlivesEnabledは、HTTP keep-alivesが有効かどうかを制御します。
// デフォルトでは、keep-alivesは常に有効になっています。非常にリソースが制限された環境またはシャットダウン中のサーバーのみ、それらを無効にする必要があります。
func (srv *Server) SetKeepAlivesEnabled(v bool)

// ListenAndServeは、TCPネットワークアドレスaddrでリッスンし、
// Serveを呼び出して着信接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// ハンドラは通常nilであり、その場合はDefaultServeMuxが使用されます。
//
// ListenAndServeは常に非nilのエラーを返します。
func ListenAndServe(addr string, handler Handler) error

// ListenAndServeTLSは、ListenAndServeと同じように動作しますが、HTTPS接続が必要です。
// さらに、サーバーの証明書と対応する秘密鍵が含まれるファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error

// ListenAndServeTLSは、TCPネットワークアドレスsrv.Addrでリッスンし、
// ServeTLSを呼び出して着信TLS接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// サーバーのTLSConfig.CertificatesまたはTLSConfig.GetCertificateがどちらも設定されていない場合、
// サーバーの証明書と対応する秘密鍵が含まれるファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
//
// srv.Addrが空白の場合、":https"が使用されます。
//
// ListenAndServeTLSは常に非nilのエラーを返します。ShutdownまたはCloseの後、
// 返されるエラーはErrServerClosedです。
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error

// ListenAndServeTLSは、TCPネットワークアドレスsrv.Addrでリッスンし、
// ServeTLSを呼び出して着信TLS接続のリクエストを処理します。
// 受け入れられた接続は、TCP keep-alivesを有効にするように構成されます。
//
// サーバーのTLSConfig.CertificatesまたはTLSConfig.GetCertificateがどちらも設定されていない場合、
// サーバーの証明書と対応する秘密鍵が含まれるファイルを提供する必要があります。
// 証明書が認証局によって署名されている場合、certFileはサーバーの証明書、中間証明書、およびCAの証明書を連結したものである必要があります。
//
// srv.Addrが空白の場合、":https"が使用されます。
//
// ListenAndServeTLSは常に非nilのエラーを返します。ShutdownまたはCloseの後、
// 返されるエラーはErrServerClosedです。
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

// ErrHandlerTimeoutは、タイムアウトしたハンドラ内のResponseWriter Write呼び出しで返されます。
var ErrHandlerTimeout = errors.New("http: Handler timeout")

var _ Pusher = (*timeoutWriter)(nil)

// MaxBytesHandlerは、ResponseWriterとRequest.BodyをMaxBytesReaderでラップしてhを実行するハンドラを返します。
func MaxBytesHandler(h Handler, n int64) Handler
