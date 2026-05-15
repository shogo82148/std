// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// scannerパッケージはGo言語のソーステキストのためのスキャナを実装します。
// ソースとして[]byteを受け取り、Scanメソッドへの繰り返し呼び出しを通じてトークン化します。
package scanner

import (
	"github.com/shogo82148/std/go/token"
)

// ErrorHandlerは [Scanner.Init] に提供することができます。構文エラーが発生し、ハンドラがインストールされている場合、ハンドラは位置とエラーメッセージとともに呼び出されます。位置は問題のあるトークンの始まりを指します。
type ErrorHandler func(pos token.Position, msg string)

// スキャナは、指定されたテキストを処理する間、スキャンの内部状態を保持します。それは他のデータ構造の一部として割り当てられることもできますが、使用前に [Scanner.Init] によって初期化される必要があります。
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

	endPosValid bool
	endPos      token.Pos

	// 公開される状態 - 変更しても問題ありません
	ErrorCount int
}

// Mode値はフラグのセット（または0）です。
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
// [Scanner.Scan] の呼び出しは、構文エラーが発生した場合にエラーハンドラerrを呼び出し、errがnilでない場合にはエラーハンドラerrを呼び出します。
// また、エラーが発生する毎に [Scanner] のErrorCountフィールドが1増加します。モードパラメーターはコメントの扱い方を決定します。
//
// Initはファイルの最初の文字にエラーがある場合、errを呼び出す場合があります。
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler, mode Mode)

// Endは最後にスキャンしたトークンの直後の位置を返します。
// まだ [Scanner.Scan] が呼び出されていない場合、Endは [token.NoPos] を返します。
func (s *Scanner) End() token.Pos

// Scanは次のトークンをスキャンし、トークン位置、トークン、
// および必要に応じてそのリテラル文字列を返します。ソースの終端は
// [token.EOF] で示されます。
//
// 返されたトークンがリテラル（[token.IDENT], [token.INT], [token.FLOAT],
// [token.IMAG], [token.CHAR], [token.STRING]）または [token.COMMENT] の場合、
// リテラル文字列は対応する値を持ちます。
//
// 返されたトークンがキーワードの場合、リテラル文字列はそのキーワードです。
//
// 返されたトークンが [token.SEMICOLON] の場合、対応する
// リテラル文字列は、ソース中にセミコロンが存在したときは ";"、
// 改行またはEOFによってセミコロンが補挿されたときは "\n" です。
// 改行が /*...*/ コメント内にある場合、SEMICOLONトークンは
// COMMENTトークンの直後に合成されます。その位置は、
// コメント内の実際の改行位置です。
//
// 返されたトークンが [token.ILLEGAL] の場合、リテラル文字列は
// 問題のある文字です。
//
// それ以外の場合、Scanは空のリテラル文字列を返します。
//
// より寛容な解析のため、Scanは構文エラーに遭遇した場合でも、
// 可能であれば有効なトークンを返します。したがって、結果として得られる
// トークン列に不正なトークンが含まれていなくても、クライアントは
// エラーが発生しなかったとは仮定できません。代わりに、
// スキャナのErrorCount、またはエラーハンドラがインストールされている場合は
// その呼び出し回数を確認する必要があります。
//
// Scanは、Initでファイルセットに追加されたファイルへ行情報を追加します。
// トークン位置はそのファイルに対する相対位置であり、
// したがってファイルセットに対する相対位置でもあります。
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string)
