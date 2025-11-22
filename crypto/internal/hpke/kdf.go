// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

// The KDF is one of the three components of an HPKE ciphersuite, implementing
// key derivation.
type KDF interface {
	ID() uint16
	oneStage() bool
	size() int
	labeledDerive(suiteID, inputKey []byte, label string, context []byte, length uint16) ([]byte, error)
	labeledExtract(suiteID, salt []byte, label string, inputKey []byte) ([]byte, error)
	labeledExpand(suiteID, randomKey []byte, label string, info []byte, length uint16) ([]byte, error)
}

// NewKDF returns the KDF implementation for the given KDF ID.
//
// Applications are encouraged to use specific implementations like [HKDFSHA256]
// instead, unless runtime agility is required.
func NewKDF(id uint16) (KDF, error)

// HKDFSHA256 returns an HKDF-SHA256 KDF implementation.
func HKDFSHA256() KDF

// HKDFSHA384 returns an HKDF-SHA384 KDF implementation.
func HKDFSHA384() KDF

// HKDFSHA512 returns an HKDF-SHA512 KDF implementation.
func HKDFSHA512() KDF

// SHAKE128 returns a SHAKE128 KDF implementation.
func SHAKE128() KDF

// SHAKE256 returns a SHAKE256 KDF implementation.
func SHAKE256() KDF
