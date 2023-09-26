// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

// encoder represents an ASN.1 element that is waiting to be marshaled.

// Marshal returns the ASN.1 encoding of val.
//
// In addition to the struct tags recognized by Unmarshal, the following can be
// used:
//
//	ia5:         causes strings to be marshaled as ASN.1, IA5String values
//	omitempty:   causes empty slices to be skipped
//	printable:   causes strings to be marshaled as ASN.1, PrintableString values
//	utf8:        causes strings to be marshaled as ASN.1, UTF8String values
//	utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
//	generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
func Marshal(val any) ([]byte, error)

// MarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func MarshalWithParams(val any, params string) ([]byte, error)
