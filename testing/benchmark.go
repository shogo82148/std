// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// InternalBenchmarkは内部の型ですが、他のパッケージからも利用できるように公開されています。
// これは"go test"コマンドの実装の一部です。
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}

// Bはベンチマーク関数に渡され、ベンチマークのタイミングや反復回数の管理を行う型です。
//
// ベンチマークは、そのBenchmark関数がreturnするか、[B.FailNow]、[B.Fatal]、[B.Fatalf]、[B.SkipNow]、[B.Skip]、[B.Skipf] のいずれかのメソッドを呼び出すことで終了します。
// これらのメソッドは、Benchmark関数を実行しているゴルーチンからのみ呼び出す必要があります。
// その他の報告用メソッド（[B.Log] や [B.Error] のバリエーションなど）は、複数のゴルーチンから同時に呼び出すことができます。
//
// テストと同様に、ベンチマークのログは実行中に蓄積され、終了時に標準出力に出力されます。ただし、ベンチマークのログは常に出力されるため、ベンチマーク結果に影響を与える可能性がある出力を隠すことはありません。
type B struct {
	common
	importPath       string
	bstate           *benchState
	N                int
	previousN        int
	previousDuration time.Duration
	benchFunc        func(b *B)
	benchTime        durationOrCountFlag
	bytes            int64
	missingBytes     bool
	timerOn          bool
	showAllocResult  bool
	result           BenchmarkResult
	parallelism      int
	// memStats.MallocsとmemStats.TotalAllocの初期状態。
	startAllocs uint64
	startBytes  uint64
	// 実行後のこのテストのネット合計。
	netAllocs uint64
	netBytes  uint64
	// ReportMetricによって収集される追加のメトリクス。
	extra map[string]float64

	// loop tracks the state of B.Loop
	loop struct {
		n uint64

		i uint64

		done bool
	}
}

// StartTimerはテストの計測を開始します。この関数はベンチマークが開始する前に自動的に呼び出されますが、[B.StopTimer] を呼び出した後に計測を再開するためにも使用することができます。
func (b *B) StartTimer()

// StopTimerはテストのタイミングを停止します。これは、計測したくない複雑な初期化を実行する間にタイマーを一時停止するために使用することができます。
func (b *B) StopTimer()

// ResetTimerは経過したベンチマーク時間とメモリ割り当てのカウンターをゼロにし、
// ユーザーが報告したメトリクスを削除します。
// タイマーが実行中かどうかには影響しません。
func (b *B) ResetTimer()

// SetBytesは単一の操作で処理されたバイト数を記録します。
// これが呼び出された場合、ベンチマークはns/opとMB/sを報告します。
func (b *B) SetBytes(n int64)

// ReportAllocsはこのベンチマークのためにmallocの統計情報を有効にします。
// これは-test.benchmemを設定するのと同じ効果ですが、ReportAllocsを呼び出すベンチマーク関数にのみ影響します。
func (b *B) ReportAllocs()

// Elapsedはベンチマークの計測された経過時間を返します。
// Elapsedによって報告される期間は、[B.StartTimer]、[B.StopTimer]、[B.ResetTimer]
// によって計測される期間と一致します。
func (b *B) Elapsed() time.Duration

// ReportMetricは報告されたベンチマーク結果に「n unit」を追加します。
// もしメトリックが反復ごとのものであれば、呼び出し元はb.Nで割る必要があります。
// また、単位は通常"/op"で終わるべきです。
// 同じ単位の以前の報告値は、ReportMetricによって上書きされます。
// unitが空の文字列または空白を含む場合、ReportMetricはパニックを起こします。
// unitが通常ベンチマークフレームワーク自体によって報告される単位である場合
// （例："allocs/op"）、ReportMetricはそのメトリックを上書きします。
// "ns/op"を0に設定すると、組み込まれたメトリックは抑制されます。
func (b *B) ReportMetric(n float64, unit string)

// Loopはベンチマークを継続して実行すべき間、trueを返します。
//
// 一般的なベンチマークの構造例：
//
//	func Benchmark(b *testing.B) {
//		... セットアップ ...
//		for b.Loop() {
//			... 計測対象のコード ...
//		}
//		... クリーンアップ ...
//	}
//
// Loopはベンチマーク内で最初に呼び出されたときにベンチマークタイマーをリセットします。
// そのため、ベンチマークループ開始前のセットアップ処理は計測に含まれません。
// 同様に、falseを返すときにタイマーを停止するため、クリーンアップ処理も計測されません。
//
// "for b.Loop() { ... }" ループの本体内では、ループ内で呼び出される関数の引数や戻り値が生存し続け、
// コンパイラによるループ本体の完全な最適化が防がれます。現在は、b.Loopループ内で呼び出される関数のインライン化を無効化することで実現されています。
// これはループの波括弧内で構文的に呼び出される関数にのみ適用され、ループ条件は必ず "b.Loop()" と記述する必要があります。
// ループ内から呼び出される関数では通常通り最適化が行われます。
//
// Loopがfalseを返した後、b.Nには実行された総イテレーション数が格納されるため、ベンチマークはb.Nを使って他の平均値を計算できます。
//
// Loop導入以前は、ベンチマークは0からb.Nまでの明示的なループを含む必要がありました。
// ベンチマークはLoopを使うかb.Nまでのループを含むか、どちらか一方にすべきです。
// Loopはベンチマークタイマーの管理をより自動化し、各ベンチマーク関数を計測ごとに一度だけ実行します。
// 一方、b.Nベースのベンチマークはベンチマーク関数（および関連するセットアップ・クリーンアップ）を複数回実行する必要があります。
func (b *B) Loop() bool

