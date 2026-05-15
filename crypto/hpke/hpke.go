// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hpke は [RFC 9180] で定義されたハイブリッド公開鍵暗号化 (HPKE)
// を実装します。
//
// [RFC 9180]: https://www.rfc-editor.org/rfc/rfc9180.html
package hpke

// Sender は送信 HPKE コンテキストです。特定の KEM
// カプセル化鍵、つまり公開鍵で作成され、ステートフルで
// 各 [Sender.Seal] 呼び出しで適切にノンスカウンターを増やします。
type Sender struct {
	*context
}

// Recipient は受信 HPKE コンテキストです。特定の KEM
// カプセル化解除鍵、つまり秘密鍵で作成され、ステートフルで
// 各成功した [Recipient.Open] 呼び出しで適切にノンスカウンターを増やします。
type Recipient struct {
	*context
}

// NewSender は、提供された KEM カプセル化鍵、つまり公開鍵のための
// 送信 HPKE コンテキストを返します。KEM、KDF、AEAD の組み合わせで
// 定義された暗号スイートを使用します。
//
// info パラメーターは、送信者と受信者間で一致しなくてはいけない
// 追加公開情報です。
//
// 返された enc 暗号文は、対応する KEM カプセル化解除鍵を
// 履行化するために、対応する受信 HPKE コンテキストを
// 使用できます。
func NewSender(pk PublicKey, kdf KDF, aead AEAD, info []byte) (enc []byte, s *Sender, err error)

// NewRecipient は、提供された KEM カプセル化解除鍵、つまり秘密鍵のための
// 受信 HPKE コンテキストを返します。KEM、KDF、AEAD の組み合わせで
// 定義された暗号スイートを使用します。
//
// enc パラメーターは、対応する KEM カプセル化鍵を持つ一致した
// 送信 HPKE コンテキストで作成された必要があります。info パラメーターは
// 送信者と受信者間で一致しなくてはいけない追加公開情報です。
func NewRecipient(enc []byte, k PrivateKey, kdf KDF, aead AEAD, info []byte) (*Recipient, error)

// Seal は、提供された平文を暗号化し、理想的に追加公開データ aad に
// 紐付けします。
//
// Seal は各呼び出しに従ってノンスカウンターを使用し、受信側の Open は
// Seal と同じ順序で呼び出す必要があります。
func (s *Sender) Seal(aad, plaintext []byte) ([]byte, error)

// Seal は [NewSender] のように一回限りの送信 HPKE コンテキストを初期化し
// てから、[Sender.Seal] のように提供された平文を暗号化します（aad なし）。
// Seal はカプセル化鍵と暗号文の連結を返します。
func Seal(pk PublicKey, kdf KDF, aead AEAD, info, plaintext []byte) ([]byte, error)

// Export は、送信者と受信者間で一致する共有鍵から導出した
// 秘密値を生成します。length は 65535 以下である必要があります。
func (s *Sender) Export(exporterContext string, length int) ([]byte, error)

// Open は、提供された暗号文を複号化し、理想的に追加公開データ aad に
// 紐付けします。複号が失敗した場合、エラーを返します。
//
// Open は各成功した呼び出しに従ってノンスカウンターを使用し、
// 送信側の Seal と同じ順序で呼び出す必要があります。
func (r *Recipient) Open(aad, ciphertext []byte) ([]byte, error)

// Open は [NewRecipient] のように一回限りの受信 HPKE コンテキストを
// 初期化してから、[Recipient.Open] のように提供された暗号文を複号化します
// （aad なし）。ciphertext はカプセル化鍵と実際の暗号文の連結である
// 必要があります。
func Open(k PrivateKey, kdf KDF, aead AEAD, info, ciphertext []byte) ([]byte, error)

// Export は、送信者と受信者間で一致する共有鍵から導出した
// 秘密値を生成します。length は 65535 以下である必要があります。
func (r *Recipient) Export(exporterContext string, length int) ([]byte, error)
