// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ファイルASTから例となる関数を抽出します。

package doc

import (
	"github.com/shogo82148/std/go/ast"
)

// Example はテストソースファイル内で見つかった関数の例を表します。
type Example struct {
	Name        string
	Suffix      string
	Doc         string
	Code        ast.Node
	Play        *ast.File
	Comments    []*ast.CommentGroup
	Output      string
	Unordered   bool
	EmptyOutput bool
	Order       int
}

<<<<<<< HEAD
// ExamplesはtestFilesで見つかった例を、Nameフィールドでソートして返します。
// Orderフィールドには、例が出現した順序が記録されます。
// Suffixフィールドは、Examplesが直接呼び出された場合は値が入りませんが、
// NewFromFilesによってtest.goファイルで見つかった例にのみ値が入ります。
=======
// Examples returns the examples found in testFiles, sorted by Name field.
// The Order fields record the order in which the examples were encountered.
// The Suffix field is not populated when Examples is called directly, it is
// only populated by [NewFromFiles] for examples it finds in _test.go files.
>>>>>>> upstream/master
//
// プレイ可能な例は、名前が"_test"で終わるパッケージにある必要があります。
// 例は、次のいずれかの場合に「プレイ可能」です（Playフィールドが非nilである場合）：
//   - 例の関数が自己完結している場合：関数は他のパッケージの識別子（または"int"などの予め宣言された識別子）のみを参照し、テストファイルにドットインポートが含まれていない。
//   - テストファイル全体が例である場合：ファイルには正確に1つの例関数、テスト、fuzzテスト、またはベンチマーク関数が含まれ、例関数以外の少なくとも1つのトップレベル関数、型、変数、または定数宣言が存在する。
func Examples(testFiles ...*ast.File) []*Example
