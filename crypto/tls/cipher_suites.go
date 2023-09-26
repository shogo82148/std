// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// a keyAgreement implements the client and server side of a TLS key agreement
// protocol by generating and processing key exchange messages.

// A cipherSuite is a specific combination of key agreement, cipher and MAC
// function. All cipher suites currently assume RSA key agreement.

// ssl30MAC implements the SSLv3 MAC function, as defined in
// www.mozilla.org/projects/security/pki/nss/ssl/draft302.txt section 5.2.3.1

// tls10MAC implements the TLS 1.0 MAC function. RFC 2246, section 6.2.3.

// A list of the possible cipher suite ids. Taken from
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xml
const (
	TLS_RSA_WITH_RC4_128_SHA            uint16 = 0x0005
	TLS_RSA_WITH_3DES_EDE_CBC_SHA       uint16 = 0x000a
	TLS_RSA_WITH_AES_128_CBC_SHA        uint16 = 0x002f
	TLS_RSA_WITH_AES_256_CBC_SHA        uint16 = 0x0035
	TLS_ECDHE_RSA_WITH_RC4_128_SHA      uint16 = 0xc011
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA uint16 = 0xc012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA  uint16 = 0xc013
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA  uint16 = 0xc014
)
