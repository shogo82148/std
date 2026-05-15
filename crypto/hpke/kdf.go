// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

// KDF は HPKE 暗号スイートの 3 つのコンポーネントの 1 つで、鍵導出を
// 実装します。
type KDF interface {
	ID() uint16
	oneStage() bool
	size() int
	labeledDerive(suiteID, inputKey []byte, label string, context []byte, length uint16) ([]byte, error)
	labeledExtract(suiteID, salt []byte, label string, inputKey []byte) ([]byte, error)
	labeledExpand(suiteID, randomKey []byte, label string, info []byte, length uint16) ([]byte, error)
}

// NewKDF は与えられた KDF ID の KDF 実装を返します。
//
// アプリケーションは、ランタイム可変性が必要でない限り、
// [HKDFSHA256] などの特定の実装を使用してください。
func NewKDF(id uint16) (KDF, error)

// HKDFSHA256 は HKDF-SHA256 KDF 実装を返します。
func HKDFSHA256() KDF

// HKDFSHA384 は HKDF-SHA384 KDF 実装を返します。
func HKDFSHA384() KDF

// HKDFSHA512 は HKDF-SHA512 KDF 実装を返します。
func HKDFSHA512() KDF

// SHAKE128 は SHAKE128 KDF 実装を返します。
func SHAKE128() KDF

// SHAKE256 は SHAKE256 KDF 実装を返します。
func SHAKE256() KDF
