// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crypto collects common cryptographic constants.
package crypto

import (
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

// Hash identifies a cryptographic hash function that is implemented in another
// package.
type Hash uint

// HashFunc simply returns the value of h so that Hash implements SignerOpts.
func (h Hash) HashFunc() Hash

func (h Hash) String() string

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
	SHA3_224
	SHA3_256
	SHA3_384
	SHA3_512
	SHA512_224
	SHA512_256
	BLAKE2s_256
	BLAKE2b_256
	BLAKE2b_384
	BLAKE2b_512
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

// Signer is an interface for an opaque private key that can be used for
// signing operations. For example, an RSA key kept in a hardware module.
type Signer interface {
	Public() PublicKey

	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}

// SignerOpts contains options for signing with a Signer.
type SignerOpts interface {
	HashFunc() Hash
}

// Decrypter is an interface for an opaque private key that can be used for
// asymmetric decryption operations. An example would be an RSA key
// kept in a hardware module.
type Decrypter interface {
	Public() PublicKey

	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts interface{}
