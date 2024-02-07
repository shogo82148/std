// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rsaは、PKCS #1およびRFC 8017で指定されたRSA暗号化を実装します。
//
// RSAは、このパッケージで使用される単一の基本操作であり、公開鍵暗号化または公開鍵署名のいずれかを実装するために使用されます。
//
// RSAの暗号化および署名の元となる仕様は、PKCS #1であり、デフォルトでは "RSA暗号化"および "RSA署名"という用語は、通常PKCS #1バージョン1.5を指します。ただし、その仕様には欠陥があり、新しい設計はできる限りOAEPとPSSと呼ばれるバージョン2を使用するべきです。
//
// このパッケージには2つのセットのインタフェースが含まれています。より抽象的なインタフェースが不要な場合は、v1.5 / OAEPでの暗号化/復号化、およびv1.5 / PSSでの署名/検証のための関数があります。公開鍵原則に対して抽象化する必要がある場合は、PrivateKey型がcryptoパッケージのDecrypterおよびSignerインタフェースを実装します。
//
// このパッケージの操作は、一部の操作を除いて、一定の時間アルゴリズムを使用して実装されています。[GenerateKey]、[PrivateKey.Precompute]、および[PrivateKey.Validate]を除くすべての他の操作は、関連する値のビットサイズのみ漏洩し、すべての値は選択したキーサイズに依存します。
package rsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/internal/bigmod"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

<<<<<<< HEAD
=======
// A PublicKey represents the public part of an RSA key.
//
// The value of the modulus N is considered secret by this library and protected
// from leaking through timing side-channels. However, neither the value of the
// exponent E nor the precise bit size of N are similarly protected.
>>>>>>> upstream/release-branch.go1.22
type PublicKey struct {
	N *big.Int
	E int
}

// Size はバイト単位での剰余サイズを返します。この公開鍵による生の署名および暗号文のサイズは同じです。
func (pub *PublicKey) Size() int

// Equalは、pubとxが同じ値を持っているかどうかを報告します。
func (pub *PublicKey) Equal(x crypto.PublicKey) bool

// OAEPOptionsは、crypto.Decrypterインタフェースを使用してOAEP復号化にオプションを渡すためのインタフェースです。
type OAEPOptions struct {
	// Hashはマスク生成時に使用されるハッシュ関数です。
	Hash crypto.Hash

	// MGFHashはMGF1で使用されるハッシュ関数です。
	// ゼロの場合、代わりにHashが使用されます。
	MGFHash crypto.Hash

	// ラベルは、暗号化の際に使用される値と等しい任意のバイトストリングです。
	Label []byte
}

// PrivateKeyはRSAキーを表します
type PrivateKey struct {
	PublicKey
	D      *big.Int
	Primes []*big.Int

	// PrecomputedはRSAの操作を高速化するために事前計算された値を含んでいます。
	// 利用可能な場合は、PrivateKey.Precomputeを呼び出して生成されなければなりません。
	// また、変更してはいけません。
	Precomputed PrecomputedValues
}

// Publicはprivに対応する公開鍵を返します。
func (priv *PrivateKey) Public() crypto.PublicKey

// Equalはprivとxが等価な値を持つかどうかを報告します。事前計算された値は無視されます。
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool

// privという秘密鍵を使用して、乱数をrandから読み取り、署名digestを生成します。optsが *[PSSOptions] の場合、PSSアルゴリズムが使用されます。それ以外の場合は、PKCS #1 v1.5が使用されます。digestは、opts.HashFunc()を使用して入力メッセージのハッシュ値を計算した結果でなければなりません。
//
// このメソッドは [crypto.Signer] を実装しており、例えばハードウェアモジュールに保持される秘密部分をサポートするインターフェースです。一般的な使用法では、このパッケージ内のSign*関数を直接使用するべきです。
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)

// Decryptはprivで暗号文を復号化します。optsがnilまたは *[PKCS1v15DecryptOptions] 型の場合、PKCS #1 v1.5 復号化が実行されます。それ以外の場合、optsは *[OAEPOptions] 型でなければなりませんし、OAEP 復号化が行われます。
func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)

type PrecomputedValues struct {
	Dp, Dq *big.Int
	Qinv   *big.Int

	// CRTValuesは3番目以降の素数に使用されます。歴史的な偶然により、
	// 最初の2つの素数のためのCRTはPKCS #1で異なる方法で処理されますが、
	// この相互運用性は十分に重要です。
	//
	// 廃止予定:これらの値は、後方互換性のためにまだPrecomputeによって
	// 埋められていますが、使用されていません。マルチプライムRSAは非常に稀ですが、
	// 複雑さを制限するためにこのパッケージでCRTの最適化なしで実装されています。
	CRTValues []CRTValue

	n, p, q *bigmod.Modulus
}

