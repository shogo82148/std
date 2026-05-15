// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkem は、[NIST FIPS 203] で指定された量子耐性鍵カプセル化方式
// ML-KEM (以前は Kyber として知られていた) を実装します。
//
// ほとんどのアプリケーションは、[DecapsulationKey768] と [EncapsulationKey768]
// で実装された ML-KEM-768 パラメータセットを使用すべきです。
//
// [NIST FIPS 203]: https://doi.org/10.6028/NIST.FIPS.203
package mlkem

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/internal/fips140/mlkem"
)

const (
	// SharedKeySize は、ML-KEM で生成される共有鍵のサイズです。
	SharedKeySize = 32

	// SeedSize は、カプセル化解除鍵を生成するために使用されるシードのサイズです。
	SeedSize = 64

	// CiphertextSize768 は、ML-KEM-768 で生成される暗号文のサイズです。
	CiphertextSize768 = 1088

	// EncapsulationKeySize768 は、ML-KEM-768 カプセル化鍵のサイズです。
	EncapsulationKeySize768 = 1184

	// CiphertextSize1024 は、ML-KEM-1024 で生成される暗号文のサイズです。
	CiphertextSize1024 = 1568

	// EncapsulationKeySize1024 は、ML-KEM-1024 カプセル化鍵のサイズです。
)

// DecapsulationKey768 は、暗号文から共有鍵をカプセル化解除するために使用される
// 秘密鍵です。様々な事前計算値が含まれています。
type DecapsulationKey768 struct {
	key *mlkem.DecapsulationKey768
}

// GenerateKey768 は、安全なソースからランダムバイトを取得して、新しい
// カプセル化解除鍵を生成します。カプセル化解除鍵は秘密に保つ必要があります。
func GenerateKey768() (*DecapsulationKey768, error)

// NewDecapsulationKey768 は、"d || z" 形式の 64 バイトシードから
// カプセル化解除鍵を展開します。シードは均一にランダムである必要があります。
func NewDecapsulationKey768(seed []byte) (*DecapsulationKey768, error)

// Bytes は、カプセル化解除鍵を "d || z" 形式の 64 バイトシードとして返します。
//
// カプセル化解除鍵は秘密に保つ必要があります。
func (dk *DecapsulationKey768) Bytes() []byte

// Decapsulate は、暗号文とカプセル化解除鍵から共有鍵を生成します。
// 暗号文が有効でない場合、Decapsulate はエラーを返します。
//
// 共有鍵は秘密に保つ必要があります。
func (dk *DecapsulationKey768) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)

// EncapsulationKey は、暗号文を生成するために必要な公開カプセル化鍵を
// 返します。
func (dk *DecapsulationKey768) EncapsulationKey() *EncapsulationKey768

// Encapsulator は、[DecapsulationKey768.EncapsulationKey] のように
// カプセル化鍵を返します。
//
// これは [crypto.Decapsulator] を実装します。
func (dk *DecapsulationKey768) Encapsulator() crypto.Encapsulator

var _ crypto.Decapsulator = (*DecapsulationKey768)(nil)

// EncapsulationKey768 は、対応する DecapsulationKey768 によって
// カプセル化解除される暗号文を生成するために使用される公開鍵です。
type EncapsulationKey768 struct {
	key *mlkem.EncapsulationKey768
}

// NewEncapsulationKey768 は、カプセル化鍵をエンコード形式から解析します。
// カプセル化鍵が有効でない場合、NewEncapsulationKey768 はエラーを返します。
func NewEncapsulationKey768(encapsulationKey []byte) (*EncapsulationKey768, error)

// Bytes は、カプセル化鍵をバイトスライスとして返します。
func (ek *EncapsulationKey768) Bytes() []byte

// Encapsulate は、カプセル化鍵から共有鍵と関連する暗号文を生成します。
// 安全なソースからランダムバイトを取得します。
//
// 共有鍵は秘密に保つ必要があります。
//
// テストの場合、非ランダム化されたカプセル化は [crypto/mlkem/mlkemtest]
// パッケージで提供されています。
func (ek *EncapsulationKey768) Encapsulate() (sharedKey, ciphertext []byte)

// DecapsulationKey1024 は、暗号文から共有鍵をカプセル化解除するために使用される
// 秘密鍵です。様々な事前計算値が含まれています。
type DecapsulationKey1024 struct {
	key *mlkem.DecapsulationKey1024
}

// GenerateKey1024 は、安全なソースからランダムバイトを取得して、新しい
// カプセル化解除鍵を生成します。カプセル化解除鍵は秘密に保つ必要があります。
func GenerateKey1024() (*DecapsulationKey1024, error)

// NewDecapsulationKey1024 は、"d || z" 形式の 64 バイトシードから
// カプセル化解除鍵を展開します。シードは均一にランダムである必要があります。
func NewDecapsulationKey1024(seed []byte) (*DecapsulationKey1024, error)

// Bytes は、カプセル化解除鍵を "d || z" 形式の 64 バイトシードとして返します。
//
// カプセル化解除鍵は秘密に保つ必要があります。
func (dk *DecapsulationKey1024) Bytes() []byte

// Decapsulate は、暗号文とカプセル化解除鍵から共有鍵を生成します。
// 暗号文が有効でない場合、Decapsulate はエラーを返します。
//
// 共有鍵は秘密に保つ必要があります。
func (dk *DecapsulationKey1024) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)

// EncapsulationKey は、暗号文を生成するために必要な公開カプセル化鍵を
// 返します。
func (dk *DecapsulationKey1024) EncapsulationKey() *EncapsulationKey1024

// Encapsulator は、[DecapsulationKey1024.EncapsulationKey] のように
// カプセル化鍵を返します。
//
// これは [crypto.Decapsulator] を実装します。
func (dk *DecapsulationKey1024) Encapsulator() crypto.Encapsulator

var _ crypto.Decapsulator = (*DecapsulationKey1024)(nil)

// EncapsulationKey1024 は、対応する DecapsulationKey1024 によって
// カプセル化解除される暗号文を生成するために使用される公開鍵です。
type EncapsulationKey1024 struct {
	key *mlkem.EncapsulationKey1024
}

// NewEncapsulationKey1024 は、カプセル化鍵をエンコード形式から解析します。
// カプセル化鍵が有効でない場合、NewEncapsulationKey1024 はエラーを返します。
func NewEncapsulationKey1024(encapsulationKey []byte) (*EncapsulationKey1024, error)

// Bytes は、カプセル化鍵をバイトスライスとして返します。
func (ek *EncapsulationKey1024) Bytes() []byte

// Encapsulate は、カプセル化鍵から共有鍵と関連する暗号文を生成します。
// 安全なソースからランダムバイトを取得します。
//
// 共有鍵は秘密に保つ必要があります。
//
// テストの場合、非ランダム化されたカプセル化は [crypto/mlkem/mlkemtest]
// パッケージで提供されています。
func (ek *EncapsulationKey1024) Encapsulate() (sharedKey, ciphertext []byte)
