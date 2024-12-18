// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CFB (Cipher Feedback) Mode.

package cipher

// NewCFBEncrypter returns a [Stream] which encrypts with cipher feedback mode,
// using the given [Block]. The iv must be the same length as the [Block]'s block
// size.
//
// Deprecated: CFB mode is not authenticated, which generally enables active
// attacks to manipulate and recover the plaintext. It is recommended that
// applications use [AEAD] modes instead. The standard library implementation of
// CFB is also unoptimized and not validated as part of the FIPS 140-3 module.
// If an unauthenticated [Stream] mode is required, use [NewCTR] instead.
func NewCFBEncrypter(block Block, iv []byte) Stream

// NewCFBDecrypter returns a [Stream] which decrypts with cipher feedback mode,
// using the given [Block]. The iv must be the same length as the [Block]'s block
// size.
//
// Deprecated: CFB mode is not authenticated, which generally enables active
// attacks to manipulate and recover the plaintext. It is recommended that
// applications use [AEAD] modes instead. The standard library implementation of
// CFB is also unoptimized and not validated as part of the FIPS 140-3 module.
// If an unauthenticated [Stream] mode is required, use [NewCTR] instead.
func NewCFBDecrypter(block Block, iv []byte) Stream
