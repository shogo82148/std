// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/iter"
)

// ErrDuplicateNameは、JSONトークンが重複したJSONオブジェクト名となるため、
// エンコードまたはデコードできないことを示します。
// このエラーは生成時に [SyntacticError] で直接ラップされます。
//
// 重複したJSONオブジェクトメンバーの名前は以下のように抽出できます:
//
//	err := ...
//	var serr jsontext.SyntacticError
//	if errors.As(err, &serr) && serr.Err == jsontext.ErrDuplicateName {
//		ptr := serr.JSONPointer // 重複名へのJSONポインタ
//		name := ptr.LastToken() // 重複した名前そのもの
//		...
//	}
//
// このエラーは [AllowDuplicateNames] がfalseの場合のみ返されます。
var ErrDuplicateName = errors.New("duplicate object member name")

// ErrNonStringNameは、JSONトークンが文字列でないため、
// RFC 8259の4節で要求されるJSONオブジェクト名として
// エンコードまたはデコードできないことを示します。
// このエラーは生成時に [SyntacticError] で直接ラップされます。
var ErrNonStringName = errors.New("object member name must be a string")

// Pointerは、トップレベルのJSON値のルートからの特定のJSON値を参照する
// JSONポインタ（RFC 6901）です。
//
// Pointerはスラッシュ区切りのトークンのリストであり、各トークンは
// JSONオブジェクト名またはJSON配列要素へのインデックス（10進数でエンコード）です。
// 配列インデックスとオブジェクト名（たまたま10進数でエンコードされたもの）を
// 区別することは、ポインタが参照するトップレベルJSON値の構造を知らない限り不可能です。
//
// 特定の値へのポインタの表現は常に1つだけなので、Pointer値の比較は
// 両者がまったく同じ値を指しているかどうかの判定と同じです。
type Pointer string

// IsValidは、pがRFC 6901に従った有効なJSONポインタかどうかを報告します。
// 2つの有効なポインタを連結しても有効なポインタになることに注意してください。
func (p Pointer) IsValid() bool

// Containsは、pが指すJSON値が
// pcが指すJSON値と等しいか、またはそれを含んでいるかどうかを報告します。
func (p Pointer) Contains(pc Pointer) bool

// Parentは最後のトークンを取り除き、残りのポインタを返します。
// 空のpの親は空文字列です。
func (p Pointer) Parent() Pointer

// LastTokenは、ポインタ内の最後のトークンを返します。
// 空のpの最後のトークンは空文字列です。
func (p Pointer) LastToken() string

// AppendTokenは、トークンをpの末尾に追加し、完全なポインタを返します。
func (p Pointer) AppendToken(tok string) Pointer

// Tokensは、JSONポインタ内の参照トークンに対するイテレータを返します。
// 最初のトークンから最後のトークンまで（途中で停止しない限り）順に返します。
func (p Pointer) Tokens() iter.Seq[string]
