// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP Request reading and parsing.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/mime/multipart"
	"github.com/shogo82148/std/net/url"
)

// ErrMissingFileは、FormFileが提供されたファイルフィールド名がリクエストに存在しないか、ファイルフィールドではない場合に返されます。
var ErrMissingFile = errors.New("http: no such file")

// ProtocolErrorは、HTTPプロトコルエラーを表します。
//
// Deprecated: httpパッケージのすべてのプロトコルエラーに関連するエラーがProtocolError型ではありません。
type ProtocolError struct {
	ErrorString string
}

func (pe *ProtocolError) Error() string

// Isは、http.ErrNotSupportedがerrors.ErrUnsupportedに一致するようにします。
func (pe *ProtocolError) Is(err error) bool

var (
	// ErrNotSupportedは、機能がサポートされていないことを示します。
	//
	// ResponseControllerメソッドによって、ハンドラがメソッドをサポートしていないことを示すために返され、
	// Pusher実装のPushメソッドによって、HTTP/2 Pushサポートが利用できないことを示すために返されます。
	ErrNotSupported = &ProtocolError{"feature not supported"}

	// Deprecated: ErrUnexpectedTrailerは、net/httpパッケージの何も返さなくなりました。
	// 呼び出し元は、この変数とエラーを比較すべきではありません。
	ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}

	// ErrMissingBoundaryは、リクエストのContent-Typeに「boundary」パラメータが含まれていない場合に、Request.MultipartReaderによって返されます。
	ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}

	/// ErrNotMultipartは、リクエストのContent-Typeがmultipart/form-dataでない場合、Request.MultipartReaderによって返されます。
	ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}

	// Deprecated: ErrHeaderTooLongは、net/httpパッケージの何も返さなくなりました。
	// 呼び出し元は、この変数とエラーを比較すべきではありません。
	ErrHeaderTooLong = &ProtocolError{"header too long"}

	// Deprecated: ErrShortBodyは、net/httpパッケージの何も返さなくなりました。
	// 呼び出し元は、この変数とエラーを比較すべきではありません。
	ErrShortBody = &ProtocolError{"entity body too short"}

	// Deprecated: ErrMissingContentLengthは、net/httpパッケージの何も返さなくなりました。
	// 呼び出し元は、この変数とエラーを比較すべきではありません。
	ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
)

