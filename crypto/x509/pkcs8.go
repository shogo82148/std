// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// pkcs8 reflects an ASN.1, PKCS#8 PrivateKey. See
// ftp://ftp.rsasecurity.com/pub/pkcs/pkcs-8/pkcs-8v1_2.asn.

// ParsePKCS8PrivateKey parses an unencrypted, PKCS#8 private key. See
// http://www.rsa.com/rsalabs/node.asp?id=2130
func ParsePKCS8PrivateKey(der []byte) (key interface{}, err error)
