// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package des

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// The DES block size in bytes.
const BlockSize = 8

type KeySizeError int

func (k KeySizeError) Error() string

// desCipher is an instance of DES encryption.

// NewCipher creates and returns a new cipher.Block.
func NewCipher(key []byte) (cipher.Block, error)

// A tripleDESCipher is an instance of TripleDES encryption.

// NewTripleDESCipher creates and returns a new cipher.Block.
func NewTripleDESCipher(key []byte) (cipher.Block, error)
