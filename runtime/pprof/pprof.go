// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pprofは、pprof視覚化ツールで期待される形式でランタイムプロファイリングデータを書き込みます。
// # Goプログラムのプロファイリング
//
// Goプログラムをプロファイリングする最初のステップは、プロファイリングを有効にすることです。
// 標準のテストパッケージでビルドされたベンチマークのプロファイリングをサポートするためには、go testに組み込まれています。
// たとえば、次のコマンドは現在のディレクトリでベンチマークを実行し、CPUプロファイルとメモリプロファイルをcpu.profとmem.profに書き込みます：
//
//	go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
//
// スタンドアロンのプログラムに同等のプロファイリングサポートを追加するには、以下のようなコードをmain関数に追加します：
//
//	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
//	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
//
//	func main() {
//	    flag.Parse()
//	    if *cpuprofile != "" {
//	        f, err := os.Create(*cpuprofile)
//	        if err != nil {
//	            log.Fatal("could not create CPU profile: ", err)
//	        }
//	        defer f.Close() // エラーハンドリングは例外です
//	        if err := pprof.StartCPUProfile(f); err != nil {
//	            log.Fatal("could not start CPU profile: ", err)
//	        }
//	        defer pprof.StopCPUProfile()
//	    }
//
//	    // ... プログラムの残り ...
//
//	    if *memprofile != "" {
//	        f, err := os.Create(*memprofile)
//	        if err != nil {
//	            log.Fatal("could not create memory profile: ", err)
//	        }
//	        defer f.Close() // エラーハンドリングは例外です
//	        runtime.GC() // 最新の統計情報を取得
//	        if err := pprof.WriteHeapProfile(f); err != nil {
//	            log.Fatal("could not write memory profile: ", err)
//	        }
//	    }
//	}
//
// プロファイリングデータへの標準のHTTPインターフェースもあります。以下の行を追加すると、/debug/pprof/の下にハンドラがインストールされ、ライブプロファイルをダウンロードすることができます：
//
//	import _ "net/http/pprof"
//
// 詳細については、net/http/pprofパッケージを参照してください。
// プロファイルはpprofツールで可視化することができます：
//
//	go tool pprof cpu.prof
//
// pprofコマンドラインからは多くのコマンドが利用できます。
// よく使用されるコマンドには、「top」（プログラムのホットスポットの要約を表示する）や、「web」（ホットスポットとその呼び出しグラフの対話型グラフを開く）があります。
// すべてのpprofコマンドに関する情報については、「help」を使用してください。
// pprofに関する詳細情報は、次を参照してください
// https://github.com/google/pprof/blob/master/doc/README.md.
package pprof

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// A Profile is a collection of stack traces showing the call sequences
// that led to instances of a particular event, such as allocation.
// Packages can create and maintain their own profiles; the most common
// use is for tracking resources that must be explicitly closed, such as files
// or network connections.
//
// A Profile's methods can be called from multiple goroutines simultaneously.
//
// Each Profile has a unique name. A few profiles are predefined:
//
//	goroutine    - stack traces of all current goroutines
//	heap         - a sampling of memory allocations of live objects
//	allocs       - a sampling of all past memory allocations
//	threadcreate - stack traces that led to the creation of new OS threads
//	block        - stack traces that led to blocking on synchronization primitives
//	mutex        - stack traces of holders of contended mutexes
//
// These predefined profiles maintain themselves and panic on an explicit
// [Profile.Add] or [Profile.Remove] method call.
//
// The heap profile reports statistics as of the most recently completed
// garbage collection; it elides more recent allocation to avoid skewing
// the profile away from live data and toward garbage.
// If there has been no garbage collection at all, the heap profile reports
// all known allocations. This exception helps mainly in programs running
// without garbage collection enabled, usually for debugging purposes.
//
// The heap profile tracks both the allocation sites for all live objects in
// the application memory and for all objects allocated since the program start.
// Pprof's -inuse_space, -inuse_objects, -alloc_space, and -alloc_objects
// flags select which to display, defaulting to -inuse_space (live objects,
// scaled by size).
//
// The allocs profile is the same as the heap profile but changes the default
// pprof display to -alloc_space, the total number of bytes allocated since
// the program began (including garbage-collected bytes).
//
// The CPU profile is not available as a Profile. It has a special API,
// the [StartCPUProfile] and [StopCPUProfile] functions, because it streams
// output to a writer during profiling.
type Profile struct {
	name  string
	mu    sync.Mutex
	m     map[any][]uintptr
	count func() int
	write func(io.Writer, int) error
}

// NewProfileは指定された名前で新しいプロファイルを作成します。
// すでにその名前のプロファイルが存在する場合、NewProfileはパニックを引き起こします。
// 各パッケージごとに別の名前空間を作成するために、'import/path.'接頭辞を使用するのが一般的です。
// pprofデータを読み取るさまざまなツールとの互換性のために、プロファイル名にはスペースを含めないでください。
func NewProfile(name string) *Profile

// Lookupは指定された名前のプロフィールを返します。存在しない場合はnilを返します。
func Lookup(name string) *Profile

