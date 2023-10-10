// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scannerはGo言語のソーステキストのためのスキャナを実装します。
// ソースとして[]byteを受け取り、Scanメソッドへの繰り返し呼び出しを通じてトークン化します。
package scanner

import (
	"github.com/shogo82148/std/go/token"
)

// ErrorHandlerはScanner.Initに提供することができます。構文エラーが発生し、ハンドラがインストールされている場合、ハンドラは位置とエラーメッセージとともに呼び出されます。位置は問題のあるトークンの始まりを指します。
type ErrorHandler func(pos token.Position, msg string)

// スキャナは、指定されたテキストを処理する間、スキャンの内部状態を保持します。それは他のデータ構造の一部として割り当てられることもできますが、使用前にInitによって初期化される必要があります。
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
// Scanの呼び出しは、構文エラーが発生した場合にエラーハンドラerrを呼び出し、errがnilでない場合にはエラーハンドラerrを呼び出します。
// また、エラーが発生する毎にスキャナのErrorCountフィールドが1増加します。モードパラメーターはコメントの扱い方を決定します。
//
// Initはファイルの最初の文字にエラーがある場合、errを呼び出す場合があります。
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler, mode Mode)

// Scanは次のトークンをスキャンして、トークンの位置、トークン、およびリテラル文字列（適用可能な場合）を返します。ソースの終了はtoken.EOFで示されます。
// 返されるトークンがリテラル（token.IDENT、token.INT、token.FLOAT、token.IMAG、token.CHAR、token.STRING）またはtoken.COMMENTである場合、リテラル文字列は対応する値を持ちます。
// 返されるトークンがキーワードの場合、リテラル文字列はキーワードです。
// 返されるトークンがtoken.SEMICOLONの場合、対応するリテラル文字列は、ソースにセミコロンが存在する場合は";"であり、改行またはEOFの場合は"\n"です。
// 返されるトークンがtoken.ILLEGALの場合、リテラル文字列は違反している文字です。
// その他の場合、Scanは空のリテラル文字列を返します。
// より許容度の高い解析のために、構文エラーが発生した場合でも、可能な限り有効なトークンを返します。したがって、結果のトークンのシーケンスに違法なトークンが含まれていない場合でも、クライアントはエラーが発生したとは推測できません。代わりに、スキャナーのErrorCountまたはエラーハンドラーの呼び出し回数を確認する必要があります（インストールされている場合）。
// ScanはInitでファイルセットに追加されたファイルに行情報を追加します。トークンの位置はそのファイルとファイルセットに対して相対的です。
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string)
