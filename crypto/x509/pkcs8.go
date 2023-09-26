// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// pkcs8 reflects an ASN.1, PKCS #8 PrivateKey. See
// ftp://ftp.rsasecurity.com/pub/pkcs/pkcs-8/pkcs-8v1_2.asn
// and RFC 5208.

// ParsePKCS8PrivateKey parses an unencrypted private key in PKCS #8, ASN.1 DER form.
//
// It returns a *rsa.PrivateKey, a *ecdsa.PrivateKey, a ed25519.PrivateKey (not
// a pointer), or a *ecdh.PrivateKey (for X25519). More types might be supported
// in the future.
//
// This kind of key is commonly encoded in PEM blocks of type "PRIVATE KEY".
func ParsePKCS8PrivateKey(der []byte) (key any, err error)

// MarshalPKCS8PrivateKey converts a private key to PKCS #8, ASN.1 DER form.
//
// The following key types are currently supported: *rsa.PrivateKey,
// *ecdsa.PrivateKey, ed25519.PrivateKey (not a pointer), and *ecdh.PrivateKey.
// Unsupported key types result in an error.
//
// This kind of key is commonly encoded in PEM blocks of type "PRIVATE KEY".
func MarshalPKCS8PrivateKey(key any) ([]byte, error)
