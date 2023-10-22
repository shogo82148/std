// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scannerはGo言語のソーステキストのためのスキャナを実装します。
// ソースとして[]byteを受け取り、Scanメソッドへの繰り返し呼び出しを通じてトークン化します。
package scanner

import (
	"github.com/shogo82148/std/go/token"
)

<<<<<<< HEAD
// ErrorHandlerはScanner.Initに提供することができます。構文エラーが発生し、ハンドラがインストールされている場合、ハンドラは位置とエラーメッセージとともに呼び出されます。位置は問題のあるトークンの始まりを指します。
type ErrorHandler func(pos token.Position, msg string)

// スキャナは、指定されたテキストを処理する間、スキャンの内部状態を保持します。それは他のデータ構造の一部として割り当てられることもできますが、使用前にInitによって初期化される必要があります。
=======
// An ErrorHandler may be provided to [Scanner.Init]. If a syntax error is
// encountered and a handler was installed, the handler is called with a
// position and an error message. The position points to the beginning of
// the offending token.
type ErrorHandler func(pos token.Position, msg string)

// A Scanner holds the scanner's internal state while processing
// a given text. It can be allocated as part of another data
// structure but must be initialized via [Scanner.Init] before use.
>>>>>>> upstream/master
type Scanner struct {
	// 変更不可な状態
	file *token.File
	dir  string
	src  []byte
	err  ErrorHandler
	mode Mode

	// スキャン状態
	ch         rune
	offset     int
	rdOffset   int
	lineOffset int
	insertSemi bool
	nlPos      token.Pos

	// 公開される状態 - 変更しても問題ありません
	ErrorCount int
}

// モード値はフラグの集合体です（または0）。
// これらはスキャナの動作を制御します。
type Mode uint

const (
	ScanComments Mode = 1 << iota
)

// Init関数は、スキャナsを準備してテキストsrcをトークンに分割することができるように、
// スキャナをsrcの先頭に位置付けします。スキャナは位置情報のためにファイルセットファイルを使用し、
// 各行に対して行情報を追加します。同じファイルを再スキャンする際には、既に存在する行情報は無視されるため、
// 同じファイルを再使用しても問題ありません。もしファイルのサイズがsrcのサイズと一致しない場合は、
// Initはパニックを引き起こします。
//
<<<<<<< HEAD
// Scanの呼び出しは、構文エラーが発生した場合にエラーハンドラerrを呼び出し、errがnilでない場合にはエラーハンドラerrを呼び出します。
// また、エラーが発生する毎にスキャナのErrorCountフィールドが1増加します。モードパラメーターはコメントの扱い方を決定します。
=======
// Calls to [Scanner.Scan] will invoke the error handler err if they encounter a
// syntax error and err is not nil. Also, for each error encountered,
// the [Scanner] field ErrorCount is incremented by one. The mode parameter
// determines how comments are handled.
>>>>>>> upstream/master
//
// Initはファイルの最初の文字にエラーがある場合、errを呼び出す場合があります。
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler, mode Mode)

<<<<<<< HEAD
// Scanは次のトークンをスキャンして、トークンの位置、トークン、およびリテラル文字列（適用可能な場合）を返します。ソースの終了はtoken.EOFで示されます。
// 返されるトークンがリテラル（token.IDENT、token.INT、token.FLOAT、token.IMAG、token.CHAR、token.STRING）またはtoken.COMMENTである場合、リテラル文字列は対応する値を持ちます。
// 返されるトークンがキーワードの場合、リテラル文字列はキーワードです。
// 返されるトークンがtoken.SEMICOLONの場合、対応するリテラル文字列は、ソースにセミコロンが存在する場合は";"であり、改行またはEOFの場合は"\n"です。
// 返されるトークンがtoken.ILLEGALの場合、リテラル文字列は違反している文字です。
// その他の場合、Scanは空のリテラル文字列を返します。
// より許容度の高い解析のために、構文エラーが発生した場合でも、可能な限り有効なトークンを返します。したがって、結果のトークンのシーケンスに違法なトークンが含まれていない場合でも、クライアントはエラーが発生したとは推測できません。代わりに、スキャナーのErrorCountまたはエラーハンドラーの呼び出し回数を確認する必要があります（インストールされている場合）。
// ScanはInitでファイルセットに追加されたファイルに行情報を追加します。トークンの位置はそのファイルとファイルセットに対して相対的です。
=======
// Scan scans the next token and returns the token position, the token,
// and its literal string if applicable. The source end is indicated by
// [token.EOF].
//
// If the returned token is a literal ([token.IDENT], [token.INT], [token.FLOAT],
// [token.IMAG], [token.CHAR], [token.STRING]) or [token.COMMENT], the literal string
// has the corresponding value.
//
// If the returned token is a keyword, the literal string is the keyword.
//
// If the returned token is [token.SEMICOLON], the corresponding
// literal string is ";" if the semicolon was present in the source,
// and "\n" if the semicolon was inserted because of a newline or
// at EOF.
//
// If the returned token is [token.ILLEGAL], the literal string is the
// offending character.
//
// In all other cases, Scan returns an empty literal string.
//
// For more tolerant parsing, Scan will return a valid token if
// possible even if a syntax error was encountered. Thus, even
// if the resulting token sequence contains no illegal tokens,
// a client may not assume that no error occurred. Instead it
// must check the scanner's ErrorCount or the number of calls
// of the error handler, if there was one installed.
//
// Scan adds line information to the file added to the file
// set with Init. Token positions are relative to that file
// and thus relative to the file set.
>>>>>>> upstream/master
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string)
