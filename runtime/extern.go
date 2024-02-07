// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージruntimeには、goroutineを制御するための関数など、Goのランタイムシステムとやり取りする操作が含まれています。
また、reflectパッケージで使用されるランタイムタイプシステムのプログラマブルインターフェースに関するドキュメントについては、reflectのドキュメントを参照してください。

# 環境変数

以下の環境変数（ホストオペレーティングシステムによっては$nameまたは%name%）は、Goプログラムのランタイム動作を制御します。意味や使用法はリリースごとに変更される可能性があります。

GOGC変数は、初期のガベージコレクションのターゲットパーセンテージを設定します。
前回のコレクション後に残されたライブデータに対する新しく割り当てられたデータの比率がこのパーセンテージに達すると、コレクションがトリガーされます。
デフォルトはGOGC=100です。GOGC=offに設定すると、ガベージコレクタが完全に無効になります。
[runtime/debug.SetGCPercent] を使用すると、このパーセンテージを実行時に変更できます。

GOMEMLIMIT変数は、ランタイムのソフトメモリ制限を設定します。
このメモリ制限には、Goヒープとランタイムによって管理されるすべてのその他のメモリが含まれますが、
バイナリ自体のマッピング、他の言語で管理されるメモリ、およびGoプログラムの代わりにオペレーティングシステムに保持されるメモリなど、外部メモリソースは除外されます。
GOMEMLIMITは、オプションの単位接尾辞を持つバイト単位の数値です。
サポートされる接尾辞には、B、KiB、MiB、GiB、およびTiBが含まれます。
これらの接尾辞は、IEC 80000-13標準で定義されるバイトの量を表します。
つまり、2の累乗に基づいています：KiBは2^10バイトを意味し、MiBは2^20バイトを意味します。
デフォルト設定はmath.MaxInt64であり、これによりメモリ制限が無効になります。
[runtime/debug.SetMemoryLimit] を使用すると、この制限を実行時に変更できます。

GODEBUG変数は、ランタイム内のデバッグ変数を制御します。
これは、これらの名前付き変数を設定するname=valペアのカンマ区切りリストです。

	allocfreetrace: allocfreetrace=1を設定すると、すべての割り当てがプロファイルされ、各オブジェクトの割り当てと解放時にスタックトレースが出力されます。

	clobberfree: clobberfree=1を設定すると、ガベージコレクタがオブジェクトを解放するときに、オブジェクトのメモリ内容を悪い内容で上書きします。

	cpu.*: cpu.all=offは、すべてのオプションの命令セット拡張機能の使用を無効にします。
	cpu.extension=offは、指定された命令セット拡張機能からの命令の使用を無効にします。
	拡張機能は、内部/cpuパッケージにリストされているsse41やavxなどの命令セット拡張機能の小文字の名前です。
	例えば、cpu.avx=offは、AVX命令のランタイム検出とそれによる使用を無効にします。

	cgocheck: cgocheck=0を設定すると、cgoを使用してGoポインタを非Goコードに誤って渡すパッケージのすべてのチェックが無効になります。
	cgocheck=1（デフォルト）を設定すると、比較的安価なチェックが有効になりますが、一部のエラーを見逃す可能性があります。
	より完全で遅いcgocheckモードは、GOEXPERIMENTを使用して有効にできます（再ビルドが必要です）。
	詳細については、https://pkg.go.dev/internal/goexperiment を参照してください。

<<<<<<< HEAD
	dontfreezetheworld: デフォルトでは、致命的なパニックまたは例外の開始は「世界を凍結」し、実行中のすべてのスレッドをプリエンプトして、
	実行中のすべてのgoroutineを停止します。これにより、すべてのgoroutineをトレースバックし、パニックの発生地点に近い状態を保持することができます。
	dontfreezetheworld=1を設定すると、このプリエンプションが無効になり、パニック処理中にgoroutineが引き続き実行されるようになります。
	ただし、スケジューラに自然に入るgoroutineは引き続き停止します。これは、スケジューラのランタイムデバッグ時に有用であり、
	freezetheworldはスケジューラの状態を変更するため、問題を隠す可能性があるためです。
