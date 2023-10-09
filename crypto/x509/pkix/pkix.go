// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージ pkix には、ASN.1 パースおよび X.509 証明書、CRL、OCSP のシリアル化に使用される共有の低レベルの構造体が含まれています。
package pkix

import (
	"github.com/shogo82148/std/encoding/asn1"
	"github.com/shogo82148/std/math/big"
	"github.com/shogo82148/std/time"
)

// AlgorithmIdentifierは、同名のASN.1構造を表します。RFC 5280のセクション4.1.1.2を参照してください。
type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}

type RDNSequence []RelativeDistinguishedNameSET

// Stringは、シーケンスrの文字列表現を返します。
// おおよそRFC 2253の特定名の構文に従います。
func (r RDNSequence) String() string

type RelativeDistinguishedNameSET []AttributeTypeAndValue

// AttributeTypeAndValueは、RFC 5280、セクション4.1.2.4で同名のASN.1構造体を反映しています。
type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value any
}

// AttributeTypeAndValueSETは、RFC 2986（PKCS＃10）からのAttributeTypeAndValueシーケンスの集合を表す。
type AttributeTypeAndValueSET struct {
	Type  asn1.ObjectIdentifier
	Value [][]AttributeTypeAndValue `asn1:"set"`
}

// Extensionは同名のASN.1構造を表します。RFC 5280、セクション4.2を参照してください。
type Extension struct {
	Id       asn1.ObjectIdentifier
	Critical bool `asn1:"optional"`
	Value    []byte
}

// NameはX.509の識別名を表します。これにはDNの一般的な要素のみが含まれます。なお、NameはX.509の構造の近似値です。正確な表現が必要な場合は、生のsubjectまたはissuerをRDNSequenceとしてasn1.Unmarshalしてください。
type Name struct {
	Country, Organization, OrganizationalUnit []string
	Locality, Province                        []string
	StreetAddress, PostalCode                 []string
	SerialNumber, CommonName                  string

	// Namesにはすべての解析された属性が含まれています。識別名を解析する際に、
	// このフィールドを使用して、このパッケージでは解析されない非標準の属性を抽出できます。
	// RDNSequencesに統合する際には、Namesフィールドは無視されますが、ExtraNamesを参照してください。
	Names []AttributeTypeAndValue

	// ExtraNamesには、マーシャリングされる任意の識別名にコピーされる属性が含まれています。値は、同じOIDを持つ属性を上書きします。ExtraNamesフィールドは、パース時には埋め込まれません。Namesを参照してください。
	ExtraNames []AttributeTypeAndValue
}

// FillFromRDNSequence は与えられた RDNSequence から n を埋めます。
// 複数エントリの RDN は平坦化され、すべてのエントリは関連する n フィールドに追加され、グルーピングは保持されません。
func (n *Name) FillFromRDNSequence(rdns *RDNSequence)

// ToRDNSequenceはnを単一のRDNSequenceに変換します。次の属性は複数値のRDNとしてエンコードされます：
// - 国
// - 組織
// - 組織単位
// - 地域
// - 県
// - 住所
// - 郵便番号
// 各ExtraNamesエントリは個別のRDNとしてエンコードされます。
func (n Name) ToRDNSequence() (ret RDNSequence)

// Stringはnの文字列形式を返します。ほぼ、RFC 2253の識別名の構文に従います。
func (n Name) String() string

// CertificateListは同名のASN.1構造を表します。RFC 5280、セクション5.1を参照してください。署名を検証するためにCertificate.CheckCRLSignatureを使用します。
//
// 廃止予定: 代わりにx509.RevocationListを使用するべきです。
type CertificateList struct {
	TBSCertList        TBSCertificateList
	SignatureAlgorithm AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

// HasExpiredは、certListがこの時点で更新されるべきかどうかを報告します。
func (certList *CertificateList) HasExpired(now time.Time) bool

// TBSCertificateListは、同じ名前のASN.1構造を表します。RFC 5280、セクション5.1を参照してください。
//
// 廃止予定：代わりにx509.RevocationListを使用するべきです。
type TBSCertificateList struct {
	Raw                 asn1.RawContent
	Version             int `asn1:"optional,default:0"`
	Signature           AlgorithmIdentifier
	Issuer              RDNSequence
	ThisUpdate          time.Time
	NextUpdate          time.Time            `asn1:"optional"`
	RevokedCertificates []RevokedCertificate `asn1:"optional"`
	Extensions          []Extension          `asn1:"tag:0,optional,explicit"`
}

// RevokedCertificateは同名のASN.1構造を表します。詳細はRFC 5280のセクション5.1を参照してください。
type RevokedCertificate struct {
	SerialNumber   *big.Int
	RevocationTime time.Time
	Extensions     []Extension `asn1:"optional"`
}
