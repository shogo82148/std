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
	// The time at which the output method (Log, Info, etc.) was called.
	Time time.Time

	// The log message.
	Message string

	// The level of the event.
	Level Level

	// The program counter at the time the record was constructed, as determined
	// by runtime.Callers. If zero, no program counter is available.
	//
	// The only valid use for this value is as an argument to
	// [runtime.CallersFrames]. In particular, it must not be passed to
	// [runtime.FuncForPC].
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

<<<<<<< HEAD
// NewRecordは、指定された引数からRecordを作成します。
// Recordに属性を追加するには、 [Record.AddAttrs] を使用します。
=======
// NewRecord creates a [Record] from the given arguments.
// Use [Record.AddAttrs] to add attributes to the Record.
>>>>>>> upstream/master
//
// NewRecordは、 [Handler] をバックエンドとしてサポートするログAPIに使用することを想定しています。
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record

// Cloneは、共有状態のないレコードのコピーを返します。
// オリジナルのレコードとクローンの両方を変更できます。
// 互いに干渉しません。
func (r Record) Clone() Record

<<<<<<< HEAD
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
=======
// NumAttrs returns the number of attributes in the [Record].
func (r Record) NumAttrs() int

// Attrs calls f on each Attr in the [Record].
// Iteration stops if f returns false.
func (r Record) Attrs(f func(Attr) bool)

// AddAttrs appends the given Attrs to the [Record]'s list of Attrs.
// It omits empty groups.
func (r *Record) AddAttrs(attrs ...Attr)

// Add converts the args to Attrs as described in [Logger.Log],
// then appends the Attrs to the [Record]'s list of Attrs.
// It omits empty groups.
>>>>>>> upstream/master
func (r *Record) Add(args ...any)

// Sourceは、ソースコードの行の場所を記述します。
type Source struct {
	// Function is the package path-qualified function name containing the
	// source line. If non-empty, this string uniquely identifies a single
	// function in the program. This may be the empty string if not known.
	Function string `json:"function"`
	// File and Line are the file name and line number (1-based) of the source
	// line. These may be the empty string and zero, respectively, if not known.
	File string `json:"file"`
	Line int    `json:"line"`
}
