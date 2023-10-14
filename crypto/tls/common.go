// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

const (
	VersionTLS10 = 0x0301
	VersionTLS11 = 0x0302
	VersionTLS12 = 0x0303
	VersionTLS13 = 0x0304

	// Deprecated: SSLv3は暗号的に破損しており、
	// このパッケージではサポートされていません。golang.org/issue/32716を参照してください。
	VersionSSL30 = 0x0300
)

// VersionNameは提供されたTLSバージョン番号の名前（例："TLS 1.3"）を返します。
// もしバージョンがこのパッケージで実装されていない場合は、値のフォールバック表現を返します。
func VersionName(version uint16) string

// CurveIDは楕円曲線のためのTLS識別子のタイプです。参照：
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8.
//
// TLS 1.3では、このタイプはNamedGroupと呼ばれますが、現在のところ、
// このライブラリは楕円曲線ベースのグループのみをサポートしています。RFC 8446、セクション4.2.7を参照してください。
type CurveID uint16

const (
	CurveP256 CurveID = 23
	CurveP384 CurveID = 24
	CurveP521 CurveID = 25
	X25519    CurveID = 29
)

// ConnectionStateは接続に関する基本的なTLSの詳細を記録します。
type ConnectionState struct {
	// Versionは接続で使用されるTLSバージョンです（例：VersionTLS12）。
	Version uint16

	// HandshakeComplete はハンドシェイクが完了している場合に true です。
	HandshakeComplete bool

	// DidResume は、この接続がセッションチケットや類似のメカニズムによって前のセッションから正常に再開された場合に true です。
	DidResume bool

	// CipherSuiteは、接続のためにネゴシエートされた暗号スイートです（例： TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256、TLS_AES_128_GCM_SHA256）。
	CipherSuite uint16

	// NegotiatedProtocolはALPNで交渉されたアプリケーションプロトコルです。
	NegotiatedProtocol string

	// NegotiatedProtocolIsMutualは相互のNPN（Next Protocol Negotiation）交渉を示すために使われます。
	//
	// 廃止予定：この値は常にtrueです。
	NegotiatedProtocolIsMutual bool

	// ServerNameはクライアントが送信するServer Name Indication拡張機能の値です。これはサーバーとクライアントの両方で利用できます。
	ServerName string

	// PeerCertificates は、ピアから送信された証明書が解析されたものであり、送信順に格納されています。最初の要素は、接続が検証されるリーフ証明書です。
	//
	// クライアント側では空にすることはできません。サーバー側では、Config.ClientAuth が RequireAnyClientCert または RequireAndVerifyClientCert でない場合、空にすることができます。
	//
	// PeerCertificates およびその内容は変更してはいけません。
	PeerCertificates []*x509.Certificate

	// VerifiedChainsは、最初の要素がPeerCertificates[0]であり、最後の要素がConfig.RootCAs（クライアント側）またはConfig.ClientCAs（サーバー側）からの要素の1つ以上のチェーンのリストです。
	//
	// クライアント側では、Config.InsecureSkipVerifyがfalseの場合に設定されます。サーバー側では、Config.ClientAuthがVerifyClientCertIfGiven（かつピアが証明書を提供した場合）またはRequireAndVerifyClientCertである場合に設定されます。
	//
	// VerifiedChainsおよびその内容は変更しないでください。
	VerifiedChains [][]*x509.Certificate

	// SignedCertificateTimestampsは、もしあれば、ピアからのTLSハンドシェイクによって提供されるリーフ証明書のSCTのリストです。
	SignedCertificateTimestamps [][]byte

	// OCSPResponseは、ピアから提供される、必要に応じてリーフ証明書のステープル化されたオンライン証明書ステータスプロトコル（OCSP）レスポンスです。
	OCSPResponse []byte

	// TLSUniqueには、"tls-unique"チャネルバインディングの値が格納されています（RFC 5929、セクション3を参照）。
	// この値は、TLS 1.3接続や拡張されたマスターシークレット（RFC 7627）をサポートしていない再開接続ではnilになります。
	TLSUnique []byte

	// ekmはExportKeyingMaterialを介して公開されるクロージャです。
	ekm func(label string, context []byte, length int) ([]byte, error)
}

