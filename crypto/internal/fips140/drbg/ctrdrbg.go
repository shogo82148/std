// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package drbg

import (
	"github.com/shogo82148/std/crypto/internal/fips140/aes"
)

// Counter is an SP 800-90A Rev. 1 CTR_DRBG instantiated with AES-256.
//
// Per Table 3, it has a security strength of 256 bits, a seed size of 384 bits,
// a counter length of 128 bits, a reseed interval of 2^48 requests, and a
// maximum request size of 2^19 bits (2^16 bytes, 64 KiB).
//
// We support a narrow range of parameters that fit the needs of our RNG:
// AES-256, no derivation function, no personalization string, no prediction
// resistance, and 384-bit additional input.
type Counter struct {
	// c is instantiated with K as the key and V as the counter.
	c aes.CTR

	reseedCounter uint64
}

const (
	SeedSize = keySize + aes.BlockSize
)

func NewCounter(entropy *[SeedSize]byte) *Counter

func (c *Counter) Reseed(entropy, additionalInput *[SeedSize]byte)

// Generate produces at most maxRequestSize bytes of random data in out.
func (c *Counter) Generate(out []byte, additionalInput *[SeedSize]byte) (reseedRequired bool)
