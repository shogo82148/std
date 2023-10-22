// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package constraintはビルド制約行の解析と評価を実装しています。
// ビルド制約自体のドキュメントについては、https://golang.org/cmd/go/#hdr-Build_constraintsを参照してください。
//
// このパッケージは、オリジナルの「// +build」構文と、Go 1.17で追加された「//go:build」構文の両方を解析します。
// 「//go:build」構文の詳細については、https://golang.org/design/draft-gobuildを参照してください。
package constraint

// Exprはビルドタグの制約式です。
// 内部の具体的な型は*[AndExpr]、*[OrExpr]、*[NotExpr]、または*[TagExpr] です。
type Expr interface {
	String() string

	Eval(ok func(tag string) bool) bool

	isExpr()
}

// TagExprは、単一のタグTagのための [Expr] です。
type TagExpr struct {
	Tag string
}

func (x *TagExpr) Eval(ok func(tag string) bool) bool

func (x *TagExpr) String() string

// NotExprは式!X（Xの否定）を表します。
type NotExpr struct {
	X Expr
}

func (x *NotExpr) Eval(ok func(tag string) bool) bool

func (x *NotExpr) String() string

// AndExprは式X && Yを表します。
type AndExpr struct {
	X, Y Expr
}

func (x *AndExpr) Eval(ok func(tag string) bool) bool

func (x *AndExpr) String() string

// OrExprはX || Yを表します。
type OrExpr struct {
	X, Y Expr
}

func (x *OrExpr) Eval(ok func(tag string) bool) bool

func (x *OrExpr) String() string

// SyntaxErrorは解析されたビルド式の構文エラーを報告します。
type SyntaxError struct {
	Offset int
	Err    string
}

func (e *SyntaxError) Error() string

// Parseは、形式「//go:build ...」または「// +build ...」の単一のビルド制約行を解析し、対応するブール式を返します。
func Parse(line string) (Expr, error)

// IsGoBuildは、テキストの行が「//go:build」の制約であるかどうかを報告します。
// これは、テキストのプレフィックスのみをチェックし、式自体の解析は行いません。
func IsGoBuild(line string) bool

// IsPlusBuildはテキストの行が "// +build" 制約であるかどうかを報告します。
// これはテキストの接頭辞のみをチェックし、式そのものの解析は行いません。
func IsPlusBuild(line string) bool

// PlusBuildLinesはビルド式xに評価される「// +build」の行のシーケンスを返します。
// 式が直接「// +build」の行に変換できるほど複雑でない場合、PlusBuildLinesはエラーを返します。
func PlusBuildLines(x Expr) ([]string, error)
