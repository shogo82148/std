// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// ecdsaパッケージは、FIPS 186-4およびSEC 1、バージョン2.0で定義されている楕円曲線デジタル署名アルゴリズムを実装しています。
=======
// Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as
// defined in [FIPS 186-5].
>>>>>>> upstream/release-branch.go1.25
//
// このパッケージによって生成される署名は確定的ではありませんが、エントロピーは秘密鍵とメッセージと混合され、ランダム性源の故障の場合には同じレベルのセキュリティを実現します。
//
<<<<<<< HEAD
// 私有鍵を扱う操作は、[elliptic.P224]、[elliptic.P256]、[elliptic.P384]、または [elliptic.P521] によって返される
// [elliptic.Curve]が使用される限り、定数時間アルゴリズムを使用して実装されています。
=======
// Operations involving private keys are implemented using constant-time
// algorithms, as long as an [elliptic.Curve] returned by [elliptic.P224],
// [elliptic.P256], [elliptic.P384], or [elliptic.P521] is used.
//
// [FIPS 186-5]: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf
>>>>>>> upstream/release-branch.go1.25
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

	// X, Y are the coordinates of the public key point.
	//
	// Modifying the raw coordinates can produce invalid keys, and may
	// invalidate internal optimizations; moreover, [big.Int] methods are not
	// suitable for operating on cryptographic values. To encode and decode
	// PublicKey values, use [PublicKey.Bytes] and [ParseUncompressedPublicKey]
	// or [crypto/x509.MarshalPKIXPublicKey] and [crypto/x509.ParsePKIXPublicKey].
	// For ECDH, use [crypto/ecdh]. For lower-level elliptic curve operations,
	// use a third-party module like filippo.io/nistec.
	//
	// These fields will be deprecated in Go 1.26.
	X, Y *big.Int
}

// ECDHはkを[ecdh.PublicKey]として返します。もしキーが[ecdh.Curve.NewPublicKey]の定義に従って無効であるか、暗号/crypto/ecdhでサポートされていないCurveであればエラーが返されます。
func (k *PublicKey) ECDH() (*ecdh.PublicKey, error)

// Equalは、pubとxが同じ値を持つかどうかを報告します。
//
// 2つのキーは、同じCurve値を持っている場合にのみ同じ値と見なされます。
// [elliptic.P256] とelliptic.P256().Params()は異なる値です。
// 後者は一般的な定数時間実装ではないためです。
func (pub *PublicKey) Equal(x crypto.PublicKey) bool

<<<<<<< HEAD
// PrivateKeyはECDSAの秘密鍵を表します。
=======
// ParseUncompressedPublicKey parses a public key encoded as an uncompressed
// point according to SEC 1, Version 2.0, Section 2.3.3 (also known as the X9.62
// uncompressed format). It returns an error if the point is not in uncompressed
// form, is not on the curve, or is the point at infinity.
//
// curve must be one of [elliptic.P224], [elliptic.P256], [elliptic.P384], or
// [elliptic.P521], or ParseUncompressedPublicKey returns an error.
//
// ParseUncompressedPublicKey accepts the same format as
// [ecdh.Curve.NewPublicKey] does for NIST curves, but returns a [PublicKey]
// instead of an [ecdh.PublicKey].
//
// Note that public keys are more commonly encoded in DER (or PEM) format, which
// can be parsed with [crypto/x509.ParsePKIXPublicKey] (and [encoding/pem]).
func ParseUncompressedPublicKey(curve elliptic.Curve, data []byte) (*PublicKey, error)

// Bytes encodes the public key as an uncompressed point according to SEC 1,
// Version 2.0, Section 2.3.3 (also known as the X9.62 uncompressed format).
// It returns an error if the public key is invalid.
//
// PublicKey.Curve must be one of [elliptic.P224], [elliptic.P256],
// [elliptic.P384], or [elliptic.P521], or Bytes returns an error.
//
// Bytes returns the same format as [ecdh.PublicKey.Bytes] does for NIST curves.
//
// Note that public keys are more commonly encoded in DER (or PEM) format, which
// can be generated with [crypto/x509.MarshalPKIXPublicKey] (and [encoding/pem]).
func (pub *PublicKey) Bytes() ([]byte, error)

