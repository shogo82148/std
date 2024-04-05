// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

<<<<<<< HEAD
// これらはTime.Formatとtime.Parseで使用するための事前定義されたレイアウトです。
// これらのレイアウトで使用される参照時間は、特定のタイムスタンプです：
=======
// These are predefined layouts for use in [Time.Format] and [time.Parse].
// The reference time used in these layouts is the specific time stamp:
>>>>>>> upstream/master
//
// 01/02 03:04:05PM '06 -0700
//
<<<<<<< HEAD
// (2006年1月2日15時04分05秒、GMTより7時間西のタイムゾーン)。
// その値は、以下にリストされている定数であるLayoutとして記録されます。UNIX
// タイムとしては、1136239445です。MSTはGMT-0700であるため、参照時間は
// UNIXのdateコマンドで次のように表示されます：
=======
// (January 2, 15:04:05, 2006, in time zone seven hours west of GMT).
// That value is recorded as the constant named [Layout], listed below. As a Unix
// time, this is 1136239445. Since MST is GMT-0700, the reference would be
// printed by the Unix date command as:
>>>>>>> upstream/master
//
// Mon Jan 2 15:04:05 MST 2006
//
// 月の数字を日の前に置くアメリカの慣習を採用していることは、歴史的な誤りです。
//
// Time.Formatの例では、レイアウト文字列の動作を詳しく説明しており、参考になります。
//
<<<<<<< HEAD
// RFC822、RFC850、RFC1123のフォーマットは、ローカル時間にのみ適用する必要があります。
// UTC時間に適用する場合、時差表示として「UTC」が使用されますが、厳密にはこれらのRFCでは
// その場合に「GMT」の使用が必要です。
// 一般に、RFC1123の代わりにRFC1123Zを使用するべきです。
// また、新しいプロトコルにはRFC3339を優先すべきです。
// RFC3339、RFC822、RFC822Z、RFC1123、RFC1123Zは、フォーマット用途に有用です。
// time.Parseで使用する場合、これらはRFCで許可されているすべての時間形式を受け付けず、
// 正式に定義されていない時間形式を受け入れます。
// RFC3339Nano形式は秒の末尾のゼロを削除するため、フォーマット後に正しくソートされない場合があります。
=======
// Note that the [RFC822], [RFC850], and [RFC1123] formats should be applied
// only to local times. Applying them to UTC times will use "UTC" as the
// time zone abbreviation, while strictly speaking those RFCs require the
// use of "GMT" in that case.
// In general [RFC1123Z] should be used instead of [RFC1123] for servers
// that insist on that format, and [RFC3339] should be preferred for new protocols.
// [RFC3339], [RFC822], [RFC822Z], [RFC1123], and [RFC1123Z] are useful for formatting;
// when used with time.Parse they do not accept all the time formats
// permitted by the RFCs and they do accept time formats not formally defined.
// The [RFC3339Nano] format removes trailing zeros from the seconds field
// and thus may not sort correctly once formatted.
>>>>>>> upstream/master
//
// ほとんどのプログラムは、FormatやParseに渡すための定義済みの定数の1つを使用できます。
// カスタムレイアウト文字列を作成する場合以外は、このコメントの残りは無視してかまいません。
//
<<<<<<< HEAD
// 独自のフォーマットを定義するには、参照時間があなたの方法でどのように
// フォーマットされるかを書き出してください。ANSIC、StampMicro、Kitchenなどの
// 定数の値を参照してください。モデルは、参照時間がどのようになっているかを実証し、
// FormatとParseメソッドが一般的な時間値に同じ変換を適用できるようにすることです。
=======
// To define your own format, write down what the reference time would look like
// formatted your way; see the values of constants like [ANSIC], [StampMicro] or
// [Kitchen] for examples. The model is to demonstrate what the reference time
// looks like so that the Format and Parse methods can apply the same
// transformation to a general time value.
>>>>>>> upstream/master
//
// 以下はレイアウト文字列のコンポーネントの概要です。各要素は参照時間の要素のフォーマットを
// 例示しています。これらの値のみが認識されます。参照時間の一部として認識されないレイアウト文字列の
// テキストは、Formatでそのまま出力され、Parseの入力にそのまま表示されると予想されます。
//
// 年: "2006" "06"
// 月: "Jan" "January" "01" "1"
// 曜日: "Mon" "Monday"
// 月の日にち: "2" "_2" "02"
// 年の日にち: "__2" "002"
// 時: "15" "3" "03" (PMまたはAM)
// 分: "4" "04"
// 秒: "5" "05"
// AM/PMマーク: "PM"
//
// 数値のタイムゾーンオフセットは以下のようにフォーマットされます：
//
// "-0700" ±hhmm
// "-07:00" ±hh:mm
// "-07" ±hh
// "-070000" ±hhmmss
// "-07:00:00" ±hh:mm:ss
//
// フォーマット内の符号をZに置き換えると、UTCゾーンのオフセットではなく
// ZがプリントされるISO 8601の動作になります。
// したがって：
//
// "Z0700" Zまたは±hhmm
// "Z07:00" Zまたは±hh:mm
// "Z07" Zまたは±hh
// "Z070000" Zまたは±hhmmss
// "Z07:00:00" Zまたは±hh:mm:ss
//
// フォーマット文字列内では、"_2"と"__2"の下線は、次の数字が複数桁である場合に
// 数字に置き換えられる可能性のあるスペースを表します。
// これにより、固定桁のUNIXタイムフォーマットとの互換性が保たれます。先頭のゼロは
// ゼロパディングされた値を表します。
//
// フォーマット002は空白でパディングされ、ゼロでパディングされた3文字の年の日を表します。
// パディングされていない年の日フォーマットはありません。
//
// カンマまたは小数点の後にゼロが1つ以上続く場合、指定した小数桁数で出力される小数秒を表します。
// カンマまたは小数点の後に9が1つ以上続く場合、指定した小数桁数で出力され、末尾のゼロは削除されます。
// たとえば、「15:04:05,000」または「15:04:05.000」はミリ秒の精度でフォーマットまたは解析します。
//
// いくつかの有効なレイアウトは、space paddingのようなフォーマットやzone情報のZのような
// フォーマットのため、time.Parseには無効な時間値です。
const (
	Layout      = "01/02 03:04:05PM '06 -0700"
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700"
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// 便利なタイムスタンプ。
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
)

