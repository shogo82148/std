// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package randは、暗号学的に安全な乱数生成器を実装しています。
package rand

import "github.com/shogo82148/std/io"

// Readerは暗号学的に安全な乱数生成器のグローバルで共有されたインスタンスです。
//
<<<<<<< HEAD
// Linux、FreeBSD、Dragonfly、NetBSD、Solarisでは、Readerは利用可能な場合はgetrandom(2)を、
// それ以外の場合は/dev/urandomを使用します。
// OpenBSDとmacOSでは、Readerはgetentropy(2)を使用します。
// その他のUnix系システムでは、Readerは/dev/urandomから読み取ります。
// Windowsシステムでは、ReaderはProcessPrng APIを使用します。
// JS/Wasmでは、ReaderはWeb Crypto APIを使用します。
// WASIP1/Wasmでは、Readerはwasi_snapshot_preview1のrandom_getを使用します。
=======
//   - On Linux, FreeBSD, Dragonfly, and Solaris, Reader uses getrandom(2)
//     if available, and /dev/urandom otherwise.
//   - On macOS and iOS, Reader uses arc4random_buf(3).
//   - On OpenBSD and NetBSD, Reader uses getentropy(2).
//   - On other Unix-like systems, Reader reads from /dev/urandom.
//   - On Windows, Reader uses the ProcessPrng API.
//   - On js/wasm, Reader uses the Web Crypto API.
//   - On wasip1/wasm, Reader uses random_get from wasi_snapshot_preview1.
>>>>>>> upstream/master
var Reader io.Reader

// Readはio.ReadFullを使ってReader.Readを呼び出すヘルパー関数です。
// 帰り値として、n == len(b) は err == nil の場合に限り成り立ちます。
func Read(b []byte) (n int, err error)
