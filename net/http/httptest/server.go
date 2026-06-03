// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// サーバーの実装

package httptest

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/internal/nettest"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/testing"
)

// Serverは、エンドツーエンドのHTTPテストで使用するHTTPサーバーです。
//
// ほとんどのテストでは、[NewTestServer] でサーバーを作成するべきです。
// [Server.Client] メソッドは、テストサーバーにリクエストを送るクライアントを返します。
//
//	// テストサーバーを作成してリクエストを送信する。
//	server := httptest.NewTestServer(t, handler)
//	resp, err := server.Client().Get("http://www.example.com/")
//
// # Configuration
//
// テストでは、使用前にServerの設定を変更できます。
// [Server.Client]、[Server.Start]、[Server.StartTLS] のいずれかを
// 最初に呼び出した後は、設定を変更してはいけません。
//
//	// 使用前にテストサーバーを設定する。
//	server := httptest.NewTestServer(t, handler)
//	server.Config.MaxHeaderBytes = 1024
//	resp, err := server.Client().Get("http://www.example.com/")
//
// # Tests
//
// [NewTestServer] で作成されたサーバーは次の動作をします。
//
//   - サーバーハンドラが [http.ErrAbortHandler] 以外の値でパニックした場合、
//     テストを失敗させます。
//   - テスト終了時にサーバーを停止するためのCleanup関数を登録します。
//
// それ以外の方法で作成したサーバーは、[Server.Close] で手動停止する必要があります。
//
// # In-Memory Network
//
// Serverは、インメモリネットワーク実装を使うことも、
// ローカルネットワークのループバックインターフェースで待ち受けることもできます。
// 多くのテストでは、ポート枯渇や一時的なネットワーク問題を避けられ、
// [testing/synctest] パッケージとも相性のよいインメモリネットワークを使うべきです。
//
// インメモリネットワークを使うには、[NewTestServer] でサーバーを作成します。
// [Server.Start] や [Server.StartTLS] は呼び出さないでください。
//
// インメモリネットワーク使用時、[Server.Client] が返す [http.Client] は
// すべてのリクエストをこのサーバーへ送るよう設定されます。
// クライアントは宛先アドレスやホスト名に関係なく、
// HTTPおよびHTTPSのリクエストをサーバーへ向けます。
// リクエストのベースURLとして [Server.URL] を使う必要はありません。
//
//	server := httptest.NewTestServer(t, handler)
//	client := server.Client()
//
//	// これらのリクエストはすべてテストサーバーに送信される。
//	// https:// のリクエストはインメモリネットワーク上でTLSを使用する。
//	_, _ = client.Get("http://www.example.com/")
//	_, _ = client.Get("https://go.dev/")
//	_, _ = client.Get("http://10.0.0.1/")
//
// インメモリネットワーク使用時は、[Server.Listener] フィールドは設定されません。
//
// # Loopback Network
//
// ループバックインターフェースで待ち受けるには、[Server.Start] または [Server.StartTLS] を呼び出します。
// サーバーはシステムが選んだポートで待ち受けます。
//
// ループバックサーバーは、[Server.Start] で起動した場合はHTTP、
// [Server.StartTLS] で起動した場合はHTTPSを提供します。
//
// ループバックネットワーク使用時、[Server.Client] が返す [http.Client] は、
// ホスト名が "example.com" または ".example.com" のサブドメインである
// リクエストをサーバーへ送るよう設定されます。
//
// リクエストはサーバーのループバックアドレスにも送信できます。
// [Server.URL] フィールドにはサーバーのアドレスを含むベースURLが設定されます。
//
//	server := httptest.NewTestServer(t, handler)
//	server.Start()
//	client := server.Client()
//
//	// このリクエストはテストサーバーに送信される。
//	_, _ = server.Client().Get(server.URL + "/")
//
//	// このリクエスト（http.DefaultClientを使用）も、
//	// server.URL にサーバーのローカルIPアドレスが含まれるためテストサーバーに送信される。
//	_, _ = http.Get(server.URL + "/")
type Server struct {
	// URLはサーバーのベースURLで、末尾スラッシュのない
	// http://address:port 形式です。
	//
	// Client、Start、StartTLS のいずれかを最初に呼び出したときに設定されます。
	//
	// ループバックで待ち受けるサーバーでは、アドレスはサーバーのループバックIPアドレスです。
	//
	// インメモリネットワークを使うサーバーでは、このアドレスは "example.com" です。
	// インメモリネットワークを使うサーバーへのリクエストでは任意のアドレスを使用でき、
	// このベースURLを使う必要はありません。
	URL string

	// Listenerは、ループバックで待ち受けるサーバーのネットワークリスナーです。
	// インメモリネットワークを使うサーバーでは設定されません。
	Listener net.Listener

	// EnableHTTP2は、サーバーでHTTP/2を有効にするかどうかを制御します。
	// Client、Start、StartTLS を呼び出す前に設定する必要があります。
	EnableHTTP2 bool

	// TLSはオプションのTLS構成であり、新しい構成でポピュレートされます
	// TLSが開始された後に。 StartTLSが呼び出される前に開始されていないサーバーに設定されている場合、既存のフィールドは新しい構成にコピーされます。
	TLS *tls.Config

	// Configは、Client、Start、StartTLS を呼び出す前であれば変更できます。
	Config *http.Server

	t testing.TB

	// certificate is a parsed version of the TLS config certificate, if present.
	certificate *x509.Certificate

	// startOnce is used to start fakenet servers once.
	startOnce sync.Once

	// started indicates whether the server has been started.
	started bool

	// Fake network listeners, one for HTTP and one for HTTPS.
	fakeListener    *nettest.Listener
	fakeTLSListener *nettest.Listener

	// wg counts the number of outstanding HTTP requests on this server.
	// Close blocks until all requests are finished.
	wg sync.WaitGroup

	mu     sync.Mutex
	closed bool
	conns  map[net.Conn]http.ConnState

	// client はサーバーとの通信用に設定されています。
	// Close が呼び出されると、自動的にトランスポートが閉じられます。
	client *http.Client
}

