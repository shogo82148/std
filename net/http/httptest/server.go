// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// サーバーの実装

package httptest

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
)

// Serverは、エンドツーエンドのHTTPテストで使用するために、
// ローカルループバックインターフェース上のシステムが選んだポートでリッスンするHTTPサーバーです。
type Server struct {
	URL      string
	Listener net.Listener

	// EnableHTTP2は、サーバー上でHTTP/2が有効かどうかを制御します。
	// NewUnstartedServerを呼び出すときとServer.StartTLSを呼び出すときの間に設定する必要があります。
	EnableHTTP2 bool

	// TLSはオプションのTLS構成であり、新しい構成でポピュレートされます
	// TLSが開始された後に。 StartTLSが呼び出される前に開始されていないサーバーに設定されている場合、既存のフィールドは新しい構成にコピーされます。
	TLS *tls.Config

	// NewUnstartedServer を呼び出した後、Start または StartTLS を実行する前に、Config を変更することができます。
	Config *http.Server

	// certificate はTLS設定の証明書の解析バージョンです。存在する場合にのみ使用されます。
	certificate *x509.Certificate

	// wgはこのサーバー上の未処理のHTTPリクエストの数をカウントします。
	// Closeはすべてのリクエストが終了するまでブロックします。
	wg sync.WaitGroup

	mu     sync.Mutex
	closed bool
	conns  map[net.Conn]http.ConnState

	// client はサーバーとの通信用に設定されています。
	// Close が呼び出されると、自動的にトランスポートが閉じられます。
	client *http.Client
}

<<<<<<< HEAD
// NewServer は新しい Server を起動して返します。
// 使用が終わったら、呼び出し元は Close を呼び出してシャットダウンする必要があります。
func NewServer(handler http.Handler) *Server

// NewUnstartedServerは新しいServerを返しますが、開始はしません。
=======
// NewServer starts and returns a new [Server].
// The caller should call Close when finished, to shut it down.
func NewServer(handler http.Handler) *Server

// NewUnstartedServer returns a new [Server] but doesn't start it.
>>>>>>> upstream/release-branch.go1.22
//
// 設定を変更した後、呼び出し元はStartまたはStartTLSを呼び出す必要があります。
//
// 使用し終えたらCloseを呼び出してシャットダウンする必要があります。
func NewUnstartedServer(handler http.Handler) *Server

// Start はNewUnstartedServerからサーバーを起動します。
func (s *Server) Start()

// StartTLSは、NewUnstartedServerからサーバー上でTLSを開始します。
func (s *Server) StartTLS()

<<<<<<< HEAD
// NewTLSServerはTLSを使用して新しいサーバーを起動し、それを返します。
// 終了時には、呼び出し元はシャットダウンするためにCloseを呼び出す必要があります。
=======
// NewTLSServer starts and returns a new [Server] using TLS.
// The caller should call Close when finished, to shut it down.
>>>>>>> upstream/release-branch.go1.22
func NewTLSServer(handler http.Handler) *Server

// Close はサーバーをシャットダウンし、このサーバーに対して保留中のすべてのリクエストが完了するまでブロックします。
func (s *Server) Close()

// CloseClientConnectionsはテストサーバーへのすべてのオープン中のHTTP接続を閉じます。
func (s *Server) CloseClientConnections()

// Certificateは、サーバーがTLSを使用していない場合はnil、それ以外の場合はサーバーが使用する証明書を返します。
func (s *Server) Certificate() *x509.Certificate

<<<<<<< HEAD
// Clientは、サーバーへのリクエストを行うために設定されたHTTPクライアントを返します。
// サーバーのTLSテスト証明書を信頼するように設定されており、Server.Close時にアイドル接続をクローズします。
=======
// Client returns an HTTP client configured for making requests to the server.
// It is configured to trust the server's TLS test certificate and will
// close its idle connections on [Server.Close].
>>>>>>> upstream/release-branch.go1.22
func (s *Server) Client() *http.Client
