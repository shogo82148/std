// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージtlsは、RFC 5246で指定されているTLS 1.2と、RFC 8446で指定されているTLS 1.3を部分的に実装しています。
package tls

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/net"
)

// Serverは、基礎となるトランスポートとしてconnを使用して新しいTLSサーバーサイド接続を返します。
// configの設定は、非nilであり、少なくとも1つの証明書を含んでいるか、またはGetCertificateを設定している必要があります。
func Server(conn net.Conn, config *Config) *Conn

// Client は、基礎となるトランスポートとして conn を使用して
// 新しい TLS クライアント側の接続を返します。
// ユーザーは、config に ServerName または InsecureSkipVerify のどちらかを設定する必要があります。config は nil であってはなりません。
func Client(conn net.Conn, config *Config) *Conn

// NewListenerは、内部Listenerからの接続を受け入れ、それぞれの接続を [Server] でラップするListenerを作成します。
// 設定configはnilではなく、少なくとも1つの証明書を含むか、それ以外の場合はGetCertificateを設定する必要があります。
func NewListener(inner net.Listener, config *Config) net.Listener

// Listenは、net.Listenを使用して、指定されたネットワークアドレスで接続を受け入れるTLSリスナーを作成します。
// 設定（config）は、nil以外である必要があり、少なくとも1つの証明書を含んでいるか、GetCertificateを設定している必要があります。
func Listen(network, laddr string, config *Config) (net.Listener, error)

// DialWithDialerは、dialer.Dialを使用して指定されたネットワークアドレスに接続し、
// TLSハンドシェイクを開始し、結果のTLS接続を返します。dialerで指定されたタイムアウトや
// デッドラインは、接続とTLSハンドシェイク全体に適用されます。
//
// DialWithDialerは、nilの設定をゼロの設定として解釈します。
// デフォルトの内容については、 [Config] のドキュメントを参照してください。
//
// DialWithDialerは、内部的にcontext.Backgroundを使用します。
// コンテキストを指定するには、NetDialerを必要なダイアラに設定した
// [Dialer.DialContext] を使用してください。
func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error)

// Dial関数は、net.Dialを使用して指定されたネットワークアドレスに接続し、
// 次にTLSハンドシェイクを開始して、結果として得られるTLSコネクションを返します。
// Dialはnilのconfigurationをゼロのconfigurationと同等だと解釈します。
// デフォルトの設定については、Configのドキュメンテーションを参照してください。
func Dial(network, addr string, config *Config) (*Conn, error)

// Dialerは、設定と基礎となる接続のためのダイアラーを使用してTLS接続をダイアルします。
type Dialer struct {

	// NetDialerはTLS接続の基礎となるTCP接続に使用するオプションのダイアラです。
	// nilのNetDialerはnet.Dialerのゼロ値と同じです。
	NetDialer *net.Dialer

	// Config は新しい接続に使用するTLSの設定です。
	// nilの設定はゼロと同等であり、デフォルトの設定についてはConfigのドキュメントを参照してください。
	Config *Config
}

// Dialは指定されたネットワークアドレスに接続し、TLSハンドシェイクを開始し、結果のTLS接続を返します。
//
// 返される [Conn] は、もし存在する場合は常に *[Conn] 型です。
//
// 内部的にDialはcontext.Backgroundを使用しますが、コンテキストを指定するにはDialContextを使用してください。
func (d *Dialer) Dial(network, addr string) (net.Conn, error)

// DialContextは指定されたネットワークアドレスに接続し、TLSハンドシェイクを開始し、結果のTLS接続を返します。
// 提供されたContextはnil以外である必要があります。接続が完了する前にContextが期限切れになった場合、エラーが返されます。接続成功後、Contextが期限切れになっても接続には影響しません。
// 返される [Conn] （あれば）は常に *[Conn] 型です。
func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error)

// LoadX509KeyPairは、ペアのファイルから公開/秘密鍵ペアを読み込んで解析します。
// ファイルにはPEMエンコードされたデータが含まれている必要があります。証明書ファイルには、
// リーフ証明書に続く中間証明書が含まれている場合があり、証明書チェーンを形成します。
// 成功時の戻り値では、Certificate.Leafが設定されます。
//
// Go 1.23以前ではCertificate.Leafはnilのままで、解析された証明書は破棄されました。
// この挙動は、GODEBUG環境変数で"x509keypairleaf=0"を設定することで再度有効にすることができます。
func LoadX509KeyPair(certFile, keyFile string) (Certificate, error)

// X509KeyPairは、ペアのPEMエンコードされたデータから公開/秘密鍵ペアを解析します。成功時の戻り値では、Certificate.Leafが設定されます。
//
// Go 1.23以前ではCertificate.Leafはnilのままで、解析された証明書は破棄されました。この挙動は、GODEBUG環境変数で"x509keypairleaf=0"
// を設定することで再度有効にすることができます。
func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (Certificate, error)
