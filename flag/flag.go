// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package flagは、コマンドラインのフラグ解析を実装します。

# 使用法

flag.String（）、Bool（）、Int（）などを使用してフラグを定義します。

これは、ポインタnFlagに格納された整数フラグ-nを宣言し、型*intを持つものです。

	import "flag"
	var nFlag = flag.Int（"n"、1234、"フラグnのヘルプメッセージ"）

好きな場合、Var（）関数を使用してフラグを変数にバインドすることもできます。

	var flagvar int
	func init（）{
		flag.IntVar（&flagvar、"flagname"、1234、"フラグ名のヘルプメッセージ"）
	}

または、フラグを値インターフェースを満たすカスタムフラグに作成し、次にフラグ解析と結びつけることもできます（ポインタレシーバを使用）。

	flag.Var（&flagVal、「名前」、「flagnameのヘルプメッセージ」）

そのようなフラグの場合、デフォルト値は変数の初期値そのままです。

すべてのフラグが定義された後、次を呼び出して

	flag.Parse（）

定義されたフラグをコマンドラインで解析します。

その後、フラグを直接使用できます。フラグ自体を使用している場合、それらはすべてポインタです。変数にバインドする場合、値です。

	ipの値は、*ipを出力します。
	flagvarの値は、flagvarを出力します。

解析後、フラグに続く引数は、スライスflag.Args（）または個別にflag.Arg（i）として使用できます。
引数は0からflag.NArg（）-1までインデックス付けされます。

# コマンドラインフラグの句法

次の形式が許可されています。

	-flag
	--flag   // ダブルダッシュも許可されています
	-flag=x
	-flag x  // ブールフラグ以外

1つまたは2つのダッシュを使用できます。それらは同等です。
最後の形式は、0、falseなどという名前のファイルがある場合にはブールフラグでは許可されていません。

	cmd -x *

Unixシェルのワイルドカードである、Starは、「-」のコマンドの意味が変わるため、ブールフラグをオフにするには、-flag=false形式を使用する必要があります。

フラグ解析は、最初の非フラグ引数（「-」は非フラグ引数です）または終端子「--」の直前で停止します。

整数フラグは、1234、0664、0x1234などを受け入れ、負の値にすることもできます。
ブールフラグは以下のどれかです：

	1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False

期間フラグは、time.ParseDurationで有効な入力を受け入れます。

コマンドラインフラグのデフォルトセットは、トップレベルの関数によって制御されます。FlagSet型は、コマンドラインインターフェースのサブコマンドを実装するための独立したフラグセットを定義するために使用します。FlagSetのメソッドは、コマンドラインフラグセットのトップレベルの関数と同様です。
*/
package flag

