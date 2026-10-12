// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

const (
	// PSSSaltLengthAutoは、署名時にはPSS署名のソルトを可能な限り大きくし、
	// 検証時には自動検出するようにします。
	//
	// FIPS 140-3モードで署名する場合、ソルト長は署名で使用される
	// ハッシュ関数の長さに制限されます。
	PSSSaltLengthAuto = 0
	// PSSSaltLengthEqualsHashは、ソルト長を署名で使用されるハッシュの長さと
	// 等しくします。
	PSSSaltLengthEqualsHash = -1
)

// PSSOptionsはPSS署名の作成と検証のためのオプションを含みます。
type PSSOptions struct {
	// SaltLengthはPSS署名で使用されるソルトの長さを制御します。
	// 正のバイト数か、特別なPSSSaltLength定数のいずれかを指定できます。
	SaltLength int

	// Hashはメッセージダイジェストの生成に使用されるハッシュ関数です。
	// ゼロでない場合、SignPSSに渡されたハッシュ関数を上書きします。
	// PrivateKey.Signを使用する場合は必須です。
	Hash crypto.Hash
}

// HashFuncはopts.Hashを返すことで、[PSSOptions] が [crypto.SignerOpts] を実装するようにします。
func (opts *PSSOptions) HashFunc() crypto.Hash

// SignPSSはPSSを使用してdigestの署名を計算します。
//
// digestは、指定されたハッシュ関数を使用して入力メッセージをハッシュした結果でなければなりません。
// opts引数はnilでもよく、その場合は妥当なデフォルト値が使用されます。
// opts.Hashが設定されている場合、hashを上書きします。
//
// 署名はメッセージ、鍵、ソルトサイズに応じて、randomからのバイトを使用して
// ランダム化されます。ほとんどのアプリケーションはrandomとして
// [crypto/rand.Reader] を使用するべきです。
func SignPSS(random io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, opts *PSSOptions) ([]byte, error)

// VerifyPSSはPSS署名を検証します。
//
// 有効な署名であることは、nilエラーを返すことで示されます。digestは、
// 指定されたハッシュ関数を使用して入力メッセージをハッシュした結果でなければなりません。
// opts引数はnilでもよく、その場合は妥当なデフォルト値が使用されます。
// opts.Hashは無視されます。
//
// 入力は機密とはみなされず、タイミングサイドチャネルを通じて、
// あるいは攻撃者が入力の一部を制御している場合に漏洩する可能性があります。
func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error

// EncryptOAEPは与えられたメッセージをRSA-OAEPで暗号化します。
//
// OAEPはランダムオラクルとして使用されるハッシュ関数によってパラメーター化されます。
// 特定のメッセージの暗号化と復号には同じハッシュ関数を使用する必要があり、
// sha256.New()が妥当な選択です。
//
// randomパラメーターはエントロピー源として使用され、同じメッセージを2回暗号化しても
// 同じ暗号文にならないようにします。
// ほとんどのアプリケーションはrandomとして [crypto/rand.Reader] を使用するべきです。
//
// labelパラメーターには暗号化されない任意のデータを含めることができますが、
// メッセージに重要なコンテキストを与えます。例えば、ある公開鍵が2種類の
// メッセージの暗号化に使用される場合、異なるlabel値を使用することで、
// ある目的の暗号文が攻撃者によって別の目的に使用されないようにできます。
// 必要なければ空でもかまいません。
//
// メッセージの長さは、公開モジュラスの長さからハッシュ長の2倍を引き、
// さらに2を引いた長さ以下でなければなりません。
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)

// EncryptOAEPWithOptionsは、指定されたオプションを使用して
// 与えられたメッセージをRSA-OAEPで暗号化します。
//
// この関数は、OAEPとMGF1のハッシュを個別に指定する必要がある場合にのみ
// [EncryptOAEP] の代わりに使用するべきです。
//
// 詳細は [EncryptOAEP] を参照してください。
func EncryptOAEPWithOptions(random io.Reader, pub *PublicKey, msg []byte, opts *OAEPOptions) ([]byte, error)

// DecryptOAEPはRSA-OAEPを使用して暗号文を復号します。
//
// OAEPはランダムオラクルとして使用されるハッシュ関数によってパラメーター化されます。
// 特定のメッセージの暗号化と復号には同じハッシュ関数を使用する必要があり、
// sha256.New()が妥当な選択です。
//
// randomパラメーターはレガシーであり無視されます。nilでもかまいません。
//
// labelパラメーターは暗号化時に指定した値と一致しなければなりません。
// 詳細は [EncryptOAEP] を参照してください。
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)

// SignPKCS1v15は、RSA PKCS #1 v1.5のRSASSA-PKCS1-V1_5-SIGNを使用して
// hashedの署名を計算します。hashedは、指定されたハッシュ関数を使用して
// 入力メッセージをハッシュした結果でなければならないことに注意してください。
// hashがゼロの場合、hashedは直接署名されます。これは相互運用性のため以外には
// 推奨されません。
//
// randomパラメーターはレガシーであり無視されます。nilでもかまいません。
//
// この関数は決定的です。したがって、可能なメッセージの集合が小さい場合、
// 攻撃者はメッセージから署名へのマップを構築し、署名されたメッセージを
// 特定できる可能性があります。いつものことですが、署名は真正性を提供するものであり、
// 機密性を提供するものではありません。
func SignPKCS1v15(random io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)

// VerifyPKCS1v15はRSA PKCS #1 v1.5署名を検証します。
// hashedは指定されたハッシュ関数を使用して入力メッセージをハッシュした結果であり、
// sigは署名です。有効な署名であることは、nilエラーを返すことで示されます。
// hashがゼロの場合、hashedが直接使用されます。これは相互運用性のため以外には
// 推奨されません。
//
// 入力は機密とはみなされず、タイミングサイドチャネルを通じて、
// あるいは攻撃者が入力の一部を制御している場合に漏洩する可能性があります。
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