// ExportKeyingMaterialは、RFC 5705で定義されているように、エクスポートされたキーマテリアルの長さのバイトを新しいスライスで返します。contextがnilの場合、シードの一部として使用されません。Config.Renegotiation経由で再協議が許可された接続の場合、この関数はエラーを返します。
// 返された値が接続に固有でない場合、条件があります。RFC 5705およびRFC 7627のセキュリティに関する考慮事項セクション、およびhttps://mitls.org/pages/attacks/3SHAKE#channelbindingsを参照してください。
func (cs *ConnectionState) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)

// ClientAuthTypeは、TLSクライアント認証に関するサーバーのポリシーを宣言します。
type ClientAuthType int

const (

	// NoClientCertはハンドシェイク中にクライアント証明書を要求しないことを示し、
	// 送信された証明書があっても検証されないことを意味します。
	NoClientCert ClientAuthType = iota

	// RequestClientCertは、ハンドシェイク中にクライアント証明書の要求を示しますが、クライアントが証明書を送信することは必要ありません。
	RequestClientCert

	// RequireAnyClientCertは、ハンドシェイク中にクライアント証明書を要求し、
	// クライアントから少なくとも1つの証明書の送信が必要であることを示しますが、
	// その証明書が有効である必要はありません。
	RequireAnyClientCert

	// VerifyClientCertIfGivenは、ハンドシェイク中にクライアント証明書の要求をすることを示しますが、クライアントが証明書を送信する必要はありません。ただし、クライアントが証明書を送信する場合は、その証明書が有効であることが必要です。
	VerifyClientCertIfGiven

	// RequireAndVerifyClientCertは、ハンドシェイク中にクライアント証明書の要求が行われることを示し、クライアントが少なくとも1つの有効な証明書を送信する必要があることを示します。
	RequireAndVerifyClientCert
)

// ClientSessionCacheは、クライアントが特定のサーバーとTLSセッションを再開するために使用できるClientSessionStateオブジェクトのキャッシュです。ClientSessionCacheの実装は、異なるゴルーチンから同時に呼び出されることを想定しています。TLS 1.2までは、SessionIDベースの再開ではなく、チケットベースの再開のみがサポートされています。TLS 1.3では、これらがPSKモードにマージされ、このインターフェースを介してサポートされています。
type ClientSessionCache interface {
	// Get searches for a ClientSessionState associated with the given key.
	// On return, ok is true if one was found.
	Get(sessionKey string) (session *ClientSessionState, ok bool)

	// Put adds the ClientSessionState to the cache with the given key. It might
	// get called multiple times in a connection if a TLS 1.3 server provides
	// more than one session ticket. If called with a nil *ClientSessionState,
	// it should remove the cache entry.
	Put(sessionKey string, cs *ClientSessionState)
}

// SignatureSchemeは、TLSでサポートされる署名アルゴリズムを識別します。RFC 8446、セクション4.2.3を参照してください。
type SignatureScheme uint16

const (
	// RSASSA-PKCS1-v1_5 アルゴリズム。
	PKCS1WithSHA256 SignatureScheme = 0x0401
	PKCS1WithSHA384 SignatureScheme = 0x0501
	PKCS1WithSHA512 SignatureScheme = 0x0601

	// 公開キーOID rsaEncryption を使用した RSASSA-PSS アルゴリズム。
	PSSWithSHA256 SignatureScheme = 0x0804
	PSSWithSHA384 SignatureScheme = 0x0805
	PSSWithSHA512 SignatureScheme = 0x0806

	// ECDSAアルゴリズム。TLS 1.3では特定の曲線に制約される。
	ECDSAWithP256AndSHA256 SignatureScheme = 0x0403
	ECDSAWithP384AndSHA384 SignatureScheme = 0x0503
	ECDSAWithP521AndSHA512 SignatureScheme = 0x0603

	// EdDSAアルゴリズム。
	Ed25519 SignatureScheme = 0x0807

	// TLS 1.2用の旧バージョンの署名およびハッシュアルゴリズム。
	PKCS1WithSHA1 SignatureScheme = 0x0201
	ECDSAWithSHA1 SignatureScheme = 0x0203
)

