// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// randパッケージは、暗号学的に安全な乱数生成器を実装しています。
package rand

import (
	"github.com/shogo82148/std/io"
)

// Readerは暗号学的に安全な乱数生成器のグローバルで共有されたインスタンスです。
// 並行使用に対して安全です。
//
//   - Linux、FreeBSD、Dragonfly、およびSolarisでは、Readerはgetrandom(2)を使用します。
//   - レガシーLinux（< 3.17）では、Readerは初回使用時に/dev/urandomを開きます。
//   - macOS、iOS、およびOpenBSDでは、Readerはarc4random_buf(3)を使用します。
//   - NetBSDでは、Readerはkern.arandom sysctlを使用します。
//   - Windowsでは、ReaderはProcessPrng APIを使用します。
//   - js/wasmでは、ReaderはWeb Crypto APIを使用します。
//   - wasip1/wasmでは、Readerはrandom_getを使用します。
//
// FIPS 140-3モードでは、出力はSP 800-90A Rev. 1
// 決定論的乱数ビット生成器（DRBG）を通過します。
var Reader io.Reader

// Readは暗号学的に安全な乱数バイトでbを埋めます。エラーを返すことはなく、
// 常にbを完全に埋めます。
//
// Readは [Reader] に対して [io.ReadFull] を呼び出し、エラーが返された場合は
// プログラムを回復不可能にクラッシュさせます。デフォルトのReaderは、
// レガシーLinuxシステム以外では決してエラーを返さないと文書化されているオペレーティングシステムAPIを使用します。
func Read(b []byte) (n int, err error)
