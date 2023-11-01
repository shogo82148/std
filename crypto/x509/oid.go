// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/encoding/asn1"
)

// An OID represents an ASN.1 OBJECT IDENTIFIER.
type OID struct {
	der []byte
}

// OIDFromInts creates a new OID using ints, each integer is a separate component.
func OIDFromInts(oid []uint64) (OID, error)

// Equal returns true when oid and other represents the same Object Identifier.
func (oid OID) Equal(other OID) bool

// EqualASN1OID returns whether an OID equals an asn1.ObjectIdentifier. If
// asn1.ObjectIdentifier cannot represent the OID specified by oid, because
// a component of OID requires more than 31 bits, it returns false.
func (oid OID) EqualASN1OID(other asn1.ObjectIdentifier) bool

// Strings returns the string representation of the Object Identifier.
func (oid OID) String() string
