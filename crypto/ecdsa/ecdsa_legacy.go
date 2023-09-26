// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdsa

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Sign signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key's curve order, the hash will be truncated to that length. It
// returns the signature as a pair of integers. Most applications should use
// SignASN1 instead of dealing directly with r, s.
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

// Verify verifies the signature in r, s of hash using the public key, pub. Its
// return value records whether the signature is valid. Most applications should
// use VerifyASN1 instead of dealing directly with r, s.
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