// PrivateKey represents an ECDSA private key.
>>>>>>> upstream/release-branch.go1.25
type PrivateKey struct {
	PublicKey

	// D is the private scalar value.
	//
	// Modifying the raw value can produce invalid keys, and may
	// invalidate internal optimizations; moreover, [big.Int] methods are not
	// suitable for operating on cryptographic values. To encode and decode
	// PrivateKey values, use [PrivateKey.Bytes] and [ParseRawPrivateKey] or
	// [crypto/x509.MarshalPKCS8PrivateKey] and [crypto/x509.ParsePKCS8PrivateKey].
	// For ECDH, use [crypto/ecdh].
	//
	// This field will be deprecated in Go 1.26.
	D *big.Int
}

// ECDHは [ecdh.PrivateKey] としてkを返します。キーが [ecdh.Curve.NewPrivateKey] の定義に従って無効である場合や、Curveが [crypto/ecdh] でサポートされていない場合はエラーを返します。
func (k *PrivateKey) ECDH() (*ecdh.PrivateKey, error)

// Publicはprivに対応する公開鍵を返します。
func (priv *PrivateKey) Public() crypto.PublicKey

// Equalはprivとxが同じ値を持つかどうかを報告します。
//
// Curveが比較される方法の詳細については、 [PublicKey.Equal] を参照してください。
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool

<<<<<<< HEAD
// privを使用してダイジェストに署名し、randからランダム性を読み取ります。opts引数
// は現在は使用されていませんが、crypto.Signerインターフェースに準拠するために、
// メッセージのダイジェストに使用されるハッシュ関数であるべきです。
//
// このメソッドはcrypto.Signerを実装しており、たとえばハードウェアモジュールにプライベートパートが
// 保持されているキーをサポートするインターフェースです。一般的にはこのパッケージの [SignASN1] 関数を直接使用できます。
=======
// ParseRawPrivateKey parses a private key encoded as a fixed-length big-endian
// integer, according to SEC 1, Version 2.0, Section 2.3.6 (sometimes referred
// to as the raw format). It returns an error if the value is not reduced modulo
// the curve's order, or if it's zero.
//
// curve must be one of [elliptic.P224], [elliptic.P256], [elliptic.P384], or
// [elliptic.P521], or ParseRawPrivateKey returns an error.
//
// ParseRawPrivateKey accepts the same format as [ecdh.Curve.NewPrivateKey] does
// for NIST curves, but returns a [PrivateKey] instead of an [ecdh.PrivateKey].
//
// Note that private keys are more commonly encoded in ASN.1 or PKCS#8 format,
// which can be parsed with [crypto/x509.ParseECPrivateKey] or
// [crypto/x509.ParsePKCS8PrivateKey] (and [encoding/pem]).
func ParseRawPrivateKey(curve elliptic.Curve, data []byte) (*PrivateKey, error)

// Bytes encodes the private key as a fixed-length big-endian integer according
// to SEC 1, Version 2.0, Section 2.3.6 (sometimes referred to as the raw
// format). It returns an error if the private key is invalid.
//
// PrivateKey.Curve must be one of [elliptic.P224], [elliptic.P256],
// [elliptic.P384], or [elliptic.P521], or Bytes returns an error.
//
// Bytes returns the same format as [ecdh.PrivateKey.Bytes] does for NIST curves.
//
// Note that private keys are more commonly encoded in ASN.1 or PKCS#8 format,
// which can be generated with [crypto/x509.MarshalECPrivateKey] or
// [crypto/x509.MarshalPKCS8PrivateKey] (and [encoding/pem]).
func (priv *PrivateKey) Bytes() ([]byte, error)

// Sign signs a hash (which should be the result of hashing a larger message
// with opts.HashFunc()) using the private key, priv. If the hash is longer than
// the bit-length of the private key's curve order, the hash will be truncated
// to that length. It returns the ASN.1 encoded signature, like [SignASN1].
//
// If rand is not nil, the signature is randomized. Most applications should use
// [crypto/rand.Reader] as rand. Note that the returned signature does not
// depend deterministically on the bytes read from rand, and may change between
// calls and/or between versions.
//
// If rand is nil, Sign will produce a deterministic signature according to RFC
// 6979. When producing a deterministic signature, opts.HashFunc() must be the
// function used to produce digest and priv.Curve must be one of
// [elliptic.P224], [elliptic.P256], [elliptic.P384], or [elliptic.P521].
>>>>>>> upstream/release-branch.go1.25
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
//
// 入力は機密とはみなされず、タイミングのサイドチャネルを通じて、または攻撃者が入力の一部を制御している場合に漏洩する可能性があります。
func VerifyASN1(pub *PublicKey, hash, sig []byte) bool
