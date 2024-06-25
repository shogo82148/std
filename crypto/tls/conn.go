// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TLSの低レベル接続とレコードレイヤー

package tls

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// Connはセキュア接続を表します。
// net.Connインターフェースを実装しています。
type Conn struct {
	// 定数
	conn        net.Conn
	isClient    bool
	handshakeFn func(context.Context) error
	quic        *quicState

	// isHandshakeCompleteは、接続が現在アプリケーションデータを転送している場合（つまり、ハンドシェイク処理を行っていない場合）にtrueです。
	// isHandshakeCompleteがtrueである場合、handshakeErr == nilとなります。
	isHandshakeComplete atomic.Bool
	// ハンドシェイクの後の定数; ハンドシェイクミューテックスによって保護される
	handshakeMutex sync.Mutex
	handshakeErr   error
	vers           uint16
	haveVers       bool
	config         *Config

	// handshakesはこれまでに接続で行われたハンドシェイクの回数を数えます。再交渉が無効化されている場合、これは0または1です。
	handshakes       int
	extMasterSecret  bool
	didResume        bool
	didHRR           bool
	cipherSuite      uint16
	curveID          CurveID
	ocspResponse     []byte
	scts             [][]byte
	peerCertificates []*x509.Certificate

	// activeCertHandlesにはpeerCertificates内の証明書のキャッシュハンドルが格納されており、アクティブな参照を追跡するために使用されます。
	activeCertHandles []*activeCert

	// verifiedChainsには、私たちが構築した証明書チェーンが含まれています。
	// これは、サーバーが提示した証明書チェーンとは異なります。
	verifiedChains [][]*x509.Certificate
	// serverName には、クライアントが指定したサーバー名が含まれています。
	serverName string

	// secureRenegotiation は、サーバーが安全な再ネゴシエーション拡張を返した場合は true です。
	// これは、サーバーの場合には無意味であり、再ネゴシエーションはサポートされていません。
	secureRenegotiation bool
	// ekmはキーマテリアルをエクスポートするためのクロージャです。
	ekm func(label string, context []byte, length int) ([]byte, error)

	// resumptionSecretは、処理または送信するための再開マスターシークレットです。または、NewSessionTicketメッセージを送信するためのものです。
	resumptionSecret []byte

	// ticketKeys は、この接続の有効なセッションチケットキーのセットです。
	// 最初のキーは新しいチケットの暗号化に使用され、すべてのキーがデコードの試行に使用されます。
	ticketKeys []ticketKey

	// clientFinishedIsFirst は、最新のハンドシェイク中にクライアントが最初のFinishedメッセージを送信した場合にtrueです。これは記録されます。最初に送信されたFinishedメッセージはtls-uniqueチャネルバインディング値です。
	clientFinishedIsFirst bool

	// closeNotifyErrはalertCloseNotifyレコードの送信エラーです。
	closeNotifyErr error

	// closeNotifySent は、Conn が alertCloseNotify レコードを送信しようとした場合にtrueです。
	closeNotifySent bool

	// clientFinishedとserverFinishedには、最新のハンドシェイクでクライアントまたはサーバーが送信したFinishedメッセージが含まれています。これは、再ネゴシエーション拡張とtls-uniqueチャネルバインディングをサポートするために保持されます。
	clientFinished [12]byte
	serverFinished [12]byte

	// clientProtocolは、ALPNプロトコルの協議結果です。
	clientProtocol string

	// 入力/出力
	in, out   halfConn
	rawInput  bytes.Buffer
	input     bytes.Reader
	hand      bytes.Buffer
	buffering bool
	sendBuf   []byte

	// bytesSentは送信されたアプリケーションデータのバイト数をカウントします。
	// packetsSentはパケットの数をカウントします。
	bytesSent   int64
	packetsSent int64

	// retryCount は Conn.readRecord によって受信された、連続して進展しないレコードの数をカウントします。つまり、ハンドシェイクを進行させず、またアプリケーションデータを送信しないレコードです。in.Mutex によって保護されています。
	retryCount int

	// activeCallはCloseが呼び出されたかどうかを最下位ビットで示します。
	// 残りのビット数はConn.Write内のゴルーチンの数です。
	activeCall atomic.Int32

	tmp [16]byte
}

// LocalAddrはローカルネットワークアドレスを返します。
func (c *Conn) LocalAddr() net.Addr

// RemoteAddrはリモートネットワークアドレスを返します。
func (c *Conn) RemoteAddr() net.Addr

// SetDeadlineは接続に関連付けられた読み込みと書き込みのタイムアウトを設定します。
// tのゼロ値は、 [Conn.Read] と  [Conn.Write] がタイムアウトしないことを意味します。
// 書き込みがタイムアウトした後、TLSの状態が破損し、将来の書き込みは同じエラーを返します。
func (c *Conn) SetDeadline(t time.Time) error

