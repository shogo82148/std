// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージcryptoは一般的な暗号定数を収集します。
package crypto

import (
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

// Hashは別のパッケージで実装されている暗号ハッシュ関数を識別します。
type Hash uint

<<<<<<< HEAD
// HashFunc は単に h の値を返すだけであり、Hash が SignerOpts を実装することを保証します。
=======
// HashFunc simply returns the value of h so that [Hash] implements [SignerOpts].
>>>>>>> upstream/master
func (h Hash) HashFunc() Hash

func (h Hash) String() string

const (
	MD4 Hash = 1 + iota
	MD5
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	MD5SHA1
	RIPEMD160
	SHA3_224
	SHA3_256
	SHA3_384
	SHA3_512
	SHA512_224
	SHA512_256
	BLAKE2s_256
	BLAKE2b_256
	BLAKE2b_384
	BLAKE2b_512
)

// Sizeは与えられたハッシュ関数から生成されるダイジェストの長さ（バイト単位）を返します。
// この関数は、対象のハッシュ関数がプログラムにリンクされている必要はありません。
func (h Hash) Size() int

// Newは指定されたハッシュ関数を計算する新しいhash.Hashを返します。ハッシュ関数がバイナリにリンクされていない場合、Newはパニックを発生させます。
func (h Hash) New() hash.Hash

// Availableは、与えられたハッシュ関数がバイナリにリンクされているかどうかを示します。
func (h Hash) Available() bool

// RegisterHash は与えられたハッシュ関数の新しいインスタンスを返す関数を登録します。
// これはハッシュ関数を実装するパッケージの init 関数から呼び出されることを意図しています。
func RegisterHash(h Hash, f func() hash.Hash)

// PublicKeyは未指定のアルゴリズムを使用して公開鍵を表します。
//
// このタイプは、後方互換性のための空のインターフェースですが、
// 標準ライブラリのすべての公開鍵タイプは、以下のインターフェースを実装しています。
//
//	interface{
//	    Equal(x crypto.PublicKey) bool
//	}
//
// アプリケーション内での型安全性向上のために使用することができます。
type PublicKey any

// PrivateKeyは不特定のアルゴリズムを使用して秘密鍵を表します。
//
// この型は、後方互換性のために空のインターフェースとして定義されていますが、
// 標準ライブラリのすべての秘密鍵タイプは以下のインターフェースを実装します。
//
//	interface{
//	    Public() crypto.PublicKey
//	    Equal(x crypto.PrivateKey) bool
//	}
//
<<<<<<< HEAD
// また、SignerやDecrypterなどの特定の目的のインターフェースも実装しており、
// アプリケーション内で型の安全性を向上させるために使用できます。
=======
// as well as purpose-specific interfaces such as [Signer] and [Decrypter], which
// can be used for increased type safety within applications.
>>>>>>> upstream/master
type PrivateKey any

// Signerは、署名操作に使用される不透明な秘密鍵のインターフェースです。たとえば、ハードウェアモジュールに保存されているRSA鍵などがあります。
type Signer interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Sign signs digest with the private key, possibly using entropy from
	// rand. For an RSA key, the resulting signature should be either a
	// PKCS #1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA
	// key, it should be a DER-serialised, ASN.1 signature structure.
	//
	// Hash implements the SignerOpts interface and, in most cases, one can
	// simply pass in the hash function used as opts. Sign may also attempt
	// to type assert opts to other types in order to obtain algorithm
	// specific values. See the documentation in each package for details.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest) and the hash function (as opts) to Sign.
	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}

<<<<<<< HEAD
// SignerOptsはSignerでの署名に対するオプションを含んでいます。
=======
// SignerOpts contains options for signing with a [Signer].
>>>>>>> upstream/master
type SignerOpts interface {
	// HashFunc returns an identifier for the hash function used to produce
	// the message passed to Signer.Sign, or else zero to indicate that no
	// hashing was done.
	HashFunc() Hash
}

// Decrypterは、非対称復号化操作に使用できる不透明な秘密鍵のインターフェースです。例えば、ハードウェアモジュールに保管されるRSA鍵があります。
type Decrypter interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Decrypt decrypts msg. The opts argument should be appropriate for
	// the primitive used. See the documentation in each implementation for
	// details.
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts any
