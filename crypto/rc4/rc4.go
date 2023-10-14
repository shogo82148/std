// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージrc4はBruce Schneierの「応用暗号化」で定義されている
// RC4暗号を実装しています。
//
// RC4は暗号学的に脆弱であり、安全なアプリケーションには使用すべきではありません。
package rc4

// Cipherは特定のキーを使用したRC4のインスタンスです。
type Cipher struct {
	s    [256]uint32
	i, j uint8
}

type KeySizeError int

func (k KeySizeError) Error() string

<<<<<<< HEAD
// NewCipherは新しいCipherを作成し、返します。キーアーギュメントはRC4キーであり、少なくとも1バイト、最大256バイトである必要があります。
func NewCipher(key []byte) (*Cipher, error)

// Resetはキーデータをゼロ化し、Cipherを使用できなくします。
=======
// NewCipher creates and returns a new [Cipher]. The key argument should be the
// RC4 key, at least 1 byte and at most 256 bytes.
func NewCipher(key []byte) (*Cipher, error)

// Reset zeros the key data and makes the [Cipher] unusable.
>>>>>>> upstream/master
//
// Deprecated: Resetはキーがプロセスのメモリから完全に削除されることを保証できません。
func (c *Cipher) Reset()

// XORKeyStreamは、キーストリームを使用してsrcとXOR演算した結果をdstに設定します。
// Dstとsrcは完全に重なるか、まったく重ならない必要があります。
func (c *Cipher) XORKeyStream(dst, src []byte)
