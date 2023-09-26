// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smtp

// toServerEmptyAuth is an implementation of Auth that only implements
// the Start method, and returns "FOOAUTH", nil, nil. Notably, it returns
// zero bytes for "toServer" so we can test that we don't send spaces at
// the end of the line. See TestClientAuthTrimSpace.

// localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:
// go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
// 		--ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h

// localhostKey is the private key for localhostCert.
