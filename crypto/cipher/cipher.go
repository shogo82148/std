// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// cipherパッケージは、低レベルのブロック暗号実装を包み込むことができる標準のブロック暗号モードを実装しています。
// 詳細はhttps://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html
// およびNIST Special Publication 800-38Aを参照してください。
package cipher

// A Block represents an implementation of block cipher
// using a given key. It provides the capability to encrypt
// or decrypt individual blocks. The mode implementations
// extend that capability to streams of blocks.
// ブロックは与えられた鍵を使用したブロック暗号の実装を表します。個々のブロックを暗号化または復号する機能を提供します。モードの実装は、ブロックのストリームにこの機能を拡張します。
type Block interface {
	BlockSize() int

	Encrypt(dst, src []byte)

	Decrypt(dst, src []byte)
}

// Streamはストリーム暗号を表します。
type Stream interface {
	XORKeyStream(dst, src []byte)
}

// BlockModeは、ブロックベースのモード（CBC、ECBなど）で動作するブロック暗号を表します。
type BlockMode interface {
	BlockSize() int

	CryptBlocks(dst, src []byte)
}
