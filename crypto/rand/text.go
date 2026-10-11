// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// Textは、標準のRFC 4648 base32アルファベットを使用した暗号論的にランダムな文字列を返します。
// これは、秘密の文字列、トークン、パスワード、その他のテキストが必要な場合に使用するためのものです。
// 結果には少なくとも128ビットのランダム性が含まれており、総当たりによる推測攻撃を防ぎ、
// 衝突の可能性を限りなく小さくするのに十分です。
// 将来のバージョンでは、これらの特性を維持するために必要に応じてより長いテキストを返す可能性があります。
func Text() string
