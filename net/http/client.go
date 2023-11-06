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

// DefaultClientはデフォルトのClientであり、Get、Head、およびPostに使用されます。
var DefaultClient = &Client{}

// RoundTripperは、指定されたRequestに対するResponseを取得するための単一のHTTPトランザクションを実行する能力を表すインターフェースです。
//
// RoundTripperは、複数のゴルーチンによる同時使用に対して安全である必要があります。
type RoundTripper interface {
	// RoundTripは単一のHTTPトランザクションを実行し、
	// 提供されたRequestのResponseを返します。
	//
	// RoundTripはレスポンスを解釈しようとしてはなりません。特に、
	// RoundTripはレスポンスを取得した場合、レスポンスのHTTPステータスコードに関係なく
	// err == nilを返さなければなりません。非nilのerrはレスポンスの取得に失敗した場合に
	// 予約されるべきです。同様に、RoundTripはリダイレクト、認証、またはクッキーなどの
	// 高レベルのプロトコル詳細を処理しようとしてはなりません。
	//
	// RoundTripは、リクエストのBodyを消費して閉じる以外に、
	// リクエストを変更してはなりません。RoundTripは、
	// 別のgoroutineでリクエストのフィールドを読むことができます。呼び出し元は、
	// レスポンスのBodyが閉じられるまで、リクエストを変更したり再利用したりしてはなりません。
	//
	// RoundTripは常にボディを閉じる必要があります。これにはエラー時も含まれますが、
	// 実装によっては、RoundTripが返った後でも別のgoroutineで行うことがあります。
	// これは、後続のリクエストのボディを再利用したい呼び出し元は、
	// それを行う前にClose呼び出しを待つように手配する必要があることを意味します。
	//
	// RequestのURLとHeaderフィールドは初期化されていなければなりません。
	RoundTrip(*Request) (*Response, error)
}

// ErrSchemeMismatchは、サーバーがHTTPSクライアントにHTTPレスポンスを返した場合に返されます。
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

// Getは指定されたURLにGETを発行します。レスポンスが以下のリダイレクトコードのいずれかである場合、
// Getはリダイレクトをフォローします。最大10回のリダイレクトが許可されます：
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// リダイレクトが多すぎる場合やHTTPプロトコルエラーがあった場合にエラーが返されます。
// 2xx以外のレスポンスはエラーを引き起こしません。
// 返される任意のエラーは*url.Error型になります。
// url.Error値のTimeoutメソッドは、リクエストがタイムアウトした場合にtrueを報告します。
//
// errがnilの場合、respは常に非nilのresp.Bodyを含みます。
// 読み取りが完了したら、呼び出し元はresp.Bodyを閉じる必要があります。
//
// GetはDefaultClient.Getのラッパーです。
//
// カスタムヘッダーでリクエストを行うには、NewRequestと
// DefaultClient.Doを使用します。
//
// 指定したcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
func (c *Client) Get(url string) (resp *Response, err error)

// ErrUseLastResponseは、Client.CheckRedirectフックによって返され、リダイレクトの処理方法を制御するために使用できます。
// これが返されると、次のリクエストは送信されず、最も最近のレスポンスがそのボディが閉じられていない状態で返されます。
var ErrUseLastResponse = errors.New("net/http: use last response")

// DoはHTTPリクエストを送信し、クライアントの設定したポリシー（リダイレクト、クッキー、認証など）に従ってHTTPレスポンスを返します。
//
// クライアントポリシー（CheckRedirectなど）によってトリガーされた場合、またはHTTP転送が失敗した場合（ネットワーク接続の問題など）、エラーが返されます。2xx以外のステータスコードはエラーを引き起こしません。
//
// 返されるエラーがnilの場合、Responseにはユーザーが閉じなければならない非nilのBodyが含まれます。BodyがEOFまで完全に読み取られずに閉じられない場合、クライアントの基本となるRoundTripper（通常はTransport）は、次の「keep-alive」リクエストのためにサーバーへの永続的なTCP接続を再利用できないかもしれません。
//
// リクエストのBodyがnilでない場合、それは基本となるTransportによって閉じられます。エラーが発生した場合も同様です。
//
// エラーが発生した場合、任意のResponseは無視できます。非nilのResponseと非nilのエラーは、CheckRedirectが失敗した場合にのみ返されます。その場合でも、返されたResponse.Bodyはすでに閉じられています。
//
// 通常、Doの代わりにGet、Post、またはPostFormが使用されます。
//
// サーバーがリダイレクトを返すと、Clientは最初にCheckRedirect関数を使用してリダイレクトをフォローするかどうかを決定します。許可されると、301、302、または303のリダイレクトは、メソッドがGET（元のリクエストがHEADだった場合はHEAD）でボディがない後続のリクエストを引き起こします。307または308のリダイレクトは、Request.GetBody関数が定義されている場合、元のHTTPメソッドとボディを保持します。NewRequest関数は、一般的な標準ライブラリボディタイプのGetBodyを自動的に設定します。
//
// すべての返されるエラーは*url.Error型です。url.ErrorのTimeoutメソッドは、リクエストがタイムアウトした場合にtrueを報告します。
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

// PostForm関数は、指定されたデータのキーと値のペアをリクエストボディにエンコードして、URLにPOSTリクエストを送信します。
//
// Content-Typeヘッダーはapplication/x-www-form-urlencodedに設定されます。
// 他のヘッダーを設定するには、NewRequestとClient.Doを使用します。
//
// errがnilの場合、respは常に非nilのresp.Bodyを含みます。
// 読み取り後、呼び出し元はresp.Bodyを閉じなければなりません。
//
// リダイレクトの処理については、Client.Doメソッドのドキュメンテーションを参照してください。
//
// 指定したcontext.Contextでリクエストを作成するには、NewRequestWithContextとClient.Doを使用します。
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)

// Headは、指定されたURLに対してHEADリクエストを発行します。レスポンスが以下のリダイレクトコードのいずれかである場合、
// Headはリダイレクションをフォローします。最大10回のリダイレクトが許可されます：
//
//	301 (Moved Permanently)
//	302 (Found)
//	303 (See Other)
//	307 (Temporary Redirect)
//	308 (Permanent Redirect)
//
// HeadはDefaultClient.Headのラッパーです。
//
// 指定したcontext.Contextでリクエストを作成するには、NewRequestWithContextとDefaultClient.Doを使用します。
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

// CloseIdleConnectionsは、以前は接続されていたが現在は"keep-alive"状態、つまりアイドル状態にあるTransport上のアイドル接続を閉じます。
// これは現在アクティブな接続を中断しません。
//
// クライアントのTransportにCloseIdleConnectionsメソッドがない場合、このメソッドは何もしません。
func (c *Client) CloseIdleConnections()
