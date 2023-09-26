// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rc4 implements RC4 encryption, as defined in Bruce Schneier's
// Applied Cryptography.
//
// RC4 is cryptographically broken and should not be used for secure
// applications.
package rc4

// A Cipher is an instance of RC4 using a particular key.
type Cipher struct {
	s    [256]uint32
	i, j uint8
}

type KeySizeError int

func (k KeySizeError) Error() string

// NewCipher creates and returns a new Cipher. The key argument should be the
// RC4 key, at least 1 byte and at most 256 bytes.
func NewCipher(key []byte) (*Cipher, error)

// Reset zeros the key data and makes the Cipher unusable.
//
// Deprecated: Reset can't guarantee that the key will be entirely removed from
// the process's memory.
func (c *Cipher) Reset()

// XORKeyStream sets dst to the result of XORing src with the key stream.
// Dst and src must overlap entirely or not at all.
func (c *Cipher) XORKeyStream(dst, src []byte)
