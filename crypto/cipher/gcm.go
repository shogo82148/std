// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

// AEADは関連データを含めた認証暗号化を提供する暗号モードです。手法の説明については、以下を参照してください。
// https://en.wikipedia.org/wiki/Authenticated_encryption.
type AEAD interface {
	// NonceSize returns the size of the nonce that must be passed to Seal
	// and Open.
	NonceSize() int

	// Overhead returns the maximum difference between the lengths of a
	// plaintext and its ciphertext.
	Overhead() int

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	//
	// To reuse plaintext's storage for the encrypted output, use plaintext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	Seal(dst, nonce, plaintext, additionalData []byte) []byte

	// Open decrypts and authenticates ciphertext, authenticates the
	// additional data and, if successful, appends the resulting plaintext
	// to dst, returning the updated slice. The nonce must be NonceSize()
	// bytes long and both it and the additional data must match the
	// value passed to Seal.
	//
	// To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	//
	// Even if the function fails, the contents of dst, up to its capacity,
	// may be overwritten.
	Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}

// NewGCMは、標準のnonce長でラップされた128ビットのブロック暗号を返します。
//
<<<<<<< HEAD
// 一般的に、GCMのこの実装で実行されるGHASH操作は一定時間ではありません。
// aes.NewCipherで生成された基礎のブロックが、AESのハードウェアサポートを持つシステムである場合は例外です。詳細については、crypto/aesパッケージのドキュメントを参照してください。
func NewGCM(cipher Block) (AEAD, error)

// NewGCMWithNonceSize は、与えられた長さの非スタンダードなノンスを受け付ける、128-bitのブロック暗号をGalios Counter Modeでラップしたものを返します。長さはゼロであってはいけません。
// 他の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。他のユーザーは、より高速でミス使用に対してより抵抗力のあるNewGCMを使用すべきです。
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)

// NewGCMWithTagSizeは、指定された128ビットのブロック暗号をGalois Counter Modeでラップし、指定された長さのタグを生成します。
// 12バイトから16バイトのタグサイズが許可されています。
// 非標準のタグ長を使用する既存の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。その他のユーザーは、誤用に対してより耐性があるNewGCMを使用するべきです。
=======
// In general, the GHASH operation performed by this implementation of GCM is not constant-time.
// An exception is when the underlying [Block] was created by aes.NewCipher
// on systems with hardware support for AES. See the [crypto/aes] package documentation for details.
func NewGCM(cipher Block) (AEAD, error)

// NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois
// Counter Mode, which accepts nonces of the given length. The length must not
// be zero.
//
// Only use this function if you require compatibility with an existing
// cryptosystem that uses non-standard nonce lengths. All other users should use
// [NewGCM], which is faster and more resistant to misuse.
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)

// NewGCMWithTagSize returns the given 128-bit, block cipher wrapped in Galois
// Counter Mode, which generates tags with the given length.
//
// Tag sizes between 12 and 16 bytes are allowed.
//
// Only use this function if you require compatibility with an existing
// cryptosystem that uses non-standard tag lengths. All other users should use
// [NewGCM], which is more resistant to misuse.
>>>>>>> upstream/master
func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)
