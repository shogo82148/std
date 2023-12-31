// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: run 'go generate' (which will run 'go test -generate') to update the "Supported metrics" list.
//go:generate go test -run=Docs -generate

/*
パッケージmetricsは、Goランタイムによってエクスポートされる実装定義の
メトリクスにアクセスするための安定したインターフェースを提供します。このパッケージは、
既存の関数である [runtime.ReadMemStats] や [debug.ReadGCStats] に似ていますが、
かなり一般的です。

このパッケージで定義されているメトリクスのセットは、ランタイム自体が進化するにつれて進化し、
また、関連するメトリクスのセットが交差しない可能性のあるGoの実装間でのバリエーションを可能にします。

# Interface

メトリクスは、例えば、構造体のフィールド名ではなく、文字列キーによって指定されます。
サポートされているメトリクスの完全なリストは、常にAllによって返されるDescriptionsのスライスにあります。
各Descriptionには、メトリクスに関する有用な情報も含まれています。

したがって、このAPIのユーザーは、Allが返すスライスによって定義されたサポートされているメトリクスをサンプリングすることを推奨します。
これにより、Goのバージョン間での互換性が保たれます。もちろん、特定のメトリクスを読み取ることが重要な状況が生じることもあります。
これらのケースでは、ユーザーはビルドタグを使用することを推奨します。メトリクスは非推奨となり削除されることもありますが、
これは特定のGoの実装における非常に大きな変更と一致する、例外的で稀なイベントと考えてください。

各メトリックキーには、メトリックの値の形式を説明する「種類」もあります。
このパッケージのユーザーを壊さないために、特定のメトリックの「種類」は変更されないことが保証されています。
それが変更しなければならない場合、新しいキーと新しい「種類」で新しいメトリックが導入されます。

# Metric key format

前述の通り、メトリックキーは文字列です。その形式はシンプルで明確に定義されており、
人間と機械の両方が読み取れるように設計されています。それは二つのコンポーネントに分割され、
コロンで区切られています：ルート化されたパスと単位です。キーに単位を含める選択は、
互換性によって動機づけられています：メトリックの単位が変更されると、そのセマンティクスも
おそらく変更され、新しいキーが導入されるべきです。

メトリックキーのパスと単位の形式の詳細な定義については、
Description構造体のNameフィールドのドキュメンテーションを参照してください。

# A note about floats

このパッケージは、値が浮動小数点表現を持つメトリクスをサポートしています。
使いやすさを向上させるために、このパッケージはNaNや無限大といった
浮動小数点数のクラスを生成しないことを約束します。

# Supported metrics

以下は、辞書順に並べたサポートされているメトリクスの完全なリストです。

	/cgo/go-to-c-calls:calls
		現在のプロセスによってGoからCへ行われた呼び出しの数。

	/cpu/classes/gc/mark/assist:cpu-seconds
		GCがアプリケーションに遅れを取らないように、GCのタスクを実行するためにゴルーチンが費やした
		CPU時間の推定総量。このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/gc/mark/dedicated:cpu-seconds
		GCのタスクに専用のプロセッサ（GOMAXPROCSによって定義）でGCのタスクを実行するために費やした
		CPU時間の推定総量。このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/gc/mark/idle:cpu-seconds
		Goスケジューラが他の用途で使用できなかった余剰のCPUリソースでGCのタスクを実行するために費やした
		CPU時間の推定総量。これはGCのCPU時間の合計から差し引くべきで、強制的なGCのCPU時間の尺度を得るためです。
		このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/gc/pause:cpu-seconds
		GCによってアプリケーションが一時停止されている間に費やされたCPU時間の推定総量。
		一時停止中に実行されているスレッドが1つだけでも、他の何も実行できないため、
		これはGOMAXPROCS倍の一時停止遅延として計算されます。これは、各サンプルが取得された時点で
		GOMAXPROCSで乗算される場合の/gc/pause:secondsのサンプルの正確な合計です。
		このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/gc/total:cpu-seconds
		GCタスクの実行に費やされたCPU時間の推定総量。このメトリックは過大評価であり、
		システムのCPU時間の測定値とは直接比較できません。他の/cpu/classesメトリックとのみ比較してください。
		/cpu/classes/gc内のすべてのメトリックの合計。

	/cpu/classes/idle:cpu-seconds
		GoまたはGoランタイムコードの実行に使用されなかった利用可能なCPU時間の推定総量。
		つまり、/cpu/classes/total:cpu-secondsの未使用部分です。このメトリックは過大評価であり、
		システムのCPU時間の測定値とは直接比較できません。他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/scavenge/assist:cpu-seconds
		メモリ圧力に対する応答として、使用されていないメモリを積極的に基盤となるプラットフォームに返すために費やした
		CPU時間の推定総量。このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/scavenge/background:cpu-seconds
		未使用のメモリを基盤となるプラットフォームに返すためのバックグラウンドタスクを実行するために費やした
		CPU時間の推定総量。このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。

	/cpu/classes/scavenge/total:cpu-seconds
		未使用のメモリを基盤となるプラットフォームに返すタスクを実行するために費やした
		CPU時間の推定総量。このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。/cpu/classes/scavenge内のすべてのメトリックの合計。

	/cpu/classes/total:cpu-seconds
		GOMAXPROCSによって定義された、ユーザーのGoコードまたはGoランタイムの利用可能なCPU時間の推定総量。
		つまり、このプロセスが実行されている壁時計の期間をGOMAXPROCSで積分したものです。
		このメトリックは過大評価であり、システムのCPU時間の測定値とは直接比較できません。
		他の/cpu/classesメトリックとのみ比較してください。/cpu/classes内のすべてのメトリックの合計。

	/cpu/classes/user:cpu-seconds
		ユーザーのGoコードの実行に費やされたCPU時間の推定総量。これには、Goランタイムで費やされた
		わずかな時間も含まれる可能性があります。このメトリックは過大評価であり、システムのCPU時間の測定値とは
		直接比較できません。他の/cpu/classesメトリックとのみ比較してください。

	/gc/cycles/automatic:gc-cycles
		Goランタイムによって生成された完了したGCサイクルの数。

	/gc/cycles/forced:gc-cycles
		アプリケーションによって強制された完了したGCサイクルの数。

	/gc/cycles/total:gc-cycles
		すべての完了したGCサイクルの数。

	/gc/gogc:percent
		ユーザーによって設定されたヒープサイズの目標パーセンテージ、それ以外の場合は
		100。この値はGOGC環境変数とruntime/debug.SetGCPercent関数によって設定されます。

	/gc/gomemlimit:bytes
		ユーザーによって設定されたGoランタイムのメモリ制限、それ以外の場合は
		math.MaxInt64。この値はGOMEMLIMIT環境変数とruntime/debug.SetMemoryLimit関数によって設定されます。

	/gc/heap/allocs-by-size:bytes
		近似的なサイズ別のヒープ割り当ての分布。
		バケットのカウントは単調に増加します。これには、/gc/heap/tiny/allocs:objectsによって定義された
		小さなオブジェクトは含まれません、小さなブロックのみが含まれます。

	/gc/heap/allocs:bytes
		アプリケーションによってヒープに割り当てられたメモリの累積合計。

	/gc/heap/allocs:objects
		アプリケーションによって引き起こされたヒープ割り当ての累積カウント。
		これには、/gc/heap/tiny/allocs:objectsによって定義された小さなオブジェクトは含まれません、
		小さなブロックのみが含まれます。

	/gc/heap/frees-by-size:bytes
		近似的なサイズ別の解放されたヒープ割り当ての分布。
		バケットのカウントは単調に増加します。これには、/gc/heap/tiny/allocs:objectsによって定義された
		小さなオブジェクトは含まれません、小さなブロックのみが含まれます。

	/gc/heap/frees:bytes
		ガベージコレクタによって解放されたヒープメモリの累積合計。

	/gc/heap/frees:objects
		ガベージコレクタによってストレージが解放されたヒープ割り当ての累積カウント。
		これには、/gc/heap/tiny/allocs:objectsによって定義された小さなオブジェクトは含まれません、
		小さなブロックのみが含まれます。

	/gc/heap/goal:bytes
		GCサイクルの終了時のヒープサイズの目標。

	/gc/heap/live:bytes
		前回のGCによってマークされた生存オブジェクトが占めるヒープメモリ。

	/gc/heap/objects:objects
		ヒープメモリを占有している、生存しているか未掃除のオブジェクトの数。

	/gc/heap/tiny/allocs:objects
		ブロックにまとめられた小さな割り当ての数。
		これらの割り当ては、各個別の割り当てがランタイムによって追跡されていないため、
		他の割り当てとは別にカウントされます、ブロックのみが追跡されます。各ブロックはすでに
		allocs-by-sizeとfrees-by-sizeで計算されています。

	/gc/limiter/last-enabled:gc-cycle
		GC CPUリミッターが最後に有効にされたGCサイクル。
		このメトリックは、リミッターがGCのCPU時間が高すぎるときにメモリをCPU時間と交換するため、
		メモリ不足エラーの根本原因を診断するのに役立ちます。これは、SetMemoryLimitの使用時に最も発生しやすいです。
		最初のGCサイクルはサイクル1なので、値が0の場合はそれが一度も有効にされなかったことを示します。

	/gc/pauses:seconds
		個々のGC関連の世界停止一時停止の遅延の分布。
		バケットのカウントは単調に増加します。

	/gc/scan/globals:bytes
		スキャン可能なグローバル変数スペースの総量。

	/gc/scan/heap:bytes
		スキャン可能なヒープスペースの総量。

	/gc/scan/stack:bytes
		最後のGCサイクルでスキャンされたスタックのバイト数。

	/gc/scan/total:bytes
		スキャン可能なスペースの総量。/gc/scan内のすべてのメトリックの合計。

	/gc/stack/starting-size:bytes
		新しいゴルーチンのスタックサイズ。

	/godebug/non-default-behavior/execerrdot:events
		非デフォルトのGODEBUG=execerrdot=...設定により、os/execパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/gocachehash:events
		非デフォルトのGODEBUG=gocachehash=...設定により、cmd/goパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/gocachetest:events
		非デフォルトのGODEBUG=gocachetest=...設定により、cmd/goパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/gocacheverify:events
		非デフォルトのGODEBUG=gocacheverify=...設定により、cmd/goパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/http2client:events
		非デフォルトのGODEBUG=http2client=...設定により、net/httpパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/http2server:events
		非デフォルトのGODEBUG=http2server=...設定により、net/httpパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/installgoroot:events
		非デフォルトのGODEBUG=installgoroot=...設定により、go/buildパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/jstmpllitinterp:events
		非デフォルトのGODEBUG=jstmpllitinterp=...設定により、html/templateパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/multipartmaxheaders:events
		非デフォルトのGODEBUG=multipartmaxheaders=...設定により、mime/multipartパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/multipartmaxparts:events
		非デフォルトのGODEBUG=multipartmaxparts=...設定により、mime/multipartパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/multipathtcp:events
		非デフォルトのGODEBUG=multipathtcp=...設定により、netパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/panicnil:events
		非デフォルトのGODEBUG=panicnil=...設定により、ランタイムパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/randautoseed:events
		非デフォルトのGODEBUG=randautoseed=...設定により、math/randパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/tarinsecurepath:events
		非デフォルトのGODEBUG=tarinsecurepath=...設定により、archive/tarパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/tlsmaxrsasize:events
		非デフォルトのGODEBUG=tlsmaxrsasize=...設定により、crypto/tlsパッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/x509sha1:events
		非デフォルトのGODEBUG=x509sha1=...設定により、crypto/x509パッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/x509usefallbackroots:events
		非デフォルトのGODEBUG=x509usefallbackroots=...設定により、crypto/x509パッケージが実行した
		非デフォルトの動作の数。

	/godebug/non-default-behavior/zipinsecurepath:events
		非デフォルトのGODEBUG=zipinsecurepath=...設定により、archive/zipパッケージが実行した
		非デフォルトの動作の数。

	/memory/classes/heap/free:bytes
		完全に空き、かつ基礎となるシステムに返すことが可能なメモリですが、まだ返されていません。
		このメトリックは、物理メモリにバックアップされた空きアドレススペースのランタイムの推定値です。

	/memory/classes/heap/objects:bytes
		生存しているオブジェクトとまだガベージコレクタによってフリーマークされていない死んだオブジェクトが占有しているメモリ。

	/memory/classes/heap/released:bytes
		完全に空き、かつ基礎となるシステムに返されたメモリ。このメトリックは、
		まだプロセスにマップされているが物理メモリにバックアップされていない空きアドレススペースのランタイムの推定値です。

	/memory/classes/heap/stacks:bytes
		スタックスペースのためにヒープから割り当てられたメモリで、現在使用中であるかどうかに関わらず予約されています。
		現在、これはすべてのゴルーチンのスタックメモリを表しています。また、非cgoプログラムのすべてのOSスレッドスタックも含まれます。
		注意してください、将来的にスタックは異なる方法で割り当てられる可能性があり、これが変更されるかもしれません。

	/memory/classes/heap/unused:bytes
		ヒープオブジェクト用に予約されているが、現在はヒープオブジェクトを保持するためには使用されていないメモリ。

	/memory/classes/metadata/mcache/free:bytes
		ランタイムmcache構造体のために予約されているが、使用中ではないメモリ。

	/memory/classes/metadata/mcache/inuse:bytes
		現在使用中のランタイムmcache構造体によって占有されているメモリ。

	/memory/classes/metadata/mspan/free:bytes
		ランタイムmspan構造体のために予約されているが、使用中ではないメモリ。

	/memory/classes/metadata/mspan/inuse:bytes
		現在使用中のランタイムmspan構造体によって占有されているメモリ。

	/memory/classes/metadata/other:bytes
		ランタイムメタデータを保持するために予約または使用されているメモリ。

	/memory/classes/os-stacks:bytes
		基礎となるオペレーティングシステムによって割り当てられたスタックメモリ。
		非cgoプログラムでは、このメトリックは現在ゼロです。これは将来変更される可能性があります。
		cgoプログラムでは、このメトリックにはOSから直接割り当てられたOSスレッドスタックが含まれます。
		現在、これはc-sharedとc-archiveビルドモードの一つのスタックのみを計算し、
		OSからの他のスタックソースは測定されていません。これも将来変更される可能性があります。

	/memory/classes/other:bytes
		実行トレースバッファ、ランタイムのデバッグ用の構造体、ファイナライザーとプロファイラーの特殊なもの、その他に使用されるメモリ。

	/memory/classes/profiling/buckets:bytes
		プロファイリングに使用されるスタックトレースのハッシュマップに使用されるメモリ。

	/memory/classes/total:bytes
		Goランタイムによって現在のプロセスに読み書き可能としてマップされたすべてのメモリ。
		これには、cgo経由またはsyscallパッケージ経由で呼び出されたコードによってマップされたメモリは含まれません。
		/memory/classes内のすべてのメトリックの合計。

	/sched/gomaxprocs:threads
		現在のruntime.GOMAXPROCS設定、または同時にユーザーレベルのGoコードを実行できる
		オペレーティングシステムのスレッド数。

	/sched/goroutines:goroutines
		生存しているゴルーチンの数。

	/sched/latencies:seconds
		実際に実行される前に、スケジューラ内でゴルーチンが実行可能な状態で過ごした時間の分布。
		バケットの数は単調に増加します。

	/sync/mutex/wait/total:seconds
		ゴルーチンがsync.Mutexまたはsync.RWMutexでブロックされて過ごした累計時間の概算。
		このメトリックは、ロック競合の全体的な変化を特定するのに役立ちます。
		より詳細な競合データについては、runtime/pprofパッケージを使用してミューテックスまたはブロックプロファイルを収集します。
*/
package metrics
