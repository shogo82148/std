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
// NewTickerは、各チック後にチャネルに現在の時刻を送信するチッカーを返します。チックの期間はduration引数で指定されます。チッカーは、遅い受信者のために時間間隔を調整するか、チックを削除します。 dの期間はゼロより大きくする必要があります。そうでない場合、NewTickerはパニックを起こします。チッカーを停止して関連するリソースを解放します。
=======
// NewTicker returns a new Ticker containing a channel that will send
// the current time on the channel after each tick. The period of the
// ticks is specified by the duration argument. The ticker will adjust
// the time interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will
// panic.
//
// Before Go 1.23, the garbage collector did not recover
// tickers that had not yet expired or been stopped, so code often
// immediately deferred t.Stop after calling NewTicker, to make
// the ticker recoverable when it was no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced
// tickers, even if they haven't been stopped.
// The Stop method is no longer necessary to help the garbage collector.
// (Code may of course still want to call Stop to stop the ticker for other reasons.)
>>>>>>> upstream/master
func NewTicker(d Duration) *Ticker

// Stopはティッカーを停止します。Stop後は、もうティックが送信されません。
// Stopはチャネルを閉じないため、チャネルから読み取り中の並行ゴルーチンが誤った「ティック」を見ることを防ぎます。
func (t *Ticker) Stop()

// Resetはタイマーを停止し、指定された期間でタイマーをリセットします。
// 新しい期間が経過すると、次のティックが到着します。期間dは0より大きくなければなりません。もしそうでない場合、Resetはパニックを引き起こします。
func (t *Ticker) Reset(d Duration)

<<<<<<< HEAD
// Tickは、ティッキングチャネルにのみアクセスを提供するためのNewTickerの便利なラッパーです。Tickerをシャットダウンする必要のないクライアントに便利ですが、シャットダウンの方法がないため、ガベージコレクタによって元のTickerは回収されません。"リーク"します。
// NewTickerとは異なり、d <= 0の場合にはnilが返されます。
=======
// Tick is a convenience wrapper for NewTicker providing access to the ticking
// channel only. Unlike NewTicker, Tick will return nil if d <= 0.
//
// Before Go 1.23, this documentation warned that the underlying
// Ticker would never be recovered by the garbage collector, and that
// if efficiency was a concern, code should use NewTicker instead and
// call Ticker.Stop when the ticker is no longer needed.
// As of Go 1.23, the garbage collector can recover unreferenced
// tickers, even if they haven't been stopped.
// The Stop method is no longer necessary to help the garbage collector.
// There is no longer any reason to prefer NewTicker when Tick will do.
>>>>>>> upstream/master
func Tick(d Duration) <-chan Time
