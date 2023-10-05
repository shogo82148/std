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

	// The following fields are for requests matched by ServeMux.
	pat         *pattern
	matches     []string
	otherValues map[string]string
}

// Contextは、リクエストのコンテキストを返します。コンテキストを変更するには、CloneまたはWithContextを使用してください。
//
// 返されるコンテキストは常にnilではありません。デフォルトでは、バックグラウンドコンテキストになります。
//
// 出力クライアントリクエストの場合、コンテキストはキャンセルを制御します。
//
// 入力サーバーリクエストの場合、クライアントの接続が閉じられたとき、リクエストがキャンセルされたとき（HTTP/2で）、またはServeHTTPメソッドが返されたときに、コンテキストがキャンセルされます。
func (r *Request) Context() context.Context

// WithContextは、そのコンテキストをctxに変更したrの浅いコピーを返します。提供されたctxはnilであってはなりません。
//
// 出力クライアントリクエストの場合、コンテキストはリクエストとそのレスポンスのライフタイム全体を制御します：接続の取得、リクエストの送信、レスポンスヘッダーとボディの読み取り。
//
// コンテキストを持つ新しいリクエストを作成するには、NewRequestWithContextを使用します。
// 新しいコンテキストを持つリクエストのディープコピーを作成するには、Request.Cloneを使用します。
func (r *Request) WithContext(ctx context.Context) *Request

// Cloneは、そのコンテキストをctxに変更したrのディープコピーを返します。提供されたctxはnilであってはなりません。
//
// 出力クライアントリクエストの場合、コンテキストはリクエストとそのレスポンスのライフタイム全体を制御します：接続の取得、リクエストの送信、レスポンスヘッダーとボディの読み取り。
//
// 新しいコンテキストを持つリクエストを作成するには、NewRequestWithContextを使用します。
// コンテキストを変更せずに新しいリクエストを作成するには、Request.WithContextを使用します。
func (r *Request) Clone(ctx context.Context) *Request

// ProtoAtLeastは、リクエストで使用されるHTTPプロトコルがmajor.minor以上であるかどうかを報告します。
func (r *Request) ProtoAtLeast(major, minor int) bool

// UserAgentは、リクエストで送信された場合にクライアントのUser-Agentを返します。
func (r *Request) UserAgent() string

// Cookiesは、リクエストで送信されたHTTPクッキーを解析して返します。
func (r *Request) Cookies() []*Cookie

// ErrNoCookieは、Cookieメソッドがクッキーを見つけられなかった場合にRequestによって返されます。
var ErrNoCookie = errors.New("http: named cookie not present")

// Cookieは、リクエストで提供された名前付きクッキーを返します。クッキーが見つからない場合はErrNoCookieを返します。
// 複数のクッキーが指定された名前に一致する場合、1つのクッキーのみが返されます。
func (r *Request) Cookie(name string) (*Cookie, error)

// AddCookieは、リクエストにクッキーを追加します。RFC 6265セクション5.4に従い、
// AddCookieは1つ以上のCookieヘッダーフィールドを添付しません。つまり、すべてのクッキーが、
// セミコロンで区切られた同じ行に書き込まれます。
// AddCookieは、cの名前と値をサニタイズするだけで、すでにリクエストに存在するCookieヘッダーをサニタイズしません。
func (r *Request) AddCookie(c *Cookie)

// Refererは、リクエストで送信された場合に参照元のURLを返します。
//
// Refererは、HTTPの初期の日々からの誤りで、リクエスト自体で誤ってスペルがされています。
// この値はHeader["Referer"]としてHeaderマップから取得することもできますが、
// メソッドとして利用可能にすることの利点は、代替の（正しい英語の）スペルreq.Referrer()を使用するプログラムをコンパイラが診断できるが、
// Header["Referrer"]を使用するプログラムを診断できないことです。
func (r *Request) Referer() string

// MultipartReaderは、これがmultipart/form-dataまたはmultipart/mixed POSTリクエストである場合、MIMEマルチパートリーダーを返します。
// それ以外の場合はnilとエラーを返します。
// リクエストボディをストリームとして処理するために、ParseMultipartFormの代わりにこの関数を使用してください。
func (r *Request) MultipartReader() (*multipart.Reader, error)

