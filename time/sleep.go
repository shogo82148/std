// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Sleepは現在のゴルーチンを少なくともdの期間だけ一時停止します。
// 負数またはゼロの期間は、Sleepがすぐに戻る原因となります。
func Sleep(d Duration)

// Timer型は単一のイベントを表します。
// Timerが時間切れになると、現在の時間がCに送信されます。
// ただし、TimerがAfterFuncによって作成された場合は除きます。
// TimerはNewTimerまたはAfterFuncで作成する必要があります。
type Timer struct {
	C <-chan Time
	r runtimeTimer
}

// Stop prevents the Timer from firing.
// It returns true if the call stops the timer, false if the timer has already
// expired or been stopped.
// Stop does not close the channel, to prevent a read from the channel succeeding
// incorrectly.
//
// To ensure the channel is empty after a call to Stop, check the
// return value and drain the channel.
// For example, assuming the program has not received from t.C already:
//
//	if !t.Stop() {
//		<-t.C
//	}
//
// This cannot be done concurrent to other receives from the Timer's
// channel or other calls to the Timer's Stop method.
//
// For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer
// has already expired and the function f has been started in its own goroutine;
// Stop does not wait for f to complete before returning.
// If the caller needs to know whether f is completed, it must coordinate
// with f explicitly.
func (t *Timer) Stop() bool

// NewTimerは、少なくともdの期間経過後に、
// そのチャネルに現在の時刻を送信する新しいTimerを作成します。
func NewTimer(d Duration) *Timer

// Resetはタイマーを期間d後に期限切れにする。
// タイマーがアクティブであった場合はtrue、期限切れまたは停止された場合はfalseを返す。
//
// NewTimerで作成されたタイマーの場合、Resetは停止または期限切れのタイマーにのみ呼び出すべきです。チャネルが空になっている必要があります。
//
// プログラムがすでにt.Cから値を受け取っている場合、タイマーは期限切れであり、チャネルは空になっているため、t.Resetを直接使用できます。
// しかし、まだt.Cから値を受け取っていない場合、タイマーは停止しなければならず、Stopがタイマーが停止される前に期限切れしたことを報告した場合は、
// チャネルを明示的に空にする必要があります：
//
//	if !t.Stop() {
//		<-t.C
//	}
//	t.Reset(d)
//
// これは、他の受信操作と同時に行うべきではありません。
//
// Resetの返り値を正しく使用することはできないことに注意してください。チャネルの空にする操作と新しいタイマーの期限切れとの競合状態があります。
// 上記に説明したように、Resetは必ず停止または期限切れのチャネルに対して呼び出すべきです。返り値は、既存のプログラムとの互換性を保つために存在します。
//
// AfterFunc(d, f)で作成されたタイマーの場合、Resetはfが実行されるタイミングを再スケジュールするかどうかによってtrueまたはfalseを返します。
// Resetがfalseを返す場合、Resetは前のfが完了するのを待たずにリターンせず、また後続のgoroutineが前のgoroutineと同時に実行されないことを保証しません。
// 前のfが完了したかどうかを知る必要がある場合、明示的にfと調整する必要があります。
func (t *Timer) Reset(d Duration) bool

// durationが経過するのを待って、その後で現在の時刻を返されたチャネルに送信します。
// これはNewTimer(d).Cと同等です。
// ガベージコレクタによってTimerが回収されるのは、タイマーが発火するまでではありません。
// 効率が問題の場合は、代わりにNewTimerを使用し、タイマーが不要になった場合はTimer.Stopを呼び出してください。
func After(d Duration) <-chan Time

// AfterFuncは、指定した時間が経過した後、fを自身のゴルーチンで呼び出します。
// Stopメソッドを使用して呼び出しをキャンセルするために使用できるTimerを返します。
// 返されたTimerのCフィールドは使用されず、nilになります。
func AfterFunc(d Duration, f func()) *Timer