import (
	"github.com/shogo82148/std/encoding"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

// ErrHelpは、-helpまたは-hフラグが呼び出されたが、そのようなフラグが定義されていない場合に返されるエラーです。
var ErrHelp = errors.New("flag: help requested")

// Valueはフラグに格納されている動的な値へのインターフェースです。
// （デフォルト値は文字列で表されます。）
//
// Valueがtrueを返すIsBoolFlag() boolメソッドを持つ場合、
// コマンドラインパーサーは-nameを-name=trueと等価にし、次のコマンドライン引数を使用しないようにします。
//
// flagが存在するごとに、Setメソッドがコマンドラインの順序で一度呼び出されます。
// flagパッケージはゼロ値のレシーバ（nilポインタなど）でStringメソッドを呼び出す場合があります。
type Value interface {
	String() string
	Set(string) error
}

// GetterはValueの内容を取得することを可能にするインターフェースです。
// これはGo 1以降の互換性規則のために、Valueインターフェースの一部ではなく、ラップされています。
// このパッケージで提供されるすべてのValue型は、Getterインターフェースを満たしますが、Funcによって使用される型は満たしません。
type Getter interface {
	Value
	Get() any
}

// ErrorHandlingは、parseが失敗した場合にFlagSet.Parseの動作を定義します。
type ErrorHandling int

// これらの定数は、パースが失敗した場合にFlagSet.Parseが説明された動作をするようにします。
const (
	ContinueOnError ErrorHandling = iota
	ExitOnError
	PanicOnError
)

// FlagSetは定義されたフラグの集合を表します。FlagSetのゼロ値は名前を持たず、ContinueOnErrorエラーハンドリングを持ちます。
// フラグの名前はFlagSet内でユニークでなければなりません。既に使用されている名前でフラグを定義しようとすると、パニックが発生します。
type FlagSet struct {

	// Usage はフラグの解析中にエラーが発生した場合に呼び出される関数です。
	// このフィールドは、カスタムのエラーハンドラを指すように変更できる関数（メソッドではありません）です。
	// Usage が呼び出された後に何が起こるかは、ErrorHandling の設定に依存します。
	// コマンドラインでは ExitOnError がデフォルトであり、Usage の呼び出し後にプログラムが終了します。
	Usage func()

	name          string
	parsed        bool
	actual        map[string]*Flag
	formal        map[string]*Flag
	args          []string
	errorHandling ErrorHandling
	output        io.Writer
	undef         map[string]string
}

// Flagはフラグの状態を表します。
type Flag struct {
	Name     string
	Usage    string
	Value    Value
	DefValue string
}

// Output は使用方法やエラーメッセージのための出力先を返します。 output が設定されていない場合や nil に設定されている場合は、os.Stderr が返されます。
func (f *FlagSet) Output() io.Writer

// Nameはフラグセットの名前を返します。
func (f *FlagSet) Name() string

// ErrorHandlingはフラグセットのエラーハンドリングの動作を返します。
func (f *FlagSet) ErrorHandling() ErrorHandling

// SetOutputは使用法やエラーメッセージの出力先を設定します。
// もしoutputがnilの場合、os.Stderrが使用されます。
func (f *FlagSet) SetOutput(output io.Writer)

// VisitAllは辞書順にフラグを訪れ、それぞれについてfnを呼び出します。
// 設定されていないフラグも含めて、すべてのフラグを訪れます。
func (f *FlagSet) VisitAll(fn func(*Flag))

// VisitAllはコマンドラインフラグを辞書順に訪れ、それぞれに対してfnを呼び出します。設定されていないフラグも含めて、すべてのフラグを訪れます。
func VisitAll(fn func(*Flag))

// Visitは辞書順にフラグを訪れ、それぞれに対してfnを呼び出します。
// 設定されているフラグのみを訪れます。
func (f *FlagSet) Visit(fn func(*Flag))

// Visitは辞書順でコマンドラインフラグを訪問し、各フラグに対してfnを呼び出します。
// 設定されたフラグのみを訪問します。
func Visit(fn func(*Flag))

// Lookupは指定されたフラグの構造体を返します。存在しない場合はnilを返します。
func (f *FlagSet) Lookup(name string) *Flag

// Lookupは指定されたコマンドラインフラグのFlag構造体を返します。存在しない場合はnilを返します。
func Lookup(name string) *Flag

// Setは指定したフラグの値を設定します。
func (f *FlagSet) Set(name, value string) error

// Setは名前付きのコマンドラインフラグの値を設定します。
func Set(name, value string) error

// UnquoteUsageは、フラグの使用法の場所から引用符で囲まれた名前を取り出し、
// それとその引用符を取り除いた使用法を返します。
// "a `name` to show"と与えられた場合、("name", "a name to show")を返します。
// もし引用符がない場合、その名前はフラグの値の型の教養された推測であり、もしフラグがブール値であれば空の文字列です。
func UnquoteUsage(flag *Flag) (name string, usage string)

// PrintDefaultsは、設定されていない限り、標準エラー出力に、セット内のすべての定義されたコマンドラインフラグのデフォルト値を表示します。詳細については、グローバル関数PrintDefaultsのドキュメントを参照してください。
func (f *FlagSet) PrintDefaults()

// PrintDefaultsは、標準エラー出力に、設定がない場合はデフォルトの設定を表示する使用方法メッセージを出力します。
// 整数値のフラグxに対して、デフォルトの出力形式は次のようになります。
//
//	-x int
//		xの使用方法メッセージ（デフォルト 7）
//
// 使用方法メッセージは、boolフラグ以外の場合には別の行に表示されます。boolフラグの場合は、型は省略され、フラグ名が1バイトの場合は使用方法メッセージが同じ行に表示されます。デフォルト値が型のゼロ値である場合、カッコ内のデフォルトは省略されます。ここではintと表示されていますが、フラグの使用方法文字列にバッククォートで囲まれた名前を記述することで、リストされる型を変更することができます。メッセージ中の最初のこのような項目が、メッセージ内で表示される引数の名前として扱われ、メッセージが表示される際にはバッククォートが剥がされます。例えば、以下のようにすると、
//
//	flag.String("I", "", "includeファイルを検索する`ディレクトリ`")
//
// 出力は次のようになります。
//
//	-I ディレクトリ
//		ディレクトリを検索するincludeファイル。
//
// フラグメッセージの出力先を変更するには、CommandLine.SetOutputを呼び出します。
func PrintDefaults()

// Usageは、CommandLineの出力（デフォルトでos.Stderr）に、定義されたすべてのコマンドラインフラグに関する使用法メッセージを出力します。
// フラグの解析中にエラーが発生したときに呼び出されます。
// この関数はカスタム関数を指すように変更できる変数です。
// デフォルトでは、簡単なヘッダーが表示され、PrintDefaultsが呼び出されます。
// 出力のフォーマットや、それを制御する方法の詳細については、PrintDefaultsのドキュメントを参照してください。
// カスタムのUsage関数ではプログラムを終了することも選択できますが、デフォルトでは終了は常に発生します。
// なぜなら、コマンドラインのエラーハンドリングストラテジーはExitOnErrorに設定されているからです。
var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	PrintDefaults()
}

