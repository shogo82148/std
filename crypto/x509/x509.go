// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージx509はX.509規格の一部を実装しています。
//
// 証明書、証明書署名要求、証明書失効リスト、エンコードされた公開および秘密鍵の解析および生成を可能にします。
// チェーンビルダーを備えた証明書検証機能も提供します。
//
// このパッケージはIETF（RFC 2459/3280/5280）によって定義されたX.509技術プロファイルを対象としており、CA/Browser Forum Baseline Requirementsによってさらに制限されています。
// 主な目標は、公に信頼されるTLS証明書エコシステムとそのポリシーおよび制約との互換性を提供することであり、これらのプロファイル外の機能には最小限のサポートしかありません。
//
// macOSおよびWindowsでは、証明書の検証はシステムAPIによって処理されますが、パッケージはオペレーティングシステム間で一貫した検証ルールを適用することを目指しています。
package x509

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509/pkix"
	"github.com/shogo82148/std/encoding/asn1"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"
)

<<<<<<< HEAD
// ParsePKIXPublicKeyはPKIX、ASN.1 DER形式の公開鍵を解析します。エンコードされた公開鍵はSubjectPublicKeyInfo構造体です（RFC 5280、セクション4.1を参照）。
// *rsa.PublicKey、*dsa.PublicKey、*ecdsa.PublicKey、ed25519.PublicKey（ポインタではない）、または*ecdh.PublicKey（X25519用）を返します。
// 将来的にはさらに多くの種類がサポートされるかもしれません。
// この種類の鍵は、一般的に「PUBLIC KEY」というタイプのPEMブロックでエンコードされます。
=======
// ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form. The encoded
// public key is a SubjectPublicKeyInfo structure (see RFC 5280, Section 4.1).
//
// It returns a *[rsa.PublicKey], *[dsa.PublicKey], *[ecdsa.PublicKey],
// [ed25519.PublicKey] (not a pointer), or *[ecdh.PublicKey] (for X25519).
// More types might be supported in the future.
//
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
>>>>>>> upstream/master
func ParsePKIXPublicKey(derBytes []byte) (pub any, err error)

// MarshalPKIXPublicKeyは公開鍵をPKIX、ASN.1 DER形式に変換します。
// エンコードされた公開鍵はSubjectPublicKeyInfo構造体です
// （RFC 5280、セクション4.1を参照）。
//
<<<<<<< HEAD
// 現在サポートされているキータイプは次のとおりです：*rsa.PublicKey、*ecdsa.PublicKey、ed25519.PublicKey（ポインタではありません）、*ecdh.PublicKey。
// サポートされていないキータイプはエラーとなります。
=======
// The following key types are currently supported: *[rsa.PublicKey],
// *[ecdsa.PublicKey], [ed25519.PublicKey] (not a pointer), and *[ecdh.PublicKey].
// Unsupported key types result in an error.
>>>>>>> upstream/master
//
// この種類のキーは一般的には"type 'PUBLIC KEY'のPEMブロックでエンコードされます。
func MarshalPKIXPublicKey(pub any) ([]byte, error)

type SignatureAlgorithm int

const (
	UnknownSignatureAlgorithm SignatureAlgorithm = iota

	MD2WithRSA
	MD5WithRSA
	SHA1WithRSA
	SHA256WithRSA
	SHA384WithRSA
	SHA512WithRSA
	DSAWithSHA1
	DSAWithSHA256
	ECDSAWithSHA1
	ECDSAWithSHA256
	ECDSAWithSHA384
	ECDSAWithSHA512
	SHA256WithRSAPSS
	SHA384WithRSAPSS
	SHA512WithRSAPSS
	PureEd25519
)

func (algo SignatureAlgorithm) String() string

type PublicKeyAlgorithm int

const (
	UnknownPublicKeyAlgorithm PublicKeyAlgorithm = iota
	RSA
	DSA
	ECDSA
	Ed25519
)

func (algo PublicKeyAlgorithm) String() string

// KeyUsageは、与えられたキーに対して有効なアクションのセットを表します。これはKeyUsage*の定数のビットマップです。
type KeyUsage int

