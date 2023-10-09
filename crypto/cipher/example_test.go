// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/crypto/aes"
	"github.com/shogo82148/std/crypto/cipher"
	"github.com/shogo82148/std/crypto/rand"
	"github.com/shogo82148/std/encoding/hex"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
)

func ExampleNewGCM_encrypt() {

	// 安全な場所から秘密鍵を読み込み、複数のSeal/Open呼び出しで再利用します。
	//（もちろん、実際の用途にはこの例の鍵を使用しないでください。）
	// パスフレーズを鍵に変換したい場合は、bcryptやscryptのような適切な
	// パッケージを使用してください。
	// デコードされた鍵は16バイト（AES-128）または32バイト（AES-256）である必要があります。
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("exampleplaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// 同じキーで2^32以上のランダムなノンスを使用しないでください。繰り返しのリスクがあるためです。
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}

func ExampleNewGCM_decrypt() {

	// 安全な場所から秘密のキーを読み込み、複数のSeal/Open呼び出し間で再利用してください。
	// （もちろん、実際の用途にはこの例のキーを使用しないでください。）
	// パスフレーズをキーに変換したい場合は、bcryptやscryptなどの適切なパッケージを使用してください。
	// キーをデコードすると、16バイト（AES-128）または32バイト（AES-256）である必要があります。
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	ciphertext, _ := hex.DecodeString("c3aaa29f002ca75870806e44086700f62ce4d43e902b3888e23ceff797a7a471")
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
	// Output: exampleplaintext
}

func ExampleNewCBCDecrypter() {

	// 安全な場所から秘密鍵を読み込んで、複数の NewCipher 呼び出し間で再利用してください。
	// （もちろん、実際の用途にはこの例の鍵を使用しないでください。）
	// パスフレーズを鍵に変換したい場合は、bcrypt や scrypt のような適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("73c86d43a9d700a253a96c85b0f6b03ac9792e0e757f869cca306bd3cba1c62b")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IVは一意である必要がありますが、セキュリティは必要ありません。
	// そのため、しばしば暗号文の先頭に含まれます。
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBCモードは常に完全なブロックで動作します。
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks は、引数が同じであればその場で処理されます。
	mode.CryptBlocks(ciphertext, ciphertext)

	// もし元の平文の長さがブロックの倍数でない場合、暗号化する際に追加する必要があるパディングがこの時点で削除されます。例としては、https://tools.ietf.org/html/rfc5246#section-6.2.3.2 を参照してください。ただし、パディングオラクルを作成しないために、暗号文を複合化する前に必ず認証すること（つまり、crypto/hmacを使用すること）が非常に重要です。

	fmt.Printf("%s\n", ciphertext)
	// Output: exampleplaintext
}

func ExampleNewCBCEncrypter() {

	// 安全な場所から秘密の鍵をロードし、複数の NewCipher 呼び出しで再利用します。
	// (もちろん、実際の目的にはこの例の鍵を使用しないでください。)
	// もしパスフレーズを鍵に変換したい場合は、bcrypt や scrypt のような適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("exampleplaintext")

	// CBCモードでは、平文はブロック単位で処理されるため、次の完全なブロックまでパディングする必要がある場合があります。このようなパディングの例については、次を参照してください：https://tools.ietf.org/html/rfc5246#section-6.2.3.2。ここでは、平文が既に正しい長さであると仮定します。
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IVはユニークである必要がありますが、セキュリティは求められません。
	// そのため、一般的には暗号文の先頭に含めることがあります。
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 暗号文は、安全にするために暗号化されるだけでなく、
	// (つまり、crypto/hmacを使用することによって)認証されている必要があることを忘れないことが重要です。

	fmt.Printf("%x\n", ciphertext)
}

func ExampleNewCFBDecrypter() {

	// 安全な場所から秘密キーを読み込み、複数のNewCipher呼び出しで再利用してください。
	// （もちろん、実際にはこの例のキーを使用しないでください。）パスフレーズをキーに変換したい場合は、bcryptやscryptなど適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("7dd015f06bec7f1b8f6559dad89f4131da62261786845100056b353194ad")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IVは一意である必要がありますが、安全性は問われません。したがって、通常は暗号文の先頭に含まれます。
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// もし2つの引数が同じ場合、XORKeyStreamはインプレースで動作することができます。
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext)
	// Output: some plaintext
}

