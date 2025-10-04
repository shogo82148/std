// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/sync/atomic"
)

// Levelは、ログイベントの重要度または深刻度を表します。
// レベルが高いほど、イベントはより重要または深刻です。
type Level int

// 一般的なレベルの名前。
//
// レベル番号は本質的に任意ですが、3つの制約を満たすように選択しました。
// 任意のシステムは、別の番号付けスキームにマップできます。
//
// まず、デフォルトのレベルをInfoにしたかったため、Levelsはintであり、
// Infoはintのデフォルト値であるゼロです。
//
// 2番目に、レベルを使用してロガーの冗長性を指定することを簡単にしたかったです。
// より深刻なイベントは、より高いレベルを意味するため、
// より小さい（または負の）レベルのイベントを受け入れるロガーは、より冗長なロガーを意味します。
// ロガーの冗長性は、したがってイベントの深刻度の否定であり、
// デフォルトの冗長性0は、INFO以上のすべてのイベントを受け入れます。
//
// 3番目に、名前付きレベルを持つスキームを収容するために、レベル間に余裕が必要でした。
// たとえば、Google Cloud Loggingは、InfoとWarnの間にNoticeレベルを定義しています。
// これらの中間レベルはわずかであるため、数字の間のギャップは大きくする必要はありません。
// 私たちのギャップ4はOpenTelemetryのマッピングに一致します。
// OpenTelemetryのDEBUG、INFO、WARN、ERROR範囲から9を引くと、
// 対応するslog Level範囲に変換されます。
// OpenTelemetryにはTRACEとFATALという名前がありますが、slogにはありません。
// ただし、適切な整数を使用することで、これらのOpenTelemetryレベルをslog Levelsとして表すことができます。
const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

// Stringは、レベルの名前を返します。
// レベルに名前がある場合、その名前を大文字で返します。
// レベルが名前付き値の間にある場合、
// 大文字の名前に整数が追加されます。
// 例：
//
//	LevelWarn.String() => "WARN"
//	(LevelInfo+2).String() => "INFO+2"
func (l Level) String() string

// MarshalJSONは、 [Level.String] の出力を引用符で囲んで、
// [encoding/json.Marshaler] を実装します。
func (l Level) MarshalJSON() ([]byte, error)

// UnmarshalJSONは、 [encoding/json.Unmarshaler] を実装します。
// [Level.MarshalJSON] によって生成された任意の文字列を受け入れ、
// 大文字小文字を区別しません。
// また、出力上異なる文字列になる数値オフセットも受け入れます。
// たとえば、"Error-8"は "INFO" としてマーシャルされます。
func (l *Level) UnmarshalJSON(data []byte) error

<<<<<<< HEAD
// MarshalTextは、 [Level.String] を呼び出して、
// [encoding.TextMarshaler] を実装します。
=======
// AppendText implements [encoding.TextAppender]
// by calling [Level.String].
func (l Level) AppendText(b []byte) ([]byte, error)

// MarshalText implements [encoding.TextMarshaler]
// by calling [Level.AppendText].
>>>>>>> upstream/release-branch.go1.25
func (l Level) MarshalText() ([]byte, error)

// UnmarshalTextは、 [encoding.TextUnmarshaler] を実装します。
// [Level.MarshalText] によって生成された任意の文字列を受け入れ、
// 大文字小文字を区別しません。
// また、出力上異なる文字列になる数値オフセットも受け入れます。
// たとえば、"Error-8"は "INFO" としてマーシャルされます。
func (l *Level) UnmarshalText(data []byte) error

// Levelはレシーバーを返します。
// [Leveler] を実装します。
func (l Level) Level() Level

// LevelVarは、[Level] 変数を表し、[Handler] レベルを動的に変更するために使用されます。
// [Leveler] を実装すると同時に、Setメソッドも実装しており、
// 複数のゴルーチンから使用することができます。
// ゼロ値のLevelVarは [LevelInfo] に対応します。
type LevelVar struct {
	val atomic.Int64
}

// Levelは、vのレベルを返します。
func (v *LevelVar) Level() Level

// Setは、vのレベルをlに設定します。
func (v *LevelVar) Set(l Level)

func (v *LevelVar) String() string

<<<<<<< HEAD
// MarshalTextは、 [Level.MarshalText] を呼び出して、
// [encoding.TextMarshaler] を実装します。
=======
// AppendText implements [encoding.TextAppender]
// by calling [Level.AppendText].
func (v *LevelVar) AppendText(b []byte) ([]byte, error)

// MarshalText implements [encoding.TextMarshaler]
// by calling [LevelVar.AppendText].
>>>>>>> upstream/release-branch.go1.25
func (v *LevelVar) MarshalText() ([]byte, error)

// UnmarshalTextは、 [Level.UnmarshalText] を呼び出して、
// [encoding.TextUnmarshaler] を実装します。
func (v *LevelVar) UnmarshalText(data []byte) error

// Levelerは、[Level] 値を提供します。
//
// Level自体がLevelerを実装しているため、
// [HandlerOptions] など、Levelerが必要な場所では通常、Level値を提供します。
// レベルを動的に変更する必要があるクライアントは、
// *[LevelVar]などのより複雑なLeveler実装を提供できます。
type Leveler interface {
	Level() Level
}
