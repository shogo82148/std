// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.runtimesecret

// secretパッケージは、前方秘匿性のために、ユーザープログラムからは見えないメモリを
// ゼロクリアするためのヘルパー関数を含みます。詳細については
// https://en.wikipedia.org/wiki/Forward_secrecy を参照してください。
//
// このパッケージ（runtime/secret）は実験的なものであり、
// Go 1の互換性の約束の対象ではありません。
// GOEXPERIMENT=runtimesecret環境変数を設定してビルドした場合にのみ存在します。
package secret