=======
	disablethp: setting disablethp=1 on Linux disables transparent huge pages for the heap.
	It has no effect on other platforms. disablethp is meant for compatibility with versions
	of Go before 1.21, which stopped working around a Linux kernel default that can result
	in significant memory overuse. See https://go.dev/issue/64332. This setting will be
	removed in a future release, so operators should tweak their Linux configuration to suit
	their needs before then. See https://go.dev/doc/gc-guide#Linux_transparent_huge_pages.

	dontfreezetheworld: by default, the start of a fatal panic or throw
	"freezes the world", preempting all threads to stop all running
	goroutines, which makes it possible to traceback all goroutines, and
	keeps their state close to the point of panic. Setting
	dontfreezetheworld=1 disables this preemption, allowing goroutines to
	continue executing during panic processing. Note that goroutines that
	naturally enter the scheduler will still stop. This can be useful when
	debugging the runtime scheduler, as freezetheworld perturbs scheduler
	state and thus may hide problems.
>>>>>>> upstream/release-branch.go1.21

	efence: efence=1を設定すると、アロケータがユニークなページ上に各オブジェクトを割り当て、アドレスを再利用しないモードで実行されるようになります。

	gccheckmark: gccheckmark=1を設定すると、世界が停止している間に2回目のマークパスを実行して、ガベージコレクタの並列マークフェーズの検証を有効にします。
	2回目のパスで並列マークで見つからなかった到達可能なオブジェクトが見つかった場合、ガベージコレクタはパニックを引き起こします。

	go gcpacertrace: gcpacertrace=1を設定すると、ガベージコレクタが並列ペーサーの内部状態に関する情報を出力します。

	gcshrinkstackoff: gcshrinkstackoff=1を設定すると、ゴルーチンをより小さなスタックに移動しないようにできます。
	このモードでは、ゴルーチンのスタックは成長するだけで、縮小されません。

	gcstoptheworld: gcstoptheworld=1を設定すると、並列ガベージコレクションが無効になり、すべてのガベージコレクションがストップ・ザ・ワールドイベントになります。
	gcstoptheworld=2を設定すると、ガベージコレクションが終了した後に並列スイープも無効になります。

	gctrace: gctrace=1を設定すると、ガベージコレクタが各コレクションで標準エラーに1行の要約を出力し、
	回収されたメモリ量と一時停止の長さをまとめます。この行の形式は変更される可能性があります。
	以下に説明するのは、各フィールドの関連するruntime/metricsメトリックも含まれています。現在の形式は次のとおりです。
		gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # MB stacks, #MB globals, # P
	フィールドは次のようになります。
		# GC番号
		@#s          プログラム開始以来の秒数
		#%           プログラム開始以来のGCに費やされた時間の割合
		#+...+#      GCフェーズのウォールクロック/CPU時間
		#->#-># MB   GC開始時、GC終了時、およびライブヒープのヒープサイズ、または/gc/scan/heap:bytes
		# MB goal    ゴールヒープサイズ、または/gc/heap/goal:bytes
		# MB stacks  スキャン可能なスタックサイズの推定値、または/gc/scan/stack:bytes
		# MB globals スキャン可能なグローバルサイズ、または/gc/scan/globals:bytes
		# P          使用されたプロセッサの数、または/sched/gomaxprocs:threads
	フェーズは、ストップ・ザ・ワールド（STW）スイープ終了、並列マーク・スキャン、およびSTWマーク終了です。
	マーク/スキャンのCPU時間は、アシスト時間（割り当てと同時に実行されるGC）、
	バックグラウンドGC時間、およびアイドルGC時間に分割されます。行が「(forced)」で終わる場合、
	このGCはruntime.GC()呼び出しによって強制されました。

	harddecommit: harddecommit=1を設定すると、OSに返されるメモリに対しても保護が解除されるようになります。
	これはWindowsでの唯一の動作モードですが、他のプラットフォームでのスキャベンジャー関連の問題のデバッグに役立ちます。
	現在、Linuxのみでサポートされています。

	inittrace: inittrace=1を設定すると、ランタイムが、実行時間とメモリ割り当てを要約した、
	各パッケージのinit作業ごとに標準エラーに1行の情報を出力します。
	プラグインの読み込み時に実行されるinitsと、ユーザー定義とコンパイラ生成のinit作業の両方を持たないパッケージについては、
	情報は出力されません。この行の形式は変更される可能性があります。現在の形式は次のとおりです。
		init # @#ms, # ms clock, # bytes, # allocs
	フィールドは次のようになります。
		init #      パッケージ名
		@# ms       プログラム開始以来、initが開始されたときのミリ秒単位の時間
		# clock     パッケージ初期化作業のウォールクロック時間
		# bytes     ヒープに割り当てられたメモリ
		# allocs    ヒープ割り当ての数

	madvdontneed: madvdontneed=0を設定すると、Linuxではメモリをカーネルに返すときにMADV_DONTNEEDの代わりにMADV_FREEを使用します。
	これはより効率的ですが、OSがメモリ圧力下にある場合にのみRSS数が減少することを意味します。
	BSDおよびIllumos/Solarisでは、madvdontneed=1を設定すると、MADV_FREEの代わりにMADV_DONTNEEDを使用します。
	これはより効率的ではありませんが、RSS数がより速く減少するようになります。

	memprofilerate: memprofilerate=Xを設定すると、runtime.MemProfileRateの値が更新されます。
	0に設定すると、メモリプロファイリングが無効になります。デフォルト値についてはMemProfileRateの説明を参照してください。

	pagetrace: pagetrace=/path/to/fileを設定すると、ページイベントのトレースが書き出され、
	x/debug/cmd/pagetraceツールを使用して表示、分析、および可視化できます。この機能を有効にするには、
	プログラムをGOEXPERIMENT=pagetraceでビルドしてください。
	この機能は、セットUIDバイナリの場合にセキュリティリスクを導入するため、
	プログラムがセットUIDバイナリである場合はこの機能を有効にしないでください。
	現在、Windows、plan9、js/wasmではサポートされていません。
	一部のアプリケーションでこのオプションを設定すると、大きなトレースが生成される場合があるため、注意して使用してください。

	invalidptr: invalidptr=1（デフォルト）は、ポインタ型の場所で無効なポインタ値（たとえば1）が見つかった場合、
	ガベージコレクタとスタックコピープログラムをクラッシュさせます。invalidptr=0に設定すると、このチェックが無効になります。
	これは、バグのあるコードを診断するための一時的なワークアラウンドとしてのみ使用する必要があります。
	本当の修正は、整数をポインタ型の場所に格納しないことです。

	sbrk: sbrk=1を設定すると、メモリアロケータとガベージコレクタが置き換えられ、オペレーティングシステムからメモリを取得し、メモリを回収しない単純なアロケータになります。

	scavtrace: scavtrace=1を設定すると、ランタイムが、スキャベンジャーによって実行された作業量、
	オペレーティングシステムに返された総メモリ量、および物理メモリ使用量の推定を要約した、標準エラーに1行の情報を出力します。
	この行の形式は変更される可能性があります。現在の形式は次のとおりです。
		scav # KiB work (bg), # KiB work (eager), # KiB total, #% util
	extern.goファイルから、フィールドは次のようになります。
		# KiB work (bg)    前回の行以降にバックグラウンドでOSに返されたメモリ量
		# KiB work (eager) 前回の行以降にイーガーモードでOSに返されたメモリ量
		# KiB now          現在OSに返されているアドレス空間の量
		#% util            スキャベンジングされていないヒープメモリのうち、使用中の割合
	もし行が"(forced)"で終わる場合、スキャベンジングはdebug.FreeOSMemory()呼び出しによって強制されました。

	scheddetail: schedtrace=Xおよびscheddetail=1を設定すると、
	スケジューラがXミリ秒ごとに詳細な複数行情報を出力し、スケジューラ、プロセッサ、スレッド、およびゴルーチンの状態を説明します。

	schedtrace: schedtrace=Xを設定すると、スケジューラがXミリ秒ごとに、スケジューラの状態を要約した1行を標準エラーに出力します。

	tracebackancestors: tracebackancestors=Nを設定すると、トレースバックに、ゴルーチンが作成されたスタックが追加されます。
	Nは報告する祖先ゴルーチンの数を制限します。これにより、runtime.Stackが返す情報も拡張されます。
	祖先のゴルーチンIDは、作成時のゴルーチンIDを参照します。このIDは、別のゴルーチンで再利用される可能性があります。Nを0に設定すると、祖先情報は報告されません。

	tracefpunwindoff: tracefpunwindoff=1を設定すると、実行トレーサーがフレームポインタアンワインディングの代わりに
	ランタイムのデフォルトスタックアンワインダーを使用するようになります。これにより、トレーサーのオーバーヘッドが増加しますが、
	フレームポインタアンワインディングによって引き起こされる予期しないリグレッションのワークアラウンドやデバッグに役立つ場合があります。

	asyncpreemptoff: asyncpreemptoff=1を設定すると、シグナルベースの非同期ゴルーチンプリエンプションが無効になります。
	これにより、一部のループが長時間プリエンプションされなくなり、GCとゴルーチンスケジューリングが遅れる場合があります。
	これは、非同期にプリエンプションされたゴルーチンに対して使用される保守的なスタックスキャンも無効にするため、GCの問題をデバッグするために役立ちます。

