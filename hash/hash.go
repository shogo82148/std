// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hash provides interfaces for hash functions.
package hash

import "github.com/shogo82148/std/io"

// Hash is the common interface implemented by all hash functions.
//
// Hash implementations in the standard library (e.g. [hash/crc32] and
// [crypto/sha256]) implement the [encoding.BinaryMarshaler], [encoding.BinaryAppender],
// [encoding.BinaryUnmarshaler] and [Cloner] interfaces. Marshaling a hash implementation
// allows its internal state to be saved and used for additional processing
// later, without having to re-write the data previously written to the hash.
// The hash state may contain portions of the input in its original form,
// which users are expected to handle for any possible security implications.
//
// Compatibility: Any future changes to hash or crypto packages will endeavor
// to maintain compatibility with state encoded using previous versions.
// That is, any released versions of the packages should be able to
// decode data written with any previously released version,
// subject to issues such as security fixes.
// See the Go compatibility document for background: https://golang.org/doc/go1compat
type Hash interface {
	io.Writer

	Sum(b []byte) []byte

	Reset()

	Size() int

	BlockSize() int
}

// Hash32 is the common interface implemented by all 32-bit hash functions.
type Hash32 interface {
	Hash
	Sum32() uint32
}

// Hash64 is the common interface implemented by all 64-bit hash functions.
type Hash64 interface {
	Hash
	Sum64() uint64
}

// A Cloner is a hash function whose state can be cloned.
//
// All [Hash] implementations in the standard library implement this interface,
// unless GOFIPS140=v1.0.0 is set.
//
// If a hash can only determine at runtime if it can be cloned,
// (e.g., if it wraps another hash), it may return [errors.ErrUnsupported].
type Cloner interface {
	Hash
	Clone() (Cloner, error)
}

// XOF (extendable output function) is a hash function with arbitrary or unlimited output length.
type XOF interface {
	io.Writer

	io.Reader

	Reset()

	BlockSize() int
}
