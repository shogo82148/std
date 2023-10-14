// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/time"
)

func ExampleDuration() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func ExampleDuration_Round() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	}
	// Output:
	// d.Round(   1ns) = 1h15m30.918273645s
	// d.Round(   1µs) = 1h15m30.918274s
	// d.Round(   1ms) = 1h15m30.918s
	// d.Round(    1s) = 1h15m31s
	// d.Round(    2s) = 1h15m30s
	// d.Round(  1m0s) = 1h16m0s
	// d.Round( 10m0s) = 1h20m0s
	// d.Round(1h0m0s) = 1h0m0s
}

func ExampleDuration_String() {
	fmt.Println(1*time.Hour + 2*time.Minute + 300*time.Millisecond)
	fmt.Println(300 * time.Millisecond)
	// Output:
	// 1h2m0.3s
	// 300ms
}

func ExampleDuration_Truncate() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, t := range trunc {
		fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())
	}
	// Output:
	// d.Truncate(   1ns) = 1h15m30.918273645s
	// d.Truncate(   1µs) = 1h15m30.918273s
	// d.Truncate(   1ms) = 1h15m30.918s
	// d.Truncate(    1s) = 1h15m30s
	// d.Truncate(    2s) = 1h15m30s
	// d.Truncate(  1m0s) = 1h15m0s
	// d.Truncate( 10m0s) = 1h10m0s
	// d.Truncate(1h0m0s) = 1h0m0s
}

func ExampleParseDuration() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// このパッケージでは、マイクロのための誤ったが一般的な接頭辞uも受け入れます。
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
	// Output:
	// 10h0m0s
	// 1h10m10s
	// There are 4210 seconds in 1h10m10s.
	// There are 1000 nanoseconds in 1µs.
	// There are 1.00e-06 seconds in 1µs.
}

func ExampleDuration_Hours() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.", h.Hours())
	// Output: I've got 4.5 hours of work left.
}

func ExampleDuration_Microseconds() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d microseconds.\n", u.Microseconds())
	// Output:
	// One second is 1000000 microseconds.
}

func ExampleDuration_Milliseconds() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d milliseconds.\n", u.Milliseconds())
	// Output:
	// One second is 1000 milliseconds.
}

func ExampleDuration_Minutes() {
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
	// Output: The movie is 90 minutes long.
}

func ExampleDuration_Nanoseconds() {
	u, _ := time.ParseDuration("1µs")
	fmt.Printf("One microsecond is %d nanoseconds.\n", u.Nanoseconds())
	// Output:
	// One microsecond is 1000 nanoseconds.
}

func ExampleDuration_Seconds() {
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.", m.Seconds())
	// Output: Take off in t-90 seconds.
}

func ExampleAfter() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}

func ExampleSleep() {
	time.Sleep(100 * time.Millisecond)
}

func ExampleTick() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

func ExampleMonth() {
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}

func ExampleDate() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	// Output: Go launched at 2009-11-10 15:00:00 -0800 PST
}

func ExampleNewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

