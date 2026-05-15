// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/ecdh"
)

// MLKEM768X25519 は draft-ietf-hpke-pq から MLKEM768-X25519（別名 X-Wing）を
// 実装する KEM を返します。
func MLKEM768X25519() KEM

// MLKEM768P256 は draft-ietf-hpke-pq から MLKEM768-P256 を実装する KEM を返します。
func MLKEM768P256() KEM

// MLKEM1024P384 は draft-ietf-hpke-pq から MLKEM1024-P384 を実装する KEM を返します。
func MLKEM1024P384() KEM

// NewHybridPublicKey は以下のいずれかを実装する PublicKey を返します。
//
//   - MLKEM768-X25519 (a.k.a. X-Wing)
//   - MLKEM768-P256
//   - MLKEM1024-P384
//
// draft-ietf-hpke-pq から、t の基礎となる曲線（[ecdh.X25519]、[ecdh.P256]、
// または [ecdh.P384]）と pq の型（*[mlkem.EncapsulationKey768] または
// *[mlkem.EncapsulationKey1024]）に応じています。
//
// この関数は、既にインスタンス化された crypto/ecdh および crypto/mlkem 公開鍵を
// 持つアプリケーション用です。それ以外の場合、アプリケーションは例えば
// [MLKEM768X25519] の [KEM.NewPublicKey] メソッドを使用するべきです。
func NewHybridPublicKey(pq crypto.Encapsulator, t *ecdh.PublicKey) (PublicKey, error)

// NewHybridPrivateKey は以下のいずれかを実装する PrivateKey を返します。
//
//   - MLKEM768-X25519 (a.k.a. X-Wing)
//   - MLKEM768-P256
//   - MLKEM1024-P384
//
// draft-ietf-hpke-pq から、t の基礎となる曲線（[ecdh.X25519]、[ecdh.P256]、
// または [ecdh.P384]）と pq.Encapsulator() の型（*[mlkem.EncapsulationKey768]
// または *[mlkem.EncapsulationKey1024]）に応じています。
//
// この関数は、既にインスタンス化された crypto/ecdh および crypto/mlkem 秘密鍵、
// または [ecdh.KeyExchanger] および [crypto.Decapsulator] の別の実装（例えば
// ハードウェアキー）を持つアプリケーション用です。それ以外の場合、アプリケーション
// は例えば [MLKEM768X25519] の [KEM.NewPrivateKey] メソッドを使用するべきです。
func NewHybridPrivateKey(pq crypto.Decapsulator, t ecdh.KeyExchanger) (PrivateKey, error)

// MLKEM768 は draft-ietf-hpke-pq から ML-KEM-768 を実装する KEM を返します。
func MLKEM768() KEM

// MLKEM1024 は draft-ietf-hpke-pq から ML-KEM-1024 を実装する KEM を返します。
func MLKEM1024() KEM

// NewMLKEMPublicKey は以下のいずれかを実装する KEMPublicKey を返します。
//
//   - ML-KEM-768
//   - ML-KEM-1024
//
// draft-ietf-hpke-pq から、pub の型（*[mlkem.EncapsulationKey768] または
// *[mlkem.EncapsulationKey1024]）に応じています。
//
// この関数は、既にインスタンス化された crypto/mlkem 公開鍵を持つアプリケーション用です。
// それ以外の場合、アプリケーションは例えば [MLKEM768] の [KEM.NewPublicKey]
// メソッドを使用するべきです。
func NewMLKEMPublicKey(pub crypto.Encapsulator) (PublicKey, error)

// NewMLKEMPrivateKey は以下のいずれかを実装する KEMPrivateKey を返します。
//
//   - ML-KEM-768
//   - ML-KEM-1024
//
// draft-ietf-hpke-pq から、priv.Encapsulator() の型（*[mlkem.EncapsulationKey768]
// または *[mlkem.EncapsulationKey1024]）に応じています。
//
// この関数は、既にインスタンス化された crypto/mlkem 秘密鍵を持つアプリケーション用です。
// それ以外の場合、アプリケーションは例えば [MLKEM768] の [KEM.NewPrivateKey]
// メソッドを使用するべきです。
func NewMLKEMPrivateKey(priv crypto.Decapsulator) (PrivateKey, error)
