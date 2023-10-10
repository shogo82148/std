// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// 正規表現（RegExp）は正規表現構文木のノードです。
type Regexp struct {
	Op       Op
	Flags    Flags
	Sub      []*Regexp
	Sub0     [1]*Regexp
	Rune     []rune
	Rune0    [2]rune
	Min, Max int
	Cap      int
	Name     string
}

// Opは単一の正規表現演算子です。
type Op uint8

const (
	OpNoMatch Op = 1 + iota
	OpEmptyMatch
	OpLiteral
	OpCharClass
	OpAnyCharNotNL
	OpAnyChar
	OpBeginLine
	OpEndLine
	OpBeginText
	OpEndText
	OpWordBoundary
	OpNoWordBoundary
	OpCapture
	OpStar
	OpPlus
	OpQuest
	OpRepeat
	OpConcat
	OpAlternate
)

// Equalはxとyが同じ構造を持っているかどうかを報告します。
func (x *Regexp) Equal(y *Regexp) bool

func (re *Regexp) String() string

// MaxCapは正規表現を辿って最大のキャプチャーインデックスを見つけます。
func (re *Regexp) MaxCap() int

// CapNamesは正規表現を走査してキャプチャグループの名前を見つけます。
func (re *Regexp) CapNames() []string
