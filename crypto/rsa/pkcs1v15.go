// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/io"
)

// PKCS1v15DecryptOptionsは、 [crypto.Decrypter] インターフェースを使用してPKCS #1 v1.5復号化にオプションを渡すためのものです。
type PKCS1v15DecryptOptions struct {

	// SessionKeyLenは、復号化されているセッションキーの長さです。
	// ゼロでない場合、復号化中のパディングエラーにより、エラーが返される代わりに、この長さのランダムな平文が返されます。
	// これらの代替は一定の時間で発生します。
	SessionKeyLen int
}

// EncryptPKCS1v15は、与えられたメッセージをRSAとPKCS #1 v1.5のパディングスキームで暗号化します。メッセージの長さは、公開モジュラスの11バイトを引いた長さ以下である必要があります。
// ランダムパラメータはエントロピーソースとして使用され、同じメッセージを2回暗号化しても同じ暗号文が生成されないようにします。ほとんどのアプリケーションでは、[crypto/rand.Reader]をランダム関数として使用することが推奨されます。ただし、返される暗号文はランダムから読み取られたバイトに対して決定論的に依存せず、呼び出しやバージョンによって変わる場合があります。
// 注意：セッションキー以外の平文を暗号化するためにこの関数を使用することは危険です。新しいプロトコルではRSA OAEPを使用してください。
func EncryptPKCS1v15(random io.Reader, pub *PublicKey, msg []byte) ([]byte, error)

// DecryptPKCS1v15は、RSAとPKCS #1 v1.5のパディングスキームを使用して平文を復号化します。
// ランダムパラメータは旧式であり、無視されるため、nilである可能性があります。
//
// この関数がエラーを返すかどうかによって、秘密情報が漏洩する可能性があることに注意してください。
// 攻撃者がこの関数を繰り返し実行させ、各インスタンスがエラーを返すかどうかを学ぶことができれば、
// 秘密鍵を持っているかのように復号化し、署名を偽造することができます。この問題を解決する方法として、
// DecryptPKCS1v15SessionKeyを参照してください。
func DecryptPKCS1v15(random io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)

// DecryptPKCS1v15SessionKeyは、RSAとPKCS #1 v1.5のパディングスキームを使用してセッションキーを復号化します。randomパラメータは旧式であり、無視されることがあります。nilでも構いません。
// DecryptPKCS1v15SessionKeyは、暗号文の長さが正しくない場合や、暗号文が公開モジュラスよりも大きい場合にエラーを返します。それ以外の場合はエラーは返されません。パディングが有効な場合、結果の平文メッセージはkeyにコピーされます。そうでない場合、keyは変更されません。これらの代替は一定時間内に発生します。この関数の使用者は、事前にランダムなセッションキーを生成し、その値でプロトコルを継続することが意図されています。
// セッションキーが小さすぎる場合、攻撃者による総当たり攻撃が可能になる場合があります。攻撃者がそれを行えば、ランダムな値が使用されたかどうか（同じ暗号文に対しては異なる値になるため）と、したがってパディングが正しいかどうかを学ぶことができます。これはまた、この関数の目的を阻害します。少なくとも16バイトのキーを使用することで、この攻撃に対して保護されます。
// このメソッドは、RFC 3218 Section 2.3.2 で説明されているBleichenbacher選択暗号文攻撃[0]に対する保護を実装しています。これらの保護は、Bleichenbacher攻撃を非常に困難にしますが、保護はDecryptPKCS1v15SessionKeyを使用するプロトコルの残りの部分がこれらの考慮事項に基づいて設計されている場合にのみ効果的です。特に、復号化されたセッションキーを使用する後続の操作が、キーに関する情報（たとえば、静的なキーかランダムキーか）を漏洩させる場合、これらの緩和策は無効になります。このメソッドは非常に注意して使用する必要があり、通常は既存のプロトコル（TLSなど）との互換性のために絶対に必要な場合にのみ使用するべきです。
//   - [0] “Chosen Ciphertext Attacks Against Protocols Based on the RSA Encryption
//     Standard PKCS #1”, Daniel Bleichenbacher, Advances in Cryptology (Crypto'98)
//   - [1] RFC 3218, Preventing the Million Message Attack on CMS,
//     https://www.rfc-editor.org/rfc/rfc3218.html
func DecryptPKCS1v15SessionKey(random io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error

// SignPKCS1v15は、RSASSA-PKCS1-V1_5-SIGNを使用して、ハッシュされたデータの署名を計算します。ハッシュされたデータは、与えられたハッシュ関数を使用して入力メッセージをハッシュした結果でなければなりません。ハッシュがゼロの場合は、ハッシュされたデータが直接署名されます。これは相互運用性のためにのみ推奨されません。
// randomパラメータはレガシーであり、無視されることがあります。nilにすることもできます。
// この関数は決定的です。したがって、可能なメッセージのセットが小さい場合、攻撃者はメッセージから署名へのマップを作成し、署名されたメッセージを特定する可能性があります。いつものように、署名は真正さを提供し、機密性は提供しません。
func SignPKCS1v15(random io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)

// VerifyPKCS1v15は、RSA PKCS #1 v1.5の署名を検証します。
// hashedは、入力メッセージを指定されたハッシュ関数でハッシュ化した結果で、
// sigは署名です。有効な署名の場合、nilエラーが返されます。
// hashがゼロであれば、hashedは直接使用されます。
// これは相互運用性以外の用途には適していません。
//
// 入力は機密とはみなされず、タイミングのサイドチャネルを通じて、または攻撃者が入力の一部を制御している場合に漏洩する可能性があります。
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
