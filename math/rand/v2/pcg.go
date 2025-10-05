// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// PCGは、128ビットの内部状態を持つPCGジェネレータです。
// ゼロのPCGは、NewPCG(0, 0)と同等です。
type PCG struct {
	hi uint64
	lo uint64
}

// NewPCGは、与えられた値でシードされた新しいPCGを返します。
func NewPCG(seed1, seed2 uint64) *PCG

// Seedは、PCGをNewPCG(seed1, seed2)と同じように動作するようにリセットします。
func (p *PCG) Seed(seed1, seed2 uint64)

// AppendBinaryは [encoding.BinaryAppender] インターフェースを実装します。
func (p *PCG) AppendBinary(b []byte) ([]byte, error)

// MarshalBinaryは [encoding.BinaryMarshaler] インターフェースを実装します。
func (p *PCG) MarshalBinary() ([]byte, error)

// UnmarshalBinaryは [encoding.BinaryUnmarshaler] インターフェースを実装します。
func (p *PCG) UnmarshalBinary(data []byte) error

// Uint64は、一様に分布したランダムなuint64の値を返します。
func (p *PCG) Uint64() uint64
