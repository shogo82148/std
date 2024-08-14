// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージtimeは時間を測定し表示する機能を提供します。
//
// 暦の計算は常にグレゴリオ暦を前提としており、閏秒は考慮しません。
//
// # モノトニッククロック
//
// オペレーティングシステムは「壁掛け時計（wall clock）」と「モノトニッククロック（monotonic clock）」の2つを提供しています。壁掛け時計はクロック同期により変更される可能性がありますが、モノトニッククロックは変更されません。一般的なルールは、壁掛け時計は時刻を表示するために使用し、モノトニッククロックは時間を測定するために使用することです。このパッケージでは、[time.Now] が返すTimeには壁掛け時計の読み取り結果とモノトニッククロックの読み取り結果の両方が含まれています。後の時刻表示操作は壁掛け時計の読み取り結果を使用し、後の時間測定操作（比較や差分の計算など）はモノトニッククロックの読み取り結果を使用します。
//
// たとえ壁掛け時計が操作中に変更された場合でも、以下のコードは常に約20ミリ秒の経過時間を計算します。
//
//	start := time.Now()
//	... 20ミリ秒かかる処理 ...
//	t := time.Now()
//	elapsed := t.Sub(start)
//
// [time.Since](start)、[time.Until](deadline)、time.Now().Before(deadline)などの他のイディオムも、壁掛け時計のリセットに対して同様に頑健です。
//
// このセクションの残りの部分では、操作がモノトニッククロックを使用する方法の詳細を述べますが、これらの詳細を理解することはこのパッケージの使用には必要ありません。
//
// time.Nowが返すTimeにはモノトニッククロックの読み取り結果が含まれています。Time tがモノトニッククロックの読み取り結果を持つ場合、t.Addは壁掛け時計とモノトニッククロックの読み取り結果の両方に同じ期間を加算して結果を計算します。t.AddDate(y, m, d)、t.Round(d)、t.Truncate(d)は壁掛け時計の計算なので、結果からモノトニッククロックの読み取り結果は常に除去されます。t.In、t.Local、t.UTCは壁掛け時計の解釈への影響のために使用されますが、結果からモノトニッククロックの読み取り結果も常に除去されます。モノトニッククロックの読み取り結果を除去する正確な方法は、t = t.Round(0)を使用することです。
//
// tとuのいずれもモノトニッククロックの読み取り結果を含む場合、t.After(u)、t.Before(u)、t.Equal(u)、t.Compare(u)、t.Sub(u)は壁掛け時計の読み取り結果を無視してモノトニッククロックの読み取り結果だけを使用して実行されます。tまたはuのいずれかがモノトニッククロックの読み取り結果を含まない場合、これらの操作は壁掛け時計の読み取り結果を使用します。
//
// 一部のシステムでは、コンピュータがスリープ状態になるとモノトニッククロックが停止します。
// そのようなシステムでは、t.Sub(u) は t と u の間に実際に経過した時間を正確に反映しないかもしれません。
// [Since]、[Until]、[Before]、[After]、[Add]、[Sub]、[Equal]、[Compare] など、
// 時間を引き算する他の関数やメソッドにも同様のことが当てはまります。場合によっては、
// モノトニッククロックを取り除いて正確な結果を得る必要があるかもしれません。
//
// モノトニッククロックの読み取り結果には、現在のプロセスの外部では意味がありません。t.GobEncode、t.MarshalBinary、t.MarshalJSON、t.MarshalTextによって生成されるシリアル化された形式では、モノトニッククロックの読み取り結果は省略され、t.Formatはそれに対するフォーマットを提供しません。同様に、コンストラクタ [time.Date]、[time.Parse]、[time.ParseInLocation]、および [time.Unix]、およびアンマーシャラーt.GobDecode、t.UnmarshalBinary、t.UnmarshalJSON、およびt.UnmarshalTextは常にモノトニッククロックの読み取り結果のない時刻を作成します。
//
// モノトニッククロックの読み取り結果は [Time] の値にのみ存在します。[Duration] の値やt.Unixおよび関連する関数が返すUnix時刻には含まれていません。
//
// Goの==演算子は、時間の瞬間だけでなく、[Location] とモノトニッククロックの読み取り結果も比較します。Time型の等値テストについては、Time型のドキュメントを参照してください。
//
// デバッグ用に、t.Stringの結果には、存在する場合はモノトニッククロックの読み取りが含まれます。
// t != uの場合、異なるモノトニッククロックの読み取りによって、t.String()とu.String()の出力に差異が見られます。
//
// # タイマーの解像度
//
// [Timer] の解像度は、Goランタイム、オペレーティングシステム、
// および基礎となるハードウェアによって異なります。
// Unixでは、解像度は約1msです。
// Windowsバージョン1803以降では、解像度は約0.5msです。
// 古いWindowsバージョンでは、デフォルトの解像度は約16msですが、
// [golang.org/x/sys/windows.TimeBeginPeriod] を使用して高解像度を要求することができます。
package time