// ClientHelloInfoには、GetCertificateおよびGetConfigForClientのコールバックにおいて、
// アプリケーションロジックをガイドするためのClientHelloメッセージからの情報が含まれています。
type ClientHelloInfo struct {

	// CipherSuitesはクライアントがサポートするCipherSuitesをリストアップしています（例：TLS_AES_128_GCM_SHA256、TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256）。
	CipherSuites []uint16

	// ServerNameは、クライアントがリクエストしたサーバーの名前を示します。
	// 仮想ホスティングをサポートするために、クライアントがSNIを使用している場合にのみ、ServerNameが設定されます（RFC 4366、セクション3.1を参照）。
	ServerName string

	// SupportedCurvesはクライアントでサポートされている楕円曲線をリストアップしています。
	// SupportedCurvesは、サポートされている楕円曲線拡張が使用されている場合にのみ設定されます(RFC 4492、セクション5.1.1を参照)。
	SupportedCurves []CurveID

	// SupportedPointsはクライアントがサポートするポイントフォーマットをリストアップしています。
	// SupportedPointsは、サポートされているポイントフォーマット拡張が利用されている場合にのみ設定されます（RFC 4492、セクション5.1.2を参照）。
	SupportedPoints []uint8

	// SignatureSchemesは、クライアントが検証可能な署名とハッシュ方式のリストです。SignatureSchemesは、Signature Algorithms拡張が使用されている場合にのみ設定されます（RFC 5246、セクション7.4.1.4.1を参照）。
	SignatureSchemes []SignatureScheme

	// SupportedProtosはクライアントがサポートしているアプリケーションプロトコルのリストです。
	// SupportedProtosは、アプリケーション層プロトコルネゴシエーション拡張が使用されている場合にのみ設定されます（RFC 7301、セクション3.1を参照）。
	//
	// サーバーは、GetConfigForClientの返り値でConfig.NextProtosを設定することでプロトコルを選択できます。
	SupportedProtos []string

	// SupportedVersions はクライアントでサポートされているTLSのバージョンをリストにします。
	// TLSバージョン1.3未満では、これはクライアントがアドバタイズする最大のバージョンから予測されるため、最大の値以外の値は使用されると拒否される可能性があります。
	SupportedVersions []uint16

	// Connは接続の基礎となるnet.Connです。この接続から読み取ったり、書き込んだりしないでください。それはTLS接続の失敗を引き起こします。
	Conn net.Conn

	// configはGetCertificateまたはGetConfigForClientの呼び出し元で埋め込まれ、
	// SupportsCertificateと共に使用されます。
	config *Config

	// ctxは進行中のハンドシェイクのコンテキストです。
	ctx context.Context
}

// Contextは進行中のハンドシェイクのコンテキストを返します。
// このコンテキストは、HandshakeContextに渡されたコンテキストの子孫であり、
// ハンドシェイクが終了するとキャンセルされます。
func (c *ClientHelloInfo) Context() context.Context

// CertificateRequestInfo は、クライアントから証明書とコントロールの証明を要求するために使用される、
// サーバーの CertificateRequest メッセージからの情報を含んでいます。
type CertificateRequestInfo struct {

	// AcceptableCAsは、ゼロまたは複数のDERエンコードされたX.501の区別名を含んでいます。これらは、サーバーが返される証明書の署名元として望むルートまたは中間CAの名前です。空のスライスは、サーバーが好みを持たないことを示します。
	AcceptableCAs [][]byte

	// SignatureSchemesは、サーバーが検証を行いたい署名スキームをリストアップします。
	SignatureSchemes []SignatureScheme

	// Versionは、この接続のために交渉されたTLSのバージョンです。
	Version uint16

	// ctxは進行中のハンドシェイクのコンテキストです。
	ctx context.Context
}

// Contextは進行中のハンドシェイクのコンテキストを返します。
// このコンテキストは、HandshakeContextに渡されたコンテキストの子であり、
// ハンドシェイクが終了するとキャンセルされます。
func (c *CertificateRequestInfo) Context() context.Context

