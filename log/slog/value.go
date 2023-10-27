// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/time"
)

// Valueは、任意のGo値を表すことができますが、type anyとは異なり、
// 大部分の小さな値を割り当てなしで表現できます。
// ゼロ値のValueはnilに対応します。
type Value struct {
	_ [0]func()
	// num holds the value for Kinds Int64, Uint64, Float64, Bool and Duration,
	// the string length for KindString, and nanoseconds since the epoch for KindTime.
	num uint64
	// If any is of type Kind, then the value is in num as described above.
	// If any is of type *time.Location, then the Kind is Time and time.Time value
	// can be constructed from the Unix nanos in num and the location (monotonic time
	// is not preserved).
	// If any is of type stringptr, then the Kind is String and the string value
	// consists of the length in num and the pointer in any.
	// Otherwise, the Kind is Any and any is the value.
	// (This implies that Attrs cannot store values of type Kind, *time.Location
	// or stringptr.)
	any any
}

<<<<<<< HEAD
// KindはValueの種類です。
=======
// Kind is the kind of a [Value].
>>>>>>> upstream/master
type Kind int

const (
	KindAny Kind = iota
	KindBool
	KindDuration
	KindFloat64
	KindInt64
	KindString
	KindTime
	KindUint64
	KindGroup
	KindLogValuer
)

func (k Kind) String() string

// Kindは、Valueの種類を返します。
func (v Value) Kind() Kind

<<<<<<< HEAD
// StringValueは、文字列の新しいValueを返します。
func StringValue(value string) Value

// IntValueは、intのValueを返します。
func IntValue(v int) Value

// Int64Valueは、int64のValueを返します。
func Int64Value(v int64) Value

// Uint64Valueは、uint64のValueを返します。
func Uint64Value(v uint64) Value

// Float64Valueは、浮動小数点数のValueを返します。
func Float64Value(v float64) Value

// BoolValueは、boolのValueを返します。
func BoolValue(v bool) Value

// TimeValueは、time.TimeのValueを返します。
// monotonic部分は破棄されます。
func TimeValue(v time.Time) Value

// DurationValueは、time.DurationのValueを返します。
func DurationValue(v time.Duration) Value

// GroupValueは、Attrのリストの新しいValueを返します。
// 呼び出し元は、引数スライスを後で変更しないでください。
func GroupValue(as ...Attr) Value

// AnyValueは、提供された値のValueを返します。
=======
// StringValue returns a new [Value] for a string.
func StringValue(value string) Value

// IntValue returns a [Value] for an int.
func IntValue(v int) Value

// Int64Value returns a [Value] for an int64.
func Int64Value(v int64) Value

// Uint64Value returns a [Value] for a uint64.
func Uint64Value(v uint64) Value

// Float64Value returns a [Value] for a floating-point number.
func Float64Value(v float64) Value

// BoolValue returns a [Value] for a bool.
func BoolValue(v bool) Value

// TimeValue returns a [Value] for a [time.Time].
// It discards the monotonic portion.
func TimeValue(v time.Time) Value

// DurationValue returns a [Value] for a [time.Duration].
func DurationValue(v time.Duration) Value

// GroupValue returns a new [Value] for a list of Attrs.
// The caller must not subsequently mutate the argument slice.
func GroupValue(as ...Attr) Value

// AnyValue returns a [Value] for the supplied value.
>>>>>>> upstream/master
//
// 提供された値がValue型の場合、変更されずに返されます。
//
<<<<<<< HEAD
// Goの事前宣言されたstring、bool、または（非複素）数値型のいずれかの値が与えられた場合、
// AnyValueは、String、Bool、Uint64、Int64、またはFloat64の種類のValueを返します。
// 元の数値型の幅は保持されません。
//
// time.Timeまたはtime.Duration値が与えられた場合、AnyValueは、KindTimeまたはKindDurationのValueを返します。
// monotonic timeは保持されません。
//
// nil、または数値型の基礎型である名前付き型を含む、すべての他の型の値の場合、
// AnyValueは、KindAnyの種類のValueを返します。
=======
// Given a value of one of Go's predeclared string, bool, or
// (non-complex) numeric types, AnyValue returns a Value of kind
// [KindString], [KindBool], [KindUint64], [KindInt64], or [KindFloat64].
// The width of the original numeric type is not preserved.
//
// Given a [time.Time] or [time.Duration] value, AnyValue returns a Value of kind
// [KindTime] or [KindDuration]. The monotonic time is not preserved.
//
// For nil, or values of all other types, including named types whose
// underlying type is numeric, AnyValue returns a value of kind [KindAny].
>>>>>>> upstream/master
func AnyValue(v any) Value

