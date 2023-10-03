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
// * 「Authorization」、「WWW-Authenticate」、「Cookie」などの機密性の高いヘッダーを、信頼できないターゲットに転送する場合。
// これらのヘッダーは、初期ドメインのサブドメインマッチまたは完全一致ではないドメインにリダイレクトする場合には無視されます。
// たとえば、「foo.com」から「foo.com」または「sub.foo.com」にリダイレクトする場合、機密性の高いヘッダーが転送されますが、「bar.com」にリダイレクトする場合は転送されません。
//
// * 非nilのCookie Jarで「Cookie」ヘッダーを転送する場合。
// 各リダイレクトは、Cookie Jarの状態を変更する可能性があるため、初期リクエストで設定されたCookieを変更する可能性があります。
// 「Cookie」ヘッダーを転送する場合、変更されたCookieは省略され、Jarが更新された値でこれらの変更されたCookieを挿入することが期待されます(元の値が一致する場合)。
// Jarがnilの場合、初期Cookieは変更せずに転送されます。
type Client struct {
	Transport RoundTripper

	CheckRedirect func(req *Request, via []*Request) error

	Jar CookieJar

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
// リクエストBodyがnilでない場合、下層のTransportによってクローズされます。エラーが発生した場合でも同様です。
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

// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
//  1. On Read error or close, the stop func is called.
//  2. On Read failure, if reqDidTimeout is true, the error is wrapped and
//     marked as net.Error that hit its timeout.
