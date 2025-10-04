// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
hmacパッケージは、米国連邦情報処理標準出版物198で定義されているキー付きハッシュメッセージ認証コード（HMAC）を実装しています。
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

<<<<<<< HEAD
// Newは指定した [hash.Hash] タイプとキーを使用して新しいHMACハッシュを返します。
// [crypto/sha256] からのsha256.NewのようなNew関数はhとして使用できます。
// hは呼び出されるたびに新しいハッシュを返す必要があります。
// 標準ライブラリの他のハッシュ実装とは異なり、返されたハッシュは [encoding.BinaryMarshaler] または [encoding.BinaryUnmarshaler] を実装していません。
=======
// New returns a new HMAC hash using the given [hash.Hash] type and key.
// New functions like [crypto/sha256.New] can be used as h.
// h must return a new Hash every time it is called.
// Note that unlike other hash implementations in the standard library,
// the returned Hash does not implement [encoding.BinaryMarshaler]
// or [encoding.BinaryUnmarshaler].
>>>>>>> upstream/release-branch.go1.25
func New(h func() hash.Hash, key []byte) hash.Hash

// Equalは、タイミング情報を漏らさずに2つのMACを比較します。
func Equal(mac1, mac2 []byte) bool
