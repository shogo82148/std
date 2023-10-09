// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/time"
)

type InvalidReason int

const (

	// NotAuthorizedToSignは、他のCA証明書としてマークされていない証明書によって署名された場合に発生する結果です。
	NotAuthorizedToSign InvalidReason = iota

	// VerifyOptionsで指定された時間に基づき、証明書が期限切れとなった結果を返します。
	Expired

	// CANotAuthorizedForThisNameは、中間またはルート証明書に、葉証明書でDNSまたはその他の名前（IPアドレスを含む）を許可しない制約がある場合に発生します。
	CANotAuthorizedForThisName

	// TooManyIntermediatesは、パスの長さ制約が違反された場合に発生します。
	TooManyIntermediates

	// IncompatibleUsageは、証明書のキーの使用法が異なる目的でのみ使用できることを示す場合に発生します。
	IncompatibleUsage

	// NameMismatchは、親の証明書のサブジェクト名が子の発行者名と一致しない場合に発生します。
	NameMismatch
	// NameConstraintsWithoutSANsは、過去のエラーであり、もはや返されなくなりました。
	NameConstraintsWithoutSANs

	// UnconstrainedNameは、CA証明書に許容される名前制約が含まれているが、
	// リーフ証明書にはサポートされていないまたは制約のないタイプの名前が含まれている場合の結果です。
	UnconstrainedName

	// TooManyConstraintsは、証明書を検証するために必要な比較操作の数が、VerifyOptions.MaxConstraintComparisonsで設定された制限を超える場合に発生します。この制限は、CPU時間の過剰な消費を防ぐために存在します。
	TooManyConstraints

	// CANotAuthorizedForExtKeyUsage は、中間証明書またはルート証明書が要求された拡張キー使用法を許可しない場合に発生します。
	CANotAuthorizedForExtKeyUsage
)

// CertificateInvalidErrorは、奇妙なエラーが発生した場合に結果が返されます。このライブラリのユーザーはおそらく、これらのエラーを統一的に処理したいと考えるでしょう。
type CertificateInvalidError struct {
	Cert   *Certificate
	Reason InvalidReason
	Detail string
}

func (e CertificateInvalidError) Error() string

// HostnameErrorは、許可された名前のセットが要求された名前と一致しない場合に発生します。
type HostnameError struct {
	Certificate *Certificate
	Host        string
}

func (h HostnameError) Error() string

// UnknownAuthorityErrorは、証明書の発行者が不明な場合に発生します。
type UnknownAuthorityError struct {
	Cert *Certificate

	// hintErrには、権限が見つからなかった原因を特定するのに役立つかもしれないエラーが含まれています。
	hintErr error

	// hintCertには、hintErrのエラーのために却下された可能性のある認証局の証明書が含まれています。
	hintCert *Certificate
}

func (e UnknownAuthorityError) Error() string

// システムのルート証明書の読み込みに失敗した場合、SystemRootsErrorが発生します。
type SystemRootsError struct {
	Err error
}

func (se SystemRootsError) Error() string

func (se SystemRootsError) Unwrap() error

// VerifyOptionsにはCertificate.Verifyのパラメータが含まれています。
type VerifyOptions struct {

	// DNSNameが設定されている場合は、Certificate.VerifyHostnameまたはプラットフォームの検証器で葉証明書と照合されます。
	DNSName string

	// Intermediatesは、信頼アンカーではないが、リーフ証明書からルート証明書までのチェーンを形成するために使用できるオプションの証明書のプールです。
	Intermediates *CertPool

	// Rootsは、リーフ証明書がチェーンアップするために必要な信頼できるルート証明書のセットです。nilの場合、システムのルートまたはプラットフォームの検証器が使用されます。
	Roots *CertPool

	// CurrentTimeは、チェーン内のすべての証明書の有効性を確認するために使用されます。
	// ゼロの場合、現在の時刻が使用されます。
	CurrentTime time.Time

	// KeyUsagesは受け入れ可能な拡張キー利用法（Extended Key Usage）の値を指定します。リストされた値のいずれかを許可する場合、チェーンは受け入れられます。空のリストはExtKeyUsageServerAuthを意味します。どんなキー利用法でも受け入れる場合は、ExtKeyUsageAnyを含めてください。
	KeyUsages []ExtKeyUsage

	// MaxConstraintComparisionsは、指定された証明書の名前制約をチェックする際に行う比較の最大数です。
	// ゼロの場合、適切なデフォルト値が使用されます。この制限によって、病的な証明書が検証時に過剰なCPU時間を消費するのを防ぎます。
	// この制限は、プラットフォームの検証ツールには適用されません。
	MaxConstraintComparisions int
}

// Verifyは、オプションのRootsの証明書を使用して、cからaの証明書までの1つ以上のチェーンを構築し、検証を試みます。成功すると、最初のチェーン要素はcで、最後の要素はopts.Rootsから来ます。
// opts.Rootsがnilの場合、プラットフォームの検証プログラムが使用される可能性があります。この場合、検証の詳細は以下の説明と異なる場合があります。システムのルートが利用できない場合、返されるエラーはSystemRootsError型です。
// 中間証明書の名前制約は、opts.DNSNameだけでなく、チェーン内で要求されるすべての名前に適用されます。したがって、中間証明書が許可しない場合でも、葉がexample.comを要求することは無効です。ただし、DirectoryName制約はサポートされていません。
// 名前制約の検証は、RFC 5280の規則に従います。ただし、DNS名の制約は、電子メールやURIで定義された先頭のピリオド形式を使用できます。制約に先頭のピリオドがある場合、有効な制約名を構成するために少なくとも1つの追加のラベルが前置される必要があります。
// 拡張キー使用（Extended Key Usage）の値は、チェーンが入れ子になった形で強制されます。したがって、EKUを列挙する中間証明書やルート証明書は、そのリストに含まれていないEKUを持つ葉がアサートすることを防ぎます。（これは明示されていませんが、CAが発行できる証明書の種類を制限するために一般的に行われています。）
// SHA1WithRSAおよびECDSAWithSHA1の署名を使用する証明書はサポートされておらず、チェーンの構築には使用されません。
// 返されるチェーンのc以外の証明書は変更しないでください。
// 警告：この関数はリボケーション（証明書の失効チェック）を行いません。
func (c *Certificate) Verify(opts VerifyOptions) (chains [][]*Certificate, err error)

// VerifyHostnameは指定されたホストに対して、cが有効な証明書であればnilを返します。
// それ以外の場合、ミスマッチを説明するエラーを返します。
//
// IPアドレスはオプションで角括弧で囲まれ、IPAddressesフィールドと照合されます。
// それ以外の名前は、DNSNamesフィールドで大文字小文字を区別せずにチェックされます。
// 名前が有効なホスト名である場合、証明書のフィールドにはワイルドカードが完全な左端のラベルとして含まれている場合があります（例: *.example.com）。
//
// レガシーのCommon Nameフィールドは無視されることに注意してください。
func (c *Certificate) VerifyHostname(h string) error
