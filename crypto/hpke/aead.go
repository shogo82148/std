// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// The AEAD is one of the three components of an HPKE ciphersuite, implementing
// symmetric encryption.
type AEAD interface {
	ID() uint16
	keySize() int
	nonceSize() int
	aead(key []byte) (cipher.AEAD, error)
}

// NewAEAD returns the AEAD implementation for the given AEAD ID.
//
// Applications are encouraged to use specific implementations like [AES128GCM]
// or [ChaCha20Poly1305] instead, unless runtime agility is required.
func NewAEAD(id uint16) (AEAD, error)

// AES128GCM returns an AES-128-GCM AEAD implementation.
func AES128GCM() AEAD

// AES256GCM returns an AES-256-GCM AEAD implementation.
func AES256GCM() AEAD

// ChaCha20Poly1305 returns a ChaCha20Poly1305 AEAD implementation.
func ChaCha20Poly1305() AEAD

// ExportOnly returns a placeholder AEAD implementation that cannot encrypt or
// decrypt, but only export secrets with [Sender.Export] or [Recipient.Export].
//
// When this is used, [Sender.Seal] and [Recipient.Open] return errors.
func ExportOnly() AEAD