// Writeは、ワイヤフォーマットでHTTP/1.1リクエスト（ヘッダーとボディ）を書き込みます。
// このメソッドは、リクエストの以下のフィールドを参照します。
//
//	Host
//	URL
//	Method（デフォルトは "GET"）
//	Header
//	ContentLength
//	TransferEncoding
//	Body
//
// Bodyが存在し、Content-Lengthが0以下であり、TransferEncodingが "identity"に設定されていない場合、
// Writeはヘッダーに "Transfer-Encoding: chunked"を追加します。Bodyは送信後に閉じられます。
func (r *Request) Write(w io.Writer) error

// WriteProxyは、Writeと似ていますが、HTTPプロキシが期待する形式でリクエストを書き込みます。
// 特に、WriteProxyは、RFC 7230のセクション5.3に従って、スキームとホストを含む絶対URIでリクエストの最初のRequest-URI行を書き込みます。
// WriteProxyは、r.Host または r.URL.Host を使用して、Hostヘッダーも書き込みます。
func (r *Request) WriteProxy(w io.Writer) error

// ParseHTTPVersionは、RFC 7230、セクション2.6に従ってHTTPバージョン文字列を解析します。
// "HTTP/1.0"は(1, 0, true)を返します。注意："HTTP/2"のようにマイナーバージョンがない文字列は無効です。
func ParseHTTPVersion(vers string) (major, minor int, ok bool)

// NewRequestWithContextは、context.Backgroundを使用してNewRequestWithContextをラップします。
func NewRequest(method, url string, body io.Reader) (*Request, error)

// NewRequestWithContextは、メソッド、URL、およびオプションのボディが与えられた場合に新しいRequestを返します。
//
<<<<<<< HEAD
// If the provided body is also an io.Closer, the returned
// Request.Body is set to body and will be closed (possibly
// asynchronously) by the Client methods Do, Post, and PostForm,
// and Transport.RoundTrip.
=======
// 提供されたボディがio.Closerでもある場合、返されたRequest.Bodyはbodyに設定され、ClientメソッドDo、Post、およびPostForm、およびTransport.RoundTripによって閉じられます。
>>>>>>> release-branch.go1.21
//
// NewRequestWithContextは、Client.Do または Transport.RoundTrip で使用するためのRequestを返します。
// Server Handlerをテストするためのリクエストを作成するには、net/http/httptestパッケージのNewRequest関数を使用するか、
// ReadRequestを使用するか、またはRequestフィールドを手動で更新します。送信元のクライアントリクエストの場合、
// コンテキストはリクエストとその応答の全寿命を制御します：接続の取得、リクエストの送信、および応答ヘッダーとボディの読み取り。
// 入力リクエストフィールドと出力リクエストフィールドの違いについては、Requestタイプのドキュメントを参照してください。
//
// bodyが *bytes.Buffer、 *bytes.Reader、または *strings.Readerの場合、返されたリクエストのContentLengthはその正確な値に設定されます（-1の代わりに）、
// GetBodyが作成されます（307および308のリダイレクトがボディを再生できるように）、およびContentLengthが0の場合はBodyがNoBodyに設定されます。
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)

// BasicAuthは、リクエストがHTTP Basic認証を使用する場合、リクエストのAuthorizationヘッダーで提供されるユーザー名とパスワードを返します。
// RFC 2617、セクション2を参照してください。
func (r *Request) BasicAuth() (username, password string, ok bool)

// SetBasicAuthは、提供されたユーザー名とパスワードを使用して、HTTP Basic認証を使用するようにリクエストのAuthorizationヘッダーを設定します。
//
// HTTP Basic認証では、提供されたユーザー名とパスワードは暗号化されません。 HTTPSリクエストでのみ使用することが一般的です。
//
// ユーザー名にはコロンを含めることはできません。一部のプロトコルでは、ユーザー名とパスワードを事前にエスケープする追加の要件がある場合があります。たとえば、OAuth2と一緒に使用する場合、両方の引数を最初にurl.QueryEscapeでURLエンコードする必要があります。
func (r *Request) SetBasicAuth(username, password string)

// ReadRequestは、bから受信したリクエストを読み取り、解析します。
//
// ReadRequestは、低レベルの関数であり、特殊なアプリケーションにのみ使用する必要があります。ほとんどのコードは、Serverを使用してリクエストを読み取り、Handlerインターフェイスを介して処理する必要があります。 ReadRequestは、HTTP / 1.xリクエストのみをサポートしています。 HTTP / 2の場合は、golang.org/x/net/http2を使用してください。
func ReadRequest(b *bufio.Reader) (*Request, error)

