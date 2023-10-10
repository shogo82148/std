// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 暗号ブロックチェーン (CBC) モード。

// CBCは、ブロック暗号を適用する前に、前の暗号文ブロックと各平文ブロックをXOR演算（チェイニング）することで機密性を提供します。

// NIST SP 800-38A、pp 10-11を参照してください。

package cipher

// NewCBCEncrypterは、与えられたBlockを使用して、暗号ブロック連鎖モードで暗号化するBlockModeを返します。ivの長さは、Blockのブロックサイズと同じでなければなりません。
func NewCBCEncrypter(b Block, iv []byte) BlockMode

// NewCBCDecrypterは、与えられたBlockを使用して、暗号ブロックチェーンモードで復号化するためのBlockModeを返します。ivの長さは、Blockのブロックサイズと同じでなければならず、データの暗号化に使用されたivと一致する必要があります。
func NewCBCDecrypter(b Block, iv []byte) BlockMode
