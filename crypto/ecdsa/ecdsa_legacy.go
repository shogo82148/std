// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdsa

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Signは、秘密鍵privを使用してハッシュ（より大きなメッセージをハッシュした結果である必要があります）に署名します。
// ハッシュが秘密鍵の曲線位数のビット長より長い場合、ハッシュはその長さに切り詰められます。
// 署名は整数のペアとして返されます。ほとんどのアプリケーションでは、r、sと直接取り扱う代わりに
// [SignASN1] を使用する必要があります。
//
// 署名はランダム化されます。Go 1.26以降、安全なランダムバイトソースが常に使用され、
// GODEBUG=cryptocustomrand=1が設定されない限りReaderは無視されます。
// この設定は将来のGoリリースで削除されます。代わりに [testing/cryptotest.SetGlobalRandom] を使用してください。
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

// Verifyは、公開鍵pubを使用してハッシュのr、sの署名を検証します。
// 戻り値は署名の有効性を記録します。ほとんどのアプリケーションでは、r、sと直接取り扱う代わりに、VerifyASN1を使用する必要があります。
//
// 入力は機密とはみなされず、タイミングのサイドチャネルを通じて、または攻撃者が入力の一部を制御している場合に漏洩する可能性があります。
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
