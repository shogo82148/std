// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"github.com/shogo82148/std/time"
)

// GCStatsは最近のガベージコレクションに関する情報を収集します。
type GCStats struct {
	LastGC         time.Time
	NumGC          int64
	PauseTotal     time.Duration
	Pause          []time.Duration
	PauseEnd       []time.Time
	PauseQuantiles []time.Duration
}

// ReadGCStatsはゴミ収集に関する統計情報をstatsに読み込みます。
// 一時停止履歴のエントリ数はシステムに依存します。
// stats.Pauseスライスは十分に大きければ再利用され、そうでなければ再割り当てされます。
// ReadGCStatsはstats.Pauseスライスの容量をフルに使用する可能性があります。
// もしstats.PauseQuantilesが空でない場合、ReadGCStatsは一時停止時間の分布を要約した
// 分位数をstats.PauseQuantilesに埋め込みます。
// 例えば、len(stats.PauseQuantiles)が5の場合、最小値、25%、50%、75%、最大値の一時停止時間が埋め込まれます。
func ReadGCStats(stats *GCStats)

// SetGCPercentはガベージコレクションの目標パーセンテージを設定します：
// 前回のコレクション後に残ったデータの生存データに対する最新割り当てデータの比率が
// このパーセンテージに達した時にコレクションがトリガーされます。
// SetGCPercentは以前の設定を返します。
// 初期設定は起動時のGOGC環境変数の値、または変数が設定されていない場合は100です。
// この設定はメモリ制限を維持するために効果的に減少させることができます。
// マイナスのパーセンテージは、メモリ制限が達成されない限り、ガベージコレクションを
// 実質的に無効にします。
// 詳細については、SetMemoryLimitを参照してください。
func SetGCPercent(percent int) int

// FreeOSMemoryはガベージコレクションを強制的に行い、可能なだけ多くのメモリをオペレーティングシステムに返す試みをします。
// (これが呼ばれなくても、ランタイムはバックグラウンドで徐々にメモリをオペレーティングシステムに返します。)
func FreeOSMemory()

// SetMaxStackは、個々のゴルーチンスタックが使用可能なメモリの最大量を設定します。
// スタックを成長させながらこの制限を超える場合、プログラムはクラッシュします。
// SetMaxStackは、前の設定を返します。
// 初期設定は、64ビットシステムでは1GB、32ビットシステムでは250MBです。
// SetMaxStackに提供される値に関係なく、システムによって設定された最大スタック制限がある場合があります。
//
// SetMaxStackは、無限再帰に入るゴルーチンによって引き起こされるダメージを制限するために主に役立ちます。
// これは将来のスタックの成長のみを制限します。
func SetMaxStack(bytes int) int

// SetMaxThreadsはGoプログラムが使用できるオペレーティングシステムの最大スレッド数を設定します。これ以上のスレッドを使用しようとすると、プログラムはクラッシュします。
// SetMaxThreadsは前の設定を返します。
// 初期設定は10,000スレッドです。
//
// 制限はオペレーティングシステムのスレッド数を制御しますが、ゴルーチンの数を制御しません。Goプログラムは、既存のすべてのスレッドがシステムコール、CGOコールにブロックされているか、runtime.LockOSThreadの使用により他のゴルーチンにロックされている場合にのみ新しいスレッドを作成します。
//
// SetMaxThreadsは、無制限の数のスレッドを作成するプログラムによる被害を制限するために主に役立ちます。アイデアは、プログラムがオペレーティングシステムをダウンさせる前にプログラム自体をダウンさせることです。
func SetMaxThreads(threads int) int

// SetPanicOnFaultは、プログラムが予期しない（非nil）アドレスでの障害が発生した場合、ランタイムの動作を制御します。
// このような障害は通常、ランタイムのメモリ破損などのバグによって引き起こされるため、デフォルトの応答はプログラムをクラッシュさせることです。
// メモリマップドファイルやメモリの安全でない操作を行うプログラムは、非nilアドレスでの障害を劇的な状況で引き起こすかもしれません。
// SetPanicOnFaultは、そのようなプログラムがクラッシュではなくパニックのみを要求できるようにします。
// ランタイムがパニックするときにランタイムがパニックするruntime.Errorには、追加のメソッドが存在する場合があります：
// Addr() uintptr
// もしAddrメソッドが存在する場合、それは障害を引き起こしたメモリアドレスを返します。
// Addrの結果はベストエフォートであり、結果の信頼性はプラットフォームに依存する可能性があります。
// SetPanicOnFaultは現在のゴルーチンにのみ適用されます。
// それは前の設定を返します。
func SetPanicOnFault(enabled bool) bool

