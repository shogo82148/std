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

// Profileは、特定のイベント（例えば、割り当て）へのインスタンスにつながる呼び出しシーケンスを示すスタックトレースの集合です。
// パッケージは自身のプロファイルを作成し、維持することができます。最も一般的な使用例は、
// ファイルやネットワーク接続のような、明示的に閉じる必要があるリソースの追跡です。
//
// プロファイルのメソッドは、複数のゴルーチンから同時に呼び出すことができます。
//
// 各プロファイルには一意の名前があります。いくつかのプロファイルは事前に定義されています：
//
//	goroutine    - 現在のすべてのゴルーチンのスタックトレース
//	heap         - 生存しているオブジェクトのメモリ割り当てのサンプリング
//	allocs       - 過去のすべてのメモリ割り当てのサンプリング
//	threadcreate - 新しいOSスレッドの作成につながったスタックトレース
//	block        - 同期プリミティブでのブロックにつながったスタックトレース
//	mutex        - 競合するミューテックスの保持者のスタックトレース
//
<<<<<<< HEAD
// These predefined profiles maintain themselves and panic on an explicit
// [Profile.Add] or [Profile.Remove] method call.
//
// The CPU profile is not available as a Profile. It has a special API,
// the [StartCPUProfile] and [StopCPUProfile] functions, because it streams
// output to a writer during profiling.
//
// # Heap profile
=======
// これらの事前定義されたプロファイルは自身を維持し、明示的な
// AddまたはRemoveメソッド呼び出しでパニックします。
>>>>>>> release-branch.go1.21
//
// ヒーププロファイルは、最も最近に完了したガベージコレクション時点の統計を報告します。
// これは、プロファイルを生データからガベージに偏らせるのを避けるため、より最近の割り当てを省略します。
// ガベージコレクションが一度も行われていない場合、ヒーププロファイルはすべての既知の割り当てを報告します。
// この例外は主に、通常はデバッグ目的で、ガベージコレクションが有効になっていないプログラムで役立ちます。
//
// ヒーププロファイルは、アプリケーションメモリ内のすべてのライブオブジェクトの割り当て場所と、
// プログラム開始以降に割り当てられたすべてのオブジェクトを追跡します。
// Pprofの -inuse_space、-inuse_objects、-alloc_space、および -alloc_objects
// フラグは、表示するものを選択し、デフォルトは -inuse_space（ライブオブジェクト、サイズによってスケーリング）です。
//
<<<<<<< HEAD
// # Allocs profile
//
// The allocs profile is the same as the heap profile but changes the default
// pprof display to -alloc_space, the total number of bytes allocated since
// the program began (including garbage-collected bytes).
//
// # Block profile
//
// The block profile tracks time spent blocked on synchronization primitives,
// such as [sync.Mutex], [sync.RWMutex], [sync.WaitGroup], [sync.Cond], and
// channel send/receive/select.
//
// Stack traces correspond to the location that blocked (for example,
// [sync.Mutex.Lock]).
//
// Sample values correspond to cumulative time spent blocked at that stack
// trace, subject to time-based sampling specified by
// [runtime.SetBlockProfileRate].
//
// # Mutex profile
//
// The mutex profile tracks contention on mutexes, such as [sync.Mutex],
// [sync.RWMutex], and runtime-internal locks.
//
// Stack traces correspond to the end of the critical section causing
// contention. For example, a lock held for a long time while other goroutines
// are waiting to acquire the lock will report contention when the lock is
// finally unlocked (that is, at [sync.Mutex.Unlock]).
//
// Sample values correspond to the approximate cumulative time other goroutines
// spent blocked waiting for the lock, subject to event-based sampling
// specified by [runtime.SetMutexProfileFraction]. For example, if a caller
// holds a lock for 1s while 5 other goroutines are waiting for the entire
// second to acquire the lock, its unlock call stack will report 5s of
// contention.
//
// Runtime-internal locks are always reported at the location
// "runtime._LostContendedRuntimeLock". More detailed stack traces for
// runtime-internal locks can be obtained by setting
// `GODEBUG=runtimecontentionstacks=1` (see package [runtime] docs for
// caveats).
=======
// allocsプロファイルはヒーププロファイルと同じですが、デフォルトの
// pprof表示を -alloc_space（プログラム開始以降に割り当てられたバイト数の合計、
// ガベージコレクションされたバイトを含む）に変更します。
//
// CPUプロファイルはProfileとして利用できません。これは特別なAPIを持っており、
// StartCPUProfileとStopCPUProfile関数があります。これはプロファイリング中に
// 出力をライターにストリームします。
>>>>>>> release-branch.go1.21
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

// Nameはこのプロフィールの名前を返します。プロフィールを再取得するために [Lookup] に渡すことができます。
func (p *Profile) Name() string

// Countは現在のプロファイル内の実行スタックの数を返します。
func (p *Profile) Count() int

// Addは現在の実行スタックを、値と関連付けてプロファイルに追加します。
// Addは値を内部のマップに保存するため、値はマップのキーとして使用するのに適しており、対応する [Profile.Remove] 呼び出しまでガベージコレクトされません。
// Addはもしプロファイルにすでに値のスタックが含まれている場合、パニックを発生させます。
//
// skipパラメータは [runtime.Caller] のskipと同じ意味を持ち、スタックトレースが始まる場所を制御します。
// skip=0を渡すと、Addを呼び出した関数からトレースが始まります。例えば、以下のような実行スタックがあるとします:
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

// WriteHeapProfileは、[Lookup]("heap").WriteTo(w, 0)の略記です。
// 後方互換性のために保持されています。
func WriteHeapProfile(w io.Writer) error

// StartCPUProfileは現在のプロセスに対してCPUプロファイリングを有効にします。
// プロファイリング中は、プロファイルがバッファリングされ、wに書き込まれます。
// StartCPUProfileは、すでにプロファイリングが有効な場合にエラーを返します。
//
// Unix系システムでは、-buildmode=c-archiveまたは-buildmode=c-sharedで
// ビルドされたGoコードでは、デフォルトではStartCPUProfileは動作しません。
// StartCPUProfileはSIGPROFシグナルを利用していますが、
// そのシグナルはGoが使用するものではなく、
// メインプログラムのSIGPROFシグナルハンドラ（あれば）に送信されます。
// 関数 [os/signal.Notify] を [syscall.SIGPROF] に対して呼び出すことで、
// それが動作するようにすることができますが、その場合、
// メインプログラムで実行されているプロファイリングが壊れる可能性があります。
func StartCPUProfile(w io.Writer) error

// StopCPUProfileは現在のCPUプロファイルを停止します（もし存在する場合）。
// StopCPUProfileはプロファイルの書き込みが完了するまでのみ戻ります。
func StopCPUProfile()
