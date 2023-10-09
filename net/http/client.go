// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTPクライアント。RFC 7230から7235を参照してください。
//
// これは高レベルのClientインターフェースです。
// 低レベルの実装はtransport.goにあります。

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
// * 「Authorization」、「WWW-Authenticate」、「Cookie」などの機密性の高いヘッダーを、信頼できないターゲットに転送する場合。
// これらのヘッダーは、初期ドメインのサブドメインマッチまたは完全一致ではないドメインにリダイレクトする場合には無視されます。
// たとえば、「foo.com」から「foo.com」または「sub.foo.com」にリダイレクトする場合、機密性の高いヘッダーが転送されますが、「bar.com」にリダイレクトする場合は転送されません。
//
// * 非nilのCookie Jarで「Cookie」ヘッダーを転送する場合。
// 各リダイレクトは、Cookie Jarの状態を変更する可能性があるため、初期リクエストで設定されたCookieを変更する可能性があります。
// 「Cookie」ヘッダーを転送する場合、変更されたCookieは省略され、Jarが更新された値でこれらの変更されたCookieを挿入することが期待されます(元の値が一致する場合)。
// Jarがnilの場合、初期Cookieは変更せずに転送されます。
type Client struct {
	// Transportは個別のHTTPリクエストが行われるメカニズムを指定します。
	// nilの場合、DefaultTransportが使用されます。
	Transport RoundTripper

	// CheckRedirectはリダイレクトの処理ポリシーを指定します。
	// CheckRedirectがnilでない場合、クライアントはHTTPリダイレクトに従う前にそれを呼び出します。
	// 引数のreqとviaは、直前のリクエストとこれまでに行われたリクエストです。最も古いものから順に渡されます。
	// CheckRedirectがエラーを返す場合、ClientのGetメソッドはリクエストreqを発行せずに、
	// 前のResponse（そのBodyが閉じられている）とCheckRedirectのエラー（url.Errorでラップされたもの）の両方を返します。
	// 特別な場合として、CheckRedirectがErrUseLastResponseを返す場合、
	// 最新のレスポンスがそのBodyが閉じていない状態で返され、エラーはnilです。
	//
	// CheckRedirectがnilの場合、Clientはデフォルトのポリシーを使用します。
	// デフォルトのポリシーは、連続した10個のリクエストの後に停止することです。
	CheckRedirect func(req *Request, via []*Request) error

	// Jarはクッキージャーを指定します。
	//
	// Jarは、出力リクエストに関連するクッキーを挿入するために使用され、
	// すべての入力レスポンスのクッキー値で更新されます。
	// Jarは、クライアントがフォローするすべてのリダイレクトで参照されます。
	//
	// Jarがnilの場合、リクエストに明示的に設定されていない場合にのみクッキーが送信されます。
	Jar CookieJar

	// Timeoutは、このClientによって行われるリクエストのタイムアウト時間を指定します。
	// タイムアウトには、接続時間、リダイレクト、レスポンスボディの読み取りなどが含まれます。
	// Get、Head、Post、またはDoが返却された後でも、タイマーは実行し続け、
	// Response.Bodyの読み取りを中断します。
	//
	// Timeoutがゼロの場合は、タイムアウトはありません。
	//
	// Clientは、RequestのContextが終了した場合と同様に、
	// ベースとなるTransportへのリクエストをキャンセルします。
	//
	// 互換性のために、ClientはTransportにCancelRequestメソッドも使用しますが、
	// 新しいRoundTripperの実装はCancelRequestではなく、
	// リクエストのContextを使用してキャンセルするべきです。
	Timeout time.Duration
}

// DefaultClient is the default Client and is used for Get, Head, and Post.
var DefaultClient = &Client{}

// RoundTripperは、指定されたRequestに対するResponseを取得するための単一のHTTPトランザクションを実行する能力を表すインターフェースです。
//
// RoundTripperは、複数のゴルーチンによる同時使用に対して安全である必要があります。
// RoundTripperは、複数のゴルーチンによる同時使用に対して安全である必要があります。
type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
}

// ErrSchemeMismatchは、サーバーがHTTPSクライアントにHTTPレスポンスを返した場合に返されます。
// ErrSchemeMismatch is returned when a server returns an HTTP response to an HTTPS client.
var ErrSchemeMismatch = errors.New("http: server gave HTTP response to HTTPS client")

// Getは、指定されたURLにGETを発行します。レスポンスが次のリダイレクトコードの1つである場合、
// Getはリダイレクトに従います。最大10回のリダイレクトまで:
//
//	301（Moved Permanently）
//	302（Found）
//	303（See Other）
//	307（Temporary Redirect）
//	308（Permanent Redirect）
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
// カスタムヘッダーでリクエストを作成するには、NewRequestとDefaultClient.Doを使用します。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
func Get(url string) (resp *Response, err error)

