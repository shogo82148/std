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
	Time time.Time

	Message string

	Level Level

	PC uintptr

	front [nAttrsInline]Attr

	nFront int

	back []Attr
}

// NewRecordは、指定された引数からRecordを作成します。
// Recordに属性を追加するには、 [Record.AddAttrs] を使用します。
//
// NewRecordは、 [Handler] をバックエンドとしてサポートするログAPIに使用することを想定しています。
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record

// Cloneは、共有状態のないレコードのコピーを返します。
// オリジナルのレコードとクローンの両方を変更できます。
// 互いに干渉しません。
func (r Record) Clone() Record

// NumAttrsは、Recordの属性の数を返します。
func (r Record) NumAttrs() int

// Attrsは、Record内の各Attrに対してfを呼び出します。
// fがfalseを返すと、反復処理が停止します。
func (r Record) Attrs(f func(Attr) bool)

// AddAttrsは、指定されたAttrsをRecordのAttrsリストに追加します。
// 空のグループは省略されます。
func (r *Record) AddAttrs(attrs ...Attr)

// Addは、[Logger.Log]で説明されているように、argsをAttrsに変換し、
// RecordのAttrsリストにAttrsを追加します。
// 空のグループは省略されます。
func (r *Record) Add(args ...any)

// Sourceは、ソースコードの行の場所を記述します。
type Source struct {
	Function string `json:"function"`

	File string `json:"file"`
	Line int    `json:"line"`
}