// Requestは、サーバーによって受信されたHTTPリクエストまたはクライアントによって送信されるHTTPリクエストを表します。
//
// フィールドの意味は、クライアントとサーバーの使用方法でわずかに異なります。
// 以下のフィールドに関する注意事項に加えて、Request.WriteおよびRoundTripperのドキュメントを参照してください。
type Request struct {
	// Methodは、HTTPメソッド（GET、POST、PUTなど）を指定します。
	// クライアントリクエストの場合、空の文字列はGETを意味します。
	//
	// GoのHTTPクライアントは、CONNECTメソッドでリクエストを送信することをサポートしていません。
	// 詳細については、Transportのドキュメントを参照してください。
	Method string

	// URLは、サーバーリクエストの場合に要求されるURI（URI）を指定するか、
	// クライアントリクエストの場合にアクセスするURLを指定します。
	//
	// サーバーリクエストの場合、URLはRequestURIに格納されたRequest-Lineで指定されたURIから解析されます。
	// ほとんどのリクエストでは、PathとRawQuery以外のフィールドは空になります。 （RFC 7230、セクション5.3を参照）
	//
	// クライアントリクエストの場合、URLのHostは接続するサーバーを指定し、
	// RequestのHostフィールドはHTTPリクエストで送信するHostヘッダー値をオプションで指定します。
	URL *url.URL

	// サーバーリクエストのプロトコルバージョン。
	//
	// クライアントリクエストの場合、これらのフィールドは無視されます。
	// HTTPクライアントコードは常にHTTP/1.1またはHTTP/2を使用します。
	// 詳細については、Transportのドキュメントを参照してください。
	Proto      string
	ProtoMajor int
	ProtoMinor int

	// Headerは、サーバーによって受信されたリクエストヘッダーフィールドまたはクライアントによって送信されるリクエストヘッダーフィールドを含みます。
	//
	// サーバーがヘッダーラインを含むリクエストを受信した場合、
	//
	//	Host: example.com
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	fOO: Bar
	//	foo: two
	//
	// その場合、
	//
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
	//
	// 入力リクエストの場合、HostヘッダーはRequest.Hostフィールドに昇格し、Headerマップから削除されます。
	//
	// HTTPは、ヘッダー名が大文字小文字を区別しないことを定義しています。リクエストパーサーは、CanonicalHeaderKeyを使用してこれを実装し、ハイフンの後に続く最初の文字と任意の文字を大文字にし、残りを小文字にします。
	//
	// クライアントリクエストの場合、Content-LengthやConnectionなどの特定のヘッダーは必要に応じて自動的に書き込まれ、Headerの値は無視される場合があります。Request.Writeメソッドのドキュメントを参照してください。
	Header Header

	// Bodyは、リクエストの本文です。
	//
	// クライアントリクエストの場合、nilのBodyは、GETリクエストなど、リクエストに本文がないことを意味します。
	// HTTPクライアントのTransportは、Closeメソッドを呼び出す責任があります。
	//
	// サーバーリクエストの場合、Request Bodyは常にnilではありませんが、本文が存在しない場合はすぐにEOFが返されます。
	// サーバーはリクエストボディを閉じます。ServeHTTPハンドラーは閉じる必要はありません。
	//
	// Bodyは、ReadがCloseと同時に呼び出されることを許可する必要があります。
	// 特に、Closeを呼び出すと、入力を待機しているReadがブロックされている場合は、それを解除する必要があります。
	Body io.ReadCloser

	// GetBodyは、Bodyの新しいコピーを返すオプションの関数を定義します。
	// リダイレクトにより、Bodyを複数回読み取る必要がある場合にクライアントリクエストで使用されます。
	// GetBodyの使用には、Bodyを設定する必要があります。
	//
	// サーバーリクエストの場合、使用されません。
	GetBody func() (io.ReadCloser, error)

	// ContentLengthは、関連するコンテンツの長さを記録します。
	// 値-1は、長さが不明であることを示します。
	// 値が0以上の場合、Bodyから指定されたバイト数を読み取ることができます。
	//
	// クライアントリクエストの場合、Bodyがnilでない場合、値0は不明として扱われます。
	ContentLength int64

	// TransferEncodingは、最も外側から最も内側までの転送エンコーディングをリストします。
	// 空のリストは「identity」エンコーディングを示します。
	// TransferEncodingは通常無視できます。チャンクエンコーディングは、
	// リクエストの送信と受信時に必要に応じて自動的に追加および削除されます。
	TransferEncoding []string

	// Closeは、このリクエストに応答した後（サーバーの場合）またはこのリクエストを送信してその応答を読み取った後（クライアントの場合）に接続を閉じるかどうかを示します。
	//
	// サーバーリクエストの場合、HTTPサーバーがこれを自動的に処理し、このフィールドはHandlersによって必要ではありません。
	//
	// クライアントリクエストの場合、このフィールドを設定すると、 Transport.DisableKeepAlives が設定された場合と同様に、同じホストへのリクエスト間でTCP接続の再利用が防止されます。
	Close bool

	// サーバーリクエストの場合、HostはURLが要求されるホストを指定します。
	// HTTP/1（RFC 7230、セクション5.4による）の場合、これは「Host」ヘッダーの値またはURL自体で指定されたホスト名のいずれかです。
	// HTTP/2の場合、これは「:authority」疑似ヘッダーフィールドの値です。
	// 形式は「host:port」にすることができます。
	// 国際ドメイン名の場合、HostはPunycodeまたはUnicode形式である場合があります。
	// 必要に応じて、golang.org/x/net/idna を使用してどちらの形式にも変換できます。
	// DNSリバインディング攻撃を防ぐために、サーバーハンドラーは、Hostヘッダーがハンドラー自身が権限を持つ値であることを検証する必要があります。
	// 含まれるServeMuxは、特定のホスト名に登録されたパターンをサポートし、その登録されたハンドラーを保護します。
	//
	// クライアントリクエストの場合、HostはオプションでHostヘッダーを上書きして送信します。
	// 空の場合、Request.Write メソッドは URL.Host の値を使用します。
	// Hostには国際ドメイン名が含まれる場合があります。
	Host string

	// Formには、URLフィールドのクエリパラメータとPATCH、POST、またはPUTフォームデータを含む解析されたフォームデータが含まれます。
	// ParseFormが呼び出された後にのみこのフィールドが使用可能です。
	// HTTPクライアントはFormを無視し、代わりにBodyを使用します。
	Form url.Values

	// PostFormには、PATCH、POST、またはPUTボディパラメータから解析されたフォームデータが含まれます。
	//
	// ParseFormが呼び出された後にのみこのフィールドが使用可能です。
	// HTTPクライアントはPostFormを無視し、代わりにBodyを使用します。
	PostForm url.Values

	// MultipartFormは、ファイルのアップロードを含む解析されたマルチパートフォームです。
	// ParseMultipartFormが呼び出された後にのみこのフィールドが使用可能です。
	// HTTPクライアントはMultipartFormを無視し、代わりにBodyを使用します。
	MultipartForm *multipart.Form

	// Trailerは、リクエスト本文の後に送信される追加のヘッダーを指定します。
	//
	// サーバーリクエストの場合、Trailerマップには最初にトレーラーキーのみが含まれ、nil値が含まれます。
	// （クライアントは、後で送信するトレーラーを宣言します。）
	// ハンドラーがBodyから読み取っている間、Trailerに参照しないでください。
	// Bodyからの読み取りがEOFを返した後、Trailerを再度読み取ると、クライアントが送信した場合は非nil値が含まれます。
	//
	// クライアントリクエストの場合、Trailerは、後で送信するトレーラーキーを含むマップで初期化する必要があります。
	// 値はnilまたは最終値である場合があります。ContentLengthは0または-1である必要があります。
	// HTTPリクエストが送信された後、マップの値は、リクエストボディが読み取られている間に更新できます。
	// ボディがEOFを返した後、呼び出し元はTrailerを変更してはいけません。
	//
	// 少数のHTTPクライアント、サーバー、またはプロキシがHTTPトレーラーをサポートしています。
	Trailer Header

	// RemoteAddrは、HTTPサーバーやその他のソフトウェアが、通常はログ記録のためにリクエストを送信したネットワークアドレスを記録できるようにします。
	// このフィールドはReadRequestによって入力されず、定義された形式はありません。
	// このパッケージのHTTPサーバーは、ハンドラーを呼び出す前にRemoteAddrを「IP:port」アドレスに設定します。
	// このフィールドはHTTPクライアントによって無視されます。
	RemoteAddr string

	// RequestURIは、クライアントがサーバーに送信したRequest-Line（RFC 7230、セクション3.1.1）の変更されていないリクエストターゲットです。
	// 通常、URLフィールドを代わりに使用する必要があります。
	// HTTPクライアントリクエストでこのフィールドを設定することはエラーです。
	RequestURI string

	// TLSは、HTTPサーバーやその他のソフトウェアが、リクエストを受信したTLS接続に関する情報を記録できるようにします。
	// このフィールドはReadRequestによって入力されず、定義された形式はありません。
	// このパッケージのHTTPサーバーは、ハンドラーを呼び出す前にTLS有効な接続のためにこのフィールドを設定します。
	// それ以外の場合は、このフィールドをnilのままにします。
	// このフィールドはHTTPクライアントによって無視されます。
	TLS *tls.ConnectionState

	// Cancelは、クライアントリクエストがキャンセルされたと見なす必要があることを示すチャネルのオプションです。
	// RoundTripperのすべての実装がCancelをサポートしているわけではありません。
	//
	// サーバーリクエストの場合、このフィールドは適用されません。
	//
	// Deprecated: NewRequestWithContextでRequestのコンテキストを設定してください。
	// RequestのCancelフィールドとコンテキストの両方が設定されている場合、Cancelが尊重されるかどうかは未定義です。
	Cancel <-chan struct{}

	// Responseは、このリクエストが作成されたクライアントリダイレクト中にのみ設定されるリダイレクトレスポンスです。
	Response *Response

	// ctx is either the client or server context. It should only
	// be modified via copying the whole Request using Clone or WithContext.
	// It is unexported to prevent people from using Context wrong
	// and mutating the contexts held by callers of the same request.
	ctx context.Context
}

