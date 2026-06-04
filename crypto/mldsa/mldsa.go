// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mldsaは、[FIPS 204] で規定された耐量子計算機性ML-DSA署名方式を実装します。
//
// このパッケージは、[FIPS 140-3 Go Cryptographic Module] v1.0.0 を使用している場合は利用できません。
// その場合、[GenerateKey]、[NewPrivateKey]、[NewPublicKey]、[Verify] は
// エラーを返します。v1.26.0以降を使用している場合は利用できます。
//
// [FIPS 204]: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.204.pdf
// [FIPS 140-3 Go Cryptographic Module]: https://go.dev/doc/security/fips140
package mldsa

import "github.com/shogo82148/std/crypto"

const (
	PrivateKeySize = 32

	MLDSA44PublicKeySize = 1312
	MLDSA65PublicKeySize = 1952
	MLDSA87PublicKeySize = 2592

	MLDSA44SignatureSize = 2420
	MLDSA65SignatureSize = 3309
	MLDSA87SignatureSize = 4627
)

// Parametersは、FIPS 204で定義された固定パラメータセットの1つを表します。
//
// ほとんどのアプリケーションでは [MLDSA44] を使用するべきです。
//
// [MLDSA44]、[MLDSA65]、[MLDSA87] を複数回呼び出しても、
// それぞれ同じ値が返されます。この値は等価比較やswitch文に利用できます。
// 返される値は並行利用しても安全です。
type Parameters struct {
	name          string
	publicKeySize int
	signatureSize int
}

// MLDSA44は、FIPS 204で定義されたML-DSA-44パラメータセットを返します。
func MLDSA44() Parameters

// MLDSA65は、FIPS 204で定義されたML-DSA-65パラメータセットを返します。
func MLDSA65() Parameters

// MLDSA87は、FIPS 204で定義されたML-DSA-87パラメータセットを返します。
func MLDSA87() Parameters

// PublicKeySizeは、このパラメータセットの公開鍵サイズをバイト単位で返します。
func (params Parameters) PublicKeySize() int

// SignatureSizeは、このパラメータセットの署名サイズをバイト単位で返します。
func (params Parameters) SignatureSize() int

// Stringは、パラメータセットの名前（例: "ML-DSA-44"）を返します。
func (params Parameters) String() string

// Optionsは、ML-DSA署名の生成と検証のための追加オプションを保持します。
type Options struct {
	// Contextは、異なる目的で作成された署名を区別するために使用できます。
	// 長さは最大255バイトで、デフォルトでは空です。
	//
	// 署名の生成時と検証時には、同じcontextを使用する必要があります。
	Context string
}

// HashFuncは [crypto.SignerOpts] インターフェースを実装するために0を返します。
func (opts *Options) HashFunc() crypto.Hash