// MaxBytesReaderは、io.LimitReaderに似ていますが、着信リクエストボディのサイズを制限するために使用されます。
// io.LimitReaderとは異なり、MaxBytesReaderの結果はReadCloserであり、制限を超えたReadに対して *MaxBytesError 型の非nilエラーを返し、Closeメソッドが呼び出されたときに基になるリーダーを閉じます。
//
// MaxBytesReaderは、クライアントが誤ってまたは悪意を持って大きなリクエストを送信してサーバーのリソースを浪費することを防止します。可能であれば、ResponseWriterに制限に達した後に接続を閉じるように指示します。
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser

// MaxBytesErrorは、MaxBytesReaderの読み取り制限を超えた場合にMaxBytesReaderによって返されます。
type MaxBytesError struct {
	Limit int64
}

func (e *MaxBytesError) Error() string

// ParseFormはr.Formとr.PostFormを埋めます。
//
// すべてのリクエストに対して、ParseFormはURLから生のクエリを解析し、r.Formを更新します。
//
// POST、PUT、およびPATCHリクエストの場合、それはまた、リクエストボディを読み取り、フォームとして解析し、その結果をr.PostFormとr.Formの両方に入れます。リクエストボディのパラメータは、r.FormのURLクエリ文字列値より優先されます。
//
// リクエストボディのサイズがすでにMaxBytesReaderによって制限されていない場合、サイズは10MBに制限されます。
//
// 他のHTTPメソッド、またはContent-Typeがapplication/x-www-form-urlencodedでない場合、リクエストボディは読み取られず、r.PostFormはnilでない空の値に初期化されます。
//
// ParseMultipartFormは自動的にParseFormを呼び出します。
// ParseFormは冪等です。
func (r *Request) ParseForm() error

// ParseMultipartFormは、リクエストボディをmultipart/form-dataとして解析します。
// リクエストボディ全体が解析され、そのファイルパーツの最大maxMemoryバイトがメモリに格納され、残りは一時ファイルに格納されます。
// ParseMultipartFormは必要に応じてParseFormを呼び出します。
// ParseFormがエラーを返した場合、ParseMultipartFormはそれを返しますが、リクエストボディの解析を続けます。
// ParseMultipartFormを1回呼び出した後、以降の呼び出しは効果がありません。
func (r *Request) ParseMultipartForm(maxMemory int64) error

<<<<<<< HEAD
// FormValue returns the first value for the named component of the query.
// POST, PUT, and PATCH body parameters take precedence over URL query string values.
// FormValue calls ParseMultipartForm and ParseForm if necessary and ignores
// any errors returned by these functions.
// If key is not present, FormValue returns the empty string.
// To access multiple values of the same key, call ParseForm and
// then inspect Request.Form directly.
func (r *Request) FormValue(key string) string

// PostFormValue returns the first value for the named component of the POST,
// PUT, or PATCH request body. URL query parameters are ignored.
// PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores
// any errors returned by these functions.
// If key is not present, PostFormValue returns the empty string.
=======
// FormValueは、クエリの名前付きコンポーネントの最初の値を返します。
// POSTおよびPUTボディパラメータは、URLクエリ文字列値より優先されます。
// FormValueは必要に応じてParseMultipartFormおよびParseFormを呼び出し、これらの関数によって返されたエラーを無視します。
// キーが存在しない場合、FormValueは空の文字列を返します。
// 同じキーの複数の値にアクセスするには、ParseFormを呼び出して、その後Request.Formを直接調べます。
func (r *Request) FormValue(key string) string

// PostFormValueは、POST、PATCH、またはPUTリクエストボディの名前付きコンポーネントの最初の値を返します。
// URLクエリパラメータは無視されます。
// PostFormValueは必要に応じてParseMultipartFormおよびParseFormを呼び出し、これらの関数によって返されたエラーを無視します。
// キーが存在しない場合、PostFormValueは空の文字列を返します。
>>>>>>> release-branch.go1.21
func (r *Request) PostFormValue(key string) string

// FormFileは、指定されたフォームキーの最初のファイルを返します。
// 必要に応じて、FormFileはParseMultipartFormおよびParseFormを呼び出します。
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)

// PathValue returns the value for the named path wildcard in the ServeMux pattern
// that matched the request.
// It returns the empty string if the request was not matched against a pattern
// or there is no such wildcard in the pattern.
func (r *Request) PathValue(name string) string

func (r *Request) SetPathValue(name, value string)
