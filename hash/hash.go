// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//hash パッケージはハッシュ関数のためのインターフェースを提供します。
package hash

import "github.com/shogo82148/std/io"

// Hashはすべてのハッシュ関数で実装される共通のインターフェースです。
//
// 標準ライブラリのハッシュ実装（例：hash/crc32およびcrypto/sha256）では、
// encoding.BinaryMarshalerおよびencoding.BinaryUnmarshalerインターフェースが実装されます。
// ハッシュ実装をマーシャリングすると、内部状態を保存し、後で追加の処理に使用することができます。
// 以前にハッシュに書き込まれたデータを再度書き直すことなく、利用することができます。
// ハッシュの状態には、入力の一部がそのまま含まれる場合がありますので、
// ユーザーは可能なセキュリティの影響に対応することが期待されています。
//
// 互換性：ハッシュまたは暗号パッケージへの将来の変更は、
// 以前のバージョンでエンコードされた状態を保持することを目指します。
// つまり、パッケージのリリースバージョンは、
// 以前のリリースバージョンで書かれたデータをデコードすることができるはずです。
// ただし、セキュリティ修正などの問題により、異なる結果となる場合があります。
// 背景については、Goの互換性文書を参照してください：https://golang.org/doc/go1compat
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}

// Hash32はすべての32ビットハッシュ関数によって実装される共通インターフェースです。
type Hash32 interface {
	Hash
	Sum32() uint32
}

// Hash64は、すべての64ビットハッシュ関数によって実装される共通のインタフェースです。
type Hash64 interface {
	Hash
	Sum64() uint64
}
