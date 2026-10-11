// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Sleepは現在のゴルーチンを少なくともdの期間だけ一時停止します。
// 負数またはゼロの期間は、Sleepがすぐに戻る原因となります。
func Sleep(d Duration)

// Timer型は単一のイベントを表します。
// Timerが時間切れになると、現在の時間がCに送信されます。
// ただし、Timerが [AfterFunc] によって作成された場合は除きます。
// Timerは [NewTimer] またはAfterFuncで作成する必要があります。
type Timer struct {
	C         <-chan Time
	initTimer bool
}

// Stopは[Timer]の発火を防ぎます。
// 呼び出しがタイマーを停止した場合はtrueを返し、タイマーがすでに期限切れまたは
// 停止されていた場合はfalseを返します。
//
// [AfterFunc](d, f)で作成された関数ベースのタイマーの場合、
// t.Stopがfalseを返したならば、タイマーはすでに期限切れになっており、
// 関数fは独自のゴルーチンで開始されています。
// Stopは戻る前にfの完了を待ちません。
// 呼び出し元がfが完了したかどうかを知る必要がある場合は、
// fと明示的に連携する必要があります。
//
// NewTimer(d)で作成されたチャネルベースのタイマーの場合、Go 1.23以降、
// Stopが戻った後のt.Cからの受信は、Stop以前の古い時刻値を受信するのではなく
// ブロックすることが保証されます。
// プログラムがまだt.Cから受信しておらず、タイマーが実行中の場合、
// Stopはtrueを返すことが保証されます。
// Go 1.23より前では、Stopを安全に使用する唯一の方法は、Stopがfalseを返した場合に
// 余分な<-t.Cを挿入して潜在的な古い値を排出することでした。
// 詳細については[NewTimer]のドキュメントを参照してください。
func (t *Timer) Stop() bool

// NewTimerは、少なくともdの期間経過後に、
// そのチャネルに現在の時刻を送信する新しいTimerを作成します。
//
// Before Go 1.23, the garbage collector did not recover
// timers that had not yet expired or been stopped, so code often
// immediately deferred t.Stop after calling NewTimer, to make
// the timer recoverable when it was no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced
// timers, even if they haven't expired or been stopped.
// The Stop method is no longer necessary to help the garbage collector.
// (Code may of course still want to call Stop to stop the timer for other reasons.)
//
// Before Go 1.23, the channel associated with a Timer was
// asynchronous (buffered, capacity 1), which meant that
// stale time values could be received even after [Timer.Stop]
// or [Timer.Reset] returned.
// As of Go 1.23, the channel is synchronous (unbuffered, capacity 0),
// eliminating the possibility of those stale values.
func NewTimer(d Duration) *Timer

// Resetはタイマーを期間d後に期限切れにする。
// タイマーがアクティブであった場合はtrue、期限切れまたは停止された場合はfalseを返す。
//
// For a func-based timer created with [AfterFunc](d, f), Reset either reschedules
// when f will run, in which case Reset returns true, or schedules f
// to run again, in which case it returns false.
// When Reset returns false, Reset neither waits for the prior f to
// complete before returning nor does it guarantee that the subsequent
// goroutine running f does not run concurrently with the prior
// one. If the caller needs to know whether the prior execution of
// f is completed, it must coordinate with f explicitly.
//
// For a chan-based timer created with NewTimer, as of Go 1.23,
// any receive from t.C after Reset has returned is guaranteed not
// to receive a time value corresponding to the previous timer settings;
// if the program has not received from t.C already and the timer is
// running, Reset is guaranteed to return true.
// Before Go 1.23, the only safe way to use Reset was to call [Timer.Stop]
// and explicitly drain the timer first.
// See the [NewTimer] documentation for more details.
func (t *Timer) Reset(d Duration) bool

// Afterは期間が経過するのを待ち、返されたチャネルに現在時刻を送信します。
// これは[NewTimer](d).Cと同等です。
//
// Go 1.23より前では、このドキュメントは、基礎となる[Timer]はタイマーが発火するまで
// ガベージコレクタによって回収されないこと、そして効率が問題になる場合は
// コードは代わりにNewTimerを使用し、タイマーが不要になったら[Timer.Stop]を
// 呼び出すべきであることを警告していました。
// Go 1.23以降、ガベージコレクタは参照されていない停止されていないタイマーを回収できます。
// Afterで十分な場合にNewTimerを優先する理由はありません。
func After(d Duration) <-chan Time

// AfterFuncは、指定した時間が経過した後、fを自身のゴルーチンで呼び出します。
// Stopメソッドを使用して呼び出しをキャンセルするために使用できる [Timer] を返します。
// 返されたTimerのCフィールドは使用されず、nilになります。
func AfterFunc(d Duration, f func()) *Timer