// Context returns the request's context. To change the context, use
// Clone or WithContext.
//
// The returned context is always non-nil; it defaults to the
// background context.
//
// For outgoing client requests, the context controls cancellation.
//
// For incoming server requests, the context is canceled when the
// client's connection closes, the request is canceled (with HTTP/2),
// or when the ServeHTTP method returns.
func (r *Request) Context() context.Context

// WithContext returns a shallow copy of r with its context changed
// to ctx. The provided ctx must be non-nil.
//
// For outgoing client request, the context controls the entire
// lifetime of a request and its response: obtaining a connection,
// sending the request, and reading the response headers and body.
//
// To create a new request with a context, use NewRequestWithContext.
// To make a deep copy of a request with a new context, use Request.Clone.
func (r *Request) WithContext(ctx context.Context) *Request

// Clone returns a deep copy of r with its context changed to ctx.
// The provided ctx must be non-nil.
//
// For an outgoing client request, the context controls the entire
// lifetime of a request and its response: obtaining a connection,
// sending the request, and reading the response headers and body.
func (r *Request) Clone(ctx context.Context) *Request

// ProtoAtLeast reports whether the HTTP protocol used
// in the request is at least major.minor.
func (r *Request) ProtoAtLeast(major, minor int) bool

// UserAgent returns the client's User-Agent, if sent in the request.
func (r *Request) UserAgent() string

