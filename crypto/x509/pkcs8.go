// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// pkcs8 reflects an ASN.1, PKCS#8 PrivateKey. See
// ftp://ftp.rsasecurity.com/pub/pkcs/pkcs-8/pkcs-8v1_2.asn
// and RFC 5208.

// ParsePKCS8PrivateKey parses an unencrypted, PKCS#8 private key.
// See RFC 5208.
func ParsePKCS8PrivateKey(der []byte) (key interface{}, err error)

// MarshalPKCS8PrivateKey converts a private key to PKCS#8 encoded form.
// The following key types are supported: *rsa.PrivateKey, *ecdsa.PublicKey.
// Unsupported key types result in an error.
//
// See RFC 5208.
func MarshalPKCS8PrivateKey(key interface{}) ([]byte, error)
