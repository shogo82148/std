// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージdsaは、FIPS 186-3で定義されたデジタル署名アルゴリズムを実装します。
//
// このパッケージのDSA操作は、定数時間アルゴリズムを使用して実装されていません。
//
// 廃止予定：DSAはレガシーアルゴリズムであり、パッケージcrypto/ed25519で実装されたEd25519などのモダンな代替方法を代わりに使用する必要があります。1024ビットのモジュラス（L1024N160パラメーター）を持つキーは、暗号学的に弱く、より大きなキーは一般的にサポートされていません。FIPS 186-5では、DSAが署名生成については承認されなくなっています。
package dsa

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Parametersはキーのドメインパラメータを表します。これらのパラメータは多くのキーで共有することができます。Qのビット長は8の倍数でなければなりません。
type Parameters struct {
	P, Q, G *big.Int
}

// PublicKeyはDSA公開鍵を表します。
type PublicKey struct {
	Parameters
	Y *big.Int
}

// PrivateKeyはDSAの秘密鍵を表します。
type PrivateKey struct {
	PublicKey
	X *big.Int
}

// ErrInvalidPublicKey は、このコードで使用できない公開鍵の場合に発生します。
// FIPSはDSAキーの形式に対して非常に厳格ですが、他のコードではそうでない場合もあります。
// したがって、他のコードで生成された可能性がある鍵を使用する場合は、このエラーを処理する必要があります。
var ErrInvalidPublicKey = errors.New("crypto/dsa: invalid public key")

// ParameterSizesは、DSAパラメータの受け入れ可能な素数のビット長の列挙です。 FIPS 186-3、セクション4.2を参照してください。
type ParameterSizes int

const (
	L1024N160 ParameterSizes = iota
	L2048N224
	L2048N256
	L3072N256
)

// GenerateParameters はランダムで有効なDSAパラメータをparamsに生成します。
// この関数は高速なマシンでも数秒かかる場合があります。
func GenerateParameters(params *Parameters, rand io.Reader, sizes ParameterSizes) error

// GenerateKey は公開鍵と秘密鍵のペアを生成します。PrivateKeyのパラメータはすでに有効である必要があります（GenerateParametersを参照してください）。
func GenerateKey(priv *PrivateKey, rand io.Reader) error

// Signは、ハッシュ関数を使って（より大きなメッセージのハッシュ結果である必要があります）、秘密鍵privを使って任意の長さのハッシュに署名します。署名は2つの整数のペアとして返されます。秘密鍵のセキュリティはrandのエントロピーに依存します。
// なお、FIPS 186-3のセクション4.6では、ハッシュは部分群のバイト長に切り詰める必要があると指定されています。この関数自体ではその切り詰めを実行しません。
// 注意してください。攻撃者の制御下にあるPrivateKeyを使用してSignを呼び出すことは、任意の量のCPUを必要とする場合があります。
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

// Verifyは、公開鍵pubを使用してハッシュのr、sの署名を検証します。署名が有効かどうかを報告します。
//
// FIPS 186-3のセクション4.6では、ハッシュは部分群のバイト長に切り詰める必要があると指定されています。この関数自体ではその切り詰めを行いません。
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
