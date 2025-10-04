// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ecdhパッケージはNIST曲線とCurve25519上での楕円曲線ディフィー・ヘルマンを実装します。
package ecdh

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/internal/boring"
	"github.com/shogo82148/std/crypto/internal/fips140/ecdh"
	"github.com/shogo82148/std/io"
)

type Curve interface {
	GenerateKey(rand io.Reader) (*PrivateKey, error)

	NewPrivateKey(key []byte) (*PrivateKey, error)

	NewPublicKey(key []byte) (*PublicKey, error)

	ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error)
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
	fips      *ecdh.PublicKey
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
	publicKey  *PublicKey
	boring     *boring.PrivateKeyECDH
	fips       *ecdh.PrivateKey
}

// ECDHはECDH交換を実行し、共有秘密を返します。[PrivateKey]
// と [PublicKey] は同じ曲線を使用する必要があります。
//
// NIST曲線の場合、これはSEC 1, Version 2.0,
// Section 3.3.1で指定されたECDHを実行し、SEC 1,
// Version 2.0, Section 2.3.5に従ってエンコードされたx座標を返します。結果は決して無限遠点にはなりません。
// これは、NIST SP 800-56A Rev. 3, Section 6.1.2.2で指定された
// Ephemeral Unified Modelスキームの共有秘密計算としても知られています。
//
// [X25519] の場合、これはRFC 7748 Section 6.1 で指定されたようにECDHを実行します。結果が全て0値の場合、ECDHはエラーを返します。
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
