// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

type CTR struct {
	b          Block
	ivlo, ivhi uint64
	offset     uint64
}

func NewCTR(b *Block, iv []byte) *CTR

func (c *CTR) XORKeyStream(dst, src []byte)

// RoundToBlock is used by CTR_DRBG, which discards the rightmost unused bits at
// each request. It rounds the offset up to the next block boundary.
func RoundToBlock(c *CTR)

// XORKeyStreamAt behaves like XORKeyStream but keeps no state, and instead
// seeks into the keystream by the given bytes offset from the start (ignoring
// any XORKetStream calls). This allows for random access into the keystream, up
// to 16 EiB from the start.
func (c *CTR) XORKeyStreamAt(dst, src []byte, offset uint64)
