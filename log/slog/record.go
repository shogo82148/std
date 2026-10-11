// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/time"
)

const nAttrsInline = 4

// Recordは、ログイベントに関する情報を保持します。
// Recordのコピーは状態を共有します。
// Recordをコピーしてから変更しないでください。
// 新しいRecordを作成するには、 [NewRecord] を呼び出します。
// 共有状態のないコピーを作成するには、 [Record.Clone] を使用します。
type Record struct {
	// 出力メソッド（Log、Infoなど）が呼び出された時刻。
	Time time.Time

	// ログメッセージ。
	Message string

	// イベントのレベル。
	Level Level

	// レコードが構築された時点のプログラムカウンタで、runtime.Callersによって決定されます。
	// ゼロの場合、プログラムカウンタは利用できません。
	//
	// この値の唯一の有効な使用法は、[runtime.CallersFrames]の引数としての使用です。
	// 特に、[runtime.FuncForPC]に渡してはいけません。
	PC uintptr

	// Allocation optimization: an inline array sized to hold
	// the majority of log calls (based on examination of open-source
	// code). It holds the start of the list of Attrs.
	front [nAttrsInline]Attr

	// The number of Attrs in front.
	nFront int

	// The list of Attrs except for those in front.
	// Invariants:
	//   - len(back) > 0 iff nFront == len(front)
	//   - Unused array elements are zero. Used to detect mistakes.
	back []Attr
}

// NewRecordは、指定された引数から [Record] を作成します。
// Recordに属性を追加するには、 [Record.AddAttrs] を使用します。
//
// NewRecordは、 [Handler] をバックエンドとしてサポートするログAPIに使用することを想定しています。
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record

// Cloneは、共有状態のないレコードのコピーを返します。
// オリジナルのレコードとクローンの両方を変更できます。
// 互いに干渉しません。
func (r Record) Clone() Record

// NumAttrsは、[Record] の属性の数を返します。
func (r Record) NumAttrs() int

// Attrsは、[Record] 内の各Attrに対してfを呼び出します。
// fがfalseを返すと、反復処理が停止します。
func (r Record) Attrs(f func(Attr) bool)

// AddAttrsは、指定されたAttrsを [Record] のAttrsリストに追加します。
// 空のグループは省略されます。
func (r *Record) AddAttrs(attrs ...Attr)

// Addは、[Logger.Log]で説明されているように、argsをAttrsに変換し、
// [Record] のAttrsリストにAttrsを追加します。
// 空のグループは省略されます。
func (r *Record) Add(args ...any)

// Sourceは、ソースコードの行の場所を記述します。
type Source struct {
	// Functionは、ソース行を含むパッケージパスで修飾された関数名です。
	// 空でない場合、この文字列はプログラム内の単一の関数を一意に識別します。
	// 不明な場合は空文字列になることがあります。
	Function string `json:"function"`
	// FileとLineは、ソース行のファイル名と行番号（1始まり）です。
	// 不明な場合は、それぞれ空文字列とゼロになることがあります。
	File string `json:"file"`
	Line int    `json:"line"`
}

// Sourceは、rのPCを使用してログイベントの新しいSourceを返します。
// PCフィールドがゼロの場合、つまりRecordが必要な情報なしで作成されたか、
// 位置が利用できない場合は、nilが返されます。
func (r Record) Source() *Source
