// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pprofパッケージは、pprof視覚化ツールで期待される形式でランタイムプロファイリングデータを書き込みます。
//
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
//	        defer f.Close() // エラーハンドリングは例のため省略
//	        runtime.GC() // 最新の統計情報を取得
//	        // Lookup("allocs")は go test -memprofile と同様のプロファイルを作成します。
//	        // または、Lookup("heap")を使うことで、
//	        // デフォルトのインデックスが inuse_space となるプロファイルを取得できます。
//	        if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
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
//
// https://github.com/google/pprof/blob/main/doc/README.md.
package pprof

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// Profileは、アロケーションなど特定のイベントのインスタンスに至った
// 呼び出しシーケンスを示すスタックトレースのコレクションです。
// パッケージは独自のプロファイルを作成・管理できます。最も一般的な
// 用途は、ファイルやネットワーク接続など明示的にクローズが必要なリソースの
// 追跡です。
//
// Profileのメソッドは複数のゴルーチンから同時に呼び出すことができます。
//
// 各Profileは一意の名前を持ちます。いくつかのプロファイルがあらかじめ定義されています：
//
//	goroutine      - 現在のすべてのゴルーチンのスタックトレース
//	goroutineleak  - リークしたすべてのゴルーチンのスタックトレース
//	allocs         - 過去のすべてのメモリアロケーションのサンプリング
//	heap           - 生きているオブジェクトのメモリアロケーションのサンプリング
//	threadcreate   - 新しいOSスレッドの作成に至ったスタックトレース
//	block          - 同期プリミティブでのブロックに至ったスタックトレース
//	mutex          - 競合したミューテックスの保持者のスタックトレース
//
// これらの定義済みプロファイルは自己管理され、明示的な
// [Profile.Add] または [Profile.Remove] メソッド呼び出しでパニックを起こします。
//
// CPUプロファイルはProfileとして利用できません。プロファイリング中に
// ライターに出力をストリーミングするため、[StartCPUProfile] と
// [StopCPUProfile] という特別なAPIを持ちます。
//
// # ヒーププロファイル
//
// ヒーププロファイルは直近に完了したガベージコレクション時点の統計を報告します。
// プロファイルが生きているデータから離れてガベージに偏らないよう、
// より最近のアロケーションを省略します。
// ガベージコレクションが一度も行われていない場合、ヒーププロファイルは
// 既知のすべてのアロケーションを報告します。この例外は主に、
// 通常デバッグ目的でガベージコレクションを無効にして実行しているプログラムで役立ちます。
//
// ヒーププロファイルは、アプリケーションメモリ内のすべての生きているオブジェクトの
// アロケーションサイトと、プログラム開始以降にアロケーションされたすべてのオブジェクトの
// 両方を追跡します。Pprofの -inuse_space、-inuse_objects、-alloc_space、および
// -alloc_objects フラグで表示内容を選択します。デフォルトは -inuse_space
// （生きているオブジェクト、サイズでスケーリング）です。
//
// # アロケーションプロファイル
//
// アロケーションプロファイルはヒーププロファイルと同じですが、デフォルトの
// pprof表示を -alloc_space（プログラム開始以降にアロケーションされた
// バイト数の合計（ガベージコレクションされたバイトを含む））に変更します。
//
// # ブロックプロファイル
//
// ブロックプロファイルは、[sync.Mutex]、[sync.RWMutex]、[sync.WaitGroup]、
// [sync.Cond]、チャネルの送受信/selectなど、同期プリミティブでブロックされた
// 時間を追跡します。
//
// スタックトレースはブロックが発生した場所（例：[sync.Mutex.Lock]）に対応します。
//
// サンプル値は、[runtime.SetBlockProfileRate] で指定された時間ベースのサンプリングに
// 従い、そのスタックトレースでブロックされた累積時間に対応します。
//
// # ミューテックスプロファイル
//
// ミューテックスプロファイルは、[sync.Mutex]、[sync.RWMutex]、
// およびランタイム内部ロックなどのミューテックスでの競合を追跡します。
//
// スタックトレースは競合を引き起こしたクリティカルセクションの終わりに対応します。
// たとえば、他のゴルーチンがロックの取得を待っている間、長時間保持されたロックは、
// そのロックが最終的にアンロックされたとき（つまり [sync.Mutex.Unlock] の時点）に
// 競合を報告します。
//
// サンプル値は、[runtime.SetMutexProfileFraction] で指定されたイベントベースの
// サンプリングに従い、他のゴルーチンがロックの取得を待ってブロックされた
// おおよその累積時間に対応します。たとえば、呼び出し元が1秒間ロックを保持し、
// 5つの他のゴルーチンがその1秒間ずっとロックの取得を待っている場合、
// アンロックの呼び出しスタックは5秒の競合を報告します。
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
