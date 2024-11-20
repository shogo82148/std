// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ed25519

import (
	"github.com/shogo82148/std/crypto/internal/fips/edwards25519"
	"github.com/shogo82148/std/io"
)

type PrivateKey struct {
	seed   [seedSize]byte
	pub    [publicKeySize]byte
	s      edwards25519.Scalar
	prefix [sha512Size / 2]byte
}

func (priv *PrivateKey) Bytes() []byte

func (priv *PrivateKey) Seed() []byte

func (priv *PrivateKey) PublicKey() []byte

type PublicKey struct {
	a      edwards25519.Point
	aBytes [32]byte
}

func (pub *PublicKey) Bytes() []byte

// GenerateKey generates a new Ed25519 private key pair.
//
// In FIPS mode, rand is ignored. Otherwise, the output of this function is
// deterministic, and equivalent to reading 32 bytes from rand, and passing them
// to [NewKeyFromSeed].
func GenerateKey(rand io.Reader) (*PrivateKey, error)

func NewPrivateKeyFromSeed(seed []byte) (*PrivateKey, error)

func NewPrivateKey(priv []byte) (*PrivateKey, error)

func NewPublicKey(pub []byte) (*PublicKey, error)

func Sign(priv *PrivateKey, message []byte) []byte

func SignPH(priv *PrivateKey, message []byte, context string) ([]byte, error)

func SignCtx(priv *PrivateKey, message []byte, context string) ([]byte, error)

func Verify(pub *PublicKey, message, sig []byte) error

func VerifyPH(pub *PublicKey, message []byte, sig []byte, context string) error

func VerifyCtx(pub *PublicKey, message []byte, sig []byte, context string) error
