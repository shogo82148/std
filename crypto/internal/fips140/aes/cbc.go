// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

type CBCEncrypter struct {
	b  Block
	iv [BlockSize]byte
}

// NewCBCEncrypter returns a [cipher.BlockMode] which encrypts in cipher block
// chaining mode, using the given Block.
func NewCBCEncrypter(b *Block, iv [BlockSize]byte) *CBCEncrypter

func (c *CBCEncrypter) BlockSize() int

func (c *CBCEncrypter) CryptBlocks(dst, src []byte)

func (x *CBCEncrypter) SetIV(iv []byte)

type CBCDecrypter struct {
	b  Block
	iv [BlockSize]byte
}

// NewCBCDecrypter returns a [cipher.BlockMode] which decrypts in cipher block
// chaining mode, using the given Block.
func NewCBCDecrypter(b *Block, iv [BlockSize]byte) *CBCDecrypter

func (c *CBCDecrypter) BlockSize() int

func (c *CBCDecrypter) CryptBlocks(dst, src []byte)

func (x *CBCDecrypter) SetIV(iv []byte)
