// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/crypto/rsa"
)

// pkcs1PrivateKey is a structure which mirrors the PKCS #1 ASN.1 for an RSA private key.

// pkcs1PublicKey reflects the ASN.1 structure of a PKCS #1 public key.

// ParsePKCS1PrivateKey parses an RSA private key in PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)

// MarshalPKCS1PrivateKey converts an RSA private key to PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
// For a more flexible key format which is not RSA specific, use
// MarshalPKCS8PrivateKey.
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte

// ParsePKCS1PublicKey parses an RSA public key in PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".
func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)

// MarshalPKCS1PublicKey converts an RSA public key to PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".
func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
