// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// vet パッケージは、標準ライブラリーにベンダリングされた
// golang.org/x/tools のコピーから、cmd/vet のアナライザーを提供します。
// これにより、cmd/vet（および将来的には go test の解析ドライバーなど、
// ツールチェーンの他の部分）が、ベンダリングされた単一のコピーを共有できます。
//
// これらは、同じコピーのドライバー（[internal/analysis/unitchecker] を参照）で実行してください。
// *analysis.Analyzer 型が同一でなければならないためです。
//
// このパッケージは import のボトルネックです。どの識別子のためであっても import すると、
// vet スイートのすべてのアナライザーに依存することになります。
// これは意図的なもので、cmd/vet のようにスイート全体を実行するプログラムのためだけのものです。
// 単一のアナライザーのために import すべきではありません。
package vet

import "golang.org/x/tools/go/analysis/suite/vet"

// Suite は cmd/vet が実行するアナライザーの集合です。
var Suite = vet.Suite
