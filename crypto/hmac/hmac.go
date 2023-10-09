// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package hmacは、米国連邦情報処理標準出版物198で定義されているキー付きハッシュメッセージ認証コード（HMAC）を実装しています。
HMACは、メッセージに署名するためにキーを使用する暗号ハッシュです。
受信者は、同じキーを使用してハッシュを再計算することで、ハッシュを検証します。

タイミングの副作用を回避するために、受信者はMACを比較するためにEqualを使用することに注意する必要があります：

	// ValidMACは、messageMACがメッセージの有効なHMACタグであるかどうかを報告します。
	func ValidMAC(message、messageMAC、key []byte) bool {
		mac := hmac.New(sha256.New、key)
		mac.Write(message)
		expectedMAC := mac.Sum(nil)
		return hmac.Equal(messageMAC、expectedMAC)
	}
*/package hmac

import (
	"github.com/shogo82148/std/hash"
)

// Newは指定したhash.Hashタイプとキーを使用して新しいHMACハッシュを返します。
// crypto/sha256からのsha256.NewのようなNew関数はhとして使用できます。
// hは呼び出されるたびに新しいハッシュを返す必要があります。
// 標準ライブラリの他のハッシュ実装とは異なり、返されたハッシュはencoding.BinaryMarshalerまたはencoding.BinaryUnmarshalerを実装していません。
func New(h func() hash.Hash, key []byte) hash.Hash

// Equalは、タイミング情報を漏らさずに2つのMACを比較します。
func Equal(mac1, mac2 []byte) bool