const (
	_ = iota
)

// Stringはフォーマット文字列を使用して書式付けされた時間を返します
//
//	"2006-01-02 15:04:05.999999999 -0700 MST"
//
// もし時間が単調なクロック読み取りを持っている場合、返される文字列は
// "m=±<value>"という最後のフィールドを含みます。ここで、valueは
// 単調なクロック読み取りを秒で表した10進数です。
//
// 返される文字列はデバッグ用途であり、安定したシリアライズされた表現には
// t.MarshalText、t.MarshalBinary、または明示的なフォーマット文字列を使用した
// t.Formatを使用してください。
func (t Time) String() string

<<<<<<< HEAD
// GoStringはfmt.GoStringerを実装し、Goソースコードで表示されるようにtをフォーマットします。
func (t Time) GoString() string

// Formatは、引数で定義されたレイアウトに従って、時刻値のテキスト表現を返します。レイアウトフォーマットの表現方法については、「Layout」という定数のドキュメントを参照してください。
//
// Time.Formatの実行可能な例は、レイアウト文字列の詳細な動作を示しており、参考になります。
func (t Time) Format(layout string) string

// AppendFormatはFormatと似ていますが、テキスト表現をbに追加し、拡張されたバッファを返します。
=======
// GoString implements [fmt.GoStringer] and formats t to be printed in Go source
// code.
func (t Time) GoString() string

// Format returns a textual representation of the time value formatted according
// to the layout defined by the argument. See the documentation for the
// constant called [Layout] to see how to represent the layout format.
//
// The executable example for [Time.Format] demonstrates the working
// of the layout string in detail and is a good reference.
func (t Time) Format(layout string) string

// AppendFormat is like [Time.Format] but appends the textual
// representation to b and returns the extended buffer.
>>>>>>> upstream/master
func (t Time) AppendFormat(b []byte, layout string) []byte