// NFlagは設定されているフラグの数を返します。
func (f *FlagSet) NFlag() int

// NFlagは設定されたコマンドラインフラグの数を返します。
func NFlag() int

// Argはi番目の引数を返します。Arg(0)はフラグが処理された後の最初の残りの引数です。存在しない要素が要求された場合、Argは空の文字列を返します。
func (f *FlagSet) Arg(i int) string

// Argはi番目のコマンドライン引数を返します。Arg(0)は、フラグが処理された後の最初の残りの引数です。要求された要素が存在しない場合、Argは空の文字列を返します。
func Arg(i int) string

// NArgはフラグ処理後に残る引数の数です。
func (f *FlagSet) NArg() int

// NArgは、フラグが処理された後の残りの引数の数です。
func NArg() int

// Argsはフラグ以外の引数を返します。
func (f *FlagSet) Args() []string

// Argsはフラグではないコマンドライン引数を返します。
func Args() []string

// BoolVarは指定された名前、デフォルト値、および使用法の文字列を持つboolフラグを定義します。
// 引数pはフラグの値を格納するためのbool変数を指すポインタです。
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)

// BoolVarは指定された名前、デフォルト値、および使用方法のあるboolフラグを定義します。
// 引数pは、フラグの値を格納するbool変数を指すポインタです。
func BoolVar(p *bool, name string, value bool, usage string)

// Boolは、指定した名前、デフォルト値、使用法の説明でboolフラグを定義します。
// 戻り値は、フラグの値を格納するbool変数のアドレスです。
func (f *FlagSet) Bool(name string, value bool, usage string) *bool

// Boolは指定された名前、デフォルト値、使用方法の文字列を持つboolフラグを定義します。
// 返り値はフラグの値を格納するbool変数のアドレスです。
func Bool(name string, value bool, usage string) *bool

// IntVarは指定された名前、デフォルト値、使用法の文字列を持つintフラグを定義します。
// 引数pはフラグの値を格納するint変数を指すポインタです。
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)

// IntVarは指定された名前、デフォルト値、使用方法の文字列を持つintフラグを定義します。
// 引数pはフラグの値を格納するint変数を指すポインタです。
func IntVar(p *int, name string, value int, usage string)

// Intは指定された名前、デフォルト値、使用法の文字列を持つintフラグを定義します。
// 返り値は、フラグの値を格納するint変数のアドレスです。
func (f *FlagSet) Int(name string, value int, usage string) *int

// Intは指定した名前、デフォルト値、使用方法の文字列を持つintフラグを定義します。
// 返り値は、フラグの値を格納するint変数のアドレスです。
func Int(name string, value int, usage string) *int

// Int64Varは指定された名前、デフォルト値、使用方法の文字列を持つint64フラグを定義します。
// 引数pは、フラグの値を格納するint64変数を指すポインタです。
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)

// Int64Varは、指定された名前、デフォルト値、使用方法の文字列を持つint64フラグを定義します。
// 引数pは、フラグの値を格納するためのint64変数を指すポインタです。
func Int64Var(p *int64, name string, value int64, usage string)

// Int64は、指定された名前、デフォルト値、および使用法の文字列を持つint64フラグを定義します。
// 戻り値は、フラグの値を格納するint64変数のアドレスです。
func (f *FlagSet) Int64(name string, value int64, usage string) *int64

// Int64は指定された名前、デフォルト値、使用方法の文字列を持つint64フラグを定義します。
// 返り値は、フラグの値を保持するint64変数のアドレスです。
func Int64(name string, value int64, usage string) *int64

// UintVarは、指定された名前、デフォルト値、および使用方法の文字列を持つuintフラグを定義します。
// 引数pは、フラグの値を格納するuint変数を指すポインタです。
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)

