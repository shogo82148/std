// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha3 は、FIPS 202 で定義された SHA-3 ハッシュアルゴリズムと
// SHAKE 拡張出力関数を実装します。
package sha3

import (
	"github.com/shogo82148/std/crypto/internal/fips140/sha3"
	"github.com/shogo82148/std/hash"
)

// Sum224 は data の SHA3-224 ハッシュを返します。
func Sum224(data []byte) [28]byte

// Sum256 は data の SHA3-256 ハッシュを返します。
func Sum256(data []byte) [32]byte

// Sum384 は data の SHA3-384 ハッシュを返します。
func Sum384(data []byte) [48]byte

// Sum512 は data の SHA3-512 ハッシュを返します。
func Sum512(data []byte) [64]byte

// SumSHAKE128 は data に SHAKE128 拡張出力関数を適用し、
// 指定したバイト長の出力を返します。
func SumSHAKE128(data []byte, length int) []byte

// SumSHAKE256 は data に SHAKE256 拡張出力関数を適用し、
// 指定したバイト長の出力を返します。
func SumSHAKE256(data []byte, length int) []byte

// SHA3 は SHA-3 ハッシュのインスタンスです。[hash.Hash] を実装します。
// ゼロ値は使用可能な SHA3-256 ハッシュです。
type SHA3 struct {
	s sha3.Digest
}

// New224 は新しい SHA3-224 ハッシュを生成します。
func New224() *SHA3

// New256 は新しい SHA3-256 ハッシュを生成します。
func New256() *SHA3

// New384 は新しい SHA3-384 ハッシュを生成します。
func New384() *SHA3

// New512 は新しい SHA3-512 ハッシュを生成します。
func New512() *SHA3

// Write はハッシュの状態にさらにデータを取り込みます。
func (s *SHA3) Write(p []byte) (n int, err error)

// Sum は現在のハッシュを b に追加し、結果のスライスを返します。
func (s *SHA3) Sum(b []byte) []byte

// Reset はハッシュを初期状態に戻します。
func (s *SHA3) Reset()

// Size は Sum が生成するバイト数を返します。
func (s *SHA3) Size() int

// BlockSize はハッシュのレートを返します。
func (s *SHA3) BlockSize() int

// MarshalBinary は [encoding.BinaryMarshaler] を実装します。
func (s *SHA3) MarshalBinary() ([]byte, error)

// AppendBinary は [encoding.BinaryAppender] を実装します。
func (s *SHA3) AppendBinary(p []byte) ([]byte, error)

// UnmarshalBinary は [encoding.BinaryUnmarshaler] を実装します。
func (s *SHA3) UnmarshalBinary(data []byte) error

// Clone は [hash.Cloner] を実装します。
func (d *SHA3) Clone() (hash.Cloner, error)

// SHAKE は SHAKE 拡張出力関数のインスタンスです。
// ゼロ値は使用可能な SHAKE256 ハッシュです。
type SHAKE struct {
	s sha3.SHAKE
}

// NewSHAKE128 は新しい SHAKE128 XOF を生成します。
func NewSHAKE128() *SHAKE

// NewSHAKE256 は新しい SHAKE256 XOF を生成します。
func NewSHAKE256() *SHAKE

// NewCSHAKE128 は新しい cSHAKE128 XOF を生成します。
//
// N は cSHAKE ベースの関数を定義するために使われ、通常の cSHAKE が必要な
// 場合は空にできます。S はドメイン分離に使うカスタマイズ用バイト列です。
// N と S がともに空の場合、これは NewSHAKE128 と等価です。
func NewCSHAKE128(N, S []byte) *SHAKE

// NewCSHAKE256 は新しい cSHAKE256 XOF を生成します。
//
// N は cSHAKE ベースの関数を定義するために使われ、通常の cSHAKE が必要な
// 場合は空にできます。S はドメイン分離に使うカスタマイズ用バイト列です。
// N と S がともに空の場合、これは NewSHAKE256 と等価です。
func NewCSHAKE256(N, S []byte) *SHAKE

// Write は XOF の状態にさらにデータを取り込みます。
//
// すでに出力を読み出している場合は panic します。
func (s *SHAKE) Write(p []byte) (n int, err error)

// Read は XOF からさらに出力を絞り出します。
//
// Read を呼んだ後に Write を呼ぶと panic します。
func (s *SHAKE) Read(p []byte) (n int, err error)

// Reset は XOF を初期状態に戻します。
func (s *SHAKE) Reset()

// BlockSize は XOF のレートを返します。
func (s *SHAKE) BlockSize() int

// MarshalBinary は [encoding.BinaryMarshaler] を実装します。
func (s *SHAKE) MarshalBinary() ([]byte, error)

// AppendBinary は [encoding.BinaryAppender] を実装します。
func (s *SHAKE) AppendBinary(p []byte) ([]byte, error)

// UnmarshalBinary は [encoding.BinaryUnmarshaler] を実装します。
func (s *SHAKE) UnmarshalBinary(data []byte) error