// Profilesは、名前でソートされたすべてのプロフィールのスライスを返します。
func Profiles() []*Profile

<<<<<<< HEAD
// Nameはこのプロフィールの名前を返します。プロフィールを再取得するためにLookupに渡すことができます。
=======
// Name returns this profile's name, which can be passed to [Lookup] to reobtain the profile.
>>>>>>> upstream/master
func (p *Profile) Name() string

// Countは現在のプロファイル内の実行スタックの数を返します。
func (p *Profile) Count() int

<<<<<<< HEAD
// Addは現在の実行スタックを、値と関連付けてプロファイルに追加します。
// Addは値を内部のマップに保存するため、値はマップのキーとして使用するのに適しており、対応するRemove呼び出しまでガベージコレクトされません。
// Addはもしプロファイルにすでに値のスタックが含まれている場合、パニックを発生させます。
//
// skipパラメータはruntime.Callerのskipと同じ意味を持ち、スタックトレースが始まる場所を制御します。
// skip=0を渡すと、Addを呼び出した関数からトレースが始まります。例えば、以下のような実行スタックがあるとします:
=======
// Add adds the current execution stack to the profile, associated with value.
// Add stores value in an internal map, so value must be suitable for use as
// a map key and will not be garbage collected until the corresponding
// call to [Profile.Remove]. Add panics if the profile already contains a stack for value.
//
// The skip parameter has the same meaning as [runtime.Caller]'s skip
// and controls where the stack trace begins. Passing skip=0 begins the
// trace in the function calling Add. For example, given this
// execution stack:
>>>>>>> upstream/master
//
//	Add
//	rpc.NewClientから呼び出される
//	mypkg.Runから呼び出される
//	main.mainから呼び出される
//
// skip=0を渡すと、スタックトレースはrpc.NewClient内でのAddの呼び出しで始まります。
// skip=1を渡すと、スタックトレースはmypkg.Run内でのNewClientの呼び出しで始まります。
func (p *Profile) Add(value any, skip int)

// Removeはプロファイルから関連付けられた実行スタックを削除します。
// valueがプロファイルに存在しない場合、何もしません。
func (p *Profile) Remove(value any)

// WriteToはプロファイルのスナップショットをwにpprof形式で書き込みます。
// wへの書き込みがエラーを返す場合、WriteToはそのエラーを返します。
// それ以外の場合、WriteToはnilを返します。
//
// debugパラメータは追加の出力を有効にします。
// debug=0を渡すと、https://github.com/google/pprof/tree/master/proto#overviewで
// 説明されているgzip圧縮されたプロトコルバッファ形式で書き込まれます。
// debug=1を渡すと、関数名と行番号をアドレスに変換したレガシーテキスト形式で書き込まれます。
// これにより、プログラマがツールなしでプロファイルを読むことができます。
//
// プリセットのプロファイルは、他のdebugの値に意味を割り当てることができます。
// たとえば、"goroutine"プロファイルの場合、debug=2は、
// ゴルーチンのスタックをGoプログラムが回復不可能なパニックによって終了する際に使用する同じ形式で表示することを意味します。
func (p *Profile) WriteTo(w io.Writer, debug int) error

<<<<<<< HEAD
// WriteHeapProfileは、Lookup("heap").WriteTo(w, 0)の略記です。
// 後方互換性のために保持されています。
=======
// WriteHeapProfile is shorthand for [Lookup]("heap").WriteTo(w, 0).
// It is preserved for backwards compatibility.
>>>>>>> upstream/master
func WriteHeapProfile(w io.Writer) error

// StartCPUProfileは現在のプロセスに対してCPUプロファイリングを有効にします。
// プロファイリング中は、プロファイルがバッファリングされ、wに書き込まれます。
// StartCPUProfileは、すでにプロファイリングが有効な場合にエラーを返します。
//
<<<<<<< HEAD
// Unix系システムでは、-buildmode=c-archiveまたは-buildmode=c-sharedで
// ビルドされたGoコードでは、デフォルトではStartCPUProfileは動作しません。
// StartCPUProfileはSIGPROFシグナルを利用していますが、
// そのシグナルはGoが使用するものではなく、
// メインプログラムのSIGPROFシグナルハンドラ（あれば）に送信されます。
// 関数os/signal.Notifyをsyscall.SIGPROFに対して呼び出すことで、
// それが動作するようにすることができますが、その場合、
// メインプログラムで実行されているプロファイリングが壊れる可能性があります。
=======
// On Unix-like systems, StartCPUProfile does not work by default for
// Go code built with -buildmode=c-archive or -buildmode=c-shared.
// StartCPUProfile relies on the SIGPROF signal, but that signal will
// be delivered to the main program's SIGPROF signal handler (if any)
// not to the one used by Go. To make it work, call [os/signal.Notify]
// for [syscall.SIGPROF], but note that doing so may break any profiling
// being done by the main program.
>>>>>>> upstream/master
func StartCPUProfile(w io.Writer) error

// StopCPUProfileは現在のCPUプロファイルを停止します（もし存在する場合）。
// StopCPUProfileはプロファイルの書き込みが完了するまでのみ戻ります。
func StopCPUProfile()
