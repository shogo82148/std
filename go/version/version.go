// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// パッケージversionは、[Goのバージョン] に対する操作を提供します。
//
// [Goのバージョン]: https://go.dev/doc/toolchain#version
=======
// Package version provides operations on [Go versions]
// in [Go toolchain name syntax]: strings like
// "go1.20", "go1.21.0", "go1.22rc2", and "go1.23.4-bigcorp".
//
// [Go versions]: https://go.dev/doc/toolchain#version
// [Go toolchain name syntax]: https://go.dev/doc/toolchain#name
>>>>>>> upstream/release-branch.go1.22
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

<<<<<<< HEAD
// Compareは、x < y、x == y、またはx > yの場合にそれぞれ-1、0、または+1を返します。
// これはGoのバージョンとして解釈されます。
// バージョンxとyは"go"プレフィックスで始まる必要があります："go1.21"であり、"1.21"ではありません。
// 無効なバージョン、空文字列を含む、は有効なバージョンよりも小さく、
// お互いに等しいと比較されます。
// 言語バージョン"go1.21"はリリース候補および最終リリース"go1.21rc1"および"go1.21.0"よりも小さいと比較されます。
// カスタムツールチェーンのサフィックスは比較中に無視されます：
// "go1.21.0"と"go1.21.0-bigcorp"は等しいです。
=======
// Compare returns -1, 0, or +1 depending on whether
// x < y, x == y, or x > y, interpreted as Go versions.
// The versions x and y must begin with a "go" prefix: "go1.21" not "1.21".
// Invalid versions, including the empty string, compare less than
// valid versions and equal to each other.
// The language version "go1.21" compares less than the
// release candidate and eventual releases "go1.21rc1" and "go1.21.0".
>>>>>>> upstream/release-branch.go1.22
func Compare(x, y string) int

// IsValidは、バージョンxが有効かどうかを報告します。
func IsValid(x string) bool
