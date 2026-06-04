// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build fips140v1.0

package mldsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/io"
)

// PrivateKeyはメモリ上のML-DSA秘密鍵です。
// [crypto.Signer] と、非公式に拡張された [crypto.PrivateKey] インターフェースを実装します。
//
// PrivateKeyは並行利用しても安全です。
type PrivateKey struct{}

// GenerateKeyは新しいランダムなML-DSA秘密鍵を生成します。
func GenerateKey(params Parameters) (*PrivateKey, error)

// NewPrivateKeyは、与えられたseedからML-DSA秘密鍵を復号します。
//
// seedの長さはちょうど [PrivateKeySize] バイトでなければなりません。
func NewPrivateKey(params Parameters, seed []byte) (*PrivateKey, error)

// Publicは、この秘密鍵に対応する [PublicKey] を返します。
//
// これは [crypto.Signer] インターフェースを実装します。
func (sk *PrivateKey) Public() crypto.PublicKey

// Equalは、skとxが同じ鍵かどうか（すなわち同じseedから導出されたかどうか）を報告します。
//
// xが*PrivateKeyでない場合、Equalはfalseを返します。
func (sk *PrivateKey) Equal(x crypto.PrivateKey) bool

// PublicKeyは、この秘密鍵に対応する [PublicKey] を返します。
func (sk *PrivateKey) PublicKey() *PublicKey

// Bytesは秘密鍵のseedを返します。
func (sk *PrivateKey) Bytes() []byte

// Signは、この秘密鍵を使って与えられたmessageの署名を返します。
//
// optsがnil、またはopts.HashFuncがゼロを返す場合、messageは直接署名されます。
// opts.HashFuncが [crypto.MLDSAMu] を返す場合、与えられたmessageは
// [pre-hashed μ message representative] でなければなりません。optsには *[Options]
// 型を指定できます。io.Reader引数は無視されます。
//
// [pre-hashed μ message representative]: https://www.rfc-editor.org/rfc/rfc9881.html#externalmu
func (sk *PrivateKey) Sign(_ io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)

// SignDeterministicは [PrivateKey.Sign] と同様に動作しますが、署名は決定的です。
func (sk *PrivateKey) SignDeterministic(message []byte, opts crypto.SignerOpts) (signature []byte, err error)

// PublicKeyはML-DSA公開鍵です。
// 非公式に拡張された [crypto.PublicKey] インターフェースを実装します。
//
// PublicKeyは並行利用しても安全です。
type PublicKey struct{}

// NewPublicKeyは、与えられたエンコーディングから新しいML-DSA公開鍵を作成します。
func NewPublicKey(params Parameters, encoding []byte) (*PublicKey, error)

// Bytesは公開鍵のエンコーディングを返します。
func (pk *PublicKey) Bytes() []byte

// Equalは、pkとxが同じ鍵かどうか（すなわち同じエンコーディングを持つかどうか）を報告します。
//
// xが*PublicKeyでない場合、Equalはfalseを返します。
func (pk *PublicKey) Equal(x crypto.PublicKey) bool

// Parametersは、この公開鍵に関連付けられたパラメータを返します。
func (pk *PublicKey) Parameters() Parameters

// Verifyは、signatureがpkによるmessageの有効な署名かどうかを報告します。
// optsがnilの場合、Optionsのゼロ値と同等です。
func Verify(pk *PublicKey, message []byte, signature []byte, opts *Options) error
