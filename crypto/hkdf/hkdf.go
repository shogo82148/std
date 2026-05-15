// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hkdf は RFC 5869 で定義されたHMAC ベースの抽出・展開鍵導出
// 関数 (HKDF) を実装します。
//
// HKDF は、限定的な入力鍵材を 1 つ以上の暗号学的に強い秘密鍵に
// 展開することを目標とした暗号学的鍵導出関数 (KDF) です。
package hkdf

import (
	"github.com/shogo82148/std/hash"
)

// Extract は入力秘密と、オプションで独立したソルトから、[Expand] で
// 使用する擬似乱数鍵を生成します。
//
// 抽出した鍵を複数の Expand 呼び出しと異なるコンテキスト値で再利用する
// 必要がある場合のみこの関数を使用してください。複数の鍵を生成する場合を
// 含むほとんどの一般的なシナリオでは、代わりに [Key] を使用してください。
func Extract[H hash.Hash](h func() H, secret, salt []byte) ([]byte, error)

// Expand は、指定されたハッシュ、鍵、およびオプションのコンテキスト情報から
// 鍵を導出し、暗号学的鍵として使用できる keyLength 長の []byte を返します。
// 抽出ステップはスキップされます。
//
// 鍵は [Extract] によって生成されたか、均一にランダムまたは
// 擬似ランダムな暗号学的に強い鍵である必要があります。RFC 5869 Section
// 3.3 を参照してください。ほとんどの一般的なシナリオでは [Key] を
// 使用することをお勧めします。
func Expand[H hash.Hash](h func() H, pseudorandomKey []byte, info string, keyLength int) ([]byte, error)

// Key は、指定されたハッシュ、秘密、ソルト、およびコンテキスト情報から
// 鍵を導出し、暗号学的鍵として使用できる keyLength 長の []byte を返します。
// ソルトとコンテキスト情報は nil にできます。
func Key[Hash hash.Hash](h func() Hash, secret, salt []byte, info string, keyLength int) ([]byte, error)
