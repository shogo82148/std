// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client. See RFC 7230 through 7235.
//
// This is the high-level Client interface.
// The low-level implementation is in transport.go.

package http

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"
)

// ClientはHTTPクライアントです。ゼロ値(DefaultClient)は、DefaultTransportを使用する使用可能なクライアントです。
//
// ClientのTransportには通常、内部状態(キャッシュされたTCP接続など)があるため、必要に応じて作成するのではなく、再利用する必要があります。
// Clientsは、複数のゴルーチンによる同時使用に対して安全です。
//
// Clientは、RoundTripper(Transportなど)よりも高レベルであり、クッキーやリダイレクトなどのHTTPの詳細も処理します。
//
// リダイレクトに従う場合、Clientは、初期リクエストに設定されたすべてのヘッダーを転送しますが、以下の場合は除外されます。
//
<<<<<<< HEAD
//   - when forwarding sensitive headers like "Authorization",
//     "WWW-Authenticate", and "Cookie" to untrusted targets.
//     These headers will be ignored when following a redirect to a domain
//     that is not a subdomain match or exact match of the initial domain.
//     For example, a redirect from "foo.com" to either "foo.com" or "sub.foo.com"
//     will forward the sensitive headers, but a redirect to "bar.com" will not.
//   - when forwarding the "Cookie" header with a non-nil cookie Jar.
//     Since each redirect may mutate the state of the cookie jar,
//     a redirect may possibly alter a cookie set in the initial request.
//     When forwarding the "Cookie" header, any mutated cookies will be omitted,
//     with the expectation that the Jar will insert those mutated cookies
//     with the updated values (assuming the origin matches).
//     If Jar is nil, the initial cookies are forwarded without change.
=======
// * 「Authorization」、「WWW-Authenticate」、「Cookie」などの機密性の高いヘッダーを、信頼できないターゲットに転送する場合。
// これらのヘッダーは、初期ドメインのサブドメインマッチまたは完全一致ではないドメインにリダイレクトする場合には無視されます。
// たとえば、「foo.com」から「foo.com」または「sub.foo.com」にリダイレクトする場合、機密性の高いヘッダーが転送されますが、「bar.com」にリダイレクトする場合は転送されません。
//
// * 非nilのCookie Jarで「Cookie」ヘッダーを転送する場合。
// 各リダイレクトは、Cookie Jarの状態を変更する可能性があるため、初期リクエストで設定されたCookieを変更する可能性があります。
// 「Cookie」ヘッダーを転送する場合、変更されたCookieは省略され、Jarが更新された値でこれらの変更されたCookieを挿入することが期待されます(元の値が一致する場合)。
// Jarがnilの場合、初期Cookieは変更せずに転送されます。
>>>>>>> release-branch.go1.21
type Client struct {
	// Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	Transport RoundTripper

	// CheckRedirect specifies the policy for handling redirects.
	// If CheckRedirect is not nil, the client calls it before
	// following an HTTP redirect. The arguments req and via are
	// the upcoming request and the requests made already, oldest
	// first. If CheckRedirect returns an error, the Client's Get
	// method returns both the previous Response (with its Body
	// closed) and CheckRedirect's error (wrapped in a url.Error)
	// instead of issuing the Request req.
	// As a special case, if CheckRedirect returns ErrUseLastResponse,
	// then the most recent response is returned with its body
	// unclosed, along with a nil error.
	//
	// If CheckRedirect is nil, the Client uses its default policy,
	// which is to stop after 10 consecutive requests.
	CheckRedirect func(req *Request, via []*Request) error

	// Jar specifies the cookie jar.
	//
	// The Jar is used to insert relevant cookies into every
	// outbound Request and is updated with the cookie values
	// of every inbound Response. The Jar is consulted for every
	// redirect that the Client follows.
	//
	// If Jar is nil, cookies are only sent if they are explicitly
	// set on the Request.
	Jar CookieJar

	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	//
	// The Client cancels requests to the underlying Transport
	// as if the Request's Context ended.
	//
	// For compatibility, the Client will also use the deprecated
	// CancelRequest method on Transport if found. New
	// RoundTripper implementations should use the Request's Context
	// for cancellation instead of implementing CancelRequest.
	Timeout time.Duration
}

