// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	_ "crypto/sha256"
	_ "crypto/sha512"
)

// Self-signed certificate using ECDSA with SHA1 & secp256r1

// Self-signed certificate using ECDSA with SHA256 & secp256r1

// Self-signed certificate using ECDSA with SHA256 & secp384r1

// Self-signed certificate using ECDSA with SHA384 & secp521r1

// Self-signed certificate using DSA with SHA1

// These CSR was generated with OpenSSL:
//  openssl req -out CSR.csr -new -sha256 -nodes -keyout privateKey.key -config openssl.cnf
//
// With openssl.cnf containing the following sections:
//   [ v3_req ]
//   basicConstraints = CA:FALSE
//   keyUsage = nonRepudiation, digitalSignature, keyEncipherment
//   subjectAltName = email:gopher@golang.org,DNS:test.example.com
//   [ req_attributes ]
//   challengePassword = ignored challenge
//   unstructuredName  = ignored unstructured name
