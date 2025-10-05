// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ecdsaパッケージは、[FIPS 186-5]で定義されている楕円曲線デジタル署名アルゴリズムを実装します。
//
// このパッケージによって生成される署名は確定的ではありませんが、エントロピーは秘密鍵とメッセージと混合され、ランダム性源の故障の場合には同じレベルのセキュリティを実現します。
//
// 秘密鍵を扱う操作は、[elliptic.P224]、[elliptic.P256]、[elliptic.P384]、
// または [elliptic.P521] によって返される [elliptic.Curve] が使用される限り、
// 定数時間アルゴリズムを使用して実装されています。
//
// [FIPS 186-5]: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf
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

	// X, Yは公開鍵の点の座標です。
	//
	// 生の座標を変更すると無効なキーが生成される可能性があり、内部最適化が
	// 無効になる可能性があります。さらに、[big.Int] メソッドは暗号値の操作には
	// 適していません。PublicKey値をエンコード・デコードするには、[PublicKey.Bytes] と
	// [ParseUncompressedPublicKey]、または [crypto/x509.MarshalPKIXPublicKey] と
	// [crypto/x509.ParsePKIXPublicKey] を使用してください。ECDHには [crypto/ecdh] を使用し、
	// 低レベルの楕円曲線操作には filippo.io/nistec のようなサードパーティモジュールを使用してください。
	//
	// これらのフィールドはGo 1.26で非推奨になります。
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

// ParseUncompressedPublicKeyは、SEC 1, Version 2.0, Section 2.3.3に従って
// 非圧縮ポイントとしてエンコードされた公開鍵を解析します（X9.62非圧縮形式とも呼ばれます）。
// ポイントが非圧縮形式でない場合、曲線上にない場合、または無限遠点の場合はエラーを返します。
//
// curveは [elliptic.P224]、[elliptic.P256]、[elliptic.P384]、または
// [elliptic.P521] のいずれかである必要があり、そうでなければParseUncompressedPublicKeyはエラーを返します。
//
// ParseUncompressedPublicKeyは、NIST曲線に対して [ecdh.Curve.NewPublicKey] と同じ形式を受け入れますが、
// [ecdh.PublicKey] の代わりに [PublicKey] を返します。
//
// 公開鍵は通常DER（またはPEM）形式でエンコードされており、
// [crypto/x509.ParsePKIXPublicKey]（および [encoding/pem]）で解析できることに注意してください。
func ParseUncompressedPublicKey(curve elliptic.Curve, data []byte) (*PublicKey, error)

// BytesはSEC 1, Version 2.0, Section 2.3.3に従って、公開鍵を非圧縮ポイントとして
// エンコードします（X9.62非圧縮形式とも呼ばれます）。
// 公開鍵が無効な場合はエラーを返します。
//
// PublicKey.Curveは [elliptic.P224]、[elliptic.P256]、
// [elliptic.P384]、または [elliptic.P521] のいずれかである必要があり、そうでなければBytesはエラーを返します。
//
// BytesはNIST曲線に対して [ecdh.PublicKey.Bytes] と同じ形式を返します。
//
// 公開鍵は通常DER（またはPEM）形式でエンコードされており、
// [crypto/x509.MarshalPKIXPublicKey]（および [encoding/pem]）で生成できることに注意してください。
func (pub *PublicKey) Bytes() ([]byte, error)

// PrivateKeyはECDSA秘密鍵を表します。
type PrivateKey struct {
	PublicKey

	// Dは秘密スカラー値です。
	//
	// 生の値を変更すると無効なキーが生成される可能性があり、内部最適化が
	// 無効になる可能性があります。さらに、[big.Int] メソッドは暗号値の操作には
	// 適していません。PrivateKey値をエンコード・デコードするには、[PrivateKey.Bytes] と
	// [ParseRawPrivateKey]、または [crypto/x509.MarshalPKCS8PrivateKey] と
	// [crypto/x509.ParsePKCS8PrivateKey] を使用してください。ECDHには [crypto/ecdh] を使用してください。
	//
	// このフィールドはGo 1.26で非推奨になります。
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

// ParseRawPrivateKeyは、SEC 1, Version 2.0, Section 2.3.6に従って、
// 固定長のビッグエンディアン整数としてエンコードされた秘密鍵を解析します（raw形式と呼ばれることもあります）。
// 値が曲線の位数を法として縮約されていない場合、またはゼロの場合はエラーを返します。
//
// curveは[elliptic.P224]、[elliptic.P256]、[elliptic.P384]、または
// [elliptic.P521] のいずれかである必要があり、そうでなければParseRawPrivateKeyはエラーを返します。
//
// ParseRawPrivateKeyは、NIST曲線に対して [ecdh.Curve.NewPrivateKey] と同じ形式を受け入れますが、
// [ecdh.PrivateKey] の代わりに [PrivateKey] を返します。
//
// 秘密鍵は通常ASN.1またはPKCS#8形式でエンコードされており、
// [crypto/x509.ParseECPrivateKey] または
// [crypto/x509.ParsePKCS8PrivateKey]（および [encoding/pem]）で解析できることに注意してください。
func ParseRawPrivateKey(curve elliptic.Curve, data []byte) (*PrivateKey, error)

// BytesはSEC 1, Version 2.0, Section 2.3.6に従って、秘密鍵を固定長のビッグエンディアン整数として
// エンコードします（raw形式と呼ばれることもあります）。秘密鍵が無効な場合はエラーを返します。
//
// PrivateKey.Curveは[elliptic.P224]、[elliptic.P256]、
// [elliptic.P384]、または [elliptic.P521] のいずれかである必要があり、そうでなければBytesはエラーを返します。
//
// BytesはNIST曲線に対して [ecdh.PrivateKey.Bytes] と同じ形式を返します。
//
// 秘密鍵は通常ASN.1またはPKCS#8形式でエンコードされており、
// [crypto/x509.MarshalECPrivateKey] または
// [crypto/x509.MarshalPKCS8PrivateKey]（および [encoding/pem]）で生成できることに注意してください。
func (priv *PrivateKey) Bytes() ([]byte, error)

// Signは秘密鍵privを使用してハッシュ（opts.HashFunc()で大きなメッセージをハッシュ化した結果であるべき）に署名します。
// ハッシュが秘密鍵の曲線位数のビット長よりも長い場合、ハッシュはその長さに切り詰められます。
// [SignASN1] と同様に、ASN.1エンコードされた署名を返します。
//
// randがnilでない場合、署名はランダム化されます。ほとんどのアプリケーションでは
// randとして [crypto/rand.Reader] を使用する必要があります。返される署名は
// randから読み取られたバイトに決定論的に依存せず、呼び出し間やバージョン間で変更される可能性があることに注意してください。
//
// randがnilの場合、SignはRFC 6979に従って決定論的署名を生成します。
// 決定論的署名を生成する場合、opts.HashFunc()はdigestを生成するために使用された関数である必要があり、
// priv.Curveは [elliptic.P224]、[elliptic.P256]、[elliptic.P384]、または [elliptic.P521] のいずれかである必要があります。
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
