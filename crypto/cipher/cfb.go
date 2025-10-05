// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CFB（Cipher Feedback）モード。

package cipher

// NewCFBEncrypterは、与えられた [Block] を使用して、暗号フィードバックモードで暗号化する [Stream] を返します。
// ivは [Block] のブロックサイズと同じ長さである必要があります。
//
// Deprecated: CFBモードは認証されておらず、一般的に平文を操作して復元するアクティブ攻撃を可能にします。
// アプリケーションでは代わりに [AEAD] モードを使用することを推奨します。CFBの標準ライブラリ実装は
// 最適化されておらず、FIPS 140-3モジュールの一部として検証されていません。
// 認証されていない [Stream] モードが必要な場合は、代わりに [NewCTR] を使用してください。
func NewCFBEncrypter(block Block, iv []byte) Stream

// NewCFBDecrypterは、与えられた [Block] を使用して、暗号フィードバックモードで復号化する [Stream] を返します。
// ivは [Block] のブロックサイズと同じ長さである必要があります。
//
// Deprecated: CFBモードは認証されておらず、一般的に平文を操作して復元するアクティブ攻撃を可能にします。
// アプリケーションでは代わりに [AEAD] モードを使用することを推奨します。CFBの標準ライブラリ実装は
// 最適化されておらず、FIPS 140-3モジュールの一部として検証されていません。
// 認証されていない [Stream] モードが必要な場合は、代わりに [NewCTR] を使用してください。
func NewCFBDecrypter(block Block, iv []byte) Stream