netおよびnet/httpパッケージも、GODEBUGのデバッグ変数を参照しています。
詳細については、それらのパッケージのドキュメントを参照してください。

GOMAXPROCS変数は、ユーザーレベルのGoコードを同時に実行できるオペレーティングシステムスレッドの数を制限します。
Goコードの代わりにシステムコールでブロックされているスレッドの数に制限はありません。
これらはGOMAXPROCSの制限には含まれません。このパッケージのGOMAXPROCS関数は、制限をクエリおよび変更します。

GORACE変数は、-raceを使用してビルドされたプログラムのレースディテクタを設定します。
詳細については、https://golang.org/doc/articles/race_detector.html を参照してください。

GOTRACEBACK変数は、Goプログラムが回復不能なパニックまたは予期しないランタイム条件によって失敗した場合に生成される出力量を制御します。
デフォルトでは、失敗は現在のゴルーチンのスタックトレースを出力し、ランタイムシステム内部の関数を省略して、終了コード2で終了します。
現在のゴルーチンが存在しない場合や、失敗がランタイム内部で発生した場合は、すべてのゴルーチンのスタックトレースが出力されます。
GOTRACEBACK=noneは、ゴルーチンのスタックトレースを完全に省略します。
GOTRACEBACK=single（デフォルト）は、上記の説明のように動作します。
GOTRACEBACK=allは、すべてのユーザー作成ゴルーチンのスタックトレースを追加します。
GOTRACEBACK=systemは、「all」と同様ですが、ランタイム関数のスタックフレームを追加します。
extern.goファイルから、フィールドは次のようになります。
ランタイムによって内部的に作成されたゴルーチンを表示します。
GOTRACEBACK=crashは、「system」と同様ですが、OS固有の方法でクラッシュします。たとえば、Unixシステムでは、クラッシュはSIGABRTを発生させてコアダンプをトリガーします。
GOTRACEBACK=werは、「crash」と同様ですが、Windows Error Reporting（WER）を無効にしません。
歴史的な理由から、GOTRACEBACK設定0、1、および2は、それぞれnone、all、およびsystemの同義語です。
runtime/debugパッケージのSetTraceback関数を使用すると、実行時に出力量を増やすことができますが、環境変数で指定された量を下回ることはできません。
https://golang.org/pkg/runtime/debug/#SetTraceback を参照してください。

