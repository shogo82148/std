// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdsa

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Sign は、ハッシュ（大きなメッセージのハッシュの結果である必要があります）をプライベートキー priv を使用して署名します。
// もしハッシュがプライベートキーの曲線順序のビット長よりも長い場合、ハッシュはその長さに切り詰められます。
// 結果の署名は、2つの整数のペアとして返されます。ほとんどのアプリケーションでは、r, s と直接扱う代わりに [SignASN1] を使用するべきです。
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

// Verifyは、公開鍵pubを使用してハッシュのr、sの署名を検証します。
// 戻り値は署名の有効性を記録します。ほとんどのアプリケーションでは、r、sと直接取り扱う代わりに、VerifyASN1を使用する必要があります。
//
// 入力は機密とはみなされず、タイミングのサイドチャネルを通じて、または攻撃者が入力の一部を制御している場合に漏洩する可能性があります。
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
