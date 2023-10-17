// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/io"
)

const (

	// PSSSaltLengthAuto は、PSS署名におけるソルトをできるだけ大きくし、検証時には自動で検出されるようにします。
	PSSSaltLengthAuto = 0

	// PSSSaltLengthEqualsHash は、署名に使用されるハッシュの長さと同じ長さのソルトを使用します。
	PSSSaltLengthEqualsHash = -1
)

// PSSOptionsはPSS署名の作成と検証のためのオプションを含んでいます。
type PSSOptions struct {

	// SaltLengthはPSS署名で使用されるソルトの長さを制御します。
	// それはバイト数の正数であるか、または特別なPSSSaltLengthの定数のいずれかです。
	SaltLength int

	// Hashはメッセージダイジェストを生成するために使用されるハッシュ関数です。ゼロでない場合、SignPSSに渡されたハッシュ関数を上書きします。PrivateKey.Signを使用する場合には必須です。
	Hash crypto.Hash
}

// HashFunc は opts.Hash を返します。これにより、 [PSSOptions] は [crypto.SignerOpts] を実装します。
func (opts *PSSOptions) HashFunc() crypto.Hash

// SignPSSはPSSを使用してダイジェストの署名を計算します。
//
// ダイジェストは、指定されたハッシュ関数を使用して入力メッセージをハッシュすることによって得られた結果である必要があります。
// opts引数は、nilの場合、適切なデフォルト値が使用されます。opts.Hashが設定されている場合は、hashが上書きされます。
//
// 署名は、メッセージ、キー、およびソルトサイズに応じてランダム化され、ランドからのバイトを使用します。
// ほとんどのアプリケーションでは、[crypto/rand.Reader]をrandとして使用するべきです。
func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, opts *PSSOptions) ([]byte, error)

// VerifyPSSはPSS署名を検証します。
//
// エラーがnilである場合、有効な署名です。ダイジェストは入力メッセージを与えられたハッシュ関数を使用してハッシュした結果である必要があります。opts引数はnilである場合、適切なデフォルト値が使用されます。opts.Hashは無視されます。
func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error