GOARCH、GOOS、GOPATH、およびGOROOT環境変数は、Goプログラムのビルドに影響を与えます
（https://golang.org/cmd/go および https://golang.org/pkg/go/build を参照）。
GOARCH、GOOS、およびGOROOTは、コンパイル時に記録され、このパッケージの定数または関数によって利用可能になりますが、
ランタイムシステムの実行には影響しません。

# セキュリティ

Unixプラットフォームでは、危険な動作を防止するために、バイナリがsetuid/setgidに設定されているか、
setuid/setgidのようなプロパティで実行されている場合、Goのランタイムシステムはわずかに異なる動作をします。
Linuxでは、補助ベクトルでAT_SECUREフラグをチェックし、BSDおよびSolaris/Illumosではissetugidシスコールをチェックし、
AIXではuid/gidが有効なuid/gidと一致するかどうかをチェックします。

ランタイムがバイナリがsetuid/setgidのようであると判断した場合、次の3つの主な処理が行われます。
  - 標準入出力ファイルディスクリプタ（0、1、2）が開いているかどうかを確認します。いずれかが閉じられている場合、/dev/nullを指すように開きます。
  - GOTRACEBACK環境変数の値を'none'に設定します。
  - プログラムを終了するシグナルが受信された場合、またはGOTRACEBACKの値を上書きする回復不能なパニックが発生した場合、ゴルーチンのスタック、レジスタ、およびその他のメモリ関連情報が省略されます。