func ExampleTime_Format() {
	// 標準のUnix形式の文字列から時刻値を解析する。
	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // 必ずエラーを確認する。たとえそれが起こらないはずでも。
		panic(err)
	}

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil { // 常にエラーチェックを行います。たとえ起こるべきでない場合でも。
		panic(err)
	}

	// time.TimeのStringerメソッドは、フォーマットなしでも便利です。
	fmt.Println("default format:", t)

	// パッケージ内の予め定義された定数は、共通のレイアウトを実装します。
	fmt.Println("Unix format:", t.Format(time.UnixDate))

	// 時間値に付けられたタイムゾーンは出力に影響を与えます。
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	fmt.Println("in Shanghai with seconds:", t.In(tz).Format("2006-01-02T15:04:05 -070000"))

	fmt.Println("in Shanghai with colon seconds:", t.In(tz).Format("2006-01-02T15:04:05 -07:00:00"))

	// この関数の残りの部分は、フォーマットに使用されるレイアウト文字列の特性を示しています。

	// Parse関数とFormatメソッドで使用されるレイアウト文字列は、
	// 参照時刻の表現方法を例示します。
	// 強調することは、参照時刻のフォーマット方法を示す必要があるということであり、
	// ユーザーが選択した時刻ではありません。そのため、各レイアウト文字列は
	// 時刻スタンプの表現です。
	// Jan 2 15:04:05 2006 MST
	// この値を覚えるための簡単な方法は、この順序で表示されたときに、
	// 上記の要素と一致する値を保持していることです。
	// 1 2 3 4 5 6 -7
	// 以下では、いくつかの例が示されています。

	// FormatやParseのほとんどの使用例は、このパッケージで定義されているような定数のレイアウト文字列を使用しますが、
	// これらの例が示すように、インタフェースは柔軟です。

	// 例の出力を見やすくするためのヘルパー関数を定義する。
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// 出力にヘッダーを表示します。
	fmt.Printf("\nFormats:\n\n")

	// シンプルなスターターの例。
	do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
	do("Basic short date", "2006/01/02", "2015/02/25")

	// 参照時刻の時は15、または午後3時です。レイアウトはどちらの形式でも表現できますが、私たちの値は朝の時間なので、AMとして表示する必要があります。両方を1つの書式文字列で表示します。小文字でも。
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}

	// レイアウト文字列に少数秒の表現が含まれていない場合、出力には表示されません。
	do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")

	// 配列文字列の中で、秒数の値の小数点以下に0または9が連続していると、小数点以下の秒数が表示されます。
	// レイアウトの桁が0の場合、指定された桁数の小数秒になります。ただし、出力には末尾に0が付きます。
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

	// レイアウト内の小数部が9sの場合、末尾のゼロは削除されます。
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")

	// Output:
	// default format: 2015-02-25 11:06:39 -0800 PST
	// Unix format: Wed Feb 25 11:06:39 PST 2015
	// Same, in UTC: Wed Feb 25 19:06:39 UTC 2015
	//in Shanghai with seconds: 2015-02-26T03:06:39 +080000
	//in Shanghai with colon seconds: 2015-02-26T03:06:39 +08:00:00
	//
	// Formats:
	//
	// Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	// Basic short date "2006/01/02" gives "2015/02/25"
	// AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
	// No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	// 0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
	// 9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"

}

func ExampleTime_Format_pad() {
	// 標準のUnix形式の文字列から時間値を解析します。
	t, err := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
	if err != nil { // 必要ないと思われるかもしれないが、常にエラーをチェックしてください。
		panic(err)
	}

	// ヘルパー関数を定義して、例の出力を見やすくします。
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// 予定された定数 Unix は日を埋めるためにアンダースコアを使用します。
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	// 固定幅で値をプリントするために、日付などの1文字または2文字になる可能性がある場合、レイアウト文字列の代わりにスペースの代わりに_を使用します。
	// ここでは、レイアウト文字列では2で値は7の日を印刷しています。
	do("No pad", "<2>", "<7>")

	// アンダースコアはスペースのパッドを表します。日付が1桁の場合。
	do("Spaces", "<_2>", "< 7>")

	// "0"は一桁の値のゼロ埋めを示します。
	do("Zeros", "<02>", "<07>")

	// 値がすでに適切な幅である場合、パディングは使用されません。
	// たとえば、値の2番目（リファレンス時刻の05）は39ですが、
	// したがって、パディングは必要ありませんが、分の部分（04, 06）には必要です。
	do("Suppressed pad", "04:05", "06:39")

	// Output:
	// Unix             "Mon Jan _2 15:04:05 MST 2006" gives "Sat Mar  7 11:06:39 PST 2015"
	// No pad           "<2>" gives "<7>"
	// Spaces           "<_2>" gives "< 7>"
	// Zeros            "<02>" gives "<07>"
	// Suppressed pad   "04:05" gives "06:39"

}

func ExampleTime_GoString() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t.GoString())
	t = t.Add(1 * time.Minute)
	fmt.Println(t.GoString())
	t = t.AddDate(0, 1, 0)
	fmt.Println(t.GoString())
	t, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	fmt.Println(t.GoString())

	// Output:
	// time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// time.Date(2009, time.November, 10, 23, 1, 0, 0, time.UTC)
	// time.Date(2009, time.December, 10, 23, 1, 0, 0, time.UTC)
	// time.Date(2013, time.February, 3, 19, 54, 0, 0, time.UTC)
}

