// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ed25519はEd25519署名アルゴリズムを実装しています。詳しくは、https://ed25519.cr.yp.to/を参照してください。
//
// これらの関数はRFC 8032で定義されている「Ed25519」関数とも互換性があります。ただし、RFC 8032の定義とは異なり、このパッケージの秘密鍵表現には公開鍵の接尾辞が含まれており、同じ鍵での複数の署名操作を効率的に行うことができます。このパッケージでは、RFC 8032の秘密鍵を「seed」と呼んでいます。
package ed25519

import (
	"github.com/shogo82148/std/crypto"

	"github.com/shogo82148/std/io"
)

const (
	// PublicKeySizeは、このパッケージで使用される公開鍵のバイト単位のサイズです。
	PublicKeySize = 32
	// PrivateKeySizeは、このパッケージで使用される秘密鍵のサイズ（バイト単位）です。
	PrivateKeySize = 64
	// SignatureSizeは、このパッケージで生成および検証される署名のサイズ（バイト単位）です。
	SignatureSize = 64
	// SeedSizeは、RFC 8032で使用されるプライベートキーのシードのサイズ（バイト単位）です。
	SeedSize = 32
)

// PublicKeyはEd25519公開鍵の型です。
type PublicKey []byte

// Equalはpubとxが同じ値を持っているかどうかを報告します。
func (pub PublicKey) Equal(x crypto.PublicKey) bool

// PrivateKeyはEd25519の秘密鍵の型です。[crypto.Signer]を実装しています。
type PrivateKey []byte

// Publicはprivに対応する[PublicKey]を返します。
func (priv PrivateKey) Public() crypto.PublicKey

// Equal は priv と x が同じ値を持っているかどうかを報告します。
func (priv PrivateKey) Equal(x crypto.PrivateKey) bool

// Seedはprivに対応するプライベートキーシードを返します。RFC 8032との互換性のために提供されています。RFC 8032のプライベートキーはこのパッケージのシードに対応します。
func (priv PrivateKey) Seed() []byte

// Signは与えられたメッセージをprivで署名します。randは無視され、nilであってもかまいません。
//
// もしopts.HashFunc()が[crypto.SHA512]の場合、事前ハッシュバリアントのEd25519phが使用され、
// メッセージはSHA-512ハッシュであることが期待されます。それ以外の場合、opts.HashFunc()は[crypto.Hash](0)である必要があり、
// メッセージはハッシュされていない状態である必要があります。なぜなら、Ed25519は署名されるメッセージに対して二回のパスを行うからです。
//
// [Options]型の値、またはcrypto.Hash(0)またはcrypto.SHA512を直接使用して、
// 純粋なEd25519またはEd25519phを選択することができます。
func (priv PrivateKey) Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)

// Optionsは[PrivateKey.Sign]または[VerifyWithOptions]と一緒に使われ、Ed25519のバリアントを選択するために使用できます。
type Options struct {
	// Ed25519 の場合、ハッシュはゼロまたは Ed25519ph の場合は crypto.SHA512 になることがあります。
	Hash crypto.Hash

	// Contextが空でない場合、Ed25519ctxを選択するかEd25519phのコンテキスト文字列を提供します。長さは最大255バイトです。
	Context string
}

// HashFuncはo.Hashを返します。
func (o *Options) HashFunc() crypto.Hash

// GenerateKeyはrandからのエントロピーを使用して公開鍵/秘密鍵のペアを生成します。
// randがnilの場合、[crypto/rand.Reader]が使用されます。
//
// この関数の出力は決定論的であり、randから[SeedSize]バイトを読み取り、[NewKeyFromSeed]に渡すことと等価です。
func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error)

// NewKeyFromSeedはシードから秘密鍵を計算します。もしseedの長さが[SeedSize]でない場合、パニックを発生させます。この関数はRFC 8032との互換性のために提供されています。RFC 8032の秘密鍵はこのパッケージのシードに対応します。
func NewKeyFromSeed(seed []byte) PrivateKey

// SignはメッセージにprivateKeyで署名し、署名を返します。もしprivateKeyの長さが[PrivateKeySize]でない場合はパニックを起こします。
func Sign(privateKey PrivateKey, message []byte) []byte

// Verifyは、publicKeyによってメッセージのsigが有効な署名かどうかを検証します。
// もしlen(publicKey)が[PublicKeySize]でない場合、パニックを引き起こします。
func Verify(publicKey PublicKey, message, sig []byte) bool

// VerifyWithOptionsは、publicKeyによってメッセージのsigが有効な署名であるかどうかを報告します。有効な署名は、nilのエラーを返すことで示されます。len(publicKey)が[PublicKeySize]でない場合、パニックが発生します。
// もしopts.Hashが[crypto.SHA512]である場合、Ed25519phとして事前にハッシュされたバリアントが使用され、messageはSHA-512ハッシュであることが想定されます。それ以外の場合、opts.Hashは[crypto.Hash](0)でなければならず、メッセージはハッシュされていない状態である必要があります。なぜなら、Ed25519は署名されるメッセージを2回処理するからです。
func VerifyWithOptions(publicKey PublicKey, message, sig []byte, opts *Options) error