// RenegotiationSupportは、TLS renegotiationのサポートレベルを列挙しています。TLS renegotiationとは、最初の接続の後に接続で追加のhandshakeを行うことです。これにより、ステートマシンが複雑化し、多くの微妙なセキュリティ上の問題の原因となっています。再交渉の開始はサポートされていませんが、再交渉リクエストの受け入れには対応している場合があります。
// 有効にされていても、サーバーはハンドシェイク間で自身の識別情報を変更することはできません（つまり、リーフ証明書は同じである必要があります）。また、同時にハンドシェイクとアプリケーションデータのフローを行うことは許可されていないため、再交渉は、再交渉と同期するプロトコル（例：HTTPS）とのみ使用できます。
// 再交渉はTLS 1.3で定義されていません。
type RenegotiationSupport int

const (
	// RenegotiateNeverは再交渉を無効にします。
	RenegotiateNever RenegotiationSupport = iota

	// RenegotiateOnceAsClientは、リモートサーバーに対して
	// 接続ごとに再交渉を一度だけ要求することを可能にします。
	RenegotiateOnceAsClient

	// RenegotiateFreelyAsClientは、リモートサーバーが繰り返し再交渉を要求できるようにします。
	RenegotiateFreelyAsClient
)

// Config構造体はTLSクライアントやサーバーを設定するために使用されます。
// ConfigがTLS関数に渡された後は変更しないでください。
// Configは再利用することができますが、tlsパッケージ自体は変更しません。
type Config struct {

	// RandはノンスやRSAブラインディングのエントロピーの源を提供します。
	// もしRandがnilの場合、TLSはパッケージcrypto/randの暗号化されたランダムリーダーを使用します。
	// このリーダーは複数のゴルーチンによる使用に安全である必要があります。
	Rand io.Reader

	// Timeはエポックからの経過秒数として現在時刻を返します。
	// Timeがnilの場合、TLSはtime.Nowを使用します。
	Time func() time.Time

	// Certificatesは、接続の相手側に提供する1つ以上の証明書チェーンを含んでいます。相手側の要件と互換性のある最初の証明書が自動的に選択されます。
	//
	// サーバーの設定では、Certificates、GetCertificate、またはGetConfigForClientのいずれかを設定する必要があります。クライアント認証を行うクライアントは、CertificatesまたはGetClientCertificateのいずれかを設定することができます。
	//
	// 注意：複数のCertificatesがあり、オプションのLeafフィールドが設定されていない場合、証明書の選択には著しいハンドシェイクごとのパフォーマンスのコストがかかります。
	Certificates []Certificate

	// NameToCertificate は、証明書名を Certificates の要素にマッピングします。
	// 証明書名は'*.example.com'のような形式であるため、必ずしもドメイン名である必要はありません。
	//
	// Deprecated: NameToCertificate では、特定の名前に対して単一の証明書の関連付けしか許可されません。
	// このフィールドを nil のままにしておくと、ライブラリが Certificates から最初の互換性のあるチェーンを選択します。
	NameToCertificate map[string]*Certificate

	// GetCertificateは与えられたClientHelloInfoに基づいて証明書を返します。
	// クライアントがSNI情報を提供する場合またはCertificatesが空の場合のみ呼び出されます。
	//
	// GetCertificateがnilであるかnilを返す場合、証明書はNameToCertificateから取得されます。
	// NameToCertificateがnilの場合、Certificatesの最良の要素が使用されます。
	//
	// 一度証明書が返されたら、変更しないでください。
	GetCertificate func(*ClientHelloInfo) (*Certificate, error)

	// GetClientCertificateは、クライアントが証明書を要求する場合に呼び出されます。
	// 設定されている場合、Certificatesの内容は無視されます。
	//
	// GetClientCertificateがエラーを返すと、ハンドシェイクは中止され、そのエラーが返されます。
	// それ以外の場合、GetClientCertificateはnilではないCertificateを返さなければなりません。
	// Certificate.Certificateが空である場合、サーバーには証明書は送信されません。
	// サーバーがこれを受け入れられない場合、ハンドシェイクを中止することがあります。
	//
	// GetClientCertificateは、再協議が発生するか、TLS 1.3が使用されている場合に、同じ接続に対して複数回呼び出される可能性があります。
	//
	// 一度Certificateが返されたら、変更しないでください。
	GetClientCertificate func(*CertificateRequestInfo) (*Certificate, error)

	// GetConfigForClientは、クライアントからClientHelloが受信された後に呼び出されます。この接続を処理するために使用されるConfigを変更するために、非nilのConfigを返すことができます。返されたConfigがnilの場合、元のConfigが使用されます。このコールバックによって返されたConfigは、後で変更できません。
	//
	// GetConfigForClientがnilの場合、Server()に渡されたConfigがすべての接続に使用されます。
	//
	// 返されたConfigに明示的にSessionTicketKeyが設定されている場合、または返されたConfigにSetSessionTicketKeysが呼び出された場合、これらのキーが使用されます。それ以外の場合、元のConfigキーが使用されます（自動的に管理される場合、回転する可能性もあります）。
	GetConfigForClient func(*ClientHelloInfo) (*Config, error)

	// VerifyPeerCertificate は、TLSクライアントまたはサーバーによる通常の証明書検証の後に呼び出されます（nilでない場合）。この関数は、ピアから提供された生のASN.1証明書と、通常の処理で検証されたチェーンを受け取ります。もしnon-nilのエラーを返す場合、ハンドシェイクは中断され、そのエラーが結果となります。
	// 通常の検証に失敗した場合、このコールバックは検討される前にハンドシェイクが中断されます。通常の検証が無効になっている場合（InsecureSkipVerifyがクライアント側で設定されている場合、またはServerAuthがRequestClientCertまたはRequireAnyClientCertの場合）、このコールバックは考慮されますが、verifiedChains引数は常にnilになります。ClientAuthがNoClientCertの場合、このコールバックはサーバーで呼び出されません。rawCertsは、ClientAuthがRequestClientCertまたはVerifyClientCertIfGivenの場合には、サーバーで空になる可能性があります。
	// このコールバックは再開された接続では呼び出されず、証明書は再検証されません。
	// verifiedChainsとその内容は変更しないでください。
	VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

	// VerifyConnectionは、通常の証明書の検証とVerifyPeerCertificateの後に、TLSクライアントまたはサーバーによって呼び出されます（nilではない場合）。非nilのエラーが返された場合、ハンドシェイクは中止されます。
	// 通常の検証が失敗した場合、このコールバックは考慮される前にハンドシェイクが中止されます。このコールバックは、InsecureSkipVerifyまたはClientAuthの設定に関係なく、再開を含むすべての接続で実行されます。
	VerifyConnection func(ConnectionState) error

	// RootCAsは、クライアントがサーバー証明書を検証する際に使用するルート証明書機関のセットを定義します。
	// RootCAsがnilの場合、TLSはホストのルートCAセットを使用します。
	RootCAs *x509.CertPool

	// NextProtosはサポートされているアプリケーションレベルのプロトコルのリストで、
	// 優先順位順に表示されます。両方のピアがALPNをサポートする場合、
	// 選択されるプロトコルはこのリストから選ばれ、相互にサポートされるプロトコルがない場合は接続が失敗します。
	// NextProtosが空であるか、ピアがALPNをサポートしていない場合、接続は成功し、
	// ConnectionState.NegotiatedProtocolは空になります。
	NextProtos []string

	// ServerNameは返された証明書のホスト名の検証に使用されます（InsecureSkipVerifyが指定されていない場合）。また、IPアドレスでない限り、クライアントのハンドシェイクに仮想ホスティングをサポートするために含まれます。
	ServerName string

	// ClientAuthは、TLSクライアント認証のサーバーポリシーを決定します。デフォルトはNoClientCertです。
	ClientAuth ClientAuthType

	// ClientCAsは、ClientAuthのポリシーに従って
	// クライアント証明書の検証が必要な場合、
	// サーバーが使用するルート証明機関のセットを定義します。
	ClientCAs *x509.CertPool

	// InsecureSkipVerifyは、クライアントがサーバーの証明書チェーンとホスト名を検証するかどうかを制御します。
	// InsecureSkipVerifyがtrueの場合、crypto/tlsパッケージは、サーバーが提示する証明書とその証明書内の任意のホスト名を受け入れます。
	// このモードでは、カスタムな検証が使用されていない限り、TLSはマンインザミドル攻撃に対して脆弱です。
	// これはテスト目的でのみ使用するか、VerifyConnectionまたはVerifyPeerCertificateと組み合わせて使用する必要があります。
	InsecureSkipVerify bool

	// CipherSuitesは有効なTLS 1.0-1.2の暗号化スイートのリストです。リストの順序は無視されます。注意事項として、TLS 1.3の暗号スイートは設定できません。
	//
	// CipherSuitesがnilの場合、安全なデフォルトのリストが使用されます。デフォルトの暗号スイートは時間とともに変更される可能性があります。
	CipherSuites []uint16

	// PreferServerCipherSuitesは古いフィールドであり、効果がありません。
	//
	// かつて、このフィールドは、サーバーがクライアントの選好またはサーバーの選好に従うかどうかを制御していました。現在、サーバーは、推測されたクライアントのハードウェア、サーバーのハードウェア、およびセキュリティを考慮した論理に基づいて、最も相互にサポートされた暗号スイートを選択します。
	//
	// Deprecated: PreferServerCipherSuitesは無視されます。
	PreferServerCipherSuites bool

	// SessionTicketsDisabledがtrueに設定されると、セッションチケットおよびPSK（再開）のサポートが無効になります。クライアントでは、ClientSessionCacheがnilの場合もセッションチケットのサポートが無効になります。
	SessionTicketsDisabled bool

	// SessionTicketKeyは、セッション再開を提供するためにTLSサーバーによって使用されます。
	// RFC 5077およびRFC 8446のPSKモードを参照してください。ゼロの場合、最初のサーバーハンドシェイクの前にランダムデータで埋められます。
	//
	// 廃止: このフィールドがゼロのままになっている場合、セッションチケットキーは自動的に毎日回転され、7日後に削除されます。
	// 回転スケジュールのカスタマイズや同じホストに対して接続を終了するサーバーの同期を行う場合は、SetSessionTicketKeysを使用してください。
	SessionTicketKey [32]byte

	// ClientSessionCacheはTLSセッション再開のためのClientSessionStateエントリのキャッシュです。
	// クライアント側でのみ使用されます。
	ClientSessionCache ClientSessionCache

	// UnwrapSessionは、[WrapSession]によって生成されたチケット/アイデンティティを、使用可能なセッションに変換するためにサーバー上で呼び出されます。
	//
	// UnwrapSessionは通常、チケット内のセッション状態を復号化するか（たとえば、[Config.EncryptTicket]を使用して）、以前に保存された状態を回復するためにチケットを使用します。セッション状態を逆シリアル化するには、[ParseSessionState]を使用する必要があります。
	//
	// UnwrapSessionがエラーを返した場合、接続は終了されます。（nil、nil）を返す場合、セッションは無視されます。crypto/tlsは、それでも返されたセッションを再開しないこともあります。
	UnwrapSession func(identity []byte, cs ConnectionState) (*SessionState, error)

	// WrapSessionはサーバーで呼び出され、セッションのチケット/アイデンティティを生成します。
	//
	// WrapSessionはセッションの状態を[SessionState.Bytes]でシリアライズする必要があります。
	// 次に、シリアライズされた状態を暗号化する（たとえば[Config.DecryptTicket]で）または状態を保存して
	// ハンドルを返すことができます。
	//
	// WrapSessionがエラーを返すと、接続は終了します。
	//
	// 警告：返り値は平文でワイヤー上やクライアントに公開されます。
	// アプリケーションはそれを暗号化し、認証する（およびキーをローテーションする）
	// または高エントロピーの識別子を返すことが責任です。
	// 正しく行わないと、現在の接続や以前の接続、将来の接続が妥協される可能性があります。
	//
	WrapSession func(ConnectionState, *SessionState) ([]byte, error)

	// MinVersionには受け入れ可能な最小のTLSバージョンが含まれています。
	//
	// デフォルトでは、クライアントとして動作する場合にはTLS 1.2が現在の最小バージョンとして使用され、サーバーとして動作する場合にはTLS 1.0が最小のバージョンとして使用されます。このパッケージでは、クライアントとしてもサーバーとしても、最小限にTLS 1.0がサポートされています。
	//
	// クライアント側のデフォルトは、GODEBUG環境変数に値"x509sha1=1"を含めることで一時的にTLS 1.0に戻すことができます。ただし、このオプションはGo 1.19で削除されます（ただし、このフィールドを明示的にVersionTLS10に設定することは引き続き可能です）。
	MinVersion uint16

	// MaxVersionには許容される最大のTLSバージョンが含まれています。
	//
	// デフォルトでは、このパッケージでサポートされている最大バージョンが使用されます。
	// 現在のバージョンはTLS 1.3です。
	MaxVersion uint16

	// CurvePreferencesには、ECDHEハンドシェイクで使用される楕円曲線が好まれる順に含まれます。空の場合、デフォルトが使用されます。クライアントは、TLS 1.3でキーシェアのタイプとして最初の選択肢を使用します。将来的には、これは変更される可能性があります。
	CurvePreferences []CurveID

	// DynamicRecordSizingDisabledはTLSレコードの適応的なサイズ調整を無効にします。
	// trueの場合、常に最大のTLSレコードサイズが使用されます。falseの場合、
	// TLSレコードのサイズはレイテンシを改善するために調整されることがあります。
	DynamicRecordSizingDisabled bool

	// Renegotiationは、再交渉がサポートされるタイプを制御します。
	// デフォルトの「なし」は、ほとんどのアプリケーションにとって正しいです。
	Renegotiation RenegotiationSupport

	// KeyLogWriterは、TLSのマスターシークレットの宛先として使用できる、NSSキーログ形式の外部プログラム（Wiresharkなど）によるTLS接続の復号化を許可するためのオプションです。
	// https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Formatを参照してください。
	// KeyLogWriterの使用はセキュリティを損なう可能性があり、デバッグ目的のみに使用するべきです。
	KeyLogWriter io.Writer

	// mutexはsessionTicketKeysとautoSessionTicketKeysを保護しています。
	mutex sync.RWMutex

	// sessionTicketKeysには、ゼロ個以上のチケットキーが含まれています。
	// 設定されている場合、それはSessionTicketKeyまたはSetSessionTicketKeysでキーが設定されたことを意味します。
	// 最初のキーは新しいチケットに使用され、後続のキーは古いチケットの復号に使用できます。
	// スライスの内容はミューテックスで保護されず、変更不可です。
	sessionTicketKeys []ticketKey

	// autoSessionTicketKeysはsessionTicketKeysと似ていますが、自動回転ロジックによって所有されています。Config.ticketKeysを参照してください。
	autoSessionTicketKeys []ticketKey
}

