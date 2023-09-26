// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// The AES block size in bytes.
const BlockSize = 16

// A cipher is an instance of AES encryption using a particular key.

type KeySizeError int

func (k KeySizeError) Error() string

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewCipher(key []byte) (cipher.Block, error)
