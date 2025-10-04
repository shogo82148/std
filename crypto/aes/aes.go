// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// aesパッケージは、U.S.連邦情報処理標準出版物197で定義されているAES暗号（以前はRijndaelとして知られていた）を実装しています。
//
// このパッケージのAES操作は、定数時間アルゴリズムを使用して実装されていません。
// ただし、AESのハードウェアサポートが有効なシステムで実行される場合は例外です。
// これらの操作は、AES-NI拡張を使用しているamd64システムや
// Message-Security-Assist拡張を使用しているs390xシステムなどが該当します。
// このようなシステムでは、NewCipherの結果がcipher.NewGCMに渡される場合、
// GCMで使用されるGHASH操作も定数時間です。
package aes

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// The AES block size in bytes.
const BlockSize = 16

type KeySizeError int

func (k KeySizeError) Error() string

// NewCipher creates and returns a new [cipher.Block].
// The key argument must be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewCipher(key []byte) (cipher.Block, error)
