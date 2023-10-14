// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package des

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// DESのブロックサイズ（単位はバイト）。
const BlockSize = 8

type KeySizeError int

func (k KeySizeError) Error() string

<<<<<<< HEAD
// NewCipherは新しいcipher.Blockを作成して返します。
func NewCipher(key []byte) (cipher.Block, error)

// NewTripleDESCipher は新しい cipher.Block を作成して返します。
=======
// NewCipher creates and returns a new [cipher.Block].
func NewCipher(key []byte) (cipher.Block, error)

// NewTripleDESCipher creates and returns a new [cipher.Block].
>>>>>>> upstream/master
func NewTripleDESCipher(key []byte) (cipher.Block, error)