// Cookies parses and returns the HTTP cookies sent with the request.
func (r *Request) Cookies() []*Cookie

// ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
var ErrNoCookie = errors.New("http: named cookie not present")

// Cookie returns the named cookie provided in the request or
// ErrNoCookie if not found.
// If multiple cookies match the given name, only one cookie will
// be returned.
func (r *Request) Cookie(name string) (*Cookie, error)

// AddCookie adds a cookie to the request. Per RFC 6265 section 5.4,
// AddCookie does not attach more than one Cookie header field. That
// means all cookies, if any, are written into the same line,
// separated by semicolon.
// AddCookie only sanitizes c's name and value, and does not sanitize
// a Cookie header already present in the request.
func (r *Request) AddCookie(c *Cookie)

// Referer returns the referring URL, if sent in the request.
//
// Referer is misspelled as in the request itself, a mistake from the
// earliest days of HTTP.  This value can also be fetched from the
// Header map as Header["Referer"]; the benefit of making it available
// as a method is that the compiler can diagnose programs that use the
// alternate (correct English) spelling req.Referrer() but cannot
// diagnose programs that use Header["Referrer"].
func (r *Request) Referer() string

// MultipartReader returns a MIME multipart reader if this is a
// multipart/form-data or a multipart/mixed POST request, else returns nil and an error.
// Use this function instead of ParseMultipartForm to
// process the request body as a stream.
func (r *Request) MultipartReader() (*multipart.Reader, error)

// Write writes an HTTP/1.1 request, which is the header and body, in wire format.
// This method consults the following fields of the request:
//
//	Host
//	URL
//	Method (defaults to "GET")
//	Header
//	ContentLength
//	TransferEncoding
//	Body
//
// If Body is present, Content-Length is <= 0 and TransferEncoding
// hasn't been set to "identity", Write adds "Transfer-Encoding:
// chunked" to the header. Body is closed after it is sent.
func (r *Request) Write(w io.Writer) error

// WriteProxy is like Write but writes the request in the form
// expected by an HTTP proxy. In particular, WriteProxy writes the
// initial Request-URI line of the request with an absolute URI, per
// section 5.3 of RFC 7230, including the scheme and host.
// In either case, WriteProxy also writes a Host header, using
// either r.Host or r.URL.Host.
func (r *Request) WriteProxy(w io.Writer) error

// ParseHTTPVersion parses an HTTP version string according to RFC 7230, section 2.6.
// "HTTP/1.0" returns (1, 0, true). Note that strings without
// a minor version, such as "HTTP/2", are not valid.
func ParseHTTPVersion(vers string) (major, minor int, ok bool)

// NewRequest wraps NewRequestWithContext using context.Background.
func NewRequest(method, url string, body io.Reader) (*Request, error)

