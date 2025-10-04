// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP逆プロキシハンドラ

package httputil

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"
)

// ProxyRequestは、[ReverseProxy] によって書き換えられるリクエストを含んでいます。
type ProxyRequest struct {
	In *http.Request

	// Outはプロキシに送信されるリクエストです。
	// Rewrite関数はこのリクエストを変更または置換する場合があります。
	// Rewriteが呼び出される前に、ホップバイホップのヘッダーはこのリクエストから削除されます。
	Out *http.Request
}

<<<<<<< HEAD
// SetURLは、ターゲットに指定されたスキーム、ホスト、およびベースパスに従って、アウトバウンドリクエストをルーティングします。
// もしターゲットのパスが"/base"であり、受信したリクエストが"/dir"である場合、ターゲットリクエストは"/base/dir"となります。
=======
// SetURL routes the outbound request to the scheme, host, and base path
// provided in target. If the target's path is "/base" and the incoming
// request was for "/dir", the target request will be for "/base/dir".
// To route requests without joining the incoming path,
// set r.Out.URL directly.
>>>>>>> upstream/release-branch.go1.25
//
// SetURLは、アウトバウンドのHostヘッダをターゲットのホストに合わせて書き換えます。
// インバウンドのリクエストのHostヘッダを保持するために（[NewSingleHostReverseProxy] のデフォルトの動作）：
//
//	rewriteFunc := func(r *httputil.ProxyRequest) {
//	    r.SetURL(url)
//	    r.Out.Host = r.In.Host
//	}
func (r *ProxyRequest) SetURL(target *url.URL)

// SetXForwardedは、出力リクエストのX-Forwarded-For、X-Forwarded-Host、およびX-Forwarded-Protoヘッダーを設定します。
//
// - X-Forwarded-Forヘッダーは、クライアントのIPアドレスに設定されます。
// - X-Forwarded-Hostヘッダーは、クライアントが要求したホスト名に設定されます。
// - X-Forwarded-Protoヘッダーは、入力リクエストがTLS対応の接続で行われたかどうかに応じて、「http」または「https」に設定されます。
//
// 出力リクエストに既存のX-Forwarded-Forヘッダーが含まれている場合、SetXForwardedはクライアントのIPアドレスを追加します。
// SetXForwardedを呼び出す前に、入力リクエストのX-Forwarded-Forヘッダー（Director関数を使用して [ReverseProxy] を使用している場合のデフォルト動作）をコピーして、
// 入力リクエストのX-Forwarded-Forヘッダーに追加します：
//
//	rewriteFunc := func(r *httputil.ProxyRequest) {
//	   r.Out.Header["X-Forwarded-For"] = r.In.Header["X-Forwarded-For"]
//	   r.SetXForwarded()
//	}
func (r *ProxyRequest) SetXForwarded()

