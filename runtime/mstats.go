// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// メモリの統計情報

package runtime

// MemStatsはメモリアロケータに関する統計情報を記録します。
type MemStats struct {

	// Allocは割り当てられたヒープオブジェクトのバイト数です。
	//
	// これはHeapAllocと同じです（以下を参照）。
	Alloc uint64

	// TotalAllocはヒープオブジェクトのために割り当てられた累積バイト数です。
	//
	// TotalAllocはヒープオブジェクトが割り当てられると増加しますが、
	// AllocとHeapAllocとは異なり、オブジェクトが解放されると減少しません。
	TotalAlloc uint64

	// SysはOSから取得したメモリの合計バイト数です。
	//
	// Sysは以下のXSysフィールドの合計です。SysはGoランタイムがヒープ、スタック、および他の内部データ構造に予約した仮想アドレススペースを計測します。ある時点では、仮想アドレススペースのすべてが物理メモリでバックアップされているわけではありませんが、一般的にはすべてバックアップされています。
	Sys uint64

	// Lookupsはランタイムによって実行されるポインターの参照の数です。
	//
	// これは主にランタイム内部のデバッグに役立ちます。
	Lookups uint64

	// Mallocsはヒープオブジェクトの割り当て数の累積数です。
	// 生存しているオブジェクトの数はMallocs - Freesです。
	Mallocs uint64

	// Frees はヒープオブジェクトが解放された累積数です。
	Frees uint64

	// HeapAllocは割り当てられたヒープオブジェクトのバイト数です。
	//
	// "割り当てられた"ヒープオブジェクトには、到達可能なオブジェクト全体と、
	// ガベージコレクタがまだ解放していない到達不能なオブジェクトが含まれます。
	// 具体的には、ヒープオブジェクトが割り当てられるとHeapAllocは増加し、
	// ヒープがスイープされて到達不能なオブジェクトが解放されるとHeapAllocは減少します。
	// スイープはGCサイクル間に段階的に行われるため、
	// これらの2つのプロセスは同時に発生し、その結果HeapAllocは滑らかに変化します
	//（ストップ・ザ・ワールドのガベージコレクタの典型的なギザギザとは対照的です）。
	HeapAlloc uint64

	// HeapSysはOSから取得されたヒープメモリのバイト数です。
	//
	// HeapSysは、ヒープのために予約された仮想アドレス空間の量を測定します。
	// これには、まだ使用されていないが予約されている仮想アドレス空間が含まれます。
	// これは物理メモリを消費しませんが、通常は小さくなります。
	// また、使用されなくなった後に物理メモリがOSに返された仮想アドレス空間も含まれます（後者の測定にはHeapReleasedを参照してください）。
	//
	// HeapSysは、ヒープが持っていた最大のサイズを推測します。
	HeapSys uint64

	// HeapIdleはアイドル状態（未使用）のスパンのバイト数です。
	//
	// アイドルスパンにはオブジェクトが含まれていません。これらのスパンは
	// OSに返却されることができます（または既に返却されているかもしれません）し、ヒープの割り当てに再利用されることもあります。
	// また、スタックメモリとして再利用されることもあります。
	//
	// HeapIdleからHeapReleasedを引いた値は、OSに返還できるメモリ量を見積もるものですが、
	// ランタイムによって保持されているため、ヒープの拡張時にOSからの追加メモリ要求なしでヒープを成長させるために利用されています。
	// もし、この差がヒープサイズよりもはるかに大きい場合、最近の一時的なライブヒープサイズの急増を示しています。
	HeapIdle uint64

	// HeapInuseは使用中スパンのバイト数です。
	//
	// 使用中スパンには少なくとも1つのオブジェクトが含まれています。
	// これらのスパンはおおよそ同じサイズの他のオブジェクトにのみ使用できます。
	//
	// HeapInuseからHeapAllocを引いた値は、特定のサイズクラスに割り当てられたメモリの量を推定しますが、現在は使用されていません。
	// これは断片化の上限であり、一般的にこのメモリは効率的に再利用できます。
	HeapInuse uint64

	// HeapReleasedはOSに返される物理メモリのバイト数です。
	//
	// これは、ヒープに再取得される前に、アイドルスパンから返された
	// ヒープメモリをカウントしています。
	HeapReleased uint64

	// HeapObjectsは割り当てられたヒープオブジェクトの数です。
	//
	// HeapAllocと同様に、オブジェクトが割り当てられると増加し、
	// ヒープが掃引され、到達不能なオブジェクトが解放されると減少します。
	HeapObjects uint64

	// StackInuse はスタックスパンのバイト数です。
	//
	// 使用中のスタックスパンには少なくとも1つのスタックがあります。これらのスパンは同じサイズの他のスタックにしか使用できません。
	//
	// StackIdle は存在しません。未使用のスタックスパンはヒープに戻されるため（そのため HeapIdle にカウントされます）。
	StackInuse uint64

	// StackSysはOSから取得したスタックメモリのバイト数です。
	//
	// StackSysはStackInuseに加えて、OSスレッドスタック用に直接
	// OSから取得したメモリです。
	//
	// cgoを使用しないプログラムでは、このメトリックは現在StackInuseと同じです
	// （しかし、これに依存するべきではなく、値は将来変わる可能性があります）。
	//
	// cgoを使用するプログラムでは、OSスレッドスタックも含まれます。
	// 現在、c-sharedおよびc-archiveビルドモードでは1つのスタックのみを考慮し、
	// 他のOSからのスタック（特にCコードによって割り当てられたスタック）は現在計測されていません。
	// これも将来変更される可能性があります。
	StackSys uint64

	// MSpanInuseは割り当てられたmspan構造体のバイト数です。
	MSpanInuse uint64

	// MSpanSysは、mspan構造体のためにOSから取得したメモリのバイトです。
	MSpanSys uint64

	// MCacheInuseは割り当てられたmcache構造体のバイト数です。
	MCacheInuse uint64

	// MCacheSysは、mcache構造体のためにオペレーティングシステムから取得されたバイト数です。
	MCacheSys uint64

	BuckHashSys uint64

	// GCSysはゴミ回収メタデータのメモリのバイト数です。
	GCSys uint64

	// OtherSys は、さまざまなオフヒープランタイム割り当てのメモリのバイト数です。
	OtherSys uint64

	// NextGCは次のGCサイクルのターゲットヒープサイズです。
	//
	// ガベージコレクタの目標は、HeapAlloc ≤ NextGCを維持することです。
	// 各GCサイクルの終了時に、次のサイクルのターゲットは
	// アクセス可能なデータの量とGOGCの値に基づいて計算されます。
	NextGC uint64

	// LastGCは、最後のガベージコレクションが終了した時刻で、
	// 1970年以降のUNIXエポックからの経過時間（ナノ秒単位）です。
	LastGC uint64

	// PauseTotalNsは、プログラムが開始されてからのGCによる累積ナノ秒数です。
	//
	// ストップ・ザ・ワールド・ポーズ中には、すべてのゴルーチンが一時停止され、
	// ガベージコレクタのみが実行されます。
	PauseTotalNs uint64

	// PauseNsは直近のGCのストップ・ザ・ワールドの一時停止時間（ナノ秒）の循環バッファです。
	//
	// 最も直近の一時停止はPauseNs[(NumGC+255)%256]にあります。一般的に、PauseNs[N%256]は直近のN%256番目のGCサイクルでの一時停止時間を記録しています。1つのGCサイクルに複数の一時停止が存在する可能性があります。これはサイクル中のすべての一時停止の合計です。
	PauseNs [256]uint64

	// PauseEndは最近のGCの一時停止終了時間の循環バッファで、1970年以降のナノ秒（UNIXエポック）で表されます。
	//
	// このバッファはPauseNsと同じ方法で埋められます。1つのGCサイクルに複数の一時停止がある場合があります。このバッファはサイクル内の最後の一時停止の終了を記録します。
	PauseEnd [256]uint64

	// NumGCは完了したGCサイクルの数です。
	NumGC uint32

	// NumForcedGC は、アプリケーションがGC関数を呼び出し、強制的に実行されたGCサイクルの回数です。
	NumForcedGC uint32

	// GCCPUFractionは、プログラム開始以来のGCによって使用された利用可能なCPU時間の割合です。
	//
	// GCCPUFractionは0から1の数値で表され、0はこのプログラムのCPUの使用が全くないことを意味します。
	// プログラムの利用可能なCPU時間は、プログラム開始時からのGOMAXPROCSの積分と定義されます。
	// つまり、GOMAXPROCSが2で、プログラムが10秒間実行されている場合、その「利用可能なCPU時間」は20秒です。
	// GCCPUFractionには、書き込みバリアのアクティビティに使用されたCPU時間は含まれません。
	//
	// これは、GODEBUG=gctrace=1によって報告されるCPUの割合と同じです。
	GCCPUFraction float64

	// EnableGCはGCが有効であることを示します。常にtrueですが、
	// GOGC=offの場合でも有効です。
	EnableGC bool

	// DebugGC は現在使用されていません。
	DebugGC bool

	// BySizeは、サイズごとの割り当て統計を報告します。
	//
	// BySize[N]は、サイズSの割り当てに関する統計を提供します。ここで、BySize[N-1].Size < S ≤ BySize[N].Sizeです。
	//
	// これは、BySize[60].Sizeより大きい割り当てを報告しません。
	BySize [61]struct {
		Size uint32

		Mallocs uint64

		Frees uint64
	}
}

// ReadMemStatsはメモリアロケータ統計情報をmに書き込みます。
//
// 返されるメモリアロケータ統計情報はReadMemStatsの呼び出し時点で最新のものです。
// これは、ヒーププロファイルとは異なり、最新のガベージコレクションサイクルのスナップショットです。
func ReadMemStats(m *MemStats)
