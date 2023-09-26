// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

type DecryptPKCS1v15Test struct {
	in, out string
}

// These test vectors were generated with `openssl rsautl -pkcs -encrypt`

// These test vectors were generated with `openssl rsautl -pkcs -encrypt`

// These vectors have been tested with
//   `openssl rsautl -verify -inkey pk -in signature | hexdump -C`
