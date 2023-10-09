// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constraint

// GoVersionは与えられたビルド条件によって暗黙的に指定される最小のGoバージョンを返します。
// 条件式にGoバージョンタグがない場合、GoVersionは空の文字列を返します。
//
// 例:
//
//	GoVersion(linux && go1.22) = "go1.22"
//	GoVersion((linux && go1.22) || (windows && go1.20)) = "go1.20" => go1.20
//	GoVersion(linux) = ""
//	GoVersion(linux || (windows && go1.22)) = ""
//	GoVersion(!go1.22) = ""
//
// GoVersionは、任意のタグまたは否定されたタグが独立してtrueである可能性があると仮定しています。
// そのため、解析はSATソルバーなしで純粋に構造的に行われます。
//「不可能」とされる部分式は結果に影響する可能性があるためです。
//
// 例:
//
//	GoVersion((linux && !linux && go1.20) || go1.21) = "go1.20"
func GoVersion(x Expr) string
