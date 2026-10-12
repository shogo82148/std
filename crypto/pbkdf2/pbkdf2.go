// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pbkdf2パッケージは、RFC 8018 (PKCS #5 v2.1) で定義されている
// 鍵導出関数PBKDF2を実装します。
//
// 鍵導出関数は、パスワードやその他の完全にはランダムでないデータに基づいて
// データを暗号化する場合に役立ちます。擬似乱数関数を使用して、
// パスワードに基づいた安全な暗号鍵を導出します。
package pbkdf2

import (
	"github.com/shogo82148/std/hash"
)

// Keyはパスワード、ソルト、反復回数から鍵を導出し、暗号鍵として使用できる
// 長さkeyLengthの[]byteを返します。鍵は、指定されたハッシュ関数を用いた
// HMAC版のPBKDF2として説明されている方法に基づいて導出されます。
//
// 例えば、HMAC-SHA-1ベースのPBKDF2鍵導出関数を使用する場合、
// AES-256用の導出鍵（32バイトの鍵が必要）は次のようにして得られます:
//
//	dk, err := pbkdf2.Key(sha1.New, "some password", salt, 4096, 32)
//
// 良質なランダムソルトを使用することを忘れないでください。RFCでは
// 少なくとも8バイトが推奨されています。
//
// 反復回数を増やすと全数探索のコストが増加しますが、導出も比例して遅くなります。
//
// keyLengthは1から(2^32 - 1) * h.Size()の間の正の整数でなければなりません。
// keyLengthをこの範囲外の値に設定するとエラーになります。
func Key[Hash hash.Hash](h func() Hash, password string, salt []byte, iter, keyLength int) ([]byte, error)
