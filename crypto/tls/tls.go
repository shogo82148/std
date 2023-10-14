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

<<<<<<< HEAD
// NewListenerは、内部Listenerからの接続を受け入れ、それぞれの接続をServerでラップするListenerを作成します。
// 設定configはnilではなく、少なくとも1つの証明書を含むか、それ以外の場合はGetCertificateを設定する必要があります。
=======
// NewListener creates a Listener which accepts connections from an inner
// Listener and wraps each connection with [Server].
// The configuration config must be non-nil and must include
// at least one certificate or else set GetCertificate.
>>>>>>> upstream/master
func NewListener(inner net.Listener, config *Config) net.Listener

// Listenは、net.Listenを使用して、指定されたネットワークアドレスで接続を受け入れるTLSリスナーを作成します。
// 設定（config）は、nil以外である必要があり、少なくとも1つの証明書を含んでいるか、GetCertificateを設定している必要があります。
func Listen(network, laddr string, config *Config) (net.Listener, error)

// DialWithDialerは、dialer.Dialを使用して指定されたネットワークアドレスに接続し、
// TLSハンドシェイクを開始し、結果のTLS接続を返します。dialerで指定されたタイムアウトや
// デッドラインは、接続とTLSハンドシェイク全体に適用されます。
//
<<<<<<< HEAD
// DialWithDialerは、nilの設定をゼロの設定として解釈します。
// デフォルトの内容については、Configのドキュメントを参照してください。
//
// DialWithDialerは、内部的にcontext.Backgroundを使用します。
// コンテキストを指定するには、NetDialerを必要なダイアラに設定した
// Dialer.DialContextを使用してください。
=======
// DialWithDialer interprets a nil configuration as equivalent to the zero
// configuration; see the documentation of [Config] for the defaults.
//
// DialWithDialer uses context.Background internally; to specify the context,
// use [Dialer.DialContext] with NetDialer set to the desired dialer.
>>>>>>> upstream/master
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
<<<<<<< HEAD
// 返されるConnは、もし存在する場合は常に*Conn型です。
//
// 内部的にDialはcontext.Backgroundを使用しますが、コンテキストを指定するにはDialContextを使用してください。
func (d *Dialer) Dial(network, addr string) (net.Conn, error)

// DialContextは指定されたネットワークアドレスに接続し、TLSハンドシェイクを開始し、結果のTLS接続を返します。
// 提供されたContextはnil以外である必要があります。接続が完了する前にContextが期限切れになった場合、エラーが返されます。接続成功後、Contextが期限切れになっても接続には影響しません。
// 返されるConn（あれば）は常に*Conn型です。
=======
// The returned [Conn], if any, will always be of type *[Conn].
//
// Dial uses context.Background internally; to specify the context,
// use [Dialer.DialContext].
func (d *Dialer) Dial(network, addr string) (net.Conn, error)

// DialContext connects to the given network address and initiates a TLS
// handshake, returning the resulting TLS connection.
//
// The provided Context must be non-nil. If the context expires before
// the connection is complete, an error is returned. Once successfully
// connected, any expiration of the context will not affect the
// connection.
//
// The returned [Conn], if any, will always be of type *[Conn].
>>>>>>> upstream/master
func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error)

// LoadX509KeyPairは、公開/秘密キーペアをペアのファイルから読み込んで解析します。
// ファイルにはPEMエンコードされたデータを含める必要があります。証明書ファイルには、
// リーフ証明書に続く中間証明書を含めて、証明書チェーンを形成することができます。成功した場合、
// Certificate.Leafはnilになります。これは、証明書の解析形式が保持されていないためです。
func LoadX509KeyPair(certFile, keyFile string) (Certificate, error)

// X509KeyPairは、一組のPEMエンコードされたデータから公開/秘密鍵のペアを解析します。成功した場合、Certificate.Leafはnilです。そのため、証明書の解析形式は保持されません。
func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (Certificate, error)
