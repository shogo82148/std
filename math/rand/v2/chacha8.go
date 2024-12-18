// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/internal/chacha8rand"
)

// ChaCha8は、ChaCha8ベースの暗号的に強力な
// 乱数生成器です。
type ChaCha8 struct {
	state chacha8rand.State

	// The last readLen bytes of readBuf are still to be consumed by Read.
	readBuf [8]byte
	readLen int
}

// NewChaCha8は、指定されたシードで初期化された新しいChaCha8を返します。
func NewChaCha8(seed [32]byte) *ChaCha8

// Seedは、ChaCha8をNewChaCha8(seed)と同じように動作するようにリセットします。
func (c *ChaCha8) Seed(seed [32]byte)

// Uint64は、一様に分布したランダムなuint64値を返します。
func (c *ChaCha8) Uint64() uint64

// Readは、pに正確にlen(p)バイトを読み込みます。
// 常にlen(p)とnilエラーを返します。
//
// ReadとUint64の呼び出しが交互に行われる場合、
// 両者によって返されるビットの順序は未定義であり、
// Readは最後のUint64の呼び出し前に生成されたビットを返すことがあります。
func (c *ChaCha8) Read(p []byte) (n int, err error)

// UnmarshalBinaryはencoding.BinaryUnmarshalerインターフェースを実装します。
func (c *ChaCha8) UnmarshalBinary(data []byte) error

// MarshalBinaryはencoding.BinaryMarshalerインターフェースを実装します。
func (c *ChaCha8) MarshalBinary() ([]byte, error)
