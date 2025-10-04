// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

// BlockSize is the AES block size in bytes.
const BlockSize = 16

// A Block is an instance of AES using a particular key.
// It is safe for concurrent use.
type Block struct {
	block
}

type KeySizeError int

func (k KeySizeError) Error() string

// New creates and returns a new [cipher.Block] implementation.
// The key argument should be the AES key, either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func New(key []byte) (*Block, error)

func (c *Block) BlockSize() int

func (c *Block) Encrypt(dst, src []byte)

func (c *Block) Decrypt(dst, src []byte)

// EncryptBlockInternal applies the AES encryption function to one block.
//
// It is an internal function meant only for the gcm package.
func EncryptBlockInternal(c *Block, dst, src []byte)
