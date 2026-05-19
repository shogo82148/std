// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/io"
)

// GenerateKey generates a new RSA key pair of the given bit size.
// bits must be at least 32.
//
// It follows the process described at c2sp.org/det-keygen, which is compliant
// with FIPS 186-5, Appendix A.1, IFC Key Pair Generation and FIPS 186-5,
// Appendix A.1.3, Generation of Random Primes that are Probably Prime.
// The prime candidates are drawn from rand, which in production will be the
// global DRBG, while in tests can be an HMAC_DRBG as specified in
// c2sp.org/det-keygen, to allow using its tests vectors.
func GenerateKey(rand io.Reader, bits int) (*PrivateKey, error)
