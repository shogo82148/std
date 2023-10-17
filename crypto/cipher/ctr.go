// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// カウンタ (CTR) モード。

// CTRは、ブロック暗号をインクリメントするカウンタを暗号化し、結果として得られるデータストリームを入力とXOR演算することで、ストリーム暗号へ変換します。

// NIST SP 800-38A、pp 13-15を参照してください。

package cipher

// NewCTRは、指定された [Block] を使用して暗号化/復号化を行う [Stream] を返します。
// ivの長さは、 [Block] のブロックサイズと同じでなければなりません。
func NewCTR(block Block, iv []byte) Stream