*/
package runtime

import (
	"github.com/shogo82148/std/internal/goarch"
	"github.com/shogo82148/std/internal/goos"
)

// Callerは、呼び出し元のゴルーチンのスタック上での関数呼び出しに関するファイルと行番号情報を報告します。
// 引数skipは、上昇するスタックフレームの数であり、0はCallerの呼び出し元を識別します。
// （歴史的な理由から、skipの意味はCallerとCallersで異なります。）
// 戻り値は、対応する呼び出しのプログラムカウンタ、ファイル名、およびファイル内の行番号を報告します。
// 情報を回復できなかった場合、ブール値okはfalseです。
func Caller(skip int) (pc uintptr, file string, line int, ok bool)

// Callersは、呼び出し元のゴルーチンのスタック上での関数呼び出しの戻りプログラムカウンタを、スライスpcで埋めます。
// 引数skipは、pcに記録する前にスキップするスタックフレームの数であり、0はCallers自体のフレームを識別し、1はCallersの呼び出し元を識別します。
// pcに書き込まれたエントリ数を返します。
//
// これらのPCを関数名や行番号などのシンボル情報に変換するには、CallersFramesを使用します。
// CallersFramesはインライン関数を考慮し、戻りプログラムカウンタを呼び出しプログラムカウンタに調整します。
// 直接PCのスライスを反復処理することは推奨されていません。また、返されたPCのいずれかに対してFuncForPCを使用することも推奨されていません。
// これらはインライン化や戻りプログラムカウンタの調整を考慮できないためです。
func Callers(skip int, pc []uintptr) int

// GOROOTは、Goツリーのルートを返します。プロセス開始時に設定されている場合はGOROOT環境変数を使用し、
// それ以外の場合はGoビルド中に使用されたルートを使用します。
func GOROOT() string

// Versionは、Goツリーのバージョン文字列を返します。
// ビルド時のコミットハッシュと日付、または可能な場合は「go1.3」のようなリリースタグです。
func Version() string

// GOOSは実行中のプログラムのオペレーティングシステムターゲットです。
// darwin、freebsd、linuxなどのいずれかです。
// GOOSとGOARCHの可能な組み合わせを表示するには、「go tool dist list」と入力してください。
const GOOS string = goos.GOOS

// GOARCHは、実行中のプログラムのアーキテクチャターゲットです。
// 386、amd64、arm、s390xなどのいずれかです。
const GOARCH string = goarch.GOARCH