// DefaultClientは、デフォルトのClientであり、Get、Head、およびPostで使用されます。
var DefaultClient = &Client{}

// RoundTripperは、指定されたRequestに対するResponseを取得するための単一のHTTPトランザクションを実行する能力を表すインターフェースです。
//
// RoundTripperは、複数のゴルーチンによる同時使用に対して安全である必要があります。
type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
}

// ErrSchemeMismatchは、サーバーがHTTPSクライアントにHTTPレスポンスを返した場合に返されます。
var ErrSchemeMismatch = errors.New("http: server gave HTTP response to HTTPS client")

// Getは、指定されたURLにGETを発行します。レスポンスが次のリダイレクトコードの1つである場合、
// Getはリダイレクトに従います。最大10回のリダイレクトまで:
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// リダイレクトが多すぎる場合や、HTTPプロトコルエラーがあった場合はエラーが返されます。
// 非2xxレスポンスはエラーを引き起こしません。
// 任意の返されたエラーは*url.Error型です。url.ErrorのTimeoutメソッドは、
// リクエストがタイムアウトした場合にtrueを報告します。
//
// errがnilの場合、respには常に非nilのresp.Bodyが含まれます。
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
//
// Getは、DefaultClient.Getのラッパーです。
//
// カスタムヘッダーでリクエストを作成するには、NewRequest と DefaultClient.Do を使用します。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContext と DefaultClient.Do を使用します。
func Get(url string) (resp *Response, err error)

// Getは、指定されたURLにGETを発行します。レスポンスが次のリダイレクトコードの1つである場合、
// Getはリダイレクトに従います。最大10回のリダイレクトまで:
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// リダイレクトが多すぎる場合や、HTTPプロトコルエラーがあった場合はエラーが返されます。
// 非2xxレスポンスはエラーを引き起こしません。
// 任意の返されたエラーは*url.Error型です。url.ErrorのTimeoutメソッドは、
// リクエストがタイムアウトした場合にtrueを報告します。
//
// errがnilの場合、respには常に非nilのresp.Bodyが含まれます。
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
//
// カスタムヘッダーでリクエストを作成するには、NewRequestとClient.Doを使用します。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとClient.Doを使用します。
func (c *Client) Get(url string) (resp *Response, err error)

// ErrUseLastResponseは、Client.CheckRedirectフックによって返されることがあります。
// リダイレクトの処理方法を制御するために使用されます。返された場合、次のリクエストは送信されず、
// 最新のレスポンスがそのボディが閉じられていないまま返されます。
var ErrUseLastResponse = errors.New("net/http: use last response")

// DoはHTTPリクエストを送信し、クライアントで設定されたポリシー(リダイレクト、クッキー、認証など)に従ってHTTPレスポンスを返します。
//
// クライアントポリシー(例えばCheckRedirect)によって引き起こされた場合、またはHTTPの送信に失敗した場合(ネットワーク接続の問題など)、エラーが返されます。非2xxステータスコードはエラーを引き起こしません。
//
// 返されたエラーがnilの場合、Responseにはユーザーが閉じる必要のある非nilのBodyが含まれます。BodyがEOFまで読み取られずに閉じられていない場合、Clientの基礎となるRoundTripper(通常はTransport)は、次の「keep-alive」リクエストのためにサーバーへの永続的なTCP接続を再利用できなくなる可能性があります。
//
<<<<<<< HEAD
// The request Body, if non-nil, will be closed by the underlying
// Transport, even on errors. The Body may be closed asynchronously after
// Do returns.
=======
// リクエストBodyがnilでない場合、下層のTransportによってクローズされます。エラーが発生した場合でも同様です。
>>>>>>> release-branch.go1.21
//
// エラーが発生した場合、任意のResponseは無視できます。非nilのResponseと非nilのエラーが返されるのは、CheckRedirectが失敗した場合だけであり、その場合でも返されたResponse.Bodyは既に閉じられています。
//
// 通常、Doの代わりにGet、Post、またはPostFormが使用されます。
//
// サーバーがリダイレクトで応答した場合、Clientは最初にCheckRedirect関数を使用して、リダイレクトをフォローするかどうかを決定します。許可された場合、301、302、または303のリダイレクトは、HTTPメソッドGET(または元のリクエストがHEADの場合はHEAD)を使用して、ボディなしで後続のリクエストを引き起こします。307または308のリダイレクトは、Request.GetBody関数が定義されている場合、元のHTTPメソッドとボディを保持します。NewRequest関数は、一般的な標準ライブラリのボディタイプに対してGetBodyを自動的に設定します。
//
// 返されるエラーはすべて*url.Error型です。url.ErrorのTimeoutメソッドは、リクエストがタイムアウトした場合にtrueを報告します。
func (c *Client) Do(req *Request) (*Response, error)

