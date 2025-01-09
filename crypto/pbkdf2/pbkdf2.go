// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pbkdf2 implements the key derivation function PBKDF2 as defined in
// RFC 8018 (PKCS #5 v2.1).
//
// A key derivation function is useful when encrypting data based on a password
// or any other not-fully-random data. It uses a pseudorandom function to derive
// a secure encryption key based on the password.
package pbkdf2

import (
	"github.com/shogo82148/std/hash"
)

// Key derives a key from the password, salt and iteration count, returning a
// []byte of length keyLength that can be used as cryptographic key. The key is
// derived based on the method described as PBKDF2 with the HMAC variant using
// the supplied hash function.
//
// For example, to use a HMAC-SHA-1 based PBKDF2 key derivation function, you
// can get a derived key for e.g. AES-256 (which needs a 32-byte key) by
// doing:
//
//	dk := pbkdf2.Key(sha1.New, []byte("some password"), salt, 4096, 32)
//
// Remember to get a good random salt. At least 8 bytes is recommended by the
// RFC.
//
// Using a higher iteration count will increase the cost of an exhaustive
// search but will also make derivation proportionally slower.
func Key[Hash hash.Hash](h func() Hash, password string, salt []byte, iter, keyLength int) ([]byte, error)