const (
	KeyUsageDigitalSignature KeyUsage = 1 << iota
	KeyUsageContentCommitment
	KeyUsageKeyEncipherment
	KeyUsageDataEncipherment
	KeyUsageKeyAgreement
	KeyUsageCertSign
	KeyUsageCRLSign
	KeyUsageEncipherOnly
	KeyUsageDecipherOnly
)

// ExtKeyUsageは、与えられたキーに対して有効な拡張アクションのセットを表します。
// ExtKeyUsage*の各定数は、ユニークなアクションを定義しています。
type ExtKeyUsage int

const (
	ExtKeyUsageAny ExtKeyUsage = iota
	ExtKeyUsageServerAuth
	ExtKeyUsageClientAuth
	ExtKeyUsageCodeSigning
	ExtKeyUsageEmailProtection
	ExtKeyUsageIPSECEndSystem
	ExtKeyUsageIPSECTunnel
	ExtKeyUsageIPSECUser
	ExtKeyUsageTimeStamping
	ExtKeyUsageOCSPSigning
	ExtKeyUsageMicrosoftServerGatedCrypto
	ExtKeyUsageNetscapeServerGatedCrypto
	ExtKeyUsageMicrosoftCommercialCodeSigning
	ExtKeyUsageMicrosoftKernelCodeSigning
)

// CertificateはX.509証明書を表します。
type Certificate struct {
	Raw                     []byte
	RawTBSCertificate       []byte
	RawSubjectPublicKeyInfo []byte
	RawSubject              []byte
	RawIssuer               []byte

	Signature          []byte
	SignatureAlgorithm SignatureAlgorithm

	PublicKeyAlgorithm PublicKeyAlgorithm
	PublicKey          any

	Version             int
	SerialNumber        *big.Int
	Issuer              pkix.Name
	Subject             pkix.Name
	NotBefore, NotAfter time.Time
	KeyUsage            KeyUsage

	// Extensionsには生のX.509拡張が含まれています。証明書を解析する際、
	// このフィールドを使用して、このパッケージによって解析されない非致命的な拡張を抽出できます。証明書をマーシャリングする際、Extensionsフィールドは無視されます。ExtraExtensionsを参照してください。
	Extensions []pkix.Extension

	// ExtraExtensionsには、任意のマーシャル化された証明書にコピーして使用される、拡張機能が含まれています。値は、他のフィールドに基づいて生成される拡張機能を上書きします。証明書の解析時にはExtraExtensionsフィールドは埋められませんが、Extensionsを参照してください。
	ExtraExtensions []pkix.Extension

	// UnhandledCriticalExtensionsは、解析時に（完全に）処理されなかった拡張IDのリストを含んでいます。このスライスが空でない場合、検証は失敗します。ただし、すべての重要な拡張を理解できるOSライブラリに検証が委任されている場合は除きます。
	//
	// ユーザーはExtensionsを使用してこれらの拡張にアクセスし、処理されたと信じられる要素をこのスライスから削除することができます。
	UnhandledCriticalExtensions []asn1.ObjectIdentifier

	ExtKeyUsage        []ExtKeyUsage
	UnknownExtKeyUsage []asn1.ObjectIdentifier

	// BasicConstraintsValidは、IsCA、MaxPathLen、およびMaxPathLenZeroが正常であるかどうかを示す。
	BasicConstraintsValid bool
	IsCA                  bool

	// MaxPathLen と MaxPathLenZero は BasicConstraints の "pathLenConstraint" の存在と値を指します。
	//
	// 証明書を解析する際に、正の非ゼロの MaxPathLen はフィールドが指定されたことを示し、-1 は指定されなかったことを示し、MaxPathLenZero が true の場合はフィールドが明示的にゼロに設定されたことを示します。MaxPathLen == 0 かつ MaxPathLenZero == false の場合は -1 と同等に扱われるべきです。
	//
	// 証明書を生成する際、未設定の pathLenConstraint は MaxPathLen == -1 または MaxPathLen と MaxPathLenZero の両方にゼロ値を使用することでリクエストすることができます。
	MaxPathLen int

	// MaxPathLenZeroは、BasicConstraintsValid==trueであるとき、
	// MaxPathLen==0は実際の最大パス長さが0であると解釈されることを示しています。
	// それ以外の場合、この組み合わせはMaxPathLenが設定されていないと解釈されます。
	MaxPathLenZero bool

	SubjectKeyId   []byte
	AuthorityKeyId []byte

	// RFC 5280、4.2.2.1（権限情報アクセス）
	OCSPServer            []string
	IssuingCertificateURL []string

	// Subject Alternate Nameの値。（ただし、パースされた証明書に無効な値が含まれている場合、これらの値は有効ではない場合があります。例えば、DNSNamesの要素が有効なDNSドメイン名であるとは限りません。）
	DNSNames       []string
	EmailAddresses []string
	IPAddresses    []net.IP
	URIs           []*url.URL

	// 名前の制約
	PermittedDNSDomainsCritical bool
	PermittedDNSDomains         []string
	ExcludedDNSDomains          []string
	PermittedIPRanges           []*net.IPNet
	ExcludedIPRanges            []*net.IPNet
	PermittedEmailAddresses     []string
	ExcludedEmailAddresses      []string
	PermittedURIDomains         []string
	ExcludedURIDomains          []string

	// CRL配布ポイント
	CRLDistributionPoints []string

	PolicyIdentifiers []asn1.ObjectIdentifier
}

