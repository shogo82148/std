// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkemtest は ML-KEM アルゴリズムのテスト関数を提供します。
package mlkemtest

import (
	"github.com/shogo82148/std/crypto/mlkem"
)

// Encapsulate768 は、提供されたカプセル化鍵 ek と 32 バイトのランダムネスを使用して、
// 非ランダム化された ML-KEM-768 カプセル化 (FIPS 203 の ML-KEM.Encaps_internal)
// を実装します。
//
// これは既知答えテストのためだけに使用する必要があります。
func Encapsulate768(ek *mlkem.EncapsulationKey768, random []byte) (sharedKey, ciphertext []byte, err error)

// Encapsulate1024 は、提供されたカプセル化鍵 ek と 32 バイトのランダムネスを使用して、
// 非ランダム化された ML-KEM-1024 カプセル化 (FIPS 203 の ML-KEM.Encaps_internal)
// を実装します。
//
// これは既知答えテストのためだけに使用する必要があります。
func Encapsulate1024(ek *mlkem.EncapsulationKey1024, random []byte) (sharedKey, ciphertext []byte, err error)
