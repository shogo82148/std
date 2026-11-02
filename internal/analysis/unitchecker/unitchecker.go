// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// unitchecker パッケージは、標準ライブラリーにベンダリングされた
// golang.org/x/tools のコピーから、unitchecker 解析ドライバーを cmd/vet に提供します。
// これで実行するアナライザーは、同じコピーのものでなければなりません
// （[internal/analysis/suite/vet] を参照）。*analysis.Analyzer 型が同一でなければならないためです。
package unitchecker

import (
	"golang.org/x/tools/go/analysis"
)

// Main は、指定されたアナライザーを実行する vet ツールの main 関数で、
// go コマンドから呼び出されます。この関数は戻りません。
func Main(analyzers ...*analysis.Analyzer)
