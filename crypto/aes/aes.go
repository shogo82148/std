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

// AESのブロックサイズ（バイト単位）。
const BlockSize = 16

type KeySizeError int

func (k KeySizeError) Error() string

// NewCipherは新しい[cipher.Block]を作成して返します。
// key引数はAES鍵でなければならず、16、24、32バイトのいずれかで
// AES-128、AES-192、AES-256を選択します。
func NewCipher(key []byte) (cipher.Block, error)