// UintVarは指定された名前、デフォルト値、および使用法の文字列でuintフラグを定義します。
// 引数pは、フラグの値を格納するuint変数を指すポインタです。
func UintVar(p *uint, name string, value uint, usage string)

// Uintは指定した名前、デフォルトの値、および使用方法の文字列を持つuintフラグを定義します。
// 戻り値は、フラグの値を格納するuint変数のアドレスです。
func (f *FlagSet) Uint(name string, value uint, usage string) *uint

// Uintは指定された名前、デフォルト値、使用法の文字列を持つuintフラグを定義します。
// 返り値は、フラグの値を格納するuint変数のアドレスです。
func Uint(name string, value uint, usage string) *uint

// Uint64Varは、指定された名前、デフォルト値、使用方法の文字列を持つuint64フラグを定義します。
// 引数pは、フラグの値を格納するためのuint64変数を指すポインタです。
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)

// Uint64Varは指定された名前、デフォルト値、および使用法の文字列でuint64フラグを定義します。
// 引数pはフラグの値を格納するためのuint64変数を指すポインタです。
func Uint64Var(p *uint64, name string, value uint64, usage string)

// Uint64は指定された名前、デフォルト値、使用法のテキストでuint64のフラグを定義します。
// 返り値は、フラグの値を保持するuint64変数のアドレスです。
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64

// Uint64は指定された名前、デフォルト値、使用法の文字列を持つuint64フラグを定義します。
// 返り値は、フラグの値を格納するuint64変数のアドレスです。
func Uint64(name string, value uint64, usage string) *uint64

// StringVarは指定された名前、デフォルト値、および使用法の文字列を持つ文字列フラグを定義します。
// 引数pは、フラグの値を格納するための文字列変数を指すポインタです。
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)

// StringVarは、指定された名前、デフォルト値、使用法の文字列を持つ文字列フラグを定義します。
// 引数pは、フラグの値を格納する文字列変数を指すポインタです。
func StringVar(p *string, name string, value string, usage string)

// Stringは、指定された名前、デフォルト値、および使用法の文字列で文字列フラグを定義します。
// 返り値は、フラグの値を格納する文字列変数のアドレスです。
func (f *FlagSet) String(name string, value string, usage string) *string

// Stringは指定された名前、デフォルト値、使用方法のストリングフラグを定義します。
// 返り値は、フラグの値を保存する文字列変数のアドレスです。
func String(name string, value string, usage string) *string

// Float64Varは、指定した名前、デフォルト値、使用法の文字列を持つfloat64フラグを定義します。
// 引数pは、フラグの値を格納するfloat64変数を指すポインタです。
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)

// Float64Varは指定された名前、デフォルト値、使用法の文字列を持つfloat64フラグを定義します。
// 引数pはフラグの値を保存するためのfloat64変数を指すポインタです。
func Float64Var(p *float64, name string, value float64, usage string)

// Float64は指定された名前、デフォルト値、使用方法の文字列を持つfloat64フラグを定義します。
// 返り値は、フラグの値を格納するfloat64変数のアドレスです。
func (f *FlagSet) Float64(name string, value float64, usage string) *float64

// Float64は指定された名前、デフォルト値、および使用法の文字列を持つfloat64フラグを定義します。
// 戻り値は、フラグの値を格納するfloat64変数のアドレスです。
func Float64(name string, value float64, usage string) *float64

// DurationVarは、指定された名前、デフォルト値、使用方法を持つtime.Durationフラグを定義します。
// 引数pは、フラグの値を保存するためのtime.Duration変数を指すポインタです。
// このフラグは、time.ParseDurationで受け入れ可能な値を受け入れます。
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)

// DurationVarは、指定された名前、デフォルト値、使用方法の文字列を持つtime.Durationフラグを定義します。
// 引数pは、フラグの値を格納するためのtime.Duration変数を指すポインタです。
// フラグはtime.ParseDurationで受け入れ可能な値を受け付けます。
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)

// Durationは指定された名前、デフォルト値、および使用法の文字列を持つtime.Durationフラグを定義します。
// 戻り値は、フラグの値を格納するtime.Duration変数のアドレスです。
// このフラグは、time.ParseDurationが受け入れ可能な値を受け入れます。
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration

// Durationは指定された名前、デフォルト値、および使用法の文字列を持つtime.Durationフラグを定義します。
// 戻り値は、flagの値を格納するtime.Duration変数のアドレスです。
// このフラグは、time.ParseDurationで受け入れ可能な値を受け入れます。
func Duration(name string, value time.Duration, usage string) *time.Duration

