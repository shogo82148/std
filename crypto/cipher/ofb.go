// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// OFB（Output Feedback）モード。

package cipher

// NewOFBは、ブロック暗号bを使用して暗号化または復号化する [Stream] を返します。
// 初期化ベクトルivの長さは、bのブロックサイズと等しくなければなりません。
func NewOFB(b Block, iv []byte) Stream
