// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	_ "crypto/sha256"
	_ "crypto/sha512"
)

// pemEd25519Key is the example from RFC 8410, Secrion 4.

// Self-signed certificate using ECDSA with SHA1 & secp256r1

// Self-signed certificate using ECDSA with SHA256 & secp256r1

// Self-signed certificate using ECDSA with SHA256 & secp384r1

// Self-signed certificate using ECDSA with SHA384 & secp521r1

// Self-signed certificate using DSA with SHA1

// openssl req -newkey rsa:2048 -keyout test.key -sha256 -sigopt \
// rsa_padding_mode:pss -sigopt rsa_pss_saltlen:32 -sigopt rsa_mgf1_md:sha256 \
// -x509 -days 3650 -nodes -subj '/C=US/ST=CA/L=SF/O=Test/CN=Test' -out \
// test.pem

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

// certMissingRSANULL contains an RSA public key where the AlgorithmIdentifier
// parameters are omitted rather than being an ASN.1 NULL.

// certMultipleRDN contains a RelativeDistinguishedName with two elements (the
// common name and serial number). This particular certificate was the first
// such certificate in the “Pilot” Certificate Transparency log.

// multipleURLsInCRLDPPEM contains two URLs in a single CRL DistributionPoint
// structure. It is taken from https://crt.sh/?id=12721534.

// mismatchingSigAlgIDPEM contains a certificate where the Certificate
// signatureAlgorithm and the TBSCertificate signature contain
// mismatching OIDs

// mismatchingSigAlgParamPEM contains a certificate where the Certificate
// signatureAlgorithm and the TBSCertificate signature contain
// mismatching parameters
