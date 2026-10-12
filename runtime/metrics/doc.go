// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: run 'go generate' (which will run 'go test -generate') to update the "Supported metrics" list.
//go:generate go test -run=Docs -generate

/*
metricsパッケージは、Goランタイムがエクスポートする実装定義のメトリクスに
アクセスするための安定したインターフェースを提供します。このパッケージは、
[runtime.ReadMemStats] や [runtime/debug.ReadGCStats] といった既存の関数に似ていますが、
はるかに汎用的です。

このパッケージで定義されるメトリクスの集合は、ランタイム自体の進化に伴って
変化する可能性があります。また、Goの実装ごとの違いも許容しており、
各実装の関連するメトリクスの集合が共通部分を持たないこともあります。

# Interface

メトリクスは、例えば構造体のフィールド名ではなく、文字列のキーによって指定されます。
サポートされているメトリクスの完全なリストは、 [All] が返すDescriptionのスライスで
常に取得できます。各 [Description] には、メトリクスに関する有用な情報も含まれています。

したがって、このAPIの利用者は、Goのバージョン間で互換性を保つために、
Allが返すスライスで定義されたサポート済みのメトリクスをサンプリングすることが推奨されます。
もちろん、特定のメトリクスを読み取ることが不可欠な状況もあります。
そのような場合には、ビルドタグを使用することが推奨されます。メトリクスは非推奨となり
削除される可能性がありますが、利用者はそれを、特定のGo実装における非常に大きな変更と
同時に起こる、例外的でまれな出来事と考えるべきです。

各メトリクスのキーには、メトリクスの値の形式を表す「種類（kind）」（ [ValueKind] を参照）もあります。
このパッケージの利用者を壊さないために、あるメトリクスの「種類」は変更されないことが保証されています。
変更しなければならない場合は、新しいキーと新しい「種類」を持つ新しいメトリクスが導入されます。

# Metric key format

前述のとおり、メトリクスのキーは文字列です。その形式はシンプルかつ明確に定義されており、
人間と機械の両方が読めるように設計されています。キーはコロンで区切られた2つの要素、
すなわちルートからのパスと単位に分かれます。キーに単位を含めることにしたのは互換性のためです。
メトリクスの単位が変わる場合、その意味もおそらく変わっているため、新しいキーを導入すべきです。

メトリクスキーのパスと単位の形式の正確な定義についての詳細は、
Description構造体のNameフィールドのドキュメントを参照してください。

# A note about floats

このパッケージは、値が浮動小数点数で表現されるメトリクスをサポートしています。
使いやすさを向上させるため、このパッケージは次の種類の浮動小数点値を決して生成しないことを約束します:
NaN、無限大。

# Supported metrics

以下は、サポートされているメトリクスの完全なリストで、辞書順に並んでいます。

	/cgo/go-to-c-calls:calls
		現在のプロセスによるGoからCへの呼び出しの回数。

	/cpu/classes/gc/mark/assist:cpu-seconds
		GCを支援し、GCがアプリケーションに遅れを取らないようにするために、
		ゴルーチンがGCタスクの実行に費やした推定合計CPU時間。
		このメトリクスは過大評価であり、システムのCPU時間の計測値とは
		直接比較できません。他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/gc/mark/dedicated:cpu-seconds
		GCタスク専用の（GOMAXPROCSで定義される）プロセッサー上で、
		GCタスクの実行に費やされた推定合計CPU時間。このメトリクスは
		過大評価であり、システムのCPU時間の計測値とは直接比較できません。
		他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/gc/mark/idle:cpu-seconds
		Goスケジューラーが他に用途を見つけられなかった余剰CPUリソース上で、
		GCタスクの実行に費やされた推定合計CPU時間。強制的なGCのCPU時間の
		指標を得るには、これをGCの合計CPU時間から差し引く必要があります。
		このメトリクスは過大評価であり、システムのCPU時間の計測値とは
		直接比較できません。他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/gc/pause:cpu-seconds
		GCによってアプリケーションが一時停止されていた推定合計CPU時間。
		一時停止中に1つのスレッドしか実行されていなくても、他に何も
		実行できないため、これはGOMAXPROCSに一時停止のレイテンシを掛けた値として
		計算されます。各サンプルに取得時点のGOMAXPROCSを掛けた場合、
		これは /sched/pauses/total/gc:seconds のサンプルの合計と正確に一致します。
		このメトリクスは過大評価であり、システムのCPU時間の計測値とは
		直接比較できません。他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/gc/total:cpu-seconds
		GCタスクの実行に費やされた推定合計CPU時間。このメトリクスは
		過大評価であり、システムのCPU時間の計測値とは直接比較できません。
		他の /cpu/classes メトリクスとのみ比較してください。
		/cpu/classes/gc 内のすべてのメトリクスの合計です。

	/cpu/classes/idle:cpu-seconds
		GoのコードやGoランタイムのコードの実行に費やされなかった、
		利用可能な推定合計CPU時間。言い換えると、
		/cpu/classes/total:cpu-seconds のうち使われなかった部分です。
		このメトリクスは過大評価であり、システムのCPU時間の計測値とは
		直接比較できません。他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/scavenge/assist:cpu-seconds
		メモリー圧迫に応じて、未使用のメモリーを基礎となるプラットフォームに
		積極的に返却するために費やされた推定合計CPU時間。このメトリクスは
		過大評価であり、システムのCPU時間の計測値とは直接比較できません。
		他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/scavenge/background:cpu-seconds
		未使用のメモリーを基礎となるプラットフォームに返却するバックグラウンド
		タスクの実行に費やされた推定合計CPU時間。このメトリクスは
		過大評価であり、システムのCPU時間の計測値とは直接比較できません。
		他の /cpu/classes メトリクスとのみ比較してください。

	/cpu/classes/scavenge/total:cpu-seconds
		未使用のメモリーを基礎となるプラットフォームに返却するタスクの実行に
		費やされた推定合計CPU時間。このメトリクスは過大評価であり、
		システムのCPU時間の計測値とは直接比較できません。
		他の /cpu/classes メトリクスとのみ比較してください。
		/cpu/classes/scavenge 内のすべてのメトリクスの合計です。

	/cpu/classes/total:cpu-seconds
		GOMAXPROCSで定義される、ユーザーのGoコードまたはGoランタイムが
		利用可能な推定合計CPU時間。言い換えると、このプロセスが実行されてきた
		実時間にわたってGOMAXPROCSを積分したものです。このメトリクスは
		過大評価であり、システムのCPU時間の計測値とは直接比較できません。
		他の /cpu/classes メトリクスとのみ比較してください。
		/cpu/classes 内のすべてのメトリクスの合計です。

	/cpu/classes/user:cpu-seconds
		ユーザーのGoコードの実行に費やされた推定合計CPU時間。これには、
		Goランタイムで費やされたわずかな時間も含まれる場合があります。
		このメトリクスは過大評価であり、システムのCPU時間の計測値とは
		直接比較できません。他の /cpu/classes メトリクスとのみ比較してください。

	/gc/cycles/automatic:gc-cycles
		Goランタイムによって発生し、完了したGCサイクルの回数。

	/gc/cycles/forced:gc-cycles
		アプリケーションによって強制され、完了したGCサイクルの回数。

	/gc/cycles/total:gc-cycles
		完了したすべてのGCサイクルの回数。

	/gc/gogc:percent
		ユーザーが設定したヒープサイズの目標パーセンテージ。設定されていない場合は100。
		この値は、GOGC環境変数および runtime/debug.SetGCPercent 関数によって設定されます。

	/gc/gomemlimit:bytes
		ユーザーが設定したGoランタイムのメモリー制限。設定されていない場合は
		math.MaxInt64。この値は、GOMEMLIMIT環境変数および
		runtime/debug.SetMemoryLimit 関数によって設定されます。

	/gc/heap/allocs-by-size:bytes
		おおよそのサイズごとのヒープ割り当ての分布。
		バケットのカウントは単調に増加します。これには
		/gc/heap/tiny/allocs:objects で定義される微小オブジェクトは含まれず、
		微小ブロックのみが含まれることに注意してください。

	/gc/heap/allocs:bytes
		アプリケーションによってヒープに割り当てられたメモリーの累積合計。

	/gc/heap/allocs:objects
		アプリケーションによって引き起こされたヒープ割り当ての累積回数。
		これには /gc/heap/tiny/allocs:objects で定義される微小オブジェクトは
		含まれず、微小ブロックのみが含まれることに注意してください。

	/gc/heap/frees-by-size:bytes
		おおよそのサイズごとの、解放されたヒープ割り当ての分布。
		バケットのカウントは単調に増加します。これには
		/gc/heap/tiny/allocs:objects で定義される微小オブジェクトは含まれず、
		微小ブロックのみが含まれることに注意してください。

	/gc/heap/frees:bytes
		ガベージコレクターによって解放されたヒープメモリーの累積合計。

	/gc/heap/frees:objects
		ガベージコレクターによって記憶領域が解放されたヒープ割り当ての累積回数。
		これには /gc/heap/tiny/allocs:objects で定義される微小オブジェクトは
		含まれず、微小ブロックのみが含まれることに注意してください。

	/gc/heap/goal:bytes
		GCサイクル終了時のヒープサイズの目標。

	/gc/heap/live:bytes
		前回のGCでマークされた生存オブジェクトが占めるヒープメモリー。

	/gc/heap/objects:objects
		ヒープメモリーを占めている、生存中またはスイープされていないオブジェクトの数。

	/gc/heap/tiny/allocs:objects
		まとめてブロックに詰め込まれる小さな割り当ての回数。
		個々の割り当てはランタイムによって追跡されず、そのブロックのみが
		追跡されるため、これらの割り当ては他の割り当てとは別に数えられます。
		各ブロックは、allocs-by-size と frees-by-size ですでに計上されています。

	/gc/limiter/last-enabled:gc-cycle
		GCのCPUリミッターが最後に有効になったときのGCサイクル。
		リミッターはGCのCPU時間が高くなりすぎたときにメモリーとCPU時間を
		引き換えにするため、このメトリクスはメモリー不足エラーの根本原因を
		診断するのに役立ちます。これはSetMemoryLimitを使用している場合に
		最も起こりやすいです。最初のGCサイクルはサイクル1なので、
		値が0の場合は一度も有効になっていないことを示します。

	/gc/pauses:seconds
		非推奨です。同一の /sched/pauses/total/gc:seconds を使用してください。

	/gc/scan/globals:bytes
		スキャン可能なグローバル変数領域の合計量。

	/gc/scan/heap:bytes
		スキャン可能なヒープ領域の合計量。

	/gc/scan/stack:bytes
		前回のGCサイクルでスキャンされたスタックのバイト数。

	/gc/scan/total:bytes
		スキャン可能な領域の合計量。/gc/scan 内のすべてのメトリクスの合計です。

	/gc/stack/starting-size:bytes
		新しいゴルーチンのスタックサイズ。

	/godebug/non-default-behavior/allowmultiplevcs:events
		デフォルト以外の GODEBUG=allowmultiplevcs=... 設定により、
		cmd/go パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/containermaxprocs:events
		デフォルト以外の GODEBUG=containermaxprocs=... 設定により、
		runtime パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/cryptocustomrand:events
		デフォルト以外の GODEBUG=cryptocustomrand=... 設定により、
		crypto パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/embedfollowsymlinks:events
		デフォルト以外の GODEBUG=embedfollowsymlinks=... 設定により、
		cmd/go パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/execerrdot:events
		デフォルト以外の GODEBUG=execerrdot=... 設定により、
		os/exec パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/fips140ems:events
		デフォルト以外の GODEBUG=fips140ems=... 設定により、
		crypto/tls パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/gocachehash:events
		デフォルト以外の GODEBUG=gocachehash=... 設定により、
		cmd/go パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/gocachetest:events
		デフォルト以外の GODEBUG=gocachetest=... 設定により、
		cmd/go パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/gocacheverify:events
		デフォルト以外の GODEBUG=gocacheverify=... 設定により、
		cmd/go パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/gotestjsonbuildtext:events
		デフォルト以外の GODEBUG=gotestjsonbuildtext=... 設定により、
		cmd/go パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/htmlmetacontenturlescape:events
		デフォルト以外の GODEBUG=htmlmetacontenturlescape=... 設定により、
		html/template パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/http2client:events
		デフォルト以外の GODEBUG=http2client=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/http2server:events
		デフォルト以外の GODEBUG=http2server=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/httpcookiemaxnum:events
		デフォルト以外の GODEBUG=httpcookiemaxnum=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/httplaxcontentlength:events
		デフォルト以外の GODEBUG=httplaxcontentlength=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/httpmuxgo121:events
		デフォルト以外の GODEBUG=httpmuxgo121=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/httpservecontentkeepheaders:events
		デフォルト以外の GODEBUG=httpservecontentkeepheaders=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/httpservecontentmaxranges:events
		デフォルト以外の GODEBUG=httpservecontentmaxranges=... 設定により、
		net/http パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/installgoroot:events
		デフォルト以外の GODEBUG=installgoroot=... 設定により、
		go/build パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/multipartmaxheaders:events
		デフォルト以外の GODEBUG=multipartmaxheaders=... 設定により、
		mime/multipart パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/multipartmaxparts:events
		デフォルト以外の GODEBUG=multipartmaxparts=... 設定により、
		mime/multipart パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/multipathtcp:events
		デフォルト以外の GODEBUG=multipathtcp=... 設定により、
		net パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/netedns0:events
		デフォルト以外の GODEBUG=netedns0=... 設定により、
		net パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/panicnil:events
		デフォルト以外の GODEBUG=panicnil=... 設定により、
		runtime パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/randautoseed:events
		デフォルト以外の GODEBUG=randautoseed=... 設定により、
		math/rand パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/randseednop:events
		デフォルト以外の GODEBUG=randseednop=... 設定により、
		math/rand パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/rsa1024min:events
		デフォルト以外の GODEBUG=rsa1024min=... 設定により、
		crypto/rsa パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/tarinsecurepath:events
		デフォルト以外の GODEBUG=tarinsecurepath=... 設定により、
		archive/tar パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/tlsmaxrsasize:events
		デフォルト以外の GODEBUG=tlsmaxrsasize=... 設定により、
		crypto/tls パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/tlssha1:events
		デフォルト以外の GODEBUG=tlssha1=... 設定により、
		crypto/tls パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/updatemaxprocs:events
		デフォルト以外の GODEBUG=updatemaxprocs=... 設定により、
		runtime パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/urlmaxqueryparams:events
		デフォルト以外の GODEBUG=urlmaxqueryparams=... 設定により、
		net/url パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/urlstrictcolons:events
		デフォルト以外の GODEBUG=urlstrictcolons=... 設定により、
		net/url パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/winreadlinkvolume:events
		デフォルト以外の GODEBUG=winreadlinkvolume=... 設定により、
		os パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/winsymlink:events
		デフォルト以外の GODEBUG=winsymlink=... 設定により、
		os パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/x509negativeserial:events
		デフォルト以外の GODEBUG=x509negativeserial=... 設定により、
		crypto/x509 パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/x509rsacrt:events
		デフォルト以外の GODEBUG=x509rsacrt=... 設定により、
		crypto/x509 パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/x509sha256skid:events
		デフォルト以外の GODEBUG=x509sha256skid=... 設定により、
		crypto/x509 パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/x509sslcertoverrideplatform:events
		デフォルト以外の GODEBUG=x509sslcertoverrideplatform=... 設定により、
		crypto/x509 パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/x509usefallbackroots:events
		デフォルト以外の GODEBUG=x509usefallbackroots=... 設定により、
		crypto/x509 パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/x509usepolicies:events
		デフォルト以外の GODEBUG=x509usepolicies=... 設定により、
		crypto/x509 パッケージで実行されたデフォルト以外の動作の回数。

	/godebug/non-default-behavior/zipinsecurepath:events
		デフォルト以外の GODEBUG=zipinsecurepath=... 設定により、
		archive/zip パッケージで実行されたデフォルト以外の動作の回数。

	/memory/classes/heap/free:bytes
		完全に空いていて基礎となるシステムに返却可能であるものの、
		まだ返却されていないメモリー。このメトリクスは、物理メモリーに
		裏付けられた空きアドレス空間のランタイムによる推定値です。

	/memory/classes/heap/objects:bytes
		生存オブジェクトと、ガベージコレクターによってまだ解放済みとして
		マークされていない死んだオブジェクトが占めるメモリー。

	/memory/classes/heap/released:bytes
		完全に空いていて、基礎となるシステムに返却済みのメモリー。
		このメトリクスは、プロセスにまだマップされているものの、
		物理メモリーに裏付けられていない空きアドレス空間のランタイムによる推定値です。

	/memory/classes/heap/stacks:bytes
		スタック領域のために予約された、ヒープから割り当てられたメモリー。
		現在使用中かどうかは問いません。現在、これはゴルーチンのすべての
		スタックメモリーを表します。また、cgoを使用しないプログラムでは、
		すべてのOSスレッドのスタックも含まれます。スタックは将来異なる方法で
		割り当てられる可能性があり、これは変更される可能性があることに注意してください。

	/memory/classes/heap/unused:bytes
		ヒープオブジェクトのために予約されているものの、現在は
		ヒープオブジェクトの保持に使われていないメモリー。

	/memory/classes/metadata/mcache/free:bytes
		ランタイムのmcache構造体のために予約されているものの、
		使用されていないメモリー。

	/memory/classes/metadata/mcache/inuse:bytes
		現在使用されているランタイムのmcache構造体が占めるメモリー。

	/memory/classes/metadata/mspan/free:bytes
		ランタイムのmspan構造体のために予約されているものの、
		使用されていないメモリー。

	/memory/classes/metadata/mspan/inuse:bytes
		現在使用されているランタイムのmspan構造体が占めるメモリー。

	/memory/classes/metadata/other:bytes
		ランタイムのメタデータのために予約されている、またはその保持に使われているメモリー。

	/memory/classes/os-stacks:bytes
		基礎となるオペレーティングシステムによって割り当てられたスタックメモリー。
		cgoを使用しないプログラムでは、このメトリクスは現在ゼロです。これは
		将来変更される可能性があります。cgoを使用するプログラムでは、このメトリクスには
		OSから直接割り当てられたOSスレッドのスタックが含まれます。現在、
		c-sharedおよびc-archiveビルドモードでは1つのスタックのみが計上され、
		OSからのその他のスタックの割り当て元は計測されません。
		これも将来変更される可能性があります。

	/memory/classes/other:bytes
		実行トレースのバッファー、ランタイムのデバッグ用の構造体、
		ファイナライザーとプロファイラーの特殊レコードなどに使われるメモリー。

	/memory/classes/profiling/buckets:bytes
		プロファイリングに使われるスタックトレースのハッシュマップが使用するメモリー。

	/memory/classes/total:bytes
		Goランタイムによって現在のプロセスに読み書き可能としてマップされたすべてのメモリー。
		これには、cgo経由またはsyscallパッケージ経由で呼び出されたコードによって
		マップされたメモリーは含まれないことに注意してください。
		/memory/classes 内のすべてのメトリクスの合計です。

	/sched/gomaxprocs:threads
		現在の runtime.GOMAXPROCS の設定、つまりユーザーレベルのGoコードを
		同時に実行できるオペレーティングシステムのスレッドの数。

	/sched/goroutines:goroutines
		生存しているゴルーチンの数。

	/sched/latencies:seconds
		ゴルーチンが実際に実行される前に、実行可能状態でスケジューラー内で
		過ごした時間の分布。バケットのカウントは単調に増加します。

	/sched/pauses/stopping/gc:seconds
		GC関連のストップ・ザ・ワールドにおける、個々の停止レイテンシの分布。
		これは、ワールドを停止することを決定してから、すべてのPが停止するまでに
		かかる時間です。これは、GC関連のストップ・ザ・ワールドの合計時間
		（/sched/pauses/total/gc:seconds）の一部です。この間、一部のスレッドが
		実行されている可能性があります。バケットのカウントは単調に増加します。

	/sched/pauses/stopping/other:seconds
		GCに関連しないストップ・ザ・ワールドにおける、個々の停止レイテンシの分布。
		これは、ワールドを停止することを決定してから、すべてのPが停止するまでに
		かかる時間です。これは、GCに関連しないストップ・ザ・ワールドの合計時間
		（/sched/pauses/total/other:seconds）の一部です。この間、一部のスレッドが
		実行されている可能性があります。バケットのカウントは単調に増加します。

	/sched/pauses/total/gc:seconds
		GC関連のストップ・ザ・ワールドにおける、個々の一時停止レイテンシの分布。
		これは、ワールドを停止することを決定してから、ワールドが再び開始されるまでの
		時間です。この時間の一部は、すべてのスレッドを停止させるために費やされ
		（これは /sched/pauses/stopping/gc:seconds で直接計測されます）、
		その間一部のスレッドはまだ実行されている可能性があります。
		バケットのカウントは単調に増加します。

	/sched/pauses/total/other:seconds
		GCに関連しないストップ・ザ・ワールドにおける、個々の一時停止レイテンシの分布。
		これは、ワールドを停止することを決定してから、ワールドが再び開始されるまでの
		時間です。この時間の一部は、すべてのスレッドを停止させるために費やされます
		（/sched/pauses/stopping/other:seconds で直接計測されます）。
		バケットのカウントは単調に増加します。

	/sync/mutex/wait/total:seconds
		ゴルーチンが sync.Mutex 、 sync.RWMutex 、またはランタイム内部のロックで
		ブロックされていたおおよその累積時間。このメトリクスは、ロック競合の
		全体的な変化を特定するのに役立ちます。より詳細な競合データを得るには、
		runtime/pprofパッケージを使用してmutexプロファイルまたはblockプロファイルを
		収集してください。
*/
package metrics
