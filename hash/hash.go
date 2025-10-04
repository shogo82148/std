// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// hash パッケージはハッシュ関数のためのインターフェースを提供します。
package hash

import "github.com/shogo82148/std/io"

// Hashはすべてのハッシュ関数で実装される共通のインターフェースです。
//
<<<<<<< HEAD
// 標準ライブラリのハッシュ実装（例：[hash/crc32] や [crypto/sha256]）は、
// [encoding.BinaryMarshaler]、および
// [encoding.BinaryUnmarshaler] インターフェースを実装しています。
// ハッシュ実装をマーシャリングすることで、その内部状態を保存し、
// 後で追加の処理に使用することができます。これにより、ハッシュに以前に書き込まれたデータを
// 再度書き込む必要がなくなります。ハッシュ状態には、入力の一部が元の形式で含まれている場合があり、
// ユーザーはそのセキュリティ上の影響を考慮する必要があります。
=======
// Hash implementations in the standard library (e.g. [hash/crc32] and
// [crypto/sha256]) implement the [encoding.BinaryMarshaler], [encoding.BinaryAppender],
// [encoding.BinaryUnmarshaler] and [Cloner] interfaces. Marshaling a hash implementation
// allows its internal state to be saved and used for additional processing
// later, without having to re-write the data previously written to the hash.
// The hash state may contain portions of the input in its original form,
// which users are expected to handle for any possible security implications.
>>>>>>> upstream/release-branch.go1.25
//
// 互換性：ハッシュまたは暗号パッケージへの将来の変更は、
// 以前のバージョンでエンコードされた状態を保持することを目指します。
// つまり、パッケージのリリースバージョンは、
// 以前のリリースバージョンで書かれたデータをデコードすることができるはずです。
// ただし、セキュリティ修正などの問題により、異なる結果となる場合があります。
// 背景については、Goの互換性文書を参照してください：https://golang.org/doc/go1compat
type Hash interface {
	io.Writer

	Sum(b []byte) []byte

	Reset()

	Size() int

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

// A Cloner is a hash function whose state can be cloned, returning a value with
// equivalent and independent state.
//
// All [Hash] implementations in the standard library implement this interface,
// unless GOFIPS140=v1.0.0 is set.
//
// If a hash can only determine at runtime if it can be cloned (e.g. if it wraps
// another hash), Clone may return an error wrapping [errors.ErrUnsupported].
// Otherwise, Clone must always return a nil error.
type Cloner interface {
	Hash
	Clone() (Cloner, error)
}

// XOF (extendable output function) is a hash function with arbitrary or unlimited output length.
type XOF interface {
	io.Writer

	io.Reader

	Reset()

	BlockSize() int
}
