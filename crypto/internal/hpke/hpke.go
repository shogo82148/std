// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/crypto/ecdh"

	"golang.org/x/crypto/chacha20poly1305"
)

var SupportedKEMs = map[uint16]struct {
	curve   ecdh.Curve
	hash    crypto.Hash
	nSecret uint16
}{

	0x0020: {ecdh.X25519(), crypto.SHA256, 32},
}

type Sender struct {
	aead cipher.AEAD
	kem  *dhKEM

	sharedSecret []byte

	suiteID []byte

	key            []byte
	baseNonce      []byte
	exporterSecret []byte

	seqNum uint128
}

var SupportedAEADs = map[uint16]struct {
	keySize   int
	nonceSize int
	aead      func([]byte) (cipher.AEAD, error)
}{

	0x0001: {keySize: 16, nonceSize: 12, aead: aesGCMNew},
	0x0002: {keySize: 32, nonceSize: 12, aead: aesGCMNew},
	0x0003: {keySize: chacha20poly1305.KeySize, nonceSize: chacha20poly1305.NonceSize, aead: chacha20poly1305.New},
}

var SupportedKDFs = map[uint16]func() *hkdfKDF{

	0x0001: func() *hkdfKDF { return &hkdfKDF{crypto.SHA256} },
}

func SetupSender(kemID, kdfID, aeadID uint16, pub crypto.PublicKey, info []byte) ([]byte, *Sender, error)

func (s *Sender) Seal(aad, plaintext []byte) ([]byte, error)

func SuiteID(kemID, kdfID, aeadID uint16) []byte

func ParseHPKEPublicKey(kemID uint16, bytes []byte) (*ecdh.PublicKey, error)