// Anyは、vの値をanyとして返します。
func (v Value) Any() any

<<<<<<< HEAD
// Stringは、Valueの値をfmt.Sprintのようにフォーマットした文字列として返します。
// Int64、Float64などのメソッドは、vが間違った種類の場合にpanicしますが、
// Stringは決してpanicしません。
=======
// String returns Value's value as a string, formatted like [fmt.Sprint]. Unlike
// the methods Int64, Float64, and so on, which panic if v is of the
// wrong kind, String never panics.
>>>>>>> upstream/master
func (v Value) String() string

// Int64は、vの値をint64として返します。vが符号付き整数でない場合はpanicします。
func (v Value) Int64() int64

// Uint64は、vの値をuint64として返します。vが符号なし整数でない場合はpanicします。
func (v Value) Uint64() uint64

// Boolは、vの値をboolとして返します。vがboolでない場合はpanicします。
func (v Value) Bool() bool

<<<<<<< HEAD
// Durationは、vの値をtime.Durationとして返します。vがtime.Durationでない場合はpanicします。
func (a Value) Duration() time.Duration
=======
// Duration returns v's value as a [time.Duration]. It panics
// if v is not a time.Duration.
func (v Value) Duration() time.Duration
>>>>>>> upstream/master

// Float64は、vの値をfloat64として返します。vがfloat64でない場合はpanicします。
func (v Value) Float64() float64

<<<<<<< HEAD
// Timeは、vの値をtime.Timeとして返します。vがtime.Timeでない場合はpanicします。
=======
// Time returns v's value as a [time.Time]. It panics
// if v is not a time.Time.
>>>>>>> upstream/master
func (v Value) Time() time.Time

// LogValuerは、vの値をLogValuerとして返します。vがLogValuerでない場合はpanicします。
func (v Value) LogValuer() LogValuer

<<<<<<< HEAD
// Groupは、vの値を[]Attrとして返します。
// vのKindがKindGroupでない場合はpanicします。
=======
// Group returns v's value as a []Attr.
// It panics if v's [Kind] is not [KindGroup].
>>>>>>> upstream/master
func (v Value) Group() []Attr

// Equalは、vとwが同じGo値を表しているかどうかを報告します。
func (v Value) Equal(w Value) bool

// LogValuerは、自身をログ出力するためにValueに変換できる任意のGo値です。
//
// このメカニズムは、高価な操作を必要になるまで遅延させるために使用できます。
// また、単一の値を複数のコンポーネントに展開するために使用できます。
type LogValuer interface {
	LogValue() Value
}

<<<<<<< HEAD
// Resolveは、vがLogValuerを実装している間、vのLogValueを繰り返し呼び出し、結果を返します。
// vがグループに解決された場合、グループの属性の値は再帰的に解決されません。
// LogValueの呼び出し回数が閾値を超えた場合、エラーを含むValueが返されます。
// Resolveの戻り値は、KindLogValuerの種類ではないことが保証されています。
=======
// Resolve repeatedly calls LogValue on v while it implements [LogValuer],
// and returns the result.
// If v resolves to a group, the group's attributes' values are not recursively
// resolved.
// If the number of LogValue calls exceeds a threshold, a Value containing an
// error is returned.
// Resolve's return value is guaranteed not to be of Kind [KindLogValuer].
>>>>>>> upstream/master
func (v Value) Resolve() (rv Value)