// NewTestServerは、テスト用の新しい [Server] を返します。
// サーバーはデフォルトでインメモリネットワーク実装を使用します。
//
// handlerがnilの場合、サーバーはすべてのリクエストに500レスポンスを返します。
// [http.DefaultServeMux] は使用しません。
//
// 詳細は [Server] のドキュメントを参照してください。
func NewTestServer(t testing.TB, handler http.Handler) *Server

// NewServerは、ローカルネットワークのループバックインターフェースで待ち受ける
// 新しい [Server] を起動して返します。
// これは [NewUnstartedServer] を呼び出してから [Server.Start] を呼ぶのと同等です。
//
// 呼び出し側は、終了時に [Server.Close] を呼んで停止する必要があります。
//
// ほとんどのユーザーは代わりに [NewTestServer] を使うべきです。
// 詳細は [Server] のドキュメントを参照してください。
func NewServer(handler http.Handler) *Server

// NewUnstartedServerは、ローカルネットワークのループバックインターフェースで待ち受ける
// 新しい [Server] を返します。サーバーは起動しません。
//
// サーバー設定を変更した後、呼び出し側は [Server.Start] または [Server.StartTLS] を
// 呼び出す必要があります。
//
// 呼び出し側は、終了時に [Server.Close] を呼んで停止する必要があります。
//
// ほとんどのユーザーは代わりに [NewTestServer] を使うべきです。
// 詳細は [Server] のドキュメントを参照してください。
func NewUnstartedServer(handler http.Handler) *Server

// Startは、ローカルのループバックネットワークインターフェースでサーバーを起動します。
//
// サーバーは [NewTestServer] または [NewUnstartedServer] で作成されている必要があります。
func (s *Server) Start()

// StartTLSは、ローカルのループバックネットワークインターフェース上のサーバーでTLSを起動します。
//
// サーバーは [NewTestServer] または [NewUnstartedServer] で作成されている必要があります。
func (s *Server) StartTLS()

// NewTLSServerは、TLSを使用し、ローカルネットワークのループバックインターフェースで待ち受ける
// 新しい [Server] を起動して返します。
// これは [NewUnstartedServer] を呼び出してから [Server.StartTLS] を呼ぶのと同等です。
//
// 呼び出し側は、終了時に [Server.Close] を呼んで停止する必要があります。
//
// ほとんどのユーザーは代わりに [NewTestServer] を使うべきです。
// 詳細は [Server] のドキュメントを参照してください。
func NewTLSServer(handler http.Handler) *Server

// Close はサーバーをシャットダウンし、このサーバーに対して保留中のすべてのリクエストが完了するまでブロックします。
func (s *Server) Close()

// CloseClientConnectionsはテストサーバーへのすべてのオープン中のHTTP接続を閉じます。
func (s *Server) CloseClientConnections()

// Certificateは、サーバーがTLSを使用していない場合はnil、それ以外の場合はサーバーが使用する証明書を返します。
func (s *Server) Certificate() *x509.Certificate

// Clientは、サーバーへのリクエスト送信用に設定されたHTTPクライアントを返します。
// このクライアントはサーバーのTLSテスト証明書を信頼するよう設定されており、
// [Server.Close] 時にアイドル接続を閉じます。
func (s *Server) Client() *http.Client