// ErrUnsupportedAlgorithmは、現在実装されていないアルゴリズムを使用して操作を実行しようとした結果です。
var ErrUnsupportedAlgorithm = errors.New("x509: cannot verify signature: algorithm unimplemented")

<<<<<<< HEAD
// InsecureAlgorithmErrorは、署名の生成に使用されるSignatureAlgorithmが安全でないことを示し、署名が拒否されたことを示します。
=======
// An InsecureAlgorithmError indicates that the [SignatureAlgorithm] used to
// generate the signature is not secure, and the signature has been rejected.
>>>>>>> upstream/master
//
// SHA-1署名のサポートを一時的に復元するには、GODEBUG環境変数に値"x509sha1=1"を含めます。ただし、このオプションは将来のリリースで削除される予定です。
type InsecureAlgorithmError SignatureAlgorithm

func (e InsecureAlgorithmError) Error() string

// ConstraintViolationErrorは、証明書によって許可されていない要求された使用方法がある場合に発生します。例えば、公開キーが証明書署名キーではない場合に署名のチェックを行うことなどです。
type ConstraintViolationError struct{}

func (ConstraintViolationError) Error() string

func (c *Certificate) Equal(other *Certificate) bool

// CheckSignatureFromは、c上の署名が親からの有効な署名であるかを検証します。
//
// これは非常に限定的なチェックを行う低レベルAPIであり、完全なパス検証ではありません。
// ほとんどのユーザーは[Certificate.Verify]を使用するべきです。
func (c *Certificate) CheckSignatureFrom(parent *Certificate) error

// CheckSignatureは署名がsigned fromの公開鍵の有効な署名であることを検証します。
//
// これは証明書に対して妥当性チェックを行わない低レベルのAPIです。
//
// [MD5WithRSA]の署名は拒否され、[SHA1WithRSA]と[ECDSAWithSHA1]の署名は現在受け入れられています。
func (c *Certificate) CheckSignature(algo SignatureAlgorithm, signed, signature []byte) error

// CheckCRLSignatureは、crlの署名がcからのものであることをチェックします。
//
<<<<<<< HEAD
// 廃止予定：RevocationList.CheckSignatureFromを使用してください。
=======
// Deprecated: Use [RevocationList.CheckSignatureFrom] instead.
>>>>>>> upstream/master
func (c *Certificate) CheckCRLSignature(crl *pkix.CertificateList) error

type UnhandledCriticalExtension struct{}

func (h UnhandledCriticalExtension) Error() string

