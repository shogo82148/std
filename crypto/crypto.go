// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crypto collects common cryptographic constants.
package crypto

import (
	"github.com/shogo82148/std/hash"
)

// Hash identifies a cryptographic hash function that is implemented in another
// package.
type Hash uint

const (
	MD4 Hash = 1 + iota
	MD5
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	MD5SHA1
	RIPEMD160
)

// Size returns the length, in bytes, of a digest resulting from the given hash
// function. It doesn't require that the hash function in question be linked
// into the program.
func (h Hash) Size() int

// New returns a new hash.Hash calculating the given hash function. New panics
// if the hash function is not linked into the binary.
func (h Hash) New() hash.Hash

// Available reports whether the given hash function is linked into the binary.
func (h Hash) Available() bool

// RegisterHash registers a function that returns a new instance of the given
// hash function. This is intended to be called from the init function in
// packages that implement hash functions.
func RegisterHash(h Hash, f func() hash.Hash)

// PublicKey represents a public key using an unspecified algorithm.
type PublicKey interface{}

// PrivateKey represents a private key using an unspecified algorithm.
type PrivateKey interface{}