// TextVarは指定された名前、デフォルト値、使用方法のフラグを定義します。
// 引数pは値を保持する変数へのポインタでなければならず、pはencoding.TextUnmarshalerを実装していなければなりません。
// フラグが使用された場合、フラグの値はpのUnmarshalTextメソッドに渡されます。
// デフォルト値の型はpの型と同じである必要があります。
func (f *FlagSet) TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string)

// TextVarは指定された名前、デフォルト値、使用方法を持つフラグを定義します。
// 引数pは値を保持する変数へのポインタでなければならず、pはencoding.TextUnmarshalerを実装しなければなりません。
// フラグが使用される場合、フラグの値はpのUnmarshalTextメソッドに渡されます。
// デフォルト値の型はpの型と同じでなければなりません。
func TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string)

// Funcは指定された名前と使用方法の文字列でフラグを定義します。
// フラグが見つかるたびに、fnがフラグの値で呼び出されます。
// fnが非nilのエラーを返す場合、それはフラグ値の解析エラーとして扱われます。
func (f *FlagSet) Func(name, usage string, fn func(string) error)

// Funcは指定された名前と使用法の文字列を持つフラグを定義します。
// フラグが見つかるたびに、fnがフラグの値で呼び出されます。
// もしfnが非nilのエラーを返す場合、それはフラグ値の解析エラーとして扱われます。
func Func(name, usage string, fn func(string) error)

// BoolFuncは値を必要とせず、指定された名前と使用方法の文字列でフラグを定義します。
// フラグが見つかる度に、fnがフラグの値と一緒に呼び出されます。
// もしfnが非nilのエラーを返した場合、それはフラグ値の解析エラーとして扱われます。
func (f *FlagSet) BoolFunc(name, usage string, fn func(string) error)

// BoolFuncは値を必要とせず、指定した名前と使用方法のフラグを定義します。
// フラグが表示されるたびに、fnがフラグの値で呼び出されます。
// もしfnが非nilのエラーを返した場合、それはフラグの値のパースエラーとして扱われます。
func BoolFunc(name, usage string, fn func(string) error)

// Varは指定された名前と使用方法のフラグを定義します。フラグの型と値は、通常、Valueという型の最初の引数で示され、Valueのユーザー定義の実装を保持します。例えば、呼び出し元は、Valueのメソッドを持つスライスにカンマ区切りの文字列を変換するフラグを作成することができます。特に、Setメソッドはカンマ区切りの文字列をスライスに分解します。
func (f *FlagSet) Var(value Value, name string, usage string)

// Varは指定された名前と使用方法のフラグを定義します。フラグの型と値は、通常はValue型の最初の引数で表されます。このValue型は一般的に、ユーザー定義のValue型の実装を保持します。たとえば、呼び出し側は、値のメソッドを持つスライスにコンマ区切りの文字列を変換するフラグを作成することができます。特に、Setはコンマ区切りの文字列をスライスに分解します。
func Var(value Value, name string, usage string)

// Parseは引数リストからフラグ定義を解析します。コマンド名は含まれていてはいけません。
// FlagSet内のすべてのフラグが定義され、プログラムによってフラグにアクセスされる前に呼び出す必要があります。
// 返り値は、-helpまたは-hが設定されているが定義されていない場合、ErrHelpになります。
func (f *FlagSet) Parse(arguments []string) error

// Parsedはf.Parseが呼ばれたかどうかを報告する。
func (f *FlagSet) Parsed() bool

// Parseはos.Args[1:]からコマンドラインフラグを解析します。全てのフラグが定義された後、プログラムによってフラグにアクセスされる前に呼び出す必要があります。
func Parse()

// Parsedは、コマンドラインフラグが解析されたかどうかを示します。
func Parsed() bool

// CommandLineはos.Argsから解析されたデフォルトのコマンドラインフラグのセットです。
// BoolVar、Argなどのトップレベルの関数は、CommandLineのメソッドのラッパーです。
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

// NewFlagSetは指定された名前とエラーハンドリングプロパティを持つ新しい空のフラグセットを返します。名前が空でない場合、デフォルトの使用方法メッセージとエラーメッセージに表示されます。
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet

// Initはフラグセットの名前とエラーハンドリングプロパティを設定します。
// デフォルトでは、ゼロ値のFlagSetは空の名前とContinueOnErrorのエラーハンドリングポリシーを使用します。
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
