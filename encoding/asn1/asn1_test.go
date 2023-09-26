// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

import (
	"math/big"
	"time"
)

type TestObjectIdentifierStruct struct {
	OID ObjectIdentifier
}

type TestContextSpecificTags struct {
	A int `asn1:"tag:1"`
}

type TestContextSpecificTags2 struct {
	A int `asn1:"explicit,tag:1"`
	B int
}

type TestContextSpecificTags3 struct {
	S string `asn1:"tag:1,utf8"`
}

type TestElementsAfterString struct {
	S    string
	A, B int
}

type TestBigInt struct {
	X *big.Int
}

type TestSet struct {
	Ints []int `asn1:"set"`
}

type Certificate struct {
	TBSCertificate     TBSCertificate
	SignatureAlgorithm AlgorithmIdentifier
	SignatureValue     BitString
}

type TBSCertificate struct {
	Version            int `asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber       RawValue
	SignatureAlgorithm AlgorithmIdentifier
	Issuer             RDNSequence
	Validity           Validity
	Subject            RDNSequence
	PublicKey          PublicKeyInfo
}

type AlgorithmIdentifier struct {
	Algorithm ObjectIdentifier
}

type RDNSequence []RelativeDistinguishedNameSET

type RelativeDistinguishedNameSET []AttributeTypeAndValue

type AttributeTypeAndValue struct {
	Type  ObjectIdentifier
	Value interface{}
}

type Validity struct {
	NotBefore, NotAfter time.Time
}

type PublicKeyInfo struct {
	Algorithm AlgorithmIdentifier
	PublicKey BitString
}
