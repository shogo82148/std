// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージlogはシンプルなロギングパッケージを実装します。出力のフォーマットに関するメソッドを持つ [Logger] という型を定義します。
// また、Print[f|ln]、Fatal[f|ln]、Panic[f|ln]というヘルパー関数を通じてアクセス可能な、事前定義された'standard' Loggerもあります。
// これらは、手動でLoggerを作成するよりも使いやすいです。そのロガーは標準エラーに書き込み、各ログメッセージの日付と時間を印刷します。
// すべてのログメッセージは別々の行に出力されます：印刷されるメッセージが改行で終わらない場合、ロガーは一つ追加します。
// Fatal関数は、ログメッセージを書き込んだ後に [os.Exit](1)を呼び出します。
// Panic関数は、ログメッセージを書き込んだ後にpanicを呼び出します。
package log

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
)

// これらのフラグは、[Logger] によって生成される各ログエントリの先頭にどのテキストを追加するかを定義します。
// ビットはor'ed（論理和）されて、何が印刷されるかを制御します。
// Lmsgprefixフラグを除いて、それらが表示される順序（ここにリストされている順序）や
// 形式（コメントで説明されているように）を制御する方法はありません。
// プレフィックスは、LlongfileまたはLshortfileが指定されたときにのみコロンに続きます。
// 例えば、フラグLdate | Ltime（またはLstdFlags）は、
//
//	2009/01/23 01:23:23 message
//
// フラグ Ldate | Ltime | Lmicroseconds | Llongfile は以下を生成します。
//
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	Lmsgprefix
	LstdFlags = Ldate | Ltime
)

// Loggerは、[io.Writer] に対して出力行を生成するアクティブなロギングオブジェクトを表します。
// 各ロギング操作は、WriterのWriteメソッドを一度だけ呼び出します。
// Loggerは複数のgoroutineから同時に使用することができます。これはWriterへのアクセスをシリアライズすることを保証します。
type Logger struct {
	outMu sync.Mutex
	out   io.Writer

	prefix    atomic.Pointer[string]
	flag      atomic.Int32
	isDiscard atomic.Bool
}

// Newは新しい [Logger] を作成します。out変数はログデータが書き込まれる先を設定します。
// プレフィックスは、生成された各ログ行の先頭に表示されるか、
// [Lmsgprefix] フラグが提供されている場合はログヘッダーの後に表示されます。
// flag引数はログのプロパティを定義します。
func New(out io.Writer, prefix string, flag int) *Logger

// SetOutputはロガーの出力先を設定します。
func (l *Logger) SetOutput(w io.Writer)

// Defaultは、パッケージレベルの出力関数で使用される標準ロガーを返します。
func Default() *Logger

// Outputはログイベントの出力を書き込みます。文字列sは、
// Loggerのフラグで指定されたプレフィックスの後に印刷するテキストを含みます。
// sの最後の文字がすでに改行でない場合、改行が追加されます。
// CalldepthはPCを回復するために使用され、一般性を提供しますが、
// 現時点ではすべての事前定義されたパスで2になります。
func (l *Logger) Output(calldepth int, s string) error

// Printはl.Outputを呼び出してロガーに出力します。
// 引数は [fmt.Print] の方法で処理されます。
func (l *Logger) Print(v ...any)

// Printfはl.Outputを呼び出してロガーに出力します。
// 引数は [fmt.Printf] の方法で処理されます。
func (l *Logger) Printf(format string, v ...any)

// Printlnはl.Outputを呼び出してロガーに出力します。
// 引数は [fmt.Println] の方法で処理されます。
func (l *Logger) Println(v ...any)

// Fatalはl.Print()の後に [os.Exit](1)を呼び出すのと同等です。
func (l *Logger) Fatal(v ...any)

// Fatalfはl.Printf()の後に [os.Exit](1)を呼び出すのと同等です。
func (l *Logger) Fatalf(format string, v ...any)

// Fatallnはl.Println()の後に [os.Exit](1)を呼び出すのと同等です。
func (l *Logger) Fatalln(v ...any)

// Panicはl.Print()の後にpanic()を呼び出すのと同等です。
func (l *Logger) Panic(v ...any)

// Panicfはl.Printf()の後にpanic()を呼び出すのと同等です。
func (l *Logger) Panicf(format string, v ...any)

// Paniclnはl.Println()の後にpanic()を呼び出すのと同等です。
func (l *Logger) Panicln(v ...any)

// Flagsはロガーの出力フラグを返します。
// フラグビットは [Ldate]、[Ltime] などです。
func (l *Logger) Flags() int

// SetFlagsはロガーの出力フラグを設定します。
// フラグビットは [Ldate]、[Ltime] などです。
func (l *Logger) SetFlags(flag int)

// Prefixはロガーの出力プレフィックスを返します。
func (l *Logger) Prefix() string

// SetPrefixはロガーの出力プレフィックスを設定します。
func (l *Logger) SetPrefix(prefix string)

// Writerはロガーの出力先を返します。
func (l *Logger) Writer() io.Writer

// SetOutputは標準ロガーの出力先を設定します。
func SetOutput(w io.Writer)

// Flagsは標準ロガーの出力フラグを返します。
// フラグビットは [Ldate]、[Ltime]などです。
func Flags() int

// SetFlagsは標準ロガーの出力フラグを設定します。
// フラグビットは [Ldate]、[Ltime] などです。
func SetFlags(flag int)

// Prefixは標準ロガーの出力プレフィックスを返します。
func Prefix() string

// SetPrefixは標準ロガーの出力プレフィックスを設定します。
func SetPrefix(prefix string)

// Writerは標準ロガーの出力先を返します。
func Writer() io.Writer

// PrintはOutputを呼び出して標準ロガーに出力します。
// 引数は [fmt.Print] の方法で処理されます。
func Print(v ...any)

// PrintfはOutputを呼び出して標準ロガーに出力します。
// 引数は [fmt.Printf] の方法で処理されます。
func Printf(format string, v ...any)

// PrintlnはOutputを呼び出して標準ロガーに出力します。
// 引数は [fmt.Println] の方法で処理されます。
func Println(v ...any)

// Fatalは [Print] の後に [os.Exit](1)を呼び出すのと同等です。
func Fatal(v ...any)

// Fatalfは [Printf] の後に [os.Exit](1)を呼び出すのと同等です。
func Fatalf(format string, v ...any)

// Fatallnは [Println] の後に [os.Exit](1)を呼び出すのと同等です。
func Fatalln(v ...any)

// Panicは [Print] の後にpanic()を呼び出すのと同等です。
func Panic(v ...any)

// Panicfは [Printf] の後にpanic()を呼び出すのと同等です。
func Panicf(format string, v ...any)

// Paniclnは [Println] の後にpanic()を呼び出すのと同等です。
func Panicln(v ...any)

// Outputはログイベントの出力を書き込みます。文字列sは、
// Loggerのフラグで指定されたプレフィックスの後に印刷するテキストを含みます。
// sの最後の文字がすでに改行でない場合、改行が追加されます。
// Calldepthは、[Llongfile] または [Lshortfile] が設定されている場合にファイル名と行番号を計算する際にスキップするフレームの数を表します。
// 値が1の場合、Outputの呼び出し元の詳細が印刷されます。
func Output(calldepth int, s string) error
