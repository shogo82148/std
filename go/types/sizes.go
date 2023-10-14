// "go test -run=Generate -write=all"によって生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは、Sizesを実装しています。

package types

// Sizesはパッケージunsafeのサイズ決定関数を定義します。
type Sizes interface {
	Alignof(T Type) int64

	Offsetsof(fields []*Var) []int64

	Sizeof(T Type) int64
}

// StdSizesはよく使われるサイズを作成するための便利な型です。
// 以下の単純化された仮定を行います：
//
//   - 明示的なサイズの基本型（int16など）のサイズは指定されたサイズです。
//   - 文字列とインターフェースのサイズは2 * WordSizeです。
//   - スライスのサイズは3 * WordSizeです。
//   - n要素の配列のサイズは、配列要素の型のn連続フィールドの配列のサイズに対応します。
//   - 構造体のサイズは、最後のフィールドのオフセットにそのフィールドのサイズを加えたものです。
//     すべての要素型と同様に、構造体が配列で使用される場合、そのサイズはまず構造体のアライメントの倍数に揃える必要があります。
//   - その他のすべての型のサイズはWordSizeです。
//   - 配列と構造体は仕様の定義に従ってアラインされます。その他のすべての型は最大アラインメントMaxAlignで自然にアラインされます。
//
// *StdSizesはSizesを実装しています。
type StdSizes struct {
	WordSize int64
	MaxAlign int64
}

func (s *StdSizes) Alignof(T Type) (result int64)

func (s *StdSizes) Offsetsof(fields []*Var) []int64

func (s *StdSizes) Sizeof(T Type) int64

// SizesForは、コンパイラがアーキテクチャで使用するサイズを返します。
// コンパイラ/アーキテクチャの組み合わせが不明な場合、結果はnilです。
//
// コンパイラ"gc"に対応したアーキテクチャ:
// "386", "amd64", "amd64p32", "arm", "arm64", "loong64", "mips", "mipsle",
// "mips64", "mips64le", "ppc64", "ppc64le", "riscv64", "s390x", "sparc64", "wasm"。
func SizesFor(compiler, arch string) Sizes
