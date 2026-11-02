// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// fix パッケージは、標準ライブラリーにベンダリングされた
// golang.org/x/tools のコピーから、cmd/fix のアナライザーを提供します。
// これにより、cmd/vet のアナライザーとベンダリングされた単一のコピーを共有できます。
//
// これらは、同じコピーのドライバー（[internal/analysis/unitchecker] を参照）で実行してください。
// *analysis.Analyzer 型が同一でなければならないためです。
//
// このパッケージは import のボトルネックです。どの識別子のためであっても import すると、
// fix スイートのすべてのアナライザーに依存することになります。
// これは意図的なもので、cmd/fix のようにスイート全体を実行するプログラムのためだけのものです。
// 単一のアナライザーのために import すべきではありません。
package fix

import "golang.org/x/tools/go/analysis/suite/fix"

// Suite は cmd/fix が実行するアナライザーの集合です。
var Suite = fix.Suite
