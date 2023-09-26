// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// Generated using:
//
//	openssl genrsa 1024 | openssl pkcs8 -topk8 -nocrypt

// Generated using:
//
//	openssl ecparam -genkey -name secp224r1 | openssl pkcs8 -topk8 -nocrypt

// Generated using:
//
//	openssl ecparam -genkey -name secp256r1 | openssl pkcs8 -topk8 -nocrypt

// Generated using:
//
//	openssl ecparam -genkey -name secp384r1 | openssl pkcs8 -topk8 -nocrypt

// Generated using:
//
//	openssl ecparam -genkey -name secp521r1 | openssl pkcs8 -topk8 -nocrypt
//
// Note that OpenSSL will truncate the private key if it can (i.e. it emits it
// like an integer, even though it's an OCTET STRING field). Thus if you
// regenerate this you may, randomly, find that it's a byte shorter than
// expected and the Go test will fail to recreate it exactly.

// From RFC 8410, Section 7.

// Generated using:
//
//	openssl genpkey -algorithm x25519