func ExampleParse() {

	// Time.Formatの例を参照して、time.Time値を解析するためにレイアウト文字列を定義する方法の詳細を見てください。ParseとFormatは、入力と出力を記述するために同じモデルを使用します。

	// longFormは、参照時刻がどのように所望のレイアウトで表されるかを例示します。
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortFormは望ましいレイアウトで参照時刻が表されるもう一つの方法です。これにはタイムゾーンが存在しません。
	// 注意：明示的なゾーンがない場合、時間はUTCで返されます。
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// いくつかの有効なレイアウトは、スペースの埋め込みのための _ やゾーン情報のための Z などのフォーマット指定子によって無効な時間値となります。
	// たとえば、RFC3339のレイアウト 2006-01-02T15:04:05Z07:00 は、有効なオプションの両方を扱うために、Z とタイムゾーンオフセットの両方を含んでいます。
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // レイアウトが有効な時間の値ではないため、エラーが返されます。

	// Output:
	// 2013-02-03 19:54:00 -0800 PST
	// 2013-02-03 00:00:00 +0000 UTC
	// 2006-01-02 15:04:05 +0000 UTC
	// 2006-01-02 15:04:05 +0700 +0700
	// error parsing time "2006-01-02T15:04:05Z07:00": extra text: "07:00"
}

func ExampleParseInLocation() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// これはEurope/Berlinの時刻帯でCESTという名前を探します。
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// 注意：明示的なタイムゾーンがない場合、指定された場所の時間を返します。
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	// Output:
	// 2012-07-09 05:02:00 +0200 CEST
	// 2012-07-09 00:00:00 +0200 CEST
}

func ExampleUnix() {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println(t)

	// Output:
	// 1257894000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleUnixMicro() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMicro())
	t := time.UnixMicro(umt.UnixMicro()).UTC()
	fmt.Println(t)

	// Output:
	// 1257894000000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleUnixMilli() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMilli())
	t := time.UnixMilli(umt.UnixMilli()).UTC()
	fmt.Println(t)

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleTime_Unix() {
	// Unixの10億秒、3つの方法。
	fmt.Println(time.Unix(1e9, 0).UTC())     // 1億秒
	fmt.Println(time.Unix(0, 1e18).UTC())    // 1e18ナノ秒
	fmt.Println(time.Unix(2e9, -1e18).UTC()) // 2e9 秒 - 1e18 ナノ秒

	t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
	fmt.Println(t.Unix())     // 1970年からの経過時間（秒）
	fmt.Println(t.UnixNano()) // 1970 年以来のナノ秒

	// Output:
	// 2001-09-09 01:46:40 +0000 UTC
	// 2001-09-09 01:46:40 +0000 UTC
	// 2001-09-09 01:46:40 +0000 UTC
	// 1000000000
	// 1000000000000000000
}

func ExampleTime_Round() {
	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
	// Output:
	// t.Round(   1ns) = 12:15:30.918273645
	// t.Round(   1µs) = 12:15:30.918274
	// t.Round(   1ms) = 12:15:30.918
	// t.Round(    1s) = 12:15:31
	// t.Round(    2s) = 12:15:30
	// t.Round(  1m0s) = 12:16:00
	// t.Round( 10m0s) = 12:20:00
	// t.Round(1h0m0s) = 12:00:00
}

func ExampleTime_Truncate() {
	t, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
	}

	for _, d := range trunc {
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
	// ローカルタイムゾーンの最後の深夜に丸めるには、新しい日付を作成します。
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	_ = midnight

	// Output:
	// t.Truncate(  1ns) = 12:15:30.918273645
	// t.Truncate(  1µs) = 12:15:30.918273
	// t.Truncate(  1ms) = 12:15:30.918
	// t.Truncate(   1s) = 12:15:30
	// t.Truncate(   2s) = 12:15:30
	// t.Truncate( 1m0s) = 12:15:00
	// t.Truncate(10m0s) = 12:10:00
}

func ExampleLoadLocation() {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location))
	// Output: 2018-08-30 05:00:00 -0700 PDT
}

func ExampleLocation() {
	// 中国はサマータイムを実施していません。UTCから8時間の固定オフセットを使用しています。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// システムにタイムゾーンデータベースが存在する場合、その場所を読み込むことが可能です。例えば：
	//    ニューヨーク、err := time.LoadLocation("America/New_York")

	// 時間の作成には場所が必要です。一般的な場所にはtime.Localとtime.UTCがあります。
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// UTCの時計の時間が1200であり、北京の時計の時間が2000であるが、北京は8時間進んでいるため、
	// 実際にはこの2つの日時は同じ瞬間を表しています。
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// Output:
	// true
}