// Postは、指定されたURLにPOSTを発行します。
//
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
//
// 提供されたBodyがio.Closerである場合、リクエストの後に閉じられます。
//
// Postは、DefaultClient.Postのラッパーです。
//
// カスタムヘッダーを設定するには、NewRequestとDefaultClient.Doを使用します。
//
// リダイレクトの処理方法については、Client.Doメソッドのドキュメントを参照してください。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
func Post(url, contentType string, body io.Reader) (resp *Response, err error)

// Postは、指定されたURLにPOSTを発行します。
//
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
//
// 提供されたBodyがio.Closerである場合、リクエストの後に閉じられます。
//
// カスタムヘッダーを設定するには、NewRequestとClient.Doを使用します。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとClient.Doを使用します。
//
// リダイレクトの処理方法については、Client.Doメソッドのドキュメントを参照してください。
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)

// PostFormは、データのキーと値がURLエンコードされたリクエストボディとして指定されたURLにPOSTを発行します。
//
// Content-Typeヘッダーはapplication/x-www-form-urlencodedに設定されます。
// 他のヘッダーを設定するには、NewRequestとDefaultClient.Doを使用します。
//
// errがnilの場合、respには常に非nilのresp.Bodyが含まれます。
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
//
// PostFormは、DefaultClient.PostFormのラッパーです。
//
// リダイレクトの処理方法については、Client.Doメソッドのドキュメントを参照してください。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
func PostForm(url string, data url.Values) (resp *Response, err error)

// PostFormは、データのキーと値がURLエンコードされたリクエストボディとして指定されたURLにPOSTを発行します。
//
// Content-Typeヘッダーはapplication/x-www-form-urlencodedに設定されます。
// 他のヘッダーを設定するには、NewRequestとClient.Doを使用します。
//
// errがnilの場合、respには常に非nilのresp.Bodyが含まれます。
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
//
// リダイレクトの処理方法については、Client.Doメソッドのドキュメントを参照してください。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとClient.Doを使用します。
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)

// Headは、指定されたURLにHEADを発行します。レスポンスが次のリダイレクトコードの1つである場合、
// Headはリダイレクトに従います。最大10回のリダイレクトまで:
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// Headは、DefaultClient.Headのラッパーです。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
func Head(url string) (resp *Response, err error)

// 指定されたURLにHEADを発行します。レスポンスが以下のリダイレクトコードのいずれかである場合、
// HeadはClientのCheckRedirect関数を呼び出した後にリダイレクトに従います。
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// 指定されたcontext.Contextでリクエストを行うには、NewRequestWithContextとClient.Doを使用してください。
func (c *Client) Head(url string) (resp *Response, err error)

// CloseIdleConnectionsは、以前のリクエストから接続されていたが現在は「keep-alive」状態でアイドル状態にある
// Transport上の接続を閉じます。現在使用中の接続は中断しません。
//
// ClientのTransportにCloseIdleConnectionsメソッドがない場合、このメソッドは何もしません。
func (c *Client) CloseIdleConnections()
