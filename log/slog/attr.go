// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/time"
)

// Attrはキーと値のペアです。
type Attr struct {
	Key   string
	Value Value
}

// Stringは文字列値のAttrを返します。
func String(key, value string) Attr

// Int64はint64のAttrを返します。
func Int64(key string, value int64) Attr

// Intはintをint64に変換し、その値を持つAttrを返します。
func Int(key string, value int) Attr

// Uint64はuint64のAttrを返します。
func Uint64(key string, v uint64) Attr

// Float64は浮動小数点数のAttrを返します。
func Float64(key string, v float64) Attr

// BoolはboolのAttrを返します。
func Bool(key string, v bool) Attr

// Timeは [time.Time] のAttrを返します。
// monotonic部分は破棄されます。
func Time(key string, v time.Time) Attr

// Durationは [time.Duration] のAttrを返します。
func Duration(key string, v time.Duration) Attr

// GroupはGroup [Value] のAttrを返します。
// 最初の引数はキーで、残りの引数は [Logger.Log] と同様にAttrsに変換されます。
//
// Groupを使用して、ログ行の単一のキーの下に複数のキー-値ペアを収集するか、
// LogValueの結果として単一の値を複数のAttrsとしてログに記録するために使用します。
func Group(key string, args ...any) Attr

// GroupAttrsは与えられたAttrsからなるGroup [Value] のAttrを返します。
//
// GroupAttrsは [Attr] 値のみを受け入れる、より効率的な [Group] のバージョンです。
func GroupAttrs(key string, attrs ...Attr) Attr

// Anyは指定された値のAttrを返します。
// 値の扱い方については [AnyValue] を参照してください。
func Any(key string, value any) Attr

// Equalはaとbが等しいキーと値を持つかどうかを報告します。
func (a Attr) Equal(b Attr) bool

func (a Attr) String() string