// BenchmarkResult contains the results of a benchmark run.
type BenchmarkResult struct {
	N         int
	T         time.Duration
	Bytes     int64
	MemAllocs uint64
	MemBytes  uint64

	// Extra records additional metrics reported by ReportMetric.
	// ExtraはReportMetricによって報告された追加のメトリクスを記録します。
	Extra map[string]float64
}

// NsPerOpは"ns/op"メトリックを返します。
func (r BenchmarkResult) NsPerOp() int64

// AllocsPerOpは「allocs/op」メトリックスを返します。
// このメトリックスはr.MemAllocs / r.Nで計算されます。
func (r BenchmarkResult) AllocsPerOp() int64

// AllocedBytesPerOpは「B/op」メトリックを返します。
// これはr.MemBytes / r.Nで計算されます。
func (r BenchmarkResult) AllocedBytesPerOp() int64

// Stringは、ベンチマーク結果の概要を返します。
// これは、https://golang.org/design/14313-benchmark-format からの
// ベンチマーク結果行の形式に従いますが、ベンチマーク名は含まれません。
// 追加のメトリクスは、同じ名前の組み込みメトリクスを上書きします。
// Stringは、allocs/opやB/opを含みません。これらは [BenchmarkResult.MemString] によって報告されます。
func (r BenchmarkResult) String() string

// MemStringは、'go test'と同じ形式でr.AllocedBytesPerOpとr.AllocsPerOpを返します。
func (r BenchmarkResult) MemString() string

// RunBenchmarksは内部関数ですが、クロスパッケージであるためにエクスポートされています。
// これは"go test"コマンドの実装の一部です。
func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)

// 指定された名前でサブベンチマークとしてベンチマークを実行します。 フェイルが発生したかどうかを報告します。
//
// サブベンチマークは他のどんなベンチマークとも同じです。 Runを少なくとも1回呼び出すベンチマークは自体は計測されず、N=1で1回呼び出されます。
func (b *B) Run(name string, f func(b *B)) bool

// PBはRunParallelによって並列ベンチマークの実行に使用されます。
type PB struct {
	globalN *atomic.Uint64
	grain   uint64
	cache   uint64
	bN      uint64
}

// Nextは、さらに実行するイテレーションがあるかどうかを返します。
func (pb *PB) Next() bool

// RunParallelは、ベンチマークを並行して実行します。
// 複数のゴルーチンを作成し、b.Nの反復をそれらの間で分散します。
// ゴルーチンの数はデフォルトでGOMAXPROCSです。CPUに依存しないベンチマークの並列性を
// 増加させるためには、RunParallelの前に[B.SetParallelism]を呼び出します。
// RunParallelは通常、go test -cpuフラグと一緒に使用されます。
//
// body関数は各ゴルーチンで実行されます。それは任意の
// ゴルーチンローカルの状態を設定し、その後pb.Nextがfalseを返すまで反復します。
// それは[B.StartTimer]、[B.StopTimer]、または[B.ResetTimer]関数を
// 使用すべきではありません、なぜならそれらはグローバルな影響を持つからです。また、[B.Run]を呼び出すべきでもありません。
//
// RunParallelは、ベンチマーク全体の壁時計時間（ns/op値）を報告します。これは並列ゴルーチンごとの壁時計時間またはCPU時間の合計ではありません。
func (b *B) RunParallel(body func(*PB))

// SetParallelismは、[B.RunParallel] によって使用されるゴルーチンの数をp*GOMAXPROCSに設定します。
// CPUに依存するベンチマークでは、通常SetParallelismを呼び出す必要はありません。
// pが1未満の場合、この呼び出しは効果がありません。
func (b *B) SetParallelism(p int)

// Benchmarkは単一の関数をベンチマークします。これは、"go test"コマンドを使用しない
// カスタムベンチマークを作成するのに便利です。
//
// もしfがテストフラグに依存しているなら、Benchmarkを呼び出す前と
// [flag.Parse] を呼び出す前に、それらのフラグを登録するために [Init] を使用する必要があります。
//
// もしfがRunを呼び出すなら、結果は単一のベンチマーク内で連続して
// Runを呼び出さないすべてのサブベンチマークを実行するための推定値になります。
func Benchmark(f func(b *B)) BenchmarkResult
