// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Tickerは、時計の間隔で "ticks" を配信するチャネルを保持します。
type Ticker struct {
	C          <-chan Time
	initTicker bool
}

<<<<<<< HEAD
// NewTickerは、各ティック後にチャネル上の現在の時間を送信するチャネルを含む新しいTickerを返します。
// ティックの周期は、duration引数で指定されます。Tickerは、受信速度が遅い場合に時間間隔を調整したり、
// ティックをドロップしたりします。duration dはゼロより大きくなければならず、そうでない場合、
// NewTickerはパニックを起こします。
=======
// NewTicker returns a new [Ticker] containing a channel that will send
// the current time on the channel after each tick. The period of the
// ticks is specified by the duration argument. The ticker will adjust
// the time interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will
// panic.
>>>>>>> upstream/master
//
// Go 1.23より前では、ガベージコレクタはまだ期限切れになっていないか停止していない
// tickerを回収しなかったため、コードはしばしばNewTickerを呼び出した直後にt.Stopを即時に遅延させ、
// tickerが不要になったときに回収可能にしました。
// Go 1.23以降では、ガベージコレクタは参照されていないtickerを回収できます、
// たとえそれらが停止していなくても。
// Stopメソッドはもはやガベージコレクタを助けるためには必要ありません。
// （もちろん、コードは他の理由でtickerを停止させるためにStopを呼び出すことを望むかもしれません。）
func NewTicker(d Duration) *Ticker

// Stopはティッカーを停止します。Stop後は、もうティックが送信されません。
// Stopはチャネルを閉じないため、チャネルから読み取り中の並行ゴルーチンが誤った「ティック」を見ることを防ぎます。
func (t *Ticker) Stop()

// Resetはタイマーを停止し、指定された期間でタイマーをリセットします。
// 新しい期間が経過すると、次のティックが到着します。期間dは0より大きくなければなりません。もしそうでない場合、Resetはパニックを引き起こします。
func (t *Ticker) Reset(d Duration)

// Tick is a convenience wrapper for [NewTicker] providing access to the ticking
// channel only. Unlike NewTicker, Tick will return nil if d <= 0.
//
// Before Go 1.23, this documentation warned that the underlying
// [Ticker] would never be recovered by the garbage collector, and that
// if efficiency was a concern, code should use NewTicker instead and
// call [Ticker.Stop] when the ticker is no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced
// tickers, even if they haven't been stopped.
// The Stop method is no longer necessary to help the garbage collector.
// There is no longer any reason to prefer NewTicker when Tick will do.
func Tick(d Duration) <-chan Time