// WriteHeapDumpはヒープとその中のオブジェクトの説明を指定されたファイルディスクリプタに書き込みます。
//
// WriteHeapDumpはヒープダンプが完全に書き込まれるまですべてのゴルーチンの実行を一時停止します。したがって、ファイルディスクリプタは、同じGoプロセスのもう一方のエンドがあるパイプやソケットに接続してはいけません。代わりに、一時ファイルまたはネットワークソケットを使用してください。
//
// ヒープダンプの形式はhttps://golang.org/s/go15heapdumpで定義されています。
func WriteHeapDump(fd uintptr)

// SetTracebackは、ランタイムがパニックや内部ランタイムエラーによる終了前に出力するトレースバックの詳細度を設定します。
// level引数は、GOTRACEBACK環境変数と同じ値を受け取ります。例えば、SetTraceback("all")は、プログラムがクラッシュしたときにすべてのゴルーチンを出力することを保証します。
// 詳細については、パッケージランタイムのドキュメントを参照してください。
// SetTracebackが環境変数よりも低いレベルで呼び出された場合、呼び出しは無視されます。
func SetTraceback(level string)

// SetMemoryLimitはランタイムにソフトメモリ制限を提供します。
//
// ランタイムは、ガベージコレクションの頻度の調整やメモリをより積極的に
// 下位システムに返却するなど、このメモリ制限を尊重するためにいくつかの
// プロセスを実行します。この制限は、GOGC=off (または、SetGCPercent(-1)が実行された場合も)であっても尊重されます。
//
// 入力制限はバイト単位で提供され、Goランタイムによってマップされた、管理された、リリースされていないすべてのメモリを含みます。特に、
// Goバイナリによって使用されるスペースや、Go以外のメモリ(プロセスの
// 代わりに下位システムによって管理されるメモリや、同じプロセス内の
// 非Goコードによって管理されるメモリなど)は考慮されません。除外される
// メモリソースの例には、プロセスのために保持されたOSカーネルメモリ、
// Cコードによって割り当てられたメモリ、syscall.Mmapによってマップされた
// メモリ(それはGoランタイムによって管理されていないため)が含まれます。
//
// より具体的には、次の式が制限としてランタイムが維持しようとする値を
// 正確に反映しています：
//
//	runtime.MemStats.Sys - runtime.MemStats.HeapReleased
//
// またはruntime/metricsパッケージを使用して：
//
//	/memory/classes/total:bytes - /memory/classes/heap/released:bytes
//
// ゼロの制限やGoランタイムによって使用されるメモリ量よりも低い制限は、
// ガベージコレクタがほぼ連続的に実行される原因となるかもしれません。
// ただし、アプリケーションは引き続き進行する可能性があります。
//
<<<<<<< HEAD
// メモリ制限は常にGoランタイムによって尊重されるため、この動作を無効にするには制限を非常に高い値に設定します。
// math.MaxInt64は制限を無効にするための規定値ですが、下位システムの利用可能なメモリよりもはるかに大きな値でも同様に機能します。
=======
// The memory limit is always respected by the Go runtime, so to
// effectively disable this behavior, set the limit very high.
// [math.MaxInt64] is the canonical value for disabling the limit,
// but values much greater than the available memory on the underlying
// system work just as well.
>>>>>>> upstream/master
//
// 詳細なガイドと共に、ソフトメモリ制限について詳しく説明したガイドおよび
// さまざまな一般的な使用例とシナリオについては、
// https://go.dev/doc/gc-guideを参照してください。
//
// 初期設定はmath.MaxInt64ですが、GOMEMLIMIT環境変数が設定されている場合は、初期設定を提供します。GOMEMLIMITは、バイト単位の数値で、
// オプションの単位接尾辞を持ちます。サポートされる接尾辞には、B、KiB、
// MiB、GiB、TiBがあります。これらの接尾辞は、IEC 80000-13規格で定義されるバイトの数量を表します。つまり、KiBは2^10バイトを、
// MiBは2^20バイトを意味し、それ以降も同様です。
//
// SetMemoryLimitは以前に設定されたメモリ制限を返します。
// 負の入力は制限を調整せず、現在設定されたメモリ制限を取得することができます。
func SetMemoryLimit(limit int64) int64
