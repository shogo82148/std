// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"github.com/shogo82148/std/time"
)

// InternalBenchmarkは内部の型ですが、他のパッケージからも利用できるように公開されています。
// これは"go test"コマンドの実装の一部です。
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}

<<<<<<< HEAD
// Bはベンチマークのタイミングを管理し、実行する繰り返し回数を指定するためにベンチマーク関数に渡される型です。
=======
// B is a type passed to [Benchmark] functions to manage benchmark
// timing and to specify the number of iterations to run.
>>>>>>> upstream/master
//
// ベンチマーク関数がリターンするか、またはFailNow、Fatal、Fatalf、SkipNow、Skip、Skipfのいずれかのメソッドを呼び出すことでベンチマークは終了します。これらのメソッドはベンチマーク関数を実行しているゴルーチンからのみ呼び出す必要があります。
// ログやエラーのバリエーションといった他の報告用メソッドは、複数のゴルーチンから同時に呼び出すことができます。
//
// テストと同様に、ベンチマークのログは実行中に蓄積され、終了時に標準出力に出力されます。ただし、ベンチマークのログは常に出力されるため、ベンチマーク結果に影響を与える可能性がある出力を隠すことはありません。
type B struct {
	common
	importPath       string
	context          *benchContext
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
}

<<<<<<< HEAD
// StartTimerはテストの計測を開始します。この関数はベンチマークが開始する前に自動的に呼び出されますが、StopTimerを呼び出した後に計測を再開するためにも使用することができます。
=======
// StartTimer starts timing a test. This function is called automatically
// before a benchmark starts, but it can also be used to resume timing after
// a call to [B.StopTimer].
>>>>>>> upstream/master
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

<<<<<<< HEAD
// Elapsedはベンチマークの計測された経過時間を返します。
// Elapsedによって報告される期間は、StartTimer、StopTimer、ResetTimer
// によって計測される期間と一致します。
=======
// Elapsed returns the measured elapsed time of the benchmark.
// The duration reported by Elapsed matches the one measured by
// [B.StartTimer], [B.StopTimer], and [B.ResetTimer].
>>>>>>> upstream/master
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

// BenchmarkResultはベンチマークの実行結果を含んでいます。
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

<<<<<<< HEAD
// Stringはベンチマークの結果の概要を返します。
// これはhttps://golang.org/design/14313-benchmark-formatのベンチマーク結果行の形式に従います。
// ベンチマーク名を含めないでください。
// 追加のメトリクスは同じ名前の組み込みメトリクスを上書きします。
// allocs/opやB/opはMemStringによって報告されるため、Stringには含まれません。
=======
// String returns a summary of the benchmark results.
// It follows the benchmark result line format from
// https://golang.org/design/14313-benchmark-format, not including the
// benchmark name.
// Extra metrics override built-in metrics of the same name.
// String does not include allocs/op or B/op, since those are reported
// by [BenchmarkResult.MemString].
>>>>>>> upstream/master
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
	globalN *uint64
	grain   uint64
	cache   uint64
	bN      uint64
}

// Nextは、さらに実行するイテレーションがあるかどうかを返します。
func (pb *PB) Next() bool

<<<<<<< HEAD
// RunParallelはベンチマークを並列で実行します。
// 複数のゴルーチンを作成し、b.N回の反復をそれらのゴルーチンに分配します。
// ゴルーチンの数はデフォルトでGOMAXPROCSになります。CPUにバウンドしていないベンチマークの並列処理を増やすためには、RunParallelの前にSetParallelismを呼び出してください。
// RunParallelは通常、go test -cpuフラグと一緒に使用されます。
//
// body関数は各ゴルーチンで実行されます。これはゴルーチン固有の状態を設定し、pb.Nextがfalseを返すまで反復します。
// StartTimer、StopTimer、ResetTimer関数は使用しないでください。これらはグローバルな影響を持ちます。また、Runも呼び出さないでください。
=======
// RunParallel runs a benchmark in parallel.
// It creates multiple goroutines and distributes b.N iterations among them.
// The number of goroutines defaults to GOMAXPROCS. To increase parallelism for
// non-CPU-bound benchmarks, call [B.SetParallelism] before RunParallel.
// RunParallel is usually used with the go test -cpu flag.
//
// The body function will be run in each goroutine. It should set up any
// goroutine-local state and then iterate until pb.Next returns false.
// It should not use the [B.StartTimer], [B.StopTimer], or [B.ResetTimer] functions,
// because they have global effect. It should also not call [B.Run].
>>>>>>> upstream/master
//
// RunParallelは、ベンチマーク全体の壁時計時間（ns/op値）を報告します。これは並列ゴルーチンごとの壁時計時間またはCPU時間の合計ではありません。
func (b *B) RunParallel(body func(*PB))

<<<<<<< HEAD
// SetParallelismはRunParallelが使用するゴルーチンの数をp*GOMAXPROCSに設定します。
// 通常、CPUバウンドのベンチマークではSetParallelismを呼び出す必要はありません。
// pが1未満の場合、この呼び出しは効果がありません。
func (b *B) SetParallelism(p int)

// Benchmarkは単一の関数をベンチマークします。これは、"go test"コマンドを使用しないカスタムベンチマークの作成に役立ちます。
// もしfがテストフラグに依存している場合は、Benchmarkを呼び出す前とflag.Parseを呼び出す前に、Initを使用してこれらのフラグを登録する必要があります。
// もしfがRunを呼び出す場合、結果はRunを呼び出さないすべてのサブベンチマークをシーケンスで実行した場合の推定値となります。
=======
// SetParallelism sets the number of goroutines used by [B.RunParallel] to p*GOMAXPROCS.
// There is usually no need to call SetParallelism for CPU-bound benchmarks.
// If p is less than 1, this call will have no effect.
func (b *B) SetParallelism(p int)

// Benchmark benchmarks a single function. It is useful for creating
// custom benchmarks that do not use the "go test" command.
//
// If f depends on testing flags, then [Init] must be used to register
// those flags before calling Benchmark and before calling [flag.Parse].
//
// If f calls Run, the result will be an estimate of running all its
// subbenchmarks that don't call Run in sequence in a single benchmark.
>>>>>>> upstream/master
func Benchmark(f func(b *B)) BenchmarkResult
