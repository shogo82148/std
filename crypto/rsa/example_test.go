// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa_test

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/aes"
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/crypto/rand"
	"github.com/shogo82148/std/crypto/rsa"
	"github.com/shogo82148/std/crypto/sha256"
	"github.com/shogo82148/std/encoding/hex"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
)

// RSAは非常に限られた量のデータしか暗号化できません。したがって、合理的な量のデータを暗号化するためには、一般的にハイブリッド方式が使用されます。具体的には、RSAはAES-GCMのような対称プリミティブの鍵を暗号化するために使用されます。
// 暗号化する前に、データは既知の構造に埋め込むことで「パディング」されます。これにはいくつかの理由がありますが、最も明らかな理由は、指数関数がモジュラスよりも大きい値になるようにするためです（そうしないと平方根で復号化できてしまいます）。
// これらの設計では、PKCS #1 v1.5を使用する場合、受信したRSAメッセージが形式に適合しているか（つまり、復号化の結果が正しくパディングされたメッセージか）を漏らさないようにすることが重要です。そのためにDecryptPKCS1v15SessionKeyはこの状況に対応しており、復号化された対称鍵が適切な形式であれば、ランダムなキーを含むバッファ上で一定時間内に対称鍵をコピーします。したがって、RSAの結果が形式に適合していない場合は、実装が一定時間内にランダムなキーを使用します。
func ExampleDecryptPKCS1v15SessionKey() {

	// ハイブリッド方式では、少なくとも16バイトの対称鍵を使用する必要があります。ここでは、RSA復号が正しく形成されていない場合に使用されるランダムな鍵を読み取ります。
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic("RNG failure")
	}

	rsaCiphertext, _ := hex.DecodeString("aabbccddeeff")

	if err := rsa.DecryptPKCS1v15SessionKey(nil, rsaPrivateKey, rsaCiphertext, key); err != nil {

		// 発生したエラーは「公開される」ものであり、秘密情報なしでも判断できます。（例えば、RSA公開鍵の長さが不可能な場合など）
		fmt.Fprintf(os.Stderr, "Error from RSA decryption: %s\n", err)
		return
	}

	// 与えられたキーを使用して、対称スキームを使ってより大きな暗号文を複合することができます。
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes.NewCipher failed: " + err.Error())
	}

	// キーがランダムであるため、固定されたNonceを使用することは許容されます。
	// (キー、Nonce)のペアは依然として一意である必要があります。
	var zeroNonce [12]byte
	aead, err := cipher.NewGCM(block)
	if err != nil {
		panic("cipher.NewGCM failed: " + err.Error())
	}
	ciphertext, _ := hex.DecodeString("00112233445566")
	plaintext, err := aead.Open(nil, zeroNonce[:], ciphertext, nil)
	if err != nil {

		// RSAの暗号文の形式が不正です。AES-GCMの鍵が正しくないため、復号化はここで失敗します。
		fmt.Fprintf(os.Stderr, "Error decrypting: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", plaintext)
}

func ExampleSignPKCS1v15() {
	message := []byte("message to be signed")

	// 直接署名できるのは小さなメッセージだけです。そのため、メッセージ自体ではなくそのハッシュを署名します。これにはハッシュ関数が衝突耐性がある必要があります。SHA-256は、執筆時点（2016年）では最も弱いハッシュ関数です。
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	fmt.Printf("Signature: %x\n", signature)
}

func ExampleVerifyPKCS1v15() {
	message := []byte("message to be signed")
	signature, _ := hex.DecodeString("ad2766728615cc7a746cc553916380ca7bfa4f8983b990913bc69eb0556539a350ff0f8fe65ddfd3ebe91fe1c299c2fac135bc8c61e26be44ee259f2f80c1530")

	// 直接署名できるのは小さなメッセージのみです。そのため、メッセージ自体ではなく、メッセージのハッシュが署名されます。これには、ハッシュ関数が衝突耐性を持つ必要があります。SHA-256は、書かれた時点（2016年）で使用すべき最も安全なハッシュ関数です。
	hashed := sha256.Sum256(message)

	err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	}

	// シグネチャは公開鍵からのメッセージの有効な署名です。
}

func ExampleEncryptOAEP() {
	secretMessage := []byte("send reinforcements, we're going to advance")
	label := []byte("orders")

	// crypto/rand.Readerは暗号化関数のランダム化において十分なエントロピー源です。
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &test2048Key.PublicKey, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	// 暗号化はランダムな関数のため、暗号文は毎回異なるものとなります。
	fmt.Printf("Ciphertext: %x\n", ciphertext)
}

func ExampleDecryptOAEP() {
	ciphertext, _ := hex.DecodeString("4d1ee10e8f286390258c51a5e80802844c3e6358ad6690b7285218a7c7ed7fc3a4c7b950fbd04d4b0239cc060dcc7065ca6f84c1756deb71ca5685cadbb82be025e16449b905c568a19c088a1abfad54bf7ecc67a7df39943ec511091a34c0f2348d04e058fcff4d55644de3cd1d580791d4524b92f3e91695582e6e340a1c50b6c6d78e80b4e42c5b4d45e479b492de42bbd39cc642ebb80226bb5200020d501b24a37bcc2ec7f34e596b4fd6b063de4858dbf5a4e3dd18e262eda0ec2d19dbd8e890d672b63d368768360b20c0b6b8592a438fa275e5fa7f60bef0dd39673fd3989cc54d2cb80c08fcd19dacbc265ee1c6014616b0e04ea0328c2a04e73460")
	label := []byte("orders")

	plaintext, err := rsa.DecryptOAEP(sha256.New(), nil, test2048Key, ciphertext, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", plaintext)

	// 暗号化は機密性のみを提供することを覚えておいてください。
	// メッセージが正当性を想定した前に、暗号文には署名する必要があります。さらに、メッセージは順序が変更される可能性も考慮してください。
}