// SetReadDeadlineは基礎となる接続の読み込みの期限を設定します。
// tのゼロ値は、 [Conn.Read] がタイムアウトしないことを意味します。
func (c *Conn) SetReadDeadline(t time.Time) error

// SetWriteDeadlineは、基礎となる接続に書き込みの期限を設定します。
// tのゼロ値は、 [Conn.Write] がタイムアウトしないことを意味します。
// 書き込みがタイムアウトした後、TLSの状態が壊れるため、以降の書き込みは同じエラーを返します。
func (c *Conn) SetWriteDeadline(t time.Time) error

// NetConnはcによってラップされた基になる接続を返します。
// ただし、この接続に直接書き込みまたは読み込みを行うと、TLSセッションが破損することに注意してください。
func (c *Conn) NetConn() net.Conn

// RecordHeaderError は、TLS レコードヘッダが無効な場合に返されます。
type RecordHeaderError struct {
	// Msgはエラーを説明する人が読みやすい文字列を含んでいます。
	Msg string

	// RecordHeaderには、エラーを引き起こしたTLSレコードヘッダの5バイトが含まれています。
	RecordHeader [5]byte

	// Connは、クライアントが初期ハンドシェイクを送信しているが、TLSのように見えない場合の基礎となるnet.Connを提供します。
	// ハンドシェイクが既に行われているか、接続にTLSアラートが書き込まれている場合は、nilです。
	Conn net.Conn
}

func (e RecordHeaderError) Error() string

// Writeは接続にデータを書き込みます。
//
// [Conn.Handshake] を呼び出すため、無期限のブロッキングを防ぐために
// ハンドシェイクが完了していない場合、Writeを呼び出す前に
// [Conn.Read] とWriteの両方に期限を設定する必要があります。
// [Conn.SetDeadline] 、 [Conn.SetReadDeadline] 、および [Conn.SetWriteDeadline] を参照してください。
func (c *Conn) Write(b []byte) (int, error)

// Readは接続からデータを読み込みます。
//
// [Conn.Handshake] を呼び出すため、ハンドシェイクがまだ完了していない場合、
// Readが呼び出される前にReadと [Conn.Write] の両方にデッドラインを設定する必要があります
// 無制限のブロッキングを防ぐためです。
// [Conn.SetDeadline] 、 [Conn.SetReadDeadline] 、および [Conn.SetWriteDeadline] を参照してください。
func (c *Conn) Read(b []byte) (int, error)

// Closeは接続を閉じます。
func (c *Conn) Close() error

// CloseWriteは接続の書き込み側をシャットダウンします。ハンドシェイクが完了した後に一度だけ呼び出され、基礎となる接続上でCloseWriteを呼び出しません。ほとんどの呼び出し元は単に [Conn.Close] を使用すべきです。
func (c *Conn) CloseWrite() error

// Handshakeはクライアントまたはサーバーのハンドシェイクプロトコルを実行します。
// まだ実行されていない場合、ほとんどのこのパッケージの使用では、明示的にHandshakeを呼び出す必要はありません：最初の [Conn.Read] または [Conn.Write] が自動的に呼び出します。
//
// ハンドシェイクのキャンセルやタイムアウトの設定に関して制御するためには、 [Conn.HandshakeContext] または [Dialer] のDialContextメソッドを使用します。
//
// サーバーまたはクライアントが送信する証明書のRSAキーサイズは、拒否サービス攻撃を防ぐために、8192ビットに制限されています。この制限は、GODEBUG環境変数（例：GODEBUG=tlsmaxrsasize=4096）のtlsmaxrsasizeを設定することで上書きすることができます。
func (c *Conn) Handshake() error

// HandshakeContextは、クライアントまたはサーバーのハンドシェイクプロトコルを実行します。
// まだ実行されていない場合、提供されたコンテキストはnil以外である必要があります。
// ハンドシェイクが完了する前にコンテキストがキャンセルされた場合、ハンドシェイクは中断され、エラーが返されます。
// ハンドシェイクが完了すると、コンテキストのキャンセルは接続に影響を与えません。
//
// このパッケージのほとんどの使用では、明示的にHandshakeContextを呼び出す必要はありません：最初の [Conn.Read] または [Conn.Write] が自動的に呼び出します。
func (c *Conn) HandshakeContext(ctx context.Context) error

// ConnectionState関数は、接続に関する基本的なTLSの詳細を返します。
func (c *Conn) ConnectionState() ConnectionState

// OCSPResponseは、TLSサーバーからステープルされたOCSP応答を返します（クライアント接続の場合のみ有効）。
func (c *Conn) OCSPResponse() []byte

// VerifyHostnameは、ホストに接続するためのピア証明書チェーンが有効かどうかを確認します。有効であれば、nilを返します。そうでなければ、問題を説明するエラーを返します。
func (c *Conn) VerifyHostname(host string) error
