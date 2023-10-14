// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// AESのブロックサイズ（バイト単位）。
const BlockSize = 16

type KeySizeError int

func (k KeySizeError) Error() string

<<<<<<< HEAD
// NewCipherは新しいcipher.Blockを作成して返します。
// key引数はAESキーである必要があります。
// AES-128、AES-192、またはAES-256を選択するために、
// 16バイト、24バイト、または32バイトのいずれかを指定します。
=======
// NewCipher creates and returns a new [cipher.Block].
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
>>>>>>> upstream/master
func NewCipher(key []byte) (cipher.Block, error)
