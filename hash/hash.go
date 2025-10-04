// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// hash パッケージはハッシュ関数のためのインターフェースを提供します。
package hash

import "github.com/shogo82148/std/io"

// Hashはすべてのハッシュ関数で実装される共通のインターフェースです。
//
// 標準ライブラリのハッシュ実装（例：[hash/crc32] や
// [crypto/sha256]）は、[encoding.BinaryMarshaler]、[encoding.BinaryAppender]、
// [encoding.BinaryUnmarshaler]、および [Cloner] インターフェースを実装しています。ハッシュ実装の
// マーシャリングにより、その内部状態を保存し、以前にハッシュに書き込まれたデータを
// 再書き込みすることなく、後で追加処理に使用することができます。
// ハッシュ状態には、入力の一部が元の形で含まれている可能性があり、
// ユーザーは可能なセキュリティ上の影響を処理することが期待されます。
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

// Clonerは状態をクローンできるハッシュ関数で、
// 同等で独立した状態を持つ値を返します。
//
// 標準ライブラリのすべての [Hash] 実装は、GOFIPS140=v1.0.0が設定されていない限り、
// このインターフェースを実装しています。
//
// ハッシュが実行時にのみクローン可能かどうかを判断できる場合（例：別のハッシュをラップしている場合）、
// CloneはErrUnsupportedをラップしたエラーを返すことがあります。
// それ以外の場合、Cloneは常にnilエラーを返さなければなりません。
type Cloner interface {
	Hash
	Clone() (Cloner, error)
}

// XOF (extendable output function) は任意または無制限の出力長を持つハッシュ関数です。
type XOF interface {
	io.Writer

	io.Reader

	Reset()

	BlockSize() int
}
