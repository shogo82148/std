// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto/ecdh"
)

// KEM は、HPKE 暗号スイートの 3 つのコンポーネントの 1 つである
// 鍵カプセル化メカニズムです。
type KEM interface {
	ID() uint16

	GenerateKey() (PrivateKey, error)

	NewPublicKey([]byte) (PublicKey, error)

	NewPrivateKey([]byte) (PrivateKey, error)

	DeriveKeyPair(ikm []byte) (PrivateKey, error)

	encSize() int
}

// NewKEM は与えられた KEM ID の KEM 実装を返します。
//
// アプリケーションは、ランタイム可変性が必要でない限り、
// [DHKEM] や [MLKEM768X25519] などの特定の実装を使用してください。
func NewKEM(id uint16) (KEM, error)

// PublicKey は、HPKE 暗号スイートの 3 つのコンポーネントの 1 つである
// KEM の、カプセル化鍵（つまり公開鍵）による実装です。
//
// PublicKey は通常、対応する [KEM] または [PrivateKey] のメソッドから
// 取得されます。例えば [KEM.NewPublicKey] または [PrivateKey.PublicKey] など。
type PublicKey interface {
	KEM() KEM

	Bytes() []byte

	encap() (sharedSecret, enc []byte, err error)
}

// PrivateKey は、HPKE 暗号スイートの 3 つのコンポーネントの 1 つである
// KEM の、カプセル化解除鍵（つまり秘密鍵）による実装です。
//
// PrivateKey は通常、対応する [KEM] のメソッドから取得されます。
// 例えば [KEM.GenerateKey] または [KEM.NewPrivateKey] など。
type PrivateKey interface {
	KEM() KEM

	Bytes() ([]byte, error)

	PublicKey() PublicKey

	decap(enc []byte) (sharedSecret []byte, err error)
}

// DHKEM は以下のいずれかを実装する KEM を返します。
//
//   - DHKEM(P-256, HKDF-SHA256)
//   - DHKEM(P-384, HKDF-SHA384)
//   - DHKEM(P-521, HKDF-SHA512)
//   - DHKEM(X25519, HKDF-SHA256)
//
// curve によります。
func DHKEM(curve ecdh.Curve) KEM

// NewDHKEMPublicKey は以下のいずれかを実装する PublicKey を返します。
//
//   - DHKEM(P-256, HKDF-SHA256)
//   - DHKEM(P-384, HKDF-SHA384)
//   - DHKEM(P-521, HKDF-SHA512)
//   - DHKEM(X25519, HKDF-SHA256)
//
// pub の基礎となる曲線 ([ecdh.X25519]、[ecdh.P256]、
// [ecdh.P384]、または [ecdh.P521]) によります。
//
// この関数は、既にインスタンス化された crypto/ecdh 公開鍵を持つ
// アプリケーション用です。それ以外の場合、アプリケーションは
// [DHKEM] の [KEM.NewPublicKey] メソッドを使用してください。
func NewDHKEMPublicKey(pub *ecdh.PublicKey) (PublicKey, error)

// NewDHKEMPrivateKey は以下のいずれかを実装する PrivateKey を返します。
//
//   - DHKEM(P-256, HKDF-SHA256)
//   - DHKEM(P-384, HKDF-SHA384)
//   - DHKEM(P-521, HKDF-SHA512)
//   - DHKEM(X25519, HKDF-SHA256)
//
// priv の基礎となる曲線 ([ecdh.X25519]、[ecdh.P256]、
// [ecdh.P384]、または [ecdh.P521]) によります。
//
// この関数は、既にインスタンス化された crypto/ecdh 秘密鍵、または
// [ecdh.KeyExchanger] の別の実装（例えばハードウェアキー）を持つ
// アプリケーション用です。それ以外の場合、アプリケーションは
// [DHKEM] の [KEM.NewPrivateKey] メソッドを使用してください。
func NewDHKEMPrivateKey(priv ecdh.KeyExchanger) (PrivateKey, error)
