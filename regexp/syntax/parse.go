// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// エラーは正規表現の解析に失敗し、問題のある表現を示します。
type Error struct {
	Code ErrorCode
	Expr string
}

func (e *Error) Error() string

// 「ErrorCode」は正規表現の解析に失敗したことを説明します。
type ErrorCode string

const (
	// 予期しないエラー
	ErrInternalError ErrorCode = "regexp/syntax: internal error"

	// パースエラー
	ErrInvalidCharClass      ErrorCode = "invalid character class"
	ErrInvalidCharRange      ErrorCode = "invalid character class range"
	ErrInvalidEscape         ErrorCode = "invalid escape sequence"
	ErrInvalidNamedCapture   ErrorCode = "invalid named capture"
	ErrInvalidPerlOp         ErrorCode = "invalid or unsupported Perl syntax"
	ErrInvalidRepeatOp       ErrorCode = "invalid nested repetition operator"
	ErrInvalidRepeatSize     ErrorCode = "invalid repeat count"
	ErrInvalidUTF8           ErrorCode = "invalid UTF-8"
	ErrMissingBracket        ErrorCode = "missing closing ]"
	ErrMissingParen          ErrorCode = "missing closing )"
	ErrMissingRepeatArgument ErrorCode = "missing argument to repetition operator"
	ErrTrailingBackslash     ErrorCode = "trailing backslash at end of expression"
	ErrUnexpectedParen       ErrorCode = "unexpected )"
	ErrNestingDepth          ErrorCode = "expression nests too deeply"
	ErrLarge                 ErrorCode = "expression too large"
)

func (e ErrorCode) String() string

// Flagsはパーサーの動作を制御し、正規表現のコンテキストに関する情報を記録します。
type Flags uint16

const (
	FoldCase Flags = 1 << iota
	Literal
	ClassNL
	DotNL
	OneLine
	NonGreedy
	PerlX
	UnicodeGroups
	WasDollar
	Simple

	MatchNL = ClassNL | DotNL

	Perl        = ClassNL | OneLine | PerlX | UnicodeGroups
	POSIX Flags = 0
)

// Parseは指定されたフラグによって制御された正規表現文字列sを解析し、正規表現の解析木を返します。構文はトップレベルのコメントに記載されています。
func Parse(s string, flags Flags) (*Regexp, error)