// CRTValueには事前計算された中国剰余定理の値が含まれています。
type CRTValue struct {
	Exp   *big.Int
	Coeff *big.Int
	R     *big.Int
}

// Validateはキーに基本的な正当性チェックを行います。
// キーが正当であれば、nilを返します。それ以外の場合は、問題を説明するエラーが返されます。
func (priv *PrivateKey) Validate() error

// GenerateKeyは指定されたビットサイズのランダムなRSA秘密鍵を生成します。
//
// ほとんどのアプリケーションはrandとして[crypto/rand.Reader]を使用するべきです。注意：返される鍵は、randから読み取られたバイトに従属的に決定されず、呼び出しやバージョンによって変わる可能性があります。
func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)

// GenerateMultiPrimeKeyは指定されたビットサイズとランダムソースで、マルチプライムRSA鍵ペアを生成します。
//
// "[On the Security of Multi-prime RSA]"のテーブル1には、与えられたビットサイズの最大プライム数が示されています。
//
// 公開鍵は2つのプライムの場合と互換性がありますが、秘密鍵は異なります。したがって、特定の形式でマルチプライムの秘密鍵をエクスポートしたり、他のコードに後からインポートすることができない場合があります。
//
// このパッケージではマルチプライムRSAのCRT最適化を実装していないため、2つ以上のプライムを持つキーのパフォーマンスは悪くなります。
//
// Deprecated: 上記のセキュリティ、互換性、およびパフォーマンスの理由により、2つ以外のプライム数でこの関数を使用することはお勧めしません。代わりに [GenerateKey] を使用してください。
//
// [On the Security of Multi-prime RSA]: http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf
func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error)

// ErrMessageTooLong は、鍵のサイズに対して大きすぎるメッセージを暗号化または署名しようとした場合に返されます。 [SignPSS] を使用する場合、塩のサイズが大きすぎる場合にも返されることがあります。
var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA key size")

// EncryptOAEPはRSA-OAEPで与えられたメッセージを暗号化します。
//
// OAEPはランダムオラクルとして使用されるハッシュ関数でパラメータ化されています。
// 暗号化と復号化は同じハッシュ関数を使用する必要があります。
// sha256.New()は妥当な選択肢です。
//
// ランダムパラメータはエントロピーのソースとして使用され、同じメッセージを2回暗号化しても
// 同じ暗号文にならないようにします。
// ほとんどのアプリケーションでは、[crypto/rand.Reader]をランダムとして使用するべきです。
//
// ラベルパラメータには暗号化されない任意のデータを含めることができます。
// ただし、このデータはメッセージに重要な文脈を与えます。
// 例えば、特定の公開鍵が2つのタイプのメッセージを暗号化する場合、異なるラベル値を使用して
// 攻撃者が別の目的で暗号文を使用できないようにすることができます。
// 必要ない場合は空にしても構いません。
//
// メッセージは公開モジュラスの長さからハッシュの2倍の長さを引いた値より長くすることはできません。
// さらに2を引いた長さ以下である必要があります。
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)

// ErrDecryptionはメッセージの復号に失敗したことを表します。
// 適応攻撃を避けるため、故意に曖昧さを持たせています。
var ErrDecryption = errors.New("crypto/rsa: decryption error")

// ErrVerificationは署名を検証できなかったことを表します。
// 自己適応攻撃を避けるために、意図的にあいまいです。
var ErrVerification = errors.New("crypto/rsa: verification error")

// Precomputeは未来の秘密鍵操作を高速化するためのいくつかの計算を実行します。
func (priv *PrivateKey) Precompute()

// DecryptOAEPはRSA-OAEPを使用して暗号文を復号化します。
//
// OAEPはランダムオラクルとして使用されるハッシュ関数でパラメータ化されます。
// 特定のメッセージの暗号化および復号化は、同じハッシュ関数を使用する必要があります。
// sha256.New()が妥当な選択肢です。
//
// ランダムパラメータは旧式で無視され、nilであることができます。
//
// ラベルパラメータは暗号化時に指定した値と一致する必要があります。
// 詳細については、 [EncryptOAEP] を参照してください。
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
