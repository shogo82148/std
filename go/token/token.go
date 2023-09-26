// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
package token

// Token is the set of lexical tokens of the Go programming language.
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT
	INT
	FLOAT
	IMAG
	CHAR
	STRING

	// Operators and delimiters
	ADD
	SUB
	MUL
	QUO
	REM

	AND
	OR
	XOR
	SHL
	SHR
	AND_NOT

	ADD_ASSIGN
	SUB_ASSIGN
	MUL_ASSIGN
	QUO_ASSIGN
	REM_ASSIGN

	AND_ASSIGN
	OR_ASSIGN
	XOR_ASSIGN
	SHL_ASSIGN
	SHR_ASSIGN
	AND_NOT_ASSIGN

	LAND
	LOR
	ARROW
	INC
	DEC

	EQL
	LSS
	GTR
	ASSIGN
	NOT

	NEQ
	LEQ
	GEQ
	DEFINE
	ELLIPSIS

	LPAREN
	LBRACK
	LBRACE
	COMMA
	PERIOD

	RPAREN
	RBRACK
	RBRACE
	SEMICOLON
	COLON

	// Keywords
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR
)

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
func (tok Token) String() string

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence corresponds serves as "catch-all" precedence for
// selector, indexing, and other operator and delimiter tokens.
const (
	LowestPrec  = 0
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
func (op Token) Precedence() int

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Token

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
func (tok Token) IsLiteral() bool

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (tok Token) IsOperator() bool

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok Token) IsKeyword() bool
