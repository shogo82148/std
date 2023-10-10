// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// ErrorListでは、エラーは*Errorで表されます。
// Posは、有効な場合は問題のあるトークンの先頭を指し、エラーの状況はMsgで説明されます。
type Error struct {
	Pos token.Position
	Msg string
}

// Errorはerrorインターフェースを実装します。
func (e Error) Error() string

// ErrorListは*Errorsのリストです。
// ErrorListのゼロ値は使用する準備ができた空のErrorListです。
type ErrorList []*Error

// Addは、指定された位置とエラーメッセージを持つエラーをErrorListに追加します。
func (p *ErrorList) Add(pos token.Position, msg string)

// ResetはErrorListのエラーをリセットします。
func (p *ErrorList) Reset()

// ErrorListはsort Interfaceを実装します。
func (p ErrorList) Len() int
func (p ErrorList) Swap(i, j int)

func (p ErrorList) Less(i, j int) bool

// Sort関数は、ErrorListをソートします。*Errorのエントリは位置で、他のエラーはエラーメッセージでソートされ、*Errorのエントリの前に配置されます。
func (p ErrorList) Sort()

// RemoveMultiplesはErrorListをソートし、1行ごとに最初のエラー以外を削除します。
func (p *ErrorList) RemoveMultiples()

// ErrorListはerrorインターフェースを実装しています。
func (p ErrorList) Error() string

// Errはこのエラーリストに相当するエラーを返します。
// リストが空の場合、Errはnilを返します。
func (p ErrorList) Err() error

// PrintErrorは、errパラメータがErrorListの場合、エラーリストを1行ごとにwに出力します。それ以外の場合は、err文字列を出力します。
func PrintError(w io.Writer, err error)
