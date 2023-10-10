// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package des_test

import "github.com/shogo82148/std/crypto/des"

func ExampleNewTripleDESCipher() {

	// NewTripleDESCipherは、最初の8バイトを16バイトのキーの複製として使用することで、EDE2が必要な場合にも使用することができます。
	ede2Key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	_, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	// 暗号化と復号化にcipher.Blockを使用する方法は、crypto/cipherを参照してください。
}
