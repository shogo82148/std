// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto/ecdh"
)

// A KEM is a Key Encapsulation Mechanism, one of the three components of an
// HPKE ciphersuite.
type KEM interface {
	ID() uint16

	GenerateKey() (PrivateKey, error)

	NewPublicKey([]byte) (PublicKey, error)

	NewPrivateKey([]byte) (PrivateKey, error)

	DeriveKeyPair(ikm []byte) (PrivateKey, error)

	encSize() int
}

// NewKEM returns the KEM implementation for the given KEM ID.
//
// Applications are encouraged to use specific implementations like [DHKEM] or
// [MLKEM768X25519] instead, unless runtime agility is required.
func NewKEM(id uint16) (KEM, error)

// A PublicKey is an instantiation of a KEM (one of the three components of an
// HPKE ciphersuite) with an encapsulation key (i.e. the public key).
//
// A PublicKey is usually obtained from a method of the corresponding [KEM] or
// [PrivateKey], such as [KEM.NewPublicKey] or [PrivateKey.PublicKey].
type PublicKey interface {
	KEM() KEM

	Bytes() []byte

	encap() (sharedSecret, enc []byte, err error)
}

// A PrivateKey is an instantiation of a KEM (one of the three components of
// an HPKE ciphersuite) with a decapsulation key (i.e. the secret key).
//
// A PrivateKey is usually obtained from a method of the corresponding [KEM],
// such as [KEM.GenerateKey] or [KEM.NewPrivateKey].
type PrivateKey interface {
	KEM() KEM

	Bytes() ([]byte, error)

	PublicKey() PublicKey

	decap(enc []byte) (sharedSecret []byte, err error)
}

// DHKEM returns a KEM implementing one of
//
//   - DHKEM(P-256, HKDF-SHA256)
//   - DHKEM(P-384, HKDF-SHA384)
//   - DHKEM(P-521, HKDF-SHA512)
//   - DHKEM(X25519, HKDF-SHA256)
//
// depending on curve.
func DHKEM(curve ecdh.Curve) KEM

// NewDHKEMPublicKey returns a PublicKey implementing
//
//   - DHKEM(P-256, HKDF-SHA256)
//   - DHKEM(P-384, HKDF-SHA384)
//   - DHKEM(P-521, HKDF-SHA512)
//   - DHKEM(X25519, HKDF-SHA256)
//
// depending on the underlying curve of pub ([ecdh.X25519], [ecdh.P256],
// [ecdh.P384], or [ecdh.P521]).
//
// This function is meant for applications that already have an instantiated
// crypto/ecdh public key. Otherwise, applications should use the
// [KEM.NewPublicKey] method of [DHKEM].
func NewDHKEMPublicKey(pub *ecdh.PublicKey) (PublicKey, error)

// NewDHKEMPrivateKey returns a PrivateKey implementing
//
//   - DHKEM(P-256, HKDF-SHA256)
//   - DHKEM(P-384, HKDF-SHA384)
//   - DHKEM(P-521, HKDF-SHA512)
//   - DHKEM(X25519, HKDF-SHA256)
//
// depending on the underlying curve of priv ([ecdh.X25519], [ecdh.P256],
// [ecdh.P384], or [ecdh.P521]).
//
// This function is meant for applications that already have an instantiated
// crypto/ecdh private key, or another implementation of a [ecdh.KeyExchanger]
// (e.g. a hardware key). Otherwise, applications should use the
// [KEM.NewPrivateKey] method of [DHKEM].
func NewDHKEMPrivateKey(priv ecdh.KeyExchanger) (PrivateKey, error)