// Getは、指定されたURLにGETを発行します。
// Get sends a GET request to the specified URL.
// レスポンスが次のリダイレクトコードの1つである場合、
// If the response is one of the following redirect codes,
// Getはリダイレクトに従います。最大10回のリダイレクトまで:
// Get follows the redirect up to a maximum of 10 times:
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// リダイレクトが多すぎる場合や、HTTPプロトコルエラーがあった場合はエラーが返されます。
// If there are too many redirects or if there is an HTTP protocol error, an error will be returned.
// 非2xxレスポンスはエラーを引き起こしません。
// Non-2xx responses do not cause an error.
// 任意の返されたエラーは*url.Error型です。url.ErrorのTimeoutメソッドは、
// Any returned error is of type *url.Error. The Timeout method of url.Error reports true if the request times out.
// errがnilの場合、respには常に非nilのresp.Bodyが含まれます。
// If err is nil, resp will always contain a non-nil resp.Body.
// 呼び出し元は、resp.Bodyの読み取りが完了したらresp.Bodyを閉じる必要があります。
// The caller must close resp.Body when reading from it is complete.
// カスタムヘッダーでリクエストを作成するには、NewRequestとClient.Doを使用します。
// To create a request with custom headers, use NewRequest and Client.Do.
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとClient.Doを使用します。
// To create a request with a specified context.Context, use NewRequestWithContext and Client.Do.
func (c *Client) Get(url string) (resp *Response, err error)

// ErrUseLastResponseは、Client.CheckRedirectフックによって返されることがあります。
// リダイレクトの処理方法を制御するために使用されます。返された場合、次のリクエストは送信されず、
// 最新のレスポンスがそのボディが閉じられていないまま返されます。
// ErrUseLastResponseはClient.CheckRedirectフックによって返されることがあります。
// これはリダイレクトの処理方法を制御するために使用されます。もしErrUseLastResponseが返された場合、次のリクエストは送信されず、
// 最新のレスポンスがそのまま返されます。その時、最新のレスポンスのボディは閉じられていないままとなります。
var ErrUseLastResponse = errors.New("net/http: use last response")

// Do sends an HTTP request and returns an HTTP response according to the client's set policies (such as redirect, cookies, authentication, etc).
//
// If triggered by a client policy (such as CheckRedirect), or if the HTTP transmission fails (due to network connection issues, etc), an error is returned. Non-2xx status codes do not cause an error.
//
// If the returned error is nil, the Response will contain a non-nil Body that the user must close. If the Body is not closed without being fully read until EOF, the underlying RoundTripper of the Client (usually the Transport) may be unable to reuse the persistent TCP connection to the server for the next "keep-alive" request.
//
// If the request Body is not nil, it will be closed by the underlying Transport. The same applies even in case of an error.
//
// If an error occurs, any Response can be ignored. A non-nil Response and non-nil error are only returned if CheckRedirect fails, and even in that case, the returned Response.Body is already closed.
//
// Normally, Get, Post, or PostForm are used instead of Do.
//
// If the server responds with a redirect, the Client first uses the CheckRedirect function to determine whether to follow the redirect. If allowed, a redirect of 301, 302, or 303 will cause a subsequent request with method GET (or HEAD if the original request was HEAD) and no body. A redirect of 307 or 308 will preserve the original HTTP method and body if the Request.GetBody function is defined. The NewRequest function automatically sets GetBody for common standard library body types.
//
// All returned errors are of type *url.Error. The Timeout method of url.Error reports true if the request timed out.
func (c *Client) Do(req *Request) (*Response, error)

// Postは、指定されたURLに対してPOSTメソッドを送信します。
//
// 呼び出し元は、resp.Bodyの読み込みが完了したらresp.Bodyを閉じる必要があります。
//
// もし提供されたBodyがio.Closerを実装している場合は、リクエストの後で閉じられます。
//
// Postは、DefaultClient.Postのラッパーです。
//
// カスタムヘッダーを設定するには、NewRequestとDefaultClient.Doを使用します。
//
// リダイレクトの処理方法については、Client.Doメソッドのドキュメントを参照してください。
//
// 指定されたcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
func Post(url, contentType string, body io.Reader) (resp *Response, err error)

// Postは、指定されたURLにPOSTリクエストを送信します。
//
// 呼び出し元は、resp.Bodyの読み込みが完了したらresp.Bodyを閉じる必要があります。
//
// 提供されたBodyがio.Closerインターフェースを実装している場合、リクエストの後に閉じられます。
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

// PostForm function sends a POST request to the URL with the specified data key-value pairs encoded in the request body.
//
// The Content-Type header is set to application/x-www-form-urlencoded.
// To set other headers, use NewRequest and Client.Do.
//
// If err is nil, resp always contains a non-nil resp.Body.
// The caller must close resp.Body after reading from it.
//
// Refer to the documentation of Client.Do method for handling redirects.
//
// To create a request with the specified context.Context, use NewRequestWithContext and Client.Do.
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)

// Head issues a HEAD request to the specified URL. If the response is one of the following redirect codes,
// Head will follow the redirection. Up to 10 redirects are allowed:
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// Head is a wrapper for DefaultClient.Head.
//
// To create a request with the specified context.Context, use NewRequestWithContext and DefaultClient.Do.
func Head(url string) (resp *Response, err error)

// 指定されたURLにHEADリクエストを送信します。もしレスポンスが以下のいずれかのリダイレクトコードである場合、
// Headメソッドはリダイレクトに従う前にClientのCheckRedirect関数を呼び出すことがあります。
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// 指定されたcontext.Contextを使用してリクエストを送信する場合は、NewRequestWithContextとClient.Doを使用してください。
func (c *Client) Head(url string) (resp *Response, err error)

// CloseIdleConnections closes idle connections on the Transport that were
// previously connected but are now in the "keep-alive" state, meaning they are
// idle. It does not interrupt the currently active connections.
//
// If the Client's Transport does not have the CloseIdleConnections method,
// this method does nothing.
func (c *Client) CloseIdleConnections()