// NewRequestWithContext returns a new Request given a method, URL, and
// optional body.
//
// If the provided body is also an io.Closer, the returned
// Request.Body is set to body and will be closed by the Client
// methods Do, Post, and PostForm, and Transport.RoundTrip.
//
// NewRequestWithContext returns a Request suitable for use with
// Client.Do or Transport.RoundTrip. To create a request for use with
// testing a Server Handler, either use the NewRequest function in the
// net/http/httptest package, use ReadRequest, or manually update the
// Request fields. For an outgoing client request, the context
// controls the entire lifetime of a request and its response:
// obtaining a connection, sending the request, and reading the
// response headers and body. See the Request type's documentation for
// the difference between inbound and outbound request fields.
//
// If body is of type *bytes.Buffer, *bytes.Reader, or
// *strings.Reader, the returned request's ContentLength is set to its
// exact value (instead of -1), GetBody is populated (so 307 and 308
// redirects can replay the body), and Body is set to NoBody if the
// ContentLength is 0.
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)

// BasicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func (r *Request) BasicAuth() (username, password string, ok bool)

// SetBasicAuth sets the request's Authorization header to use HTTP
// Basic Authentication with the provided username and password.
//
// With HTTP Basic Authentication the provided username and password
// are not encrypted. It should generally only be used in an HTTPS
// request.
//
// The username may not contain a colon. Some protocols may impose
// additional requirements on pre-escaping the username and
// password. For instance, when used with OAuth2, both arguments must
// be URL encoded first with url.QueryEscape.
func (r *Request) SetBasicAuth(username, password string)

// ReadRequest reads and parses an incoming request from b.
//
// ReadRequest is a low-level function and should only be used for
// specialized applications; most code should use the Server to read
// requests and handle them via the Handler interface. ReadRequest
// only supports HTTP/1.x requests. For HTTP/2, use golang.org/x/net/http2.
func ReadRequest(b *bufio.Reader) (*Request, error)

// MaxBytesReader is similar to io.LimitReader but is intended for
// limiting the size of incoming request bodies. In contrast to
// io.LimitReader, MaxBytesReader's result is a ReadCloser, returns a
// non-nil error of type *MaxBytesError for a Read beyond the limit,
// and closes the underlying reader when its Close method is called.
//
// MaxBytesReader prevents clients from accidentally or maliciously
// sending a large request and wasting server resources. If possible,
// it tells the ResponseWriter to close the connection after the limit
// has been reached.
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser

// MaxBytesError is returned by MaxBytesReader when its read limit is exceeded.
type MaxBytesError struct {
	Limit int64
}

func (e *MaxBytesError) Error() string

// ParseForm populates r.Form and r.PostForm.
//
// For all requests, ParseForm parses the raw query from the URL and updates
// r.Form.
//
// For POST, PUT, and PATCH requests, it also reads the request body, parses it
// as a form and puts the results into both r.PostForm and r.Form. Request body
// parameters take precedence over URL query string values in r.Form.
//
// If the request Body's size has not already been limited by MaxBytesReader,
// the size is capped at 10MB.
//
// For other HTTP methods, or when the Content-Type is not
// application/x-www-form-urlencoded, the request Body is not read, and
// r.PostForm is initialized to a non-nil, empty value.
//
// ParseMultipartForm calls ParseForm automatically.
// ParseForm is idempotent.
func (r *Request) ParseForm() error

// ParseMultipartForm parses a request body as multipart/form-data.
// The whole request body is parsed and up to a total of maxMemory bytes of
// its file parts are stored in memory, with the remainder stored on
// disk in temporary files.
// ParseMultipartForm calls ParseForm if necessary.
// If ParseForm returns an error, ParseMultipartForm returns it but also
// continues parsing the request body.
// After one call to ParseMultipartForm, subsequent calls have no effect.
func (r *Request) ParseMultipartForm(maxMemory int64) error

// FormValue returns the first value for the named component of the query.
// POST and PUT body parameters take precedence over URL query string values.
// FormValue calls ParseMultipartForm and ParseForm if necessary and ignores
// any errors returned by these functions.
// If key is not present, FormValue returns the empty string.
// To access multiple values of the same key, call ParseForm and
// then inspect Request.Form directly.
func (r *Request) FormValue(key string) string

// PostFormValue returns the first value for the named component of the POST,
// PATCH, or PUT request body. URL query parameters are ignored.
// PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores
// any errors returned by these functions.
// If key is not present, PostFormValue returns the empty string.
func (r *Request) PostFormValue(key string) string

// FormFile returns the first file for the provided form key.
// FormFile calls ParseMultipartForm and ParseForm if necessary.
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
