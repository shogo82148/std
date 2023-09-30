// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ecdh implements Elliptic Curve Diffie-Hellman over
// NIST curves and Curve25519.
package ecdh

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

type Curve interface {
	GenerateKey(rand io.Reader) (*PrivateKey, error)

	NewPrivateKey(key []byte) (*PrivateKey, error)

	NewPublicKey(key []byte) (*PublicKey, error)

	ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error)

	privateKeyToPublicKey(*PrivateKey) *PublicKey
}

// PublicKey is an ECDH public key, usually a peer's ECDH share sent over the wire.
//
// These keys can be parsed with [crypto/x509.ParsePKIXPublicKey] and encoded
// with [crypto/x509.MarshalPKIXPublicKey]. For NIST curves, they then need to
// be converted with [crypto/ecdsa.PublicKey.ECDH] after parsing.
type PublicKey struct {
	curve     Curve
	publicKey []byte
	boring    *boring.PublicKeyECDH
}

// Bytes returns a copy of the encoding of the public key.
func (k *PublicKey) Bytes() []byte

// Equal returns whether x represents the same public key as k.
//
// Note that there can be equivalent public keys with different encodings which
// would return false from this check but behave the same way as inputs to ECDH.
//
// This check is performed in constant time as long as the key types and their
// curve match.
func (k *PublicKey) Equal(x crypto.PublicKey) bool

func (k *PublicKey) Curve() Curve

// PrivateKey is an ECDH private key, usually kept secret.
//
// These keys can be parsed with [crypto/x509.ParsePKCS8PrivateKey] and encoded
// with [crypto/x509.MarshalPKCS8PrivateKey]. For NIST curves, they then need to
// be converted with [crypto/ecdsa.PrivateKey.ECDH] after parsing.
type PrivateKey struct {
	curve      Curve
	privateKey []byte
	boring     *boring.PrivateKeyECDH

	publicKey     *PublicKey
	publicKeyOnce sync.Once
}

// ECDH performs an ECDH exchange and returns the shared secret. The PrivateKey
// and PublicKey must use the same curve.
//
// For NIST curves, this performs ECDH as specified in SEC 1, Version 2.0,
// Section 3.3.1, and returns the x-coordinate encoded according to SEC 1,
// Version 2.0, Section 2.3.5. The result is never the point at infinity.
//
// For X25519, this performs ECDH as specified in RFC 7748, Section 6.1. If
// the result is the all-zero value, ECDH returns an error.
func (k *PrivateKey) ECDH(remote *PublicKey) ([]byte, error)

// Bytes returns a copy of the encoding of the private key.
func (k *PrivateKey) Bytes() []byte

// Equal returns whether x represents the same private key as k.
//
// Note that there can be equivalent private keys with different encodings which
// would return false from this check but behave the same way as inputs to ECDH.
//
// This check is performed in constant time as long as the key types and their
// curve match.
func (k *PrivateKey) Equal(x crypto.PrivateKey) bool

func (k *PrivateKey) Curve() Curve

func (k *PrivateKey) PublicKey() *PublicKey

// Public implements the implicit interface of all standard library private
// keys. See the docs of crypto.PrivateKey.
func (k *PrivateKey) Public() crypto.PublicKey
