// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pkix contains shared, low level structures used for ASN.1 parsing
// and serialization of X.509 certificates, CRL and OCSP.
package pkix

import (
	"github.com/shogo82148/std/encoding/asn1"
	"github.com/shogo82148/std/math/big"
	"github.com/shogo82148/std/time"
)

// AlgorithmIdentifier represents the ASN.1 structure of the same name. See RFC
// 5280, section 4.1.1.2.
type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}

type RDNSequence []RelativeDistinguishedNameSET

// String returns a string representation of the sequence r,
// roughly following the RFC 2253 Distinguished Names syntax.
func (r RDNSequence) String() string

type RelativeDistinguishedNameSET []AttributeTypeAndValue

// AttributeTypeAndValue mirrors the ASN.1 structure of the same name in
// RFC 5280, Section 4.1.2.4.
type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value interface{}
}

// AttributeTypeAndValueSET represents a set of ASN.1 sequences of
// AttributeTypeAndValue sequences from RFC 2986 (PKCS #10).
type AttributeTypeAndValueSET struct {
	Type  asn1.ObjectIdentifier
	Value [][]AttributeTypeAndValue `asn1:"set"`
}

// Extension represents the ASN.1 structure of the same name. See RFC
// 5280, section 4.2.
type Extension struct {
	Id       asn1.ObjectIdentifier
	Critical bool `asn1:"optional"`
	Value    []byte
}

// Name represents an X.509 distinguished name. This only includes the common
// elements of a DN. Note that Name is only an approximation of the X.509
// structure. If an accurate representation is needed, asn1.Unmarshal the raw
// subject or issuer as an RDNSequence.
type Name struct {
	Country, Organization, OrganizationalUnit []string
	Locality, Province                        []string
	StreetAddress, PostalCode                 []string
	SerialNumber, CommonName                  string

	Names []AttributeTypeAndValue

	ExtraNames []AttributeTypeAndValue
}

// FillFromRDNSequence populates n from the provided RDNSequence.
// Multi-entry RDNs are flattened, all entries are added to the
// relevant n fields, and the grouping is not preserved.
func (n *Name) FillFromRDNSequence(rdns *RDNSequence)

// ToRDNSequence converts n into a single RDNSequence. The following
// attributes are encoded as multi-value RDNs:
//
//   - Country
//   - Organization
//   - OrganizationalUnit
//   - Locality
//   - Province
//   - StreetAddress
//   - PostalCode
//
// Each ExtraNames entry is encoded as an individual RDN.
func (n Name) ToRDNSequence() (ret RDNSequence)

// String returns the string form of n, roughly following
// the RFC 2253 Distinguished Names syntax.
func (n Name) String() string

// CertificateList represents the ASN.1 structure of the same name. See RFC
// 5280, section 5.1. Use Certificate.CheckCRLSignature to verify the
// signature.
type CertificateList struct {
	TBSCertList        TBSCertificateList
	SignatureAlgorithm AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

// HasExpired reports whether certList should have been updated by now.
func (certList *CertificateList) HasExpired(now time.Time) bool

// TBSCertificateList represents the ASN.1 structure of the same name. See RFC
// 5280, section 5.1.
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

// RevokedCertificate represents the ASN.1 structure of the same name. See RFC
// 5280, section 5.1.
type RevokedCertificate struct {
	SerialNumber   *big.Int
	RevocationTime time.Time
	Extensions     []Extension `asn1:"optional"`
}