func ExampleTime_Add() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)

	// Output:
	// start = 2009-01-01 12:00:00 +0000 UTC
	// start.Add(time.Second * 10) = 2009-01-01 12:00:10 +0000 UTC
	// start.Add(time.Minute * 10) = 2009-01-01 12:10:00 +0000 UTC
	// start.Add(time.Hour * 10) = 2009-01-01 22:00:00 +0000 UTC
	// start.Add(time.Hour * 24 * 10) = 2009-01-11 12:00:00 +0000 UTC
}

func ExampleTime_AddDate() {
	start := time.Date(2023, 03, 25, 12, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	dayDuration := oneDayLater.Sub(start)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	zurich, err := time.LoadLocation("Europe/Zurich")
	if err != nil {
		panic(err)
	}
	// This was the day before a daylight saving time transition in Zürich.
	startZurich := time.Date(2023, 03, 25, 12, 0, 0, 0, zurich)
	oneDayLaterZurich := startZurich.AddDate(0, 0, 1)
	dayDurationZurich := oneDayLaterZurich.Sub(startZurich)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)
	fmt.Printf("oneDayLaterZurich: startZurich.AddDate(0, 0, 1) = %v\n", oneDayLaterZurich)
	fmt.Printf("Day duration in UTC: %v | Day duration in Zürich: %v\n", dayDuration, dayDurationZurich)

	// Output:
	// oneDayLater: start.AddDate(0, 0, 1) = 2023-03-26 12:00:00 +0000 UTC
	// oneMonthLater: start.AddDate(0, 1, 0) = 2023-04-25 12:00:00 +0000 UTC
	// oneYearLater: start.AddDate(1, 0, 0) = 2024-03-25 12:00:00 +0000 UTC
	// oneDayLaterZurich: startZurich.AddDate(0, 0, 1) = 2023-03-26 12:00:00 +0200 CEST
	// Day duration in UTC: 24h0m0s | Day duration in Zürich: 23h0m0s
}

func ExampleTime_After() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear3000AfterYear2000 := year3000.After(year2000) // 本当
	isYear2000AfterYear3000 := year2000.After(year3000) // 偽

	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

	// Output:
	// year3000.After(year2000) = true
	// year2000.After(year3000) = false
}

func ExampleTime_Before() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // 真
	isYear3000BeforeYear2000 := year3000.Before(year2000) // false

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

	// Output:
	// year2000.Before(year3000) = true
	// year3000.Before(year2000) = false
}

func ExampleTime_Date() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

	// Output:
	// year = 2000
	// month = February
	// day = 1
}

func ExampleTime_Day() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	day := d.Day()

	fmt.Printf("day = %v\n", day)

	// Output:
	// day = 1
}

func ExampleTime_Equal() {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// equal演算子とは異なり、Equalはd1とd2が同じ瞬間であることを時差が異なるものとして認識しています。
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)

	datesEqualUsingEqualOperator := d1 == d2
	datesEqualUsingFunction := d1.Equal(d2)

	fmt.Printf("datesEqualUsingEqualOperator = %v\n", datesEqualUsingEqualOperator)
	fmt.Printf("datesEqualUsingFunction = %v\n", datesEqualUsingFunction)

	// Output:
	// datesEqualUsingEqualOperator = false
	// datesEqualUsingFunction = true
}

func ExampleTime_String() {
	timeWithNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC)
	withNanoseconds := timeWithNanoseconds.String()

	timeWithoutNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 0, time.UTC)
	withoutNanoseconds := timeWithoutNanoseconds.String()

	fmt.Printf("withNanoseconds = %v\n", string(withNanoseconds))
	fmt.Printf("withoutNanoseconds = %v\n", string(withoutNanoseconds))

	// Output:
	// withNanoseconds = 2000-02-01 12:13:14.000000015 +0000 UTC
	// withoutNanoseconds = 2000-02-01 12:13:14 +0000 UTC
}

func ExampleTime_Sub() {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference)

	// Output:
	// difference = 12h0m0s
}

func ExampleTime_AppendFormat() {
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))

	// Output:
	// Time: 11:00AM
}

func ExampleFixedZone() {
	loc := time.FixedZone("UTC-8", -8*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t.Format(time.RFC822))
	// Output: The time is: 10 Nov 09 23:00 UTC-8
}
