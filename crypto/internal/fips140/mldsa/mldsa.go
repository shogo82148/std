// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mldsa

const (
	PrivateKeySize = 32

	PublicKeySize44 = 32 + 4*n*10/8
	PublicKeySize65 = 32 + 6*n*10/8
	PublicKeySize87 = 32 + 8*n*10/8

	SignatureSize44 = 128/4 + 4*n*(17+1)/8 + 80 + 4
	SignatureSize65 = 192/4 + 5*n*(19+1)/8 + 55 + 6
	SignatureSize87 = 256/4 + 7*n*(19+1)/8 + 75 + 8
)

type PrivateKey struct {
	seed [32]byte
	pub  PublicKey
	s1   [maxL]nttElement
	s2   [maxK]nttElement
	t0   [maxK]nttElement
	k    [32]byte
}

func (priv *PrivateKey) Equal(x *PrivateKey) bool

func (priv *PrivateKey) Bytes() []byte

func (priv *PrivateKey) PublicKey() *PublicKey

type PublicKey struct {
	raw [maxPubKeySize]byte
	p   parameters
	a   [maxK * maxL]nttElement
	t1  [maxK]nttElement
	tr  [64]byte
}

func (pub *PublicKey) Equal(x *PublicKey) bool

func (pub *PublicKey) Bytes() []byte

func (pub *PublicKey) Parameters() string

func GenerateKey44() *PrivateKey

func GenerateKey65() *PrivateKey

func GenerateKey87() *PrivateKey

func NewPrivateKey44(seed []byte) (*PrivateKey, error)

func NewPrivateKey65(seed []byte) (*PrivateKey, error)

func NewPrivateKey87(seed []byte) (*PrivateKey, error)

func NewPublicKey44(pk []byte) (*PublicKey, error)

func NewPublicKey65(pk []byte) (*PublicKey, error)

func NewPublicKey87(pk []byte) (*PublicKey, error)

func Sign(priv *PrivateKey, msg []byte, context string) ([]byte, error)

func SignDeterministic(priv *PrivateKey, msg []byte, context string) ([]byte, error)

func TestingOnlySignWithRandom(priv *PrivateKey, msg []byte, context string, random []byte) ([]byte, error)

func SignExternalMu(priv *PrivateKey, μ []byte) ([]byte, error)

func SignExternalMuDeterministic(priv *PrivateKey, μ []byte) ([]byte, error)

func TestingOnlySignExternalMuWithRandom(priv *PrivateKey, μ []byte, random []byte) ([]byte, error)

func Verify(pub *PublicKey, msg, sig []byte, context string) error

func VerifyExternalMu(pub *PublicKey, μ []byte, sig []byte) error
