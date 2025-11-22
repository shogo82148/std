// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/ecdh"
)

// MLKEM768X25519 returns a KEM implementing MLKEM768-X25519 (a.k.a. X-Wing)
// from draft-ietf-hpke-pq.
func MLKEM768X25519() KEM

// MLKEM768P256 returns a KEM implementing MLKEM768-P256 from draft-ietf-hpke-pq.
func MLKEM768P256() KEM

// MLKEM1024P384 returns a KEM implementing MLKEM1024-P384 from draft-ietf-hpke-pq.
func MLKEM1024P384() KEM

// NewHybridPublicKey returns a PublicKey implementing one of
//
//   - MLKEM768-X25519 (a.k.a. X-Wing)
//   - MLKEM768-P256
//   - MLKEM1024-P384
//
// from draft-ietf-hpke-pq, depending on the underlying curve of t
// ([ecdh.X25519], [ecdh.P256], or [ecdh.P384]) and the type of pq (either
// *[mlkem.EncapsulationKey768] or *[mlkem.EncapsulationKey1024]).
//
// This function is meant for applications that already have instantiated
// crypto/ecdh and crypto/mlkem public keys. Otherwise, applications should use
// the [KEM.NewPublicKey] method of e.g. [MLKEM768X25519].
func NewHybridPublicKey(pq crypto.Encapsulator, t *ecdh.PublicKey) (PublicKey, error)

// NewHybridPrivateKey returns a PrivateKey implementing
//
//   - MLKEM768-X25519 (a.k.a. X-Wing)
//   - MLKEM768-P256
//   - MLKEM1024-P384
//
// from draft-ietf-hpke-pq, depending on the underlying curve of t
// ([ecdh.X25519], [ecdh.P256], or [ecdh.P384]) and the type of pq.Encapsulator()
// (either *[mlkem.EncapsulationKey768] or *[mlkem.EncapsulationKey1024]).
//
// This function is meant for applications that already have instantiated
// crypto/ecdh and crypto/mlkem private keys, or another implementation of a
// [ecdh.KeyExchanger] and [crypto.Decapsulator] (e.g. a hardware key).
// Otherwise, applications should use the [KEM.NewPrivateKey] method of e.g.
// [MLKEM768X25519].
func NewHybridPrivateKey(pq crypto.Decapsulator, t ecdh.KeyExchanger) (PrivateKey, error)

// MLKEM768 returns a KEM implementing ML-KEM-768 from draft-ietf-hpke-pq.
func MLKEM768() KEM

// MLKEM1024 returns a KEM implementing ML-KEM-1024 from draft-ietf-hpke-pq.
func MLKEM1024() KEM

// NewMLKEMPublicKey returns a KEMPublicKey implementing
//
//   - ML-KEM-768
//   - ML-KEM-1024
//
// from draft-ietf-hpke-pq, depending on the type of pub
// (*[mlkem.EncapsulationKey768] or *[mlkem.EncapsulationKey1024]).
//
// This function is meant for applications that already have an instantiated
// crypto/mlkem public key. Otherwise, applications should use the
// [KEM.NewPublicKey] method of e.g. [MLKEM768].
func NewMLKEMPublicKey(pub crypto.Encapsulator) (PublicKey, error)

// NewMLKEMPrivateKey returns a KEMPrivateKey implementing
//
//   - ML-KEM-768
//   - ML-KEM-1024
//
// from draft-ietf-hpke-pq, depending on the type of priv.Encapsulator()
// (either *[mlkem.EncapsulationKey768] or *[mlkem.EncapsulationKey1024]).
//
// This function is meant for applications that already have an instantiated
// crypto/mlkem private key. Otherwise, applications should use the
// [KEM.NewPrivateKey] method of e.g. [MLKEM768].
func NewMLKEMPrivateKey(priv crypto.Decapsulator) (PrivateKey, error)