<<<<<<< HEAD
// Cloneはcの浅いクローンを返します。cがnilの場合はnilを返します。TLSクライアントやサーバーによって
// 同時に使用されているConfigをクローンすることは安全です。
=======
// Clone returns a shallow clone of c or nil if c is nil. It is safe to clone a [Config] that is
// being used concurrently by a TLS client or server.
>>>>>>> upstream/master
func (c *Config) Clone() *Config

// SetSessionTicketKeysはサーバーのセッションチケットのキーを更新します。
//
// 新しいチケットを作成する際に最初のキーが使用され、すべてのキーはチケットの解読に使用できます。
// セッションチケットのキーをローテーションするために、この関数を実行しても問題ありません（サーバーが実行中である場合）。
// 関数はキーが空の場合、パニックを発生させます。
//
// この関数を呼び出すと、自動的なセッションチケットキーのローテーションが無効になります。
//
// 同じホストに接続を終了する複数のサーバーがある場合、すべてのサーバーは同じセッションチケットのキーを持つ必要があります。
// セッションチケットのキーが漏洩した場合、以前に記録されたおよび将来のTLS接続でこれらのキーが使用される可能性があります。
// これにより、接続が危険にさらされる可能性があります。
func (c *Config) SetSessionTicketKeys(keys [][32]byte)

