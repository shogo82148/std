// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

// NewGCMは、与えられた128ビットのブロック暗号を標準のナンス長でGalois Counter Modeでラップしたものを返します。
//
// 一般的に、GCMのこの実装で実行されるGHASH操作は一定時間ではありません。
// aes.NewCipherで生成された基礎の [Block] が、AESのハードウェアサポートを持つシステムである場合は例外です。詳細については、 [crypto/aes] パッケージのドキュメントを参照してください。
func NewGCM(cipher Block) (AEAD, error)

// NewGCMWithNonceSize は、与えられた長さの非スタンダードなノンスを受け付ける、128-bitのブロック暗号をGalios Counter Modeでラップしたものを返します。長さはゼロであってはいけません。
// 他の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。他のユーザーは、より高速でミス使用に対してより抵抗力のある [NewGCM] を使用すべきです。
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)

// NewGCMWithTagSizeは、指定された128ビットのブロック暗号をGalois Counter Modeでラップし、指定された長さのタグを生成します。
// 12バイトから16バイトのタグサイズが許可されています。
// 非標準のタグ長を使用する既存の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。その他のユーザーは、誤用に対してより耐性がある [NewGCM] を使用するべきです。
func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)

// NewGCMWithRandomNonceは、与えられた暗号をランダムに生成されたナンスを使用してGalois Counter
// Modeでラップしたものを返します。暗号は [crypto/aes.NewCipher] によって作成されている必要があります。
//
// この関数は96ビットのランダムナンスを生成し、Sealによって暗号文の先頭に付加され、
// Openによって暗号文から抽出されます。AEADのNonceSizeは0であり、
// Overheadは28バイト（ナンスサイズとタグサイズの組み合わせ）です。
//
// ランダムナンスの衝突リスクを無視できるレベルに制限するため、
// 与えられた鍵は2^32個を超えるメッセージの暗号化に使用してはいけません。
func NewGCMWithRandomNonce(cipher Block) (AEAD, error)
