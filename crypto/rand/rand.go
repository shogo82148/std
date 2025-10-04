// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// randパッケージは、暗号学的に安全な乱数生成器を実装しています。
package rand

import (
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// Readerは暗号学的に安全な乱数生成器のグローバルで共有されたインスタンスです。
//
//   - Linux、FreeBSD、Dragonfly、Solarisでは、Readerは利用可能な場合はgetrandom(2)を使用し、
//     それ以外の場合は/dev/urandomを使用します。
//   - macOSとiOSでは、Readerはarc4random_buf(3)を使用します。
//   - OpenBSDとNetBSDでは、Readerはgetentropy(2)を使用します。
//   - 他のUnix系システムでは、Readerは/dev/urandomから読み取ります。
//   - Windowsでは、ReaderはProcessPrng APIを使用します。
//   - js/wasmでは、ReaderはWeb Crypto APIを使用します。
//   - wasip1/wasmでは、Readerはwasi_snapshot_preview1からrandom_getを使用します。
var Reader io.Reader

// Readはio.ReadFullを使ってReader.Readを呼び出すヘルパー関数です。
// 帰り値として、n == len(b) は err == nil の場合に限り成り立ちます。
=======
// Reader is a global, shared instance of a cryptographically
// secure random number generator. It is safe for concurrent use.
//
//   - On Linux, FreeBSD, Dragonfly, and Solaris, Reader uses getrandom(2).
//   - On legacy Linux (< 3.17), Reader opens /dev/urandom on first use.
//   - On macOS, iOS, and OpenBSD Reader, uses arc4random_buf(3).
//   - On NetBSD, Reader uses the kern.arandom sysctl.
//   - On Windows, Reader uses the ProcessPrng API.
//   - On js/wasm, Reader uses the Web Crypto API.
//   - On wasip1/wasm, Reader uses random_get.
//
// In FIPS 140-3 mode, the output passes through an SP 800-90A Rev. 1
// Deterministric Random Bit Generator (DRBG).
var Reader io.Reader

// Read fills b with cryptographically secure random bytes. It never returns an
// error, and always fills b entirely.
//
// Read calls [io.ReadFull] on [Reader] and crashes the program irrecoverably if
// an error is returned. The default Reader uses operating system APIs that are
// documented to never return an error on all but legacy Linux systems.
>>>>>>> upstream/release-branch.go1.25
func Read(b []byte) (n int, err error)
