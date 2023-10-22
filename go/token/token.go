// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージtokenは、Goプログラミング言語の字句トークンとトークンに対する基本的な操作（印刷、述語）を定義する定数を表します。
package token

// TokenはGoプログラミング言語の字句トークンの集合です。
type Token int

// トークンのリスト。
const (
	// 特別なトークン
	ILLEGAL Token = iota
	EOF
	COMMENT

	// 識別子と基本型のリテラル
	// (これらのトークンはリテラルのクラスを表します)
	IDENT
	INT
	FLOAT
	IMAG
	CHAR
	STRING

	// オペレータと区切り文字
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

	// キーワード
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

	// 追加のトークン、特別な方法で処理される
	TILDE
)

<<<<<<< HEAD
// Stringはtokに対応する文字列を返します。
// 演算子、区切り文字、キーワードの場合、文字列は実際のトークン文字列です（たとえば、トークンADDの場合、文字列は"+"です）。
// それ以外のすべてのトークンに対して、文字列はトークンの定数名に対応します（たとえば、トークンIDENTの場合、文字列は"IDENT"です）。
=======
// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token [ADD], the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token [IDENT], the string is "IDENT").
>>>>>>> upstream/master
func (tok Token) String() string

// 優先順位ベースの式の解析のための定数のセット。
// 非演算子は最低の優先度を持ち、1から始まる演算子が続きます。
// 最高の優先度はセレクタ、インデックス、その他の演算子や区切り記号トークンのための「キャッチオール」優先度として機能します。
const (
	LowestPrec  = 0
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
// Precedenceは、バイナリ演算子opの演算子の優先度を返します。opがバイナリ演算子でない場合、結果はLowestPrecedenceになります。
func (op Token) Precedence() int

<<<<<<< HEAD
// Lookupは識別子をキーワードトークンまたはIDENT（キーワードでない場合）にマップします。
=======
// Lookup maps an identifier to its keyword token or [IDENT] (if not a keyword).
>>>>>>> upstream/master
func Lookup(ident string) Token

// IsLiteral は、識別子と基本型のリテラルに対応するトークンに対して true を返します。それ以外の場合は、false を返します。
func (tok Token) IsLiteral() bool

// IsOperatorはオペレーターや区切り記号に対応するトークンに対してtrueを返し、
// それ以外の場合はfalseを返します。
func (tok Token) IsOperator() bool

// IsKeywordはキーワードに対応するトークンに対してtrueを返し、それ以外の場合はfalseを返します。
func (tok Token) IsKeyword() bool

// IsExported は、name が大文字で始まるかどうかを報告します。
func IsExported(name string) bool

// IsKeywordは、nameがGoのキーワード（"func"や"return"など）であるかどうかを報告します。
func IsKeyword(name string) bool

// IsIdentifierは、nameがGoの識別子であるかどうかを報告します。つまり、
// 最初の文字が数字でない、文字、数字、アンダースコアで構成された空でない文字列です。キーワードは識別子ではありません。
func IsIdentifier(name string) bool
