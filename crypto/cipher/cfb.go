// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CFB（Cipher Feedback）モード。

package cipher

// NewCFBEncrypterは、与えられたブロックを使用して、暗号フィードバックモードで暗号化するストリームを返します。IVはブロックのブロックサイズと同じ長さでなければなりません。
func NewCFBEncrypter(block Block, iv []byte) Stream

// NewCFBDecrypterは、暗号フィードバックモードで復号化するStreamを返します。
// ブロックとして指定されたものを使用します。IVはブロックのサイズと同じ長さでなければならない。
func NewCFBDecrypter(block Block, iv []byte) Stream
