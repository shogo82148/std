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

type KemID uint16

const DHKEM_X25519_HKDF_SHA256 = 0x0020

var SupportedKEMs = map[uint16]struct {
	curve   ecdh.Curve
	hash    crypto.Hash
	nSecret uint16
}{

	DHKEM_X25519_HKDF_SHA256: {ecdh.X25519(), crypto.SHA256, 32},
}

type Sender struct {
	*context
}

type Recipient struct {
	*context
}

type AEADID uint16

const (
	AEAD_AES_128_GCM      = 0x0001
	AEAD_AES_256_GCM      = 0x0002
	AEAD_ChaCha20Poly1305 = 0x0003
)

var SupportedAEADs = map[uint16]struct {
	keySize   int
	nonceSize int
	aead      func([]byte) (cipher.AEAD, error)
}{

	AEAD_AES_128_GCM:      {keySize: 16, nonceSize: 12, aead: aesGCMNew},
	AEAD_AES_256_GCM:      {keySize: 32, nonceSize: 12, aead: aesGCMNew},
	AEAD_ChaCha20Poly1305: {keySize: chacha20poly1305.KeySize, nonceSize: chacha20poly1305.NonceSize, aead: chacha20poly1305.New},
}

type KDFID uint16

const KDF_HKDF_SHA256 = 0x0001

var SupportedKDFs = map[uint16]func() *hkdfKDF{

	KDF_HKDF_SHA256: func() *hkdfKDF { return &hkdfKDF{crypto.SHA256} },
}

func SetupSender(kemID, kdfID, aeadID uint16, pub *ecdh.PublicKey, info []byte) ([]byte, *Sender, error)

func SetupRecipient(kemID, kdfID, aeadID uint16, priv *ecdh.PrivateKey, info, encPubEph []byte) (*Recipient, error)

func (s *Sender) Seal(aad, plaintext []byte) ([]byte, error)

func (r *Recipient) Open(aad, ciphertext []byte) ([]byte, error)

func ParseHPKEPublicKey(kemID uint16, bytes []byte) (*ecdh.PublicKey, error)

func ParseHPKEPrivateKey(kemID uint16, bytes []byte) (*ecdh.PrivateKey, error)
