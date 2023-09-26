// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

// ASN.1 tags represent the type of the following object.
const (
	TagBoolean         = 1
	TagInteger         = 2
	TagBitString       = 3
	TagOctetString     = 4
	TagNull            = 5
	TagOID             = 6
	TagEnum            = 10
	TagUTF8String      = 12
	TagSequence        = 16
	TagSet             = 17
	TagNumericString   = 18
	TagPrintableString = 19
	TagT61String       = 20
	TagIA5String       = 22
	TagUTCTime         = 23
	TagGeneralizedTime = 24
	TagGeneralString   = 27
	TagBMPString       = 30
)

// ASN.1 class types represent the namespace of the tag.
const (
	ClassUniversal       = 0
	ClassApplication     = 1
	ClassContextSpecific = 2
	ClassPrivate         = 3
)

// fieldParameters is the parsed representation of tag string from a structure field.
