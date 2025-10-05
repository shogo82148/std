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
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/encoding/hex"
	"github.com/shogo82148/std/encoding/pem"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
)

func ExampleGenerateKey() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating RSA key: %s", err)
		return
	}

	der, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling RSA private key: %s", err)
		return
	}

	fmt.Printf("%s", pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	}))
}

func ExampleGenerateKey_testKey() {
	// これはRFC 9500, Section 2.1の安全でないテスト専用キーです。
	// 遅いキー生成を避けるためにテストで使用できます。
	block, _ := pem.Decode([]byte(strings.ReplaceAll(
		`-----BEGIN RSA TESTING KEY-----
MIIEowIBAAKCAQEAsPnoGUOnrpiSqt4XynxA+HRP7S+BSObI6qJ7fQAVSPtRkqso
tWxQYLEYzNEx5ZSHTGypibVsJylvCfuToDTfMul8b/CZjP2Ob0LdpYrNH6l5hvFE
89FU1nZQF15oVLOpUgA7wGiHuEVawrGfey92UE68mOyUVXGweJIVDdxqdMoPvNNU
l86BU02vlBiESxOuox+dWmuVV7vfYZ79Toh/LUK43YvJh+rhv4nKuF7iHjVjBd9s
B6iDjj70HFldzOQ9r8SRI+9NirupPTkF5AKNe6kUhKJ1luB7S27ZkvB3tSTT3P59
3VVJvnzOjaA1z6Cz+4+eRvcysqhrRgFlwI9TEwIDAQABAoIBAEEYiyDP29vCzx/+
dS3LqnI5BjUuJhXUnc6AWX/PCgVAO+8A+gZRgvct7PtZb0sM6P9ZcLrweomlGezI
FrL0/6xQaa8bBr/ve/a8155OgcjFo6fZEw3Dz7ra5fbSiPmu4/b/kvrg+Br1l77J
aun6uUAs1f5B9wW+vbR7tzbT/mxaUeDiBzKpe15GwcvbJtdIVMa2YErtRjc1/5B2
BGVXyvlJv0SIlcIEMsHgnAFOp1ZgQ08aDzvilLq8XVMOahAhP1O2A3X8hKdXPyrx
IVWE9bS9ptTo+eF6eNl+d7htpKGEZHUxinoQpWEBTv+iOoHsVunkEJ3vjLP3lyI/
fY0NQ1ECgYEA3RBXAjgvIys2gfU3keImF8e/TprLge1I2vbWmV2j6rZCg5r/AS0u
pii5CvJ5/T5vfJPNgPBy8B/yRDs+6PJO1GmnlhOkG9JAIPkv0RBZvR0PMBtbp6nT
Y3yo1lwamBVBfY6rc0sLTzosZh2aGoLzrHNMQFMGaauORzBFpY5lU50CgYEAzPHl
u5DI6Xgep1vr8QvCUuEesCOgJg8Yh1UqVoY/SmQh6MYAv1I9bLGwrb3WW/7kqIoD
fj0aQV5buVZI2loMomtU9KY5SFIsPV+JuUpy7/+VE01ZQM5FdY8wiYCQiVZYju9X
Wz5LxMNoz+gT7pwlLCsC4N+R8aoBk404aF1gum8CgYAJ7VTq7Zj4TFV7Soa/T1eE
k9y8a+kdoYk3BASpCHJ29M5R2KEA7YV9wrBklHTz8VzSTFTbKHEQ5W5csAhoL5Fo
qoHzFFi3Qx7MHESQb9qHyolHEMNx6QdsHUn7rlEnaTTyrXh3ifQtD6C0yTmFXUIS
CW9wKApOrnyKJ9nI0HcuZQKBgQCMtoV6e9VGX4AEfpuHvAAnMYQFgeBiYTkBKltQ
XwozhH63uMMomUmtSG87Sz1TmrXadjAhy8gsG6I0pWaN7QgBuFnzQ/HOkwTm+qKw
AsrZt4zeXNwsH7QXHEJCFnCmqw9QzEoZTrNtHJHpNboBuVnYcoueZEJrP8OnUG3r
UjmopwKBgAqB2KYYMUqAOvYcBnEfLDmyZv9BTVNHbR2lKkMYqv5LlvDaBxVfilE0
2riO4p6BaAdvzXjKeRrGNEKoHNBpOSfYCOM16NjL8hIZB1CaV3WbT5oY+jp7Mzd5
7d56RZOE+ERK2uz/7JX9VSsM/LbH9pJibd4e8mikDS9ntciqOH/3
-----END RSA TESTING KEY-----`, "TESTING KEY", "PRIVATE KEY")))
	testRSA2048, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	fmt.Println("Private key bit size:", testRSA2048.N.BitLen())
}

// RSAは非常に限られた量のデータしか暗号化できません。合理的な量の
// データを暗号化するために、ハイブリッドスキームが一般的に
// 使用されます：RSAはAES-GCMのような対称プリミティブの
// キーを暗号化するために使用されます。
//
// 暗号化の前に、データは既知の
// 構造に埋め込むことで「パディング」されます。これは多くの理由で行われますが、最も
// 明らかなのは、べき乗が剰余より大きくなるように
// 値が十分に大きいことを保証することです。（そうでなければ平方根で
// 復号化される可能性があります。）
//
// これらの設計では、PKCS #1 v1.5を使用する場合、受信したRSAメッセージが
// 正しい形式であったかどうか（つまり、復号化の結果が正しく
// パディングされたメッセージであるかどうか）を開示することを避けることが
// 極めて重要です。これは秘密情報を漏洩するためです。
// DecryptPKCS1v15SessionKeyはこの状況のために設計されており、
// 復号化された対称キー（正しい形式の場合）を、ランダムキーを含む
// バッファに定数時間でコピーします。したがって、RSAの結果が
// 正しい形式でない場合、実装は定数時間でランダムキーを使用します。
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