// CreateCertificateは、テンプレートに基づいて新しいX.509 v3証明書を作成します。
// 現在のテンプレートの以下のメンバーが使用されています：
//
//   - AuthorityKeyId
//   - BasicConstraintsValid
//   - CRLDistributionPoints
//   - DNSNames
//   - EmailAddresses
//   - ExcludedDNSDomains
//   - ExcludedEmailAddresses
//   - ExcludedIPRanges
//   - ExcludedURIDomains
//   - ExtKeyUsage
//   - ExtraExtensions
//   - IPAddresses
//   - IsCA
//   - IssuingCertificateURL
//   - KeyUsage
//   - MaxPathLen
//   - MaxPathLenZero
//   - NotAfter
//   - NotBefore
//   - OCSPServer
//   - PermittedDNSDomains
//   - PermittedDNSDomainsCritical
//   - PermittedEmailAddresses
//   - PermittedIPRanges
//   - PermittedURIDomains
//   - PolicyIdentifiers
//   - SerialNumber
//   - SignatureAlgorithm
//   - Subject
//   - SubjectKeyId
//   - URIs
//   - UnknownExtKeyUsage
//
// 証明書は親によって署名されます。親がテンプレートと等しい場合、証明書は自己署名です。pubパラメータは生成される証明書の公開鍵であり、privは署名者の秘密鍵です。
//
// 返されるスライスはDERエンコーディングされた証明書です。
//
// 現在サポートされている鍵のタイプは*rsa.PublicKey、*ecdsa.PublicKey、およびed25519.PublicKeyです。pubはサポートされている鍵のタイプである必要があり、privはサポートされている公開鍵を持つcrypto.Signerである必要があります。
//
// AuthorityKeyIdは、親のSubjectKeyIdから取得されます（存在する場合）、ただし証明書が自己署名でない場合はテンプレートの値が使用されます。
//
// テンプレートのSubjectKeyIdが空で、テンプレートがCAである場合、SubjectKeyIdは公開鍵のハッシュから生成されます。
func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv any) ([]byte, error)

<<<<<<< HEAD
// ParseCRLは指定されたバイトからCRLを解析します。PEMエンコードされたCRLがDERエンコードされるべき場所に表示されることがよくありますが、この関数は前方にゴミがない限り、PEMエンコーディングを透過的に処理します。
// 廃止予定: 代わりにParseRevocationListを使用してください。
=======
// ParseCRL parses a CRL from the given bytes. It's often the case that PEM
// encoded CRLs will appear where they should be DER encoded, so this function
// will transparently handle PEM encoding as long as there isn't any leading
// garbage.
//
// Deprecated: Use [ParseRevocationList] instead.
>>>>>>> upstream/master
func ParseCRL(crlBytes []byte) (*pkix.CertificateList, error)

// ParseDERCRLは与えられたバイトからDER形式でエンコードされたCRLをパースします。
//
<<<<<<< HEAD
// 非推奨: 代わりにParseRevocationListを使用してください。
=======
// Deprecated: Use [ParseRevocationList] instead.
>>>>>>> upstream/master
func ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error)

// CreateCRLは、指定された失効した証明書のリストを含む、この証明書によって署名されたDERエンコードされたCRLを返します。
//
<<<<<<< HEAD
// 廃止予定: このメソッドはRFC 5280準拠のX.509 v2 CRLを生成しません。
// 標準に準拠したCRLを生成するためには、代わりにCreateRevocationListを使用してください。
=======
// Deprecated: this method does not generate an RFC 5280 conformant X.509 v2 CRL.
// To generate a standards compliant CRL, use [CreateRevocationList] instead.
>>>>>>> upstream/master
func (c *Certificate) CreateCRL(rand io.Reader, priv any, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) (crlBytes []byte, err error)

// CertificateRequestはPKCS #10、証明書署名リクエストを表します。
type CertificateRequest struct {
	Raw                      []byte
	RawTBSCertificateRequest []byte
	RawSubjectPublicKeyInfo  []byte
	RawSubject               []byte

	Version            int
	Signature          []byte
	SignatureAlgorithm SignatureAlgorithm

	PublicKeyAlgorithm PublicKeyAlgorithm
	PublicKey          any

	Subject pkix.Name

	// Attributesには、以下のCSR属性が含まれています。pkix.AttributeTypeAndValueSETとして解析できます。
	//
	// 廃止予定: 解析および生成には、requestedExtensions属性の代わりにExtensionsおよびExtraExtensionsを使用してください。
	Attributes []pkix.AttributeTypeAndValueSET

	// Extensionsは、すべてのリクエストされた拡張子を生の形式で保持しています。CSRを解析する際に、このパッケージで解析されない拡張子を抽出するために使用できます。
	Extensions []pkix.Extension

	// ExtraExtensionsは、任意のCSRにコピーされる拡張機能を含みます。
	// CreateCertificateRequestによってマーシャリングされます。
	// 値は他のフィールドに基づいて生成される拡張機能を上書きしますが、
	// Attributesで指定された拡張機能によっては上書きされます。
	//
	// ExtraExtensionsフィールドはParseCertificateRequestでは使用されず、
	// 代わりにExtensionsを参照してください。
	ExtraExtensions []pkix.Extension

	// Subject Alternate Nameの値。
	DNSNames       []string
	EmailAddresses []string
	IPAddresses    []net.IP
	URIs           []*url.URL
}