// ParseErrorは時間文字列を解析する際の問題を記述します。
type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
	Message    string
}

// ErrorはParseErrorの文字列表現を返します。
func (e *ParseError) Error() string

<<<<<<< HEAD
// Parseは書式指定された文字列を解析し、それが表す時間の値を返します。
// 解析するには、レイアウトとして提供された形式文字列（レイアウト）を使用して、
// 第1引数として提供された解析可能な第2引数が必要です。
//
// Time.Formatの例はレイアウト文字列の動作を詳しく説明しており、参考となります。
=======
// Parse parses a formatted string and returns the time value it represents.
// See the documentation for the constant called [Layout] to see how to
// represent the format. The second argument must be parseable using
// the format string (layout) provided as the first argument.
//
// The example for [Time.Format] demonstrates the working of the layout string
// in detail and is a good reference.
>>>>>>> upstream/master
//
// 解析（Parse）時には、レイアウトがその存在を示していなくても、入力には秒の次に少数秒
// フィールドが直ちに続く場合があります。その場合、最大連続桁の直後にカンマまたは
// 小数点が続くものを少数秒として解析します。少数秒はナノ秒まで切り捨てられます。
//
// レイアウトから省略された要素はゼロであるか、ゼロが不可能な場合は1であるとみなされます。
// したがって、"3:04pm"を解析すると、年0月1日15:04:00 UTCに対応する時間が返されます
// （年が0のため、この時間はゼロ時刻より前です）。
// 年は0000〜9999の範囲内である必要があります。曜日は構文チェックされますが、
// それ以外は無視されます。
//
// 2桁の年06を指定するレイアウトの場合、値NN >= 69は19NNとして扱われ、
// 値NN < 69は20NNとして扱われます。
//
// このコメントの残りの部分は、タイムゾーンの処理方法について説明しています。
//
// タイムゾーン指示がない場合、ParseはUTCで時間を返します。
//
<<<<<<< HEAD
// -0700のようなタイムゾーンオフセットを持つ時間を解析する場合、オフセットが現在の場所（ローカル）で使用されている
// タイムゾーンに対応している場合、Parseはその場所とタイムゾーンを使用して時間を返します。
// そうでない場合は、与えられたゾーンオフセットで時間が固定された、架空の場所として時間を記録します。
//
// MSTのようなタイムゾーン省略形を持つ時間を解析する場合、現在の場所で定義されたオフセットがある場合、それを使用します。
// "UTC"というタイムゾーン省略形は、場所に関係なくUTCとして認識されます。
// タイムゾーン省略形が不明な場合、Parseは与えられたゾーン省略形とゼロオフセットの架空の場所に
// 時間を記録します。この選択肢は、そのような時間をレイアウトの変更なしで解析および再フォーマットできますが、
// 表現に使用される正確な瞬間は実際のゾーンオフセットによって異なります。
// そのような問題を回避するためには、数値のゾーンオフセットを使用する時間レイアウトを使用するか、ParseInLocationを使用してください。
=======
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location ([Local]), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation "UTC" is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use [ParseInLocation].
>>>>>>> upstream/master
func Parse(layout, value string) (Time, error)

// ParseInLocationはParseと似ていますが、2つの重要な違いがあります。
// まず、タイムゾーン情報がない場合、Parseは時間をUTCとして解釈しますが、
// ParseInLocationは指定された場所の時間として解釈します。
// さらに、ゾーンオフセットや略語が与えられた場合、Parseはそれをローカルの場所と照合しようとしますが、
// ParseInLocationは指定された場所を使用します。
func ParseInLocation(layout, value string, loc *Location) (Time, error)

// ParseDurationは期間文字列を解析します。
// 期間文字列は、可能性のある符号付きの連続した
// 小数点数、オプションの小数部および単位接尾辞からなります。
// 例： "300ms"、"-1.5h"、または "2h45m"。
// 有効な時間単位は「ns」、「us」（または「µs」）、「ms」、「s」、「m」、「h」です。
func ParseDuration(s string) (Duration, error)