// ReverseProxyは、受信したリクエストを別のサーバーに送信し、レスポンスをクライアントにプロキシするHTTPハンドラです。
//
<<<<<<< HEAD
// もし基礎となるトランスポートがClientTrace.Got1xxResponseをサポートしている場合、1xxのレスポンスはクライアントに転送されます。
=======
// 1xx responses are forwarded to the client if the underlying
// transport supports ClientTrace.Got1xxResponse.
//
// Hop-by-hop headers (see RFC 9110, section 7.6.1), including
// Connection, Proxy-Connection, Keep-Alive, Proxy-Authenticate,
// Proxy-Authorization, TE, Trailer, Transfer-Encoding, and Upgrade,
// are removed from client requests and backend responses.
// The Rewrite function may be used to add hop-by-hop headers to the request,
// and the ModifyResponse function may be used to remove them from the response.
>>>>>>> upstream/release-branch.go1.25
type ReverseProxy struct {

	// Rewriteは、リクエストを変更してTransportを使用して送信される新しいリクエストに変換する関数でなければなりません。
	// そのレスポンスは、元のクライアントに変更せずにコピーされます。
	// Rewriteは、戻る後に提供されたProxyRequestまたはその内容にアクセスしてはいけません。
	//
	// Forwarded、X-Forwarded、X-Forwarded-Host、およびX-Forwarded-Protoヘッダーは、
	// Rewriteが呼び出される前に送信リクエストから削除されます。また、ProxyRequest.SetXForwardedメソッドも参照してください。
	//
	// 解析できないクエリパラメータは、Rewriteが呼び出される前に送信リクエストから削除されます。
	// Rewrite関数は、インバウンドURLのRawQueryをアウトバウンドURLにコピーして、元のパラメータ文字列を保持することがあります。
	// 注意：これは、プロキシのクエリパラメータの解釈がダウンストリームサーバーと一致しない場合にセキュリティの問題を引き起こす可能性があります。
	//
	// RewriteまたはDirectorのいずれか一つのみ設定できます。
	Rewrite func(*ProxyRequest)

	// Director（ディレクター）は、リクエストを変更して新しいリクエストをTransport（トランスポート）を使用して送信します。そのレスポンスは、元のクライアントに変更せずにコピーされます。Directorは、戻った後に提供されたリクエストにアクセスしてはいけません。
	// デフォルトでは、X-Forwarded-ForヘッダーはクライアントのIPアドレスの値に設定されます。もし既にX-Forwarded-Forヘッダーが存在する場合、クライアントのIPは既存の値に追加されます。ただし、特殊なケースとして、リクエストのRequest.Headerマップにヘッダーが存在しているが、値がnilである場合（Director関数によって設定された場合など）、X-Forwarded-Forヘッダーは変更されません。
	// IPスプーフィングを防ぐために、クライアントまたは信頼できないプロキシから送られてきたプリエクスティングのX-Forwarded-Forヘッダーを削除するようにしてください。
	// ディレクターが戻った後にリクエストからホップバイホップヘッダーが削除されます。これにより、ディレクターが追加したヘッダーも削除される可能性があります。リクエストの変更を保持するためには、リライト関数を使用してください。
	// ディレクターが戻った後、リクエストのRequest.Formが設定されている場合は、解析できないクエリパラメータが送信先のリクエストから削除されます。
	// RewriteまたはDirectorのうち、最大1つが設定できます。
	Director func(*http.Request)

	// プロキシリクエストを実行するために使用されるトランスポートです。
	// nil の場合、http.DefaultTransport が使用されます。
	Transport http.RoundTripper

	// FlushIntervalは、レスポンスボディをクライアントにコピーする際のフラッシュ間隔を指定します。
	// ゼロの場合、定期的なフラッシュは行われません。
	// 負の値は、クライアントへの各書き込みの直後にすぐにフラッシュすることを意味します。
	// FlushIntervalは、ReverseProxyがストリーミングレスポンスとしてレスポンスを認識するか、またはContentLengthが-1の場合は無視されます。
	// このようなレスポンスの場合、書き込みはすぐにクライアントにフラッシュされます。
	FlushInterval time.Duration

	// ErrorLogは、リクエストをプロキシする際に発生したエラーのオプションのロガーを指定します。
	// nilの場合、ログはlogパッケージの標準ロガーを使用して行われます。
	ErrorLog *log.Logger

	// BufferPoolは、io.CopyBufferがHTTPのレスポンスボディをコピーする際に使用するバイトスライスを取得するためのオプションのバッファプールを指定します。
	BufferPool BufferPool

	// ModifyResponseはオプションの関数であり、バックエンドからのレスポンスを変更する役割を持ちます。
	// この関数は、バックエンドからのレスポンスがある場合に呼び出されます（HTTPのステータスコードに関係なく）。
	// バックエンドに到達できない場合は、オプションのErrorHandlerが呼び出され、ModifyResponseは呼び出されません。
	//
<<<<<<< HEAD
	// ModifyResponseがエラーを返す場合、それに対してErrorHandlerが呼び出されます。
	// ErrorHandlerがnilの場合は、デフォルトの実装が使用されます。
=======
	// Hop-by-hop headers are removed from the response before
	// calling ModifyResponse. ModifyResponse may need to remove
	// additional headers to fit its deployment model, such as Alt-Svc.
	//
	// If ModifyResponse returns an error, ErrorHandler is called
	// with its error value. If ErrorHandler is nil, its default
	// implementation is used.
>>>>>>> upstream/release-branch.go1.25
	ModifyResponse func(*http.Response) error

	// ErrorHandlerは、バックエンドに到達したエラーやModifyResponseからのエラーを処理するオプションの関数です。
	//
	// nilの場合、デフォルトでは提供されたエラーをログに記録し、502 Status Bad Gatewayレスポンスを返します。
	ErrorHandler func(http.ResponseWriter, *http.Request, error)
}

// BufferPoolは [io.CopyBuffer] で使用するための一時的なバイトスライスを取得および返却するためのインターフェースです。
type BufferPool interface {
	Get() []byte
	Put([]byte)
}

// NewSingleHostReverseProxyは、URLを指定されたスキーム、ホスト、およびベースパスにルーティングする新しい [ReverseProxy] を返します。ターゲットのパスが"/base"であり、受信したリクエストが"/dir"である場合、ターゲットのリクエストは/base/dirになります。
// NewSingleHostReverseProxyは、Hostヘッダーを書き換えません。
//
// NewSingleHostReverseProxyが提供する以上のカスタマイズをするには、Rewrite関数を使用して直接ReverseProxyを使用してください。ProxyRequest SetURLメソッドを使用してアウトバウンドリクエストをルーティングすることができます（ただし、SetURLはデフォルトでアウトバウンドリクエストのHostヘッダーを書き換えます）。
//
//	proxy := &ReverseProxy{
//			Rewrite: func(r *ProxyRequest) {
//				r.SetURL(target)
//				r.Out.Host = r.In.Host // 必要に応じて
//			},
//		}
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)