// CreateCertificateRequestは、テンプレートを基に新しい証明書リクエストを作成します。テンプレートの以下のメンバーが使用されます：
//   - SignatureAlgorithm
//   - Subject
//   - DNSNames
//   - EmailAddresses
//   - IPAddresses
//   - URIs
//   - ExtraExtensions
//   - Attributes (非推奨)
//
// privはCSRに署名するための秘密鍵であり、対応する公開鍵はCSRに含まれます。privはcrypto.Signerを実装しており、そのPublic()メソッドは*rsa.PublicKeyまたは*ecdsa.PublicKeyまたはed25519.PublicKeyを返さなければなりません。(*rsa.PrivateKey、*ecdsa.PrivateKey、またはed25519.PrivateKeyもこれを満たします。)
// 返されるスライスはDERエンコードされた証明書リクエストです。
func CreateCertificateRequest(rand io.Reader, template *CertificateRequest, priv any) (csr []byte, err error)

// ParseCertificateRequestは与えられたASN.1 DERデータから単一の証明書リクエストを解析します。
func ParseCertificateRequest(asn1Data []byte) (*CertificateRequest, error)

// CheckSignatureはcの署名が有効かどうかを報告します。
func (c *CertificateRequest) CheckSignature() error

// RevocationListEntryは、CRLのrevokedCertificatesシーケンスのエントリを表します。
type RevocationListEntry struct {

	// Raw は revokedCertificates エントリの生のバイトを含んでいます。
	// CRL を解析する際に設定され、CRL を生成する際には無視されます。
	Raw []byte

	// SerialNumberは失効証明書のシリアル番号を表します。CRLを作成する際に使用され、
	// CRLを解析する際にも設定されます。nilであってはいけません。
	SerialNumber *big.Int

	// RevocationTimeは証明書の失効日時を表します。
	// CRLを作成する際に使用され、CRLを解析する際に設定されます。
	// ゼロの時間であってはなりません。
	RevocationTime time.Time

	// ReasonCodeは、RFC 5280セクション5.3.1で指定された整数の列挙値を使用して、回復の理由を表します。CRLを作成する場合、ゼロ値はreasonCode拡張機能が省略される結果になります。CRLを解析する際、ゼロ値はreasonCode拡張機能が存在しないこと（0/Unspecifiedのデフォルト回収理由を意味する）を表すか、reasonCode拡張機能が存在し、明示的に0/Unspecifiedの値を含んでいることを表す可能性があります（これはDERエンコーディングルールによらないで発生する可能性がありますが、実際に発生することがあります）。
	ReasonCode int

	// Extensionsには生のX.509拡張が含まれています。CRLエントリを解析する際、
	// このフィールドを使用して、このパッケージでパースされない非クリティカルな拡張を取得できます。
	// CRLエントリをマーシャル化する際、Extensionsフィールドは無視されます。ExtraExtensionsを参照してください。
	Extensions []pkix.Extension

	// ExtraExtensionsには、任意のマーシャルされたCRLエントリにコピーするための拡張機能が含まれています。値は、他のフィールドに基づいて生成される拡張機能を上書きします。ExtraExtensionsフィールドは、CRLエントリの解析時には値が設定されません。Extensionsを参照してください。
	ExtraExtensions []pkix.Extension
}

