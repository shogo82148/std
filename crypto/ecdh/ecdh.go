// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージecdhはNIST曲線とCurve25519上での楕円曲線ディフィー・ヘルマンを実装します。
package ecdh

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/internal/boring"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

type Curve interface {
	// GenerateKey generates a random PrivateKey.
	//
	// Most applications should use [crypto/rand.Reader] as rand. Note that the
	// returned key does not depend deterministically on the bytes read from rand,
	// and may change between calls and/or between versions.
	GenerateKey(rand io.Reader) (*PrivateKey, error)

	// NewPrivateKey checks that key is valid and returns a PrivateKey.
	//
	// For NIST curves, this follows SEC 1, Version 2.0, Section 2.3.6, which
	// amounts to decoding the bytes as a fixed length big endian integer and
	// checking that the result is lower than the order of the curve. The zero
	// private key is also rejected, as the encoding of the corresponding public
	// key would be irregular.
	//
	// For X25519, this only checks the scalar length.
	NewPrivateKey(key []byte) (*PrivateKey, error)

	// NewPublicKey checks that key is valid and returns a PublicKey.
	//
	// For NIST curves, this decodes an uncompressed point according to SEC 1,
	// Version 2.0, Section 2.3.4. Compressed encodings and the point at
	// infinity are rejected.
	//
	// For X25519, this only checks the u-coordinate length. Adversarially
	// selected public keys can cause ECDH to return an error.
	NewPublicKey(key []byte) (*PublicKey, error)

	// ecdh performs an ECDH exchange and returns the shared secret. It's exposed
	// as the PrivateKey.ECDH method.
	//
	// The private method also allow us to expand the ECDH interface with more
	// methods in the future without breaking backwards compatibility.
	ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error)

	// privateKeyToPublicKey converts a PrivateKey to a PublicKey. It's exposed
	// as the PrivateKey.PublicKey method.
	//
	// This method always succeeds: for X25519, the zero key can't be
	// constructed due to clamping; for NIST curves, it is rejected by
	// NewPrivateKey.
	privateKeyToPublicKey(*PrivateKey) *PublicKey
}

// PublicKeyは通常、ワイヤ経由で送信されるECDHの共有キーです。
//
// これらのキーは[crypto/x509.ParsePKIXPublicKey]で解析可能であり、
// [crypto/x509.MarshalPKIXPublicKey]を使用してエンコードすることもできます。
// NIST曲線の場合、解析後に[crypto/ecdsa.PublicKey.ECDH]で変換する必要があります。
type PublicKey struct {
	curve     Curve
	publicKey []byte
	boring    *boring.PublicKeyECDH
}

// Bytesは公開鍵のエンコードのコピーを返します。
func (k *PublicKey) Bytes() []byte

// Equalは、xがkと同じ公開鍵を表すかどうかを返します。
//
// 注意：異なるエンコーディングを持つ同等の公開鍵があり、このチェックではfalseを返しますが、ECDHの入力として同じ結果を示す可能性があります。
//
// キーのタイプとその曲線が一致している限り、このチェックは一定時間で実行されます。
func (k *PublicKey) Equal(x crypto.PublicKey) bool

func (k *PublicKey) Curve() Curve

// PrivateKeyは通常秘密に保持されるECDHの秘密鍵です。
//
// これらの鍵は[crypto/x509.ParsePKCS8PrivateKey]でパースでき、[crypto/x509.MarshalPKCS8PrivateKey]
// でエンコードすることができます。NIST曲線の場合、パース後に[crypto/ecdsa.PrivateKey.ECDH]で変換する必要があります。
type PrivateKey struct {
	curve      Curve
	privateKey []byte
	boring     *boring.PrivateKeyECDH

	// publicKeyは、公開鍵を公開鍵の一度セットしています。これにより、スカラー乗算を行わずにNewPrivateKeyで秘密鍵を読み込むことができます。
	publicKey     *PublicKey
	publicKeyOnce sync.Once
}

// ECDH performs an ECDH exchange and returns the shared secret. The [PrivateKey]
// and [PublicKey] must use the same curve.
//
// NIST曲線の場合、これはSEC 1バージョン2.0セクション3.3.1で指定されたようにECDHを実行し、SEC 1バージョン2.0セクション2.3.5に従ってエンコードされたx座標を返します。結果は決して無限遠点ではありません。
//
// [X25519] の場合、これはRFC 7748セクション6.1で指定されたようにECDHを実行します。結果が全て0値の場合、ECDHはエラーを返します。
func (k *PrivateKey) ECDH(remote *PublicKey) ([]byte, error)

// Bytesは、プライベートキーのエンコーディングのコピーを返します。
func (k *PrivateKey) Bytes() []byte

// Equalは、xがkと同じ秘密鍵を表しているかどうかを返します。
//
// ただし、異なるエンコーディングを持つ等価な秘密鍵が存在する場合があることに注意してください。
// この場合、このチェックではfalseが返されますが、 [ECDH] の入力としては同じように機能します。
//
// このチェックは、キーのタイプと曲線が一致している限り、一定の時間で実行されます。
func (k *PrivateKey) Equal(x crypto.PrivateKey) bool

func (k *PrivateKey) Curve() Curve

func (k *PrivateKey) PublicKey() *PublicKey

// Publicは、すべての標準ライブラリの非公開キーの暗黙のインターフェースを実装します。 [crypto.PrivateKey] のドキュメントを参照してください。
func (k *PrivateKey) Public() crypto.PublicKey
