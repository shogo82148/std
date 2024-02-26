// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージversionは、[Goツールチェーン名の構文] のような文字列、
// つまり "go1.20"、"go1.21.0"、"go1.22rc2"、"go1.23.4-bigcorp"などの
// [Goバージョン] に対する操作を提供します。
//
// [Goバージョン]: https://go.dev/doc/toolchain#version
// [Goツールチェーン名の構文]: https://go.dev/doc/toolchain#name
package version

// Langは、バージョンxのGo言語バージョンを返します。
// もしxが有効なバージョンでない場合、Langは空の文字列を返します。
// 例えば：
//
//	Lang("go1.21rc2") = "go1.21"
//	Lang("go1.21.2") = "go1.21"
//	Lang("go1.21") = "go1.21"
//	Lang("go1") = "go1"
//	Lang("bad") = ""
//	Lang("1.21") = ""
func Lang(x string) string

// Compareは、x < y、x == y、またはx > yの場合にそれぞれ-1、0、または+1を返します。
// これはGoのバージョンとして解釈されます。
// バージョンxとyは"go"プレフィックスで始まる必要があります："go1.21"であり、"1.21"ではありません。
// 無効なバージョン、空文字列を含む、は有効なバージョンよりも小さく、
// お互いに等しいと比較されます。
// 言語バージョン"go1.21"はリリース候補および最終リリース"go1.21rc1"および"go1.21.0"よりも小さいと比較されます。
func Compare(x, y string) int

// IsValidは、バージョンxが有効かどうかを報告します。
func IsValid(x string) bool