// SupportsCertificateは、提供された証明書が
// ClientHelloを送信したクライアントによってサポートされている場合にはnilを返します。そうでない場合は、互換性のない理由を説明するエラーを返します。
//
<<<<<<< HEAD
// このClientHelloInfoがGetConfigForClientまたはGetCertificateコールバックに渡された場合、このメソッドは関連するConfigを考慮に入れます。ただし、GetConfigForClientが異なるConfigを返す場合、このメソッドでは変更を考慮することができません。
=======
// If this [ClientHelloInfo] was passed to a GetConfigForClient or GetCertificate
// callback, this method will take into account the associated [Config]. Note that
// if GetConfigForClient returns a different [Config], the change can't be
// accounted for by this method.
>>>>>>> upstream/master
//
// c.Leafが設定されていない場合、この関数はx509.ParseCertificateを呼び出しますが、それはかなりのパフォーマンスコストを伴うことになります。
func (chi *ClientHelloInfo) SupportsCertificate(c *Certificate) error

// SupportsCertificateは、提供された証明書がCertificateRequestを送信したサーバーによってサポートされている場合はnilを返します。それ以外の場合、非互換性の理由を説明するエラーが返されます。
func (cri *CertificateRequestInfo) SupportsCertificate(c *Certificate) error

// BuildNameToCertificateはc.Certificatesを解析し、各リーフ証明書のCommonNameとSubjectAlternateNameフィールドからc.NameToCertificateを構築します。
// 廃止されました: NameToCertificateは特定の名前に対して単一の証明書の関連付けしか許可しません。そのフィールドをnilのままにしておき、ライブラリに最初に互換性のあるチェーンを選択させます。
func (c *Config) BuildNameToCertificate()

