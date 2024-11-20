// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gcm

import (
	"github.com/shogo82148/std/crypto/internal/fips/aes"
)

// GCM represents a Galois Counter Mode with a specific key.
type GCM struct {
	cipher    aes.Block
	nonceSize int
	tagSize   int
	gcmPlatformData
}

func New(cipher *aes.Block, nonceSize, tagSize int) (*GCM, error)

func (g *GCM) NonceSize() int

func (g *GCM) Overhead() int

func (g *GCM) Seal(dst, nonce, plaintext, data []byte) []byte

func (g *GCM) Open(dst, nonce, ciphertext, data []byte) ([]byte, error)
