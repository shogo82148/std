// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// KeyBytes returns the encoding of a key.
//
// It is usually a method expression like (*ecdh.PrivateKey).Bytes.
type KeyBytes[K any] interface {
	func(K) []byte | func(K) ([]byte, error)
}

// TestPrivateKey performs a set of tests on a private key type, checking the
// Equal and Public methods documented in crypto.PrivateKey, and its encoding
// like [TestKeyEncoding].
//
// k1 and k2 must be different keys.
func TestPrivateKey[K privateKey, B KeyBytes[K]](t *testing.T, k1, k2 K, newKey func([]byte) (K, error), keyBytes B)

// TestPublicKey performs a set of tests on a public key type, checking the
// Equal method documented in crypto.PublicKey, and its encoding like
// [TestKeyEncoding].
//
// k1 and k2 must be different keys.
func TestPublicKey[K publicKey, B KeyBytes[K]](t *testing.T, k1, k2 K, newKey func([]byte) (K, error), keyBytes B)

// TestKeyEncoding performs a set of tests on the fixed-length encoding of a key
// type, checking that it round-trips, that inputs of the wrong length are
// rejected, and that keys don't share memory with the input of newKey or with
// the output of keyBytes.
//
// newKey parses the encoding and keyBytes produces it, like NewPrivateKey and
// PrivateKey.Bytes in crypto/ecdh. k should come from a function other than
// newKey, such as GenerateKey, so that both are tested.
func TestKeyEncoding[K any, B KeyBytes[K]](t *testing.T, k K, newKey func([]byte) (K, error), keyBytes B)
