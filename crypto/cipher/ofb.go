// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// OFB（Output Feedback）モード。

package cipher

// NewOFBは、ブロック暗号bを出力フィードバックモードで使用して暗号化または復号化する [Stream] を返します。
// 初期化ベクトルivの長さは、bのブロックサイズと等しくなければなりません。
//
// Deprecated: OFBモードは認証されておらず、一般的に平文を操作して復元するアクティブ攻撃を可能にします。
// アプリケーションでは代わりに [AEAD] モードを使用することを推奨します。OFBの標準ライブラリ実装は
// 最適化されておらず、FIPS 140-3モジュールの一部として検証されていません。
// 認証されていない [Stream] モードが必要な場合は、代わりに [NewCTR] を使用してください。
func NewOFB(b Block, iv []byte) Stream