// Timeは納秒単位の精度で時刻を表します。
//
// Timeを使用するプログラムでは通常、値として保存し、渡すべきです。
// つまり、時刻変数や構造体のフィールドは*time.Timeではなく [time.Time] の型であるべきです。
//
// Timeの値は、[Time.GobDecode]、[Time.UnmarshalBinary]、[Time.UnmarshalJSON]、[Time.UnmarshalText] メソッドを除いて、
// 複数のゴルーチンで同時に使用できます。
//
// Timeの瞬間は [Time.Before]、[Time.After]、[Time.Equal] メソッドを使って比較することができます。
// [Time.Sub] メソッドは2つの瞬間を引いて [Duration] を生成します。
// [Time.Add] メソッドはTimeとDurationを足してTimeを生成します。
//
// Time型のゼロ値は、UTCでの1年1月1日00:00:00.000000000です。
// この時刻は実際にはほとんど使われないため、[Time.IsZero] メソッドは明示的に初期化されていない時刻を検出するための簡単な方法です。
//
// 各時刻には関連する [Location] があります。[Time.Local]、[Time.UTC]、および Time.In メソッドは、特定のLocationを持つTimeを返します。
// これらのメソッドを使用してTime値のLocationを変更しても、それが表す実際の瞬間は変更されず、解釈するタイムゾーンのみが変更されます。
//
// [Time.GobEncode]、[Time.MarshalBinary]、[Time.MarshalJSON]、[Time.MarshalText] メソッドによって保存されるTime値の表現には、[Time.Location] のオフセットが格納されますが、
// 場所の名前は格納されません。そのため、夏時間に関する情報が失われます。
//
// 必要な「壁時計」の読み取りに加えて、Timeにはオプションのプロセスの単調な時計の読み取りが含まれることがあります。
// 比較や減算のための追加の精度を提供するためです。
// 詳細については、パッケージのドキュメントの「単調な時計」のセクションを参照してください。
//
// Goの==演算子は、時刻の瞬間だけでなく、Locationと単調な時計の読み取りも比較します。
// そのため、Timeの値は、まずすべての値に同じLocationが設定されていることを保証してから、
// マップやデータベースのキーとして使用するべきではありません。これはUTCまたはLocalメソッドの使用によって実現でき、
// 単調な時計の読み取りはt = t.Round(0)と設定することで削除する必要があります。
// 一般的には、t.Equal(u)をt == uよりも優先し、t.Equalは最も正確な比較を使用し、
// 1つの引数のみが単調な時計の読み取りを持つ場合のケースを正しく処理します。
type Time struct {

	// wallとextは、壁時の秒、壁時のナノ秒、およびオプションの単調クロックの読み取り値（ナノ秒単位）をエンコードします。
	// 上位から下位のビット位置に従い、wallは1ビットのフラグ（hasMonotonic）、33ビットの秒フィールド、および30ビットの壁時のナノ秒フィールドをエンコードします。
	// ナノ秒のフィールドの範囲は[0, 999999999]です。
	// hasMonotonicビットが0の場合、33ビットのフィールドはゼロでなければならず、完全な符号付きの64ビットの壁秒は、Jan 1 year 1からの時間がextに格納されます。
	// hasMonotonicビットが1の場合、33ビットのフィールドは1885年1月1日以降の33ビットの符号なし壁秒を保持し、extはプロセスの開始からの時間が符号付きの64ビットの単調クロックの読み取り値（ナノ秒単位）を保持します。
	wall uint64
	ext  int64

	// locはTimeに対応する分、時、月、日、年を決定するために使用されるべき場所を指定します。
	// nilの場所はUTCを意味します。
	// すべてのUTC時刻はloc==nilで表され、loc==&utcLocではありません。
	loc *Location
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool

// tがuより後の時刻であるかどうかを報告します。
func (t Time) After(u Time) bool

// Beforeは、時刻tがuよりも前であるかどうかを報告します。
func (t Time) Before(u Time) bool

// Compare関数は時刻tとuを比較します。tがuより前の場合は-1を返し、tがuより後の場合は+1を返します。同じである場合は0を返します。
func (t Time) Compare(u Time) int

// Equalは、tとuが同じ瞬間を表しているかどうかを報告します。
// 異なる場所にある場合でも、2つの時刻が等しい場合があります。
// たとえば、6:00 +0200と4:00 UTCはEqualです。
// Timeのドキュメントを参照して、==を使用する際の注意点を確認してください。
// ほとんどのコードはEqualを使用する必要があります。
func (t Time) Equal(u Time) bool

// Monthは年の月を指定します（1月= 1、...）。
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// Stringは月の英語名（「January」、「February」、...）を返します。
func (m Month) String() string

// Weekdayは、週の日を指定します（日曜日 = 0、...）。
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Stringは日の英語名（"Sunday", "Monday", ...）を返します。
func (d Weekday) String() string

// Dateはtが発生する年、月、日を返します。
func (t Time) Date() (year int, month Month, day int)

// Yearはtが発生する年を返します。
func (t Time) Year() int

// Month は t で指定された年の月を返します。
func (t Time) Month() Month

// Dayはtによって指定された月の日を返します。
func (t Time) Day() int

// Weekdayはtに指定された曜日を返します。
func (t Time) Weekday() Weekday

// ISOWeekは、tが発生するISO 8601の年と週番号を返します。
// 週は1から53までの範囲です。年nの1月1日から1月3日は年n-1の週52または53に属する場合があり、12月29日から12月31日は年n+1の週1に属する場合があります。
func (t Time) ISOWeek() (year, week int)

// Clock 関数は、t で指定された日の時間、分、秒を返します。
func (t Time) Clock() (hour, min, sec int)

// Hourはtで指定された日の中の時間を返します。範囲は[0, 23]です。
func (t Time) Hour() int

// Minuteは時間tで指定された時間内の分のオフセットを、[0、59]の範囲で返します。
func (t Time) Minute() int

// Secondは、tで指定された分の中の2番目のオフセットを、[0, 59]の範囲で返します。
func (t Time) Second() int

// Nanosecondはtで指定された秒の中でのナノ秒オフセットを[0, 999999999]の範囲で返します。
func (t Time) Nanosecond() int

// YearDay は与えられた t によって指定される年の日を返します。
// 非閏年の場合、範囲は [1, 365] であり、閏年の場合は [1, 366] です。
func (t Time) YearDay() int

// Durationは2つの瞬間の経過時間をint64ナノ秒カウントとして表します。
// この表現は、最大表現可能な期間を約290年に制限しています。
type Duration int64

// 共通の期間です。Dayやそれ以上の単位の定義はありません。
// 夏時間ゾーンの移行時に混乱を避けるためです。
//
// [Duration] の単位数を数えるには、次のように割ります：
//
//	second := time.Second
//	fmt.Print(int64(second/time.Millisecond)) // 1000と出力されます
//
// 整数の単位数をDurationに変換するには、次のように掛けます：
//
//	seconds := 10
//	fmt.Print(time.Duration(seconds)*time.Second) // 10sと出力されます
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

// Stringは、期間を「72h3m0.5s」という形式の文字列で返します。
// 先頭のゼロ単位は省略されます。特別なケースとして、1秒未満の期間は、先頭の数字が0以外になるように、より小さな単位（ミリ秒、マイクロ秒、ナノ秒）を使用してフォーマットされます。ゼロ期間は「0s」としてフォーマットされます。
func (d Duration) String() string

// Nanosecondsは、期間を整数のナノ秒数として返します。
func (d Duration) Nanoseconds() int64

// Microsecondsは、整数のマイクロ秒数として期間を返します。
func (d Duration) Microseconds() int64

// Millisecondsは、期間を整数のミリ秒カウントとして返します。
func (d Duration) Milliseconds() int64

// Secondsは、秒単位の浮動小数点数としての期間を返します。
func (d Duration) Seconds() float64

// Minutesは、分単位の浮動小数点数としての期間を返します。
func (d Duration) Minutes() float64

// Hoursは時間を浮動小数点数として返します。
func (d Duration) Hours() float64

// Truncate は d を 0 に向かって丸めた結果を、m の倍数にして返します。
// もし m <= 0 であれば、Truncate は d をそのまま返します。
func (d Duration) Truncate(m Duration) Duration

// Round関数は、dを最も近いmの倍数に丸めた結果を返します。
// 半分の値の丸め方は、ゼロから離れるように丸めます。
// 結果が [Duration] に格納できる最大（または最小）値を超える場合、
// Round関数は最大（または最小）のdurationを返します。
// m <= 0の場合、Round関数はdをそのまま返します。
func (d Duration) Round(m Duration) Duration

// Absはdの絶対値を返します。
// 特殊なケースとして、Duration([math.MinInt64])はDuration([math.MaxInt64])に変換され、
// その大きさが1ナノ秒減少します。
func (d Duration) Abs() Duration

// Addは時間tにdを加えた時間を返します。
func (t Time) Add(d Duration) Time

// Subは期間t-uを返します。結果が [Duration] に格納できる最大値（または最小値）を超える場合、最大（または最小）の期間が返されます。
// 期間dのt-dを計算するためには、t.Add(-d)を使用してください。
func (t Time) Sub(u Time) Duration

// Since(t)は、tから経過した時間を返します。
// これはtime.Now().Sub(t)の省略形です。
func Since(t Time) Duration

// Untilはtまでの時間を返します。
// これはt.Sub(time.Now())の省略形です。
func Until(t Time) Duration

// AddDateはtに指定された年数、月数、日数を加算した時間を返します。
// 例えば、2011年1月1日にAddDate(-1, 2, 3)を適用すると、
// 2010年3月4日が返されます。
//
// 日付は基本的にタイムゾーンに結び付けられており、日などの暦期間には固定された期間がありません。
// AddDateは、Time値のLocationを使用してこれらの期間を決定します。
// つまり、同じAddDate引数でも、基本となるTime値とそのLocationに応じて、絶対時間のシフトが異なる場合があります。
// たとえば、3月27日の12:00に適用されるAddDate(0、0、1)は常に3月28日の12:00を返します。
// 一部の場所や年では、これは24時間のシフトです。他の場所や年では、夏時間の移行により23時間のシフトです。
//
// AddDateはDateと同じように結果を正規化します。
// つまり、10月31日に1ヶ月を追加すると11月31日となり、
// これは正規化されたフォームである12月1日になります。
func (t Time) AddDate(years int, months int, days int) Time

// 現在のローカル時間を返します。
func Now() Time

// UTCは場所をUTCに設定したtを返します。
func (t Time) UTC() Time

// Local は t の位置をローカルの時間に設定して返します。
func (t Time) Local() Time

// In は、同じ時刻インスタンスを表す t のコピーを返しますが、表示目的でコピーの場所情報を loc に設定します。
//
// loc が nil の場合、In はパニックを発生させます。
func (t Time) In(loc *Location) Time

// Locationは、tに関連付けられたタイムゾーン情報を返します。
func (t Time) Location() *Location

// Zone関数は、時刻tに有効なタイムゾーンを計算し、タイムゾーンの省略名（「CET」など）とUTCから東へのオフセットを秒単位で返します。
func (t Time) Zone() (name string, offset int)

// ZoneBoundsは、時刻tに有効なタイムゾーンの範囲を返します。
// ゾーンはstartで始まり、次のゾーンはendで始まります。
// ゾーンが時間の開始時点で始まる場合、startはゼロの時間として返されます。
// ゾーンが永遠に続く場合、endはゼロの時間として返されます。
// 返された時間の場所はtと同じになります。
func (t Time) ZoneBounds() (start, end Time)

// UnixはtをUnix時刻として返します。これは1970年1月1日UTCからの経過秒数です。結果はtに関連付けられた場所に依存しません。
// Unix系のオペレーティングシステムはしばしば時間を32ビットの秒数として記録しますが、ここで返される値は64ビットのため、過去や未来の数十億年に対しても有効です。
func (t Time) Unix() int64

// UnixMilliはtをUnix時刻として返します。つまり、1970年1月1日UTCから経過したミリ秒の数です。int64で表現できない（1970年より292万年以上前または後の日付）場合、結果は未定義です。結果はtに関連付けられた場所に依存しません。
func (t Time) UnixMilli() int64

// UnixMicroは、tをUnix時間として返します。これは、1970年1月1日UTCから経過したマイクロ秒の数です。int64で表現できない場合（年-290307以前または年294246以降の日付）、結果は未定義です。結果は、tに関連付けられた場所に依存しません。
func (t Time) UnixMicro() int64

// UnixNanoはtをUnix時刻として返します。これは、1970年1月1日UTCから経過したナノ秒数です。Unix時刻がint64で表現できない場合（1678年以前または2262年以降の日付）、結果は未定義です。なお、これはゼロのTimeに対してUnixNanoを呼び出した結果も未定義であることを意味します。結果はtに関連付けられた場所に依存しません。
func (t Time) UnixNano() int64

// MarshalBinaryはencoding.BinaryMarshalerインターフェースを実装します。
func (t Time) MarshalBinary() ([]byte, error)

// UnmarshalBinaryはencoding.BinaryUnmarshalerインターフェースを実装します。
func (t *Time) UnmarshalBinary(data []byte) error

// GobEncodeはgob.GobEncoderインターフェースを実装します。
func (t Time) GobEncode() ([]byte, error)

// GobDecodeはgob.GobDecoderインターフェースを実装します。
func (t *Time) GobDecode(data []byte) error

// MarshalJSONは [encoding/json.Marshaler] インターフェースを実装します。
// 時刻はRFC 3339形式で引用符で囲まれた文字列です。秒未満の精度もあります。
// タイムスタンプが有効なRFC 3339として表現できない場合
// （例：年が範囲外の場合）、エラーが報告されます。
func (t Time) MarshalJSON() ([]byte, error)

// UnmarshalJSONは [encoding/json.Unmarshaler] インターフェースを実装します。
// 時刻はRFC 3339形式でクォートされた文字列である必要があります。
func (t *Time) UnmarshalJSON(data []byte) error

// MarshalTextは [encoding.TextMarshaler] インターフェースを実装します。
// 時間はRFC 3339形式でサブ秒の精度でフォーマットされます。
// タイムスタンプが有効なRFC 3339として表現できない場合（例：年が範囲外の場合）、エラーが報告されます。
func (t Time) MarshalText() ([]byte, error)

// UnmarshalTextは [encoding.TextUnmarshaler] インターフェースを実装します。
// 時刻はRFC 3339形式である必要があります。
func (t *Time) UnmarshalText(data []byte) error

// Unixは、与えられたUnix時刻に対応する現地時刻を返します。
// secは1970年1月1日UTCからの秒数であり、nsecはナノ秒単位です。
// nsecは[0、999999999]の範囲外を指定することもできます。
// すべてのsec値に対応する時間値が存在するわけではありません。そのような値の一つは1 << 63-1（最大のint64値）です。
func Unix(sec int64, nsec int64) Time

// UnixMilliは与えられたUnix時刻に対応するローカルの時刻を返します。
// 1970年1月1日UTCからのミリ秒で表されます。
func UnixMilli(msec int64) Time

// UnixMicroは、与えられたUnix時間に対応するローカル時間を返します。
// usecは1970年1月1日UTCからのマイクロ秒です。
func UnixMicro(usec int64) Time

// IsDSTは、設定された場所の時刻が夏時間（Daylight Savings Time）かどうかを報告します。
func (t Time) IsDST() bool

// Date関数は、与えられた場所と時間に対応するTimeZoneでの、 yyyy-mm-dd hh:mm:ss + nsecナノ秒のTimeを返します。
//
// 月、日、時、分、秒、nsecの値は通常の範囲外である場合でも、変換中に正規化されます。
// 例えば、10月32日は11月1日に変換されます。
//
// 夏時間の変遷では、一部の時間が飛ばされたり、繰り返されることがあります。
// 例えば、アメリカ合衆国では、2011年3月13日の2時15分は存在せず、
// 2011年11月6日の1時15分は2回存在します。そのような場合、
// タイムゾーンの選択、つまり時間が明確でなくなります。
// Date関数は、変遷に関与する2つのタイムゾーンのうちの1つで正確な時間を返しますが、
// どちらかは保証されません。
//
// locがnilの場合、Date関数はパニックを発生させます。
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

// Truncate関数は、tをd（ゼロ時点からの倍数）に切り捨てた結果を返します。
// もしd <= 0の場合、Truncateはmonotonicなクロックの読み取りを除いたtを返しますが、それ以外は変更されません。
//
// Truncateは、時間をゼロ時点からの絶対経過時間として操作します。
// 時間の表現形式ではなく、時刻を操作します。したがって、Truncate(Hour)は、
// 時刻のLocationに依存して、非ゼロの分を持つ時刻を返す場合があります。
func (t Time) Truncate(d Duration) Time

// Roundはtを最も近いdの倍数に丸めた結果を返します（ゼロ時からの経過時間を使います）。
// 半分の値の丸め方は切り上げです。
// d <= 0の場合、Roundはtをモノトニックな時計読み取りを除いて変更せずに返します。
//
// Roundは時間の表示形式ではなく、ゼロ時からの絶対経過時間として動作します。
// したがって、Round(Hour)は時間のLocationによって非ゼロの分を持つことがあります。
func (t Time) Round(d Duration) Time
