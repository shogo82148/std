// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gcm

import (
	"github.com/shogo82148/std/crypto/internal/fips140/aes"
)

// SealWithRandomNonce encrypts plaintext to out, and writes a random nonce to
// nonce. nonce must be 12 bytes, and out must be 16 bytes longer than plaintext.
// out and plaintext may overlap exactly or not at all. additionalData and out
// must not overlap.
//
// This complies with FIPS 140-3 IG C.H Scenario 2.
//
// Note that this is NOT a [cipher.AEAD].Seal method.
func SealWithRandomNonce(g *GCM, nonce, out, plaintext, additionalData []byte)

// NewGCMWithCounterNonce returns a new AEAD that works like GCM, but enforces
// the construction of deterministic nonces. The nonce must be 96 bits, the
// first 32 bits must be an encoding of the module name, and the last 64 bits
// must be a counter.
//
// This complies with FIPS 140-3 IG C.H Scenario 3.
func NewGCMWithCounterNonce(cipher *aes.Block) (*GCMWithCounterNonce, error)

type GCMWithCounterNonce struct {
	g         GCM
	ready     bool
	fixedName uint32
	start     uint64
	next      uint64
}

func (g *GCMWithCounterNonce) NonceSize() int

func (g *GCMWithCounterNonce) Overhead() int

func (g *GCMWithCounterNonce) Seal(dst, nonce, plaintext, data []byte) []byte

func (g *GCMWithCounterNonce) Open(dst, nonce, ciphertext, data []byte) ([]byte, error)

// NewGCMForTLS12 returns a new AEAD that works like GCM, but enforces the
// construction of nonces as specified in RFC 5288, Section 3 and RFC 9325,
// Section 7.2.1.
//
// This complies with FIPS 140-3 IG C.H Scenario 1.a.
func NewGCMForTLS12(cipher *aes.Block) (*GCMForTLS12, error)

type GCMForTLS12 struct {
	g    GCM
	next uint64
}

func (g *GCMForTLS12) NonceSize() int

func (g *GCMForTLS12) Overhead() int

func (g *GCMForTLS12) Seal(dst, nonce, plaintext, data []byte) []byte

func (g *GCMForTLS12) Open(dst, nonce, ciphertext, data []byte) ([]byte, error)

// NewGCMForTLS13 returns a new AEAD that works like GCM, but enforces the
// construction of nonces as specified in RFC 8446, Section 5.3.
func NewGCMForTLS13(cipher *aes.Block) (*GCMForTLS13, error)

type GCMForTLS13 struct {
	g     GCM
	ready bool
	mask  uint64
	next  uint64
}

func (g *GCMForTLS13) NonceSize() int

func (g *GCMForTLS13) Overhead() int

func (g *GCMForTLS13) Seal(dst, nonce, plaintext, data []byte) []byte

func (g *GCMForTLS13) Open(dst, nonce, ciphertext, data []byte) ([]byte, error)

// NewGCMForSSH returns a new AEAD that works like GCM, but enforces the
// construction of nonces as specified in RFC 5647.
//
// This complies with FIPS 140-3 IG C.H Scenario 1.d.
func NewGCMForSSH(cipher *aes.Block) (*GCMForSSH, error)

type GCMForSSH struct {
	g     GCM
	ready bool
	start uint64
	next  uint64
}

func (g *GCMForSSH) NonceSize() int

func (g *GCMForSSH) Overhead() int

func (g *GCMForSSH) Seal(dst, nonce, plaintext, data []byte) []byte

func (g *GCMForSSH) Open(dst, nonce, ciphertext, data []byte) ([]byte, error)
