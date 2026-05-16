// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cryptotestは決定論的な乱数ソースのテストを提供します。
package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// SetGlobalRandomは、テストtの実行中にグローバルで決定論的な
// 暗号学的乱数ソースを設定します。これはcrypto/randと、
// crypto/... パッケージ内の暗号学的乱数のすべての暗黙的ソースに影響します。
//
// SetGlobalRandomは、同じテスト内で複数回呼び出して、
// 乱数ストリームをリセットしたりseedを変更したりできます。
//
// SetGlobalRandomはプロセス全体に影響するため、
// 並列テストや並列な先祖テストを持つテストでは使用できません。
//
// 暗号アルゴリズムが乱数をどのように使うかは一般に仕様化されておらず、
// 時間とともに変わる可能性がある点に注意してください。したがって、
// テストが暗号関数の特定の出力を期待している場合、SetGlobalRandomを
// 使用していても将来失敗する可能性があります。
//
// SetGlobalRandomは、Go Cryptographic Module v1.0.0 を対象にビルドする場合
// サポートされません（つまり [crypto/fips140.Version] が "v1.0.0" を
// 返す場合）。
func SetGlobalRandom(t *testing.T, seed uint64)
