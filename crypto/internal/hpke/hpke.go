// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hpke implements Hybrid Public Key Encryption (HPKE) as defined in
// [RFC 9180].
//
// [RFC 9180]: https://www.rfc-editor.org/rfc/rfc9180.html
package hpke

// Sender is a sending HPKE context. It is instantiated with a specific KEM
// encapsulation key (i.e. the public key), and it is stateful, incrementing the
// nonce counter for each [Sender.Seal] call.
type Sender struct {
	*context
}

// Recipient is a receiving HPKE context. It is instantiated with a specific KEM
// decapsulation key (i.e. the secret key), and it is stateful, incrementing the
// nonce counter for each successful [Recipient.Open] call.
type Recipient struct {
	*context
}

// NewSender returns a sending HPKE context for the provided KEM encapsulation
// key (i.e. the public key), and using the ciphersuite defined by the
// combination of KEM, KDF, and AEAD.
//
// The info parameter is additional public information that must match between
// sender and recipient.
//
// The returned enc ciphertext can be used to instantiate a matching receiving
// HPKE context with the corresponding KEM decapsulation key.
func NewSender(pk PublicKey, kdf KDF, aead AEAD, info []byte) (enc []byte, s *Sender, err error)

// NewRecipient returns a receiving HPKE context for the provided KEM
// decapsulation key (i.e. the secret key), and using the ciphersuite defined by
// the combination of KEM, KDF, and AEAD.
//
// The enc parameter must have been produced by a matching sending HPKE context
// with the corresponding KEM encapsulation key. The info parameter is
// additional public information that must match between sender and recipient.
func NewRecipient(enc []byte, k PrivateKey, kdf KDF, aead AEAD, info []byte) (*Recipient, error)

// Seal encrypts the provided plaintext, optionally binding to the additional
// public data aad.
//
// Seal uses incrementing counters for each call, and Open on the receiving side
// must be called in the same order as Seal.
func (s *Sender) Seal(aad, plaintext []byte) ([]byte, error)

// Seal instantiates a single-use HPKE sending HPKE context like [NewSender],
// and then encrypts the provided plaintext like [Sender.Seal] (with no aad).
// Seal returns the concatenation of the encapsulated key and the ciphertext.
func Seal(pk PublicKey, kdf KDF, aead AEAD, info, plaintext []byte) ([]byte, error)

// Export produces a secret value derived from the shared key between sender and
// recipient. length must be at most 65,535.
func (s *Sender) Export(exporterContext string, length int) ([]byte, error)

// Open decrypts the provided ciphertext, optionally binding to the additional
// public data aad, or returns an error if decryption fails.
//
// Open uses incrementing counters for each successful call, and must be called
// in the same order as Seal on the sending side.
func (r *Recipient) Open(aad, ciphertext []byte) ([]byte, error)

// Open instantiates a single-use HPKE receiving HPKE context like [NewRecipient],
// and then decrypts the provided ciphertext like [Recipient.Open] (with no aad).
// ciphertext must be the concatenation of the encapsulated key and the actual ciphertext.
func Open(k PrivateKey, kdf KDF, aead AEAD, info, ciphertext []byte) ([]byte, error)

// Export produces a secret value derived from the shared key between sender and
// recipient. length must be at most 65,535.
func (r *Recipient) Export(exporterContext string, length int) ([]byte, error)