// ボタン構造体は、最初にリーフ（最下位）のボタンから始まり、その上位にある1つ以上のボタンのチェーンです。
type Certificate struct {
	Certificate [][]byte

	// PrivateKeyは、Leafの公開鍵に対応する秘密鍵を含んでいます。
	// これは、RSA、ECDSA、またはEd25519 PublicKeyを使用してcrypto.Signerを実装する必要があります。
	// TLS 1.2までのサーバーの場合、RSA PublicKeyを使用してcrypto.Decrypterも実装できます。
	PrivateKey crypto.PrivateKey

	// SupportedSignatureAlgorithmsは、PrivateKeyが使用できる署名アルゴリズムを制限するオプションのリストです。
	SupportedSignatureAlgorithms []SignatureScheme

	// OCSPStapleには、リクエストするクライアントに提供されるオプションのOCSP応答が含まれています。
	OCSPStaple []byte

	// SignedCertificateTimestampsには、要求するクライアントに提供されるオプションの署名付き証明書タイムスタンプのリストが含まれています。
	SignedCertificateTimestamps [][]byte

	// Leafは、パースされたリーフ証明書の形式であり、x509.ParseCertificateを使用して初期化することができます。
	// デフォルトではnilである場合、リーフ証明書は必要に応じてパースされます。
	Leaf *x509.Certificate
}

<<<<<<< HEAD
// NewLRUClientSessionCacheは、与えられた容量を使用してLRU戦略を採用したClientSessionCacheを返します。容量が1未満の場合、代わりにデフォルトの容量が使用されます。
=======
// NewLRUClientSessionCache returns a [ClientSessionCache] with the given
// capacity that uses an LRU strategy. If capacity is < 1, a default capacity
// is used instead.
>>>>>>> upstream/master
func NewLRUClientSessionCache(capacity int) ClientSessionCache

// CertificateVerificationError は、ハンドシェイク中に証明書の検証が失敗した場合に返されます。
type CertificateVerificationError struct {
	// UnverifiedCertificatesおよびその内容は変更しないでください。
	UnverifiedCertificates []*x509.Certificate
	Err                    error
}

func (e *CertificateVerificationError) Error() string

func (e *CertificateVerificationError) Unwrap() error
