// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージecdsaは、FIPS 186-4およびSEC 1、バージョン2.0で定義されている楕円曲線デジタル署名アルゴリズムを実装しています。
//
// このパッケージによって生成される署名は確定的ではありませんが、エントロピーは秘密鍵とメッセージと混合され、ランダム性源の故障の場合には同じレベルのセキュリティを実現します。
package ecdsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/ecdh"
	"github.com/shogo82148/std/crypto/elliptic"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// PublicKeyはECDSA公開鍵を表します。
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

// ECDHはkを[ecdh.PublicKey]として返します。もしキーが[ecdh.Curve.NewPublicKey]の定義に従って無効であるか、暗号/crypto/ecdhでサポートされていないCurveであればエラーが返されます。
func (k *PublicKey) ECDH() (*ecdh.PublicKey, error)

// Equalは、pubとxが同じ値を持つかどうかを報告します。
//
<<<<<<< HEAD
// 2つのキーは、同じCurve値を持っている場合にのみ同じ値と見なされます。
// elliptic.P256()とelliptic.P256().Params()は異なる値です。
// 後者は一般的な定数時間実装ではないためです。
=======
// Two keys are only considered to have the same value if they have the same Curve value.
// Note that for example [elliptic.P256] and elliptic.P256().Params() are different
// values, as the latter is a generic not constant time implementation.
>>>>>>> upstream/master
func (pub *PublicKey) Equal(x crypto.PublicKey) bool

// PrivateKeyはECDSAの秘密鍵を表します。
type PrivateKey struct {
	PublicKey
	D *big.Int
}

<<<<<<< HEAD
// ECDHは[ecdh.PrivateKey]としてkを返します。キーが[ecdh.Curve.NewPrivateKey]の定義に従って無効である場合や、Curveがcrypto/ecdhでサポートされていない場合はエラーを返します。
=======
// ECDH returns k as a [ecdh.PrivateKey]. It returns an error if the key is
// invalid according to the definition of [ecdh.Curve.NewPrivateKey], or if the
// Curve is not supported by [crypto/ecdh].
>>>>>>> upstream/master
func (k *PrivateKey) ECDH() (*ecdh.PrivateKey, error)

// Publicはprivに対応する公開鍵を返します。
func (priv *PrivateKey) Public() crypto.PublicKey

// Equalはprivとxが同じ値を持つかどうかを報告します。
//
<<<<<<< HEAD
// Curveが比較される方法の詳細については、PublicKey.Equalを参照してください。
=======
// See [PublicKey.Equal] for details on how Curve is compared.
>>>>>>> upstream/master
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool

// privを使用してダイジェストに署名し、randからランダム性を読み取ります。opts引数
// は現在は使用されていませんが、crypto.Signerインターフェースに準拠するために、
// メッセージのダイジェストに使用されるハッシュ関数であるべきです。
//
<<<<<<< HEAD
// このメソッドはcrypto.Signerを実装しており、たとえばハードウェアモジュールにプライベートパートが
// 保持されているキーをサポートするインターフェースです。一般的にはこのパッケージのSignASN1関数を直接使用できます。
=======
// This method implements crypto.Signer, which is an interface to support keys
// where the private part is kept in, for example, a hardware module. Common
// uses can use the [SignASN1] function in this package directly.
>>>>>>> upstream/master
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)

// GenerateKeyは指定された曲線の新しいECDSA秘密鍵を生成します。
//
// ほとんどのアプリケーションではrandとして[crypto/rand.Reader]を使用する必要があります。注意点として、randから読み取ったバイトに決定的な依存関係を持たない返された鍵は、呼び出し間やバージョン間で変更される可能性があります。
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)

// SignASN1は、大きなメッセージをハッシュ化した結果であるハッシュに対して、
// 秘密鍵privを使用して署名します。ハッシュが秘密鍵の曲線順序のビット長よりも長い場合、
// ハッシュはその長さに切り詰められます。ASN.1エンコードされた署名を返します。
//
// 署名はランダム化されています。ほとんどのアプリケーションでは、[crypto/rand.Reader]
// を rand として使用する必要があります。なお、返される署名はrandから読み取られた
// バイトに対して決定論的に依存しないことに注意してください。また、呼び出しやバージョン間で変更される可能性があります。
func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error)

// VerifyASN1は公開鍵pubを使用してハッシュのASN.1エンコードされた署名sigを検証します。
// 返り値は署名が有効かどうかを示します。
func VerifyASN1(pub *PublicKey, hash, sig []byte) bool
