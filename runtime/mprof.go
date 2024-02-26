// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// マロックのプロファイリング。
// tcmallocのアルゴリズムに基づいて作られたコードです。短いです。

package runtime

// SetBlockProfileRateは、ブロッキングイベントの割合を制御します。
// プロファイラは、ブロックされた時間がrateナノ秒ごとに平均1つのブロッキングイベントをサンプリングすることを目指しています。
//
// プロファイルにすべてのブロッキングイベントを含めるには、rate = 1を渡します。
// プロファイリングを完全にオフにするには、rate <= 0を渡します。
func SetBlockProfileRate(rate int)

// SetMutexProfileFractionは、mutexの衝突イベントのうち、
// プロファイルに報告される割合を制御します。平均して1/rateのイベントが報告されます。
// 以前のrateが返されます。
//
// プロファイリングを完全に無効にするには、rateに0を渡します。
// 現在のrateだけを読み取るには、rateに0より小さい値を渡します。
// (n>1の場合、サンプリングの詳細が変更される場合があります。)
func SetMutexProfileFraction(rate int) int

// StackRecordは単一の実行スタックを説明します。
type StackRecord struct {
	Stack0 [32]uintptr
}

// Stackは、レコードに関連付けられたスタックトレースを返します。
// これはr.Stack0のプレフィックスです。
func (r *StackRecord) Stack() []uintptr

// MemProfileRateは、メモリプロファイルに記録および報告されるメモリ割り当ての割合を制御します。
// プロファイラは、MemProfileRateバイトあたり平均1回の割り当てをサンプリングすることを目指しています。
// プロファイルにすべての割り当てブロックを含めるには、MemProfileRateを1に設定します。
// プロファイリングを完全にオフにするには、MemProfileRateを0に設定します。
// メモリプロファイルを処理するツールは、プロファイルの割合がプログラムの生存期間全体で一定であり、現在の値と等しいと想定しています。
// メモリプロファイリングの割合を変更するプログラムは、できるだけ早く（たとえば、mainの開始時などに）1度だけ行う必要があります。
var MemProfileRate int = 512 * 1024

// MemProfileRecordは、特定の呼び出しシーケンス（スタックトレース）によって割り当てられた生きているオブジェクトを記述します。
type MemProfileRecord struct {
	AllocBytes, FreeBytes     int64
	AllocObjects, FreeObjects int64
	Stack0                    [32]uintptr
}

// InUseBytesは使用中のバイト数（AllocBytes - FreeBytes）を返します。
func (r *MemProfileRecord) InUseBytes() int64

// InUseObjectsは使用中のオブジェクトの数を返します（AllocObjects - FreeObjects）。
func (r *MemProfileRecord) InUseObjects() int64

// Stackは、レコードに関連付けられたスタックトレースを返します。
// r.Stack0のプレフィックスです。
func (r *MemProfileRecord) Stack() []uintptr

// MemProfileは、割り当てられたメモリと解放されたメモリのプロファイルを、割り当ての場所別に返します。
// MemProfileは、現在のメモリプロファイルのレコード数であるnを返します。
// もしlen(p) >= nであれば、MemProfileはプロファイルをpにコピーし、nとtrueを返します。
// もしlen(p) < nであれば、MemProfileはpを変更せずに、nとfalseを返します。
// inuseZeroがtrueの場合、プロファイルにはr.AllocBytes > 0 かつ r.AllocBytes == r.FreeBytesのアロケーションレコードが含まれます。
// これは、メモリが割り当てられたがランタイムにすべて解放された場所です。
// 返されるプロファイルは、最大で2つのガベージコレクションサイクル前のものです。
// これは、プロファイルがアロケーションに偏った結果にならないようにするためです。
// アロケーションはリアルタイムで発生しますが、解放はガベージコレクタがスイーピングを実行するまで遅延されるため、
// プロファイルはガベージコレクタによって解放されるチャンスを持ったアロケーションのみをカウントします。
// 多くのクライアントは、runtime/pprofパッケージまたはtestingパッケージの-test.memprofileフラグを直接呼び出す代わりに使用するべきです。
func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool)

// BlockProfileRecordは、特定の呼び出しシーケンス（スタックトレース）で発生したブロッキングイベントを記述します。
type BlockProfileRecord struct {
	Count  int64
	Cycles int64
	StackRecord
}

// BlockProfileは現在のブロッキングプロファイルのレコード数nを返します。
// もしlen(p) >= nの場合、BlockProfileはプロファイルをpにコピーし、nとtrueを返します。
// もしlen(p) < nの場合、BlockProfileはpを変更せずに、nとfalseを返します。
//
<<<<<<< HEAD
// ほとんどのクライアントは、 [runtime/pprof] パッケージや
// [testing] パッケージの-test.blockprofileフラグを使用して、
// BlockProfileを直接呼び出す代わりに使用すべきです。
=======
// Most clients should use the [runtime/pprof] package or
// the [testing] package's -test.blockprofile flag instead
// of calling BlockProfile directly.
>>>>>>> upstream/master
func BlockProfile(p []BlockProfileRecord) (n int, ok bool)

// MutexProfileは現在のmutexプロファイルのレコード数であるnを返します。
// もしlen(p) >= nならば、MutexProfileはプロファイルをpにコピーしてnとtrueを返します。
// そうでなければ、MutexProfileはpを変更せずにnとfalseを返します。
//
// ほとんどのクライアントは、MutexProfileを直接呼び出す代わりに [runtime/pprof] パッケージを使用するべきです。
func MutexProfile(p []BlockProfileRecord) (n int, ok bool)

// ThreadCreateProfileはスレッド作成プロファイル内のレコード数であるnを返します。
// もし、len(p) >= nならば、ThreadCreateProfileはプロファイルをpにコピーしてn, trueを返します。
// もし、len(p) < nならば、ThreadCreateProfileはpを変更せずにn, falseを返します。
//
// 大抵のクライアントは直接ThreadCreateProfileを呼び出す代わりに、runtime/pprofパッケージを使用すべきです。
func ThreadCreateProfile(p []StackRecord) (n int, ok bool)

// GoroutineProfileはアクティブなゴルーチンスタックプロファイルのレコード数であるnを返します。
// もしlen(p)がn以上であれば、GoroutineProfileはプロファイルをpにコピーしnとtrueを返します。
// もしlen(p)がn未満であれば、GoroutineProfileはpを変更せずにnとfalseを返します。
//
// ほとんどのクライアントは直接GoroutineProfileを呼び出す代わりに [runtime/pprof] パッケージを使用するべきです。
func GoroutineProfile(p []StackRecord) (n int, ok bool)

// Stackは呼び出し元のゴルーチンのスタックトレースをbufに書き込み、
// bufに書き込まれたバイト数を返します。
// allがtrueの場合、現在のゴルーチンのトレースの後に、
// 他のすべてのゴルーチンのスタックトレースをbufに書き込みます。
func Stack(buf []byte, all bool) int
