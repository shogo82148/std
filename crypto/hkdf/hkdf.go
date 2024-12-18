// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hkdf implements the HMAC-based Extract-and-Expand Key Derivation
// Function (HKDF) as defined in RFC 5869.
//
// HKDF is a cryptographic key derivation function (KDF) with the goal of
// expanding limited input keying material into one or more cryptographically
// strong secret keys.
package hkdf

import (
	"github.com/shogo82148/std/hash"
)

// Extract generates a pseudorandom key for use with [Expand] from an input
// secret and an optional independent salt.
//
// Only use this function if you need to reuse the extracted key with multiple
// Expand invocations and different context values. Most common scenarios,
// including the generation of multiple keys, should use [Key] instead.
func Extract[H hash.Hash](h func() H, secret, salt []byte) ([]byte, error)

// Expand derives a key from the given hash, key, and optional context info,
// returning a []byte of length keyLength that can be used as cryptographic key.
// The extraction step is skipped.
//
// The key should have been generated by [Extract], or be a uniformly
// random or pseudorandom cryptographically strong key. See RFC 5869, Section
// 3.3. Most common scenarios will want to use [Key] instead.
func Expand[H hash.Hash](h func() H, pseudorandomKey []byte, info string, keyLength int) ([]byte, error)

// Key derives a key from the given hash, secret, salt and context info,
// returning a []byte of length keyLength that can be used as cryptographic key.
// Salt and info can be nil.
func Key[Hash hash.Hash](h func() Hash, secret, salt []byte, info string, keyLength int) ([]byte, error)