func ExampleNewCFBEncrypter() {

	// 安全な場所から秘密鍵を読み込み、複数の NewCipher 呼び出しで再利用してください。 （明らかに、実際の何かのためにこの例の鍵を使用しないでください）。 パスフレーズを鍵に変換したい場合は、bcrypt や scrypt のような適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IVは一意である必要がありますが、安全である必要はありません。したがって、一般的には、暗号文の先頭にIVを含めることがあります。
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// 暗号文は、安全性を確保するために、暗号化だけでなく、認証（crypto/hmac の使用によって）も行われる必要があることを覚えておくことが重要です。
	fmt.Printf("%x\n", ciphertext)
}

func ExampleNewCTR() {

	// 安全な場所から秘密キーを読み込み、複数のNewCipher呼び出しで再利用します。
	// （もちろん、実際の用途にはこの例のキーを使用しないでください。）
	// パスフレーズをキーに変換したい場合は、bcryptやscryptのような適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IVは一意である必要がありますが、セキュリティは必要ありません。そのため、一般的には暗号文の先頭に含まれます。
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// 暗号文は安全にするために、暗号化するだけでなく、認証（つまりcrypto/hmacを使用すること）することも重要であることを忘れないようにする必要があります。

	// CTR モードは暗号化と復号化の両方に同じですので、NewCTR を使ってその暗号文を復号化することもできます。

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2)
	// Output: some plaintext
}

func ExampleNewOFB() {

	// 安全な場所から秘密鍵を読み込み、複数の NewCipher 呼び出しで再利用します。
	//（もちろん、実際の用途にはこの例の鍵を使用しないでください。）もしパスフレーズを鍵に変換したい場合は、bcrypt や scrypt のような適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IVは一意である必要がありますが、セキュリティは必要ありません。そのため、一般的には
	// 暗号文の先頭に含まれています。
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// 暗号文だけでなく、(crypto/hmacを使用して)認証も行われる必要があることを覚えておくことは重要です。これによってセキュリティが確保されます。

	// OFBモードは暗号化と復号化の両方において同じですので、NewOFBを使ってその暗号文を復号化することも可能です。

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewOFB(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2)
	// Output: some plaintext
}

func ExampleStreamReader() {

	// 安全な場所から秘密の鍵をロードし、複数の NewCipher 呼び出しで再利用してください。
	// （もちろん、これは実際には使用しないでください。）パスフレーズを鍵に変換したい場合は、bcrypt や scrypt のような適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	encrypted, _ := hex.DecodeString("cf0495cc6f75dafc23948538e79904a9")
	bReader := bytes.NewReader(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// もしキーがそれぞれの暗号文ごとにユニークである場合、ゼロの初期化ベクトル（IV）を使用しても問題ありません。
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	reader := &cipher.StreamReader{S: stream, R: bReader}
	// 入力を出力ストリームにコピーし、逐次復号化する。
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}

	// この例では、暗号化されたデータの認証を省略しているため、単純化されています。実際にこのようにStreamReaderを使用する場合、攻撃者は出力の任意のビットを反転させることができます。

	// Output: some secret text
}

func ExampleStreamWriter() {

	// 安全な場所から秘密キーを読み込み、複数の NewCipher 呼び出しで再利用します。
	// （もちろん、実際の用途でこの例のキーを使用しないでください。）
	// もしパスフレーズをキーに変換したい場合は、bcrypt や scrypt のような
	// 適切なパッケージを使用してください。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	bReader := bytes.NewReader([]byte("some secret text"))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// キーが各暗号文ごとにユニークな場合、ゼロのIVを使用することは問題ありません。
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}
	// 入力を出力バッファにコピーし、進行中に暗号化します。
	if _, err := io.Copy(writer, bReader); err != nil {
		panic(err)
	}

	// この例は暗号化されたデータの認証を省略して簡略化しています。実際にStreamReaderをこのように使用する場合、攻撃者が復号化された結果内の任意のビットを反転させる可能性があります。

	fmt.Printf("%x\n", out.Bytes())
	// Output: cf0495cc6f75dafc23948538e79904a9
}