<<<<<<< HEAD
// RevocationList は RFC 5280 で指定されている Certificate Revocation List (CRL) を表します。
=======
// RevocationList represents a [Certificate] Revocation List (CRL) as specified
// by RFC 5280.
>>>>>>> upstream/master
type RevocationList struct {

	// Raw はCRL（tbsCertList、signatureAlgorithm、およびsignatureValue）の完全なASN.1 DERコンテンツを含んでいます。
	Raw []byte

	// RawTBSRevocationList はASN.1 DERのtbsCertList部分のみを含みます。
	RawTBSRevocationList []byte
	// RawIssuerにはDERエンコードされた発行者が含まれています。
	RawIssuer []byte

	// Issuerには発行証明書のDNが含まれています。
	Issuer pkix.Name

	// AuthorityKeyIdは、発行証明書に関連付けられた公開鍵を識別するために使用されます。CRLを解析する際、authorityKeyIdentifier拡張から取得されます。CRLを作成する際には無視されます。拡張は発行証明書自体から取得されます。
	AuthorityKeyId []byte

	Signature []byte

	// SignatureAlgorithmは、CRLを署名する際に使用する署名アルゴリズムを決定するために使用されます。
	// もし0の場合、署名キーのデフォルトアルゴリズムが使用されます。
	SignatureAlgorithm SignatureAlgorithm

	// RevokedCertificateEntriesは、CRLのrevokedCertificatesシーケンスを表します。
	// CRLを作成するときに使用され、CRLを解析するときにも入力されます。
	// CRLを作成する際には、空またはnilである場合、revokedCertificates ASN.1シーケンスはCRLから完全に省略されます。
	RevokedCertificateEntries []RevocationListEntry

	// RevokedCertificatesはRevokedCertificateEntriesが空の場合、
	// CRL内のrevokedCertificatesシーケンスを埋めるために使用されます。
	// RevokedCertificatesは空またはnilである場合、空のCRLが作成されます。
	//
	// Deprecated: 代わりにRevokedCertificateEntriesを使用してください。
	RevokedCertificates []pkix.RevokedCertificate

	// Numberは、CRL内のX.509 v2 cRLNumber拡張を埋めるために使用されます。
	// これは特定のCRLスコープとCRL発行者に対して単調に増加するシーケンス番号である必要があります。
	// また、CRLを解析する際には、cRLNumber拡張からも値が入力されます。
	Number *big.Int

	// ThisUpdateはCRLのthisUpdateフィールドに格納されるために使用され、CRLの発行日を示します。
	ThisUpdate time.Time

	// NextUpdateはCRLのnextUpdateフィールドを埋めるために使用されます。これは次のCRLが発行される日付を示しています。NextUpdateはThisUpdateよりも大きくなければなりません。
	NextUpdate time.Time

	// Extensionsは生のX.509拡張を含んでいます。CRLを作成する際は、
	// Extensionsフィールドは無視されます。ExtraExtensionsを参照してください。
	Extensions []pkix.Extension

	// ExtraExtensionsには、CRLに直接追加する必要がある追加の拡張機能が含まれています。
	ExtraExtensions []pkix.Extension
}

<<<<<<< HEAD
// CreateRevocationListは、テンプレートに基づいてRFC 5280に準拠した新しいX.509 v2証明書失効リストを作成します。
// CRLは、privによって署名されます。これは、発行者証明書の公開キーに関連付けられた秘密キーである必要があります。
// 発行者はnilではなく、キーカオ必須使用方法のcrlSignビットが設定されている必要があります。
// 発行者の識別名CRLフィールドと権限キー識別子拡張は、発行者証明書を使用してポピュレートされます。発行者にはSubjectKeyIdが設定されている必要があります。
=======
// CreateRevocationList creates a new X.509 v2 [Certificate] Revocation List,
// according to RFC 5280, based on template.
//
// The CRL is signed by priv which should be the private key associated with
// the public key in the issuer certificate.
//
// The issuer may not be nil, and the crlSign bit must be set in [KeyUsage] in
// order to use it as a CRL issuer.
//
// The issuer distinguished name CRL field and authority key identifier
// extension are populated using the issuer certificate. issuer must have
// SubjectKeyId set.
>>>>>>> upstream/master
func CreateRevocationList(rand io.Reader, template *RevocationList, issuer *Certificate, priv crypto.Signer) ([]byte, error)

// CheckSignatureFromは、rlの署名が発行元の有効な署名であることを確認します。
func (rl *RevocationList) CheckSignatureFrom(parent *Certificate) error
